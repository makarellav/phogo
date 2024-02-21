package views

import (
	"bytes"
	"fmt"
	"github.com/gorilla/csrf"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTmpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data any) {
	tmpl, err := t.htmlTmpl.Clone()

	if err != nil {
		fmt.Printf("cloning template: %v\n", err)
		http.Error(w, "There was an error rendering the page", http.StatusInternalServerError)

		return
	}

	tmpl = tmpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
	})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)

	if err != nil {
		log.Printf("failed to execute a template: %v", err)
		http.Error(w, "Something went terribly wrong...", http.StatusInternalServerError)

		return
	}

	_, err = io.Copy(w, &buf)

	if err != nil {
		log.Printf("failed to execute a template: %v", err)
		http.Error(w, "Something went terribly wrong...", http.StatusInternalServerError)

		return
	}
}

func MustParseFS(fs fs.FS, patterns ...string) Template {
	tmpl := template.New(patterns[0])
	tmpl = tmpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return `<!-- -->`
			},
		},
	)

	tmpl, err := tmpl.ParseFS(fs, patterns...)

	tmpl = template.Must(tmpl, err)

	return Template{
		htmlTmpl: tmpl,
	}
}
