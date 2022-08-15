package main

import "net/http"

func getUser(req *http.Request) user {
	var u user

	// クッキーを取得
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// セッションがあればユーザーDB情報を取得
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	// クッキーを取得
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	// ユーザーDB情報を取得
	un := dbSessions[c.Value]
	_, ok := dbUsers[un] // 情報があればtrue、なければfalse
	return ok
}
