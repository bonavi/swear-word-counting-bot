package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"swearBot/internal/ddl/messagesDDL"
	saverModel "swearBot/internal/services/saver/model"
)

func (r *SaverRepository) SaveMessage(ctx context.Context, req saverModel.SaveMessageReq) error {
	return r.db.Exec(ctx, sq.
		Insert(messagesDDL.Table).
		SetMap(map[string]any{
			messagesDDL.ColumnID:                 req.ID,
			messagesDDL.ColumnThreadID:           req.ThreadID,
			messagesDDL.ColumnUnixtime:           req.Unixtime,
			messagesDDL.ColumnOriginalMessageID:  req.OriginalMessageID,
			messagesDDL.ColumnOriginalSignature:  req.OriginalSignature,
			messagesDDL.ColumnOriginalSenderName: req.OriginalSenderName,
			messagesDDL.ColumnOriginalUnixtime:   req.OriginalUnixtime,
			messagesDDL.ColumnAutomaticForward:   req.AutomaticForward,
			messagesDDL.ColumnLastEdit:           req.LastEdit,
			messagesDDL.ColumnTopicMessage:       req.TopicMessage,
			messagesDDL.ColumnProtected:          req.Protected,
			messagesDDL.ColumnAlbumID:            req.AlbumID,
			messagesDDL.ColumnSignature:          req.Signature,
			messagesDDL.ColumnText:               req.Text,
			messagesDDL.ColumnPayload:            req.Payload,
			messagesDDL.ColumnCaption:            req.Caption,
			messagesDDL.ColumnNewGroupTitle:      req.NewGroupTitle,
			messagesDDL.ColumnGroupPhotoDeleted:  req.GroupPhotoDeleted,
			messagesDDL.ColumnGroupCreated:       req.GroupCreated,
			messagesDDL.ColumnSuperGroupCreated:  req.SuperGroupCreated,
			messagesDDL.ColumnChannelCreated:     req.ChannelCreated,
			messagesDDL.ColumnMigrateTo:          req.MigrateTo,
			messagesDDL.ColumnMigrateFrom:        req.MigrateFrom,
			messagesDDL.ColumnConnectedWebsite:   req.ConnectedWebsite,
			messagesDDL.ColumnSenderBoostCount:   req.SenderBoostCount,
			messagesDDL.ColumnHasMediaSpoiler:    req.HasMediaSpoiler,
		}).
		Suffix(fmt.Sprintf(`ON CONFLICT (%s) DO NOTHING`, messagesDDL.ColumnID)),
	)
}
