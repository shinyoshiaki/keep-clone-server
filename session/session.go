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
		key  string
		time time.Time
	}
)

var sessions = map[string]*session{}

func Set(id string, key string) {
	s := &session{
		key:  key,
		time: time.Now().Add(time.Hour * 1),
	}
	sessions[id] = s
}

func Get(id string) string {
	s := sessions[id]
	if s == nil {
		fmt.Println("session null")
		return ""
	}
	if time.Now().Unix() > s.time.Unix() {
		fmt.Println("session timeout")
		return ""
	}
	return s.key
}

func GenSession(code string) string {
	rand.Seed(time.Now().UnixNano())
	sessionKey := hash.Sha1(strconv.Itoa(rand.Int()))
	Set(code, sessionKey)

	return sessionKey
}

func IsLogin(code string, token string) bool {
	sessionKey := Get(code)
	fmt.Println("islogin:", sessionKey, token)

	if sessionKey == token {
		return true
	}
	return false
}
