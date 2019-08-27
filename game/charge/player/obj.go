package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	chargeentity "fgame/fgame/game/charge/entity"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//充值对象
type PlayerChargeObject struct {
	player     player.Player
	id         int64
	chargeType int32
	chargeId   int32
	chargeNum  int32
	orderId    string
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerChargeObject(pl player.Player) *PlayerChargeObject {
	o := &PlayerChargeObject{
		player: pl,
	}
	return o
}

func convertNewPlayerChargeObjectToEntity(o *PlayerChargeObject) (*chargeentity.PlayerChargeEntity, error) {
	e := &chargeentity.PlayerChargeEntity{
		Id:         o.id,
		PlayerId:   o.GetPlayerId(),
		ChargeType: o.chargeType,
		ChargeId:   o.chargeId,
		ChargeNum:  o.chargeNum,
		OrderId:    o.orderId,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerChargeObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerChargeObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerChargeObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerChargeObjectToEntity(o)
	return e, err
}

func (o *PlayerChargeObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chargeentity.PlayerChargeEntity)

	o.id = pse.Id

	o.chargeType = pse.ChargeType
	o.chargeNum = pse.ChargeNum
	o.chargeId = pse.ChargeId
	o.orderId = pse.OrderId
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerChargeObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Charge"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerChargeObject) GetChargeNum() int32 {

	return o.chargeNum
}

//档次首充对象
type PlayerFirstChargeRecordObject struct {
	player     player.Player
	id         int64
	chargeType int32
	chargeId   int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerFirstChargeRecordObject(pl player.Player) *PlayerFirstChargeRecordObject {
	o := &PlayerFirstChargeRecordObject{
		player: pl,
	}
	return o
}

func convertNewPlayerFirstChargeRecordObjectToEntity(o *PlayerFirstChargeRecordObject) (*chargeentity.PlayerFirstChargeRecordEntity, error) {
	e := &chargeentity.PlayerFirstChargeRecordEntity{
		Id:         o.id,
		PlayerId:   o.GetPlayerId(),
		ChargeType: o.chargeType,
		ChargeId:   o.chargeId,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerFirstChargeRecordObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerFirstChargeRecordObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerFirstChargeRecordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerFirstChargeRecordObjectToEntity(o)
	return e, err
}

func (o *PlayerFirstChargeRecordObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chargeentity.PlayerFirstChargeRecordEntity)

	o.id = pse.Id

	o.chargeType = pse.ChargeType
	o.chargeId = pse.ChargeId
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerFirstChargeRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Charge"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//每日首充对象
type PlayerCycleChargeRecordObject struct {
	player          player.Player
	id              int64
	chargeNum       int64
	preDayChargeNum int64
	updateTime      int64
	createTime      int64
	deleteTime      int64
}

func NewPlayerCycleChargeRecordObject(pl player.Player) *PlayerCycleChargeRecordObject {
	o := &PlayerCycleChargeRecordObject{
		player: pl,
	}
	return o
}

func convertNewPlayerCycleChargeRecordObjectToEntity(o *PlayerCycleChargeRecordObject) (*chargeentity.PlayerCycleChargeRecordEntity, error) {
	e := &chargeentity.PlayerCycleChargeRecordEntity{
		Id:              o.id,
		PlayerId:        o.GetPlayerId(),
		ChargeNum:       o.chargeNum,
		PreDayChargeNum: o.preDayChargeNum,
		UpdateTime:      o.updateTime,
		CreateTime:      o.createTime,
		DeleteTime:      o.deleteTime,
	}
	return e, nil
}

func (o *PlayerCycleChargeRecordObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerCycleChargeRecordObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerCycleChargeRecordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerCycleChargeRecordObjectToEntity(o)
	return e, err
}

func (o *PlayerCycleChargeRecordObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chargeentity.PlayerCycleChargeRecordEntity)

	o.id = pse.Id

	o.chargeNum = pse.ChargeNum
	o.preDayChargeNum = pse.PreDayChargeNum
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerCycleChargeRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "CycleCharge"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

// 新档次首充对象
type PlayerNewFirstChargeRecordObject struct {
	player     player.Player
	id         int64
	record     []int32
	startTime  int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerNewFirstChargeRecordObject(pl player.Player) *PlayerNewFirstChargeRecordObject {
	o := &PlayerNewFirstChargeRecordObject{
		player: pl,
	}
	return o
}

func convertNewPlayerNewFirstChargeRecordObjectToEntity(o *PlayerNewFirstChargeRecordObject) (*chargeentity.PlayerNewFirstChargeRecordEntity, error) {
	recordStr, err := json.Marshal(o.record)
	if err != nil {
		return nil, err
	}

	e := &chargeentity.PlayerNewFirstChargeRecordEntity{
		Id:         o.id,
		PlayerId:   o.GetPlayerId(),
		Record:     string(recordStr),
		StartTime:  o.startTime,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerNewFirstChargeRecordObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerNewFirstChargeRecordObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerNewFirstChargeRecordObject) GetRecord() []int32 {
	return o.record
}

func (o *PlayerNewFirstChargeRecordObject) GetStartTime() int64 {
	return o.startTime
}

func (o *PlayerNewFirstChargeRecordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerNewFirstChargeRecordObjectToEntity(o)
	return e, err
}

func (o *PlayerNewFirstChargeRecordObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*chargeentity.PlayerNewFirstChargeRecordEntity)

	var record []int32
	err := json.Unmarshal([]byte(pse.Record), &record)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.record = record
	o.startTime = pse.StartTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerNewFirstChargeRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "NewFisrtCharge"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
