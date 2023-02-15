package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/fxfrancky/bookings/internal/config"
	"github.com/fxfrancky/bookings/internal/handlers"
	"github.com/fxfrancky/bookings/internal/render"
)

var app config.AppConfig
var session *scs.SessionManager

const portNumber = ":8080"

// main is the main application function
func main() {

	// Change this to true when in production
	app.InProduction = false

	// Initialize session scs
	session = scs.New()
	// Set The session duration (24 H)
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// Update the config session field
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", portNumber)
	// Use the previous function

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
