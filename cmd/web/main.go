package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mateunix/webapp-go/pkg/config"
	"github.com/mateunix/webapp-go/pkg/handlers"
	"github.com/mateunix/webapp-go/pkg/render"
)

const portNumber = ":8080"

// Func main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
