package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/common/common"
	houseentity "fgame/fgame/game/house/entity"
	housetemplate "fgame/fgame/game/house/template"
	housetypes "fgame/fgame/game/house/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/mathutils"

	"github.com/pkg/errors"
)

//房子对象
type PlayerHouseObject struct {
	player            player.Player
	id                int64
	houseIndex        int32
	houseType         housetypes.HouseType
	level             int32
	maxLevel          int32
	dayTimes          int32
	isBroken          int32
	lastBrokenTime    int64
	isRent            int32 //使用0否1是表达
	refreshUpdateTime int64
	updateTime        int64
	createTime        int64
	deleteTime        int64
}

func NewPlayerHouseObject(pl player.Player) *PlayerHouseObject {
	o := &PlayerHouseObject{
		player: pl,
	}
	return o
}

func convertNewPlayerHouseObjectToEntity(o *PlayerHouseObject) (*houseentity.PlayerHouseEntity, error) {

	e := &houseentity.PlayerHouseEntity{
		Id:                o.id,
		PlayerId:          o.player.GetId(),
		HouseIndex:        o.houseIndex,
		HouseType:         int32(o.houseType),
		Level:             o.level,
		MaxLevel:          o.maxLevel,
		DayTimes:          o.dayTimes,
		IsBroken:          o.isBroken,
		LastBrokenTime:    o.lastBrokenTime,
		IsRent:            o.isRent,
		RefreshUpdateTime: o.refreshUpdateTime,
		UpdateTime:        o.updateTime,
		CreateTime:        o.createTime,
		DeleteTime:        o.deleteTime,
	}
	return e, nil
}

func (o *PlayerHouseObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerHouseObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerHouseObject) GetHouseIndex() int32 {
	return o.houseIndex
}

func (o *PlayerHouseObject) GetIsBroken() int32 {
	return o.isBroken
}

func (o *PlayerHouseObject) GetIsRent() int32 {
	return o.isRent
}

func (o *PlayerHouseObject) GetHouseLevel() int32 {
	return o.level
}

func (o *PlayerHouseObject) GetHouseMaxLevel() int32 {
	return o.maxLevel
}

func (o *PlayerHouseObject) GetHouseType() housetypes.HouseType {
	return o.houseType
}

func (o *PlayerHouseObject) GetDayTimes() int32 {
	return o.dayTimes
}

func (o *PlayerHouseObject) GetRefreshUpdateTime() int64 {
	return o.refreshUpdateTime
}

func (o *PlayerHouseObject) IsBroken() bool {
	return o.isBroken == 1
}

func (o *PlayerHouseObject) IsRent() bool {
	return o.isRent == 1
}

func (o *PlayerHouseObject) IsActivate() bool {
	return o.level > 0
}

func (o *PlayerHouseObject) IfCanBroken(now int64) bool {
	brokenCd := housetemplate.GetHouseTemplateService().GetHouseConstantTemplate().BrokenCd
	if now-o.lastBrokenTime <= brokenCd {
		return false
	}

	houseTemp := housetemplate.GetHouseTemplateService().GetHouseTemplate(o.houseIndex, o.houseType, o.level)
	if houseTemp == nil {
		return false
	}

	rate := int(houseTemp.BrokenPercent)
	isBroken := mathutils.RandomHit(int(common.MAX_RATE), rate)
	return isBroken
}

func (o *PlayerHouseObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerHouseObjectToEntity(o)
	return e, err
}

func (o *PlayerHouseObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*houseentity.PlayerHouseEntity)

	o.id = pse.Id
	o.houseIndex = pse.HouseIndex
	o.houseType = housetypes.HouseType(pse.HouseType)
	o.level = pse.Level
	o.maxLevel = pse.MaxLevel
	o.dayTimes = pse.DayTimes
	o.isBroken = pse.IsBroken
	o.lastBrokenTime = pse.LastBrokenTime
	o.isRent = pse.IsRent
	o.refreshUpdateTime = pse.RefreshUpdateTime
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerHouseObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "House"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
