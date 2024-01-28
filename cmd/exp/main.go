package main

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.html")

	if err != nil {
		log.Fatal(":(")
	}

	u := User{
		Name: "Vladyslav",
	}

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err = t.Execute(w, u)

		if err != nil {
			log.Fatal(":(")
		}
	})

	log.Fatal(http.ListenAndServe(":3000", r))
}
