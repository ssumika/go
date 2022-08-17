package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName   string
	Password   []byte
	First      string
	Last       string
	Permission string
}

type session struct {
	un    string
	ltime time.Time
}

//var dbSessions = map[string]string{} // session ID, user ID
var dbSessions = map[string]session{}
var dbUsers = map[string]user{} // user ID, user info

const sessionLength int = 30

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/vip", vip)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	showSessions()
	tpl.ExecuteTemplate(w, "index.html", u) // u をテンプレートに渡す
}

func vip(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Permission != "dog" {
		http.Error(w, "No dogs allowed.", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "vip.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	// すでにログインしている場合はこのページは必要ない
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// FORMから送信してきたときの処理
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		s := req.FormValue("permisson")

		// もしユーザーネームが既に使われていたらエラーにする
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// セッションDBを作成
		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}

		// パスワードを暗号化して保存
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u := user{un, bs, f, l, s}
		dbUsers[un] = u

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	// ログイン済みの場合はこのページは不必要
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// 情報の送信を受け取ったときの処理
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		// ユーザーDBから情報を取ってくる
		u, ok := dbUsers[un]
		// そもそもDBになかったら失敗
		if !ok {
			http.Error(w, "USER NAME - PASSWORD not match", http.StatusForbidden)
			return
		}
		// 受信したパスワードとDBのパスワードの照合
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "USER NAME - PASSWORD not match", http.StatusForbidden)
			return
		}

		// セッションを作成する
		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength // ここを追加した
		http.SetCookie(w, c)
		// dbSessions[c.Value] = un
		dbSessions[c.Value] = session{un, time.Now()}   // ここを修正した
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
		return
	}

	tpl.ExecuteTemplate(w, "login.html", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	cleanSessions()                                      // ここを追加
	http.Redirect(w, req, "/login", http.StatusSeeOther) // 303
}

func showSessions() {
	fmt.Println("***action=showSessions()")
	for k, v := range dbSessions {
		fmt.Println(k, v.un, v.ltime)
	}
	fmt.Println("")
}
