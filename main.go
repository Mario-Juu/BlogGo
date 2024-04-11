package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)



var cache map[string] *template.Template

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about")
}

func RenderTemplate(w http.ResponseWriter, page string){

	var t *template.Template
	var err error
	_, exists := cache[page]
	if !exists {
	t, err = template.ParseFiles("templates/" + page + ".page.tmpl", 
									"templates/base.layout.tmpl")
	
	cache[page] = t
} else{
	t = cache[page]
}
if err = t.Execute(w, nil); err != nil {
	log.Println(err)
	return
}
}

func main() {

	cache = make(map[string]*template.Template)

	config := Config{
		Port:    "3000",
		Env:     "dev",
		Version: "1.0.0",
	}

	// app := Application{
	// 	Config: config,
	// 	Cache: cache,
	// }

	http.HandleFunc("/", HomeHandler)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "about")
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))



	log.Printf("Starting server from %s on :%s", config.Env, config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil); err != nil{
		log.Println(err)
	}
}
