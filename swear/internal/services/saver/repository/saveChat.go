package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"swearBot/internal/ddl/chatsDDL"
	saverModel "swearBot/internal/services/saver/model"
)

func (r *SaverRepository) SaveChat(ctx context.Context, req saverModel.SaveChatReq) error {
	return r.db.Exec(ctx, sq.
		Insert(chatsDDL.Table).
		SetMap(map[string]any{
			chatsDDL.ColumnID:                       req.ID,
			chatsDDL.ColumnType:                     req.Type,
			chatsDDL.ColumnTitle:                    req.Title,
			chatsDDL.ColumnFirstName:                req.FirstName,
			chatsDDL.ColumnLastName:                 req.LastName,
			chatsDDL.ColumnUsername:                 req.Username,
			chatsDDL.ColumnBio:                      req.Bio,
			chatsDDL.ColumnDescription:              req.Description,
			chatsDDL.ColumnInviteLink:               req.InviteLink,
			chatsDDL.ColumnSlowMode:                 req.SlowMode,
			chatsDDL.ColumnStickerSet:               req.StickerSet,
			chatsDDL.ColumnCanSetStickerSet:         req.CanSetStickerSet,
			chatsDDL.ColumnCustomEmojiSetName:       req.CustomEmojiSetName,
			chatsDDL.ColumnLinkedChatId:             req.LinkedChatID,
			chatsDDL.ColumnPrivate:                  req.Private,
			chatsDDL.ColumnProtected:                req.Protected,
			chatsDDL.ColumnNoVoiceAndVideo:          req.NoVoiceAndVideo,
			chatsDDL.ColumnHasHiddenMembers:         req.HasHiddenMembers,
			chatsDDL.ColumnAggressiveAntiSpam:       req.AggressiveAntiSpam,
			chatsDDL.ColumnCustomEmojiId:            req.CustomEmojiID,
			chatsDDL.ColumnEmojiExpirationUnixtime:  req.EmojiExpirationUnixtime,
			chatsDDL.ColumnBackgroundEmojiId:        req.BackgroundEmojiID,
			chatsDDL.ColumnAccentColorId:            req.AccentColorID,
			chatsDDL.ColumnProfileAccentColorId:     req.ProfileAccentColorID,
			chatsDDL.ColumnProfileBackgroundEmojiId: req.ProfileBackgroundEmojiID,
			chatsDDL.ColumnHasVisibleHistory:        req.HasVisibleHistory,
			chatsDDL.ColumnUnrestrictBoosts:         req.UnrestrictBoosts,
			chatsDDL.ColumnCountUsers:               req.CountUsers,
		}).
		Suffix(fmt.Sprintf(`ON CONFLICT (%s) DO NOTHING`, chatsDDL.ColumnID)),
	)
}
