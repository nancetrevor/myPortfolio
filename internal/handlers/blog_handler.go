package handlers

import (
	"net/http"
	"text/template"
)

type BlogPageData struct {
	Title     string
	BlogPosts []string
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/blog.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	h := BlogPageData{
		Title: "Blog",
	}

	t.Execute(w, h)
}

// func parseMdToBlogPosts() BlogPageData {
//
// }
