package controllers

import (
	"fmt"
	"net/http"
)

const (
	CookieSession = "session"
	MaxAgeDelete  = -1
)

func newCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
}

func setCookie(w http.ResponseWriter, name, value string) {
	cookie := newCookie(name, value)

	http.SetCookie(w, cookie)
}

func ReadCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)

	if err != nil {
		return "", fmt.Errorf("read cookie: %w", err)
	}

	return cookie.Value, nil
}

func deleteCookie(w http.ResponseWriter, name string) {
	cookie := newCookie(name, "")
	cookie.MaxAge = MaxAgeDelete

	http.SetCookie(w, cookie)
}
