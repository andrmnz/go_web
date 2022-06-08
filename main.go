package main

import (
	"net/http"
	"path/filepath"

	"github.com/andrmnz/go_web/controllers"
	"github.com/andrmnz/go_web/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	tmpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))

	if err != nil {
		panic(err)
	}
	
	r.Get("/", controllers.StaticHandler(tmpl))

	tmpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))

	if err != nil {
		panic(err)
	}
	
	r.Get("/contact", controllers.StaticHandler(tmpl))

	tmpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))

	if err != nil {
		panic(err)
	}

	r.Get("/faq", controllers.StaticHandler(tmpl))
	
	
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	http.ListenAndServe(":8888", r)
}
