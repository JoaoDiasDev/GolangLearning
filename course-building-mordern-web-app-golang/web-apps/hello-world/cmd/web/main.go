package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joaodiasdev/hellowordwebapp/pkg/config"
	"github.com/joaodiasdev/hellowordwebapp/pkg/handlers"
	"github.com/joaodiasdev/hellowordwebapp/pkg/render"
)

const portNumber = ":8080"

// main is the entry point for the application
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
