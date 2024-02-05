package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/makarellav/phogo/controllers"
	"github.com/makarellav/phogo/models"
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

	db, err := models.Open(models.DevConfig())

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	userService := models.UserService{
		DB: db,
	}

	usersController := controllers.Users{
		UserService: &userService,
	}
	usersController.Templates.New = views.MustParseFS(templates.FS, "layout.gohtml", "signup.gohtml")
	usersController.Templates.SignIn = views.MustParseFS(templates.FS, "layout.gohtml", "signin.gohtml")

	r.Get("/signup", usersController.New)
	r.Get("/signin", usersController.SignIn)
	r.Post("/users", usersController.Create)
	r.Post("/signin", usersController.ProcessSignIn)
	r.Get("/users/me", usersController.CurrentUser)

	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"))

	log.Fatal(http.ListenAndServe(":3000", CSRF(r)))
}
