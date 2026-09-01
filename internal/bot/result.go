package bot

import (
	"fmt"
	"html"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/azolfagharj/telegram-commander/internal/executor"
)

const (
	telegramMaxMessageLen = 4096
	maxResultChunks       = 10
)

// buildResultHeader is the summary lines above stdout/stderr.
func buildResultHeader(node *Node, res executor.Result, err error) string {
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
	return body.String()
}

// formatCommandResult builds HTML messages for a command run.
// Short output is one message. Longer output is split so every piece
// is within Telegram's limit and every <pre> tag is closed.
func formatCommandResult(node *Node, res executor.Result, err error) []string {
	return splitResultChunks(buildResultHeader(node, res, err), res.Stdout, res.Stderr)
}

// htmlCodeBlock wraps s in a <pre> tag with HTML-escaped content.
// Telegram renders <pre> as a monospace code block identical to ``` in Markdown.
func htmlCodeBlock(s string) string {
	s = strings.TrimRight(s, "\n")
	return "<pre>" + html.EscapeString(s) + "</pre>"
}

func tooLongNote(n int) string {
	return fmt.Sprintf("\n(output too long; showing first %d bytes)", n)
}

func sectionLabel(name string, continued, curNonEmpty bool) string {
	s := fmt.Sprintf("--- %s ---\n", name)
	if continued {
		s = fmt.Sprintf("--- %s (continued) ---\n", name)
	}
	if curNonEmpty {
		return "\n" + s
	}
	return s
}

func splitResultChunks(header, stdout, stderr string) []string {
	type sec struct {
		name string
		text string
	}
	secs := []sec{{name: "stdout", text: stdout}}
	if stderr != "" {
		secs = append(secs, sec{name: "stderr", text: stderr})
	}

	var chunks []string
	cur := header
	shown := 0
	leftover := false

	flush := func() {
		if cur == "" {
			return
		}
		chunks = append(chunks, cur)
		cur = ""
	}

	maxForCurrent := func() int {
		m := telegramMaxMessageLen
		if len(chunks) >= maxResultChunks-1 {
			m -= len(tooLongNote(len(stdout) + len(stderr)))
		}
		return m
	}

	for _, s := range secs {
		if leftover {
			break
		}
		raw := strings.TrimRight(s.text, "\n")
		var lines []string
		switch {
		case s.name == "stdout" && raw == "":
			lines = []string{""}
		case raw == "":
			continue
		default:
			lines = strings.Split(raw, "\n")
		}

		continued := false
		i := 0
		for i < len(lines) {
			if len(chunks) >= maxResultChunks {
				leftover = true
				break
			}

			label := sectionLabel(s.name, continued, cur != "")
			limit := maxForCurrent()
			emptyPiece := label + htmlCodeBlock("")
			if cur != "" && len(cur)+len(emptyPiece) > limit {
				flush()
				continued = true
				continue
			}

			gathered := make([]string, 0, 8)
			for i < len(lines) {
				try := make([]string, len(gathered)+1)
				copy(try, gathered)
				try[len(gathered)] = lines[i]
				piece := label + htmlCodeBlock(strings.Join(try, "\n"))
				if len(cur)+len(piece) <= limit {
					gathered = try
					i++
					continue
				}
				break
			}

			if len(gathered) > 0 {
				body := strings.Join(gathered, "\n")
				cur += label + htmlCodeBlock(body)
				shown += len(body)
				continued = true
				if i < len(lines) {
					flush()
				}
				continue
			}

			line := lines[i]
			avail := limit - len(cur) - len(label) - len("<pre>") - len("</pre>")
			part, rest := splitRawPrefixToFit(line, avail)
			if part == "" {
				if cur != "" {
					flush()
					continued = true
					continue
				}
				_, size := utf8.DecodeRuneInString(line)
				if size == 0 {
					i++
					continue
				}
				part, rest = line[:size], line[size:]
			}
			cur += label + htmlCodeBlock(part)
			shown += len(part)
			flush()
			continued = true
			if rest != "" {
				lines[i] = rest
			} else {
				i++
			}
			if len(chunks) >= maxResultChunks && (rest != "" || i < len(lines)) {
				leftover = true
				break
			}
		}
	}

	if cur != "" {
		if len(chunks) < maxResultChunks {
			flush()
		} else {
			leftover = true
		}
	}
	if len(chunks) == 0 {
		chunks = []string{header}
	}
	if leftover {
		note := tooLongNote(shown)
		last := chunks[len(chunks)-1]
		if len(last)+len(note) <= telegramMaxMessageLen {
			chunks[len(chunks)-1] = last + note
		}
	}
	return chunks
}

// splitRawPrefixToFit takes the longest UTF-8 prefix of s whose HTML-escaped
// form is at most avail bytes.
func splitRawPrefixToFit(s string, avail int) (part, rest string) {
	if avail <= 0 || s == "" {
		return "", s
	}
	used := 0
	n := 0
	for n < len(s) {
		r, size := utf8.DecodeRuneInString(s[n:])
		esc := len(html.EscapeString(string(r)))
		if used+esc > avail {
			break
		}
		used += esc
		n += size
	}
	return s[:n], s[n:]
}
