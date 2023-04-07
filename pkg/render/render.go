package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// create template cache
	templateCache, err := createTemplateCache()

	if err != nil {
		fmt.Println("Error creating template cache :", err)
		return
	}

	// get requested template from cache
	template, ok := templateCache[tmpl]

	if !ok {
		fmt.Println("Error getting template from cache :", err)
		return
	}

	buf := new(bytes.Buffer)

	err = template.Execute(buf, nil)

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

func createTemplateCache() (map[string]*template.Template, error) {
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
