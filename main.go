package main

import (
	"flag"
	"html/template"
	"log"
)

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


	log.Printf("Starting server from %s on :%s", config.Env, config.Port)

	if err := app.Start(); err != nil{
		log.Fatal(err)
	}
}
