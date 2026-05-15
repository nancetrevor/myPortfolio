package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/lfizzikz/portfolioSite/package/markdown"
)

type BlogPage struct {
	Title     string
	BlogPosts []BlogPostData
}

type BlogPostData struct {
	Title   string
	Date    string
	Content string
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/blog.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	posts := parseMdToBlogPosts()

	t.Execute(w, posts)
}

func parseMdToBlogPosts() BlogPage {
	dir := "content/blog"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("error: ", err)
		return BlogPage{}
	}

	page := BlogPage{}

	for _, file := range files {
		file := filepath.Join(dir, file.Name())
		contents, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("error: ", err)
			return BlogPage{}
		}

		post, err := parseMdFile(string(contents))
		if err != nil {
			fmt.Println("error: ", err)
			return BlogPage{}
		}
		page.BlogPosts = append(page.BlogPosts, post)
	}
	return page
}

func parseMdFile(text string) (BlogPostData, error) {
	lines := strings.Split(text, "\n")
	post := BlogPostData{}

	post.Title = strings.TrimPrefix(lines[0], "title: ")
	post.Date = strings.TrimPrefix(lines[1], "date: ")

	body := strings.Join(lines[3:], "\n")

	html := markdown.MdToHTML([]byte(body))

	post.Content = string(html)

	return post, nil
}
