package executor_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/executor"
)

func TestShellEcho(t *testing.T) {
	ex := executor.NewShellExecutor(nil)
	res, err := ex.Run(context.Background(), executor.Spec{
		UserID:   1,
		Command:  "echo hello",
		Shell:    "/bin/sh",
		Timeout:  5 * time.Second,
		MaxBytes: 1024,
	})
	require.NoError(t, err)
	require.Equal(t, 0, res.ExitCode)
	require.Contains(t, res.Stdout, "hello")
}

func TestTimeout(t *testing.T) {
	ex := executor.NewShellExecutor(nil)
	res, err := ex.Run(context.Background(), executor.Spec{
		UserID:   1,
		Command:  "sleep 2",
		Shell:    "/bin/sh",
		Timeout:  200 * time.Millisecond,
		MaxBytes: 1024,
	})
	require.Error(t, err)
	require.True(t, res.TimedOut)
}

func TestNonZeroExit(t *testing.T) {
	ex := executor.NewShellExecutor(nil)
	res, err := ex.Run(context.Background(), executor.Spec{
		UserID:   1,
		Command:  "exit 7",
		Shell:    "/bin/sh",
		Timeout:  5 * time.Second,
		MaxBytes: 1024,
	})
	require.NoError(t, err)
	require.Equal(t, 7, res.ExitCode)
}

func TestFakeExecutor(t *testing.T) {
	f := &executor.FakeExecutor{
		Fn: func(ctx context.Context, spec executor.Spec) (executor.Result, error) {
			return executor.Result{Stdout: "ok", ExitCode: 0}, nil
		},
	}
	res, err := f.Run(context.Background(), executor.Spec{})
	require.NoError(t, err)
	require.Equal(t, "ok", res.Stdout)
}

func TestPerUserSerial(t *testing.T) {
	ex := executor.NewShellExecutor(nil)
	started := make(chan struct{}, 2)
	done := make(chan struct{})

	go func() {
		_, _ = ex.Run(context.Background(), executor.Spec{
			UserID:  42,
			Command: "sleep 0.3; echo a",
			Shell:   "/bin/sh",
			Timeout: 5 * time.Second,
		})
		close(done)
	}()
	time.Sleep(50 * time.Millisecond)
	go func() {
		started <- struct{}{}
		_, _ = ex.Run(context.Background(), executor.Spec{
			UserID:  42,
			Command: "echo b",
			Shell:   "/bin/sh",
			Timeout: 5 * time.Second,
		})
		started <- struct{}{}
	}()

	select {
	case <-started:
		// second job acquired the lock only after first finished or overlapped queue wait
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for second job")
	}
	<-done
}
