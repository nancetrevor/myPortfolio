package handlers

import (
	"net/http"
	"text/template"
)

type TerminalPageData struct {
	Title string
}

func TerminalHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/terminal.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	h := HomePageData{
		Title: "Terminal",
	}

	t.Execute(w, h)
}
