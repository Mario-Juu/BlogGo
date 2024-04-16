package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type TemplateData struct {
	Email    string
	Telefone string
	Route    string
	Errors   []string
	User     *SessionUser
	Posts []Post
}

type SessionUser struct {
	Email string
}

func getUserByCookie(r *http.Request) *SessionUser {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Println(err)
		return nil
	}
	if cookie != nil {
		return &SessionUser{Email: cookie.Value}
	}
	return nil
}

func (a *Application) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserByCookie(r)
		if user == nil {
			log.Println("User not logged in. Redirect to login page...")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		log.Println("User logged in. Continue...")
		next(w, r)
	}
}

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
		if r.Method == http.MethodGet {
			err := view.Render(w, r, nil)
			if err != nil {
				log.Println(err)
			}
		} else if r.Method == http.MethodPost {
			var data struct {
				Email    string `json:"email"`
				Password string `json:"password"`
				Success  bool   `json:"success"`
			}
			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				log.Println(err)
			}
			var user *User
			user, err = FindUserByEmail(data.Email)
			if err != nil {
				json.NewEncoder(w).Encode(data)
				return
			}

			if data.Password == user.Password {
				data.Success = true

				cookie := http.Cookie{Name: "session", Value: data.Email, Expires: time.Now().Add(time.Minute), HttpOnly: true}
				http.SetCookie(w, &cookie)
				json.NewEncoder(w).Encode(data)
				return
			} else {

				json.NewEncoder(w).Encode(data)
				return
			}
		}
	}
}
func (a *Application) ContactHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := view.Render(w, r, &TemplateData{Email: "teste@localhost", Telefone: "99999"})
		if err != nil {
			log.Println(err)
		}
	}
}
func (a *Application) PostHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts := ReadPosts()
		err := view.Render(w, r, &TemplateData{Posts: posts})
		if err != nil {
			log.Println(err)
		}
	}
}

func (a *Application) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func (a *Application) SignUpHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			err := view.Render(w, r, nil)
			if err != nil {
				log.Println(err)
			}
		} else if r.Method == http.MethodPost {
			var data struct {
				Email    string `json:"email"`
				Password string `json:"password"`
				Success  bool   `json:"success"`
			}
			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				log.Println(err)
			}

			stmt, err := db.Prepare("INSERT INTO users (email, password) values (?, ?)")
			if err != nil {
				json.NewEncoder(w).Encode(data)
				return
			}
			_, err = stmt.Exec(data.Email, data.Password)
			if err != nil {
				json.NewEncoder(w).Encode(data)
				return
			}
			data.Success = true
			cookie := http.Cookie{Name: "session", Value: data.Email, Expires: time.Now().Add(time.Minute), HttpOnly: true}
			http.SetCookie(w, &cookie)
			json.NewEncoder(w).Encode(data)
			return
		}
	}
}
func (a *Application) NewPostHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			err := view.Render(w, r, nil)
			if err != nil {
				log.Println(err)
			}
		} else if r.Method == http.MethodPost {
			var data struct {
				Title   string `json:"title"`
				Content string `json:"content"`
				Success bool   `json:"success"`
				Errors []string `json:"errors"`
			}

			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil{
				log.Println(err)
				return
			}
			title := data.Title
			content := data.Content
			errors := make([]string, 0)
			userDTO := getUserByCookie(r)
			user, err := FindUserByEmail(userDTO.Email)
			if err != nil {
				errors = append(errors, "You are not logged in")
			}
			if title == ""{
				errors = append(errors, "Title is required")
			}

			if content == ""{
				errors = append(errors, "Content is required")
			}
			if len(errors) > 0 {
				data.Errors = errors
				json.NewEncoder(w).Encode(data)
				return
			}

			post := Post{
				Title:   title,
				Content: content,
				Author:  user,
				Slug:    slugify(title),
			}
			err = CreatePost(post)
			if err != nil {
				log.Println(err)
			}
			json.NewEncoder(w).Encode(data)

		}
	}
}

func slugify(value string) string {
	value = strings.ToLower(value)
	reg := regexp.MustCompile("[^a-z0-9]+")
	value = reg.ReplaceAllString(value, "-")
	value = strings.Trim(value, "-")
	return value

}
