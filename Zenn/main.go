package main

import (
	"html/template"
	"os"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	type source struct {
		Title, Message string
		Time           time.Time
	}

	nf, _ := os.Create("index.html")
	defer nf.Close()

	var data source
	data = source{
		Title:   "test page",
		Message: "本日は晴天なり",
		Time:    time.Now(),
	}
	_ = tpl.ExecuteTemplate(nf, "tpl.gohtml", data)
}
