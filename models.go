package main

import (
	"errors"
	"html/template"
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
	Content   template.HTML
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
	posts := []Post{}
	rows, err := db.Query(`select p.id, p.title, p.slug, p.content, p.user_id, u.email, p.created_at, p.updated_at 
							from posts p join users u on p.user_id = u.id`)
	if err != nil {
		log.Println(err)
		return posts
	}

	for rows.Next() {
		var post Post
		var user User
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Slug,
			&post.Content,
			&user.Id,
			&user.Email,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return posts
		}
		post.Author = &user
		posts = append(posts, post)
	}

	return posts
}

func ReadPostsById(id int) []Post {
	posts := []Post{}
	stmt, err := db.Prepare(`select p.id, p.title, p.slug, p.content, p.user_id, u.email, p.created_at, p.updated_at 
							from posts p join users u on p.user_id = u.id where p.user_id = ?`)
	if err != nil {
		log.Println(err)
		return posts
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println(err)
		return posts
	}
	for rows.Next() {
		var post Post
		var user User
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Slug,
			&post.Content,
			&user.Id,
			&user.Email,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return posts
		}
		post.Author = &user
		posts = append(posts, post)
	}

	return posts
}


func GetPostById(id int) (*Post, error) {
	row := db.QueryRow("SELECT p.id, p.title, p.slug, p.content, p.user_id, u.email, p.created_at, p.updated_at FROM posts p JOIN users u ON p.user_id = u.id WHERE p.id = ?", id)

	var post Post
	var user User
	err := row.Scan(
			&post.Id,
			&post.Title,
			&post.Slug,
			&post.Content,
			&user.Id,
			&user.Email,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return &post, err
		}
		post.Author = &user
	return &post, nil
}

func UpdatePost(post Post) error {
	stmt, err := db.Prepare("UPDATE posts SET title = ?, content = ?, slug = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post.Title, post.Content, post.Slug, post.UpdatedAt, post.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(id int) error {
	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("no rows affected")
	}
	if err != nil {
		return err
	}
	return nil
}
