package bot

import (
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/config"
	"github.com/azolfagharj/telegram-commander/internal/executor"
)

func TestFormatCommandResultHTMLCodeBlock(t *testing.T) {
	n := &Node{Name: "Uptime"}
	chunks := formatCommandResult(n, executor.Result{
		Stdout:   " 12:00:01 up 1 day\n",
		Stderr:   "warn",
		ExitCode: 0,
		Duration: 12 * time.Millisecond,
	}, nil)
	require.Len(t, chunks, 1)
	text := chunks[0]
	require.Contains(t, text, "--- stdout ---\n<pre> 12:00:01 up 1 day</pre>")
	require.Contains(t, text, "--- stderr ---\n<pre>warn</pre>")
	require.Contains(t, text, "Button: Uptime")
	require.LessOrEqual(t, len(text), telegramMaxMessageLen)
}

func TestFormatCommandResultEscapesHTML(t *testing.T) {
	n := &Node{Name: "<script>"}
	chunks := formatCommandResult(n, executor.Result{
		Stdout: "a < b && b > c",
	}, nil)
	require.Len(t, chunks, 1)
	text := chunks[0]
	require.Contains(t, text, "Button: &lt;script&gt;")
	require.Contains(t, text, "<pre>a &lt; b &amp;&amp; b &gt; c</pre>")
	require.NotContains(t, text, "<script>")
}

func TestHTMLCodeBlockShape(t *testing.T) {
	require.Equal(t, "<pre>hi</pre>", htmlCodeBlock("hi"))
	require.Equal(t, "<pre>hi</pre>", htmlCodeBlock("hi\n"))
}

func TestSplitResultChunksClosesPre(t *testing.T) {
	header := "Button: Ports\nExit: 0\nDuration: 1ms\n"
	line := `udp UNCONN 0 0 0.0.0.0:80 0.0.0.0:* users:(("python3",pid=1,fd=2))`
	var raw strings.Builder
	for i := 0; i < 80; i++ {
		raw.WriteString(line)
		raw.WriteByte('\n')
	}
	chunks := splitResultChunks(header, raw.String(), "")
	require.Greater(t, len(chunks), 1)
	joined := strings.Builder{}
	for i, c := range chunks {
		require.LessOrEqual(t, len(c), telegramMaxMessageLen, "chunk %d", i)
		require.Equal(t, strings.Count(c, "<pre>"), strings.Count(c, "</pre>"), "chunk %d tags", i)
		require.NotContains(t, c, "<pre></pre></pre>")
		joined.WriteString(c)
	}
	require.Contains(t, chunks[0], "Button: Ports")
	require.Contains(t, chunks[1], "(continued)")
}

func TestSplitResultChunksLongLine(t *testing.T) {
	header := "Button: X\nExit: 0\nDuration: 1ms\n"
	line := strings.Repeat("a", 20000)
	chunks := splitResultChunks(header, line, "")
	require.Greater(t, len(chunks), 1)
	var rebuilt strings.Builder
	for i, c := range chunks {
		require.LessOrEqual(t, len(c), telegramMaxMessageLen, "chunk %d", i)
		require.Equal(t, strings.Count(c, "<pre>"), strings.Count(c, "</pre>"), "chunk %d tags", i)
		start := strings.Index(c, "<pre>")
		end := strings.LastIndex(c, "</pre>")
		require.Greater(t, start, -1)
		require.Greater(t, end, start)
		rebuilt.WriteString(c[start+len("<pre>") : end])
	}
	require.Equal(t, line, rebuilt.String())
}

func TestSplitResultChunksMaxThenNote(t *testing.T) {
	header := "Button: X\nExit: 0\nDuration: 1ms\n"
	var raw strings.Builder
	for i := 0; i < 2000; i++ {
		raw.WriteString(strings.Repeat("z", 80))
		raw.WriteByte('\n')
	}
	chunks := splitResultChunks(header, raw.String(), "")
	require.LessOrEqual(t, len(chunks), config.MaxOutputMessagesCeiling)
	require.Contains(t, chunks[len(chunks)-1], "output too long")
	for i, c := range chunks {
		require.LessOrEqual(t, len(c), telegramMaxMessageLen, "chunk %d", i)
		require.Equal(t, strings.Count(c, "<pre>"), strings.Count(c, "</pre>"), "chunk %d tags", i)
	}
}

func TestSplitRawPrefixToFitRunes(t *testing.T) {
	part, rest := splitRawPrefixToFit("héllo", 3) // "h" + "é" is 1+2 escaped same, é is 2 bytes
	require.True(t, len(part) > 0)
	require.Equal(t, "héllo", part+rest)
	part, rest = splitRawPrefixToFit(`a<"`, 10)
	require.Equal(t, `a<"`, part+rest)
}


func TestBuildResultFileAndCaption(t *testing.T) {
	n := &Node{Name: "Nginx Logs"}
	res := executor.Result{
		Stdout:   "line1\nline2\n",
		Stderr:   "warn\n",
		ExitCode: 0,
		Duration: 15 * time.Millisecond,
	}
	name, body := buildResultFile(n, res, nil)
	require.True(t, strings.HasPrefix(name, "nginx-logs-"))
	require.True(t, strings.HasSuffix(name, ".txt"))
	text := string(body)
	require.Contains(t, text, "Button: Nginx Logs")
	require.Contains(t, text, "--- stdout ---")
	require.Contains(t, text, "line1")
	require.Contains(t, text, "--- stderr ---")
	require.Contains(t, text, "warn")
	require.NotContains(t, text, "<pre>")

	caption := resultCaption(n, res, nil)
	require.Contains(t, caption, "Button: Nginx Logs")
	require.LessOrEqual(t, utf8.RuneCountInString(caption), telegramMaxCaptionLen)
	require.NotContains(t, caption, "stdout")
}

func TestResultDeliveryModes(t *testing.T) {
	tests := []struct {
		mode        string
		chunks      int
		maxMessages int
		wantFile    bool
	}{
		{mode: "file", chunks: 1, maxMessages: 2, wantFile: true},
		{mode: "text", chunks: 9, maxMessages: 2, wantFile: false},
		{mode: "auto", chunks: 2, maxMessages: 2, wantFile: false},
		{mode: "auto", chunks: 3, maxMessages: 2, wantFile: true},
		{mode: "AUTO", chunks: 3, maxMessages: 2, wantFile: true},
		{mode: "", chunks: 3, maxMessages: 0, wantFile: true},
	}
	for _, tt := range tests {
		t.Run(tt.mode+"/"+string(rune('0'+tt.chunks)), func(t *testing.T) {
			require.Equal(t, tt.wantFile, resultDelivery(tt.mode, tt.chunks, tt.maxMessages))
		})
	}
}

func TestResultCaptionTruncates(t *testing.T) {
	n := &Node{Name: strings.Repeat("X", 2000)}
	caption := resultCaption(n, executor.Result{ExitCode: 1}, nil)
	require.LessOrEqual(t, utf8.RuneCountInString(caption), telegramMaxCaptionLen)
	require.Contains(t, caption, "Button:")
}

func TestResultFileNameFallback(t *testing.T) {
	when := time.Date(2026, 9, 13, 12, 0, 0, 0, time.UTC)
	require.Equal(t, "output-20260913-120000.txt", resultFileName("!!!", when))
	require.Equal(t, "disk-free-20260913-120000.txt", resultFileName("Disk Free", when))
}
