package bot

import (
	"fmt"

	"github.com/go-telegram/bot/models"
)

const (
	btnHome       = "🏠 Home"
	btnBack       = "🔙 Back"
	btnRunCommand = "⌨️ Run Command"
	btnPrev       = "⏪ Prev"
	btnNext       = "Next ⏩"
	btnYes        = "✅ Yes"
	btnCancel     = "❌ Cancel"

	cbHome   = "h"
	cbBack   = "b"
	cbPrev   = "p"
	cbNext   = "n"
	cbYes    = "y"
	cbCancel = "x"
	cbOpen   = "o:"
	cbRun    = "r:"
)

// MenuBuild is one menu screen: title text plus the reply keyboard shown
// under the message box.
type MenuBuild struct {
	Title      string
	Page       int
	TotalPages int
	Reply      *models.ReplyKeyboardMarkup
	HasBack    bool
	HasPrev    bool
	HasNext    bool
}

// BuildMenu builds the reply keyboard for a category (or root when nodeID is empty).
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
		Reply:      menuReplyKeyboard(idx, slice, columns, hasBack, hasPrev, hasNext),
		HasBack:    hasBack,
		HasPrev:    hasPrev,
		HasNext:    hasNext,
	}, nil
}

// replyKeyboard wraps rows into a persistent, full-width reply keyboard.
// A reply keyboard always spans the whole chat width, so button text never
// gets squeezed the way it can inside a narrow inline keyboard.
func replyKeyboard(rows [][]models.KeyboardButton) *models.ReplyKeyboardMarkup {
	return &models.ReplyKeyboardMarkup{
		Keyboard:       rows,
		ResizeKeyboard: true,
		IsPersistent:   true,
	}
}

func navRow(hasBack, showRun bool) []models.KeyboardButton {
	nav := []models.KeyboardButton{{Text: btnHome}}
	if hasBack {
		nav = append(nav, models.KeyboardButton{Text: btnBack})
	}
	if showRun {
		nav = append(nav, models.KeyboardButton{Text: btnRunCommand})
	}
	return nav
}

func menuReplyKeyboard(idx *Index, ids []string, columns int, hasBack, hasPrev, hasNext bool) *models.ReplyKeyboardMarkup {
	var rows [][]models.KeyboardButton
	rows = append(rows, navRow(hasBack, idx.EnableRunCommand))
	rows = append(rows, chunkItems(idx, ids, columns)...)
	var pager []models.KeyboardButton
	if hasPrev {
		pager = append(pager, models.KeyboardButton{Text: btnPrev})
	}
	if hasNext {
		pager = append(pager, models.KeyboardButton{Text: btnNext})
	}
	if len(pager) > 0 {
		rows = append(rows, pager)
	}
	return replyKeyboard(rows)
}

func chunkItems(idx *Index, ids []string, columns int) [][]models.KeyboardButton {
	if columns < 1 {
		columns = 2
	}
	var rows [][]models.KeyboardButton
	var row []models.KeyboardButton
	for _, id := range ids {
		n := idx.ByID[id]
		if n == nil {
			continue
		}
		row = append(row, models.KeyboardButton{Text: n.Label()})
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

// ConfirmReplyKeyboard is Yes/Cancel plus Home (and Back when nested).
// Run Command is included when showRun is true.
func ConfirmReplyKeyboard(hasBack, showRun bool) *models.ReplyKeyboardMarkup {
	return replyKeyboard([][]models.KeyboardButton{
		navRow(hasBack, showRun),
		{
			{Text: btnYes},
			{Text: btnCancel},
		},
	})
}

// ResultReplyKeyboard is Home (and Back when nested) under a command result
// or the Run Command prompt. Run Command is included when showRun is true.
func ResultReplyKeyboard(hasBack, showRun bool) *models.ReplyKeyboardMarkup {
	return replyKeyboard([][]models.KeyboardButton{navRow(hasBack, showRun)})
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
