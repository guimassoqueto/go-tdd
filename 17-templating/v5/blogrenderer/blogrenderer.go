package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

type PostViewModel struct {
	Title, SanitizedTitle, Description, Body string
	Tags []string
}

func (p *Post) SanitizedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}
	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.SanitizedTitle}}">{{.Title}}</a></li>{{end}}</ol>`
	templ, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return err
	}
	if err := templ.Execute(w, posts); err != nil {
		return err
	}
	return nil
}