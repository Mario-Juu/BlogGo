package main

import (
	"log"
	"net/http"
)

func (a *Application) HomeHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	err := view.Render(w, TemplateData{Route: "index"})
	if err != nil {
		log.Println(err)
	}
}
}

func (a *Application) AboutHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	err := view.Render(w, TemplateData{Route: "about"})
	if err != nil {
		log.Println(err)
	}
}
}

func (a *Application) LoginHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	err := view.Render(w, TemplateData{Route: "login"})
	if err != nil {
		log.Println(err)
	}
}
}
func (a *Application) ContactHandler (view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	err := view.Render(w, TemplateData{Route: "contact"})
	if err != nil {
		log.Println(err)
	}
}
}
