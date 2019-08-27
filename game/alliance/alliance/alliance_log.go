package alliance

import (
	"fgame/fgame/core/storage"
	allianceentity "fgame/fgame/game/alliance/entity"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//仙盟日志
type AllianceLogObject struct {
	id         int64
	allianceId int64
	content    string
	updateTime int64
	createTime int64
	deleteTime int64
}

func createAllianceLogObject() *AllianceLogObject {
	o := &AllianceLogObject{}
	return o
}

func convertAllianceLogObjectToEntity(o *AllianceLogObject) (*allianceentity.AllianceLogEntity, error) {
	e := &allianceentity.AllianceLogEntity{
		Id:         o.id,
		AllianceId: o.allianceId,
		Content:    o.content,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *AllianceLogObject) GetId() int64 {
	return o.id
}

func (o *AllianceLogObject) GetDBId() int64 {
	return o.id
}

func (o *AllianceLogObject) GetAllianceId() int64 {
	return o.allianceId
}

func (o *AllianceLogObject) GetContent() string {
	return o.content
}

func (o *AllianceLogObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *AllianceLogObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertAllianceLogObjectToEntity(o)
	return e, err
}

func (o *AllianceLogObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*allianceentity.AllianceLogEntity)
	o.id = ae.Id
	o.allianceId = ae.AllianceId
	o.content = ae.Content
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *AllianceLogObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AllianceLog"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
