package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	fabaoentity "fgame/fgame/game/fabao/entity"
	fabaotypes "fgame/fgame/game/fabao/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//法宝对象
type PlayerFaBaoObject struct {
	player        player.Player
	id            int64
	playerId      int64
	advanceId     int
	faBaoId       int32
	unrealLevel   int32
	unrealNum     int32
	unrealPro     int32
	unrealList    []int
	timesNum      int32
	bless         int32
	blessTime     int64
	tongLingLevel int32
	tongLingNum   int32
	tongLingPro   int32
	hidden        int32
	power         int64
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func NewPlayerFaBaoObject(pl player.Player) *PlayerFaBaoObject {
	pwo := &PlayerFaBaoObject{
		player: pl,
	}
	return pwo
}

func convertNewPlayerFaBaoObjectToEntity(pwo *PlayerFaBaoObject) (*fabaoentity.PlayerFaBaoEntity, error) {
	unrealInfoBytes, err := json.Marshal(pwo.unrealList)
	if err != nil {
		return nil, err
	}
	e := &fabaoentity.PlayerFaBaoEntity{
		Id:            pwo.id,
		PlayerId:      pwo.playerId,
		AdvancedId:    pwo.advanceId,
		FaBaoId:       pwo.faBaoId,
		UnrealLevel:   pwo.unrealLevel,
		UnrealNum:     pwo.unrealNum,
		UnrealPro:     pwo.unrealPro,
		UnrealInfo:    string(unrealInfoBytes),
		TimesNum:      pwo.timesNum,
		Bless:         pwo.bless,
		BlessTime:     pwo.blessTime,
		TongLingLevel: pwo.tongLingLevel,
		TongLingNum:   pwo.tongLingNum,
		TongLingPro:   pwo.tongLingPro,
		Hidden:        pwo.hidden,
		Power:         pwo.power,
		UpdateTime:    pwo.updateTime,
		CreateTime:    pwo.createTime,
		DeleteTime:    pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerFaBaoObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerFaBaoObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerFaBaoObject) GetAdvancedId() int32 {
	return int32(pwo.advanceId)
}

func (pwo *PlayerFaBaoObject) GetFaBaoId() int32 {
	return pwo.faBaoId
}

func (pwo *PlayerFaBaoObject) GetUnrealLevel() int32 {
	return pwo.unrealLevel
}

func (pwo *PlayerFaBaoObject) GetUnrealNum() int32 {
	return pwo.unrealNum
}

func (pwo *PlayerFaBaoObject) GetUnrealPro() int32 {
	return pwo.unrealPro
}

func (pwo *PlayerFaBaoObject) GetUnrealList() []int {
	return pwo.unrealList
}

func (pwo *PlayerFaBaoObject) GetTimesNum() int32 {
	return pwo.timesNum
}

func (pwo *PlayerFaBaoObject) GetBless() int32 {
	return pwo.bless
}

func (pwo *PlayerFaBaoObject) GetBlessTime() int64 {
	return pwo.blessTime
}

func (pwo *PlayerFaBaoObject) GetTongLingLevel() int32 {
	return pwo.tongLingLevel
}

func (pwo *PlayerFaBaoObject) GetTongLingNum() int32 {
	return pwo.tongLingNum
}

func (pwo *PlayerFaBaoObject) GetTongLingPro() int32 {
	return pwo.tongLingPro
}

func (pwo *PlayerFaBaoObject) GetHidden() int32 {
	return pwo.hidden
}

func (pwo *PlayerFaBaoObject) GetPower() int64 {
	return pwo.power
}

func (pwo *PlayerFaBaoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerFaBaoObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerFaBaoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*fabaoentity.PlayerFaBaoEntity)

	var unrealList = make([]int, 0, 8)
	if err := json.Unmarshal([]byte(pse.UnrealInfo), &unrealList); err != nil {
		return err
	}
	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.advanceId = pse.AdvancedId
	pwo.faBaoId = pse.FaBaoId
	pwo.unrealLevel = pse.UnrealLevel
	pwo.unrealNum = pse.UnrealNum
	pwo.unrealPro = pse.UnrealPro
	pwo.unrealList = unrealList
	pwo.timesNum = pse.TimesNum
	pwo.bless = pse.Bless
	pwo.blessTime = pse.BlessTime
	pwo.tongLingLevel = pse.TongLingLevel
	pwo.tongLingNum = pse.TongLingNum
	pwo.tongLingPro = pse.TongLingPro
	pwo.hidden = pse.Hidden
	pwo.power = pse.Power
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerFaBaoObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("fabao: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

//法宝非进阶对象
type PlayerFaBaoOtherObject struct {
	player     player.Player
	id         int64
	playerId   int64
	typ        fabaotypes.FaBaoType
	faBaoId    int32
	level      int32
	upNum      int32
	upPro      int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerFaBaoOtherObject(pl player.Player) *PlayerFaBaoOtherObject {
	pwo := &PlayerFaBaoOtherObject{
		player: pl,
	}
	return pwo
}

func convertFaBaoOtherObjectToEntity(pwo *PlayerFaBaoOtherObject) (*fabaoentity.PlayerFaBaoOtherEntity, error) {

	e := &fabaoentity.PlayerFaBaoOtherEntity{
		Id:         pwo.id,
		PlayerId:   pwo.playerId,
		Typ:        int32(pwo.typ),
		FaBaoId:    pwo.faBaoId,
		Level:      pwo.level,
		UpNum:      pwo.upNum,
		UpPro:      pwo.upPro,
		UpdateTime: pwo.updateTime,
		CreateTime: pwo.createTime,
		DeleteTime: pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerFaBaoOtherObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerFaBaoOtherObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerFaBaoOtherObject) GetType() fabaotypes.FaBaoType {
	return pwo.typ
}

func (pwo *PlayerFaBaoOtherObject) GetFaBaoId() int32 {
	return pwo.faBaoId
}

func (pwo *PlayerFaBaoOtherObject) GetLevel() int32 {
	return pwo.level
}

func (pwo *PlayerFaBaoOtherObject) GetUpNum() int32 {
	return pwo.upNum
}

func (pwo *PlayerFaBaoOtherObject) GetUpPro() int32 {
	return pwo.upPro
}

func (pwo *PlayerFaBaoOtherObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertFaBaoOtherObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerFaBaoOtherObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*fabaoentity.PlayerFaBaoOtherEntity)

	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.typ = fabaotypes.FaBaoType(pse.Typ)
	pwo.faBaoId = pse.FaBaoId
	pwo.level = pse.Level
	pwo.upNum = pse.UpNum
	pwo.upPro = pse.UpPro
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerFaBaoOtherObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("fabao: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}
