package chat

import (
	"fgame/fgame/core/storage"
	chatentity "fgame/fgame/game/chat/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//聊天设置
type ChatSettingObject struct {
	id               int64
	serverId         int32
	worldVipLevel    int32
	worldLevel       int32
	allianceVipLevel int32
	allianceLevel    int32
	privateVipLevel  int32
	privateLevel     int32
	teamVipLevel     int32
	teamLevel        int32
	updateTime       int64
	createTime       int64
	deleteTime       int64
}

func createChatSettingObject() *ChatSettingObject {
	o := &ChatSettingObject{}
	return o
}

func convertChatSettingObjectToEntity(o *ChatSettingObject) (*chatentity.ChatSettingEntity, error) {
	e := &chatentity.ChatSettingEntity{
		Id:               o.id,
		ServerId:         o.serverId,
		WorldVipLevel:    o.worldVipLevel,
		WorldLevel:       o.worldLevel,
		AllianceVipLevel: o.allianceVipLevel,
		AllianceLevel:    o.allianceLevel,
		PrivateVipLevel:  o.privateVipLevel,
		PrivateLevel:     o.privateLevel,
		TeamVipLevel:     o.teamVipLevel,
		TeamLevel:        o.teamLevel,
		UpdateTime:       o.updateTime,
		CreateTime:       o.createTime,
		DeleteTime:       o.deleteTime,
	}
	return e, nil
}

func (o *ChatSettingObject) GetId() int64 {
	return o.id
}

func (o *ChatSettingObject) GetDBId() int64 {
	return o.id
}

func (o *ChatSettingObject) GetServerId() int32 {
	return o.serverId
}
func (o *ChatSettingObject) GetWorldVipLevel() int32 {
	return o.worldVipLevel
}

func (o *ChatSettingObject) GetWorldLevel() int32 {
	return o.worldLevel
}

func (o *ChatSettingObject) GetAllianceVipLevel() int32 {
	return o.allianceVipLevel
}

func (o *ChatSettingObject) GetAllianceLevel() int32 {
	return o.allianceLevel
}

func (o *ChatSettingObject) GetPrivateVipLevel() int32 {
	return o.privateVipLevel
}

func (o *ChatSettingObject) GetPrivateLevel() int32 {
	return o.privateLevel
}

func (o *ChatSettingObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertChatSettingObjectToEntity(o)
	return e, err
}

func (o *ChatSettingObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*chatentity.ChatSettingEntity)
	o.id = ae.Id
	o.serverId = ae.ServerId
	o.worldVipLevel = ae.WorldVipLevel
	o.worldLevel = ae.WorldLevel
	o.allianceVipLevel = ae.AllianceVipLevel
	o.allianceLevel = ae.AllianceLevel
	o.privateVipLevel = ae.PrivateVipLevel
	o.privateLevel = ae.PrivateLevel
	o.teamVipLevel = ae.TeamVipLevel
	o.teamLevel = ae.TeamLevel
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *ChatSettingObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Chat"))
		return
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
