package bot

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/executor"
)

func TestFormatCommandResultWrapsOutputInCodeFence(t *testing.T) {
	n := &Node{Name: "Uptime"}
	text := formatCommandResult(n, executor.Result{
		Stdout:   " 12:00:01 up 1 day\n",
		Stderr:   "warn",
		ExitCode: 0,
		Duration: 12 * time.Millisecond,
	}, nil)
	require.Contains(t, text, "--- stdout ---\n```\n 12:00:01 up 1 day\n```")
	require.Contains(t, text, "--- stderr ---\n```\nwarn\n```")
	require.Contains(t, text, "Button: Uptime")
}

func TestCodeFenceShape(t *testing.T) {
	require.Equal(t, "```\nhi\n```", codeFence("hi"))
	require.Equal(t, "```\nhi\n```", codeFence("hi\n"))
}
