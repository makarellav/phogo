package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/makarellav/phogo/controllers"
	mw "github.com/makarellav/phogo/middleware"
	"github.com/makarellav/phogo/migrations"
	"github.com/makarellav/phogo/models"
	"github.com/makarellav/phogo/templates"
	"github.com/makarellav/phogo/views"
	"log"
	"net/http"
)

func main() {
	// Setup DB
	db, err := models.Open(models.DevConfig())

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")

	if err != nil {
		log.Fatal(err)
	}

	// Setup services
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	// Setup middleware
	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"))
	userMiddleware := mw.UserMiddleware{
		SessionService: &sessionService,
	}

	// Setup controllers
	usersController := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersController.Templates.New = views.MustParseFS(templates.FS, "layout.gohtml", "signup.gohtml")
	usersController.Templates.SignIn = views.MustParseFS(templates.FS, "layout.gohtml", "signin.gohtml")

	// Setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(CSRF)
	r.Use(userMiddleware.SetUser)

	tmpl := views.MustParseFS(templates.FS, "layout.gohtml", "home.gohtml")
	r.Get("/", controllers.StaticHandler(tmpl))

	tmpl = views.MustParseFS(templates.FS, "layout.gohtml", "contact.gohtml")
	r.Get("/contact", controllers.StaticHandler(tmpl))

	tmpl = views.MustParseFS(templates.FS, "layout.gohtml", "faq.gohtml")
	r.Get("/faq", controllers.FAQ(tmpl))

	r.Get("/signup", usersController.New)
	r.Get("/signin", usersController.SignIn)
	r.Post("/users", usersController.Create)
	r.Post("/signin", usersController.ProcessSignIn)
	r.Post("/signout", usersController.ProcessSignOut)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(userMiddleware.RequireUser)

		r.Get("/", usersController.CurrentUser)
	})

	// Run the server
	log.Fatal(http.ListenAndServe(":3000", CSRF(userMiddleware.SetUser(r))))
}
