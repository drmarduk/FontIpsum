package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", fontHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fontHandler(w http.ResponseWriter, r *http.Request) {
	font := r.FormValue("font")
	if font == "" {
		font = "arial"
	}

	font = getFontName(font)

	tmpl, err := template.ParseFiles("html/index.tpl")
	if err != nil {
		log.Println(err.Error())
		return
	}
	tmpl.Execute(w, font)
}

func getFontName(f string) string {
	switch f {
	case "rambla":
		return "'Rambla', sans-serif"
	case "arial":
		return "'Arial', sans-serif"
	case "sourcepro":
		return "'Source Code Pro,"
	}
	return ""
}
