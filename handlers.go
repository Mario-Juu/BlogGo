package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type TemplateData struct {
	Email    string
	Telefone string
	Route    string
	Errors   []string
	User     *SessionUser
	Posts    []Post
	Post     Post
}

type SessionUser struct {
	Email string
}

func getUserByCookie(r *http.Request) *SessionUser {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}
	if cookie != nil {
		return &SessionUser{Email: cookie.Value}
	}
	return nil
}


func (a *Application) AuthUserEditMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserByCookie(r)
		if user == nil {
			log.Println("User not logged in. Redirect to login page...")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		next(w, r)
	}
}

func (a *Application) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserByCookie(r)
		if user == nil {
			log.Println("User not logged in. Redirect to login page...")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		next(w, r)
	}
}

func (a *Application) HomeHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts := ReadPosts()
		err := view.Render(w, r, &TemplateData{Posts: posts})
		if err != nil {
			log.Println(err)
		}
	}
}

func (a *Application) PostViewerHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		postId, _ := strconv.Atoi(id)
		post, err := GetPostById(postId)
		if err != nil {
			log.Println(err)
		}
		err = view.Render(w, r, &TemplateData{Post: *post})
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

				cookie := http.Cookie{Name: "session", Value: data.Email, Expires: time.Now().Add(time.Minute * 15), HttpOnly: true}
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
		errorMsg := r.URL.Query().Get("error")
		errors := []string{}
		if errorMsg != "" {
			errors = append(errors, errorMsg)
		}
		userEmail := getUserByCookie(r).Email
		user, err := FindUserByEmail(userEmail)
		if err != nil{
			log.Println(err)
		}
		posts := ReadPostsById(user.Id)
		err = view.Render(w, r, &TemplateData{Posts: posts, Errors: errors})
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
				log.Println(err)
				json.NewEncoder(w).Encode(data)
				return
			}
			_, err = stmt.Exec(data.Email, data.Password)
			if err != nil {
				log.Println(err)
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
				Title   string   `json:"title"`
				Content string   `json:"content"`
				Success bool     `json:"success"`
				Errors  []string `json:"errors"`
			}

			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
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
			if title == "" {
				errors = append(errors, "Title is required")
			}

			if content == "" {
				errors = append(errors, "Content is required")
			}
			if len(errors) > 0 {
				data.Errors = errors
				json.NewEncoder(w).Encode(data)
				return
			}

			post := Post{
				Title:   title,
				Content: template.HTML(content),
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

func (a *Application) EditPostHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postId := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(postId)

		if r.Method == http.MethodGet {
			post, err := GetPostById(id)
			if err != nil {
				log.Println(err)
				view.Render(w, r, &TemplateData{Post: *post, Errors: []string{err.Error()}})
				return
			}
			err = view.Render(w, r, &TemplateData{Post: *post})
			if err != nil {
				log.Println(err)
			}
		} else if r.Method == http.MethodPost {
			postGet, err := GetPostById(id)
			if err != nil {
				log.Println(err)
				view.Render(w, r, &TemplateData{Post: *postGet, Errors: []string{err.Error()}})
				return
			}
			title := r.FormValue("title")
			content := r.FormValue("content")
			if title == "" {
				title = postGet.Title
			}
			if content == "" {
				content = string(postGet.Content)
			}
			post := Post{
				Id:        id,
				Title:     title,
				Content:   template.HTML(content),
				Slug:      slugify(title),
				UpdatedAt: time.Now(),
			}
			err = UpdatePost(post)
			if err != nil {
				view.Render(w, r, &TemplateData{Post: post, Errors: []string{err.Error()}})
				return
			}
			http.Redirect(w, r, "/post", http.StatusSeeOther)
		}
	}
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	postId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(postId)
	err := DeletePost(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, fmt.Sprintf("/post?error=%s", err.Error()), http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/post", http.StatusSeeOther)

}

func slugify(value string) string {
	value = strings.ToLower(value)
	reg := regexp.MustCompile("[^a-z0-9]+")
	value = reg.ReplaceAllString(value, "-")
	value = strings.Trim(value, "-")
	return value

}
