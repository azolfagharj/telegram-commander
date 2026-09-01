package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/azolfagharj/telegram-commander/internal/executor"
)

// formatCommandResult builds the Telegram text for a command run.
// Stdout and stderr are wrapped in a code fence as plain text:
// ```
// output
// ```
func formatCommandResult(node *Node, res executor.Result, err error) string {
	var body strings.Builder
	body.WriteString(fmt.Sprintf("Button: %s\n", node.Name))
	body.WriteString(fmt.Sprintf("Exit: %d\n", res.ExitCode))
	body.WriteString(fmt.Sprintf("Duration: %s\n", res.Duration.Round(time.Millisecond)))
	if res.TimedOut {
		body.WriteString("Status: TIMED OUT\n")
	}
	if err != nil && !res.TimedOut {
		body.WriteString("Error: " + err.Error() + "\n")
	}
	if res.Truncated {
		body.WriteString("(output truncated)\n")
	}
	body.WriteString("\n--- stdout ---\n")
	body.WriteString(codeFence(res.Stdout))
	if res.Stderr != "" {
		body.WriteString("\n--- stderr ---\n")
		body.WriteString(codeFence(res.Stderr))
	}
	text := body.String()
	if len(text) > telegramMaxMessageLen {
		text = text[:telegramMaxMessageLen]
	}
	return text
}

func codeFence(s string) string {
	s = strings.TrimRight(s, "\n")
	return "```\n" + s + "\n```"
}
