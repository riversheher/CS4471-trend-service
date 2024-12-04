package web

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type Application struct {
	templateCache map[string]*template.Template
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Initialize() error {
	cache, err := NewTemplateCache()
	if err != nil {
		return err
	}

	a.templateCache = cache

	return nil
}

func (app *Application) Render(w http.ResponseWriter, page string) {

	// Grab the template set from the cache
	templateData := newTemplateData()
	ts, ok := app.templateCache[page]
	if !ok {
		fmt.Printf("Template %s does not exist", page)
		http.Error(w, "This page does not exist", http.StatusNotFound)
		return
	}

	// Render template to buffer to check for complete rendering
	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", templateData)
	if err != nil {
		fmt.Printf("Error rendering template %s", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	// Write the buffer to the response writer
	buf.WriteTo(w)
}
