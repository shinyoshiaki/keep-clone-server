package resolvers

import (
	"keep-server/model"
	"keep-server/model/memo"
	"keep-server/session"
	"strconv"
	"time"
)

var (
	db = memo.Connect()
)

func EditMemo(input model.EditMemo) *memo.Memo {

	owner := session.IsLogin(input.Token)
	if owner == "" {
		return nil
	}

	memo := memo.Memo{}
	newMemo := memo
	db.First(&newMemo, "Code = ?", input.MemoCode)

	if newMemo.Owner != owner {
		return nil
	}

	tag := ""
	for _, v := range input.Tag {
		tag += v + ","
	}

	newMemo.Time = strconv.FormatInt(time.Now().UTC().Unix(), 10)
	newMemo.Title = input.Title
	newMemo.Text = input.Text
	newMemo.Tag = tag

	db.Model(&memo).Update(&newMemo)
	db.Save(&newMemo)

	return &newMemo
}
