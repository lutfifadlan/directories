package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Renders an HTML template
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templatePath := filepath.Join("templates", tmpl+".html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Welcome to My Website",
		"Body":  "This is the homepage.",
	}
	RenderTemplate(w, "index", data)
}
