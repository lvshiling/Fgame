package template

import (
	"fgame/fgame/core/template"
	droptypes "fgame/fgame/game/drop/types"
	"fgame/fgame/game/global"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sort"
	"sync"
)

//掉落接口处理
type DropTemplateService interface {
	//获取掉落模板
	GetDropFromGroup(dropId int32) *gametemplate.DropTemplate
	//获取掉落物品
	GetDropItem(dropId int32) (itemId int32, num int32)
	//获取掉落物品
	GetDropListItems(dropList []int32) map[int32]int32

	//获取掉落等级物品
	GetDropListItemLevelList(dropList []int32) []*DropItemData
	//获取掉落等级物品
	GetDropItemLevel(dropId int32) *DropItemData
	//获取掉落等级物品
	GetDropItemLevelWithIndex(dropId int32) (int32, *DropItemData)

	//校验必掉落的
	CheckSureDrop(dropId int32) (flag bool)

	//宝库掉落接口begin//////////
	//获取掉落模板
	GetDropBaoKuFromGroup(dropId int32) *gametemplate.DropBaoKuTemplate
	//获取掉落物品
	GetDropBaoKuItem(dropId int32) (itemId int32, num int32)
	//获取掉落物品
	GetDropBaoKuListItems(dropList []int32) map[int32]int32
	//获取掉落等级物品
	GetDropBaoKuListItemLevelList(dropList []int32) []*DropItemData
	//获取掉落等级物品
	GetDropBaoKuItemLevel(dropId int32) *DropItemData
	//获取掉落等级物品
	GetDropBaoKuItemLevelWithIndex(dropId int32) (int32, *DropItemData)
	//校验必掉落的
	CheckSureDropBaoKu(dropId int32) (flag bool)
	//end/////////
}

//分组模板排序类型
type dropTemplateList []*gametemplate.DropTemplate

func (s dropTemplateList) Len() int           { return len(s) }
func (s dropTemplateList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s dropTemplateList) Less(i, j int) bool { return s[i].Id < s[j].Id }

//宝库掉落分组模板排序类型
type dropBaoKuTemplateList []*gametemplate.DropBaoKuTemplate

func (s dropBaoKuTemplateList) Len() int           { return len(s) }
func (s dropBaoKuTemplateList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s dropBaoKuTemplateList) Less(i, j int) bool { return s[i].Id < s[j].Id }

type dropTemplateService struct {
	//掉落组配置
	dropGroupMap map[int32][]*gametemplate.DropTemplate
	//掉落组概率
	dropRateGroupMap map[int32][]int64
	//默认掉落
	defaultTemplate *gametemplate.DropTemplate

	//装备宝库掉落begin
	//掉落组配置
	dropBaoKuGroupMap map[int32][]*gametemplate.DropBaoKuTemplate
	//掉落组概率
	dropBaoKuRateGroupMap map[int32][]int64
	//默认掉落
	defaultBaoKuTemplate *gametemplate.DropBaoKuTemplate
	//end
}

//初始化
func (ds *dropTemplateService) init() error {
	ds.dropGroupMap = make(map[int32][]*gametemplate.DropTemplate)
	ds.dropRateGroupMap = make(map[int32][]int64)
	ds.dropBaoKuGroupMap = make(map[int32][]*gametemplate.DropBaoKuTemplate)
	ds.dropBaoKuRateGroupMap = make(map[int32][]int64)

	//掉落
	templateMap := template.GetTemplateService().GetAll((*gametemplate.DropTemplate)(nil))
	for _, templateObject := range templateMap {
		dropTemplate, _ := templateObject.(*gametemplate.DropTemplate)

		//默认掉落
		if dropTemplate.Id == 1 {
			ds.defaultTemplate = dropTemplate
		}

		//掉落组配置
		dropList := ds.dropGroupMap[dropTemplate.DropId]
		dropList = append(dropList, dropTemplate)
		ds.dropGroupMap[dropTemplate.DropId] = dropList
	}

	// 排序
	for _, dropTempList := range ds.dropGroupMap {
		sort.Sort(dropTemplateList(dropTempList))
		for _, dropTemp := range dropTempList {
			//掉落组概率
			rate := dropTemp.Rate
			dropRateList := ds.dropRateGroupMap[dropTemp.DropId]
			dropRateList = append(dropRateList, rate)
			ds.dropRateGroupMap[dropTemp.DropId] = dropRateList
		}
	}

	//默认掉落
	if ds.defaultTemplate == nil {
		return fmt.Errorf("dropService: default template dropId=1 should be  exist ")
	}

	//宝库掉落begin
	baoKuTemplateMap := template.GetTemplateService().GetAll((*gametemplate.DropBaoKuTemplate)(nil))
	for _, templateObject := range baoKuTemplateMap {
		dropTemplate, _ := templateObject.(*gametemplate.DropBaoKuTemplate)

		//默认掉落
		if dropTemplate.Id == 1 {
			ds.defaultBaoKuTemplate = dropTemplate
		}

		//掉落组配置
		dropList := ds.dropBaoKuGroupMap[dropTemplate.DropId]
		dropList = append(dropList, dropTemplate)
		ds.dropBaoKuGroupMap[dropTemplate.DropId] = dropList
	}

	// 排序
	for _, dropTempList := range ds.dropBaoKuGroupMap {
		sort.Sort(dropBaoKuTemplateList(dropTempList))
		for _, dropTemp := range dropTempList {
			//掉落组概率
			rate := dropTemp.Rate
			dropRateList := ds.dropBaoKuRateGroupMap[dropTemp.DropId]
			dropRateList = append(dropRateList, rate)
			ds.dropBaoKuRateGroupMap[dropTemp.DropId] = dropRateList
		}
	}

	//默认掉落
	if ds.defaultBaoKuTemplate == nil {
		return fmt.Errorf("dropService: default baokudrop template dropId=1 should be  exist ")
	}
	//end

	return nil
}

//获取掉落模板
func (ds *dropTemplateService) GetDropFromGroup(dropId int32) *gametemplate.DropTemplate {
	dropRateList, ok := ds.dropRateGroupMap[dropId]
	if !ok {
		return ds.defaultTemplate
	}
	dropList, ok := ds.dropGroupMap[dropId]
	if !ok {
		return ds.defaultTemplate
	}
	index := mathutils.RandomWeightsFromTotalWeights(dropRateList, droptypes.MAX_RATE)
	//随机不到
	if index == -1 {
		return nil
	}
	return dropList[index]
}

//获取掉落模板
func (ds *dropTemplateService) GetDropFromGroupWithIndex(dropId int32) (int32, *gametemplate.DropTemplate) {
	dropRateList, ok := ds.dropRateGroupMap[dropId]
	if !ok {
		return 0, ds.defaultTemplate
	}
	dropList, ok := ds.dropGroupMap[dropId]
	if !ok {
		return 0, ds.defaultTemplate
	}
	index := mathutils.RandomWeightsFromTotalWeights(dropRateList, droptypes.MAX_RATE)
	//随机不到
	if index == -1 {
		return 0, nil
	}
	return int32(index), dropList[index]
}

//校验dropId必掉的
func (ds *dropTemplateService) CheckSureDrop(dropId int32) (flag bool) {
	_, ok := ds.dropGroupMap[dropId]
	if !ok {
		return
	}
	dropRateList, ok := ds.dropRateGroupMap[dropId]
	if !ok {
		return
	}
	totalRate := int64(0)
	for _, rate := range dropRateList {
		totalRate += rate
	}

	if totalRate < droptypes.MAX_RATE {
		return
	}
	flag = true
	return
}

//获取掉落物品
func (ds *dropTemplateService) GetDropItem(dropId int32) (itemId int32, num int32) {
	to := ds.GetDropFromGroup(dropId)
	if to == nil {
		return
	}
	num = to.RandomNum()
	itemId = to.ItemId
	return
}

//获取掉落物品
func (ds *dropTemplateService) GetDropListItems(dropList []int32) map[int32]int32 {
	if len(dropList) <= 0 {
		return nil
	}
	dropItemMap := make(map[int32]int32)
	for _, dropId := range dropList {
		itemId, num := ds.GetDropItem(dropId)
		if itemId == 0 {
			continue
		}
		_, ok := dropItemMap[itemId]
		if !ok {
			dropItemMap[itemId] += num
		} else {
			dropItemMap[itemId] = num
		}
	}
	return dropItemMap
}

//获取掉落等级物品
func (ds *dropTemplateService) GetDropItemLevel(dropId int32) (newData *DropItemData) {
	to := ds.GetDropFromGroup(dropId)
	if to == nil {
		return
	}
	itemId := to.ItemId
	if itemId == 0 {
		return
	}

	num := to.RandomNum()
	level := to.RandomGoldEquipLevel()
	bind := to.GetBindType()
	star := to.RandomGoldEquipUpstarLevel()
	attrList, isRandom := to.RandomGoldEquipAttr()
	if !isRandom {
		itemTemp := item.GetItemService().GetItem(int(itemId))
		if itemTemp == nil {
			return
		}

		if itemTemp.IsGoldEquip() {
			attrList = itemTemp.GetGoldEquipTemplate().RandomGoldEquipAttr()
			isRandom = true
		}
	}
	expireType := inventorytypes.NewItemLimitTimeTypeNone
	expireTime := int64(0)
	itemGetTime := global.GetGame().GetTimeService().Now()
	newData = CreateItemDataWithPropertyData(itemId, num, level, bind, star, attrList, isRandom, expireType, expireTime, itemGetTime)
	return
}

//获取掉落等级物品
func (ds *dropTemplateService) GetDropItemLevelWithIndex(dropId int32) (index int32, newData *DropItemData) {
	index, to := ds.GetDropFromGroupWithIndex(dropId)
	if to == nil {
		return
	}
	itemId := to.ItemId
	if itemId == 0 {
		return
	}

	num := to.RandomNum()
	level := to.RandomGoldEquipLevel()
	bind := to.GetBindType()
	star := to.RandomGoldEquipUpstarLevel()
	attrList, isRandom := to.RandomGoldEquipAttr()
	if !isRandom {
		itemTemp := item.GetItemService().GetItem(int(itemId))
		if itemTemp == nil {
			return
		}

		if itemTemp.IsGoldEquip() {
			attrList = itemTemp.GetGoldEquipTemplate().RandomGoldEquipAttr()
			isRandom = true
		}
	}
	expireType := inventorytypes.NewItemLimitTimeTypeNone
	expireTime := int64(0)
	itemGetTime := global.GetGame().GetTimeService().Now()
	newData = CreateItemDataWithPropertyData(itemId, num, level, bind, star, attrList, isRandom, expireType, expireTime, itemGetTime)
	return
}

//获取掉落等级物品
func (ds *dropTemplateService) GetDropListItemLevelList(dropList []int32) []*DropItemData {
	if len(dropList) <= 0 {
		return nil
	}
	var dropItemList []*DropItemData
	for _, dropId := range dropList {
		newData := ds.GetDropItemLevel(dropId)
		if newData == nil {
			continue
		}
		dropItemList = append(dropItemList, newData)
	}
	return dropItemList
}

//宝库掉落begin///////////////////////
//获取掉落模板
func (ds *dropTemplateService) GetDropBaoKuFromGroup(dropId int32) *gametemplate.DropBaoKuTemplate {
	dropRateList, ok := ds.dropBaoKuRateGroupMap[dropId]
	if !ok {
		return ds.defaultBaoKuTemplate
	}
	dropList, ok := ds.dropBaoKuGroupMap[dropId]
	if !ok {
		return ds.defaultBaoKuTemplate
	}
	index := mathutils.RandomWeightsFromTotalWeights(dropRateList, droptypes.MAX_RATE)
	//随机不到
	if index == -1 {
		return nil
	}
	return dropList[index]
}

//获取掉落模板
func (ds *dropTemplateService) GetDropBaoKuFromGroupWithIndex(dropId int32) (int32, *gametemplate.DropBaoKuTemplate) {
	dropRateList, ok := ds.dropBaoKuRateGroupMap[dropId]
	if !ok {
		return 0, ds.defaultBaoKuTemplate
	}
	dropList, ok := ds.dropBaoKuGroupMap[dropId]
	if !ok {
		return 0, ds.defaultBaoKuTemplate
	}
	index := mathutils.RandomWeightsFromTotalWeights(dropRateList, droptypes.MAX_RATE)
	//随机不到
	if index == -1 {
		return 0, nil
	}
	return int32(index), dropList[index]
}

//校验dropId必掉的
func (ds *dropTemplateService) CheckSureDropBaoKu(dropId int32) (flag bool) {
	_, ok := ds.dropBaoKuGroupMap[dropId]
	if !ok {
		return
	}
	dropRateList, ok := ds.dropBaoKuRateGroupMap[dropId]
	if !ok {
		return
	}
	totalRate := int64(0)
	for _, rate := range dropRateList {
		totalRate += rate
	}

	if totalRate < droptypes.MAX_RATE {
		return
	}
	flag = true
	return
}

//获取掉落物品
func (ds *dropTemplateService) GetDropBaoKuItem(dropId int32) (itemId int32, num int32) {
	to := ds.GetDropBaoKuFromGroup(dropId)
	if to == nil {
		return
	}
	num = to.RandomNum()
	itemId = to.ItemId
	return
}

//获取掉落物品
func (ds *dropTemplateService) GetDropBaoKuListItems(dropList []int32) map[int32]int32 {
	if len(dropList) <= 0 {
		return nil
	}
	dropItemMap := make(map[int32]int32)
	for _, dropId := range dropList {
		itemId, num := ds.GetDropBaoKuItem(dropId)
		if itemId == 0 {
			continue
		}
		_, ok := dropItemMap[itemId]
		if !ok {
			dropItemMap[itemId] += num
		} else {
			dropItemMap[itemId] = num
		}
	}
	return dropItemMap
}

//获取掉落等级物品
func (ds *dropTemplateService) GetDropBaoKuItemLevel(dropId int32) (newData *DropItemData) {
	to := ds.GetDropBaoKuFromGroup(dropId)
	if to == nil {
		return
	}
	itemId := to.ItemId
	if itemId == 0 {
		return
	}

	num := to.RandomNum()
	level := to.RandomGoldEquipLevel()
	bind := to.GetBindType()
	star := to.RandomGoldEquipUpstarLevel()
	attrList, isRandom := to.RandomGoldEquipAttr()
	if !isRandom {
		itemTemp := item.GetItemService().GetItem(int(itemId))
		if itemTemp == nil {
			return
		}

		if itemTemp.IsGoldEquip() {
			attrList = itemTemp.GetGoldEquipTemplate().RandomGoldEquipAttr()
			isRandom = true
		}
	}
	expireType := inventorytypes.NewItemLimitTimeTypeNone
	expireTime := int64(0)
	itemGetTime := global.GetGame().GetTimeService().Now()
	newData = CreateItemDataWithPropertyData(itemId, num, level, bind, star, attrList, isRandom, expireType, expireTime, itemGetTime)
	return
}

//获取掉落等级物品
func (ds *dropTemplateService) GetDropBaoKuItemLevelWithIndex(dropId int32) (index int32, newData *DropItemData) {
	index, to := ds.GetDropBaoKuFromGroupWithIndex(dropId)
	if to == nil {
		return
	}
	itemId := to.ItemId
	if itemId == 0 {
		return
	}

	num := to.RandomNum()
	level := to.RandomGoldEquipLevel()
	bind := to.GetBindType()
	star := to.RandomGoldEquipUpstarLevel()
	attrList, isRandom := to.RandomGoldEquipAttr()
	if !isRandom {
		itemTemp := item.GetItemService().GetItem(int(itemId))
		if itemTemp == nil {
			return
		}

		if itemTemp.IsGoldEquip() {
			attrList = itemTemp.GetGoldEquipTemplate().RandomGoldEquipAttr()
			isRandom = true
		}
	}
	expireType := inventorytypes.NewItemLimitTimeTypeNone
	expireTime := int64(0)
	itemGetTime := global.GetGame().GetTimeService().Now()
	newData = CreateItemDataWithPropertyData(itemId, num, level, bind, star, attrList, isRandom, expireType, expireTime, itemGetTime)
	return
}

//获取掉落等级物品
func (ds *dropTemplateService) GetDropBaoKuListItemLevelList(dropList []int32) []*DropItemData {
	if len(dropList) <= 0 {
		return nil
	}
	var dropItemList []*DropItemData
	for _, dropId := range dropList {
		newData := ds.GetDropBaoKuItemLevel(dropId)
		if newData == nil {
			continue
		}
		dropItemList = append(dropItemList, newData)
	}
	return dropItemList
}

//end///////////////////////

var (
	once sync.Once
	cs   *dropTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &dropTemplateService{}
		err = cs.init()
	})
	return err
}

func GetDropTemplateService() DropTemplateService {
	return cs
}
