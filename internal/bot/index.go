// Package bot implements Telegram bot handlers for the button menu.
package bot

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/azolfagharj/telegram-commander/internal/config"
)

// Node is an indexed menu node with a stable ID.
type Node struct {
	ID       string
	ParentID string
	Name     string
	Type     string
	Icon     string
	Function string
	Confirm  bool
	Timeout  time.Duration
	WorkDir  string
	Env      map[string]string
	Columns  int
	Command  string
	Path     string
	Args     string
	Params   map[string]string
	Children []string // child IDs
	Raw      config.ButtonNode
}

// Index maps node IDs and holds the root child IDs.
type Index struct {
	Roots            []string
	ByID             map[string]*Node
	DefaultColumns   int
	PageSize         int
	EnableRunCommand bool
}

// BuildIndex walks the button tree and assigns stable IDs.
func BuildIndex(buttons []config.ButtonNode, defaultColumns, pageSize int) *Index {
	idx := &Index{
		ByID:           make(map[string]*Node),
		DefaultColumns: defaultColumns,
		PageSize:       pageSize,
	}
	for i := range buttons {
		id := assignID(&buttons[i], "", fmt.Sprintf("%d", i))
		idx.Roots = append(idx.Roots, id)
		walk(idx, &buttons[i], "", id, defaultColumns)
	}
	return idx
}

func assignID(n *config.ButtonNode, parentPath, path string) string {
	if n.ID != "" {
		return n.ID
	}
	sum := sha256.Sum256([]byte(parentPath + "/" + path + "/" + n.Name + "/" + n.Type))
	return hex.EncodeToString(sum[:8])
}

func walk(idx *Index, n *config.ButtonNode, parentID, id string, defaultColumns int) {
	cols := defaultColumns
	if n.Columns != nil {
		cols = *n.Columns
	}
	node := &Node{
		ID:       id,
		ParentID: parentID,
		Name:     n.Name,
		Type:     strings.ToLower(n.Type),
		Icon:     n.Icon,
		Function: n.Function,
		Confirm:  n.Confirm,
		WorkDir:  n.WorkDir,
		Env:      n.Env,
		Columns:  cols,
		Command:  n.Command,
		Path:     n.Path,
		Args:     n.Args,
		Params:   n.Params,
		Raw:      *n,
	}
	if n.Timeout != nil {
		node.Timeout = n.Timeout.Duration
	}
	for i := range n.Items {
		childPath := id + "." + fmt.Sprintf("%d", i)
		cid := assignID(&n.Items[i], id, childPath)
		node.Children = append(node.Children, cid)
		walk(idx, &n.Items[i], id, cid, defaultColumns)
	}
	idx.ByID[id] = node
}

// Label returns the display label for a node.
func (n *Node) Label() string {
	if n.Icon != "" {
		return stripVariationSelectors(n.Icon + " " + n.Name)
	}
	return n.Name
}

// stripVariationSelectors drops emoji variation selectors (U+FE0E, U+FE0F).
// Some Telegram Android clients miscompute an inline button's width when its
// label contains one, which can make the text spill outside the button. The
// base emoji still renders fine without it.
func stripVariationSelectors(s string) string {
	return strings.Map(func(r rune) rune {
		if r == '\uFE0E' || r == '\uFE0F' {
			return -1
		}
		return r
	}, s)
}
