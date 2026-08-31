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

func collectReply(kb *bot.MenuBuild) []string {
	var texts []string
	for _, row := range kb.Reply.Keyboard {
		for _, b := range row {
			texts = append(texts, b.Text)
		}
	}
	return texts
}

func collectInline(kb *bot.MenuBuild) []string {
	var texts []string
	for _, row := range kb.Inline.InlineKeyboard {
		for _, b := range row {
			texts = append(texts, b.Text)
		}
	}
	return texts
}

func TestBuildIndexAndKeyboard(t *testing.T) {
	idx := bot.BuildIndex(sampleButtons(), 2, 8)
	require.Len(t, idx.Roots, 2)
	view, err := idx.BuildMenu("", 0)
	require.NoError(t, err)
	require.Equal(t, "Menu", view.Title)

	texts := collectReply(view)
	require.Contains(t, texts, "🏠 Home")
	require.NotContains(t, texts, "⬅️ Back")
	require.NotContains(t, texts, "📁 Cat")
	require.NotContains(t, texts, "Top")

	inline := collectInline(view)
	require.Contains(t, inline, "🏠 Home")
	require.NotContains(t, inline, "⬅️ Back")
	require.Contains(t, inline, "📁 Cat")
	require.Contains(t, inline, "Top")

	rootCat := idx.ByID[idx.Roots[0]]
	require.Equal(t, "category", rootCat.Type)
	view2, err := idx.BuildMenu(rootCat.ID, 0)
	require.NoError(t, err)
	require.Contains(t, view2.Title, "Cat")

	texts2 := collectReply(view2)
	require.Contains(t, texts2, "🏠 Home")
	require.Contains(t, texts2, "⬅️ Back")
	require.NotContains(t, texts2, "Echo")

	inline2 := collectInline(view2)
	require.Contains(t, inline2, "🏠 Home")
	require.Contains(t, inline2, "⬅️ Back")
	require.Contains(t, inline2, "Echo")
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
	view, err := idx.BuildMenu(catID, 0)
	require.NoError(t, err)
	require.Contains(t, view.Title, "1/3")
	require.True(t, view.HasNext)
	require.False(t, view.HasPrev)
	require.True(t, view.HasBack)

	reply := collectReply(view)
	require.Contains(t, reply, "Next ▶️")
	require.NotContains(t, reply, "◀️ Prev")
	require.Contains(t, reply, "⬅️ Back")

	inline := collectInline(view)
	require.Contains(t, inline, "Next ▶️")
	require.NotContains(t, inline, "◀️ Prev")
	require.Contains(t, inline, "⬅️ Back")
	require.Contains(t, inline, "🏠 Home")

	viewLast, err := idx.BuildMenu(catID, 2)
	require.NoError(t, err)
	require.True(t, viewLast.HasPrev)
	require.False(t, viewLast.HasNext)
	require.Contains(t, collectReply(viewLast), "◀️ Prev")
	require.NotContains(t, collectReply(viewLast), "Next ▶️")
}

func TestConfirmInlineKeyboard(t *testing.T) {
	kb := bot.ConfirmInlineKeyboard(true)
	require.Equal(t, "🏠 Home", kb.InlineKeyboard[0][0].Text)
	require.Equal(t, "⬅️ Back", kb.InlineKeyboard[0][1].Text)
	require.Equal(t, "✅ Yes", kb.InlineKeyboard[1][0].Text)
	require.Equal(t, "❌ Cancel", kb.InlineKeyboard[1][1].Text)

	root := bot.ConfirmInlineKeyboard(false)
	require.Len(t, root.InlineKeyboard[0], 1)
}
