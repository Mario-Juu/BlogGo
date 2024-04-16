package main

import (
	"log"
	"time"
)

type User struct {
	Id       int
	Email    string
	Password string
}

type Post struct {
	Id        int
	Title     string
	Content   string
	Slug      string
	Author    *User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	row := db.QueryRow("SELECT id, email, password FROM users WHERE email =?", email)
	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func CreatePost(post Post) error {
	stmt, err := db.Prepare("INSERT INTO posts (title, content, slug, user_id) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post.Title, post.Content, post.Slug, post.Author.Id)
	if err != nil {
		return err
	}
	return nil
}

func ReadPosts() []Post {
	posts := make([]Post, 0)
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		var user User
		var post Post
		err = rows.Scan(&post.Id, &post.Title, &post.Slug, &post.Content, &user.Id, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
		post.Author = &user
		posts = append(posts, post)
	}
	return posts
}
