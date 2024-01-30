package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/makarellav/phogo/controllers"
	"github.com/makarellav/phogo/templates"
	"github.com/makarellav/phogo/views"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tmpl := views.MustParseFS(templates.FS, "layout.gohtml", "home.gohtml")
	r.Get("/", controllers.StaticHandler(tmpl))

	tmpl = views.MustParseFS(templates.FS, "layout.gohtml", "contact.gohtml")
	r.Get("/contact", controllers.StaticHandler(tmpl))

	tmpl = views.MustParseFS(templates.FS, "layout.gohtml", "faq.gohtml")
	r.Get("/faq", controllers.FAQ(tmpl))

	tmpl = views.MustParseFS(templates.FS, "layout.gohtml", "signup.gohtml")
	r.Get("/signup", controllers.StaticHandler(tmpl))

	log.Fatal(http.ListenAndServe(":3000", r))
}
