package addmemo

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"

	"keep-server/handler/user/login"
	"keep-server/model/memo"
	"keep-server/utill/hash"
)

var (
	db = memo.Connect()
)

func Post(c echo.Context) (err error) {

	var json struct {
		Session string   `json:"session"`
		Code    string   `json:"code"`
		Title   string   `json:"title"`
		Text    string   `json:"text"`
		Tag     []string `json:"tag"`
	}

	if err = c.Bind(&json); err != nil {
		return
	}

	if login.IsLogin(c, json.Code, json.Session) == false {
		return c.JSON(http.StatusBadRequest, "not login")
	}

	rand.Seed(time.Now().UnixNano())
	hash := hash.Sha1(strconv.Itoa(rand.Int()))

	tag := ""

	for _, v := range json.Tag {
		tag += v + ","
	}

	db.Create(&memo.Memo{Code: json.Code, Hash: hash, Title: json.Title, Text: json.Text, Tag: tag})

	var result struct {
		Hash string `json:"hash"`
	}

	result.Hash = hash

	return c.JSON(http.StatusOK, result)
}
