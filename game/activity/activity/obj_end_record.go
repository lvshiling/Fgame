package activity

import (
	"fgame/fgame/core/storage"
	activityentity "fgame/fgame/game/activity/entity"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//活动结束记录对象
type ActivityEndRecordObject struct {
	id           int64
	serverId     int32
	activityType activitytypes.ActivityType
	endTime      int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func newActivityEndRecordObject() *ActivityEndRecordObject {
	o := &ActivityEndRecordObject{}
	return o
}

func convertActivityEndRecordObjectToEntity(o *ActivityEndRecordObject) (e *activityentity.ActivityEndRecordEntity, err error) {

	e = &activityentity.ActivityEndRecordEntity{
		Id:           o.id,
		ServerId:     o.serverId,
		ActivityType: int32(o.activityType),
		EndTime:      o.endTime,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *ActivityEndRecordObject) GetDBId() int64 {
	return o.id
}

func (o *ActivityEndRecordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertActivityEndRecordObjectToEntity(o)
	return e, err
}

func (o *ActivityEndRecordObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*activityentity.ActivityEndRecordEntity)

	o.id = te.Id
	o.serverId = te.ServerId
	o.endTime = te.EndTime
	o.activityType = activitytypes.ActivityType(te.ActivityType)
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *ActivityEndRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Activity_End_Record"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
