package player

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/storage"
	babytypes "fgame/fgame/game/baby/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	"fgame/fgame/game/inventory/dao"
	inventoryentity "fgame/fgame/game/inventory/entity"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	"fgame/fgame/game/inventory/inventory"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	ringtemplate "fgame/fgame/game/ring/template"
	ringtypes "fgame/fgame/game/ring/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//仓库对象
type PlayerInventoryObject struct {
	player        player.Player
	Id            int64
	SlotNum       int32
	DepotNum      int32
	MiBaoDepotNum int32
	UpdateTime    int64
	CreateTime    int64
	DeleteTime    int64
}

func NewPlayerInventoryObject(pl player.Player) *PlayerInventoryObject {
	pio := &PlayerInventoryObject{
		player: pl,
	}
	return pio
}

func convertPlayerInventoryObjectToEntity(pio *PlayerInventoryObject) *inventoryentity.PlayerInventoryEntity {
	e := &inventoryentity.PlayerInventoryEntity{
		Id:            pio.Id,
		PlayerId:      pio.player.GetId(),
		SlotNum:       pio.SlotNum,
		DepotNum:      pio.DepotNum,
		MiBaoDepotNum: pio.MiBaoDepotNum,
		UpdateTime:    pio.UpdateTime,
		CreateTime:    pio.CreateTime,
		DeleteTime:    pio.DeleteTime,
	}
	return e
}

func (pio *PlayerInventoryObject) GetPlayerId() int64 {
	return pio.player.GetId()
}

func (pio *PlayerInventoryObject) GetDBId() int64 {
	return pio.Id
}

func (pio *PlayerInventoryObject) ToEntity() (e storage.Entity, err error) {
	e = convertPlayerInventoryObjectToEntity(pio)
	return
}

func (pio *PlayerInventoryObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*inventoryentity.PlayerInventoryEntity)
	pio.Id = pse.Id
	pio.SlotNum = pse.SlotNum
	pio.DepotNum = pse.DepotNum
	pio.MiBaoDepotNum = pse.MiBaoDepotNum
	if pio.MiBaoDepotNum == 0 {
		pio.MiBaoDepotNum = constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMiBaoDepotSlotMax)
	}
	pio.UpdateTime = pse.UpdateTime
	pio.CreateTime = pse.CreateTime
	pio.DeleteTime = pse.DeleteTime
	return
}

func (pio *PlayerInventoryObject) SetModified() {
	e, err := pio.ToEntity()
	if err != nil {
		return
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pio.player.AddChangedObject(obj)
	return
}

//玩家背包管理器
type PlayerInventoryDataManager struct {
	p player.Player
	//背包对象
	playerInventoryObject *PlayerInventoryObject
	//主背包
	shoulderBagMap map[inventorytypes.BagType]*ShoulderBag
	//仓库
	depotBag *ShoulderBag
	//装备背包
	equipmentBag *BodyBag
	//秘宝仓库
	mibaoDepotBag *ShoulderBag
	//材料仓库
	materialDepotBag *ShoulderBag
	//次数限制性物品使用记录
	itemUseRecordMap map[int32]*PlayerItemUseObject
	//
	hbRunner heartbeat.HeartbeatTaskRunner
}

//获取玩家信息
func (pidm *PlayerInventoryDataManager) Player() player.Player {
	return pidm.p
}

//获取装备背包
func (pidm *PlayerInventoryDataManager) GetEquipmentBag() *BodyBag {
	return pidm.equipmentBag
}

//获取仓库
func (pidm *PlayerInventoryDataManager) getDepotBag() *ShoulderBag {
	return pidm.depotBag
}

//获取秘宝仓库
func (pidm *PlayerInventoryDataManager) getMiBaoDepotBag() *ShoulderBag {
	return pidm.mibaoDepotBag
}

// 获取材料仓库
func (pidm *PlayerInventoryDataManager) getMaterialDepotBag() *ShoulderBag {
	return pidm.materialDepotBag
}

//加载
func (pidm *PlayerInventoryDataManager) Load() (err error) {
	//加载背包
	err = pidm.loadInventory()
	if err != nil {
		return
	}
	//加载装备数据
	err = pidm.loadEquipmentSlot()
	if err != nil {
		return
	}
	//加载限制性物品使用记录
	err = pidm.loadUseRecord()
	if err != nil {
		return
	}
	return nil
}

//加载存储背包
func (pidm *PlayerInventoryDataManager) loadInventory() (err error) {
	pidm.shoulderBagMap = make(map[inventorytypes.BagType]*ShoulderBag)

	//加载背包
	inventoryEntity, err := dao.GetInventoryDao().GetInventoryEntity(pidm.p.GetId())
	if err != nil {
		return
	}
	if inventoryEntity == nil {
		err = pidm.initPlayerInventoryObject()
		if err != nil {
			return
		}
	} else {
		pidm.playerInventoryObject = NewPlayerInventoryObject(pidm.p)
		pidm.playerInventoryObject.FromEntity(inventoryEntity)
	}
	//加载物品
	items, err := dao.GetInventoryDao().GetItemList(pidm.p.GetId())
	if err != nil {
		return
	}

	itemListMap := make(map[inventorytypes.BagType][]*PlayerItemObject)
	//加载物品
	for _, item := range items {
		pio := NewPlayerItemObject(pidm.p)
		pio.FromEntity(item)
		tempItemList, _ := itemListMap[pio.BagType]
		tempItemList = append(tempItemList, pio)
		itemListMap[pio.BagType] = tempItemList
	}

	//主背包
	shoulderList, _ := itemListMap[inventorytypes.BagTypePrim]
	// 修正升星强化等级
	pidm.fixUpstarLevel(shoulderList)
	pidm.shoulderBagMap[inventorytypes.BagTypePrim] = createShoulderBag(pidm.p, shoulderList, pidm.playerInventoryObject.SlotNum)
	//宝石背包
	gemNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeGemBagTotalNum)
	gemShoulderList, _ := itemListMap[inventorytypes.BagTypeGem]
	pidm.shoulderBagMap[inventorytypes.BagTypeGem] = createShoulderBag(pidm.p, gemShoulderList, gemNum)
	//鲲背包
	kunNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeOneArenaKunSolt)
	kunShoulderList, _ := itemListMap[inventorytypes.BagTypeKun]
	pidm.shoulderBagMap[inventorytypes.BagTypeKun] = createShoulderBag(pidm.p, kunShoulderList, kunNum)
	//命格背包
	mingGeSlotNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMingGeSlotNum)
	mingGeList, _ := itemListMap[inventorytypes.BagTypeMingGe]
	pidm.shoulderBagMap[inventorytypes.BagTypeMingGe] = createShoulderBag(pidm.p, mingGeList, mingGeSlotNum)
	//神器背包
	shenQiNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeShenQiSlotNum)
	shenQiShoulderList, _ := itemListMap[inventorytypes.BagTypeShenQi]
	pidm.shoulderBagMap[inventorytypes.BagTypeShenQi] = createShoulderBag(pidm.p, shenQiShoulderList, shenQiNum)
	//器灵背包
	qiLingNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeQiLingSlotNum)
	qiLingShoulderList, _ := itemListMap[inventorytypes.BagTypeQiLing]
	pidm.shoulderBagMap[inventorytypes.BagTypeQiLing] = createShoulderBag(pidm.p, qiLingShoulderList, qiLingNum)
	//英灵背包
	yingLingNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeYingLingPuSlotNum)
	yingLingShoulderList, _ := itemListMap[inventorytypes.BagTypeYingLingPu]
	pidm.shoulderBagMap[inventorytypes.BagTypeYingLingPu] = createShoulderBag(pidm.p, yingLingShoulderList, yingLingNum)

	//屠龙装备背包
	tulongEquipNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuLongEquipSlotNum)
	tulongShoulderList, _ := itemListMap[inventorytypes.BagTypeTuLongEquip]
	pidm.shoulderBagMap[inventorytypes.BagTypeTuLongEquip] = createShoulderBag(pidm.p, tulongShoulderList, tulongEquipNum)

	//仓库
	depotItems, _ := dao.GetInventoryDao().GetDepotItemList(pidm.p.GetId())
	if err != nil {
		return
	}
	var tempItemList []*PlayerItemObject
	for _, item := range depotItems {
		pio := NewPlayerItemObject(pidm.p)
		pio.FromEntity(item)
		tempItemList = append(tempItemList, pio)
	}
	// 修正升星强化等级
	pidm.fixUpstarLevel(tempItemList)
	pidm.depotBag = createShoulderBag(pidm.p, tempItemList, pidm.playerInventoryObject.DepotNum)

	//秘宝仓库
	mibaoDepotItems, _ := dao.GetInventoryDao().GetMiBaoDepotItemList(pidm.p.GetId())
	if err != nil {
		return
	}
	var mibaoTempItemList []*PlayerItemObject
	for _, item := range mibaoDepotItems {
		pio := NewPlayerItemObject(pidm.p)
		pio.FromEntity(item)
		mibaoTempItemList = append(mibaoTempItemList, pio)
	}
	pidm.mibaoDepotBag = createShoulderBag(pidm.p, mibaoTempItemList, pidm.playerInventoryObject.MiBaoDepotNum)

	//材料仓库
	materialDepotItems, _ := dao.GetInventoryDao().GetMaterialDepotItemList(pidm.p.GetId())
	if err != nil {
		return
	}
	var materialTempItemList []*PlayerItemObject
	for _, item := range materialDepotItems {
		pio := NewPlayerItemObject(pidm.p)
		pio.FromEntity(item)
		materialTempItemList = append(materialTempItemList, pio)
	}
	materialDepotNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMaterialDepotSlotMax)
	pidm.materialDepotBag = createShoulderBag(pidm.p, materialTempItemList, materialDepotNum)

	return
}

// 修正升星强化等级
func (pidm *PlayerInventoryDataManager) fixUpstarLevel(itemObjList []*PlayerItemObject) {
	for _, itemObj := range itemObjList {
		if itemObj.IsEmpty() {
			continue
		}

		goldequipData, ok := itemObj.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
		if !ok {
			continue
		}
		itemTemp := item.GetItemService().GetItem(int(itemObj.ItemId))
		if itemTemp.GetGoldEquipTemplate() == nil {
			log.Info("itemid:", itemObj.ItemId)
			continue
		}
		maxLeve := itemTemp.GetGoldEquipTemplate().GetMaxUpstarLevel()
		goldequipData.FixUpstarLevel(maxLeve)
		itemObj.SetModified()
	}
}

//加载身上装备
func (pidm *PlayerInventoryDataManager) loadEquipmentSlot() (err error) {
	//加载槽位
	equipmentSlotList, err := dao.GetInventoryDao().GetEquipmentSlotList(pidm.p.GetId())
	if err != nil {
		return
	}
	slotList := make([]*PlayerEquipmentSlotObject, 0, len(equipmentSlotList))
	for _, slot := range equipmentSlotList {
		pio := NewPlayerEquipmentSlotObject(pidm.p)
		pio.FromEntity(slot)
		slotList = append(slotList, pio)
	}
	pidm.equipmentBag = createBodyBag(pidm.p, slotList)
	return
}

//加载物品使用记录
func (pidm *PlayerInventoryDataManager) loadUseRecord() (err error) {
	itemUseList, err := dao.GetInventoryDao().GetItemUseList(pidm.p.GetId())
	if err != nil {
		return
	}
	itemUseMap := make(map[int32]*PlayerItemUseObject)
	for _, itemUse := range itemUseList {
		pio := NewPlayerItemUseObject(pidm.p)
		pio.FromEntity(itemUse)
		itemUseMap[pio.ItemId] = pio
	}
	pidm.itemUseRecordMap = itemUseMap
	return
}

//第一次初始化
func (pidm *PlayerInventoryDataManager) initPlayerInventoryObject() (err error) {
	pio := NewPlayerInventoryObject(pidm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pio.Id = id
	pio.SlotNum = constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBagDefaultOpenNum)
	pio.DepotNum = constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDepotDefaultOpenNum)
	pio.MiBaoDepotNum = constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMiBaoDepotSlotMax)
	pio.CreateTime = now
	pidm.playerInventoryObject = pio
	pio.SetModified()

	return
}

func (pidm *PlayerInventoryDataManager) AfterLoad() (err error) {
	//刷新
	if err = pidm.refresh(); err != nil {
		return
	}
	//重置使用次数
	pidm.hbRunner.AddTask(CreateResetItemUseTask(pidm.p))
	return
}

//刷新数据
func (pidm *PlayerInventoryDataManager) refresh() (err error) {
	//第一次初始化
	err = pidm.checkExpired()
	return
}

//检查是否过期
func (pidm *PlayerInventoryDataManager) checkExpired() (err error) {
	// shoulderBag := pidm.getBag(inventorytypes.BagTypePrim)
	// if shoulderBag == nil {
	// 	panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.BagTypePrim.String()))
	// }

	// isSync := false
	// for _, itemObj := range shoulderBag.GetItemList() {
	// 	itemTemp := item.GetItemService().GetItem(int(itemObj.ItemId))
	// 	if itemTemp == nil {
	// 		continue
	// 	}

	// 	if itemTemp.GetLimitTimeType() != itemtypes.ItemLimitTimeTypeExpired {
	// 		continue
	// 	}

	// 	expireTime := itemTemp.GetExpireTime() + itemObj.ItemGetTime
	// 	now := global.GetGame().GetTimeService().Now()
	// 	if now < expireTime {
	// 		continue
	// 	}

	// 	flag := shoulderBag.RemoveItem(itemObj.ItemId, itemObj.Num)
	// 	if !flag {
	// 		return fmt.Errorf("移除过期物品失败，index:%d", itemObj.Index)
	// 	}

	// 	isSync = true
	// }

	// if isSync {
	// 	gameevent.Emit(inventoryeventtypes.EventTypeItemExpire, pidm.p, nil)
	// }
	return
}

//剩余可以添加的格子
func (pidm *PlayerInventoryDataManager) IfCanAddSlots(buyNum int32) bool {
	totalNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBagTotalNum)
	return totalNum-pidm.playerInventoryObject.SlotNum >= buyNum
}

//获取剩余可以购买的槽位
func (pidm *PlayerInventoryDataManager) NumOfRemainBuySlots() int32 {
	totalNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBagTotalNum)
	return totalNum - pidm.playerInventoryObject.SlotNum
}

//仓库剩余可以添加的格子
func (pidm *PlayerInventoryDataManager) IfCanAddDepotSlots(num int32) bool {
	totalNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDepotMaxNum)
	return totalNum-pidm.playerInventoryObject.DepotNum >= num
}

//添加格子数
func (pidm *PlayerInventoryDataManager) AddSlots(buyNum int32) bool {
	if !pidm.IfCanAddSlots(buyNum) {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	pidm.playerInventoryObject.SlotNum += buyNum
	pidm.playerInventoryObject.UpdateTime = now
	pidm.playerInventoryObject.SetModified()
	//重置背包容量
	pidm.getPrimBag().ResetCapacity(pidm.playerInventoryObject.SlotNum)
	return true
}

//添加仓库格子数
func (pidm *PlayerInventoryDataManager) AddDepotSlots(num int32) bool {
	if !pidm.IfCanAddDepotSlots(num) {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	pidm.playerInventoryObject.DepotNum += num
	pidm.playerInventoryObject.UpdateTime = now
	pidm.playerInventoryObject.SetModified()
	//重置背包容量
	pidm.getDepotBag().ResetCapacity(pidm.playerInventoryObject.DepotNum)
	return true
}

func (pidm *PlayerInventoryDataManager) GetSlots() int32 {
	return pidm.playerInventoryObject.SlotNum
}

func (pidm *PlayerInventoryDataManager) GetDepotSlots() int32 {
	return pidm.playerInventoryObject.DepotNum
}

// -------------------------v1.0----------------------------

//添加物品
func (pidm *PlayerInventoryDataManager) AddItem(itemId int32, num int32, reason commonlog.InventoryLogReason, reasonText string) bool {
	itemData := droptemplate.CreateItemData(itemId, num, 0, itemtypes.ItemBindTypeUnBind)
	return pidm.AddItemLevel(itemData, reason, reasonText)
}

//批量添加物品
func (pidm *PlayerInventoryDataManager) BatchAdd(items map[int32]int32, reason commonlog.InventoryLogReason, reasonText string) bool {
	newItems := mergeItems(items)
	if !pidm.HasEnoughSlots(newItems) {
		return false
	}
	for itemId, num := range newItems {
		pidm.addItem(itemId, num, reason, reasonText)
	}

	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChanged, pidm.p, nil)
	return true
}

//是否有足够的位置存放物品
func (pidm *PlayerInventoryDataManager) HasEnoughSlot(itemId int32, num int32) bool {
	if itemId == 0 || num <= 0 {
		panic(fmt.Errorf("inventory:id %d,has enough slot itemId %d,num %d", pidm.p.GetId(), itemId, num))
	}
	remainNum := pidm.RemainSlotForItemLevel(itemId, 0, itemtypes.ItemBindTypeUnBind)
	if remainNum >= num {
		return true
	}
	return false
}

//检查背包
func (pidm *PlayerInventoryDataManager) HasEnoughSlots(items map[int32]int32) bool {
	itemDataList := droptemplate.ConvertToItemDataList(items, itemtypes.ItemBindTypeUnBind)
	return pidm.HasEnoughSlotsOfItemLevel(itemDataList)
}

//剩余位置
func (pidm *PlayerInventoryDataManager) RemainSlotForItem(itemId int32) (maxNum int32) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return
	}
	shouldBag := pidm.getBag(itemTemplate.GetBagType())
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
	}
	return shouldBag.RemainSlotForItem(itemId)
}

//仓库是否有足够的位置存放物品
func (pidm *PlayerInventoryDataManager) HasEnoughDepotSlot(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType) bool {
	if itemId == 0 || num <= 0 {
		panic(fmt.Errorf("inventory:id %d,has enough depot slot itemId %d,num %d", pidm.p.GetId(), itemId, num))
	}
	remainNum := pidm.RemainDepotSlotForItem(itemId, level, bindType)
	if remainNum >= num {
		return true
	}
	return false
}

//仓库剩余位置
func (pidm *PlayerInventoryDataManager) RemainDepotSlotForItem(itemId int32, level int32, bindType itemtypes.ItemBindType) (maxNum int32) {
	shouldBag := pidm.getDepotBag()
	if shouldBag == nil {
		panic("inventory:仓库应该存在")
	}
	return shouldBag.RemainSlotForItemLevel(itemId, level, bindType)
}

//保存到仓库
func (pidm *PlayerInventoryDataManager) AddItemInDepot(itemId int32, num int32, level int32, bind itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) bool {
	//物品不存在
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}
	if itemTemplate.Storage == 0 {
		return false
	}

	shouldBag := pidm.getDepotBag()
	if shouldBag == nil {
		panic("inventory：仓库应该存在")
	}

	flag := shouldBag.AddLevelItemWithPropertyData(itemId, num, level, propertyData, inventorytypes.IsDepotTypeDepot, bind)
	if !flag {
		return false
	}

	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChanged, pidm.p, nil)
	return true
}

// -------------------------v1.1----------------------------

//添加含等级物品
func (pidm *PlayerInventoryDataManager) AddItemLevel(itemData *droptemplate.DropItemData, reason commonlog.InventoryLogReason, reasonText string) bool {
	itemId := itemData.GetItemId()
	num := itemData.GetNum()
	level := itemData.GetLevel()
	bind := itemData.GetBindType()
	expireType := itemData.GetExpireType()
	expireTime := itemData.GetExpireTime()
	itemGetTime := itemData.GetItemGetTime()
	//物品不存在
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}
	shouldBag := pidm.getBag(itemTemplate.GetBagType())
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
	}

	beforeNum := shouldBag.NumOfItemsWithProperty(itemId)
	baseProperty := inventorytypes.CreateItemPropertyDataBase(expireType, expireTime, itemGetTime)
	propertyData := inventory.CreatePropertyDataInterface(itemTemplate.GetItemType(), baseProperty)
	if itemTemplate.IsGoldEquip() {
		upstar := itemData.GetUpstar()
		attrList := itemData.GetAttrList()
		for i := 0; i < int(num); i++ {
			goldequipPropertyData := propertyData.(*goldequiptypes.GoldEquipPropertyData)
			goldequipPropertyData.UpstarLevel = upstar
			goldequipPropertyData.OpenLightLevel = itemData.GetOpenLightLevel()
			goldequipPropertyData.OpenTimes = itemData.GetOpenTimes()

			if itemData.IsRandomAttr {
				goldequipPropertyData.AttrList = attrList
				goldequipPropertyData.IsHadCountAttr = true
			}
			num = 1
			// flag := shouldBag.AddLevelItemWithPropertyData(itemId, 1, level, goldequipPropertyData, inventorytypes.IsDepotTypePrim, bind)
			// if !flag {
			// 	return false
			// }
		}
	} else if itemTemplate.IsBaoBaoCard() {
		babyPropertyData := propertyData.(*babytypes.BabyPropertyData)
		babyPropertyData.Quality = itemData.GetQuality()
		babyPropertyData.Sex = itemData.GetSex()
		babyPropertyData.TalentList = itemData.GetTalentList()
		babyPropertyData.Danbei = itemData.GetDanbei()
	} else if itemTemplate.IsTeRing() {
		ringPropertyData := propertyData.(*ringtypes.RingPropertyData)
		ringPropertyData.Advance = ringtemplate.GetRingTemplateService().GetRingMinAdvance(itemId)
	}

	flag := shouldBag.AddLevelItemWithPropertyData(itemId, num, level, propertyData, inventorytypes.IsDepotTypePrim, bind)
	if !flag {
		return false
	}

	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChanged, pidm.p, nil)
	data := inventoryeventtypes.CreatePlayerInventoryChangedLogEventData(itemId, beforeNum, num, reason, reasonText)
	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChangedLog, pidm.p, data)
	return true
}

//批量添加物品
func (pidm *PlayerInventoryDataManager) BatchAddOfItemLevel(itemList []*droptemplate.DropItemData, reason commonlog.InventoryLogReason, reasonText string) bool {
	newItems := mergeItemLevel(itemList)
	if !pidm.HasEnoughSlotsOfItemLevel(newItems) {
		return false
	}
	for _, itemData := range newItems {
		flag := pidm.AddItemLevel(itemData, reason, reasonText)
		if !flag {
			return false
		}
	}
	return true
}

//是否有足够的位置存放物品
func (pidm *PlayerInventoryDataManager) HasEnoughSlotItemLevel(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType) bool {
	expireType := inventorytypes.NewItemLimitTimeTypeNone
	expireTime := int64(0)
	itemGetTime := global.GetGame().GetTimeService().Now()
	return pidm.HasEnoughSlotItemLevelWithProperty(itemId, num, level, bindType, expireType, itemGetTime, expireTime)
}

//是否有足够的位置存放物品-多个
func (pidm *PlayerInventoryDataManager) HasEnoughSlotsOfItemLevel(itemList []*droptemplate.DropItemData) bool {

	needSlotMap := make(map[inventorytypes.BagType]int32)
	for _, itemData := range itemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		level := itemData.GetLevel()
		bind := itemData.GetBindType()
		expireType := itemData.GetExpireType()
		// expireTime := itemData.GetExpireTime()

		itemTemplate := item.GetItemService().GetItem(int(itemId))
		if itemTemplate == nil {
			return false
		}

		shoulderBag := pidm.getBag(itemTemplate.GetBagType())
		if shoulderBag == nil {
			panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
		}

		slotNum := shoulderBag.CountNeedSlotOfItemLevelWithProperty(itemId, num, level, bind, expireType, itemData.GetItemGetTime(), itemData.GetExpireTimestamp())
		needSlot := needSlotMap[itemTemplate.GetBagType()]
		needSlot += slotNum
		needSlotMap[itemTemplate.GetBagType()] = needSlot
	}
	for bagType, needSlot := range needSlotMap {
		shoulderBag := pidm.getBag(bagType)
		if needSlot > shoulderBag.GetEmptySlots() {
			return false
		}
	}
	return true
}

//剩余位置
func (pidm *PlayerInventoryDataManager) RemainSlotForItemLevel(itemId int32, level int32, bindType itemtypes.ItemBindType) (maxNum int32) {
	expireType := inventorytypes.NewItemLimitTimeTypeNone
	expireTime := int64(0)
	itemGetTime := global.GetGame().GetTimeService().Now()
	return pidm.RemainSlotForItemLevelWithProperty(itemId, level, bindType, expireType, itemGetTime, expireTime)
}

// -------------------------v1.2----------------------------

//添加物品(含属性接口)
func (pidm *PlayerInventoryDataManager) AddItemLevelWithPropertyData(itemId int32, num int32, level int32, bind itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData, reason commonlog.InventoryLogReason, reasonText string) bool {
	//物品不存在
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}
	shouldBag := pidm.getBag(itemTemplate.GetBagType())
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
	}
	beforeNum := shouldBag.NumOfItemsWithProperty(itemId)
	flag := shouldBag.AddLevelItemWithPropertyData(itemId, num, level, propertyData, inventorytypes.IsDepotTypePrim, bind)
	if !flag {
		return false
	}

	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChanged, pidm.p, nil)
	data := inventoryeventtypes.CreatePlayerInventoryChangedLogEventData(itemId, beforeNum, num, reason, reasonText)
	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChangedLog, pidm.p, data)
	return true
}

//是否有足够的位置存放物品(含时效性)
func (pidm *PlayerInventoryDataManager) HasEnoughSlotItemLevelWithProperty(itemId, num, level int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, itemGetTime int64, expireTime int64) bool {
	if itemId == 0 || num <= 0 {
		panic(fmt.Errorf("inventory:id %d,has enough slot itemId %d,num %d", pidm.p.GetId(), itemId, num))
	}
	remainNum := pidm.RemainSlotForItemLevelWithProperty(itemId, level, bindType, expireType, itemGetTime, expireTime)
	if remainNum >= num {
		return true
	}
	return false
}

//剩余位置（含时效性）
func (pidm *PlayerInventoryDataManager) RemainSlotForItemLevelWithProperty(itemId int32, level int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, itemGetTime int64, expireTime int64) (maxNum int32) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return
	}
	shouldBag := pidm.getBag(itemTemplate.GetBagType())
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
	}
	return shouldBag.RemainSlotForItemLevelWithProperty(itemId, level, bindType, expireType, itemGetTime, expireTime)
}

//仓库是否有足够的位置存放物品
func (pidm *PlayerInventoryDataManager) HasEnoughDepotSlotWithProperty(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType, expireType inventorytypes.NewItemLimitTimeType, itemGetTime int64, expireTime int64) bool {
	if itemId == 0 || num <= 0 {
		panic(fmt.Errorf("inventory:id %d,has enough depot slot itemId %d,num %d", pidm.p.GetId(), itemId, num))
	}
	shouldBag := pidm.getDepotBag()
	if shouldBag == nil {
		panic("inventory:仓库应该存在")
	}

	remainNum := shouldBag.RemainSlotForItemLevelWithProperty(itemId, level, bindType, expireType, itemGetTime, expireTime)
	if remainNum >= num {
		return true
	}
	return false
}

// -------------------------分割线----------------------------

//批量验证数量
func (pidm *PlayerInventoryDataManager) HasEnoughItems(items map[int32]int32) bool {
	if len(items) == 0 {
		panic(fmt.Errorf("inventory:物品不应该是空"))
	}
	for itemId, num := range items {
		if num <= 0 {
			panic(fmt.Errorf("inventory:num should more than 0"))
		}
		if !pidm.HasEnoughItem(itemId, num) {
			return false
		}
	}
	return true
}

//是否足够物品
func (pidm *PlayerInventoryDataManager) HasEnoughItem(itemId int32, num int32) bool {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}
	shouldBag := pidm.getBag(itemTemplate.GetBagType())
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
	}
	return shouldBag.NumOfItemsWithProperty(itemId) >= num
}

//数量
func (pidm *PlayerInventoryDataManager) NumOfItems(itemId int32) int32 {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return 0
	}
	shouldBag := pidm.getBag(itemTemplate.GetBagType())
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
	}
	return shouldBag.NumOfItemsWithProperty(itemId)
}

//获取物品使用列表
func (pidm *PlayerInventoryDataManager) GetItemUseAll() map[int32]*PlayerItemUseObject {
	pidm.refreshItemUseTimes()
	return pidm.itemUseRecordMap
}

func (pidm *PlayerInventoryDataManager) getItemUse(itemId int32) *PlayerItemUseObject {
	itemUse, ok := pidm.itemUseRecordMap[itemId]
	if ok {
		return itemUse
	} else {
		return nil
	}
}

//获取物品使用次数
func (pidm *PlayerInventoryDataManager) GetItemUseTimes(itemId int32) (dayUseTimes int32, totalUseTimes int32) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return 0, 0
	}

	pidm.refreshItemUseTimes()

	dayUseTimes, totalUseTimes = pidm.getItemUseTimes(itemId)
	return
}

func (pidm *PlayerInventoryDataManager) getItemUseTimes(itemId int32) (dayUseTimes int32, totalUseTimes int32) {
	itemUseObj := pidm.getItemUse(itemId)
	if itemUseObj == nil {
		return
	}

	dayUseTimes = itemUseObj.TodayTimes
	totalUseTimes = itemUseObj.TotalTimes
	return
}

func (pidm *PlayerInventoryDataManager) refreshItemUseTimes() error {
	now := global.GetGame().GetTimeService().Now()
	useMap := make(map[int32]*PlayerItemUseObject)
	for _, useItemObj := range pidm.itemUseRecordMap {
		flag, err := timeutils.IsSameDay(useItemObj.LastUseTime, now)
		if err != nil {
			return err
		}
		if !flag {
			useItemObj.TodayTimes = 0
			useItemObj.LastUseTime = now
			useItemObj.UpdateTime = now
			// useItemObj.SetModified()
			useMap[useItemObj.ItemId] = useItemObj
		}
	}
	if len(useMap) == 0 {
		return nil
	}

	//物品使用信息改变
	gameevent.Emit(inventoryeventtypes.EventTypeItemUseChanged, pidm.p, useMap)
	return nil
}

//获取已有的和需要购买的物品
func (pidm *PlayerInventoryDataManager) GetItemsAndNeedBuy(items map[int32]int32) (hasItems map[int32]int32, needBuyItems map[int32]int32) {
	if len(items) == 0 {
		panic(fmt.Errorf("inventory:获取已有和需要购买的不能为空"))
	}
	newItems := mergeItems(items)
	hasItems = make(map[int32]int32)
	needBuyItems = make(map[int32]int32)

	for itemId, num := range newItems {
		if num <= 0 {
			panic(fmt.Errorf("inventory:num should more than 0"))
		}
		currentNum := pidm.NumOfItems(itemId)
		if num <= currentNum {
			hasItems[itemId] = num
		} else {
			if currentNum != 0 {
				hasItems[itemId] = currentNum
			}
			needBuyItems[itemId] = num - currentNum
		}
	}
	return
}

//批量移除
func (pidm *PlayerInventoryDataManager) BatchRemove(items map[int32]int32, reason commonlog.InventoryLogReason, reasonText string) bool {
	if !pidm.HasEnoughItems(items) {
		return false
	}
	for itemId, num := range items {
		// TODO:xzk:优先消耗绑定物品
		flag := pidm.UseItem(itemId, num, reason, reasonText)
		if !flag {
			panic(fmt.Errorf("inventory:user item should be ok"))
		}
	}

	return true
}

//合并一样的物品
func mergeItems(items map[int32]int32) (newItems map[int32]int32) {
	newItems = make(map[int32]int32)
	for itemId, itemNum := range items {
		num, exist := newItems[itemId]
		if !exist {
			newItems[itemId] = itemNum
		} else {
			newItems[itemId] = itemNum + num
		}
	}
	return newItems
}

//合并一样的物品
func mergeItemLevel(items []*droptemplate.DropItemData) (newItemList []*droptemplate.DropItemData) {
	for _, itemData := range items {
		itemTemp := item.GetItemService().GetItem(int(itemData.ItemId))
		if itemTemp == nil {
			continue
		}

		isMerge := false
		for _, data := range newItemList {
			if data == nil {
				continue
			}
			if itemTemp.IsGoldEquip() {
				continue
			}
			if itemTemp.IsBaoBaoCard() {
				continue
			}
			if itemTemp.IsTeRing() {
				continue
			}
			if data.IsMerge(itemData) {
				data.Num += itemData.Num
				isMerge = true
			}
		}
		if !isMerge {
			newData := droptemplate.CreateItemDataWithData(itemData)
			newItemList = append(newItemList, newData)
		}
	}

	return newItemList
}

//使用位置的数量(包含过期物品)
func (pidm *PlayerInventoryDataManager) RemoveIndex(bagType inventorytypes.BagType, index int32, num int32, reason commonlog.InventoryLogReason, reasonText string) (flag bool, err error) {
	if index < 0 || num <= 0 {
		panic(fmt.Errorf("inventory:id %d,use index %d,num %d", pidm.p.GetId(), index, num))
	}

	it := pidm.FindItemByIndex(bagType, index)
	if it == nil {
		return
	}
	if it.Num < num {
		return
	}
	shoulderBag := pidm.getBag(bagType)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", bagType.String()))
	}
	itemId := it.ItemId
	beforeNum := shoulderBag.NumOfItems(itemId)
	shoulderBag.RemoveIndex(index, num)

	data := inventoryeventtypes.CreatePlayerInventoryChangedLogEventData(itemId, beforeNum, num, reason, reasonText)
	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChangedLog, pidm.p, data)
	return true, nil
}

//使用位置的数量
func (pidm *PlayerInventoryDataManager) BatchRemoveIndex(bagType inventorytypes.BagType, indexList []int32, reason commonlog.InventoryLogReason, reasonText string) (flag bool, err error) {
	shoulderBag := pidm.getBag(bagType)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", bagType.String()))
	}
	for _, index := range indexList {
		if index < 0 {
			panic(fmt.Errorf("inventory:id %d,use index %d", pidm.p.GetId(), index))
		}

		it := pidm.FindItemByIndex(bagType, index)
		if it == nil {
			return
		}
		itemId := it.ItemId
		num := it.Num
		beforeNum := shoulderBag.NumOfItems(itemId)
		shoulderBag.RemoveIndex(index, num)

		data := inventoryeventtypes.CreatePlayerInventoryChangedLogEventData(itemId, beforeNum, num, reason, reasonText)
		gameevent.Emit(inventoryeventtypes.EventTypeInventoryChangedLog, pidm.p, data)
	}

	return true, nil
}

//使用位置的数量
func (pidm *PlayerInventoryDataManager) RemoveDepotByIndex(index int32, num int32) (flag bool, err error) {
	if index < 0 || num <= 0 {
		panic(fmt.Errorf("inventory:id %d,use index %d,num %d", pidm.p.GetId(), index, num))
	}

	it := pidm.FindDepotItemByIndex(index)
	if it == nil {
		return
	}
	if it.Num < num {
		return
	}
	shoulderBag := pidm.getDepotBag()
	if shoulderBag == nil {
		panic("inventory:仓库应该存在")
	}
	shoulderBag.RemoveIndex(index, num)

	return true, nil
}

//批量移除金装
func (pidm *PlayerInventoryDataManager) BatchGoldEquipRemoveIndex(indexList []int32, reason commonlog.InventoryLogReason, reasonText string) (flag bool, err error) {
	for _, itemIndex := range indexList {
		flag, err = pidm.RemoveIndex(inventorytypes.BagTypePrim, itemIndex, 1, reason, reasonText)
		if !flag {
			return
		}
	}
	return true, nil
}

//客户端使用
//使用位置的数量
func (pidm *PlayerInventoryDataManager) UseIndex(bagType inventorytypes.BagType, index int32, num int32, chooseIndexList []int32, args string, reason commonlog.InventoryLogReason, reasonText string) (flag bool, err error) {
	if index < 0 || num <= 0 {
		panic(fmt.Errorf("inventory:id %d,use index %d,num %d", pidm.p.GetId(), index, num))
	}

	it := pidm.FindItemByIndex(bagType, index)
	if it == nil {
		return
	}
	if it.Num < num {
		return
	}
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	//添加使用作用器
	useHandler := GetUseHandler(itemTemplate.GetItemType(), itemTemplate.GetItemSubType())
	if useHandler == nil {
		return
	}
	flag, err = useHandler.Use(pidm.p, it, num, chooseIndexList, args)
	if err != nil {
		return false, err
	}
	if !flag {
		return
	}
	shoulderBag := pidm.getBag(bagType)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
	}
	beforeNum := shoulderBag.NumOfItemsWithProperty(itemId)
	if itemTemplate.IsDelAfterUse() {
		shoulderBag.RemoveIndex(index, num)
	}

	pidm.addItemUseRecord(itemId, num)

	data := inventoryeventtypes.CreatePlayerInventoryChangedLogEventData(itemId, beforeNum, num, reason, reasonText)
	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChangedLog, pidm.p, data)
	useData := inventoryeventtypes.CreatePlayerInventoryItemUseEventData(itemId, num)
	gameevent.Emit(inventoryeventtypes.EventTypeUseItem, pidm.p, useData)
	return true, nil
}

//添加物品使用记录
func (pidm *PlayerInventoryDataManager) addItemUseRecord(itemId, num int32) error {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if !itemTemplate.IsLimitDayUseTimes() && !itemTemplate.IsLimitTotalUseTimes() && !itemTemplate.IsCDItem() {
		return nil
	}
	pidm.refreshItemUseTimes()

	now := global.GetGame().GetTimeService().Now()
	itemUseObj := pidm.getItemUse(itemId)
	if itemUseObj == nil {
		itemUseObj = NewPlayerItemUseObject(pidm.p)
		id, _ := idutil.GetId()
		itemUseObj.Id = id
		itemUseObj.ItemId = itemId
		itemUseObj.TodayTimes = 0
		itemUseObj.TotalTimes = 0
		itemUseObj.CreateTime = now

		pidm.itemUseRecordMap[itemId] = itemUseObj
	}
	itemUseObj.TodayTimes += num
	itemUseObj.TotalTimes += num
	itemUseObj.LastUseTime = now
	itemUseObj.UpdateTime = now

	itemUseObj.SetModified()

	//物品使用信息改变
	useMap := make(map[int32]*PlayerItemUseObject)
	useMap[itemUseObj.ItemId] = itemUseObj
	gameevent.Emit(inventoryeventtypes.EventTypeItemUseChanged, pidm.p, useMap)
	return nil
}

//查找物品
func (pidm *PlayerInventoryDataManager) FindItemByIndex(bagType inventorytypes.BagType, index int32) *PlayerItemObject {
	if index < 0 {
		panic(fmt.Errorf("inventory:id %d,find index %d", pidm.p.GetId(), index))
	}
	shoulderBag := pidm.getBag(bagType)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", bagType.String()))
	}

	return shoulderBag.GetByIndex(index)
}

//仓库查找物品
func (pidm *PlayerInventoryDataManager) FindDepotItemByIndex(index int32) *PlayerItemObject {
	if index < 0 {
		panic(fmt.Errorf("inventory:id %d,find index %d", pidm.p.GetId(), index))
	}
	shoulderBag := pidm.getDepotBag()
	if shoulderBag == nil {
		panic("inventory:仓库应该存在")
	}

	return shoulderBag.GetByIndex(index)
}

//背包金装升星失败
func (pidm *PlayerInventoryDataManager) GoldEquipUpstarReturn(index int32, returnLevel int32) bool {
	shoulderBag := pidm.getBag(inventorytypes.BagTypePrim)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.BagTypePrim.String()))
	}

	bagItem := shoulderBag.GetByIndex(index)
	if bagItem == nil {
		return false
	}

	propertyData, ok := bagItem.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	propertyData.UpstarLevel = returnLevel

	shoulderBag.UpdateItem(index)
	return true
}

//背包金装继承
func (pidm *PlayerInventoryDataManager) GoldEquipExtend(targetIndex, useIndex int32) bool {
	shoulderBag := pidm.getBag(inventorytypes.BagTypePrim)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.BagTypePrim.String()))
	}

	targetIt := shoulderBag.GetByIndex(targetIndex)
	if targetIt == nil {
		return false
	}
	useIt := shoulderBag.GetByIndex(useIndex)
	if useIt == nil {
		return false
	}

	targetPropertyData, ok := targetIt.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	usePropertyData, ok := useIt.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	beforLevel := targetPropertyData.UpstarLevel
	targetPropertyData.UpstarLevel = usePropertyData.UpstarLevel
	usePropertyData.UpstarLevel = 0

	shoulderBag.UpdateItem(targetIndex)
	shoulderBag.UpdateItem(useIndex)

	//金装继承日志
	extendReason := commonlog.GoldEquipLogReasonExtend
	extendReasonText := fmt.Sprintf(extendReason.String(), targetIt.ItemId, useIt.ItemId)
	eventData := goldequipeventtypes.CreatePlayerGoldEquipExtendLogEventData(beforLevel, usePropertyData.UpstarLevel, extendReason, extendReasonText)
	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipExtendLog, pidm.p, eventData)
	return true
}

//背包金装升星成功
func (pidm *PlayerInventoryDataManager) GoldEquipUpstarSuccess(index int32) bool {
	shoulderBag := pidm.getBag(inventorytypes.BagTypePrim)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.BagTypePrim.String()))
	}

	bagItem := shoulderBag.GetByIndex(index)
	if bagItem == nil {
		return false
	}

	propertyData, ok := bagItem.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	propertyData.UpstarLevel += 1

	shoulderBag.UpdateItem(index)
	return true
}

//背包金装开光
func (pidm *PlayerInventoryDataManager) GoldEquipOpenLight(index int32, isSuccess bool) bool {
	shoulderBag := pidm.getBag(inventorytypes.BagTypePrim)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.BagTypePrim.String()))
	}

	bagItem := shoulderBag.GetByIndex(index)
	if bagItem == nil {
		return false
	}

	propertyData, ok := bagItem.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	if isSuccess {
		propertyData.OpenLightLevel += 1
		propertyData.OpenTimes = 0
	} else {
		propertyData.OpenTimes += 1
	}

	shoulderBag.UpdateItem(index)
	return true
}

//背包金装强化成功
func (pidm *PlayerInventoryDataManager) UpdateGoldEquipLevel(index int32) bool {
	shoulderBag := pidm.getBag(inventorytypes.BagTypePrim)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.BagTypePrim.String()))
	}

	bagItem := shoulderBag.GetByIndex(index)
	if bagItem == nil {
		return false
	}

	curLevel := bagItem.Level
	curLevel += 1
	shoulderBag.UpdateItemLevel(index, curLevel)

	newItemData := droptemplate.CreateItemData(bagItem.ItemId, bagItem.Num, bagItem.Level, bagItem.BindType)
	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipStrengSuccess, pidm.p, newItemData)
	return true
}

// 背包金装道具强化
func (pidm *PlayerInventoryDataManager) UpdateGoldEquipLevelUseItem(targetIndex, itemId int32) bool {
	shoulderBag := pidm.getBag(inventorytypes.BagTypePrim)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.BagTypePrim.String()))
	}

	bagItem := shoulderBag.GetByIndex(targetIndex)
	if bagItem == nil {
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
	shoulderBag.UpdateItemLevel(targetIndex, toUplevel)

	newItemData := droptemplate.CreateItemData(bagItem.ItemId, bagItem.Num, bagItem.Level, bagItem.BindType)
	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipStrengSuccess, pidm.p, newItemData)
	return true
}

// 背包金装道具开光
func (pidm *PlayerInventoryDataManager) UpdateGoldEquipOpenLightUseItem(targetIndex, itemId int32) bool {
	shoulderBag := pidm.getBag(inventorytypes.BagTypePrim)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.BagTypePrim.String()))
	}

	bagItem := shoulderBag.GetByIndex(targetIndex)
	if bagItem == nil || bagItem.IsEmpty() {
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

	propertyData, ok := bagItem.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	toUplevel := itemTemp.TypeFlag1
	propertyData.OpenLightLevel = toUplevel
	propertyData.OpenTimes = 0
	shoulderBag.UpdateItem(targetIndex)

	return true
}

// 背包金装道具强化升星
func (pidm *PlayerInventoryDataManager) UpdateGoldEquipUpstarUseItem(targetIndex, itemId int32) bool {
	shoulderBag := pidm.getBag(inventorytypes.BagTypePrim)
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.BagTypePrim.String()))
	}

	bagItem := shoulderBag.GetByIndex(targetIndex)
	if bagItem == nil || bagItem.IsEmpty() {
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

	propertyData, ok := bagItem.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return false
	}
	toUplevel := itemTemp.TypeFlag1
	propertyData.UpstarLevel = toUplevel
	shoulderBag.UpdateItem(targetIndex)

	return true
}

//是否可以卖
func (pidm *PlayerInventoryDataManager) IfCanSell(bagType inventorytypes.BagType, index int32, num int32) bool {
	it := pidm.FindItemByIndex(bagType, index)
	if it == nil {
		return false
	}
	if it.IsEmpty() {
		return false
	}
	itemTemplate := item.GetItemService().GetItem(int(it.ItemId))
	if itemTemplate == nil {
		return false
	}
	if !itemTemplate.CanSell() {
		return false
	}
	if it.Num < num {
		return false
	}
	return true
}

//是否使用CD
func (pidm *PlayerInventoryDataManager) IsItemIndexUseCd(index int32) bool {
	it := pidm.FindItemByIndex(inventorytypes.BagTypePrim, index)
	if it == nil {
		return true
	}
	if it.IsEmpty() {
		return true
	}
	return pidm.IsItemUseCd(it.ItemId)
}

//是否使用CD
func (pidm *PlayerInventoryDataManager) IsItemUseCd(itemId int32) bool {
	itemUse := pidm.getItemUse(itemId)
	if itemUse == nil {
		return false
	}
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return true
	}
	//使用CD
	now := global.GetGame().GetTimeService().Now()
	lastUseTime := itemUse.LastUseTime
	itemCdTime := int64(itemTemplate.CdTime)
	return now-lastUseTime < itemCdTime
}

// //临时使用物品
// func (pidm *PlayerInventoryDataManager) UseItem2(itemId int32, num int32, reason commonlog.InventoryLogReason, reasonText string) bool {
// 	if itemId == 0 || num <= 0 {
// 		panic(fmt.Errorf("inventory:id %d,use itemId %d,num %d", pidm.p.GetId(), itemId, num))
// 	}

// 	//物品不存在
// 	itemTemplate := item.GetItemService().GetItem(int(itemId))
// 	if itemTemplate == nil {
// 		return false
// 	}
// 	shouldBag := pidm.getBag(itemTemplate.GetBagType())
// 	if shouldBag == nil {
// 		panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
// 	}

// 	//添加使用作用器
// 	useHandler := use.GetUseHandler(itemTemplate.GetItemType(), itemTemplate.GetItemSubType())
// 	if useHandler == nil {
// 		return false
// 	}
// 	flag, _ := useHandler.Use(pidm.p, itemId, num, nil, "")
// 	if !flag {
// 		return false
// 	}
// 	flag = shouldBag.RemoveItem(itemId, num)
// 	if !flag {
// 		return false
// 	}
// 	pidm.addItemUseRecord(itemId, num)
// 	return true
// }

//使用物品
func (pidm *PlayerInventoryDataManager) UseItem(itemId int32, num int32, reason commonlog.InventoryLogReason, reasonText string) bool {
	if itemId == 0 || num <= 0 {
		panic(fmt.Errorf("inventory:id %d,use itemId %d,num %d", pidm.p.GetId(), itemId, num))
	}

	//物品不存在
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}
	shouldBag := pidm.getBag(itemTemplate.GetBagType())
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
	}
	beforeNum := shouldBag.NumOfItemsWithProperty(itemId)
	flag := shouldBag.RemoveItem(itemId, num)
	if !flag {
		return false
	}

	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChanged, pidm.p, nil)
	data := inventoryeventtypes.CreatePlayerInventoryChangedLogEventData(itemId, beforeNum, num, reason, reasonText)
	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChangedLog, pidm.p, data)
	return true
}

//获取变化
func (pidm *PlayerInventoryDataManager) GetChangedSlotAndReset() (itemList []*PlayerItemObject) {
	for _, shoulderBag := range pidm.shoulderBagMap {
		itemList = append(itemList, shoulderBag.GetChangedSlotAndReset()...)
	}
	return
}

func (pidm *PlayerInventoryDataManager) GetShoulderBagMap() map[inventorytypes.BagType]*ShoulderBag {
	return pidm.shoulderBagMap
}

//获取仓库变化
func (pidm *PlayerInventoryDataManager) GetDepotChangedSlotAndReset() (itemList []*PlayerItemObject) {
	return pidm.depotBag.GetChangedSlotAndReset()
}

//GM：清空所有物品
func (pidm *PlayerInventoryDataManager) GMClearAll() {
	for _, shoulderBag := range pidm.shoulderBagMap {
		shoulderBag.ClearAll()
	}
}

//GM：重置物品使用次数
func (pidm *PlayerInventoryDataManager) GMResetTimes() {
	for _, itemUse := range pidm.GetItemUseAll() {
		itemUse.TotalTimes = 0
		itemUse.TodayTimes = 0
		itemUse.SetModified()
	}

	//物品使用信息改变
	gameevent.Emit(inventoryeventtypes.EventTypeItemUseChanged, pidm.p, pidm.GetItemUseAll())
}

func (pidm *PlayerInventoryDataManager) addItem(itemId int32, num int32, reason commonlog.InventoryLogReason, reasonText string) bool {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}
	shouldBag := pidm.getBag(itemTemplate.GetBagType())
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", itemTemplate.GetBagType().String()))
	}
	beforeNum := shouldBag.NumOfItemsWithProperty(itemId)
	flag := shouldBag.AddLevelItem(itemId, num, 0, inventorytypes.IsDepotTypePrim, itemtypes.ItemBindTypeUnBind)
	if !flag {
		return false
	}

	data := inventoryeventtypes.CreatePlayerInventoryChangedLogEventData(itemId, beforeNum, num, reason, reasonText)
	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChangedLog, pidm.p, data)

	return true
}

//合并
func (pidm *PlayerInventoryDataManager) Merge(bagType inventorytypes.BagType) {
	shouldBag := pidm.getBag(bagType)
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", bagType.String()))
	}
	shouldBag.Merge()
}

//仓库合并
func (pidm *PlayerInventoryDataManager) MergeDepot() {
	shouldBag := pidm.getDepotBag()
	if shouldBag == nil {
		panic("inventory:仓库应该存在")
	}
	shouldBag.Merge()
}

//
func (pidm *PlayerInventoryDataManager) EquipmentSlotStrengthStar(pos inventorytypes.BodyPositionType) (flag bool) {
	flag = pidm.equipmentBag.StrengthStar(pos)
	return
}

//使用装备
func (pidm *PlayerInventoryDataManager) PutOn(pos inventorytypes.BodyPositionType, itemId int32, bindType itemtypes.ItemBindType) (flag bool) {
	flag = pidm.equipmentBag.PutOn(pos, itemId, bindType)
	return
}

//脱下装备
func (pidm *PlayerInventoryDataManager) TakeOff(pos inventorytypes.BodyPositionType) (itemId int32) {
	//判断是否可以脱下
	flag := pidm.IfCanTakeOff(pos)
	if !flag {
		return
	}

	itemId = pidm.equipmentBag.TakeOff(pos)
	return
}

//获取装备
func (pidm *PlayerInventoryDataManager) GetEquipByPos(pos inventorytypes.BodyPositionType) *PlayerEquipmentSlotObject {
	item := pidm.equipmentBag.GetByPosition(pos)
	if item == nil {
		return nil
	}

	return item
}

//获取装备的总等级
func (pidm *PlayerInventoryDataManager) GetEquipTotalLevel() (totalLevel int32) {
	totalLevel = 0
	for _, equipmentSlotObj := range pidm.equipmentBag.slotMap {
		totalLevel += equipmentSlotObj.Level
	}
	return
}

//镶嵌宝石数量
func (pidm *PlayerInventoryDataManager) GetEquipGemNum() (num int32) {
	num = 0
	for _, equipmentSlotObj := range pidm.equipmentBag.slotMap {
		num += int32(len(equipmentSlotObj.GemInfo))
	}
	return
}

//获取下一阶装备
func (pidm *PlayerInventoryDataManager) GetNextEquipment(pos inventorytypes.BodyPositionType) *gametemplate.ItemTemplate {
	it := pidm.GetEquipByPos(pos)
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

//是否可以卸下
func (pidm *PlayerInventoryDataManager) IfCanTakeOff(pos inventorytypes.BodyPositionType) bool {
	item := pidm.GetEquipByPos(pos)
	if item == nil {
		return false
	}
	if item.IsEmpty() {
		return false
	}
	return true
}

//装备改变
func (pidm *PlayerInventoryDataManager) GetChangedEquipmentSlotAndReset() (itemList []*PlayerEquipmentSlotObject) {
	return pidm.equipmentBag.GetChangedSlotAndReset()
}

//获取背包物品
func (pidm *PlayerInventoryDataManager) GetBagAll(bagType inventorytypes.BagType) []*PlayerItemObject {
	shouldBag := pidm.getBag(bagType)
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", bagType.String()))
	}
	return shouldBag.GetItemList()
}

//获取背包物品
func (pidm *PlayerInventoryDataManager) GetEmptySlots(bagType inventorytypes.BagType) int32 {
	shouldBag := pidm.getBag(bagType)
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", bagType.String()))
	}
	return shouldBag.GetEmptySlots()
}

//获取背包物品
func (pidm *PlayerInventoryDataManager) IsContainBindItem(itemId int32) bool {
	shouldBag := pidm.getBag(inventorytypes.BagTypePrim)
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.BagTypePrim.String()))
	}

	for _, it := range shouldBag.GetItemList() {
		if it.ItemId != itemId {
			continue
		}
		if it.BindType == itemtypes.ItemBindTypeBind {
			return true
		}
	}
	return false
}

//获取仓库物品
func (pidm *PlayerInventoryDataManager) GetDepotAll() []*PlayerItemObject {
	shouldBag := pidm.getDepotBag()
	if shouldBag == nil {
		panic("inventory:仓库应该存在")
	}
	return shouldBag.GetItemList()
}

//获取所有
func (pidm *PlayerInventoryDataManager) GetAll() (itemList []*PlayerItemObject) {
	for _, shoulderBag := range pidm.shoulderBagMap {
		itemList = append(itemList, shoulderBag.GetItemList()...)
	}
	return
}

//获取背包最后一个道具
func (pidm *PlayerInventoryDataManager) GetBageLastItem(itemId int32) (item *PlayerItemObject) {
	shoulderBag := pidm.getBag(inventorytypes.BagTypePrim)
	itemList := shoulderBag.GetItemList()
	targetIndex := int32(-1)
	for index, ito := range itemList {
		if ito.IsEmpty() {
			continue
		}

		if ito.ItemId != itemId {
			continue
		}
		if ito.BindType == itemtypes.ItemBindTypeBind {
			continue
		}

		targetIndex = int32(index)
	}

	if targetIndex >= 0 {
		return itemList[targetIndex]
	}
	return nil
}

//获取所有鲲
func (pidm *PlayerInventoryDataManager) GetAllKun() map[int32]int32 {
	itemMap := make(map[int32]int32)
	kunShoulder, exist := pidm.shoulderBagMap[inventorytypes.BagTypeKun]
	if !exist {
		return nil
	}
	itemList := kunShoulder.GetItemList()
	for _, itemObj := range itemList {
		if itemObj.ItemId == 0 {
			continue
		}
		curNum := itemMap[itemObj.ItemId]
		curNum += itemObj.Num
		itemMap[itemObj.ItemId] = curNum
	}
	return itemMap
}

////////////////////////// 秘宝仓库 材料仓库 ///////////////////

//秘宝仓库合并
func (pidm *PlayerInventoryDataManager) MergeMiBaoDepot(typ equipbaokutypes.BaoKuType) {
	var shouldBag *ShoulderBag
	if typ == equipbaokutypes.BaoKuTypeEquip {
		shouldBag = pidm.getMiBaoDepotBag()
	} else {
		shouldBag = pidm.getMaterialDepotBag()
	}
	if shouldBag == nil {
		panic("inventory:宝库应该存在")
	}
	shouldBag.Merge()
}

//获取仓库物品
func (pidm *PlayerInventoryDataManager) GetMiBaoDepotAll() []*PlayerItemObject {
	shouldBag := pidm.getMiBaoDepotBag()
	if shouldBag == nil {
		panic("inventory:宝库应该存在")
	}
	return shouldBag.GetItemList()
}

//获取仓库物品
func (pidm *PlayerInventoryDataManager) GetMaterialDepotAll() []*PlayerItemObject {
	shouldBag := pidm.getMaterialDepotBag()
	if shouldBag == nil {
		panic("inventory:宝库应该存在")
	}
	return shouldBag.GetItemList()
}

//获取秘宝仓库变化
func (pidm *PlayerInventoryDataManager) GetMiBaoDepotChangedSlotAndReset(typ equipbaokutypes.BaoKuType) (itemList []*PlayerItemObject) {
	if typ == equipbaokutypes.BaoKuTypeEquip {
		return pidm.mibaoDepotBag.GetChangedSlotAndReset()
	}

	return pidm.materialDepotBag.GetChangedSlotAndReset()
}

//秘宝仓库是否有足够的位置存放物品
func (pidm *PlayerInventoryDataManager) HasEnoughMiBaoDepotSlot(itemId int32, num int32, level int32, bindType itemtypes.ItemBindType, typ equipbaokutypes.BaoKuType) bool {
	if itemId == 0 || num <= 0 {
		panic(fmt.Errorf("inventory:id %d,has enough mibao depot slot itemId %d,num %d", pidm.p.GetId(), itemId, num))
	}
	remainNum := pidm.RemainMiBaoDepotSlotForItem(itemId, level, bindType, typ)
	if remainNum >= num {
		return true
	}
	return false
}

//秘宝仓库剩余位置
func (pidm *PlayerInventoryDataManager) RemainMiBaoDepotSlotForItem(itemId int32, level int32, bindType itemtypes.ItemBindType, typ equipbaokutypes.BaoKuType) (maxNum int32) {
	var shouldBag *ShoulderBag
	if typ == equipbaokutypes.BaoKuTypeEquip {
		shouldBag = pidm.getMiBaoDepotBag()
	} else {
		shouldBag = pidm.getMaterialDepotBag()
	}
	if shouldBag == nil {
		panic("inventory:宝库应该存在")
	}
	return shouldBag.RemainSlotForItemLevel(itemId, level, bindType)
}

//保存到秘宝仓库
func (pidm *PlayerInventoryDataManager) AddItemInMiBaoDepot(itemId int32, num int32, level int32, bind itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData, typ equipbaokutypes.BaoKuType) bool {
	//物品不存在
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}

	var shouldBag *ShoulderBag
	if typ == equipbaokutypes.BaoKuTypeEquip {
		shouldBag = pidm.getMiBaoDepotBag()
		if shouldBag == nil {
			panic("inventory：宝库应该存在")
		}
		flag := shouldBag.AddLevelItemWithPropertyData(itemId, num, level, propertyData, inventorytypes.IsDepotTypeMiBao, bind)
		if !flag {
			return false
		}
	} else {
		shouldBag = pidm.getMaterialDepotBag()
		if shouldBag == nil {
			panic("inventory：宝库应该存在")
		}
		flag := shouldBag.AddLevelItemWithPropertyData(itemId, num, level, propertyData, inventorytypes.IsDepotTypeMaterial, bind)
		if !flag {
			return false
		}
	}

	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChanged, pidm.p, nil)
	return true
}

//秘宝仓库查找物品
func (pidm *PlayerInventoryDataManager) FindMiBaoDepotItemByIndex(index int32, typ equipbaokutypes.BaoKuType) *PlayerItemObject {
	if index < 0 {
		panic(fmt.Errorf("inventory:id %d,find index %d", pidm.p.GetId(), index))
	}
	var shoulderBag *ShoulderBag
	if typ == equipbaokutypes.BaoKuTypeEquip {
		shoulderBag = pidm.getMiBaoDepotBag()
	} else {
		shoulderBag = pidm.getMaterialDepotBag()
	}
	if shoulderBag == nil {
		panic("inventory:宝库应该存在")
	}

	return shoulderBag.GetByIndex(index)
}

//秘宝仓库查找物品引索从最后开始
func (pidm *PlayerInventoryDataManager) FindMiBaoDepotItemIndexsFromEnd(typ equipbaokutypes.BaoKuType) []int32 {
	var shoulderBag *ShoulderBag
	if typ == equipbaokutypes.BaoKuTypeEquip {
		shoulderBag = pidm.getMiBaoDepotBag()
	} else {
		shoulderBag = pidm.getMaterialDepotBag()
	}
	if shoulderBag == nil {
		panic("inventory:仓库应该存在")
	}
	itemList := shoulderBag.GetItemList()
	var indexList []int32
	index := int32(len(itemList) - 1)
	for ; index >= 0; index-- {
		if itemList[index] == nil || itemList[index].IsEmpty() {
			continue
		}
		indexList = append(indexList, index)
	}
	return indexList
}

//使用位置的数量
func (pidm *PlayerInventoryDataManager) RemoveMiBaoDepotByIndex(index int32, num int32, typ equipbaokutypes.BaoKuType) (flag bool, err error) {
	if index < 0 || num <= 0 {
		panic(fmt.Errorf("inventory:id %d,use index %d,num %d", pidm.p.GetId(), index, num))
	}

	it := pidm.FindMiBaoDepotItemByIndex(index, typ)
	if it == nil {
		return
	}
	if it.Num < num {
		return
	}
	var shoulderBag *ShoulderBag
	if typ == equipbaokutypes.BaoKuTypeEquip {
		shoulderBag = pidm.getMiBaoDepotBag()
	} else {
		shoulderBag = pidm.getMaterialDepotBag()
	}
	if shoulderBag == nil {
		panic("inventory:秘宝仓库应该存在")
	}
	shoulderBag.RemoveIndex(index, num)

	return true, nil
}

//使用位置的数量
func (pidm *PlayerInventoryDataManager) BatchRemoveMiBaoDepotByIndex(indexList []int32, reason commonlog.InventoryLogReason, reasonText string, typ equipbaokutypes.BaoKuType) (flag bool, err error) {
	var shoulderBag *ShoulderBag
	if typ == equipbaokutypes.BaoKuTypeEquip {
		shoulderBag = pidm.getMiBaoDepotBag()
	} else {
		shoulderBag = pidm.getMaterialDepotBag()
	}

	if shoulderBag == nil {
		panic("inventory:秘宝仓库应该存在")
	}
	for _, index := range indexList {
		if index < 0 {
			panic(fmt.Errorf("inventory:id %d,use index %d", pidm.p.GetId(), index))
		}

		it := pidm.FindMiBaoDepotItemByIndex(index, typ)
		if it == nil {
			return
		}
		itemId := it.ItemId
		num := it.Num
		beforeNum := shoulderBag.NumOfItems(itemId)
		shoulderBag.RemoveIndex(index, num)

		data := inventoryeventtypes.CreatePlayerInventoryChangedLogEventData(itemId, beforeNum, num, reason, reasonText)
		gameevent.Emit(inventoryeventtypes.EventTypeInventoryChangedLog, pidm.p, data)
	}

	return true, nil
}

//添加含等级物品到秘宝仓库
func (pidm *PlayerInventoryDataManager) AddItemLevelMiBao(itemData *droptemplate.DropItemData, reason commonlog.InventoryLogReason, reasonText string, typ equipbaokutypes.BaoKuType) bool {
	itemId := itemData.GetItemId()
	num := itemData.GetNum()
	level := itemData.GetLevel()
	bind := itemData.GetBindType()
	expireType := itemData.GetExpireType()
	expireTime := itemData.GetExpireTime()
	itemGetTime := itemData.GetItemGetTime()
	//物品不存在
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}

	var shouldBag *ShoulderBag
	if typ == equipbaokutypes.BaoKuTypeEquip {
		shouldBag = pidm.getMiBaoDepotBag()
	} else {
		shouldBag = pidm.getMaterialDepotBag()
	}
	if shouldBag == nil {
		panic(fmt.Errorf("inventory:背包[%s]应该存在", inventorytypes.IsDepotTypeMiBao.String()))
	}

	beforeNum := shouldBag.NumOfItemsWithProperty(itemId)
	if itemTemplate.IsGoldEquip() {
		upstar := itemData.GetUpstar()
		attrList := itemData.GetAttrList()
		for i := 0; i < int(num); i++ {
			baseProperty := inventorytypes.CreateItemPropertyDataBase(expireType, expireTime, itemGetTime)
			propertyData := inventory.CreatePropertyDataInterface(itemTemplate.GetItemType(), baseProperty)
			goldequipPropertyData := propertyData.(*goldequiptypes.GoldEquipPropertyData)
			goldequipPropertyData.UpstarLevel = upstar
			if itemData.IsRandomAttr {
				goldequipPropertyData.AttrList = attrList
				goldequipPropertyData.IsHadCountAttr = true
			}
			if typ == equipbaokutypes.BaoKuTypeEquip {
				flag := shouldBag.AddLevelItemWithPropertyData(itemId, 1, level, goldequipPropertyData, inventorytypes.IsDepotTypeMiBao, bind)
				if !flag {
					return false
				}
			} else {
				flag := shouldBag.AddLevelItemWithPropertyData(itemId, 1, level, goldequipPropertyData, inventorytypes.IsDepotTypeMaterial, bind)
				if !flag {
					return false
				}
			}
		}
	} else {
		baseProperty := inventorytypes.CreateItemPropertyDataBase(expireType, expireTime, itemGetTime)
		propertyData := inventory.CreatePropertyDataInterface(itemTemplate.GetItemType(), baseProperty)
		if typ == equipbaokutypes.BaoKuTypeEquip {
			flag := shouldBag.AddLevelItemWithPropertyData(itemId, num, level, propertyData, inventorytypes.IsDepotTypeMiBao, bind)
			if !flag {
				return false
			}
		} else {
			flag := shouldBag.AddLevelItemWithPropertyData(itemId, num, level, propertyData, inventorytypes.IsDepotTypeMaterial, bind)
			if !flag {
				return false
			}
		}
	}

	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChanged, pidm.p, nil)
	data := inventoryeventtypes.CreatePlayerInventoryChangedLogEventData(itemId, beforeNum, num, reason, reasonText)
	gameevent.Emit(inventoryeventtypes.EventTypeInventoryChangedLog, pidm.p, data)
	return true
}

//批量添加物品-秘宝
func (pidm *PlayerInventoryDataManager) BatchAddOfItemLevelMiBao(itemList []*droptemplate.DropItemData, reason commonlog.InventoryLogReason, reasonText string, typ equipbaokutypes.BaoKuType) bool {
	newItems := mergeItemLevel(itemList)
	if !pidm.HasEnoughSlotsOfItemLevelMiBao(newItems, typ) {
		return false
	}
	for _, itemData := range newItems {
		flag := pidm.AddItemLevelMiBao(itemData, reason, reasonText, typ)
		if !flag {
			return false
		}
	}
	return true
}

//是否有足够的位置存放物品-多个-秘宝
func (pidm *PlayerInventoryDataManager) HasEnoughSlotsOfItemLevelMiBao(itemList []*droptemplate.DropItemData, typ equipbaokutypes.BaoKuType) bool {
	var shoulderBag *ShoulderBag
	if typ == equipbaokutypes.BaoKuTypeEquip {
		shoulderBag = pidm.getMiBaoDepotBag()
	} else {
		shoulderBag = pidm.getMaterialDepotBag()
	}
	if shoulderBag == nil {
		panic(fmt.Errorf("inventory:宝库[%s]应该存在", inventorytypes.IsDepotTypeMiBao.String()))
	}

	needSlot := int32(0)
	for _, itemData := range itemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		level := itemData.GetLevel()
		bind := itemData.GetBindType()
		expireType := itemData.GetExpireType()

		itemTemplate := item.GetItemService().GetItem(int(itemId))
		if itemTemplate == nil {
			return false
		}

		slotNum := shoulderBag.CountNeedSlotOfItemLevelWithProperty(itemId, num, level, bind, expireType, itemData.GetItemGetTime(), itemData.GetExpireTimestamp())
		needSlot += slotNum
	}

	if needSlot > shoulderBag.GetEmptySlots() {
		return false
	}
	return true
}

/////////////////////////////////////////////

//心跳
func (pidm *PlayerInventoryDataManager) Heartbeat() {
	pidm.hbRunner.Heartbeat()
}

//获取装备物品
func (pidm *PlayerInventoryDataManager) GetEquipmentSlots() []*PlayerEquipmentSlotObject {
	return pidm.equipmentBag.GetAll()
}

//获取装备物品
func (pidm *PlayerInventoryDataManager) ClearAllEquipmentGemInfo() {
	pidm.equipmentBag.ClearAllEquipmentGemInfo()
}

//获取主背包
func (pidm *PlayerInventoryDataManager) getPrimBag() *ShoulderBag {
	return pidm.getBag(inventorytypes.BagTypePrim)
}

//获取主背包
func (pidm *PlayerInventoryDataManager) getBag(bagType inventorytypes.BagType) *ShoulderBag {
	bag, exist := pidm.shoulderBagMap[bagType]
	if !exist {
		return nil
	}
	return bag
}

func (m *PlayerInventoryDataManager) ToEquipmentSlotList() (slotInfoList []*inventorytypes.EquipmentSlotInfo) {
	for _, slot := range m.equipmentBag.GetAll() {
		slotInfo := &inventorytypes.EquipmentSlotInfo{}
		slotInfo.SlotId = int32(slot.SlotId)
		slotInfo.Level = slot.Level
		slotInfo.Star = slot.Star
		slotInfo.ItemId = slot.ItemId
		slotInfo.Gems = make(map[int32]int32)
		for pos, itemId := range slot.GemInfo {
			slotInfo.Gems[int32(pos)] = itemId
		}
		slotInfoList = append(slotInfoList, slotInfo)
	}
	return
}

func CreatePlayerInventoryDataManager(p player.Player) player.PlayerDataManager {
	pidm := &PlayerInventoryDataManager{}
	pidm.p = p
	pidm.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	return pidm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerInventoryDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerInventoryDataManager))
}
