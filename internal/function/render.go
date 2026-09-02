package function

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"
	"text/template/parse"
)

// RenderRun renders the function's run template with the given params.
func (d *Definition) RenderRun(params map[string]string) (string, error) {
	values := make(map[string]string, len(d.Params))
	for _, p := range d.Params {
		if v, ok := params[p.Name]; ok && v != "" {
			values[p.Name] = v
			continue
		}
		if p.Required {
			return "", fmt.Errorf("missing required parameter %q", p.Name)
		}
		values[p.Name] = p.Default
	}
	// Allow extra params in the template map as well.
	for k, v := range params {
		if _, ok := values[k]; !ok {
			values[k] = v
		}
	}

	tmpl, err := template.New(d.Name).Option("missingkey=error").Parse(d.Run)
	if err != nil {
		return "", fmt.Errorf("parse run template: %w", err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, values); err != nil {
		return "", fmt.Errorf("execute run template: %w", err)
	}
	return buf.String(), nil
}

// ValidateTemplate parses the run template to catch syntax errors early.
func (d *Definition) ValidateTemplate() error {
	tmpl, err := template.New(d.Name).Parse(d.Run)
	if err != nil {
		return fmt.Errorf("function %q: invalid run template: %w", d.Name, err)
	}
	declared := make(map[string]struct{}, len(d.Params))
	for _, param := range d.Params {
		declared[param.Name] = struct{}{}
	}
	var undefined string
	walkTemplateNode(tmpl.Tree.Root, func(field *parse.FieldNode) bool {
		if len(field.Ident) == 0 {
			return true
		}
		if _, ok := declared[field.Ident[0]]; !ok {
			undefined = field.Ident[0]
			return false
		}
		return true
	})
	if undefined != "" {
		return fmt.Errorf("function %q: run template references undefined parameter %q", d.Name, undefined)
	}
	return nil
}

func walkTemplateNode(node parse.Node, visit func(*parse.FieldNode) bool) bool {
	if node == nil || reflect.ValueOf(node).IsNil() {
		return true
	}
	switch n := node.(type) {
	case *parse.ListNode:
		for _, child := range n.Nodes {
			if !walkTemplateNode(child, visit) {
				return false
			}
		}
	case *parse.ActionNode:
		return walkTemplateNode(n.Pipe, visit)
	case *parse.IfNode:
		return walkBranchNode(&n.BranchNode, visit)
	case *parse.RangeNode:
		return walkBranchNode(&n.BranchNode, visit)
	case *parse.WithNode:
		return walkBranchNode(&n.BranchNode, visit)
	case *parse.TemplateNode:
		return walkTemplateNode(n.Pipe, visit)
	case *parse.PipeNode:
		for _, command := range n.Cmds {
			if !walkTemplateNode(command, visit) {
				return false
			}
		}
	case *parse.CommandNode:
		for _, arg := range n.Args {
			if !walkTemplateNode(arg, visit) {
				return false
			}
		}
	case *parse.ChainNode:
		return walkTemplateNode(n.Node, visit)
	case *parse.FieldNode:
		return visit(n)
	}
	return true
}

func walkBranchNode(node *parse.BranchNode, visit func(*parse.FieldNode) bool) bool {
	return walkTemplateNode(node.Pipe, visit) &&
		walkTemplateNode(node.List, visit) &&
		walkTemplateNode(node.ElseList, visit)
}
