package app

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var templates = map[string]*template.Template{}

const templateDir = "templates"

func joinTemplateDir(files []string) []string {
	result := make([]string, len(files))
	for i := range files {
		result[i] = filepath.Join(templateDir, files[i])
	}
	return result
}

func parseTemplates(sets [][]string) error {
	for _, set := range sets {
		templateName := set[0]
		t := template.New("")
		t.Funcs(template.FuncMap{
			"templateName": func() string { return templateName },
		})
		if _, err := t.ParseFiles(joinTemplateDir(set)...); err != nil {
			return err
		}
		t = t.Lookup("ROOT")
		if t == nil {
			return fmt.Errorf("root template not found in %v", set)
		}
		templates[set[0]] = t
	}
	return nil
}

func executeTemplate(w http.ResponseWriter, name string, status int, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	t := templates[name]
	if t == nil {
		http.Error(w, fmt.Sprintf("template %s not found", name), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
