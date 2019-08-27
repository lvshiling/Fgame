package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	itemservice "fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	wushuangweaponentity "fgame/fgame/game/wushuangweapon/entity"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

type PlayerWushuangWeaponSlotObject struct {
	player     player.Player
	id         int64
	bodyPart   wushuangweapontypes.WushuangWeaponPart
	itemId     int32
	isActive   bool
	level      int32
	bindType   itemtypes.ItemBindType
	experience int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func (o *PlayerWushuangWeaponSlotObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerWushuangWeaponSlotObject) GetLevel() int32 {
	return o.level
}

// func (o *PlayerWushuangWeaponSlotObject) ChangeLevel(level int32) {
// 	o.level = level
// 	now := global.GetGame().GetTimeService().Now()
// 	o.updateTime = now
// 	o.SetModified()
// 	return
// }

func (o *PlayerWushuangWeaponSlotObject) GetBindType() itemtypes.ItemBindType {
	return o.bindType
}

func (o *PlayerWushuangWeaponSlotObject) ChangeLevel(level int32) {
	o.level = level
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
}

func (o *PlayerWushuangWeaponSlotObject) GetExperience() int64 {
	return o.experience
}

func (o *PlayerWushuangWeaponSlotObject) GetItemId() int32 {
	return o.itemId
}

//是否装备物品
func (o *PlayerWushuangWeaponSlotObject) IsEquip() bool {
	if o.itemId == 0 {
		return false
	} else {
		return true
	}
}

func (o *PlayerWushuangWeaponSlotObject) AddExperience(ex int64) {
	o.experience += ex
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
	return
}

func (o *PlayerWushuangWeaponSlotObject) Uplevel() {
	o.level += 1
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
	return
}

func (o *PlayerWushuangWeaponSlotObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerWushuangWeaponSlotObject) GetBodyPart() wushuangweapontypes.WushuangWeaponPart {
	return o.bodyPart
}

// 穿上物品修改槽位数据
func (o *PlayerWushuangWeaponSlotObject) PatchFromItemTemplate(itemTemp *gametemplate.ItemTemplate) bool {
	if o.itemId != 0 {
		return false
	}
	o.itemId = int32(itemTemp.Id)
	o.bindType = itemTemp.GetBindType()
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()

	return true
}

// 穿上物品修改槽位数据
func (o *PlayerWushuangWeaponSlotObject) PutOn(itemId int32, bindType itemtypes.ItemBindType, settingsLevel int32) bool {
	if o.itemId != 0 {
		return false
	}
	if itemId == 0 {
		return false
	}
	itemTemp := itemservice.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		return false
	}
	baseTemplate := itemTemp.GetWushuangBaseTemplate()
	if baseTemplate == nil {
		return false
	}
	o.itemId = itemId
	o.bindType = bindType
	o.updateLevel(settingsLevel)
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
	return true
}

func (o *PlayerWushuangWeaponSlotObject) updateLevel(settingsLevel int32) {
	itemId := o.itemId
	itemTemp := itemservice.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		return
	}
	baseTemplate := itemTemp.GetWushuangBaseTemplate()
	if baseTemplate == nil {
		return
	}

	lev, isBorder := baseTemplate.GetLevel(o.experience)
	if !isBorder {
		o.level = lev
	} else {
		// 解决正好是边界的问题，其余经验增加经验减少都走正常
		if lev == (settingsLevel - 1) {
			o.level = lev + 1
		} else {
			o.level = lev
		}
	}

}

func (o *PlayerWushuangWeaponSlotObject) TakeOff() {
	o.itemId = int32(0)
	o.bindType = itemtypes.ItemBindTypeUnBind
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
}

func (o *PlayerWushuangWeaponSlotObject) ToEntity() (e storage.Entity, err error) {
	isActive := int32(0)
	if o.isActive {
		isActive = 1
	}
	// switch o.isActive {
	// case false:
	// 	isActive = int32(0)
	// case true:
	// 	isActive = int32(1)
	// }
	e = &wushuangweaponentity.PlayerWushuangWeaponSlotEntity{
		Id:         o.id,
		PlayerId:   o.player.GetId(),
		ItemId:     o.itemId,
		SlotId:     int32(o.bodyPart),
		Level:      o.level,
		Experience: o.experience,
		Bind:       int32(o.bindType),
		IsActive:   isActive,
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *PlayerWushuangWeaponSlotObject) FromEntity(e storage.Entity) (err error) {
	te, _ := e.(*wushuangweaponentity.PlayerWushuangWeaponSlotEntity)

	isActive := te.IsActive != 0
	//修改
	// switch te.IsActive {
	// case int32(0):
	// 	isActive = false
	// case int32(1):
	// 	isActive = true
	// }
	o.id = te.Id
	o.itemId = te.ItemId
	o.bodyPart = wushuangweapontypes.WushuangWeaponPart(te.SlotId)
	o.level = te.Level
	o.experience = te.Experience
	o.bindType = itemtypes.ItemBindType(te.Bind)
	o.isActive = isActive
	o.updateTime = te.UpdateTime
	o.createTime = te.CreateTime
	o.deleteTime = te.DeleteTime
	return nil
}

func NewPlayerWushuangWeaponSlotObject(pl player.Player) *PlayerWushuangWeaponSlotObject {
	pwwso := &PlayerWushuangWeaponSlotObject{
		player: pl,
	}
	return pwwso
}

func initNewPlayerWushuangWeaponSlotObject(newObj *PlayerWushuangWeaponSlotObject, bodyPart wushuangweapontypes.WushuangWeaponPart) {
	now := global.GetGame().GetTimeService().Now()
	newObj.id, _ = idutil.GetId()
	newObj.itemId = int32(0)
	newObj.bodyPart = bodyPart
	newObj.level = int32(0)
	newObj.experience = int64(0)
	newObj.bindType = itemtypes.ItemBindTypeUnBind
	newObj.isActive = false
	newObj.createTime = now
	newObj.SetModified()
}

func (o *PlayerWushuangWeaponSlotObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "WushuangWeapon"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}
