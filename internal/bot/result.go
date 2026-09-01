package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/azolfagharj/telegram-commander/internal/executor"
)

// formatCommandResult builds the Telegram text for a command run.
// Stdout and stderr are wrapped in Markdown code fences so they show as a
// code block in the chat.
func formatCommandResult(node *Node, res executor.Result, err error) string {
	var body strings.Builder
	body.WriteString(escapeMarkdown(fmt.Sprintf("Button: %s\n", node.Name)))
	body.WriteString(fmt.Sprintf("Exit: %d\n", res.ExitCode))
	body.WriteString(fmt.Sprintf("Duration: %s\n", res.Duration.Round(time.Millisecond)))
	if res.TimedOut {
		body.WriteString("Status: TIMED OUT\n")
	}
	if err != nil && !res.TimedOut {
		body.WriteString(escapeMarkdown("Error: " + err.Error() + "\n"))
	}
	if res.Truncated {
		body.WriteString("(output truncated)\n")
	}
	body.WriteString("\n--- stdout ---\n")
	body.WriteString(markdownCodeBlock(res.Stdout))
	if res.Stderr != "" {
		body.WriteString("\n--- stderr ---\n")
		body.WriteString(markdownCodeBlock(res.Stderr))
	}
	text := body.String()
	if len(text) > telegramMaxMessageLen {
		text = text[:telegramMaxMessageLen]
	}
	return text
}

func markdownCodeBlock(s string) string {
	// A triple backtick inside the output would close the block early.
	s = strings.ReplaceAll(s, "```", "'''")
	return "```\n" + s + "\n```"
}

func escapeMarkdown(s string) string {
	r := strings.NewReplacer("_", "\\_", "*", "\\*", "`", "\\`", "[", "\\[")
	return r.Replace(s)
}
