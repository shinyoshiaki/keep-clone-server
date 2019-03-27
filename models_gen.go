// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package keep_server

type AllMemo struct {
	Memos []Memo `json:"memos"`
}

type GetAllMemo struct {
	Owner string `json:"owner"`
	Token string `json:"token"`
}

type GetUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Memo struct {
	Owner string `json:"owner"`
	Code  string `json:"code"`
	Hash  string `json:"hash"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Tag   string `json:"tag"`
}

type NewMemo struct {
	Owner string   `json:"owner"`
	Token string   `json:"token"`
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Tag   []string `json:"tag"`
}

type NewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type User struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
