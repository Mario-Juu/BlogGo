package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

//go:embed templates
var TemplateFS embed.FS

var funcs = template.FuncMap{
	"GetYear": func() int {
		return time.Now().Year()
	},
}

type View struct {
	Template *template.Template
	Layout   string
}

func getLayoutFiles() []string{
	files, err := filepath.Glob("templates/*.layout.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func NewView(layout string, pages ...string) (*View, error) {
	files := getLayoutFiles()
	for _, p := range pages {
		files = append(files, fmt.Sprintf("templates/%s.page.tmpl", p))
	}
	t, err := template.New("").Funcs(funcs).ParseFiles(files...)
	if err != nil {
		return nil, err
	}
	return &View{Template: t, Layout: layout}, nil
}

func (v *View) Render(w http.ResponseWriter, data any) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}



type TemplateData struct {
	Email    string
	Telefone string
	Route    string
}

func (a *Application) RenderTemplate(w http.ResponseWriter, page string, data any) error {

	var t *template.Template
	var err error
	_, exists := a.Cache[page]
	if !exists || a.Config.Env == "dev" {
		t, err = parseTemplate(page, a.Config.Env)
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			return err
		}
		a.Cache[page] = t
	} else {
		t = a.Cache[page]
	}

	return t.ExecuteTemplate(w, "base", data)
}

func parseTemplate(page, env string) (*template.Template, error) {
	if env == "dev" {
		return template.New("").Funcs(funcs).ParseFiles("templates/"+page+".page.tmpl",
			"templates/navbar.layout.tmpl",
			"templates/base.layout.tmpl")

	}
	return template.New("").Funcs(funcs).ParseFS(TemplateFS, "templates/"+page+".page.tmpl",
		"templates/navbar.layout.tmpl",
		"templates/base.layout.tmpl")

}
