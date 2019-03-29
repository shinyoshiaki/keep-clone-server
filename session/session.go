package session

import (
	"fmt"
	"keep-server/utill/hash"
	"math/rand"
	"strconv"
	"time"
)

type (
	session struct {
		time time.Time
		code string
	}
)

var sessions = map[string]*session{}

func Set(token string, code string) {
	s := &session{
		time: time.Now().Add(time.Hour * 1),
		code: code,
	}
	sessions[token] = s
}

func Get(token string) string {
	s := sessions[token]
	if s == nil {
		fmt.Println("session null")
		return ""
	}
	if time.Now().Unix() > s.time.Unix() {
		fmt.Println("session timeout")
		return ""
	}
	return s.code
}

func GenSession(code string) string {
	rand.Seed(time.Now().UnixNano())
	sessionKey := hash.Sha1(strconv.Itoa(rand.Int()))
	Set(sessionKey, code)

	return sessionKey
}

func IsLogin(token string) string {
	code := Get(token)

	if code != "" {
		return code
	}
	return ""
}
