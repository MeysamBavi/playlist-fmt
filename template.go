package main

import (
	"flag"
	"text/template"
)

const defaultTemplate = `[{{.Name}}]({{.URL}})
{{with .Description}}{{.}}{{end}}
{{range $i, $t := .Tracks}}
{{inc $i}}. [{{$t.Name}}]({{$t.URL}}) - {{index $t.Artists 0}}
{{end}}`

var templatePath = flag.String("template", "", "output template's path")

func getTemplate() (*template.Template, error) {
	funcMap := template.FuncMap{
		"inc": func(a int) int {
			return a + 1
		},
	}
	if *templatePath == "" {
		return template.New("").Funcs(funcMap).Parse(defaultTemplate)
	}
	return template.ParseFiles(*templatePath)
}
