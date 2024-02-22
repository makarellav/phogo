package controllers

import (
	"fmt"
	"github.com/makarellav/phogo/context"
	"github.com/makarellav/phogo/models"
	"net/http"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService    *models.UserService
	SessionService *models.SessionService
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")

	u.Templates.New.Execute(w, r, data)
}

func (u *Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")

	u.Templates.SignIn.Execute(w, r, data)
}

func (u *Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user, err := u.UserService.Authenticate(email, password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong...", http.StatusInternalServerError)

		return
	}

	session, err := u.SessionService.Create(user.ID)

	if err != nil {
		fmt.Println(err)

		http.Error(w, "something went wrong...", http.StatusInternalServerError)

		return
	}

	setCookie(w, CookieSession, session.Token)

	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u *Users) ProcessSignOut(w http.ResponseWriter, r *http.Request) {
	token, err := ReadCookie(r, CookieSession)

	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusFound)

		return
	}

	err = u.SessionService.Delete(token)

	if err != nil {
		fmt.Println(err)

		http.Error(w, "something went wrong...", http.StatusInternalServerError)

		return
	}

	deleteCookie(w, CookieSession)
	http.Redirect(w, r, "/signin", http.StatusFound)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user, err := u.UserService.Create(email, password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	session, err := u.SessionService.Create(user.ID)

	if err != nil {
		fmt.Println(err)

		http.Redirect(w, r, "/signin", http.StatusFound)

		return
	}

	setCookie(w, CookieSession, session.Token)

	fmt.Fprintf(w, "user created: %+v\n", user)
}

func (u *Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	context.User(r.Context())

	token, err := ReadCookie(r, CookieSession)

	if err != nil {
		fmt.Println(err)

		http.Redirect(w, r, "signin", http.StatusFound)
	}

	user, err := u.SessionService.User(token)

	if err != nil {
		fmt.Println(err)

		http.Redirect(w, r, "/signin", http.StatusFound)
	}

	fmt.Fprintf(w, "Email: %s", user.Email)
}
