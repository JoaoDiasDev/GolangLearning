package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// RenderTemplate renders a tmpl template
func RenderTemplateOld(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if we already have the template in our cache
	_, inMap := templateCache[t]
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println("error creating template cache:", err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = templateCache[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("error parsing template:", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache
	templateCache[t] = tmpl

	return nil
}