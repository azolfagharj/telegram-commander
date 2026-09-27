package bot

import (
	"context"
	"errors"
	"time"

	tgbot "github.com/go-telegram/bot"
)

// retryClock abstracts time so tests can run without a real sleep.
type retryClock struct {
	now   func() time.Time
	sleep func(ctx context.Context, d time.Duration)
}

// defaultRetryClock is the real wall-clock implementation used in production.
var defaultRetryClock = retryClock{
	now: time.Now,
	sleep: func(ctx context.Context, d time.Duration) {
		select {
		case <-time.After(d):
		case <-ctx.Done():
		}
	},
}

// isPermanentSendError reports whether a Telegram send error should never be
// retried. Everything else (network errors, 5xx, TooManyRequestsError, ...)
// is treated as transient.
func isPermanentSendError(err error) bool {
	return errors.Is(err, tgbot.ErrorForbidden) || errors.Is(err, tgbot.ErrorBadRequest)
}

// sendWithRetry runs send, retrying transient failures with exponential
// backoff (starting at initialBackoff, capped at maxBackoff) until it
// succeeds, hits a permanent error, or ttl has elapsed since the first
// failure. The ttl is a lazy check (no ticker), mirroring the confirmWait
// pattern elsewhere in this package. On a TooManyRequestsError, the wait
// honors the error's RetryAfter instead of the computed backoff.
//
// If send never succeeds, sendWithRetry returns the last error so the caller
// can log it; this helper does no logging itself.
func sendWithRetry(ctx context.Context, send func() error, initialBackoff, maxBackoff, ttl time.Duration, clock retryClock) error {
	backoff := initialBackoff
	var deadline time.Time

	err := send()
	for err != nil {
		if isPermanentSendError(err) {
			return err
		}
		if deadline.IsZero() {
			deadline = clock.now().Add(ttl)
		}
		if !clock.now().Before(deadline) {
			return err
		}

		wait := backoff
		var rle *tgbot.TooManyRequestsError
		if errors.As(err, &rle) {
			wait = time.Duration(rle.RetryAfter) * time.Second
		}
		clock.sleep(ctx, wait)

		if backoff < maxBackoff {
			backoff *= 2
			if backoff > maxBackoff {
				backoff = maxBackoff
			}
		}

		err = send()
	}
	return nil
}
