package bot

import (
	"fmt"

	"github.com/go-telegram/bot/models"
)

const (
	btnHome   = "🏠 Home"
	btnBack   = "🔙 Back"
	btnPrev   = "⏪ Prev"
	btnNext   = "Next ⏩"
	btnYes    = "✅ Yes"
	btnCancel = "❌ Cancel"

	cbHome   = "h"
	cbBack   = "b"
	cbPrev   = "p"
	cbNext   = "n"
	cbYes    = "y"
	cbCancel = "x"
	cbOpen   = "o:"
	cbRun    = "r:"
)

// MenuBuild is one menu screen: title text plus inline buttons.
type MenuBuild struct {
	Title      string
	Page       int
	TotalPages int
	Inline     *models.InlineKeyboardMarkup
	HasBack    bool
	HasPrev    bool
	HasNext    bool
}

// BuildMenu builds the inline keyboard for a category (or root when nodeID is empty).
func (idx *Index) BuildMenu(nodeID string, page int) (*MenuBuild, error) {
	children, title, columns, err := idx.menuChildren(nodeID)
	if err != nil {
		return nil, err
	}

	page, totalPages, slice := paginateIDs(children, idx.pageSize(), page)
	hasBack := nodeID != ""
	hasPrev := page > 0
	hasNext := page < totalPages-1

	if totalPages > 1 {
		title = fmt.Sprintf("%s (%d/%d)", title, page+1, totalPages)
	}

	return &MenuBuild{
		Title:      title,
		Page:       page,
		TotalPages: totalPages,
		Inline:     menuInlineKeyboard(idx, slice, columns, hasBack, hasPrev, hasNext),
		HasBack:    hasBack,
		HasPrev:    hasPrev,
		HasNext:    hasNext,
	}, nil
}

func navRow(hasBack bool) []models.InlineKeyboardButton {
	nav := []models.InlineKeyboardButton{{Text: btnHome, CallbackData: cbHome}}
	if hasBack {
		nav = append(nav, models.InlineKeyboardButton{Text: btnBack, CallbackData: cbBack})
	}
	return nav
}

func menuInlineKeyboard(idx *Index, ids []string, columns int, hasBack, hasPrev, hasNext bool) *models.InlineKeyboardMarkup {
	var rows [][]models.InlineKeyboardButton
	rows = append(rows, navRow(hasBack))
	rows = append(rows, chunkInlineItems(idx, ids, columns)...)
	var pager []models.InlineKeyboardButton
	if hasPrev {
		pager = append(pager, models.InlineKeyboardButton{Text: btnPrev, CallbackData: cbPrev})
	}
	if hasNext {
		pager = append(pager, models.InlineKeyboardButton{Text: btnNext, CallbackData: cbNext})
	}
	if len(pager) > 0 {
		rows = append(rows, pager)
	}
	return &models.InlineKeyboardMarkup{InlineKeyboard: rows}
}

func chunkInlineItems(idx *Index, ids []string, columns int) [][]models.InlineKeyboardButton {
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
		data := cbRun + n.ID
		if n.Type == "category" {
			data = cbOpen + n.ID
		}
		row = append(row, models.InlineKeyboardButton{Text: n.Label(), CallbackData: data})
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

func (idx *Index) menuChildren(nodeID string) (children []string, title string, columns int, err error) {
	if nodeID == "" {
		return idx.Roots, "Menu", idx.DefaultColumns, nil
	}
	n, ok := idx.ByID[nodeID]
	if !ok {
		return nil, "", 0, fmt.Errorf("unknown node %q", nodeID)
	}
	if n.Type != "category" {
		return nil, "", 0, fmt.Errorf("node %q is not a category", n.ID)
	}
	columns = n.Columns
	if columns < 1 {
		columns = idx.DefaultColumns
	}
	return n.Children, n.Label(), columns, nil
}

func (idx *Index) pageSize() int {
	if idx.PageSize < 1 {
		return 8
	}
	return idx.PageSize
}

func paginateIDs(ids []string, pageSize, page int) (clampedPage, totalPages int, slice []string) {
	total := len(ids)
	totalPages = 1
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
	return page, totalPages, ids[start:end]
}

func (idx *Index) clampPage(nodeID string, page int) int {
	children, _, _, err := idx.menuChildren(nodeID)
	if err != nil {
		return 0
	}
	page, _, _ = paginateIDs(children, idx.pageSize(), page)
	return page
}

// ConfirmInlineKeyboard is Yes/Cancel plus Home (and Back when nested).
func ConfirmInlineKeyboard(hasBack bool) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			navRow(hasBack),
			{
				{Text: btnYes, CallbackData: cbYes},
				{Text: btnCancel, CallbackData: cbCancel},
			},
		},
	}
}

// ResultInlineKeyboard is Home (and Back when nested) under a command result.
func ResultInlineKeyboard(hasBack bool) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{navRow(hasBack)},
	}
}

// homeReplyKeyboard is the persistent keyboard shown under the message box.
// It only ever has one button (Home) so it never needs its own layout
// logic; every other button lives in the inline keyboard on the menu
// message instead.
func homeReplyKeyboard() *models.ReplyKeyboardMarkup {
	return &models.ReplyKeyboardMarkup{
		Keyboard:       [][]models.KeyboardButton{{{Text: btnHome}}},
		ResizeKeyboard: true,
		IsPersistent:   true,
	}
}

// ChildByLabel finds a direct child of nodeID (empty = root) by display label.
func (idx *Index) ChildByLabel(nodeID, label string) *Node {
	var ids []string
	if nodeID == "" {
		ids = idx.Roots
	} else {
		n := idx.ByID[nodeID]
		if n == nil {
			return nil
		}
		ids = n.Children
	}
	for _, id := range ids {
		c := idx.ByID[id]
		if c != nil && c.Label() == label {
			return c
		}
	}
	return nil
}
