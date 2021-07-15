package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArmanurRahman/go-app/pkg/config"
	"github.com/ArmanurRahman/go-app/pkg/handlers"
	"github.com/ArmanurRahman/go-app/pkg/render"
)

const port = ":8080"

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

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	//http.HandleFunc("/devide", Devide)

	fmt.Println("Starting listining to port ", port)
	_ = http.ListenAndServe(port, nil)
}
