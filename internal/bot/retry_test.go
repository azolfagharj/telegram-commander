package bot

import (
	"context"
	"errors"
	"testing"
	"time"

	tgbot "github.com/go-telegram/bot"
)

// fakeClock is an instant, deterministic clock/sleep for tests: sleep just
// advances the fake "now" instead of blocking.
type fakeClock struct {
	t time.Time
}

func (f *fakeClock) clock() retryClock {
	return retryClock{
		now: func() time.Time { return f.t },
		sleep: func(_ context.Context, d time.Duration) {
			f.t = f.t.Add(d)
		},
	}
}

var errTransient = errors.New("temporary network blip")

func TestRetryPermanentErrorNoRetry(t *testing.T) {
	f := &fakeClock{t: time.Now()}
	calls := 0
	err := sendWithRetry(context.Background(), func() error {
		calls++
		return tgbot.ErrorBadRequest
	}, time.Second, 30*time.Second, 2*time.Minute, f.clock())

	if !errors.Is(err, tgbot.ErrorBadRequest) {
		t.Fatalf("expected ErrorBadRequest, got %v", err)
	}
	if calls != 1 {
		t.Fatalf("expected exactly 1 call for a permanent error, got %d", calls)
	}
}

func TestRetryTransientThenSucceeds(t *testing.T) {
	f := &fakeClock{t: time.Now()}
	calls := 0
	err := sendWithRetry(context.Background(), func() error {
		calls++
		if calls < 3 {
			return errTransient
		}
		return nil
	}, time.Second, 30*time.Second, 2*time.Minute, f.clock())

	if err != nil {
		t.Fatalf("expected eventual success, got %v", err)
	}
	if calls != 3 {
		t.Fatalf("expected 3 calls, got %d", calls)
	}
}

func TestRetryGivesUpAfterTTL(t *testing.T) {
	f := &fakeClock{t: time.Now()}
	calls := 0
	err := sendWithRetry(context.Background(), func() error {
		calls++
		return errTransient
	}, time.Second, 4*time.Second, 10*time.Second, f.clock())

	if !errors.Is(err, errTransient) {
		t.Fatalf("expected errTransient after giving up, got %v", err)
	}
	if calls < 2 {
		t.Fatalf("expected more than one attempt before giving up, got %d", calls)
	}
}

func TestRetryHonorsRetryAfterExactly(t *testing.T) {
	start := time.Now()
	f := &fakeClock{t: start}
	calls := 0
	// initialBackoff is 1s, but RetryAfter says 7s: the wait must be 7s, not 1s.
	err := sendWithRetry(context.Background(), func() error {
		calls++
		if calls == 1 {
			return &tgbot.TooManyRequestsError{Message: "slow down", RetryAfter: 7}
		}
		return nil
	}, time.Second, 30*time.Second, 2*time.Minute, f.clock())

	if err != nil {
		t.Fatalf("expected eventual success, got %v", err)
	}
	if calls != 2 {
		t.Fatalf("expected 2 calls, got %d", calls)
	}
	if got := f.t.Sub(start); got != 7*time.Second {
		t.Fatalf("expected the fake clock to advance by exactly 7s, got %v", got)
	}
}
