package bot

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/azolfagharj/telegram-commander/internal/executor"
)

// formatCommandResult builds an HTML message for a command run.
// Summary lines are plain escaped text; stdout and stderr are wrapped in
// <pre> tags so Telegram renders them as a monospace code block.
func formatCommandResult(node *Node, res executor.Result, err error) string {
	var body strings.Builder
	body.WriteString(html.EscapeString(fmt.Sprintf("Button: %s\n", node.Name)))
	body.WriteString(fmt.Sprintf("Exit: %d\n", res.ExitCode))
	body.WriteString(fmt.Sprintf("Duration: %s\n", res.Duration.Round(time.Millisecond)))
	if res.TimedOut {
		body.WriteString("Status: TIMED OUT\n")
	}
	if err != nil && !res.TimedOut {
		body.WriteString(html.EscapeString("Error: " + err.Error() + "\n"))
	}
	if res.Truncated {
		body.WriteString("(output truncated)\n")
	}
	body.WriteString("\n--- stdout ---\n")
	body.WriteString(htmlCodeBlock(res.Stdout))
	if res.Stderr != "" {
		body.WriteString("\n--- stderr ---\n")
		body.WriteString(htmlCodeBlock(res.Stderr))
	}
	text := body.String()
	if len(text) > telegramMaxMessageLen {
		text = text[:telegramMaxMessageLen]
	}
	return text
}

// htmlCodeBlock wraps s in a <pre> tag with HTML-escaped content.
// Telegram renders <pre> as a monospace code block identical to ``` in Markdown.
func htmlCodeBlock(s string) string {
	s = strings.TrimRight(s, "\n")
	return "<pre>" + html.EscapeString(s) + "</pre>"
}
