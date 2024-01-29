package main

import (
	"errors"
	"fmt"
)

//type Shelter struct {
//	Name    string
//	Address string
//}
//
//type Pet struct {
//	Name        string
//	Age         int
//	Weight      float64
//	Height      float64
//	SocialMedia map[string]string
//	Hobbies     []string
//	Shelter     Shelter
//}

type errorNotFound string

func (e errorNotFound) Error() string {
	return string(e)
}

const ErrNotFound = errorNotFound("not found :(")

func b() error {
	return fmt.Errorf("123 %w", ErrNotFound)
}

func main() {
	//t, err := template.ParseFiles("hello.html")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//p := Pet{
	//	Name:   "Cat",
	//	Age:    10,
	//	Weight: 2.99,
	//	Height: 52.492,
	//	SocialMedia: map[string]string{
	//		"Instagram": "https://instagram.com",
	//		"LinkedIn":  "https://linkedin.com",
	//	},
	//	Hobbies: []string{"jumping", "running", "eating"},
	//	Shelter: Shelter{
	//		Name:    "Luxury Cat Shelter",
	//		Address: "Luxury Street 1",
	//	},
	//}
	//
	//r := chi.NewRouter()
	//
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	err = t.Execute(w, p)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//})
	//
	//log.Fatal(http.ListenAndServe(":3000", r))
	err := b()
	unwrapped := errors.Unwrap(err)

	if err == ErrNotFound {
		fmt.Println(err)
	}

	if unwrapped == ErrNotFound {
		fmt.Println(err)
	}

	if errors.Is(err, ErrNotFound) {
		fmt.Println(err)
	}
}
