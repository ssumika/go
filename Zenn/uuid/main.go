package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

// 個人情報の定義
type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user infomation
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/vip", vip)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	// クッキーを読み込む。なかったら作る
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// もしすでにuserIDがsessionDBにあったらそのuserIDとuser情報を取ってくる
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// 入力フォームから情報を受け取ったらそのデータをsessionDBとuserDBに登録する
	// 入力フォームから情報を受け取ったら = RequestがPOSTだったら
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func vip(w http.ResponseWriter, req *http.Request) {

	// クッキーを取ってくる。なければルートにリダイレクトする
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// sessionDBからuserIDを取ってくる。なければルートにリダイレクトする
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// userDBからuser情報を取ってくる
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "vip.gohtml", u)
}
