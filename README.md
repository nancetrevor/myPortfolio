# Trevor Nance Portfolio

This is a static site. The Go program is used only at build time to render the
Markdown blog posts and HTML templates; Cloudflare Pages serves the generated
files and does not run a Go web server.

## Build locally

```sh
go run ./cmd/build
```

The deployable site is written to `dist/`. To preview it, serve that directory
with any static-file server, for example:

```sh
cd dist && python3 -m http.server 8080
```

## Deploy with Cloudflare Pages

Create a Pages project from this repository and set:

- Build command: `go run ./cmd/build`
- Build output directory: `dist`
- Root directory: `/`

Each push will rebuild the static pages. Add blog posts as Markdown files in
`content/blog/`, starting with `title:` and `date:` lines followed by a blank
line. Their URLs use a lowercase, hyphenated version of the filename.
