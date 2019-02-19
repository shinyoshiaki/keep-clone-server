package getmemo

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"keep-server/handler/user/login"
	"keep-server/model/memo"
)

var (
	db = memo.Connect()
)

func Get(c echo.Context) (err error) {

	var json struct {
		Session string `json:"session"`
		Code    string `json:"code"`
	}

	if err = c.Bind(&json); err != nil {
		return
	}

	if login.IsLogin(c, json.Code, json.Session) == false {
		return c.JSON(http.StatusBadRequest, "not login")
	}

	memos := []memo.Memo{}
	db.Find(&memos, "Code = ?", json.Code)

	fmt.Println(memos)

	var result struct {
		Memos []memo.Memo `json:"memos"`
	}

	result.Memos = memos

	return c.JSON(http.StatusOK, result)
}
