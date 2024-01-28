package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>This is a home page!!!</p>")
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "This is a contact page!!!<p>Email me at <a href='mailto:makarellads@gmail.com'>makarellads@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<ul>
<li>Q: Is there a free version?</li>
<li>A: Yes! We offer free trial for 30 days on any paid plans.</li>
<br />
<li>Q: What are your support hours?</li>
<li>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends</li>
<br />
<li>Q: How do I contact support?</li>
<li>A: Email us - <a href='mailto:support@lenslocked.com'>support@lenslocked.com</a></li>
</ul>`)
}

func URLParamHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>URL PARAM: %s</h1>", idParam)
}

func main() {
	r := chi.NewRouter()

	//r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)

		r.Get("/", homeHandler)
	})

	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(404)
		fmt.Fprint(w, "<h1>PAGE NOT FOUND :(</h1>")
	})
	r.Get("/url-params/{id}", URLParamHandler)

	log.Fatal(http.ListenAndServe(":3000", r))
}
