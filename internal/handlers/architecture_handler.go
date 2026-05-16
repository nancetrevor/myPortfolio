package handlers

import (
	"net/http"
	"text/template"
)

type ArchitecturePageData struct {
	Title string
}

func ArchitectureHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/mma_architecture.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	h := ArchitecturePageData{
		Title: "MMA_Architecture",
	}

	t.Execute(w, h)
}
