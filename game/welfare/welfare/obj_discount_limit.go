package welfare

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	welfareentity "fgame/fgame/game/welfare/entity"

	"github.com/pkg/errors"
)

//活动黑货商店次数限制对象
type OpenActivityDiscountLimitObject struct {
	id          int64
	serverId    int32
	groupId     int32
	discountDay int32
	timesMap    map[int32]int32
	startTime   int64
	endTime     int64
	updateTime  int64
	createTime  int64
	deleteTime  int64
}

func newOpenActivityDiscountLimitObject() *OpenActivityDiscountLimitObject {
	o := &OpenActivityDiscountLimitObject{}
	return o
}

func convertOpenActivityDiscountLimitObjectToEntity(o *OpenActivityDiscountLimitObject) (e *welfareentity.OpenActivityDiscountLimitEntity, err error) {
	data, err := json.Marshal(o.timesMap)
	if err != nil {
		return
	}
	e = &welfareentity.OpenActivityDiscountLimitEntity{
		Id:          o.id,
		ServerId:    o.serverId,
		GroupId:     o.groupId,
		DiscountDay: o.discountDay,
		TimesMap:    string(data),
		StartTime:   o.startTime,
		EndTime:     o.endTime,
		UpdateTime:  o.updateTime,
		CreateTime:  o.createTime,
		DeleteTime:  o.deleteTime,
	}
	return e, nil
}

func (o *OpenActivityDiscountLimitObject) GetDBId() int64 {
	return o.id
}

func (o *OpenActivityDiscountLimitObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertOpenActivityDiscountLimitObjectToEntity(o)
	return e, err
}

func (o *OpenActivityDiscountLimitObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.OpenActivityDiscountLimitEntity)

	newMap := make(map[int32]int32)
	err := json.Unmarshal([]byte(te.TimesMap), &newMap)
	if err != nil {
		return err
	}

	o.id = te.Id
	o.serverId = te.ServerId
	o.groupId = te.GroupId
	o.discountDay = te.DiscountDay
	o.timesMap = newMap
	o.startTime = te.StartTime
	o.endTime = te.EndTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *OpenActivityDiscountLimitObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OpenActivityDiscountLimit"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
