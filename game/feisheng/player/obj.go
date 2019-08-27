package player

import (
	"fgame/fgame/core/storage"
	feishengentity "fgame/fgame/game/feisheng/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//飞升对象
type PlayerFeiShengObject struct {
	player        player.Player
	id            int64
	feiLevel      int32
	addRate       int32
	gongDeNum     int64
	leftPotential int32
	tiZhi         int32
	liDao         int32
	jinGu         int32
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func NewPlayerFeiShengObject(pl player.Player) *PlayerFeiShengObject {
	o := &PlayerFeiShengObject{
		player: pl,
	}
	return o
}

func convertNewPlayerFeiShengObjectToEntity(o *PlayerFeiShengObject) (*feishengentity.PlayerFeiShengEntity, error) {

	e := &feishengentity.PlayerFeiShengEntity{
		Id:            o.id,
		PlayerId:      o.player.GetId(),
		FeiLevel:      o.feiLevel,
		AddRate:       o.addRate,
		GongDeNum:     o.gongDeNum,
		LeftPotential: o.leftPotential,
		TiZhi:         o.tiZhi,
		LiDao:         o.liDao,
		JinGu:         o.jinGu,
		UpdateTime:    o.updateTime,
		DeleteTime:    o.deleteTime,
		CreateTime:    o.createTime,
	}
	return e, nil
}

func (o *PlayerFeiShengObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFeiShengObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerFeiShengObject) GetFeiLevel() int32 {
	return o.feiLevel
}

func (o *PlayerFeiShengObject) GetAddRate() int32 {
	return o.addRate
}

func (o *PlayerFeiShengObject) GetGongDeNum() int64 {
	return o.gongDeNum
}

func (o *PlayerFeiShengObject) GetLeftPotential() int32 {
	return o.leftPotential
}

func (o *PlayerFeiShengObject) GetJinGu() int32 {
	return o.jinGu
}

func (o *PlayerFeiShengObject) GetLiDao() int32 {
	return o.liDao
}

func (o *PlayerFeiShengObject) GetTiZhi() int32 {
	return o.tiZhi
}

func (o *PlayerFeiShengObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerFeiShengObjectToEntity(o)
	return e, err
}

func (o *PlayerFeiShengObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*feishengentity.PlayerFeiShengEntity)

	o.id = pse.Id
	o.feiLevel = pse.FeiLevel
	o.addRate = pse.AddRate
	o.gongDeNum = pse.GongDeNum
	o.leftPotential = pse.LeftPotential
	o.tiZhi = pse.TiZhi
	o.liDao = pse.LiDao
	o.jinGu = pse.JinGu
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerFeiShengObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FEI_SHENG"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//飞升对象
type PlayerFeiShengReceiveObject struct {
	player     player.Player
	id         int64
	num        int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerFeiShengReceiveObject(pl player.Player) *PlayerFeiShengReceiveObject {
	o := &PlayerFeiShengReceiveObject{
		player: pl,
	}
	return o
}

func convertNewPlayerFeiShengReceiveObjectToEntity(o *PlayerFeiShengReceiveObject) (*feishengentity.PlayerFeiShengReceiveEntity, error) {

	e := &feishengentity.PlayerFeiShengReceiveEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Num:        o.num,
		UpdateTime: o.updateTime,
		DeleteTime: o.deleteTime,
		CreateTime: o.createTime,
	}
	return e, nil
}

func (o *PlayerFeiShengReceiveObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFeiShengReceiveObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerFeiShengReceiveObject) GetNum() int32 {
	return o.num
}

func (o *PlayerFeiShengReceiveObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerFeiShengReceiveObjectToEntity(o)
	return e, err
}

func (o *PlayerFeiShengReceiveObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*feishengentity.PlayerFeiShengReceiveEntity)

	o.id = pse.Id
	o.num = pse.Num
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerFeiShengReceiveObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "FEI_SHENG"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
