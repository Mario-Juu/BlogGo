package main

import "net/http"

func (a *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.HomeHandler(HomeView))
	mux.HandleFunc("/about", a.AboutHandler(AboutView))
	mux.HandleFunc("/login", a.LoginHandler(LoginView))
	mux.HandleFunc("/contact", a.ContactHandler(ContactView))
	mux.HandleFunc("/post", a.AuthMiddleware(a.PostHandler(PostView)))
	mux.HandleFunc("/signup", a.SignUpHandler(SignUpView))
	mux.HandleFunc("/post/viewer", a.PostViewerHandler(PostViewerView))
	mux.HandleFunc("/logout", a.LogoutHandler)
	mux.HandleFunc("/post/new", a.AuthMiddleware(a.NewPostHandler(NewPostView)))
	mux.HandleFunc("/post/edit", a.AuthMiddleware(a.EditPostHandler(EditPostView)))
	mux.HandleFunc("/post/delete", a.AuthMiddleware(DeletePostHandler))
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	return mux
}
