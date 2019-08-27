package baby

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	babyentity "fgame/fgame/game/baby/entity"
	babytypes "fgame/fgame/game/baby/types"
	"fgame/fgame/game/global"

	"github.com/pkg/errors"
)

//配偶宝宝对象
type CoupleBabyObject struct {
	id         int64
	serverId   int32
	playerId   int64
	babyList   []*babytypes.CoupleBabyData
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewCoupleBabyObject() *CoupleBabyObject {
	o := &CoupleBabyObject{}
	return o
}

func convertNewCoupleBabyObjectToEntity(o *CoupleBabyObject) (*babyentity.CoupleBabyEntity, error) {

	data, _ := json.Marshal(o.babyList)

	e := &babyentity.CoupleBabyEntity{
		Id:         o.id,
		ServerId:   o.serverId,
		PlayerId:   o.playerId,
		BabyList:   string(data),
		UpdateTime: o.updateTime,
		DeleteTime: o.deleteTime,
		CreateTime: o.createTime,
	}
	return e, nil
}

func (o *CoupleBabyObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *CoupleBabyObject) GetDBId() int64 {
	return o.id
}

func (o *CoupleBabyObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewCoupleBabyObjectToEntity(o)
	return e, err
}

func (o *CoupleBabyObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*babyentity.CoupleBabyEntity)

	var babyList []*babytypes.CoupleBabyData
	err := json.Unmarshal([]byte(pse.BabyList), &babyList)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.serverId = pse.ServerId
	o.playerId = pse.PlayerId
	o.babyList = babyList
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *CoupleBabyObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "COUPLE_BABY"))
	}

	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}
