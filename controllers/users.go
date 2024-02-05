package controllers

import (
	"fmt"
	"github.com/makarellav/phogo/models"
	"net/http"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")

	u.Templates.New.Execute(w, r, data)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")

	u.Templates.SignIn.Execute(w, r, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user, err := u.UserService.Authenticate(email, password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong...", http.StatusInternalServerError)

		return
	}

	cookie := http.Cookie{
		Name:     "email",
		Value:    user.Email,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "signed in successfully: %+v", user)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user, err := u.UserService.Create(email, password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "user created: %+v\n", user)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")

	if err != nil {
		http.Error(w, "failed to read a cookie", http.StatusInternalServerError)

		return
	}

	fmt.Fprintf(w, "%s: %s", email.Name, email.Value)
}
