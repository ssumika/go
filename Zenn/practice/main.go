package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var dbSessions = map[string]string{} // session ID, user ID
var dbUsers = map[string]user{}      // user ID, user info

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/vip", vip)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	var u user

	// クッキーを取得
	c, err := req.Cookie("session")
	if err != nil {
		u = user{}
	} else {

		// セッションがあればユーザーDB情報を取得
		if un, ok := dbSessions[c.Value]; ok {
			u, ok = dbUsers[un]
			if !ok {
				u = user{}
			}
		}
	}

	fmt.Println(c, u)

	tpl.ExecuteTemplate(w, "index.html", u)
}

func vip(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "vip.html", nil)
}

func signup(w http.ResponseWriter, req *http.Request) {
	// FORMから送信してきたときの処理
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		u := user{un, p, f, l}
		dbUsers[un] = u

		// リダイレクト
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return

	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}
