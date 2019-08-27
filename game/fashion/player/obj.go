package player

import (
	"fgame/fgame/core/storage"
	fashionentity "fgame/fgame/game/fashion/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//时装对象
type PlayerFashionObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	FashionId  int32
	Star       int32
	UpStarNum  int32
	UpStarPro  int32
	IsExpire   int32
	ActiveTime int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerFashionObject(pl player.Player) *PlayerFashionObject {
	pto := &PlayerFashionObject{
		player: pl,
	}
	return pto
}

func convertNewPlayerFashionObjectToEntity(o *PlayerFashionObject) (*fashionentity.PlayerFashionEntity, error) {
	e := &fashionentity.PlayerFashionEntity{
		Id:         o.Id,
		PlayerId:   o.PlayerId,
		FashionId:  o.FashionId,
		Star:       o.Star,
		UpStarNum:  o.UpStarNum,
		UpStarPro:  o.UpStarPro,
		IsExpire:   o.IsExpire,
		ActiveTime: o.ActiveTime,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerFashionObject) GetPlayerId() int64 {
	return o.PlayerId
}

func (o *PlayerFashionObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerFashionObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerFashionObjectToEntity(o)
	return e, err
}

func (o *PlayerFashionObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*fashionentity.PlayerFashionEntity)

	o.Id = pse.Id
	o.PlayerId = pse.PlayerId
	o.FashionId = pse.FashionId
	o.Star = pse.Star
	o.UpStarNum = pse.UpStarNum
	o.UpStarPro = pse.UpStarPro
	o.IsExpire = pse.IsExpire
	o.ActiveTime = pse.ActiveTime
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerFashionObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Fashion"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//穿戴时装对象
type PlayerFashionWearObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	FashionWear int32
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerFashionWearObject(pl player.Player) *PlayerFashionWearObject {
	pfto := &PlayerFashionWearObject{
		player: pl,
	}
	return pfto
}

func convertNewPlayerFashionWearObjectToEntity(o *PlayerFashionWearObject) (*fashionentity.PlayerWearFashionEntity, error) {
	e := &fashionentity.PlayerWearFashionEntity{
		Id:          o.Id,
		PlayerId:    o.PlayerId,
		FashionWear: o.FashionWear,
		UpdateTime:  o.UpdateTime,
		CreateTime:  o.CreateTime,
		DeleteTime:  o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerFashionWearObject) GetPlayerId() int64 {
	return o.PlayerId
}

func (o *PlayerFashionWearObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerFashionWearObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerFashionWearObjectToEntity(o)
	return e, err
}

func (o *PlayerFashionWearObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*fashionentity.PlayerWearFashionEntity)

	o.Id = pse.Id
	o.PlayerId = pse.PlayerId
	o.FashionWear = pse.FashionWear
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerFashionWearObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FashionWear"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

// 时装试用卡对象
type PlayerFashionTrialObject struct {
	player         player.Player
	id             int64
	trialFashionId int32
	expireTime     int64
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func NewPlayerFashionTrialObject(pl player.Player) *PlayerFashionTrialObject {
	pfto := &PlayerFashionTrialObject{
		player: pl,
	}
	return pfto
}

func convertNewPlayerFashionTrialObjectToEntity(o *PlayerFashionTrialObject) (*fashionentity.PlayerFashionTrialEntity, error) {
	e := &fashionentity.PlayerFashionTrialEntity{
		Id:             o.id,
		PlayerId:       o.player.GetId(),
		TrialFashionId: o.trialFashionId,
		ExpireTime:     o.expireTime,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *PlayerFashionTrialObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFashionTrialObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerFashionTrialObject) GetTrialFashionId() int32 {
	return o.trialFashionId
}

func (o *PlayerFashionTrialObject) GetExpireTime() int64 {
	return o.expireTime
}

func (o *PlayerFashionTrialObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerFashionTrialObjectToEntity(o)
	return e, err
}

func (o *PlayerFashionTrialObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*fashionentity.PlayerFashionTrialEntity)

	o.id = pse.Id
	o.trialFashionId = pse.TrialFashionId
	o.expireTime = pse.ExpireTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerFashionTrialObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FashionTrial"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
