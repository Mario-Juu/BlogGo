package main

import "net/http"

func (a *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplate(w, "index")
}

func (a *Application) AboutHandler(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplate(w, "about")
}

func (a *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplate(w, "login")
}

func (a *Application) ContactHandler(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplate(w, "contact")
}
