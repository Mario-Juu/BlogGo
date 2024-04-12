package main

import (
	"log"
	"net/http"
)

func (a *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := a.RenderTemplate(w, "index", TemplateData{Route: "index"})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (a *Application) AboutHandler(w http.ResponseWriter, r *http.Request) {
	err := a.RenderTemplate(w, "about", TemplateData{Route: "about"})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (a *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := a.RenderTemplate(w, "login", TemplateData{Route: "login"})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (a *Application) ContactHandler(w http.ResponseWriter, r *http.Request) {
	err := a.RenderTemplate(w, "contact", TemplateData{Email: "mario@gmail.com",
		Telefone: "123456789",
		Route:    "contact"})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}

}
