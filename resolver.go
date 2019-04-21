package keep_server

import (
	context "context"
	"keep-server/model"
	"keep-server/model/memo"
	"keep-server/model/user"
	resolvers_editmemo "keep-server/resolvers/editmemo"
	resolvers_removememo "keep-server/resolvers/removememo"

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
	userDB.Find(&u, "Name = ?", input.Name)
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

	memo := &Memo{
		Owner: owner,
		Code:  code,
		Time:  strconv.FormatInt(time.Now().UTC().Unix(), 10),
		Title: input.Title,
		Text:  input.Text,
		Tag:   tag,
	}

	memoDB.Create(memo)

	return memo, nil
}

func (r *mutationResolver) EditMemo(ctx context.Context, input EditMemo) (*Memo, error) {
	memo := resolvers_editmemo.EditMemo(model.EditMemo{
		input.Token,
		input.MemoCode,
		input.Title,
		input.Text,
		input.Tag,
	})

	return &Memo{
		Owner: memo.Owner,
		Code:  memo.Code,
		Time:  memo.Time,
		Title: memo.Title,
		Text:  memo.Text,
		Tag:   memo.Tag,
	}, nil
}

func (r *mutationResolver) RemoveMemo(ctx context.Context, input RemoveMemo) (*Memo, error) {
	memo := resolvers_removememo.RemoveMemo(model.RemoveMemo{
		Token:    input.Token,
		MemoCode: input.MemoCode,
	})

	return &Memo{
		Owner: memo.Owner,
		Code:  memo.Code,
		Time:  memo.Time,
		Title: memo.Title,
		Text:  memo.Text,
		Tag:   memo.Tag,
	}, nil
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

func (r *queryResolver) GetAllMemo(ctx context.Context, input GetAllMemo) ([]Memo, error) {
	owner := session.IsLogin(input.Token)
	if owner == "" {
		return nil, nil
	}

	var memos []Memo
	memoDB.Find(&memos, "Owner = ?", owner)

	log.Println(memos)

	return memos, nil
}

func (r *queryResolver) Auth(ctx context.Context, input Auth) (*string, error) {
	owner := session.IsLogin(input.Token)
	if owner == "" {
		return nil, nil
	}

	return &owner, nil
}
