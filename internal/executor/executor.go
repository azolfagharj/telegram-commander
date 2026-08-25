// Package executor runs shell commands with timeout, output limits, and per-user queues.
package executor

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"sync"
	"time"
)

// Spec describes a single command execution request.
type Spec struct {
	UserID   int64
	ButtonID string
	Button   string
	Command  string
	Shell    string
	WorkDir  string
	Env      map[string]string
	Timeout  time.Duration
	MaxBytes int
}

// Result holds the outcome of a command run.
type Result struct {
	Stdout    string
	Stderr    string
	ExitCode  int
	Duration  time.Duration
	TimedOut  bool
	Truncated bool
}

// Executor runs commands.
type Executor interface {
	Run(ctx context.Context, spec Spec) (Result, error)
}

// ShellExecutor executes commands via a shell with per-user serialization.
type ShellExecutor struct {
	audit *slog.Logger

	mu     sync.Mutex
	queues map[int64]chan struct{}
}

// NewShellExecutor creates a ShellExecutor. audit may be nil.
func NewShellExecutor(audit *slog.Logger) *ShellExecutor {
	if audit == nil {
		audit = slog.Default()
	}
	return &ShellExecutor{
		audit:  audit,
		queues: make(map[int64]chan struct{}),
	}
}

func (e *ShellExecutor) acquire(userID int64) func() {
	e.mu.Lock()
	ch, ok := e.queues[userID]
	if !ok {
		ch = make(chan struct{}, 1)
		ch <- struct{}{}
		e.queues[userID] = ch
	}
	e.mu.Unlock()
	<-ch
	return func() { ch <- struct{}{} }
}

// Run executes the command for the given user, serializing per user ID.
func (e *ShellExecutor) Run(ctx context.Context, spec Spec) (Result, error) {
	release := e.acquire(spec.UserID)
	defer release()

	start := time.Now()
	res, err := e.runOnce(ctx, spec)
	res.Duration = time.Since(start)

	e.audit.Info("command executed",
		"user_id", spec.UserID,
		"button", spec.Button,
		"button_id", spec.ButtonID,
		"command", spec.Command,
		"exit_code", res.ExitCode,
		"duration_ms", res.Duration.Milliseconds(),
		"timed_out", res.TimedOut,
		"truncated", res.Truncated,
		"error", errString(err),
	)
	return res, err
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func (e *ShellExecutor) runOnce(ctx context.Context, spec Spec) (Result, error) {
	var res Result
	shell := spec.Shell
	if shell == "" {
		shell = "/bin/bash"
	}
	timeout := spec.Timeout
	if timeout <= 0 {
		timeout = 60 * time.Second
	}
	maxBytes := spec.MaxBytes
	if maxBytes <= 0 {
		maxBytes = 512 * 1024
	}

	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(runCtx, shell, "-c", spec.Command)
	if spec.WorkDir != "" {
		cmd.Dir = spec.WorkDir
	}
	cmd.Env = os.Environ()
	for k, v := range spec.Env {
		cmd.Env = append(cmd.Env, k+"="+v)
	}

	var stdoutBuf, stderrBuf limitedBuffer
	stdoutBuf.limit = maxBytes
	stderrBuf.limit = maxBytes
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()
	res.Stdout = stdoutBuf.String()
	res.Stderr = stderrBuf.String()
	res.Truncated = stdoutBuf.truncated || stderrBuf.truncated

	if runCtx.Err() == context.DeadlineExceeded {
		res.TimedOut = true
		res.ExitCode = -1
		return res, fmt.Errorf("command timed out after %s", timeout)
	}
	if err != nil {
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			res.ExitCode = ee.ExitCode()
			return res, nil
		}
		res.ExitCode = -1
		return res, err
	}
	res.ExitCode = 0
	return res, nil
}

type limitedBuffer struct {
	buf       bytes.Buffer
	limit     int
	truncated bool
}

func (b *limitedBuffer) Write(p []byte) (int, error) {
	remaining := b.limit - b.buf.Len()
	if remaining <= 0 {
		b.truncated = true
		return len(p), nil
	}
	if len(p) > remaining {
		_, _ = b.buf.Write(p[:remaining])
		b.truncated = true
		return len(p), nil
	}
	return b.buf.Write(p)
}

func (b *limitedBuffer) String() string {
	return b.buf.String()
}

// FakeExecutor is a test double.
type FakeExecutor struct {
	Fn func(ctx context.Context, spec Spec) (Result, error)
}

// Run calls the configured function.
func (f *FakeExecutor) Run(ctx context.Context, spec Spec) (Result, error) {
	if f.Fn == nil {
		return Result{ExitCode: 0}, nil
	}
	return f.Fn(ctx, spec)
}

// Ensure limitedBuffer implements io.Writer.
var _ io.Writer = (*limitedBuffer)(nil)
