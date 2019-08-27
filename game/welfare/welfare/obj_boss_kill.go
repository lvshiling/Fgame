package welfare

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	welfareentity "fgame/fgame/game/welfare/entity"

	"github.com/pkg/errors"
)

//Boss首杀记录对象
type BossKillRecordObject struct {
	id         int64
	serverId   int32
	groupId    int32
	bossIdList []int32
	startTime  int64
	endTime    int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func newBossKillRecordObject() *BossKillRecordObject {
	o := &BossKillRecordObject{}
	return o
}

func convertNewBossKillRecordObjectToEntity(o *BossKillRecordObject) (*welfareentity.OpenActivityBossKillEntity, error) {
	data, err := json.Marshal(o.bossIdList)
	if err != nil {
		return nil, err
	}

	e := &welfareentity.OpenActivityBossKillEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		GroupId:    o.groupId,
		BossIdList: string(data),
		StartTime:  o.startTime,
		EndTime:    o.endTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *BossKillRecordObject) isExistBossId(bossId int32) bool {
	for _, idRecord := range o.bossIdList {
		if idRecord == bossId {
			return true
		}
	}
	return false
}

func (o *BossKillRecordObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *BossKillRecordObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *BossKillRecordObject) GetGroupId() int32 {
	return o.groupId
}

func (o *BossKillRecordObject) GetDBId() int64 {
	return o.id
}

func (o *BossKillRecordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewBossKillRecordObjectToEntity(o)
	return e, err
}

func (o *BossKillRecordObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*welfareentity.OpenActivityBossKillEntity)

	var bossIdList []int32
	err := json.Unmarshal([]byte(pse.BossIdList), &bossIdList)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.groupId = pse.GroupId
	o.bossIdList = bossIdList
	o.startTime = pse.StartTime
	o.endTime = pse.EndTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *BossKillRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "BossKillRecord"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}
