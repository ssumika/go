package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type dog struct {
	Kind    string
	Country string
}

type vegetable struct {
	Name  string
	Color string
}

var funcmap = template.FuncMap{
	"upper": strings.ToUpper,
	"three": myThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcmap).ParseFiles("tpl.gohtml"))
}

func myThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 3 {
		s = s[:3]
	}
	return s
}

func main() {

	cute := dog{
		Kind:    "Samoyed",
		Country: "Russia",
	}
	cute2 := dog{
		Kind:    "Akita",
		Country: "Japan",
	}

	v1 := vegetable{
		Name:  "onion",
		Color: "white",
	}
	v2 := vegetable{
		Name:  "carrot",
		Color: "orange",
	}
	dogs := []dog{cute, cute2}
	vegetables := []vegetable{v1, v2}

	data := struct {
		AAA []dog
		BBB []vegetable
	}{
		dogs,
		vegetables,
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
