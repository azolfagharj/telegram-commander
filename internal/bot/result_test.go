package bot

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

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
	require.LessOrEqual(t, len(chunks), maxResultChunks)
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
