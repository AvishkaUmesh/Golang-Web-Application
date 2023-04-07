package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if template is in cache
	_, inMap := templateCache[t]

	if !inMap {
		// if not, parse template and add to cache
		log.Println("Template", t, "not found in cache. Parsing template.")
		err = createTemplateCache(t)
		if err != nil {
			log.Println("Error creating template cache:", err)
			return
		}

	} else {
		// if it is, get from cache
		log.Println("Template", t, "found in cache")
	}

	tmpl = templateCache[t]
	// execute template
	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	// parse template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err

	}

	// add to cache
	templateCache[t] = tmpl

	return nil

}
