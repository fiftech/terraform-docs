/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	_ "embed" //nolint
	"fmt"
	"regexp"
	gotemplate "text/template"

	"github.com/terraform-docs/terraform-docs/print"
	"github.com/terraform-docs/terraform-docs/template"
	"github.com/terraform-docs/terraform-docs/terraform"
)

//go:embed templates/pretty.tmpl
var prettyTpl []byte

// pretty represents colorized pretty format.
type pretty struct {
	*print.Generator

	config   *print.Config
	template *template.Template
	settings *print.Settings
}

// NewPretty returns new instance of Pretty.
func NewPretty(config *print.Config) Type {
	settings, _ := config.Extract()

	tt := template.New(settings, &template.Item{
		Name: "pretty",
		Text: string(prettyTpl),
	})
	tt.CustomFunc(gotemplate.FuncMap{
		"colorize": func(c string, s string) string {
			r := "\033[0m"
			if !settings.ShowColor {
				c = ""
				r = ""
			}
			return fmt.Sprintf("%s%s%s", c, s, r)
		},
	})

	return &pretty{
		Generator: print.NewGenerator("pretty", config.ModuleRoot),
		config:    config,
		template:  tt,
		settings:  settings,
	}
}

// Generate a Terraform module document.
func (p *pretty) Generate(module *terraform.Module) error {
	rendered, err := p.template.Render("pretty", module)
	if err != nil {
		return err
	}

	p.Generator.Funcs(print.WithContent(regexp.MustCompile(`(\r?\n)*$`).ReplaceAllString(rendered, "")))

	return nil
}

func init() {
	register(map[string]initializerFn{
		"pretty": NewPretty,
	})
}
