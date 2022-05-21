package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("contact"))
}
func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("faq"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	http.ListenAndServe(":8888", r)
}
