package model

type AllMemo struct {
	Memos []Memo `json:"memos"`
}

type Auth struct {
	Token string `json:"token"`
}

type EditMemo struct {
	Token    string   `json:"token"`
	MemoCode string   `json:"memoCode"`
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	Tag      []string `json:"tag"`
}

type GetAllMemo struct {
	Token string `json:"token"`
}

type GetUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Memo struct {
	Owner string `json:"owner"`
	Code  string `json:"code"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Tag   string `json:"tag"`
	Time  string `json:"time"`
}

type NewMemo struct {
	Token string   `json:"token"`
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Tag   []string `json:"tag"`
}

type NewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RemoveMemo struct {
	Token    string `json:"token"`
	MemoCode string `json:"memoCode"`
}

type User struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
