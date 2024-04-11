package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates
var TemplateFS embed.FS

func (a *Application) RenderTemplate(w http.ResponseWriter, page string) {

	var t *template.Template
	var err error
	_, exists := a.Cache[page]
	if !exists || a.Config.Env == "dev" {
		t, _ = template.ParseFS(TemplateFS, "templates/"+page+".page.tmpl",
			"templates/navbar.layout.tmpl",
			"templates/base.layout.tmpl")

		a.Cache[page] = t
	} else {
		t = a.Cache[page]
	}

	type Contato struct {
		Email    string
		Telefone string
	}
	contato := Contato{
		Email:    "contato@localhost",
		Telefone: "123456789",
	}

	if err = t.Execute(w, contato); err != nil {
		log.Println(err)
		return
	}
}
