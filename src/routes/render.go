package routes

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(templatePath string) http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles(templatePath))

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			log.Printf("Error rendering template: %v", err)
		}
	}
}
