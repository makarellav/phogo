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

	var usersController controllers.Users
	usersController.Templates.New = views.MustParseFS(templates.FS, "layout.gohtml", "signup.gohtml")
	r.Handle("/signup", http.RedirectHandler("/users/new", http.StatusMovedPermanently))
	r.Get("/users/new", usersController.New)
	r.Post("/signup", usersController.Create)

	//tmpl = views.MustParseFS(templates.FS, "layout.gohtml", "exp.gohtml")
	//r.Get("/exp", controllers.StaticHandler(tmpl))
	//r.Post("/exp-create", usersController.Exp)

	log.Fatal(http.ListenAndServe(":3000", r))
}
