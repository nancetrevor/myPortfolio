package handlers

import (
	"net/http"
	"text/template"
)

type HomePageData struct {
	Title string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	h := HomePageData{
		Title: "Home",
	}

	t.Execute(w, h)
}
