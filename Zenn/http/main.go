package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	Pname string
	Plike bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	n := req.FormValue("name")
	l := req.FormValue("like") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{n, l})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		log.Fatalln(err)
	}
}
