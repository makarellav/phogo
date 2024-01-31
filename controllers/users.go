package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")

	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	fmt.Fprint(w, fmt.Sprintf("%s %s", email, password))
}

func (u Users) Exp(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	address := r.FormValue("address")
	remember := r.FormValue("remember")
	hobbies := r.Form["hobby"]
	sex := r.FormValue("sex")
	age := r.FormValue("age")
	file, info, err := r.FormFile("document")

	if err != nil {
		fmt.Println(err)
	}

	if file != nil {
		defer file.Close()
	} else {
		fmt.Println("no file")
	}

	fmt.Fprintf(w, "Email: %s\nPassword: %s\nAddress: %s\nRemember: %v\nHobbies: %v\nSex: %s\nAge: %s\nFile name: %s\n",
		email, password, address, remember, hobbies, sex, age, info.Filename)
}
