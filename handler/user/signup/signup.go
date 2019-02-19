package signup

import (
	"fmt"
	"net/http"
	"strconv"

	"keep-server/model/user"
	"keep-server/utill/hash"

	"github.com/labstack/echo"
)

var (
	db = user.Connect()
)

// SignUp signup
func SignUp(c echo.Context) (err error) {

	var json struct {
		Name     string `json:"name"`
		Password string `json:"pass"`
	}

	if err = c.Bind(&json); err != nil {
		return
	}

	if json.Name == "" || json.Password == "" {
		return c.String(http.StatusBadRequest, "null")
	}

	var u user.User
	db.Find(&u, "Name = ?", json.Name)

	if u.Name != "" {
		return c.String(http.StatusBadRequest, "exist")
	}

	count := 0
	users := []user.User{}
	db.Find(&users).Count(&count)

	code := hash.Sha1(strconv.Itoa(count))
	fmt.Println(count, code)
	pass := hash.Sha1(json.Password)

	db.Create(&user.User{Name: json.Name, Password: pass, Key: count, Code: code})
	var result struct {
		Name string `json:"name"`
		Code string `json:"code"`
		Pass string `json:"pass"`
	}
	result.Name = json.Name
	result.Code = code
	result.Pass = json.Password

	return c.JSON(http.StatusOK, result)
}
