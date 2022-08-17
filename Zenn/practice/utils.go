package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	var u user

	// クッキーを取得
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// セッションがあればユーザーDB情報を取得
	if ds, ok := dbSessions[c.Value]; ok {
		ds.ltime = time.Now()
		dbSessions[c.Value] = ds
		u = dbUsers[ds.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	// クッキーを取得
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	c.MaxAge = sessionLength // 有効期限を延長
	http.SetCookie(w, c)

	// ユーザーDB情報を取得するついでに有効期限を延長
	ds, ok := dbSessions[c.Value]
	if ok {
		ds.ltime = time.Now()
		dbSessions[c.Value] = ds
	}
	_, ok = dbUsers[ds.un] // 情報があればtrue、なければfalse
	return ok
}

func cleanSessions() {
	fmt.Println("***before_action=cleanSessions()") // チェック用
	for k, v := range dbSessions {
		fmt.Println(k, v) // チェック用
		if time.Now().Sub(v.ltime) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	fmt.Println("***after_action=cleanSessions()") // チェック用
	for k, v := range dbSessions {                 // チェック用
		fmt.Println(k, v) // チェック用
	} // チェック用
	fmt.Println("") // チェック用
}
