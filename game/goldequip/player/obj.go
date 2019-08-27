package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	goldequipentity "fgame/fgame/game/goldequip/entity"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"

	"github.com/pkg/errors"
)

//元神金装日志列表对象
type PlayerGoldEquipLogObject struct {
	player           player.Player
	id               int64
	fenJieItemIdList []int32
	rewItemStr       string
	updateTime       int64
	createTime       int64
	deleteTime       int64
}

func NewPlayerGoldEquipLogObject(p player.Player) *PlayerGoldEquipLogObject {
	o := &PlayerGoldEquipLogObject{}
	o.player = p
	return o
}

func convertNewPlayerGoldEquipLogObjectToEntity(o *PlayerGoldEquipLogObject) (*goldequipentity.PlayerGoldEquipLogEntity, error) {
	data, err := json.Marshal(o.fenJieItemIdList)
	if err != nil {
		return nil, err
	}
	e := &goldequipentity.PlayerGoldEquipLogEntity{
		Id:           o.id,
		PlayerId:     o.player.GetId(),
		FenJieItemId: string(data),
		RewItemStr:   o.rewItemStr,
		UpdateTime:   o.updateTime,
		CreateTime:   o.createTime,
		DeleteTime:   o.deleteTime,
	}
	return e, nil
}

func (o *PlayerGoldEquipLogObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *PlayerGoldEquipLogObject) GetRewItemStr() string {
	return o.rewItemStr
}
func (o *PlayerGoldEquipLogObject) GetFenJieItemIdList() []int32 {
	return o.fenJieItemIdList
}

func (o *PlayerGoldEquipLogObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *PlayerGoldEquipLogObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerGoldEquipLogObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerGoldEquipLogObjectToEntity(o)
	return e, err
}

func (o *PlayerGoldEquipLogObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*goldequipentity.PlayerGoldEquipLogEntity)

	var itemIdList []int32
	err := json.Unmarshal([]byte(pse.FenJieItemId), &itemIdList)
	if err != nil {
		return err
	}

	o.id = pse.Id
	o.fenJieItemIdList = itemIdList
	o.rewItemStr = pse.RewItemStr
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerGoldEquipLogObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "GoldEquipLog"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)

	return
}

//元神金装设置对象
type PlayerGoldEquipSettingObject struct {
	player         player.Player
	id             int64
	fenJieIsAuto   int32
	fenJieQuality  itemtypes.ItemQualityType
	fenJieZhuanShu int32
	isCheckOldSt   int32
	updateTime     int64
	createTime     int64
	deleteTime     int64
}

func NewPlayerGoldEquipSettingObject(p player.Player) *PlayerGoldEquipSettingObject {
	o := &PlayerGoldEquipSettingObject{}
	o.player = p
	return o
}

func convertNewPlayerGoldEquipSettingObjectToEntity(o *PlayerGoldEquipSettingObject) (*goldequipentity.PlayerGoldEquipSettingEntity, error) {
	e := &goldequipentity.PlayerGoldEquipSettingEntity{
		Id:             o.id,
		PlayerId:       o.player.GetId(),
		FenJieIsAuto:   o.fenJieIsAuto,
		FenJieQuality:  int32(o.fenJieQuality),
		FenJieZhuanShu: int32(o.fenJieZhuanShu),
		IsCheckOldSt:   o.isCheckOldSt,
		UpdateTime:     o.updateTime,
		CreateTime:     o.createTime,
		DeleteTime:     o.deleteTime,
	}
	return e, nil
}

func (o *PlayerGoldEquipSettingObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *PlayerGoldEquipSettingObject) GetFenJieQuality() itemtypes.ItemQualityType {
	return o.fenJieQuality
}

func (o *PlayerGoldEquipSettingObject) GetFenJieZhuanShu() int32 {
	return o.fenJieZhuanShu
}

func (o *PlayerGoldEquipSettingObject) GetFenJieIsAuto() int32 {
	return o.fenJieIsAuto
}

func (o *PlayerGoldEquipSettingObject) GetIsCheckOldSt() int32 {
	return o.isCheckOldSt
}

func (o *PlayerGoldEquipSettingObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *PlayerGoldEquipSettingObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerGoldEquipSettingObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerGoldEquipSettingObjectToEntity(o)
	return e, err
}

func (o *PlayerGoldEquipSettingObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*goldequipentity.PlayerGoldEquipSettingEntity)

	o.id = pse.Id
	o.fenJieQuality = itemtypes.ItemQualityType(pse.FenJieQuality)
	o.fenJieIsAuto = pse.FenJieIsAuto
	o.fenJieZhuanShu = pse.FenJieZhuanShu
	o.isCheckOldSt = pse.IsCheckOldSt
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerGoldEquipSettingObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "GoldEquipSetting"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("never reach here"))
	}

	o.player.AddChangedObject(obj)

	return
}

func (o *PlayerGoldEquipSettingObject) IsFenJieAuto() bool {

	return o.fenJieIsAuto != 0
}

func (o *PlayerGoldEquipSettingObject) IsCheckOldStLev() bool {

	return o.isCheckOldSt != 0
}

func (o *PlayerGoldEquipSettingObject) SetIsCheckOldSt() {
	now := global.GetGame().GetTimeService().Now()
	o.isCheckOldSt = int32(1)
	o.updateTime = now
	o.SetModified()
}

//元神金装对象
type PlayerGoldEquipObject struct {
	player     player.Player
	id         int64
	power      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerGoldEquipObject(p player.Player) *PlayerGoldEquipObject {
	o := &PlayerGoldEquipObject{}
	o.player = p
	return o
}

func convertNewPlayerGoldEquipObjectToEntity(o *PlayerGoldEquipObject) (*goldequipentity.PlayerGoldEquipEntity, error) {
	e := &goldequipentity.PlayerGoldEquipEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		Power:      o.power,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerGoldEquipObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerGoldEquipObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerGoldEquipObjectToEntity(o)
	return e, err
}

func (o *PlayerGoldEquipObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*goldequipentity.PlayerGoldEquipEntity)

	o.id = pse.Id
	o.power = pse.Power
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerGoldEquipObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "GoldEquip"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("never reach here"))
	}

	o.player.AddChangedObject(obj)

	return
}
