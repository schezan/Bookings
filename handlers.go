package main

import "net/http"

type Repository struct {
	App *AppConfig
}

var Repo *Repository //the repository used by the handlers

//Creates a new repository
func NewRepo(a *AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func Newhandler(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl", &TemplateData{})
}

//we have added reciever of type repository to both Home and about
// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Fucking piling on"

	RenderTemplate(w, "about.page.tmpl", &TemplateData{
		StringMap: stringMap,
	})

}
