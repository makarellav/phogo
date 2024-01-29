package views

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTmpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.htmlTmpl.Execute(w, data)

	if err != nil {
		log.Printf("failed to execute a template: %v", err)
		http.Error(w, "Something went terribly wrong...", http.StatusInternalServerError)

		return
	}
}

func MustParseFS(fs fs.FS, patterns ...string) Template {
	tmpl, err := template.ParseFS(fs, patterns...)

	tmpl = template.Must(tmpl, err)

	return Template{
		htmlTmpl: tmpl,
	}
}

func MustParse(tmplPath string) Template {
	tmpl, err := template.ParseFiles(tmplPath)

	tmpl = template.Must(tmpl, err)

	return Template{
		htmlTmpl: tmpl,
	}
}
