package resolvers_removememo

import (
	"keep-server/model"
	"keep-server/model/memo"
	"keep-server/session"
)

var (
	db = memo.Connect()
)

func RemoveMemo(input model.RemoveMemo) *memo.Memo {
	owner := session.IsLogin(input.Token)
	if owner == "" {
		return nil
	}

	memo := memo.Memo{}

	db.First(&memo, "Code = ?", input.MemoCode)

	if memo.Owner != owner {
		return nil
	}

	db.Unscoped().Delete(&memo)

	return &memo
}
