package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	goldequipentity "fgame/fgame/game/goldequip/entity"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	xiuxianbookeventtypes "fgame/fgame/game/welfare/xiuxianbook/event/types"
	"fgame/fgame/pkg/idutil"
	"fmt"

	"github.com/pkg/errors"
	"github.com/willf/bitset"
)

//元神金装身体位置包裹
type BodyBag struct {
	p       player.Player
	slotMap map[inventorytypes.BodyPositionType]*PlayerGoldEquipSlotObject
	//改变的位置
	changedBitset *bitset.BitSet
}

func (bb *BodyBag) GetAll() (slotList []*PlayerGoldEquipSlotObject) {
	for _, slot := range bb.slotMap {
		slotList = append(slotList, slot)
	}
	return
}

//获取强化等级总和
func (bb *BodyBag) GetAllStrengthenLevel() int32 {
	totalLevel := int32(0)
	for _, slot := range bb.slotMap {
		if !slot.IsEmpty() {
			totalLevel += slot.newStLevel
		}
	}
	return totalLevel
}

//获取升星等级总和
func (bb *BodyBag) GetAllUpStarLevel() int32 {
	totalLevel := int32(0)
	for _, slot := range bb.slotMap {
		if !slot.IsEmpty() {
			totalLevel += slot.level
		}
	}
	return totalLevel
}

//获取开光等级总和
func (bb *BodyBag) GetAllOpenlightLevel() int32 {
	totalLevel := int32(0)
	for _, slot := range bb.slotMap {
		if !slot.IsEmpty() {
			data, _ := slot.propertyData.(*goldequiptypes.GoldEquipPropertyData)
			totalLevel += data.OpenLightLevel
		}
	}
	return totalLevel
}

//获取根据位置
func (bb *BodyBag) GetByPosition(pos inventorytypes.BodyPositionType) *PlayerGoldEquipSlotObject {
	eso, exist := bb.slotMap[pos]
	if !exist {
		return nil
	}
	return eso
}

//穿上
func (bb *BodyBag) PutOn(pos inventorytypes.BodyPositionType, itemId int32, level int32, bind itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) bool {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if !itemTemplate.IsGoldEquip() {
		return false
	}

	bodySlot := bb.GetByPosition(pos)
	//位置不存在
	if bodySlot == nil {
		return false
	}

	if bodySlot.IsEmpty() {
		now := global.GetGame().GetTimeService().Now()
		bodySlot.itemId = itemId
		bodySlot.updateTime = now
		bodySlot.level = level
		bodySlot.bind = bind
		bodySlot.propertyData = propertyData
		bodySlot.SetModified()
		bb.changed(int(pos))
		return true
	}

	return false
}

//脱下
func (bb *BodyBag) TakeOff(pos inventorytypes.BodyPositionType) (itemId int32) {
	bodySlot := bb.GetByPosition(pos)
	if bodySlot == nil {
		return
	}
	if bodySlot.IsEmpty() {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	itemId = bodySlot.itemId
	bodySlot.itemId = 0
	bodySlot.level = 0
	bodySlot.propertyData = inventorytypes.CreateDefaultItemPropertyDataBase()
	bodySlot.updateTime = now
	bodySlot.SetModified()
	bb.changed(int(pos))
	return
}

//继承
func (bb *BodyBag) ExtendGoldEquipLevel(pos inventorytypes.BodyPositionType, upstarLevel int32) bool {
	slotItem := bb.GetByPosition(pos)
	if slotItem == nil {
		return false
	}
	if slotItem.IsEmpty() {
		return false
	}
	propertyData, ok := slotItem.propertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	propertyData.UpstarLevel = upstarLevel

	now := global.GetGame().GetTimeService().Now()
	slotItem.updateTime = now
	slotItem.SetModified()
	bb.changed(int(pos))
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, slotItem.player, nil)
	return true
}

//获取下一个新强化配置
func (bb *BodyBag) GetNextStrengthenBuWeiTemplate(pos inventorytypes.BodyPositionType) *gametemplate.GoldEquipStrengthenBuWeiTemplate {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		//正常不可能
		return nil
	}
	var nextStrengthenBuWeiTemplate *gametemplate.GoldEquipStrengthenBuWeiTemplate
	if slot.newStLevel == 0 {
		nextStrengthenBuWeiTemplate = goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipStrengthenBuWeiTemplate(pos, 1)
	} else {
		//判断槽位是否可以升级
		temp := goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipStrengthenBuWeiTemplate(pos, slot.newStLevel)
		nextStrengthenBuWeiTemplate = temp.GetNextTemplate()
	}
	return nextStrengthenBuWeiTemplate
}

//新强化升级
func (bb *BodyBag) StrengthBuWeiLevelUp(pos inventorytypes.BodyPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}
	nextTemp := bb.GetNextStrengthenBuWeiTemplate(pos)
	if nextTemp == nil {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	slot.newStLevel = nextTemp.Level
	slot.updateTime = now
	slot.SetModified()
	bb.changed(int(pos))
	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipStrengUpstarSuccess, bb.p, nil)
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, slot.player, nil)
	return true
}

//新强化升级回退
func (bb *BodyBag) StrengthBuWeiLevelReturn(pos inventorytypes.BodyPositionType, returnLevel int32) bool {
	slotItem := bb.GetByPosition(pos)
	if slotItem == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	slotItem.newStLevel = returnLevel
	slotItem.updateTime = now
	slotItem.SetModified()
	bb.changed(int(pos))
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, slotItem.player, nil)
	return true
}

//新强化修改用于数据迁移
func (bb *BodyBag) SetStrengthBuWeiLevel(pos inventorytypes.BodyPositionType, level int32) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}
	if slot.newStLevel == level {
		return true
	}
	temp := goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipStrengthenBuWeiTemplate(pos, level)
	if temp == nil {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	slot.newStLevel = level
	slot.updateTime = now
	slot.SetModified()
	return true
}

//新强化修改用于Gm
func (bb *BodyBag) GmSetStrengthBuWeiLevel(pos inventorytypes.BodyPositionType, level int32) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}
	if slot.newStLevel == level {
		return true
	}
	temp := goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipStrengthenBuWeiTemplate(pos, level)
	if temp == nil {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	slot.newStLevel = level
	slot.updateTime = now
	slot.SetModified()
	bb.changed(int(pos))
	return true
}

//更新身上升星强化等级
func (bb *BodyBag) UpdateGoldEquipStrengthBuWeiUseItem(pos inventorytypes.BodyPositionType, itemId int32) bool {
	slotItem := bb.GetByPosition(pos)
	if slotItem == nil || slotItem.IsEmpty() {
		return false
	}

	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		return false
	}
	//验证是强化圣石
	if !itemTemp.IsQiangHuaShengShi() {
		return false
	}

	slotItem.newStLevel = itemTemp.TypeFlag1
	now := global.GetGame().GetTimeService().Now()
	slotItem.updateTime = now
	slotItem.SetModified()
	bb.changed(int(pos))

	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipStrengUpstarSuccess, bb.p, nil)
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, slotItem.player, nil)
	return true
}

//更新身上金装等级
func (bb *BodyBag) UpdateGoldEquipLevel(pos inventorytypes.BodyPositionType) bool {
	slotItem := bb.GetByPosition(pos)
	if slotItem == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	slotItem.level += 1
	slotItem.updateTime = now
	slotItem.SetModified()
	bb.changed(int(pos))

	newItemData := droptemplate.CreateItemData(slotItem.itemId, 1, slotItem.level, slotItem.bind)
	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipStrengSuccess, bb.p, newItemData)
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, slotItem.player, nil)
	return true
}

//更新身上金装等级
func (bb *BodyBag) UpdateGoldEquipLevelUseItem(pos inventorytypes.BodyPositionType, itemId int32) bool {
	slotItem := bb.GetByPosition(pos)
	if slotItem == nil {
		return false
	}

	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		return false
	}
	//验证是天工锤
	if !itemTemp.IsTianGongChui() {
		return false
	}

	toUplevel := itemTemp.TypeFlag1
	now := global.GetGame().GetTimeService().Now()
	slotItem.level = toUplevel
	slotItem.updateTime = now
	slotItem.SetModified()
	bb.changed(int(pos))

	newItemData := droptemplate.CreateItemData(slotItem.itemId, 1, slotItem.level, slotItem.bind)
	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipStrengSuccess, bb.p, newItemData)
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, slotItem.player, nil)
	return true
}

//更新身上开光等级
func (bb *BodyBag) UpdateGoldEquipOpenLightUseItem(pos inventorytypes.BodyPositionType, itemId int32) bool {
	slotItem := bb.GetByPosition(pos)
	if slotItem == nil || slotItem.IsEmpty() {
		return false
	}

	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		return false
	}
	//验证是开光钻
	if !itemTemp.IsKaiGuangZuan() {
		return false
	}

	toUplevel := itemTemp.TypeFlag1
	propertyData, ok := slotItem.propertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	propertyData.OpenLightLevel = toUplevel
	propertyData.OpenTimes = 0

	now := global.GetGame().GetTimeService().Now()
	slotItem.updateTime = now
	slotItem.SetModified()
	bb.changed(int(pos))
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, slotItem.player, nil)

	return true
}

//更新身上升星强化等级
func (bb *BodyBag) UpdateGoldEquipUpstarUseItem(pos inventorytypes.BodyPositionType, itemId int32) bool {
	slotItem := bb.GetByPosition(pos)
	if slotItem == nil || slotItem.IsEmpty() {
		return false
	}

	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		return false
	}
	//验证是强化圣石
	if !itemTemp.IsQiangHuaShengShi() {
		return false
	}

	toUplevel := itemTemp.TypeFlag1
	propertyData, ok := slotItem.propertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	propertyData.UpstarLevel = toUplevel

	now := global.GetGame().GetTimeService().Now()
	slotItem.updateTime = now
	slotItem.SetModified()
	bb.changed(int(pos))

	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipStrengUpstarSuccess, bb.p, nil)
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, slotItem.player, nil)
	return true
}

//金装升星强化
func (bb *BodyBag) UpstarSuccess(pos inventorytypes.BodyPositionType) bool {
	slotItem := bb.GetByPosition(pos)
	if slotItem == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	propertyData := slotItem.propertyData.(*goldequiptypes.GoldEquipPropertyData)
	propertyData.UpstarLevel += 1
	slotItem.updateTime = now
	slotItem.SetModified()
	bb.changed(int(pos))

	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipStrengUpstarSuccess, bb.p, nil)
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, slotItem.player, nil)
	return true
}

//金装升星强化失败
func (bb *BodyBag) UpstarFaildReturn(pos inventorytypes.BodyPositionType, returnLevel int32) bool {
	slotItem := bb.GetByPosition(pos)
	if slotItem == nil {
		return false
	}
	propertyData := slotItem.propertyData.(*goldequiptypes.GoldEquipPropertyData)
	propertyData.UpstarLevel = returnLevel

	now := global.GetGame().GetTimeService().Now()
	slotItem.updateTime = now
	slotItem.SetModified()
	bb.changed(int(pos))
	return true
}

func (bb *BodyBag) GetChangedSlotAndReset() []*PlayerGoldEquipSlotObject {
	itemList := make([]*PlayerGoldEquipSlotObject, 0, 16)
	for i, valid := bb.changedBitset.NextSet(0); valid; i, valid = bb.changedBitset.NextSet(i + 1) {
		itemList = append(itemList, bb.slotMap[inventorytypes.BodyPositionType(i)])
	}
	bb.Reset()
	return itemList
}

func (bb *BodyBag) Reset() {
	bb.changedBitset.ClearAll()
}

// 是否有装备
func (bb *BodyBag) IsHadEquipment() bool {
	for _, slot := range bb.slotMap {
		if slot.itemId != 0 {
			return true
		}
	}

	return false
}

//设置改变
func (bb *BodyBag) changed(index int) {
	bb.changedBitset.Set(uint(index))
}

//是否有宝石
func (bb *BodyBag) IfEmbedGem(pos inventorytypes.BodyPositionType, order int32) (flag bool) {
	item := bb.GetByPosition(pos)
	if item == nil {
		return
	}
	if item.IsEmpty() {
		return
	}
	//获取装备槽宝石信息
	_, exist := item.GemInfo[order]
	if !exist {
		return
	}

	return true
}

//解锁宝石槽
func (bb *BodyBag) IsUnlockGem(pos inventorytypes.BodyPositionType, order int32) bool {
	item := bb.GetByPosition(pos)
	if item == nil {
		return false
	}
	//获取宝石槽信息
	_, exist := item.GemUnlockInfo[order]
	return exist
}

//解锁宝石槽
func (bb *BodyBag) UnlockGem(pos inventorytypes.BodyPositionType, order int32) {
	item := bb.GetByPosition(pos)
	if item == nil {
		return
	}
	//获取宝石槽信息
	_, exist := item.GemUnlockInfo[order]
	if exist {
		return
	}
	item.GemUnlockInfo[order] = 1
	now := global.GetGame().GetTimeService().Now()
	item.updateTime = now
	item.SetModified()
	bb.changed(int(pos))
	return
}

//脱下宝石
func (bb *BodyBag) TakeOffGem(pos inventorytypes.BodyPositionType, order int32) (itemId int32) {
	item := bb.GetByPosition(pos)
	if item == nil {
		return
	}
	if item.IsEmpty() {
		return
	}
	//获取装备槽宝石信息
	itemId, exist := item.GemInfo[order]
	if !exist {
		return
	}
	delete(item.GemInfo, order)
	now := global.GetGame().GetTimeService().Now()
	item.updateTime = now
	item.SetModified()
	bb.changed(int(pos))
	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipTakeOffGem, bb.p, nil)
	return
}

//佩戴宝石
func (bb *BodyBag) PutOnGem(pos inventorytypes.BodyPositionType, order int32, itemId int32) bool {
	//物品是不是宝石
	itemTemp := item.GetItemService().GetItem(int(itemId))
	if !itemTemp.IsGem() {
		return false
	}

	item := bb.GetByPosition(pos)
	//位置不存在
	if item == nil {
		return false
	}

	//装备不存在
	if item.IsEmpty() {
		return false
	}

	//获取装备槽宝石信息
	_, exist := item.GemInfo[order]
	if exist {
		return false
	}
	item.GemInfo[order] = itemId
	now := global.GetGame().GetTimeService().Now()
	item.updateTime = now
	item.SetModified()
	bb.changed(int(pos))

	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipEmbedGem, bb.p, nil)
	return true
}

//神铸升级
func (bb *BodyBag) UplevelGodCasting(bodyPos inventorytypes.BodyPositionType, itemId int32, sucess bool) bool {
	o := bb.GetByPosition(bodyPos)
	data, ok := o.propertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	if sucess {
		data.GodCastingTimes = 0
		o.itemId = itemId
	} else {
		data.GodCastingTimes = data.GodCastingTimes + 1
	}
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
	bb.changed(int(bodyPos))
	return true
}

//铸灵升级
func (bb *BodyBag) UplevelSpirit(bodyPos inventorytypes.BodyPositionType, spiritType goldequiptypes.SpiritType, sucess bool, bless int32) {
	o := bb.GetByPosition(bodyPos)
	info := o.CastingSpiritInfo[spiritType]
	if sucess {
		info.Level = info.Level + 1
		info.Times = 0
		info.Bless = 0
	} else {
		info.Times = info.Times + 1
		info.Bless += bless
	}
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
	bb.changed(int(bodyPos))
}

//锻魂升级
func (bb *BodyBag) UplevelSoul(bodyPos inventorytypes.BodyPositionType, soulType goldequiptypes.ForgeSoulType, sucess bool) {
	o := bb.GetByPosition(bodyPos)
	info := o.ForgeSoulInfo[soulType]
	if sucess {
		info.Level = info.Level + 1
		info.Times = 0
	} else {
		info.Times = info.Times + 1
	}
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
	bb.changed(int(bodyPos))
}

//神铸继承
func (bb *BodyBag) GodCastingInherit(bodyPos inventorytypes.BodyPositionType, itemId int32) {
	o := bb.GetByPosition(bodyPos)
	o.itemId = itemId
	now := global.GetGame().GetTimeService().Now()
	o.updateTime = now
	o.SetModified()
	bb.changed(int(bodyPos))
}

//创建身体背包
func createBodyBag(p player.Player, slotList []*PlayerGoldEquipSlotObject) *BodyBag {
	bb := &BodyBag{
		p: p,
	}

	bb.init(slotList)
	return bb
}

//初始化
func (bb *BodyBag) init(slotList []*PlayerGoldEquipSlotObject) {
	bb.changedBitset = bitset.New(64)
	bb.slotMap = make(map[inventorytypes.BodyPositionType]*PlayerGoldEquipSlotObject)
	for _, slot := range slotList {
		bb.slotMap[slot.slotId] = slot
	}
	now := global.GetGame().GetTimeService().Now()
	for slotId := inventorytypes.BodyPositionTypeWeapon; slotId <= inventorytypes.BodyPositionTypeRing; slotId++ {
		if bb.GetByPosition(slotId) != nil {
			continue
		}
		slot := createGoldEquipSlotObject(bb.p, slotId, now)
		slot.SetModified()
		bb.slotMap[slot.slotId] = slot
	}
}

func createGoldEquipSlotObject(p player.Player, slotId inventorytypes.BodyPositionType, now int64) *PlayerGoldEquipSlotObject {
	itemObject := NewPlayerGoldEquipSlotObject(p)
	itemObject.createTime = now
	itemObject.id, _ = idutil.GetId()
	itemObject.itemId = 0
	itemObject.slotId = slotId

	base := inventorytypes.CreateDefaultItemPropertyDataBase()
	itemObject.propertyData = inventory.CreatePropertyDataInterface(itemtypes.ItemTypeGoldEquip, base)
	itemObject.GemInfo = make(map[int32]int32)
	itemObject.GemUnlockInfo = make(map[int32]int32)
	itemObject.CastingSpiritInfo = make(map[goldequiptypes.SpiritType]*goldequiptypes.CastingSpiritInfo)
	itemObject.ForgeSoulInfo = make(map[goldequiptypes.ForgeSoulType]*goldequiptypes.ForgeSoulInfo)
	itemObject.createTime = now
	return itemObject
}

//玩家槽位数据
type PlayerGoldEquipSlotObject struct {
	player            player.Player
	id                int64
	playerId          int64
	slotId            inventorytypes.BodyPositionType
	itemId            int32
	level             int32 //升星等级
	newStLevel        int32 //强化等级（槽位）
	bind              itemtypes.ItemBindType
	propertyData      inventorytypes.ItemPropertyData
	GemInfo           map[int32]int32
	GemUnlockInfo     map[int32]int32
	CastingSpiritInfo map[goldequiptypes.SpiritType]*goldequiptypes.CastingSpiritInfo //铸灵等级祝福值等信息
	ForgeSoulInfo     map[goldequiptypes.ForgeSoulType]*goldequiptypes.ForgeSoulInfo  //锻魂等级信息
	updateTime        int64
	createTime        int64
	deleteTime        int64
}

func NewPlayerGoldEquipSlotObject(pl player.Player) *PlayerGoldEquipSlotObject {
	o := &PlayerGoldEquipSlotObject{
		player:   pl,
		playerId: pl.GetId(),
	}
	return o
}

func convertPlayerGoldEquipSlotObjectToEntity(o *PlayerGoldEquipSlotObject) (*goldequipentity.PlayerGoldEquipSlotEntity, error) {
	data, err := json.Marshal(o.propertyData)
	if err != nil {
		return nil, err
	}

	gemInfoBytes, err := json.Marshal(o.GemInfo)
	if err != nil {
		return nil, err
	}

	gemUnlockInfoBytes, err := json.Marshal(o.GemUnlockInfo)
	if err != nil {
		return nil, err
	}

	castingSpiritInfoBytes, err := json.Marshal(o.CastingSpiritInfo)
	if err != nil {
		return nil, err
	}

	forgeSoulInfoBytes, err := json.Marshal(o.ForgeSoulInfo)
	if err != nil {
		return nil, err
	}

	e := &goldequipentity.PlayerGoldEquipSlotEntity{
		Id:                o.id,
		PlayerId:          o.playerId,
		ItemId:            o.itemId,
		SlotId:            int32(o.slotId),
		Level:             o.level,
		NewStLevel:        o.newStLevel,
		BindType:          int32(o.bind),
		PropertyData:      string(data),
		GemInfo:           string(gemInfoBytes),
		GemUnlockInfo:     string(gemUnlockInfoBytes),
		CastingSpiritInfo: string(castingSpiritInfoBytes),
		ForgeSoulInfo:     string(forgeSoulInfoBytes),
		UpdateTime:        o.updateTime,
		CreateTime:        o.createTime,
		DeleteTime:        o.deleteTime,
	}
	return e, nil
}

func (o *PlayerGoldEquipSlotObject) GetPlayerId() int64 {
	return o.playerId
}

func (o *PlayerGoldEquipSlotObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerGoldEquipSlotObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerGoldEquipSlotObject) GetCastingSpiritInfo(spiritType goldequiptypes.SpiritType) *goldequiptypes.CastingSpiritInfo {
	return o.CastingSpiritInfo[spiritType]
}

func (o *PlayerGoldEquipSlotObject) GetForgeSoulInfo(soulType goldequiptypes.ForgeSoulType) *goldequiptypes.ForgeSoulInfo {
	return o.ForgeSoulInfo[soulType]
}

func (o *PlayerGoldEquipSlotObject) GetNewStLevel() int32 {
	return o.newStLevel
}

func (o *PlayerGoldEquipSlotObject) GetBindType() itemtypes.ItemBindType {
	return o.bind
}

func (o *PlayerGoldEquipSlotObject) GetItemId() int32 {
	return o.itemId
}

func (o *PlayerGoldEquipSlotObject) GetSlotId() inventorytypes.BodyPositionType {
	return o.slotId
}

func (o *PlayerGoldEquipSlotObject) GetPropertyData() inventorytypes.ItemPropertyData {
	return o.propertyData
}

func (o *PlayerGoldEquipSlotObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerGoldEquipSlotObjectToEntity(o)
	return
}

func (o *PlayerGoldEquipSlotObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*goldequipentity.PlayerGoldEquipSlotEntity)
	data, err := inventory.CreatePropertyData(itemtypes.ItemTypeGoldEquip, pse.PropertyData)
	if err != nil {
		return
	}

	gemInfo := make(map[int32]int32)
	err = json.Unmarshal([]byte(pse.GemInfo), &gemInfo)
	if err != nil {
		return
	}

	gemUnlockInfo := make(map[int32]int32)
	err = json.Unmarshal([]byte(pse.GemUnlockInfo), &gemUnlockInfo)
	if err != nil {
		return
	}

	castingSpiritInfo := make(map[goldequiptypes.SpiritType]*goldequiptypes.CastingSpiritInfo)
	for i := goldequiptypes.MinSpiritType; i <= goldequiptypes.MaxSpiritType; i++ {
		castingSpiritInfo[i] = &goldequiptypes.CastingSpiritInfo{}
	}
	err = json.Unmarshal([]byte(pse.CastingSpiritInfo), &castingSpiritInfo)
	if err != nil {
		return
	}

	forgeSoulInfo := make(map[goldequiptypes.ForgeSoulType]*goldequiptypes.ForgeSoulInfo)
	for i := goldequiptypes.MinForgeSoulType; i <= goldequiptypes.MaxForgeSoulType; i++ {
		forgeSoulInfo[i] = &goldequiptypes.ForgeSoulInfo{}
	}
	err = json.Unmarshal([]byte(pse.ForgeSoulInfo), &forgeSoulInfo)
	if err != nil {
		return
	}

	o.id = pse.Id
	o.playerId = pse.PlayerId
	o.itemId = pse.ItemId
	o.slotId = inventorytypes.BodyPositionType(pse.SlotId)
	o.level = pse.Level
	o.newStLevel = pse.NewStLevel
	o.bind = itemtypes.ItemBindType(pse.BindType)
	o.GemInfo = gemInfo
	o.GemUnlockInfo = gemUnlockInfo
	o.CastingSpiritInfo = castingSpiritInfo
	o.ForgeSoulInfo = forgeSoulInfo
	o.propertyData = data
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return
}

func (o *PlayerGoldEquipSlotObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "GoldEquipSlot"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	o.player.AddChangedObject(obj)
	return
}

func (o *PlayerGoldEquipSlotObject) IsEmpty() bool {
	return o.itemId == 0
}

func (o *PlayerGoldEquipSlotObject) IsFull() bool {
	return o.itemId != 0
}
