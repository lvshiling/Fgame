package item

import (
	"fgame/fgame/core/template"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/types"
	itemtypes "fgame/fgame/game/item/types"
	gametemplate "fgame/fgame/game/template"
	"sort"
	"sync"
)

type ItemService interface {
	GetEquipGemSlotTemplate(pos inventorytypes.BodyPositionType, order int32) *gametemplate.GemstoneSlotTemplate
	GetEquipStrengthenTemplate(typ inventorytypes.EquipmentStrengthenType, pos inventorytypes.BodyPositionType, level int32) *gametemplate.EquipStrengthenTemplate
	GetItem(id int) *gametemplate.ItemTemplate
	GetItemTemplate(types.ItemType, types.ItemSubType) *gametemplate.ItemTemplate
	GetItemClassMap(types.ItemType) map[types.ItemSubType]*gametemplate.ItemTemplate
	GetChargeItemTemplate(chargeType int32) *gametemplate.ItemTemplate
	GetBloodItemList() []*gametemplate.ItemTemplate
}

//快捷缓存
//物品配置的整合
type itemService struct {
	itemMap map[int]*gametemplate.ItemTemplate
	//物品分类
	itemClassMap map[types.ItemType]map[types.ItemSubType]*gametemplate.ItemTemplate
	//强化
	equipmentStrengthenTemplateMap map[inventorytypes.EquipmentStrengthenType]map[inventorytypes.BodyPositionType]map[int32]*gametemplate.EquipStrengthenTemplate
	//宝石装备槽
	equipmentGemSlotTemplateMap map[inventorytypes.BodyPositionType]map[int32]*gametemplate.GemstoneSlotTemplate
	//充值卡
	chargeItemTemplateMap map[int32]*gametemplate.ItemTemplate
	//血瓶
	bloodItemTemplateList []*gametemplate.ItemTemplate
}

func (is *itemService) init() (err error) {
	is.itemMap = make(map[int]*gametemplate.ItemTemplate)
	is.itemClassMap = make(map[types.ItemType]map[types.ItemSubType]*gametemplate.ItemTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.ItemTemplate)(nil))
	for _, templateObject := range templateMap {
		itemTemplate, _ := templateObject.(*gametemplate.ItemTemplate)
		is.itemMap[itemTemplate.TemplateId()] = itemTemplate

		typ := itemTemplate.GetItemType()
		subTyp := itemTemplate.GetItemSubType()

		if typ == types.ItemTypeDefault {
			continue
		}
		itemClassSubMap, exist := is.itemClassMap[typ]
		if !exist {
			itemClassSubMap = make(map[types.ItemSubType]*gametemplate.ItemTemplate)
			is.itemClassMap[typ] = itemClassSubMap
		}
		itemClassSubMap[subTyp] = itemTemplate
	}

	err = is.initEquipmentStrengthen()
	if err != nil {
		return
	}
	err = is.initEquipmentGemSlot()
	if err != nil {
		return
	}
	err = is.initChargeItems()
	if err != nil {
		return
	}

	//初始化血池
	err = is.initBloodItems()
	if err != nil {
		return
	}

	return nil
}

//初始化强化列表
func (is *itemService) initEquipmentStrengthen() (err error) {
	is.equipmentStrengthenTemplateMap = make(map[inventorytypes.EquipmentStrengthenType]map[inventorytypes.BodyPositionType]map[int32]*gametemplate.EquipStrengthenTemplate)

	templateEquipmentStrengthenMap := template.GetTemplateService().GetAll((*gametemplate.EquipStrengthenTemplate)(nil))
	for _, templateObject := range templateEquipmentStrengthenMap {
		equipStrengthenTemplate, _ := templateObject.(*gametemplate.EquipStrengthenTemplate)
		equipmentStrengthenPosMap, ok := is.equipmentStrengthenTemplateMap[equipStrengthenTemplate.GetEquipStrengthType()]
		if !ok {
			equipmentStrengthenPosMap = make(map[inventorytypes.BodyPositionType]map[int32]*gametemplate.EquipStrengthenTemplate)
			is.equipmentStrengthenTemplateMap[equipStrengthenTemplate.GetEquipStrengthType()] = equipmentStrengthenPosMap
		}
		equipmentStrengthenMap, ok := equipmentStrengthenPosMap[equipStrengthenTemplate.GetPosition()]
		if !ok {
			equipmentStrengthenMap = make(map[int32]*gametemplate.EquipStrengthenTemplate)
			equipmentStrengthenPosMap[equipStrengthenTemplate.GetPosition()] = equipmentStrengthenMap
		}
		equipmentStrengthenMap[equipStrengthenTemplate.Level] = equipStrengthenTemplate
	}
	return
}

//初始化宝石装备槽
func (is *itemService) initEquipmentGemSlot() (err error) {
	is.equipmentGemSlotTemplateMap = make(map[inventorytypes.BodyPositionType]map[int32]*gametemplate.GemstoneSlotTemplate)

	templateEquipmentGemSlotTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GemstoneSlotTemplate)(nil))
	for _, templateObject := range templateEquipmentGemSlotTemplateMap {
		gemstoneSlotTemplate, _ := templateObject.(*gametemplate.GemstoneSlotTemplate)
		equipmentGemSlotPosMap, ok := is.equipmentGemSlotTemplateMap[gemstoneSlotTemplate.GetBodyPosition()]
		if !ok {
			equipmentGemSlotPosMap = make(map[int32]*gametemplate.GemstoneSlotTemplate)
			is.equipmentGemSlotTemplateMap[gemstoneSlotTemplate.GetBodyPosition()] = equipmentGemSlotPosMap
		}

		equipmentGemSlotPosMap[gemstoneSlotTemplate.Order] = gemstoneSlotTemplate
	}
	return
}

//初始化充值
func (is *itemService) initChargeItems() (err error) {
	is.chargeItemTemplateMap = make(map[int32]*gametemplate.ItemTemplate)

	tempItemMap := template.GetTemplateService().GetAll((*gametemplate.ItemTemplate)(nil))
	for _, templateObject := range tempItemMap {
		tempItemTemplate, _ := templateObject.(*gametemplate.ItemTemplate)
		if tempItemTemplate.GetItemType() != itemtypes.ItemTypeFuChi {
			continue
		}
		is.chargeItemTemplateMap[tempItemTemplate.TypeFlag1] = tempItemTemplate
	}
	return
}

type bloodItemList []*gametemplate.ItemTemplate

func (adl bloodItemList) Len() int {
	return len(adl)
}

func (adl bloodItemList) Less(i, j int) bool {
	return adl[i].TypeFlag1 < adl[j].TypeFlag1
}

func (adl bloodItemList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

//初始化血池
func (is *itemService) initBloodItems() (err error) {
	is.bloodItemTemplateList = make([]*gametemplate.ItemTemplate, 0, 8)

	bloodMap, ok := is.itemClassMap[itemtypes.ItemTypeLifeOrigin]
	if !ok {
		return
	}
	for _, itemTemplate := range bloodMap {
		is.bloodItemTemplateList = append(is.bloodItemTemplateList, itemTemplate)
	}
	sort.Sort(bloodItemList(is.bloodItemTemplateList))
	return
}

//获取装备槽宝石模板
func (is *itemService) GetEquipGemSlotTemplate(pos inventorytypes.BodyPositionType, order int32) *gametemplate.GemstoneSlotTemplate {
	posMap, ok := is.equipmentGemSlotTemplateMap[pos]
	if !ok {
		return nil
	}
	to, ok := posMap[order]
	if !ok {
		return nil
	}
	return to
}

func (is *itemService) GetEquipStrengthenTemplate(typ inventorytypes.EquipmentStrengthenType, pos inventorytypes.BodyPositionType, level int32) *gametemplate.EquipStrengthenTemplate {
	posMap, ok := is.equipmentStrengthenTemplateMap[typ]
	if !ok {
		return nil
	}
	toMap, ok := posMap[pos]
	if !ok {
		return nil
	}
	to, ok := toMap[level]
	if !ok {
		return nil
	}
	return to
}
func (is *itemService) GetItem(id int) *gametemplate.ItemTemplate {
	to, ok := is.itemMap[id]
	if !ok {
		return nil
	}
	return to
}

func (is *itemService) GetItemClassMap(itemType types.ItemType) map[types.ItemSubType]*gametemplate.ItemTemplate {
	itemClassSubMap, exist := is.itemClassMap[itemType]
	if !exist {
		return nil
	}
	return itemClassSubMap
}

func (is *itemService) GetItemTemplate(itemType types.ItemType, itemSubType types.ItemSubType) *gametemplate.ItemTemplate {
	itemClassSubMap := is.GetItemClassMap(itemType)
	if itemClassSubMap == nil {
		return nil
	}
	itemTemplate, exist := itemClassSubMap[itemSubType]
	if !exist {
		return nil
	}
	return itemTemplate
}
func (is *itemService) GetChargeItemTemplate(chargeType int32) *gametemplate.ItemTemplate {
	tem, ok := is.chargeItemTemplateMap[chargeType]
	if !ok {
		return nil
	}
	return tem
}

func (is *itemService) GetBloodItemList() []*gametemplate.ItemTemplate {
	return is.bloodItemTemplateList
}

var (
	once sync.Once
	cs   *itemService
)

func Init() (err error) {
	once.Do(func() {
		cs = &itemService{}
		err = cs.init()
	})
	return err
}

func GetItemService() ItemService {
	return cs
}
