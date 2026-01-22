package routes

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates = template.Must(
	template.ParseGlob("web/templates/*.html"),
)

func RegisterWebRoutes(r *mux.Router) {

	// Static assets
	r.PathPrefix("/static/").
		Handler(http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("web/static")),
		))

	// Public pages
	r.HandleFunc("/", render("login.html")).Methods("GET")
	r.HandleFunc("/login", render("login.html")).Methods("GET")

	// Protected pages
	r.HandleFunc(
		"/dashboard",
		render("dashboard.html"),
	).Methods("GET")

	r.HandleFunc(
		"/profile",
		render("profile.html"),
	).Methods("GET")

	r.HandleFunc(
		"/scan",
		render("scan.html"),
	).Methods("GET")
}

func render(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, name, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
