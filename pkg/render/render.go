package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/fxfrancky/bookings/pkg/config"
	"github.com/fxfrancky/bookings/pkg/models"
)

var app *config.AppConfig

var functions = template.FuncMap{}

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate render a templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

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
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	// Copy my template into buf
	_ = t.Execute(buf, td)
	// Write the buffer into my writer
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser ", err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all the pages
	pages, err := filepath.Glob("./templates/*.page.tmpl")
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
		matches, err := filepath.Glob("./templates/*layout.tmpl")

		// if its the layout template
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		// Add a layout to our list
		myCache[name] = ts
	}
	return myCache, nil
}
