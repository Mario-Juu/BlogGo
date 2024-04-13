package main

import (
	"flag"
	"html/template"
	"log"
)

var LoginView *View 
var AboutView *View 
var ContactView *View 
var HomeView *View 

func CreateViews(){
	var err error
	LoginView, err = NewView("base", "login")
	if err != nil {
		log.Println(err)
	}
	AboutView, err = NewView("base", "about")
	if err != nil {
		log.Println(err)
	}
	ContactView, err = NewView("base", "contact")
	if err != nil {
		log.Println(err)
	}
	HomeView, err = NewView("base", "index")
	if err != nil {
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
	


	log.Printf("Starting server from %s on :%s", config.Env, config.Port)

	if err := app.Start(); err != nil{
		log.Fatal(err)
	}
}
