package bot_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/azolfagharj/telegram-commander/internal/bot"
	"github.com/azolfagharj/telegram-commander/internal/config"
)

func sampleButtons() []config.ButtonNode {
	return []config.ButtonNode{
		{
			Name: "Cat",
			Type: "category",
			Icon: "📁",
			Items: []config.ButtonNode{
				{Name: "Echo", Type: "button", Function: "command", Command: "echo hi"},
				{Name: "Fail", Type: "button", Function: "command", Command: "false", Confirm: true},
			},
		},
		{Name: "Top", Type: "button", Function: "command", Command: "echo top"},
	}
}

func TestBuildIndexAndKeyboard(t *testing.T) {
	idx := bot.BuildIndex(sampleButtons(), 2, 8)
	require.Len(t, idx.Roots, 2)
	kb, title, err := idx.KeyboardFor("", 0)
	require.NoError(t, err)
	require.Equal(t, "Menu", title)
	require.NotEmpty(t, kb.InlineKeyboard)

	rootCat := idx.ByID[idx.Roots[0]]
	require.Equal(t, "category", rootCat.Type)
	kb2, title2, err := idx.KeyboardFor(rootCat.ID, 0)
	require.NoError(t, err)
	require.Contains(t, title2, "Cat")
	require.NotEmpty(t, kb2.InlineKeyboard)
}

func TestParseCallback(t *testing.T) {
	kind, payload := bot.ParseCallback("home")
	require.Equal(t, "home", kind)
	kind, payload = bot.ParseCallback("n:abc")
	require.Equal(t, "nav", kind)
	require.Equal(t, "abc", payload)
	kind, payload = bot.ParseCallback("r:xyz")
	require.Equal(t, "run", kind)
	require.Equal(t, "xyz", payload)
}

func TestFindByName(t *testing.T) {
	idx := bot.BuildIndex(sampleButtons(), 2, 8)
	n := idx.FindByName("echo")
	require.NotNil(t, n)
	require.Equal(t, "Echo", n.Name)
}

func TestPagination(t *testing.T) {
	var items []config.ButtonNode
	for i := 0; i < 10; i++ {
		items = append(items, config.ButtonNode{
			Name:     string(rune('A' + i)),
			Type:     "button",
			Function: "command",
			Command:  "echo",
		})
	}
	buttons := []config.ButtonNode{{
		Name:  "Many",
		Type:  "category",
		Items: items,
	}}
	idx := bot.BuildIndex(buttons, 2, 4)
	catID := idx.Roots[0]
	kb, title, err := idx.KeyboardFor(catID, 0)
	require.NoError(t, err)
	require.Contains(t, title, "1/3")
	foundNext := false
	for _, row := range kb.InlineKeyboard {
		for _, b := range row {
			if b.Text == "Next ▶️" {
				foundNext = true
			}
		}
	}
	require.True(t, foundNext)
}

func TestConfirmKeyboard(t *testing.T) {
	kb := bot.ConfirmKeyboard("abc")
	require.Len(t, kb.InlineKeyboard, 1)
	require.Len(t, kb.InlineKeyboard[0], 2)
}
