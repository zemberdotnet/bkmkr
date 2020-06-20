package main

import (
	"html/template"
	"path/filepath"
	//	"zember.net/bkmrk/pkg/models"
)

type templateData struct {
	//	People []*models.Person
	Forms  *Forms
	People *[]map[string]interface{}
}

func newTemplateCache(dir string) (temp map[string]*template.Template, e error) {

	cache := map[string]*template.Template{}
	// We join the directory and *.page.tmpl to get all page templates
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}
	// I'm going to guess we parse them
	for _, page := range pages {
		name := filepath.Base(page)
		// If we want to add custom functions later in the program then
		// We can switch to this form
		//ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// So now we've made our base templates, we add in the other
		// parts using ts.ParseGlob. Our existing templates are ts.
		// ParseGlob pattern matches and parses any files it finds
		// Importance of file naming!!!
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*partial.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
