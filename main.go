package main

import (
	"fmt"
	"net/http"

	"github.com/andrmnz/go_web/controllers"
	models "github.com/andrmnz/go_web/models/sql"
	"github.com/andrmnz/go_web/templates"
	"github.com/andrmnz/go_web/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"layout.gohtml", "home.gohtml",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"layout.gohtml", "contact.gohtml",
	))))
	
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	userService := models.UserService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService: &userService,
	}

	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"layout.gohtml", "signup.gohtml",
	))

	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"layout.gohtml", "signin.gohtml",
	))

	r.Get("/signup", usersC.New)
	r.Get("/signin", usersC.SignIn)
	r.Post("/users", usersC.Create)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS,
		"layout.gohtml", "faq.gohtml",
	))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on port 8888")

	http.ListenAndServe(":8888", r)
}
