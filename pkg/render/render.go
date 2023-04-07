package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/AvishkaUmesh/Golang-Web-Application/pkg/config"
	"github.com/AvishkaUmesh/Golang-Web-Application/pkg/models"
)

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {

	var templateCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	template, ok := templateCache[tmpl]

	if !ok {
		fmt.Println("Error getting template from cache")
		return
	}

	buf := new(bytes.Buffer)

	// add default data
	data = AddDefaultData(data)

	err := template.Execute(buf, data)

	if err != nil {
		fmt.Println("Error executing template :", err)
		return
	}

	// render template
	_, err = buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser :", err)
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// get all pages from templates folder
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return templateCache, err
	}

	// loop through pages
	for _, page := range pages {
		// get page name
		name := filepath.Base(page)

		// parse page template
		templateSet, err := template.New(name).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		// get layout
		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return templateCache, err
		}

		// add layout to template set
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return templateCache, err
			}
		}

		// add template set to template cache
		templateCache[name] = templateSet

	}

	return templateCache, nil

}
