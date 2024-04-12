package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"time"
)

//go:embed templates
var TemplateFS embed.FS

type TemplateData struct {
	Email    string
	Telefone string
	Route    string
}


var funcs = template.FuncMap{
	"GetYear": func() int{
		return time.Now().Year()
	},
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

func parseTemplate (page, env string) (*template.Template, error){
	if( env == "dev" ){
		return template.New("").Funcs(funcs).ParseFiles( "templates/"+page+".page.tmpl",
		"templates/navbar.layout.tmpl",
		"templates/base.layout.tmpl")
		
	}
	return template.New("").Funcs(funcs).ParseFS(TemplateFS, "templates/"+page+".page.tmpl",
	"templates/navbar.layout.tmpl",
	"templates/base.layout.tmpl")
	
}