package keep_server

import (
	context "context"
	"keep-server/model/memo"
	"keep-server/model/user"
	"keep-server/session"
	"keep-server/utill/hash"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var (
	userDB = user.Connect()
	memoDB = memo.Connect()
)

type Resolver struct {
	users []User
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	var u User
	userDB.Find(u, "Name = ?", input.Name)
	if u.Name != "" {
		return nil, nil
	}

	count := 0
	users := []user.User{}
	userDB.Find(&users).Count(&count)
	code := hash.Sha1(strconv.Itoa(count))
	pass := hash.Sha1(input.Password)

	rand.Seed(time.Now().UnixNano())

	sessionKey := session.GenSession(code)

	user := &User{
		Code:     code,
		Name:     input.Name,
		Password: pass,
		Token:    sessionKey,
	}

	userDB.Create(user)

	return user, nil
}

func (r *mutationResolver) CreateMemo(ctx context.Context, input NewMemo) (*Memo, error) {
	owner := session.IsLogin(input.Token)
	if owner == "" {
		return nil, nil
	}

	rand.Seed(time.Now().UnixNano())
	code := hash.Sha1(strconv.Itoa(rand.Int()))

	tag := ""
	for _, v := range input.Tag {
		tag += v + ","
	}

	hash := hash.Sha1(input.Title + input.Text + tag)

	memo := &Memo{
		Owner: owner,
		Code:  code,
		Hash:  hash,
		Title: input.Title,
		Text:  input.Text,
		Tag:   tag,
	}

	memoDB.Create(memo)

	return memo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetUser(ctx context.Context, input GetUser) (*User, error) {
	var u User
	userDB.Find(&u, "Name = ?", input.Name)

	sessionKey := session.GenSession(u.Code)

	u.Token = sessionKey

	if hash.Sha1(input.Password) == u.Password {
		return &u, nil
	}
	return nil, nil
}

func (r *queryResolver) GetAllMemo(ctx context.Context, input GetAllMemo) (*AllMemo, error) {
	owner := session.IsLogin(input.Token)
	if owner == "" {
		return nil, nil
	}

	var memos []Memo
	memoDB.Find(&memos, "Owner = ?", owner)

	log.Println(memos)

	a := &AllMemo{
		Memos: memos,
	}

	return a, nil
}
