package config

import "html/template"

type Config struct {
	Port    string
	Env     string
	Version string
}

type Application struct {
	Config Config
	Cache  map[string]*template.Template
}