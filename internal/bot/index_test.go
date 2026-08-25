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
	require.NotEmpty(t, kb.Keyboard)

	var texts []string
	for _, row := range kb.Keyboard {
		for _, b := range row {
			texts = append(texts, b.Text)
		}
	}
	require.Contains(t, texts, "🏠 Home")
	require.NotContains(t, texts, "⬅️ Back")
	require.Contains(t, texts, "📁 Cat")
	require.Contains(t, texts, "Top")

	rootCat := idx.ByID[idx.Roots[0]]
	require.Equal(t, "category", rootCat.Type)
	kb2, title2, err := idx.KeyboardFor(rootCat.ID, 0)
	require.NoError(t, err)
	require.Contains(t, title2, "Cat")
	require.NotEmpty(t, kb2.Keyboard)

	var texts2 []string
	for _, row := range kb2.Keyboard {
		for _, b := range row {
			texts2 = append(texts2, b.Text)
		}
	}
	require.Contains(t, texts2, "🏠 Home")
	require.Contains(t, texts2, "⬅️ Back")
	require.Contains(t, texts2, "Echo")
}

func TestFindByName(t *testing.T) {
	idx := bot.BuildIndex(sampleButtons(), 2, 8)
	n := idx.FindByName("echo")
	require.NotNil(t, n)
	require.Equal(t, "Echo", n.Name)
}

func TestChildByLabel(t *testing.T) {
	idx := bot.BuildIndex(sampleButtons(), 2, 8)
	n := idx.ChildByLabel("", "📁 Cat")
	require.NotNil(t, n)
	require.Equal(t, "category", n.Type)
	echo := idx.ChildByLabel(n.ID, "Echo")
	require.NotNil(t, echo)
	require.Equal(t, "Echo", echo.Name)
	require.Nil(t, idx.ChildByLabel("", "Echo"))
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
	foundBack := false
	for _, row := range kb.Keyboard {
		for _, b := range row {
			if b.Text == "Next ▶️" {
				foundNext = true
			}
			if b.Text == "⬅️ Back" {
				foundBack = true
			}
		}
	}
	require.True(t, foundNext)
	require.True(t, foundBack)
}

func TestConfirmKeyboard(t *testing.T) {
	kb := bot.ConfirmKeyboard()
	require.Len(t, kb.Keyboard, 1)
	require.Len(t, kb.Keyboard[0], 2)
	require.Equal(t, "✅ Yes", kb.Keyboard[0][0].Text)
	require.Equal(t, "❌ Cancel", kb.Keyboard[0][1].Text)
}
