package controllers

import (
	"github.com/makarellav/phogo/views"
	"net/http"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}
