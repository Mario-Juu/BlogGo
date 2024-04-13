package main

import "net/http"

func (a *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.HomeHandler(HomeView))
	mux.HandleFunc("/about", a.AboutHandler(AboutView))
	mux.HandleFunc("/login", a.LoginHandler(LoginView))
	mux.HandleFunc("/contact", a.ContactHandler(ContactView))
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	return mux
}
