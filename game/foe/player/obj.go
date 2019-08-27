package player

import (
	"fgame/fgame/core/storage"
	foeentity "fgame/fgame/game/foe/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//玩家仇人
type PlayerFoeObject struct {
	player     player.Player
	Id         int64
	AttackId   int64
	KillTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

//记录排序
type PlayerFoeObjectList []*PlayerFoeObject

func (l PlayerFoeObjectList) Len() int {
	return len(l)
}

func (l PlayerFoeObjectList) Less(i, j int) bool {
	return l[i].KillTime < l[j].KillTime
}

func (l PlayerFoeObjectList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func newPlayerFoeObject(pl player.Player) *PlayerFoeObject {
	o := &PlayerFoeObject{
		player: pl,
	}
	return o
}

func convertPlayerFoeObjectToEntity(o *PlayerFoeObject) (e *foeentity.PlayerFoeEntity, err error) {
	e = &foeentity.PlayerFoeEntity{
		Id:         o.Id,
		PlayerId:   o.player.GetId(),
		AttackId:   o.AttackId,
		KillTime:   o.KillTime,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerFoeObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFoeObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerFoeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFoeObjectToEntity(o)
	return e, err
}

func (o *PlayerFoeObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*foeentity.PlayerFoeEntity)

	o.Id = te.Id
	o.AttackId = te.AttackId
	o.KillTime = te.KillTime
	o.UpdateTime = te.UpdateTime
	o.CreateTime = te.CreateTime
	o.DeleteTime = te.DeleteTime
	return nil
}

func (o *PlayerFoeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Foe"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//玩家仇人反馈
type PlayerFoeFeedbackObject struct {
	player       player.Player
	id           int64
	isProtect    int32
	feedbackName string
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func newPlayerFoeFeedbackObject(pl player.Player) *PlayerFoeFeedbackObject {
	o := &PlayerFoeFeedbackObject{
		player: pl,
	}
	return o
}

func convertPlayerFoeFeedbackObjectToEntity(o *PlayerFoeFeedbackObject) (e *foeentity.PlayerFoeFeedbackEntity, err error) {
	e = &foeentity.PlayerFoeFeedbackEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		IsProtected:  o.isProtect,
		FeedbackName: o.feedbackName,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *PlayerFoeFeedbackObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFoeFeedbackObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerFoeFeedbackObject) GetFeedbackName() string {
	return o.feedbackName
}

func (o *PlayerFoeFeedbackObject) GetIsProtect() int32 {
	return o.isProtect
}

func (o *PlayerFoeFeedbackObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFoeFeedbackObjectToEntity(o)
	return e, err
}

func (o *PlayerFoeFeedbackObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*foeentity.PlayerFoeFeedbackEntity)

	o.id = te.Id
	o.isProtect = te.IsProtected
	o.feedbackName = te.FeedbackName
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerFoeFeedbackObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FoeFeedback"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//玩家仇人反馈
type PlayerFoeProtectObject struct {
	player     player.Player
	id         int64
	expireTime int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func newPlayerFoeProtectObject(pl player.Player) *PlayerFoeProtectObject {
	o := &PlayerFoeProtectObject{
		player: pl,
	}
	return o
}

func convertPlayerFoeProtectObjectToEntity(o *PlayerFoeProtectObject) (e *foeentity.PlayerFoeProtectEntity, err error) {
	e = &foeentity.PlayerFoeProtectEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		ExpireTime: o.expireTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerFoeProtectObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFoeProtectObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerFoeProtectObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerFoeProtectObjectToEntity(o)
	return e, err
}

func (o *PlayerFoeProtectObject) FromEntity(e storage.Entity) error {
	te, _ := e.(*foeentity.PlayerFoeProtectEntity)

	o.id = te.Id
	o.expireTime = te.ExpireTime
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func (o *PlayerFoeProtectObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FoeProtect"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
