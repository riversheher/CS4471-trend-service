package web

import (
	"html/template"
	"path/filepath"
	"time"
)

// TemplateData holds generic data that can be passed to templates
type TemplateData struct {
	CurrentYear int
}

func newTemplateData() *TemplateData {
	return &TemplateData{
		CurrentYear: time.Now().Year(),
	}
}

// Creates a new template cache and returns a map of templates
func NewTemplateCache() (map[string]*template.Template, error) {
	// A map to act as a cache for templates
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./web/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		//Extract the file name from the full filepath
		name := filepath.Base(page)

		//Create a new template set
		ts, err := template.New(name).ParseFiles("./web/base.html")
		if err != nil {
			return nil, err
		}

		// Add partials to the template set
		ts, err = ts.ParseGlob("./web/partials/*.html")
		if err != nil {
			return nil, err
		}

		//Parse the page template file
		_, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the cache
		cache[name] = ts
	}

	return cache, nil
}
