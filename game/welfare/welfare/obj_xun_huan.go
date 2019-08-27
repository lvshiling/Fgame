package welfare

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	welfareentity "fgame/fgame/game/welfare/entity"

	"github.com/pkg/errors"
)

//循环活动
type ActivityXunHuanObject struct {
	id          int64
	serverId    int32
	arrGroup    int32
	activityDay int32
	startTime   int64
	endTime     int64
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func newActivityXunHuanObject() *ActivityXunHuanObject {
	o := &ActivityXunHuanObject{}
	return o
}

func convertActivityXunHuanObjectToEntity(o *ActivityXunHuanObject) (e *welfareentity.OpenActivityXunHuanEntity, err error) {

	e = &welfareentity.OpenActivityXunHuanEntity{
		Id:          o.id,
		ServerId:    o.serverId,
		ArrGroup:    o.arrGroup,
		ActivityDay: o.activityDay,
		StartTime:   o.startTime,
		EndTime:     o.endTime,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return e, nil
}

func (o *ActivityXunHuanObject) GetDBId() int64 {
	return o.id
}

func (o *ActivityXunHuanObject) GetArrGroup() int32 {
	return o.arrGroup
}

func (o *ActivityXunHuanObject) GetActivityDay() int32 {
	return o.activityDay
}

func (o *ActivityXunHuanObject) isOnXunHuan() bool {
	if o.activityDay != 0 && o.arrGroup != 0 {
		return true
	}

	return false
}

func (o *ActivityXunHuanObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertActivityXunHuanObjectToEntity(o)
	return e, err
}

func (o *ActivityXunHuanObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.OpenActivityXunHuanEntity)

	o.id = te.Id
	o.serverId = te.ServerId
	o.arrGroup = te.ArrGroup
	o.activityDay = te.ActivityDay
	o.startTime = te.StartTime
	o.endTime = te.EndTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *ActivityXunHuanObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OpenActivityXunHuan"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}
