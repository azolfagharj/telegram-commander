package bot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram/bot/models"
)

// KeyboardFor builds an inline keyboard for a category (or root when nodeID is empty).
func (idx *Index) KeyboardFor(nodeID string, page int) (*models.InlineKeyboardMarkup, string, error) {
	var children []string
	var title string
	var columns int
	var parentID string

	if nodeID == "" {
		children = idx.Roots
		title = "Menu"
		columns = idx.DefaultColumns
	} else {
		n, ok := idx.ByID[nodeID]
		if !ok {
			return nil, "", fmt.Errorf("unknown node %q", nodeID)
		}
		if n.Type != "category" {
			return nil, "", fmt.Errorf("node %q is not a category", nodeID)
		}
		children = n.Children
		title = n.Label()
		columns = n.Columns
		parentID = n.ParentID
		if columns < 1 {
			columns = idx.DefaultColumns
		}
	}

	pageSize := idx.PageSize
	if pageSize < 1 {
		pageSize = 8
	}
	total := len(children)
	totalPages := 1
	if total > pageSize {
		totalPages = (total + pageSize - 1) / pageSize
	}
	if page < 0 {
		page = 0
	}
	if page >= totalPages {
		page = totalPages - 1
	}

	start := page * pageSize
	end := start + pageSize
	if end > total {
		end = total
	}
	slice := children[start:end]

	rows := chunkButtons(idx, slice, columns)

	navRow := make([]models.InlineKeyboardButton, 0, 4)
	if page > 0 {
		navRow = append(navRow, models.InlineKeyboardButton{
			Text:         "◀️ Prev",
			CallbackData: cbPrefixPage + nodeID + ":" + strconv.Itoa(page-1),
		})
	}
	if page < totalPages-1 {
		navRow = append(navRow, models.InlineKeyboardButton{
			Text:         "Next ▶️",
			CallbackData: cbPrefixPage + nodeID + ":" + strconv.Itoa(page+1),
		})
	}
	if len(navRow) > 0 {
		rows = append(rows, navRow)
	}

	trail := make([]models.InlineKeyboardButton, 0, 2)
	if nodeID != "" {
		if parentID != "" {
			trail = append(trail, models.InlineKeyboardButton{
				Text:         "⬅️ Back",
				CallbackData: cbPrefixNav + parentID,
			})
		}
		trail = append(trail, models.InlineKeyboardButton{
			Text:         "🏠 Home",
			CallbackData: cbHome,
		})
	}
	if len(trail) > 0 {
		rows = append(rows, trail)
	}

	if totalPages > 1 {
		title = fmt.Sprintf("%s (%d/%d)", title, page+1, totalPages)
	}

	return &models.InlineKeyboardMarkup{InlineKeyboard: rows}, title, nil
}

func chunkButtons(idx *Index, ids []string, columns int) [][]models.InlineKeyboardButton {
	if columns < 1 {
		columns = 2
	}
	var rows [][]models.InlineKeyboardButton
	var row []models.InlineKeyboardButton
	for _, id := range ids {
		n := idx.ByID[id]
		if n == nil {
			continue
		}
		var data string
		if n.Type == "category" {
			data = cbPrefixNav + n.ID
		} else {
			data = cbPrefixRun + n.ID
		}
		row = append(row, models.InlineKeyboardButton{
			Text:         n.Label(),
			CallbackData: data,
		})
		if len(row) >= columns {
			rows = append(rows, row)
			row = nil
		}
	}
	if len(row) > 0 {
		rows = append(rows, row)
	}
	return rows
}

// ConfirmKeyboard builds yes/cancel buttons for a confirmable action.
func ConfirmKeyboard(buttonID string) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "✅ Yes", CallbackData: cbPrefixConfirm + buttonID},
				{Text: "❌ Cancel", CallbackData: cbPrefixCancel + buttonID},
			},
		},
	}
}

// ParseCallback splits a callback data string into kind and payload.
func ParseCallback(data string) (kind, payload string) {
	switch {
	case data == cbHome:
		return "home", ""
	case strings.HasPrefix(data, cbPrefixNav):
		return "nav", strings.TrimPrefix(data, cbPrefixNav)
	case strings.HasPrefix(data, cbPrefixPage):
		return "page", strings.TrimPrefix(data, cbPrefixPage)
	case strings.HasPrefix(data, cbPrefixRun):
		return "run", strings.TrimPrefix(data, cbPrefixRun)
	case strings.HasPrefix(data, cbPrefixConfirm):
		return "confirm", strings.TrimPrefix(data, cbPrefixConfirm)
	case strings.HasPrefix(data, cbPrefixCancel):
		return "cancel", strings.TrimPrefix(data, cbPrefixCancel)
	default:
		return "unknown", data
	}
}
