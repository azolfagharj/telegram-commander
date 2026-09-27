package bot

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/azolfagharj/telegram-commander/internal/config"
	"github.com/azolfagharj/telegram-commander/internal/executor"
	"github.com/azolfagharj/telegram-commander/internal/function"
)

func TestTelegramCallSerializesPerChat(t *testing.T) {
	cfg := &config.Config{}
	cfg.DeliveryRetryBackoff.Duration = time.Millisecond
	cfg.DeliveryRetryBackoffMax.Duration = time.Millisecond
	cfg.DeliveryRetryTTL.Duration = time.Second
	cfg.ApplyDefaults()

	a := NewApp(cfg, function.NewRegistry(), &executor.FakeExecutor{}, nil)

	var concurrent int32
	var maxConcurrent int32
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			_ = a.telegramCall(context.Background(), 42, func() error {
				n := atomic.AddInt32(&concurrent, 1)
				for {
					old := atomic.LoadInt32(&maxConcurrent)
					if n <= old || atomic.CompareAndSwapInt32(&maxConcurrent, old, n) {
						break
					}
				}
				time.Sleep(40 * time.Millisecond)
				atomic.AddInt32(&concurrent, -1)
				return nil
			})
		}()
	}
	wg.Wait()

	if maxConcurrent != 1 {
		t.Fatalf("expected max concurrent sends for one chat to be 1, got %d", maxConcurrent)
	}
}

func TestTelegramCallSkipsLockForChatZero(t *testing.T) {
	cfg := &config.Config{}
	cfg.DeliveryRetryBackoff.Duration = time.Millisecond
	cfg.DeliveryRetryBackoffMax.Duration = time.Millisecond
	cfg.DeliveryRetryTTL.Duration = time.Second
	cfg.ApplyDefaults()

	a := NewApp(cfg, function.NewRegistry(), &executor.FakeExecutor{}, nil)

	var concurrent int32
	var maxConcurrent int32
	started := make(chan struct{}, 2)
	release := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			_ = a.telegramCall(context.Background(), 0, func() error {
				n := atomic.AddInt32(&concurrent, 1)
				for {
					old := atomic.LoadInt32(&maxConcurrent)
					if n <= old || atomic.CompareAndSwapInt32(&maxConcurrent, old, n) {
						break
					}
				}
				started <- struct{}{}
				<-release
				atomic.AddInt32(&concurrent, -1)
				return nil
			})
		}()
	}

	// Both ops with chatID 0 should enter without waiting on each other.
	for i := 0; i < 2; i++ {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("chatID 0 should not serialize unrelated calls")
		}
	}
	close(release)
	wg.Wait()

	if maxConcurrent != 2 {
		t.Fatalf("expected concurrent chatID 0 calls, max=%d", maxConcurrent)
	}
}
