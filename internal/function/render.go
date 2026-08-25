package function

import (
	"bytes"
	"fmt"
	"text/template"
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
	_, err := template.New(d.Name).Parse(d.Run)
	if err != nil {
		return fmt.Errorf("function %q: invalid run template: %w", d.Name, err)
	}
	return nil
}
