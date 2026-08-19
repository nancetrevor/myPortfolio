// Command build renders the portfolio into static files for deployment.
package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/lfizzikz/portfolioSite/package/markdown"
)

const (
	outputDir    = "dist"
	templatesDir = "web/templates"
	postsDir     = "content/blog"
)

type blogPage struct {
	Title       string
	Description string
	BlogPosts   []blogPost
}

type blogPost struct {
	Title   string
	Date    string
	Slug    string
	Content template.HTML
}

type blogPostPage struct {
	Title string
	Post  blogPost
}

func main() {
	if err := build(); err != nil {
		fmt.Fprintln(os.Stderr, "static build failed:", err)
		os.Exit(1)
	}
}

func build() error {
	if err := os.RemoveAll(outputDir); err != nil {
		return err
	}
	if err := copyDir("web/static", filepath.Join(outputDir, "static")); err != nil {
		return fmt.Errorf("copy static assets: %w", err)
	}
	if err := copyFile("web/_redirects", filepath.Join(outputDir, "_redirects")); err != nil {
		return fmt.Errorf("copy Cloudflare redirects: %w", err)
	}
	if err := render("index.html", filepath.Join(outputDir, "index.html"), struct{ Title string }{"Home"}); err != nil {
		return err
	}
	if err := render("mma_architecture.html", filepath.Join(outputDir, "mma_architecture", "index.html"), struct{ Title string }{"MMA Architecture"}); err != nil {
		return err
	}

	posts, err := readPosts()
	if err != nil {
		return err
	}
	page := blogPage{
		Title:       "Trevor's Blog",
		Description: "Writing short blogs regarding the things I am learning in my tech journey",
		BlogPosts:   posts,
	}
	if err := render("blog.html", filepath.Join(outputDir, "blog", "index.html"), page); err != nil {
		return err
	}
	for _, post := range posts {
		if err := render("blog_post.html", filepath.Join(outputDir, "blog", post.Slug, "index.html"), blogPostPage{Title: post.Title, Post: post}); err != nil {
			return err
		}
	}
	return nil
}

func render(templateName, outputPath string, data any) error {
	t, err := template.ParseFiles(filepath.Join(templatesDir, templateName))
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return err
	}
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, data)
}

func readPosts() ([]blogPost, error) {
	entries, err := os.ReadDir(postsDir)
	if err != nil {
		return nil, err
	}
	posts := make([]blogPost, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}
		contents, err := os.ReadFile(filepath.Join(postsDir, entry.Name()))
		if err != nil {
			return nil, err
		}
		post, err := parsePost(strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())), string(contents))
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	sort.Slice(posts, func(i, j int) bool { return posts[i].Date > posts[j].Date })
	return posts, nil
}

func parsePost(filename, source string) (blogPost, error) {
	lines := strings.Split(source, "\n")
	post := blogPost{Slug: slugify(filename)}
	if len(lines) < 3 || !strings.HasPrefix(lines[0], "title: ") || !strings.HasPrefix(lines[1], "date: ") {
		return post, fmt.Errorf("%s must begin with title and date metadata", filename)
	}
	post.Title = strings.TrimSpace(strings.TrimPrefix(lines[0], "title: "))
	post.Date = strings.TrimSpace(strings.TrimPrefix(lines[1], "date: "))
	post.Content = template.HTML(markdown.MdToHTML([]byte(strings.Join(lines[3:], "\n"))))
	return post, nil
}

func slugify(value string) string {
	var out strings.Builder
	lastDash := false
	for _, r := range strings.ToLower(value) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			out.WriteRune(r)
			lastDash = false
		} else if !lastDash {
			out.WriteByte('-')
			lastDash = true
		}
	}
	return strings.Trim(out.String(), "-")
}

func copyDir(source, destination string) error {
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		target := filepath.Join(destination, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		out, err := os.Create(target)
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, in)
		return err
	})
}

func copyFile(source, destination string) error {
	in, err := os.Open(source)
	if err != nil {
		return err
	}
	defer in.Close()
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return err
	}
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}
