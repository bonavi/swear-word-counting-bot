package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"swearBot/internal/ddl/usersDDL"
	saverModel "swearBot/internal/services/saver/model"
)

func (r *SaverRepository) SaveUser(ctx context.Context, req saverModel.SaveUserReq) error {
	return r.db.Exec(ctx, sq.
		Insert(usersDDL.Table).
		SetMap(map[string]any{
			usersDDL.ColumnID:                req.ID,
			usersDDL.ColumnFirstName:         req.FirstName,
			usersDDL.ColumnLastName:          req.LastName,
			usersDDL.ColumnIsForum:           req.IsForum,
			usersDDL.ColumnUsername:          req.Username,
			usersDDL.ColumnLanguageCode:      req.LanguageCode,
			usersDDL.ColumnIsBot:             req.IsBot,
			usersDDL.ColumnIsPremium:         req.IsPremium,
			usersDDL.ColumnAddedToMenu:       req.AddedToMenu,
			usersDDL.ColumnUsernames:         req.Usernames,
			usersDDL.ColumnCustomEmojiStatus: req.CustomEmojiStatus,
		}).
		Suffix(fmt.Sprintf(`ON CONFLICT (%s) DO NOTHING`, usersDDL.ColumnID)),
	)
}
