package middleware

import (
	"github.com/makarellav/phogo/context"
	"github.com/makarellav/phogo/controllers"
	"github.com/makarellav/phogo/models"
	"net/http"
)

type UserMiddleware struct {
	SessionService *models.SessionService
}

func (umw *UserMiddleware) SetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := controllers.ReadCookie(r, controllers.CookieSession)

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		user, err := umw.SessionService.User(token)

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithUser(r.Context(), user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (umw *UserMiddleware) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := context.User(r.Context())

		if user == nil {
			http.Redirect(w, r, "/signin", http.StatusFound)

			return
		}

		next.ServeHTTP(w, r)
	})
}
