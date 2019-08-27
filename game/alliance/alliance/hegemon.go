package alliance

import (
	"fgame/fgame/core/storage"
	allianceentity "fgame/fgame/game/alliance/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

type AllianceHegemonObject struct {
	id                int64
	serverId          int32
	allianceId        int64
	winNum            int32
	defenceAllianceId int64
	updateTime        int64
	createTime        int64
	deleteTime        int64
}

func createAllianceHegemonObject() *AllianceHegemonObject {
	o := &AllianceHegemonObject{}
	return o
}

func convertAllianceHegemonObjectToEntity(o *AllianceHegemonObject) (*allianceentity.AllianceHegemonEntity, error) {
	e := &allianceentity.AllianceHegemonEntity{
		Id:                o.id,
		ServerId:          o.serverId,
		AllianceId:        o.allianceId,
		DefenceAllianceId: o.defenceAllianceId,
		WinNum:            o.winNum,
		UpdateTime:        o.updateTime,
		CreateTime:        o.createTime,
		DeleteTime:        o.deleteTime,
	}
	return e, nil
}

func (o *AllianceHegemonObject) GetId() int64 {
	return o.id
}

func (o *AllianceHegemonObject) GetDBId() int64 {
	return o.id
}

func (o *AllianceHegemonObject) GetAllianceId() int64 {
	return o.allianceId
}

func (o *AllianceHegemonObject) GetDefenceAllianceId() int64 {
	return o.defenceAllianceId
}

func (o *AllianceHegemonObject) GetWinNum() int32 {
	return o.winNum
}

func (o *AllianceHegemonObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertAllianceHegemonObjectToEntity(o)
	return e, err
}

func (o *AllianceHegemonObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*allianceentity.AllianceHegemonEntity)
	o.id = ae.Id
	o.serverId = ae.ServerId
	o.allianceId = ae.AllianceId
	o.winNum = ae.WinNum
	o.defenceAllianceId = ae.DefenceAllianceId
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *AllianceHegemonObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AllianceHegemon"))

	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
