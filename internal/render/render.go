package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/fxfrancky/bookings/internal/config"
	"github.com/fxfrancky/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig
var pathToTemplates = "./templates"

var functions = template.FuncMap{}

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data for all templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// Template render a templates using html/template
func Template(w http.ResponseWriter, tmpl string, td *models.TemplateData, r *http.Request) error {

	var tc map[string]*template.Template
	if app.UseCache {
		// Get the template cache from app congfig
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// If we have templates
	// Ok will be true if we have a tmpl value in tc, and false otherwise
	t, ok := tc[tmpl]
	if !ok {
		// template not found
		return errors.New("could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	// Copy my template into buf
	_ = t.Execute(buf, td)
	// Write the buffer into my writer
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser ", err)
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all the pages
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		// Check the current page
		// fmt.Println("Page is currently ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// Matches the layout base
		matches, err := filepath.Glob(fmt.Sprintf("%s/*layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		// if its the layout template
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		// Add a layout to our list
		myCache[name] = ts
	}
	return myCache, nil
}
