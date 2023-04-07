package handlers

import (
	"net/http"

	"github.com/AvishkaUmesh/Golang-Web-Application/pkg/render"
)

// Home is a function to handle home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is a function to handle about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
