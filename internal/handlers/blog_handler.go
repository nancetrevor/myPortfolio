package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/lfizzikz/portfolioSite/package/markdown"
)

type BlogPage struct {
	Title       string
	Description string
	BlogPosts   []BlogPostData
}

type BlogPostData struct {
	Title   string
	Date    string
	Slug    string
	Content template.HTML
}

type BlogPostPage struct {
	Title string
	Post  BlogPostData
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	posts := parseMdToBlogPosts()
	path := strings.TrimPrefix(r.URL.Path, "/blog")
	path = strings.Trim(path, "/")

	if path == "" {
		t, err := template.ParseFiles("web/templates/blog.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}

		if err := t.Execute(w, posts); err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
		}
		return
	}

	if strings.Contains(path, "/") {
		http.NotFound(w, r)
		return
	}

	for _, post := range posts.BlogPosts {
		if post.Slug != path {
			continue
		}

		t, err := template.ParseFiles("web/templates/blog_post.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}

		page := BlogPostPage{
			Title: post.Title,
			Post:  post,
		}

		if err := t.Execute(w, page); err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
		}
		return
	}

	http.NotFound(w, r)
}

func parseMdToBlogPosts() BlogPage {
	dir := "content/blog"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("error: ", err)
		return BlogPage{}
	}

	page := BlogPage{
		Title:       "Trevor's Blog",
		Description: "Writing short blogs regarding the things I am learning in my tech journey",
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		path := filepath.Join(dir, file.Name())
		contents, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("error: ", err)
			return BlogPage{}
		}

		post, err := parseMdFile(file.Name(), string(contents))
		if err != nil {
			fmt.Println("error: ", err)
			return BlogPage{}
		}
		page.BlogPosts = append(page.BlogPosts, post)
	}

	sort.Slice(page.BlogPosts, func(i, j int) bool {
		return page.BlogPosts[i].Date > page.BlogPosts[j].Date
	})
	return page
}

func parseMdFile(filename string, text string) (BlogPostData, error) {
	lines := strings.Split(text, "\n")
	post := BlogPostData{
		Slug: strings.TrimSuffix(filename, filepath.Ext(filename)),
	}

	if len(lines) < 3 {
		post.Content = template.HTML(markdown.MdToHTML([]byte(text)))
		return post, nil
	}

	post.Title = strings.ToTitle(strings.TrimPrefix(lines[0], "title: "))
	post.Date = strings.TrimPrefix(lines[1], "date: ")

	body := strings.Join(lines[3:], "\n")

	html := markdown.MdToHTML([]byte(body))

	post.Content = template.HTML(html)

	return post, nil
}
