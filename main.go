package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
)

var env = "dev"
var cache = make(map[string]*template.Template)

var LoginView *View
var AboutView *View
var ContactView *View
var HomeView *View
var PostView *View

var SignUpView *View
var NewPostView *View
var EditPostView *View
var PostViewerView *View

var db *sql.DB

func CreateViews() {
	var err error
	LoginView, err = NewView("login")
	if err != nil {
		log.Println(err)
	}
	AboutView, err = NewView("about")
	if err != nil {
		log.Println(err)
	}
	ContactView, err = NewView("contact")
	if err != nil {
		log.Println(err)
	}
	HomeView, err = NewView("index")
	if err != nil {
		log.Println(err)
	}
	PostView, err = NewView("post")
	if err != nil {
		log.Println(err)
	}
	SignUpView, err = NewView("signup")
	if err != nil {
		log.Println(err)
	}

	NewPostView, err = NewView("postnew")
	if err != nil {
		log.Println(err)
	}
	EditPostView, err = NewView("postedit")
	if err != nil {
		log.Println(err)
	}
	PostViewerView, err = NewView("postviewer")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	cache := make(map[string]*template.Template)
	config := Config{Version: "1.0.0"}
	flag.StringVar(&config.Port, "port", "8080", "HTTP Server Port")
	flag.StringVar(&config.Env, "env", "dev", "Application Environment")
	flag.Parse()

	app := Application{
		Config: config,
		Cache:  cache,
	}
	CreateViews()
	var err error
	db, err = sql.Open("mysql", "root:secret@/mysql?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	initTables()

	log.Printf("Starting server from %s on :%s", config.Env, config.Port)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func initTables() {
	log.Println("Criando as tabelas")
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id int NOT NULL AUTO_INCREMENT,
		email varchar(255) UNIQUE,
		password varchar(255),
		PRIMARY KEY (id)
	);`)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts(
		id int NOT NULL AUTO_INCREMENT,
		title varchar(255) NOT NULL,
		slug varchar (255) NOT NULL UNIQUE,
		content text,
		user_id int NOT NULL,
		created_at timestamp DEFAULT CURRENT_TIMESTAMP(),
		updated_at timestamp DEFAULT CURRENT_TIMESTAMP(),
		PRIMARY KEY (id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

}
