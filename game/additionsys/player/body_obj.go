package player

import (
	"fgame/fgame/core/storage"
	additionsysentity "fgame/fgame/game/additionsys/entity"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"
	"fmt"

	"github.com/pkg/errors"
	"github.com/willf/bitset"
)

// 装备身体位置包裹
type BodyBag struct {
	p       player.Player
	typ     additionsystypes.AdditionSysType
	slotMap map[additionsystypes.SlotPositionType]*PlayerAdditionSysSlotObject
	//改变的位置
	changedBitset *bitset.BitSet
}

func (bb *BodyBag) GetAll() (slotList []*PlayerAdditionSysSlotObject) {
	for _, slot := range bb.slotMap {
		slotList = append(slotList, slot)
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

func (bb *BodyBag) GetChangedSlotAndReset() []*PlayerAdditionSysSlotObject {
	itemList := make([]*PlayerAdditionSysSlotObject, 0, 16)
	for i, valid := bb.changedBitset.NextSet(0); valid; i, valid = bb.changedBitset.NextSet(i + 1) {
		itemList = append(itemList, bb.slotMap[additionsystypes.SlotPositionType(i)])
	}
	bb.Reset()
	return itemList
}

//获取根据位置
func (bb *BodyBag) GetByPosition(pos additionsystypes.SlotPositionType) *PlayerAdditionSysSlotObject {
	eso, exist := bb.slotMap[pos]
	if !exist {
		return nil
	}
	return eso
}

//是否可以强化升级
func (bb *BodyBag) IfCanStrengthLevel(pos additionsystypes.SlotPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}

	return bb.GetNextStrengthenTemplate(pos) != nil
}

//获取下一个强化升级
func (bb *BodyBag) GetNextStrengthenTemplate(pos additionsystypes.SlotPositionType) *gametemplate.SystemStrengthenTemplate {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		//正常不可能
		return nil
	}
	var nextStrengthenTemplate *gametemplate.SystemStrengthenTemplate
	if slot.Level == 0 {
		nextStrengthenTemplate = additionsystemplate.GetAdditionSysTemplateService().GetBodyStrengthenByArg(slot.SysType, slot.SlotId, 1)
	} else {
		//判断槽位是否可以升级
		StrengthenTemplate := additionsystemplate.GetAdditionSysTemplateService().GetBodyStrengthenByArg(slot.SysType, slot.SlotId, slot.Level)
		nextStrengthenTemplate = StrengthenTemplate.GetNextTemplate()
	}
	return nextStrengthenTemplate
}

//强化升级
func (bb *BodyBag) StrengthLevel(pos additionsystypes.SlotPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}
	//不能升星
	nextEquipmentStrengthenTemplate := bb.GetNextStrengthenTemplate(pos)

	if nextEquipmentStrengthenTemplate == nil {
		return false
	}
	slot.Level = nextEquipmentStrengthenTemplate.Level
	slot.SetModified()
	bb.changed(int(pos))
	return true
}

//强化升级回退
func (bb *BodyBag) StrengthLevelBack(pos additionsystypes.SlotPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}

	nextEquipmentStrengthenTemplate := bb.GetNextStrengthenTemplate(pos)
	if nextEquipmentStrengthenTemplate == nil {
		return false
	}
	failEquipmentStrengthenTemplate := nextEquipmentStrengthenTemplate.GetFailTemplate()
	if failEquipmentStrengthenTemplate == nil {
		return false
	}
	slot.Level = failEquipmentStrengthenTemplate.Level
	slot.SetModified()
	bb.changed(int(pos))
	return true
}

//获取下一个神铸升级
func (bb *BodyBag) GetNextShenZhuTemplate(pos additionsystypes.SlotPositionType) *gametemplate.SystemShenZhuTemplate {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		//正常不可能
		return nil
	}
	var nextTemplate *gametemplate.SystemShenZhuTemplate
	if slot.ShenZhuLev == 0 {
		nextTemplate = additionsystemplate.GetAdditionSysTemplateService().GetShenZhuByArg(pos, 1)
	} else {
		//判断槽位是否可以升级
		curTemplate := additionsystemplate.GetAdditionSysTemplateService().GetShenZhuByArg(pos, slot.ShenZhuLev)
		nextTemplate = curTemplate.GetNextTemplate()
	}
	return nextTemplate
}

//神铸升级操作
func (bb *BodyBag) ShenZhuLevel(pos additionsystypes.SlotPositionType, pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return
	}
	if sucess {
		template := bb.GetNextShenZhuTemplate(pos)
		if template == nil {
			return
		}
		slot.ShenZhuLev = template.Level
		slot.ShenZhuNum = 0
		slot.ShenZhuPro = 0
	} else {
		slot.ShenZhuNum += addTimes
		slot.ShenZhuPro += pro
	}
	now := global.GetGame().GetTimeService().Now()
	slot.UpdateTime = now
	slot.SetModified()
	bb.changed(int(pos))
	return
}

//穿上
func (bb *BodyBag) PutOn(pos additionsystypes.SlotPositionType, itemId int32, bind itemtypes.ItemBindType) bool {
	bodySlot := bb.GetByPosition(pos)
	now := global.GetGame().GetTimeService().Now()
	//位置不存在
	if bodySlot == nil {
		//创建
		slot := createAdditionSysSlotObject(bb.p, bb.typ, pos, now)
		slot.ItemId = itemId
		slot.UpdateTime = now
		slot.Bind = bind
		slot.SetModified()
		bb.changed(int(pos))
		bb.slotMap[pos] = slot
		return true
	}

	if bodySlot.IsEmpty() {

		bodySlot.ItemId = itemId
		bodySlot.UpdateTime = now
		bodySlot.Bind = bind
		bodySlot.SetModified()
		bb.changed(int(pos))
		return true
	}

	return false
}

//脱下
func (bb *BodyBag) TakeOff(pos additionsystypes.SlotPositionType) (itemId int32) {
	bodySlot := bb.GetByPosition(pos)
	if bodySlot == nil {
		return
	}
	if bodySlot.IsEmpty() {
		return
	}
	itemId = bodySlot.ItemId
	defaultInitBind := itemtypes.ItemBindTypeUnBind
	now := global.GetGame().GetTimeService().Now()
	bodySlot.ItemId = 0
	bodySlot.Bind = defaultInitBind
	bodySlot.UpdateTime = now
	bodySlot.SetModified()
	bb.changed(int(pos))
	return
}

//进阶
func (bb *BodyBag) Upgrade(pos additionsystypes.SlotPositionType) bool {
	slot := bb.GetByPosition(pos)
	if slot == nil {
		return false
	}
	if slot.IsEmpty() {
		return false
	}
	//判断槽位是否可以升阶
	equipTemplate := item.GetItemService().GetItem(int(slot.ItemId)).GetSystemEquipTemplate()
	nextItemTemplate := equipTemplate.GetNextItemTemplate()
	if nextItemTemplate == nil {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	slot.ItemId = int32(nextItemTemplate.TemplateId())
	slot.UpdateTime = now
	slot.SetModified()
	bb.changed(int(pos))
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
func (bb *BodyBag) init(typ additionsystypes.AdditionSysType, argMap map[additionsystypes.SlotPositionType]*PlayerAdditionSysSlotObject) {
	bb.typ = typ
	bb.changedBitset = bitset.New(64)
	bb.slotMap = make(map[additionsystypes.SlotPositionType]*PlayerAdditionSysSlotObject)
	for _, slot := range argMap {
		tempSlotId := slot.SlotId
		bb.slotMap[tempSlotId] = slot
	}
	// now := global.GetGame().GetTimeService().Now()
	// for slotId := additionsystypes.MinPosition; slotId <= additionsystypes.MaxPosition; slotId++ {
	// 	if bb.GetByPosition(slotId) != nil {
	// 		continue
	// 	}
	// 	slot := createAdditionSysSlotObject(bb.p, typ, slotId, now)
	// 	slot.SetModified()
	// 	bb.slotMap[slotId] = slot
	// }
}

//创建身体背包
func createBodyBag(p player.Player, typ additionsystypes.AdditionSysType, argMap map[additionsystypes.SlotPositionType]*PlayerAdditionSysSlotObject) *BodyBag {
	bb := &BodyBag{
		p: p,
	}

	bb.init(typ, argMap)
	return bb
}

func createAdditionSysSlotObject(p player.Player, typ additionsystypes.AdditionSysType, slotId additionsystypes.SlotPositionType, now int64) *PlayerAdditionSysSlotObject {
	slotObject := NewPlayerAdditionSysSlotObject(p)
	slotObject.Id, _ = idutil.GetId()
	slotObject.ItemId = 0
	slotObject.SysType = typ
	slotObject.SlotId = slotId
	slotObject.CreateTime = now
	return slotObject
}

//玩家槽位数据
type PlayerAdditionSysSlotObject struct {
	Player     player.Player
	Id         int64
	PlayerId   int64
	SysType    additionsystypes.AdditionSysType
	SlotId     additionsystypes.SlotPositionType
	ItemId     int32
	Level      int32
	ShenZhuLev int32
	ShenZhuNum int32
	ShenZhuPro int32
	Bind       itemtypes.ItemBindType
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerAdditionSysSlotObject(pl player.Player) *PlayerAdditionSysSlotObject {
	pio := &PlayerAdditionSysSlotObject{
		Player:   pl,
		PlayerId: pl.GetId(),
	}
	return pio
}

func convertPlayerAdditionSysSlotObjectToEntity(pio *PlayerAdditionSysSlotObject) (*additionsysentity.PlayerAdditionSysSlotEntity, error) {
	e := &additionsysentity.PlayerAdditionSysSlotEntity{
		Id:         pio.Id,
		PlayerId:   pio.PlayerId,
		ItemId:     pio.ItemId,
		SysType:    int32(pio.SysType),
		SlotId:     int32(pio.SlotId),
		Level:      pio.Level,
		ShenZhuLev: pio.ShenZhuLev,
		ShenZhuNum: pio.ShenZhuNum,
		ShenZhuPro: pio.ShenZhuPro,
		BindType:   int32(pio.Bind),
		UpdateTime: pio.UpdateTime,
		CreateTime: pio.CreateTime,
		DeleteTime: pio.DeleteTime,
	}
	return e, nil
}

func (pio *PlayerAdditionSysSlotObject) GetPlayerId() int64 {
	return pio.PlayerId
}

func (pio *PlayerAdditionSysSlotObject) GetDBId() int64 {
	return pio.Id
}

func (pio *PlayerAdditionSysSlotObject) GetLevel() int32 {
	return pio.Level
}

func (pio *PlayerAdditionSysSlotObject) GetShenZhuLev() int32 {
	return pio.ShenZhuLev
}

func (pio *PlayerAdditionSysSlotObject) GetShenZhuNum() int32 {
	return pio.ShenZhuNum
}

func (pio *PlayerAdditionSysSlotObject) GetShenZhuPro() int32 {
	return pio.ShenZhuPro
}

func (pio *PlayerAdditionSysSlotObject) GetBindType() itemtypes.ItemBindType {
	return pio.Bind
}

func (pio *PlayerAdditionSysSlotObject) GetItemId() int32 {
	return pio.ItemId
}

func (pio *PlayerAdditionSysSlotObject) GetSlotId() additionsystypes.SlotPositionType {
	return pio.SlotId
}

func (pio *PlayerAdditionSysSlotObject) GetSysType() additionsystypes.AdditionSysType {
	return pio.SysType
}

func (pio *PlayerAdditionSysSlotObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerAdditionSysSlotObjectToEntity(pio)
	return
}

func (pio *PlayerAdditionSysSlotObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*additionsysentity.PlayerAdditionSysSlotEntity)
	pio.Id = pse.Id
	pio.PlayerId = pse.PlayerId
	pio.ItemId = pse.ItemId
	pio.SysType = additionsystypes.AdditionSysType(pse.SysType)
	pio.SlotId = additionsystypes.SlotPositionType(pse.SlotId)
	pio.Level = pse.Level
	pio.ShenZhuLev = pse.ShenZhuLev
	pio.ShenZhuNum = pse.ShenZhuNum
	pio.ShenZhuPro = pse.ShenZhuPro
	pio.Bind = itemtypes.ItemBindType(pse.BindType)
	pio.UpdateTime = pse.UpdateTime
	pio.CreateTime = pse.CreateTime
	pio.DeleteTime = pse.DeleteTime
	return
}

func (pio *PlayerAdditionSysSlotObject) SetModified() {
	e, err := pio.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "AdditionSysSlot"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("set modified never reach here"))
	}

	pio.Player.AddChangedObject(obj)
	return
}

func (pio *PlayerAdditionSysSlotObject) IsEmpty() bool {
	return pio.ItemId == 0
}

func (pio *PlayerAdditionSysSlotObject) IsFull() bool {
	return pio.ItemId != 0
}
