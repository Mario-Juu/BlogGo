package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TemplateData struct{
	Email string
	Telefone string
	Route string
	Errors []error
	User *User
}

type User struct{
	Email string
}



func getUserByCookie(r *http.Request) *User{
	cookie, err := r.Cookie("session")
	if err != nil{
		log.Println(err)
		return nil
	}
	if cookie != nil{
		return &User{Email: cookie.Value}
	}
	return nil
}

func (a *Application) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserByCookie(r)
		if user == nil{
			log.Println("User not logged in. Redirect to login page...")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		log.Println("User logged in. Continue...")
		next(w, r)
	}}

func (a *Application) HomeHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	err := view.Render(w, r, nil)
	if err != nil {
		log.Println(err)
	}
}
}

func (a *Application) AboutHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	
	err := view.Render(w, r, nil)
	if err != nil {
		log.Println(err)
	}
}
}

func (a *Application) LoginHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet{
		err := view.Render(w, r, nil)
		if err != nil {
			log.Println(err)
		}
	} else if r.Method == http.MethodPost{
		var data struct{
			Email string `json:"email"`
			Password string `json:"password"`
			Success bool `json:"success"`
		}
		err := json.NewDecoder(r.Body).Decode(&data) 
		if err != nil {
			log.Println(err)
		}
		if data.Password == "1234"{
			data.Success = true

			cookie := http.Cookie{Name: "session", Value: data.Email, Expires: time.Now().Add(time.Minute), HttpOnly: true,}
			http.SetCookie(w, &cookie)
			json.NewEncoder(w).Encode(data)
			return
		} else{
			json.NewEncoder(w).Encode(data)
			return
		}
	}
}
}
func (a *Application) ContactHandler (view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	err := view.Render(w, r, &TemplateData{Email: "teste@localhost", Telefone: "99999"})
	if err != nil {
		log.Println(err)
	}
}
}
func (a *Application) PostHandler (view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	log.Print("renderizando post")
	err := view.Render(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	}  
}


func (a *Application) LogoutHandler (w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name: "session",
		Value: "",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
