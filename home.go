package main

import (
	"html/template"
	"net/http"
	"time"
)

var templatesDesktop = template.Must(template.ParseFiles("./template/home.html"))

type lista struct {
	Rendertime time.Duration
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	lists := lista{}
	templatesDesktop.Execute(w, lists)
}
