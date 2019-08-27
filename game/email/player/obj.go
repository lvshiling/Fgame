package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/email/entity"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//邮件对象
type PlayerEmailObject struct {
	player          player.Player
	id              int64
	isRead          int32
	isGetAttachment int32
	title           string
	content         string
	attachmentInfo  []*droptemplate.DropItemData
	updateTime      int64
	createTime      int64
	deleteTime      int64
}

func (o *PlayerEmailObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerEmailObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerEmailObject) GetEmailId() int64 {
	return o.id
}

func (o *PlayerEmailObject) GetTitle() string {
	return o.title
}

func (o *PlayerEmailObject) GetContent() string {
	return o.content
}

func (o *PlayerEmailObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *PlayerEmailObject) GetIsGetAttachment() int32 {
	return o.isGetAttachment
}

func (o *PlayerEmailObject) GetIsRead() int32 {
	return o.isRead
}

func (o *PlayerEmailObject) GetAttachmentInfo() []*droptemplate.DropItemData {
	return o.attachmentInfo
}

func NewPlayerEmailObject(pl player.Player) *PlayerEmailObject {
	o := &PlayerEmailObject{
		player: pl,
	}

	return o
}

func (o *PlayerEmailObject) FromEntity(e storage.Entity) error {
	eEntity := e.(*entity.PlayerEmailEntity)
	attachmentInfo := []*droptemplate.DropItemData{}
	if err := json.Unmarshal([]byte(eEntity.AttachementInfo), &attachmentInfo); err != nil {
		return err
	}

	o.id = eEntity.Id
	o.isRead = eEntity.IsRead
	o.isGetAttachment = eEntity.IsGetAttachment
	o.title = eEntity.Title
	o.content = eEntity.Content
	o.attachmentInfo = attachmentInfo
	o.updateTime = eEntity.UpdateTime
	o.createTime = eEntity.CreateTime
	o.deleteTime = eEntity.DeleteTime
	return nil
}
func (o *PlayerEmailObject) ToEntity() (e storage.Entity, err error) {
	emailsInfoBytes, err := json.Marshal(o.attachmentInfo)
	if err != nil {
		return nil, err
	}

	e = &entity.PlayerEmailEntity{
		Id:              o.id,
		PlayerId:        o.player.GetId(),
		IsRead:          o.isRead,
		IsGetAttachment: o.isGetAttachment,
		Title:           o.title,
		Content:         o.content,
		AttachementInfo: string(emailsInfoBytes),
		UpdateTime:      o.updateTime,
		CreateTime:      o.createTime,
		DeleteTime:      o.deleteTime,
	}
	return e, err
}

func (o *PlayerEmailObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Email"))
	}
	obj, ok := e.(playertypes.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//是否超过15天的邮件
func (o *PlayerEmailObject) isTimeOut() bool {
	now := global.GetGame().GetTimeService().Now()
	validDay := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEmailSaveDay)
	expiredTime := int64(validDay) * int64(common.DAY)

	isTimeOut := now > (o.createTime + int64(expiredTime))
	return isTimeOut
}
