package main

import "net/http"

func (a *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.HomeHandler)
	mux.HandleFunc("/about", a.AboutHandler)
	mux.HandleFunc("/login", a.LoginHandler)
	mux.HandleFunc("/contact", a.ContactHandler)
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	return mux
}
