package keep_server

import (
	"context"
	keep_server "keep-server"
	"keep-server/model/memo"
	"keep-server/session"
	"strconv"
	"time"
)

var (
	db = memo.Connect()
)

type mutationResolver struct{ *keep_server.Resolver }

func (r *mutationResolver) EditMemo(ctx context.Context, input EditMemo) (*Memo, error) {
	owner := session.IsLogin(input.Token)
	if owner == "" {
		return nil, nil
	}

	memo := memo.Memo{}
	memo.Code = input.MemoCode
	memoDB.Find(&memo)

	if memo.Owner != owner {
		return nil, nil
	}

	tag := ""
	for _, v := range input.Tag {
		tag += v + ","
	}

	memo.Time = strconv.FormatInt(time.Now().UTC().Unix(), 10)
	memo.Title = input.Title
	memo.Text = input.Text
	memo.Tag = tag

	memoDB.Save(&memo)

	ans := &Memo{
		Owner: memo.Owner,
		Code:  memo.Code,
		Time:  memo.Time,
		Title: memo.Title,
		Text:  memo.Text,
		Tag:   tag,
	}

	return ans, nil
}
