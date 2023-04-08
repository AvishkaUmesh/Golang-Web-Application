package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AvishkaUmesh/Golang-Web-Application/pkg/config"
	"github.com/AvishkaUmesh/Golang-Web-Application/pkg/handlers"
	"github.com/AvishkaUmesh/Golang-Web-Application/pkg/render"
)

const PORT = ":8080"

func main() {

	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println("Server is running on port", PORT)

	serv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
