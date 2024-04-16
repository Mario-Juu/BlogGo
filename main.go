package main

import (
	"flag"
	"html/template"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

var db *sql.DB

func CreateViews(){
	var err error
	LoginView, err = NewView( "login")
	if err != nil {
		log.Println(err)
	}
	AboutView, err = NewView( "about")
	if err != nil {
		log.Println(err)
	}
	ContactView, err = NewView( "contact")
	if err != nil {
		log.Println(err)
	}
	HomeView, err = NewView( "index")
	if err != nil {
		log.Println(err)
	}
	PostView, err = NewView( "post")
	if err != nil{
		log.Println(err)
	}
	SignUpView, err = NewView("signup")
	if err != nil{
		log.Println(err)
	}

	NewPostView, err = NewView("postnew")
	if err != nil{
		log.Println(err)
	}
}

func main() {
	cache := make(map[string]*template.Template)
	config := Config{Version: "1.0.0",
	}
	flag.StringVar(&config.Port, "port", "8080", "HTTP Server Port")
	flag.StringVar(&config.Env, "env", "dev", "Application Environment")
	flag.Parse()

	 app := Application{
	 	Config: config,
	 	Cache: cache,
	 }
	CreateViews()
	 var err error
	db, err = sql.Open("mysql", "root:secret@/mysql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	
	 

	log.Printf("Starting server from %s on :%s", config.Env, config.Port)

	if err := app.Start(); err != nil{
		log.Fatal(err)
	}
}
