package welfare

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	welfareentity "fgame/fgame/game/welfare/entity"

	"github.com/pkg/errors"
)

//活动奖励次数限制对象
type OpenActivityRewardsLimitObject struct {
	id         int64
	serverId   int32
	groupId    int32
	timesMap   map[int32]int32
	startTime  int64
	endTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func newOpenActivityRewardsLimitObject() *OpenActivityRewardsLimitObject {
	o := &OpenActivityRewardsLimitObject{}
	return o
}

func convertOpenActivityRewardsLimitObjectToEntity(o *OpenActivityRewardsLimitObject) (e *welfareentity.OpenActivityRewardsLimitEntity, err error) {
	data, err := json.Marshal(o.timesMap)
	if err != nil {
		return
	}
	e = &welfareentity.OpenActivityRewardsLimitEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		GroupId:    o.groupId,
		TimesMap:   string(data),
		StartTime:  o.startTime,
		EndTime:    o.endTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *OpenActivityRewardsLimitObject) GetDBId() int64 {
	return o.id
}

func (o *OpenActivityRewardsLimitObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertOpenActivityRewardsLimitObjectToEntity(o)
	return e, err
}

func (o *OpenActivityRewardsLimitObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*welfareentity.OpenActivityRewardsLimitEntity)

	newMap := make(map[int32]int32)
	err := json.Unmarshal([]byte(te.TimesMap), &newMap)
	if err != nil {
		return err
	}

	o.id = te.Id
	o.serverId = te.ServerId
	o.groupId = te.GroupId
	o.timesMap = newMap
	o.startTime = te.StartTime
	o.endTime = te.EndTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *OpenActivityRewardsLimitObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "OpenActivityRewardsLimit"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
