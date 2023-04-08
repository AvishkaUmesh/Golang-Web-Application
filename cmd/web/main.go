package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AvishkaUmesh/Golang-Web-Application/pkg/config"
	"github.com/AvishkaUmesh/Golang-Web-Application/pkg/handlers"
	"github.com/AvishkaUmesh/Golang-Web-Application/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
