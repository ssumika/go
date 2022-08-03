package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `<h1><a href="/set">クッキーをセット</a></h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "test",
		Value: "some value",
		Path:  "/",
	})
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("test")
	if err != nil {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
		http.Redirect(w, req, "/set", http.StatusSeeOther) // 303 // ここを変更
		return
	}

	fmt.Fprintln(w, "TEST COOKIE:", c)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("test")
	if err != nil {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
		http.Redirect(w, req, "/set", http.StatusSeeOther) // 303 // ここを変更
		return
	}
	c.MaxAge = -1 // クッキーの削除
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther) // 303
}
