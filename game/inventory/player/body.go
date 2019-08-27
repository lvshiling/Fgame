package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	gameentity "fgame/fgame/game/inventory/entity"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"

	"fgame/fgame/pkg/idutil"
	"fmt"

	"github.com/willf/bitset"
)

//身体位置包裹
type BodyBag struct {
	p       player.Player
	slotMap map[inventorytypes.BodyPositionType]*PlayerEquipmentSlotObject
	//改变的位置
	changedBitset *bitset.BitSet
}

func (bb *BodyBag) GetAll() (slotList []*PlayerEquipmentSlotObject) {
	for _, slot := range bb.slotMap {
		slotList = append(slotList, slot)
	}
	return
}

func (bb *BodyBag) ClearAllEquipmentGemInfo() {
	now := global.GetGame().GetTimeService().Now()
	for _, slot := range bb.slotMap {
		if len(slot.GemInfo) > 0 {
			slot.GemInfo = map[int32]int32{}
			slot.UpdateTime = now
			slot.SetModified()
		}
	}
	return
}

func (bb *BodyBag) Reset() {
	bb.changedBitset.ClearAll()
}

//设置改变
func (bb *BodyBag) changed(index int) {
	bb.changedBitset.Set(uint(index))
}

func (bb *BodyBag) GetChangedSlotAndReset() []*PlayerEquipmentSlotObject {
	itemList := make([]*PlayerEquipmentSlotObject, 0, 16)
	for i, valid := bb.changedBitset.NextSet(0); valid; i, valid = bb.changedBitset.NextSet(i + 1) {
		itemList = append(itemList, bb.slotMap[inventorytypes.BodyPositionType(i)])
	}
	bb.Reset()
	return itemList
}

//获取根据位置
func (bb *BodyBag) GetByPosition(pos inventorytypes.BodyPositionType) *PlayerEquipmentSlotObject {
	eso, exist := bb.slotMap[pos]
	if !exist {
		return nil
	}
	return eso
}

//是否可以升星
func (bb *BodyBag) IfCanStrengthStar(pos inventorytypes.BodyPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}
	return bb.GetNextStarEquipStrengthenTemplate(pos) != nil

}

//是否可以强化升级
func (bb *BodyBag) IfCanStrengthLevel(pos inventorytypes.BodyPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}

	return bb.GetNextUpgradeEquipStrengthenTemplate(pos) != nil
}

//获取下一个星级
func (bb *BodyBag) GetNextStarEquipStrengthenTemplate(pos inventorytypes.BodyPositionType) *gametemplate.EquipStrengthenTemplate {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		//正常不可能
		return nil
	}
	var nextEquipmentStrengthenTemplate *gametemplate.EquipStrengthenTemplate
	if slot.Star == 0 {
		nextEquipmentStrengthenTemplate = item.GetItemService().GetEquipStrengthenTemplate(inventorytypes.EquipmentStrengthenTypeStar, pos, 1)
	} else {
		//判断槽位是否可以升级
		equipmentStrengthenTemplate := item.GetItemService().GetEquipStrengthenTemplate(inventorytypes.EquipmentStrengthenTypeStar, pos, slot.Star)
		nextEquipmentStrengthenTemplate = equipmentStrengthenTemplate.GetNextEquipStrengthenTemplate()
	}
	return nextEquipmentStrengthenTemplate
}

//获取下一个强化升级
func (bb *BodyBag) GetNextUpgradeEquipStrengthenTemplate(pos inventorytypes.BodyPositionType) *gametemplate.EquipStrengthenTemplate {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		//正常不可能
		return nil
	}
	var nextEquipmentStrengthenTemplate *gametemplate.EquipStrengthenTemplate
	if slot.Level == 0 {
		nextEquipmentStrengthenTemplate = item.GetItemService().GetEquipStrengthenTemplate(inventorytypes.EquipmentStrengthenTypeUpgrade, pos, 1)
	} else {
		//判断槽位是否可以升级
		equipmentStrengthenTemplate := item.GetItemService().GetEquipStrengthenTemplate(inventorytypes.EquipmentStrengthenTypeUpgrade, pos, slot.Level)
		nextEquipmentStrengthenTemplate = equipmentStrengthenTemplate.GetNextEquipStrengthenTemplate()
	}
	return nextEquipmentStrengthenTemplate
}

//强化升星
func (bb *BodyBag) StrengthStar(pos inventorytypes.BodyPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		//正常不可能
		return false
	}
	var nextEquipmentStrengthenTemplate = bb.GetNextStarEquipStrengthenTemplate(pos)

	if nextEquipmentStrengthenTemplate == nil {
		return false
	}
	slot.Star = nextEquipmentStrengthenTemplate.Level
	slot.SetModified()
	bb.changed(int(pos))
	gameevent.Emit(inventoryeventtypes.EventTypeEquipmentUpgradeStar, bb.p, pos)
	return true
}

//强化升星回退
func (bb *BodyBag) StrengthStarBack(pos inventorytypes.BodyPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}
	var nextEquipmentStrengthenTemplate = bb.GetNextStarEquipStrengthenTemplate(pos)
	if nextEquipmentStrengthenTemplate == nil {
		return false
	}

	failEquipmentStrengthenTemplate := nextEquipmentStrengthenTemplate.GetFailEquipStrengthenTemplate()
	if failEquipmentStrengthenTemplate == nil {
		return false
	}

	slot.Star = failEquipmentStrengthenTemplate.Level
	slot.SetModified()
	bb.changed(int(pos))
	return true
}

//强化升级
func (bb *BodyBag) StrengthLevel(pos inventorytypes.BodyPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}
	//不能升星
	nextEquipmentStrengthenTemplate := bb.GetNextUpgradeEquipStrengthenTemplate(pos)

	if nextEquipmentStrengthenTemplate == nil {
		return false
	}
	slot.Level = nextEquipmentStrengthenTemplate.Level
	slot.SetModified()
	bb.changed(int(pos))
	gameevent.Emit(inventoryeventtypes.EventTypeEquipmentStrengthenLevel, bb.p, pos)
	return true
}

//强化升级回退
func (bb *BodyBag) StrengthLevelBack(pos inventorytypes.BodyPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}

	nextEquipmentStrengthenTemplate := bb.GetNextUpgradeEquipStrengthenTemplate(pos)
	if nextEquipmentStrengthenTemplate == nil {
		return false
	}
	failEquipmentStrengthenTemplate := nextEquipmentStrengthenTemplate.GetFailEquipStrengthenTemplate()
	if failEquipmentStrengthenTemplate == nil {
		return false
	}
	slot.Level = failEquipmentStrengthenTemplate.Level
	slot.SetModified()
	bb.changed(int(pos))
	return true
}

//是否可以进阶
func (bb *BodyBag) IfCanUpgrade(pos inventorytypes.BodyPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}
	if slot.IsEmpty() {
		return false
	}
	equipTemplate := item.GetItemService().GetItem(int(slot.ItemId)).GetEquipmentTemplate()
	nextItemTemplate := equipTemplate.GetNextItemTemplate()
	if nextItemTemplate == nil {
		return false
	}
	return true
}

//进阶
func (bb *BodyBag) Upgrade(pos inventorytypes.BodyPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}
	if !bb.IfCanUpgrade(pos) {
		return false
	}
	//判断槽位是否可以升阶
	equipTemplate := item.GetItemService().GetItem(int(slot.ItemId)).GetEquipmentTemplate()
	nextItemTemplate := equipTemplate.GetNextItemTemplate()
	if nextItemTemplate == nil {
		return false
	}

	slot.ItemId = int32(nextItemTemplate.TemplateId())
	slot.SetModified()
	bb.changed(int(pos))
	gameevent.Emit(inventoryeventtypes.EventTypeEquipmentUpgrade, bb.p, pos)
	return true
}

//获取下一阶装备
func (bb *BodyBag) GetNextEquipment(pos inventorytypes.BodyPositionType) *gametemplate.ItemTemplate {
	it := bb.GetByPosition(pos)
	if it == nil {
		return nil
	}
	if it.ItemId == 0 {
		return nil
	}
	//物品不存在
	itemTemplate := item.GetItemService().GetItem(int(it.ItemId))
	if itemTemplate == nil {
		return nil
	}
	equipmentTemplate := itemTemplate.GetEquipmentTemplate()
	if equipmentTemplate == nil {
		return nil
	}
	nextItemTemplate := equipmentTemplate.GetNextItemTemplate()
	if nextItemTemplate == nil {
		return nil
	}
	return nextItemTemplate
}

//穿上
func (bb *BodyBag) PutOn(pos inventorytypes.BodyPositionType, itemId int32, bindType itemtypes.ItemBindType) (flag bool) {
	//TODO 判断物品是不是装备
	item := bb.GetByPosition(pos)
	//位置不存在
	if item == nil {
		return false
	}

	if item.IsEmpty() {
		now := global.GetGame().GetTimeService().Now()
		item.ItemId = itemId
		item.BindType = bindType
		item.UpdateTime = now
		item.SetModified()
		bb.changed(int(pos))
		flag = true
	}
	gameevent.Emit(inventoryeventtypes.EventTypeEquipmentPutOn, bb.p, pos)
	return
}

//脱下
func (bb *BodyBag) TakeOff(pos inventorytypes.BodyPositionType) (itemId int32) {
	item := bb.GetByPosition(pos)
	if item == nil {
		return
	}
	if item.IsEmpty() {
		return
	}
	itemId = item.ItemId
	defaultInitBind := itemtypes.ItemBindTypeUnBind
	now := global.GetGame().GetTimeService().Now()
	item.ItemId = 0
	item.BindType = defaultInitBind
	item.UpdateTime = now
	item.SetModified()
	bb.changed(int(pos))
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
	item.UpdateTime = now
	item.SetModified()
	bb.changed(int(pos))
	gameevent.Emit(inventoryeventtypes.EventTypeEquipmentEmbedGem, bb.p, nil)
	return true
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
	item.UpdateTime = now
	item.SetModified()
	bb.changed(int(pos))
	gameevent.Emit(inventoryeventtypes.EventTypeEquipmentTakeOffGem, bb.p, nil)
	return
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

// 是否有装备
func (bb *BodyBag) IsHadEquipment() bool {
	for _, slot := range bb.slotMap {
		if slot.ItemId != 0 {
			return true
		}
	}

	return false
}

//初始化
func (bb *BodyBag) init(slotList []*PlayerEquipmentSlotObject) {
	bb.changedBitset = bitset.New(64)
	bb.slotMap = make(map[inventorytypes.BodyPositionType]*PlayerEquipmentSlotObject)
	for _, slot := range slotList {
		bb.slotMap[slot.SlotId] = slot
	}
	now := global.GetGame().GetTimeService().Now()
	for slotId := inventorytypes.BodyPositionTypeWeapon; slotId <= inventorytypes.BodyPositionTypeRing; slotId++ {
		if bb.GetByPosition(slotId) != nil {
			continue
		}
		slot := createEquipmentSlotObject(bb.p, slotId, now)
		slot.SetModified()
		bb.slotMap[slot.SlotId] = slot
	}
}

//创建身体背包
func createBodyBag(p player.Player, slotList []*PlayerEquipmentSlotObject) *BodyBag {
	bb := &BodyBag{
		p: p,
	}

	bb.init(slotList)
	return bb
}

func createEquipmentSlotObject(p player.Player, slotId inventorytypes.BodyPositionType, now int64) *PlayerEquipmentSlotObject {
	itemObject := NewPlayerEquipmentSlotObject(p)
	itemObject.CreateTime = now
	itemObject.Id, _ = idutil.GetId()
	itemObject.GemInfo = make(map[int32]int32)
	itemObject.ItemId = 0
	itemObject.SlotId = slotId
	itemObject.CreateTime = now
	return itemObject
}

//玩家槽位数据
type PlayerEquipmentSlotObject struct {
	player     player.Player
	Id         int64
	SlotId     inventorytypes.BodyPositionType
	ItemId     int32
	Star       int32
	Level      int32
	GemInfo    map[int32]int32
	BindType   itemtypes.ItemBindType
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerEquipmentSlotObject(pl player.Player) *PlayerEquipmentSlotObject {
	pio := &PlayerEquipmentSlotObject{
		player: pl,
	}
	return pio
}

func convertPlayerEquipmentSlotObjectToEntity(pio *PlayerEquipmentSlotObject) (*gameentity.PlayerEquipmentSlotEntity, error) {
	gemInfoBytes, err := json.Marshal(pio.GemInfo)
	if err != nil {
		return nil, err
	}
	e := &gameentity.PlayerEquipmentSlotEntity{
		Id:         pio.Id,
		PlayerId:   pio.player.GetId(),
		ItemId:     pio.ItemId,
		SlotId:     int32(pio.SlotId),
		Star:       pio.Star,
		Level:      pio.Level,
		GemInfo:    string(gemInfoBytes),
		BindType:   int32(pio.BindType),
		UpdateTime: pio.UpdateTime,
		CreateTime: pio.CreateTime,
		DeleteTime: pio.DeleteTime,
	}
	return e, nil
}

func (pio *PlayerEquipmentSlotObject) GetPlayerId() int64 {
	return pio.player.GetId()
}

func (pio *PlayerEquipmentSlotObject) GetDBId() int64 {
	return pio.Id
}

func (pio *PlayerEquipmentSlotObject) GetBindType() itemtypes.ItemBindType {
	return pio.BindType
}

func (pio *PlayerEquipmentSlotObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerEquipmentSlotObjectToEntity(pio)
	return
}

func (pio *PlayerEquipmentSlotObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*gameentity.PlayerEquipmentSlotEntity)
	pio.Id = pse.Id
	pio.ItemId = pse.ItemId
	pio.SlotId = inventorytypes.BodyPositionType(pse.SlotId)
	pio.Level = pse.Level
	pio.Star = pse.Star
	gemInfo := make(map[int32]int32)
	//TODO 处理错误信息
	err = json.Unmarshal([]byte(pse.GemInfo), &gemInfo)
	if err != nil {
		return
	}
	pio.GemInfo = gemInfo
	pio.BindType = itemtypes.ItemBindType(pse.BindType)
	pio.UpdateTime = pse.UpdateTime
	pio.CreateTime = pse.CreateTime
	pio.DeleteTime = pse.DeleteTime
	return
}

func (pio *PlayerEquipmentSlotObject) SetModified() {
	e, err := pio.ToEntity()
	if err != nil {
		return
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("inventory:set modified never reach here"))
	}

	pio.player.AddChangedObject(obj)
	return
}

func (pio *PlayerEquipmentSlotObject) IsEmpty() bool {
	return pio.ItemId == 0
}

func (pio *PlayerEquipmentSlotObject) IsFull() bool {
	return pio.ItemId != 0
}
