package handlers

import (
	"net/http"
	"text/template"
)

type ArchitecturePageData struct {
	Title string
}

func ArchitectureHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/architecture.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	h := ArchitecturePageData{
		Title: "Architecture",
	}

	t.Execute(w, h)
}
