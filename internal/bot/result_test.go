package bot

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/executor"
)

func TestFormatCommandResultHTMLCodeBlock(t *testing.T) {
	n := &Node{Name: "Uptime"}
	text := formatCommandResult(n, executor.Result{
		Stdout:   " 12:00:01 up 1 day\n",
		Stderr:   "warn",
		ExitCode: 0,
		Duration: 12 * time.Millisecond,
	}, nil)
	require.Contains(t, text, "--- stdout ---\n<pre> 12:00:01 up 1 day</pre>")
	require.Contains(t, text, "--- stderr ---\n<pre>warn</pre>")
	require.Contains(t, text, "Button: Uptime")
}

func TestFormatCommandResultEscapesHTML(t *testing.T) {
	n := &Node{Name: "<script>"}
	text := formatCommandResult(n, executor.Result{
		Stdout: "a < b && b > c",
	}, nil)
	require.Contains(t, text, "Button: &lt;script&gt;")
	require.Contains(t, text, "<pre>a &lt; b &amp;&amp; b &gt; c</pre>")
	require.NotContains(t, text, "<script>")
}

func TestHTMLCodeBlockShape(t *testing.T) {
	require.Equal(t, "<pre>hi</pre>", htmlCodeBlock("hi"))
	require.Equal(t, "<pre>hi</pre>", htmlCodeBlock("hi\n"))
}
