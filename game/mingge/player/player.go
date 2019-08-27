package player

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/mingge/dao"
	minggetemplate "fgame/fgame/game/mingge/template"
	minggetypes "fgame/fgame/game/mingge/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertytypes "fgame/fgame/game/property/types"
	propertyutils "fgame/fgame/game/property/utils"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/mathutils"
	"math"
)

//玩家命格管理器
type PlayerMingGeDataManager struct {
	p player.Player
	//玩家命盘
	mingGePanMap map[minggetypes.MingGeType]map[minggetypes.MingGeAllSubType]*PlayerMingGePanObject
	//玩家命盘祭炼
	mingGePanRefinedMap map[minggetypes.MingGeAllSubType]*PlayerMingGeRefinedObject
	//玩家命理
	mingLiMap        map[minggetypes.MingGongType]map[minggetypes.MingGongAllSubType]*PlayerMingLiObject
	mingGeBuchangObj *PlayerMingGeBuchangObject
	// 玩家命格
	mingGeObject *PlayerMingGeObject
}

func (m *PlayerMingGeDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerMingGeDataManager) Load() (err error) {
	err = m.loadMingPan()
	if err != nil {
		return
	}

	err = m.loadMingLi()
	if err != nil {
		return
	}

	err = m.loadMingPanRefined()
	if err != nil {
		return
	}
	err = m.loadBuchang()
	if err != nil {
		return
	}
	err = m.loadMingGe()
	if err != nil {
		return
	}
	return nil
}

func (m *PlayerMingGeDataManager) loadMingLi() (err error) {
	m.mingLiMap = make(map[minggetypes.MingGongType]map[minggetypes.MingGongAllSubType]*PlayerMingLiObject)
	mingLiList, err := dao.GetMingGeDao().GetMingLiList(m.p.GetId())
	if err != nil {
		return
	}

	for _, mingLiInfo := range mingLiList {
		obj := NewPlayerMingLiObject(m.p)
		obj.FromEntity(mingLiInfo)
		m.addMingLi(obj)
	}
	return
}

func (m *PlayerMingGeDataManager) loadMingPan() (err error) {
	m.mingGePanMap = make(map[minggetypes.MingGeType]map[minggetypes.MingGeAllSubType]*PlayerMingGePanObject)
	mingPanList, err := dao.GetMingGeDao().GetMingPanList(m.p.GetId())
	if err != nil {
		return
	}

	for _, mingPanInfo := range mingPanList {
		obj := NewPlayerMingGePanObject(m.p)
		obj.FromEntity(mingPanInfo)
		m.addMingGe(obj)
	}
	return
}

func (m *PlayerMingGeDataManager) loadMingPanRefined() (err error) {
	m.mingGePanRefinedMap = make(map[minggetypes.MingGeAllSubType]*PlayerMingGeRefinedObject)
	mingGePanRefinedList, err := dao.GetMingGeDao().GetMingGeRefinedList(m.p.GetId())
	if err != nil {
		return
	}

	for _, mingGePanRefined := range mingGePanRefinedList {
		obj := NewPlayerMingGeRefinedObject(m.p)
		obj.FromEntity(mingGePanRefined)
		m.addMingGePanRefined(obj)
	}
	return
}

func (m *PlayerMingGeDataManager) loadBuchang() (err error) {

	mingGeBuChangEntity, err := dao.GetMingGeDao().GetMingGeBuchang(m.p.GetId())
	if err != nil {
		return
	}
	if mingGeBuChangEntity == nil {
		m.initMingGeBuchang()
	} else {
		obj := NewPlayerMingGeBuchangObject(m.p)
		err = obj.FromEntity(mingGeBuChangEntity)
		if err != nil {
			return
		}
		m.mingGeBuchangObj = obj
	}

	return
}

func (m *PlayerMingGeDataManager) loadMingGe() (err error) {

	mingGeEntity, err := dao.GetMingGeDao().GetMingGeEntity(m.p.GetId())
	if err != nil {
		return
	}
	if mingGeEntity == nil {
		m.initMingGe()
	} else {
		obj := NewPlayerMingGeObject(m.p)
		err = obj.FromEntity(mingGeEntity)
		if err != nil {
			return
		}
		m.mingGeObject = obj
	}

	return
}

func (m *PlayerMingGeDataManager) initMingGeBuchang() {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	newObj := NewPlayerMingGeBuchangObject(m.p)
	newObj.id = id
	newObj.buchang = 0
	newObj.createTime = now
	newObj.SetModified()
	m.mingGeBuchangObj = newObj
	return
}

func (m *PlayerMingGeDataManager) initMingGe() {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	newObj := NewPlayerMingGeObject(m.p)
	newObj.id = id
	newObj.power = 0
	newObj.createTime = now
	newObj.SetModified()
	m.mingGeObject = newObj
	return
}

func (m *PlayerMingGeDataManager) addMingGe(obj *PlayerMingGePanObject) {
	mingGePanMap, ok := m.mingGePanMap[obj.GetMingPanType()]
	if !ok {
		mingGePanMap = make(map[minggetypes.MingGeAllSubType]*PlayerMingGePanObject)
		m.mingGePanMap[obj.GetMingPanType()] = mingGePanMap
	}
	mingGePanMap[obj.GetSubType()] = obj
}

func (m *PlayerMingGeDataManager) addMingLi(obj *PlayerMingLiObject) {
	mingLiMap, ok := m.mingLiMap[obj.GetMingGongType()]
	if !ok {
		mingLiMap = make(map[minggetypes.MingGongAllSubType]*PlayerMingLiObject)
		m.mingLiMap[obj.GetMingGongType()] = mingLiMap
	}
	mingLiMap[obj.GetSubType()] = obj
}

func (m *PlayerMingGeDataManager) addMingGePanRefined(obj *PlayerMingGeRefinedObject) {
	_, ok := m.mingGePanRefinedMap[obj.GetSubType()]
	if ok {
		return
	}
	m.mingGePanRefinedMap[obj.GetSubType()] = obj
}

//加载后
func (m *PlayerMingGeDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerMingGeDataManager) Heartbeat() {

}

func (m *PlayerMingGeDataManager) initMingLi(mingGongType minggetypes.MingGongType) {
	mingLiMap, ok := m.mingLiMap[mingGongType]
	if ok {
		return
	} else {
		mingLiMap = make(map[minggetypes.MingGongAllSubType]*PlayerMingLiObject)
		m.mingLiMap[mingGongType] = mingLiMap
	}

	for mingGongSubType := minggetypes.MingGongAllSubTypeMin; mingGongSubType <= minggetypes.MingGongAllSubTypeMax; mingGongSubType++ {
		_, ok := mingLiMap[mingGongSubType]
		if !ok {
			now := global.GetGame().GetTimeService().Now()
			obj := NewPlayerMingLiObject(m.p)
			id, _ := idutil.GetId()
			obj.id = id
			obj.mingGongType = mingGongType
			obj.subType = mingGongSubType
			obj.mingLiMap = make(map[minggetypes.MingLiSlotType]*MingLiInfo)
			obj.createTime = now
			obj.SetModified()
			mingLiMap[mingGongSubType] = obj
		}
	}
}

func (m *PlayerMingGeDataManager) initMingPanRefined(mingGeSubType minggetypes.MingGeAllSubType) (obj *PlayerMingGeRefinedObject) {
	obj, ok := m.mingGePanRefinedMap[mingGeSubType]
	if ok {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj = NewPlayerMingGeRefinedObject(m.p)
	id, _ := idutil.GetId()
	obj.id = id
	obj.subType = mingGeSubType
	obj.number = 0
	obj.star = 0
	obj.refinedNum = 0
	obj.refinedPro = 0
	obj.createTime = now
	obj.SetModified()
	m.mingGePanRefinedMap[mingGeSubType] = obj
	return
}

//获取命格盘信息
func (m *PlayerMingGeDataManager) GetMingGePanMap() map[minggetypes.MingGeType]map[minggetypes.MingGeAllSubType]*PlayerMingGePanObject {
	return m.mingGePanMap
}

//获取命盘祭炼
func (m *PlayerMingGeDataManager) GetMingGePanRefinedMap() map[minggetypes.MingGeAllSubType]*PlayerMingGeRefinedObject {
	return m.mingGePanRefinedMap
}

//获取命理信息
func (m *PlayerMingGeDataManager) GetMingLiMap() map[minggetypes.MingGongType]map[minggetypes.MingGongAllSubType]*PlayerMingLiObject {
	return m.mingLiMap
}

//获取命理根据类型
func (m *PlayerMingGeDataManager) GetMingLiByType(mingGongType minggetypes.MingGongType) map[minggetypes.MingGongAllSubType]*PlayerMingLiObject {
	mingLiMap, ok := m.mingLiMap[mingGongType]
	if !ok {
		return nil
	}
	return mingLiMap
}

//获取命盘根据类型
func (m *PlayerMingGeDataManager) GetMingGePanTypeMap(mingGeType minggetypes.MingGeType) map[minggetypes.MingGeAllSubType]*PlayerMingGePanObject {
	mingGeTypeMap, ok := m.mingGePanMap[mingGeType]
	if !ok {
		return nil
	}
	return mingGeTypeMap
}

//获取命盘祭炼根据类型
func (m *PlayerMingGeDataManager) GetMingGePanRefinedByType(mingGeSubType minggetypes.MingGeAllSubType) *PlayerMingGeRefinedObject {
	obj, ok := m.mingGePanRefinedMap[mingGeSubType]
	if !ok {
		return nil
	}
	return obj
}

//获取命理根据类型
func (m *PlayerMingGeDataManager) GetMingGeMingLiByType(mingGongType minggetypes.MingGongType) map[minggetypes.MingGongAllSubType]*PlayerMingLiObject {
	mingGongMap, ok := m.mingLiMap[mingGongType]
	if !ok {
		return nil
	}
	return mingGongMap
}

//获取命理对象
func (m *PlayerMingGeDataManager) GetMingGeMingLiByTypeAndSubType(mingGongType minggetypes.MingGongType, mingGongSubType minggetypes.MingGongAllSubType) *PlayerMingLiObject {
	mingGongMap := m.GetMingGeMingLiByType(mingGongType)
	if mingGongMap == nil {
		return nil
	}
	obj, ok := mingGongMap[mingGongSubType]
	if !ok {
		return nil
	}
	return obj
}

//获取槽位物品
func (m *PlayerMingGeDataManager) GetSlotItem(mingGeType minggetypes.MingGeType,
	mingGeSubType minggetypes.MingGeAllSubType, slot minggetypes.MingGeSlotType) (itemId int32) {

	mingGeMap, ok := m.mingGePanMap[mingGeType]
	if !ok {
		return
	}
	obj, ok := mingGeMap[mingGeSubType]
	if !ok {
		return
	}
	itemId = obj.itemMap[slot]
	return
}

// 获取命格战力
func (m *PlayerMingGeDataManager) GetPower() int64 {
	return m.mingGeObject.power
}

// 设置命格战力
func (m *PlayerMingGeDataManager) SetPower(power int64) {
	if power < 0 {
		return
	}
	if m.mingGeObject.power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.mingGeObject.power = power
	m.mingGeObject.updateTime = now
	m.mingGeObject.SetModified()
}

func (m *PlayerMingGeDataManager) slotMosaic(mingGeType minggetypes.MingGeType,
	mingGeSubType minggetypes.MingGeAllSubType, slot minggetypes.MingGeSlotType, itemId int32) {
	mingGeMap, ok := m.mingGePanMap[mingGeType]
	if !ok {
		mingGeMap = make(map[minggetypes.MingGeAllSubType]*PlayerMingGePanObject)
		m.mingGePanMap[mingGeType] = mingGeMap
	}
	now := global.GetGame().GetTimeService().Now()
	obj, ok := mingGeMap[mingGeSubType]
	if !ok {
		obj = NewPlayerMingGePanObject(m.p)
		id, _ := idutil.GetId()
		obj.id = id
		obj.mingPanType = mingGeType
		obj.subType = mingGeSubType
		obj.itemMap = make(map[minggetypes.MingGeSlotType]int32)
		obj.itemMap[slot] = itemId
		obj.createTime = now
		obj.SetModified()
		mingGeMap[mingGeSubType] = obj
		return
	}
	obj.itemMap[slot] = itemId
	obj.updateTime = now
	obj.SetModified()
}

//命盘镶嵌
func (m *PlayerMingGeDataManager) SlotMosaic(mingGeType minggetypes.MingGeType,
	mingGeSubType minggetypes.MingGeAllSubType, slot minggetypes.MingGeSlotType, itemId int32) (flag bool) {
	curItemId := m.GetSlotItem(mingGeType, mingGeSubType, slot)
	if curItemId != 0 {
		itemTemplate := item.GetItemService().GetItem(int(itemId))
		curItemTemplate := item.GetItemService().GetItem(int(curItemId))
		if curItemTemplate.GetQualityType() > itemTemplate.GetQualityType() ||
			(curItemTemplate.GetQualityType() == itemTemplate.GetQualityType() && curItemTemplate.TypeFlag2 >= itemTemplate.TypeFlag2) {
			return
		}
	}
	m.slotMosaic(mingGeType, mingGeSubType, slot, itemId)
	flag = true
	return
}

func (m *PlayerMingGeDataManager) slotUnload(mingGeType minggetypes.MingGeType,
	mingGeSubType minggetypes.MingGeAllSubType, slot minggetypes.MingGeSlotType) {
	mingGeMap, ok := m.mingGePanMap[mingGeType]
	if !ok {
		return
	}
	obj, ok := mingGeMap[mingGeSubType]
	if !ok {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	delete(obj.itemMap, slot)
	obj.updateTime = now
	obj.SetModified()
}

//命盘卸下
func (m *PlayerMingGeDataManager) SlotUnload(mingGeType minggetypes.MingGeType,
	mingGeSubType minggetypes.MingGeAllSubType, slot minggetypes.MingGeSlotType) (flag bool) {
	curItemId := m.GetSlotItem(mingGeType, mingGeSubType, slot)
	if curItemId == 0 {
		return
	}
	m.slotUnload(mingGeType, mingGeSubType, slot)
	flag = true
	return
}

//随机命盘祭炼
func (m *PlayerMingGeDataManager) RefinedRandom() (id int32, allFull bool) {

	idList := make([]int32, 0, 4)
	rateList := make([]int64, 0, 4)
	for curSubType := minggetypes.MingGeAllSubTypeMin; curSubType <= minggetypes.MingGeAllSubTypeMax; curSubType++ {
		randomNumber := int32(0)
		randomStar := int32(0)
		obj, ok := m.mingGePanRefinedMap[curSubType]
		if ok {
			randomNumber = obj.GetNumber()
			randomStar = obj.GetStar()
		}
		mingGePanTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingPanTemplate(curSubType, randomNumber, randomStar)
		if mingGePanTemplate == nil {
			continue
		}
		if mingGePanTemplate.NextId == 0 {
			continue
		}
		idList = append(idList, int32(mingGePanTemplate.TemplateId()))
		rateList = append(rateList, int64(mingGePanTemplate.Rate))
	}

	if len(rateList) == 0 {
		allFull = true
		return
	}

	index := mathutils.RandomWeights(rateList)
	return idList[index], false
}

func (m *PlayerMingGeDataManager) refinedFull(mingGeSubType minggetypes.MingGeAllSubType) (flag bool) {
	obj, ok := m.mingGePanRefinedMap[mingGeSubType]
	if !ok {
		panic("mingge:never reach here")
	}
	number := obj.GetNumber()
	star := obj.GetStar()
	mingGePanTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingPanTemplate(mingGeSubType, number, star)
	if mingGePanTemplate == nil {
		flag = true
		return
	}
	if mingGePanTemplate.NextId == 0 {
		flag = true
		return
	}
	return
}

//命盘祭炼
func (m *PlayerMingGeDataManager) Refined(mingGeSubType minggetypes.MingGeAllSubType, pro int32, sucess bool) (flag bool) {
	obj := m.GetMingGePanRefinedByType(mingGeSubType)
	if obj == nil {
		obj = m.initMingPanRefined(mingGeSubType)
	}
	flag = m.refinedFull(mingGeSubType)
	if flag {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		number := obj.GetNumber()
		star := obj.GetStar()
		mingGePanTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingPanTemplate(mingGeSubType, number, star)
		if mingGePanTemplate == nil {
			return
		}
		nextTemplate := mingGePanTemplate.GetNextMingPanTemplate()
		if nextTemplate == nil {
			return
		}
		nextNumber := nextTemplate.Number
		nextStar := nextTemplate.Star
		obj.number = nextNumber
		obj.star = nextStar
		obj.refinedNum = 0
		obj.refinedPro = pro
		obj.updateTime = now
		obj.SetModified()
	} else {
		obj.refinedNum += 1
		obj.refinedPro += pro
		obj.updateTime = now
		obj.SetModified()
	}
	return true
}

//判断命宫是否激活
func (m *PlayerMingGeDataManager) CheckMingGongActivate(level int32, zhuanShu int32) (mingGongTypeMap map[minggetypes.MingGongType]bool) {
	// if !m.p.IsFuncOpen(funcopentypes.FuncOpenTypeMingGong) {
	// 	return
	// }
	mingGongTypeMap = make(map[minggetypes.MingGongType]bool)
	for mingGongType := minggetypes.MingGongTypeGongMin; mingGongType <= minggetypes.MingGongTypeGongMax; mingGongType++ {
		mingLiMap := m.GetMingLiByType(mingGongType)
		if mingLiMap != nil {
			continue
		}
		mingGongTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingGongTemplate(mingGongType)

		needLevel := mingGongTemplate.NeedLevel
		needZhuanShu := mingGongTemplate.NeedZhuanShu
		needPower := int64(mingGongTemplate.NeedParentZhanLi)
		if level < needLevel {
			continue
		}
		if zhuanShu < needZhuanShu {
			continue
		}
		// 判断战力
		if mingGongTemplate.ParentId != 0 {
			parentMingGongType := mingGongTemplate.GetParentTemplate().GetMingGongType()
			battlePropertyMap := m.GetMingLiBattlePropertyMap(parentMingGongType)
			power := int64(0)
			if len(battlePropertyMap) != 0 {
				power = propertyutils.CulculateAllForce(battlePropertyMap)
			}
			if power < needPower {
				continue
			}
		}

		m.initMingLi(mingGongType)
		mingGongTypeMap[mingGongType] = true
	}
	return
}

//获取命理洗练所需的物品
func (m *PlayerMingGeDataManager) GetMingLiBaptizeNeedAllNum(mingGongType minggetypes.MingGongType, mingGongSubType minggetypes.MingGongAllSubType, slotTypeList []minggetypes.MingLiSlotType) (needItemMap map[int32]int32, flag bool) {
	needItemMap = make(map[int32]int32)
	obj := m.GetMingGeMingLiByTypeAndSubType(mingGongType, mingGongSubType)
	if obj == nil {
		return
	}
	for _, slotType := range slotTypeList {
		if !slotType.Vaild() {
			return
		}
		mingLiTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingLiTemplate(mingGongType, mingGongSubType)
		if mingLiTemplate == nil {
			return
		}
		baptizeNum := int32(1)
		mingLiInfo, ok := obj.mingLiMap[slotType]
		if ok {
			baptizeNum = mingLiInfo.GetTimes() + 1
		}

		useItemOne := float64(mingLiTemplate.UseItemOne)
		coefficientUse1 := mingLiTemplate.GetCoefficientUse1()
		coefficientUse2 := float64(mingLiTemplate.CoefficientUse2)
		value := math.Pow(float64(baptizeNum), coefficientUse1)
		needNum := int64(math.Ceil((useItemOne*value + coefficientUse2)))
		needItemMap[mingLiTemplate.UseItemId] += int32(needNum)
	}
	flag = true
	return
}

func (m *PlayerMingGeDataManager) getMingLiRateList(obj *PlayerMingLiObject, propertyPoolList []minggetypes.MingGePropertyType, slotType minggetypes.MingLiSlotType, rateList []int64) (curRateList []int64) {
	propertyLen := len(obj.GetMingLiMap())
	if propertyLen == 0 {
		return rateList
	}

	var curPropertyList []minggetypes.MingGePropertyType
	for curSlotType := minggetypes.MingLiSlotTypeMin; curSlotType <= minggetypes.MingLiSlotTypeMax; curSlotType++ {
		if curSlotType == slotType {
			continue
		}
		mingLiInfo, ok := obj.mingLiMap[curSlotType]
		if !ok {
			continue
		}
		curPropertyList = append(curPropertyList, mingLiInfo.MingGeProperty)
	}
	if len(curPropertyList) == 0 {
		return rateList
	}

	mingGongType := obj.GetMingGongType()
	mingGongSubType := obj.GetSubType()
	mingLiTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingLiTemplate(mingGongType, mingGongSubType)
	rate := int64(mingLiTemplate.ZhiDingRate1)
	if len(curPropertyList) != 1 {
		for i := 0; i < len(curPropertyList); i++ {
			for j := i + 1; j < len(curPropertyList); j++ {
				if curPropertyList[i] == curPropertyList[j] {
					rate = int64(mingLiTemplate.ZhiDingRate2)
				}
			}
		}
	}

	for index, propertyType := range propertyPoolList {
		isAppend := true
		for _, curPropertyType := range curPropertyList {
			if curPropertyType == propertyType {
				curRate := int64(math.Ceil(float64(rateList[index]*rate) / float64(common.MAX_RATE)))
				curRateList = append(curRateList, curRate)
				isAppend = false
				break
			}
		}
		if isAppend {
			curRateList = append(curRateList, rateList[index])
		}
	}
	return
}

func (m *PlayerMingGeDataManager) randomMingLiProperty(obj *PlayerMingLiObject, propertyPoolList []minggetypes.MingGePropertyType, slotType minggetypes.MingLiSlotType, rateList []int64, xilianLimitCount int32) (flag bool) {
	now := global.GetGame().GetTimeService().Now()
	curRateList := m.getMingLiRateList(obj, propertyPoolList, slotType, rateList)
	index := mathutils.RandomWeights(curRateList)
	if index == -1 {
		return
	}
	propertyType := propertyPoolList[index]
	mingLiInfo, ok := obj.mingLiMap[slotType]
	if !ok {
		mingLiInfo = newMingLiInfo(slotType, propertyType, 0)
		obj.mingLiMap[slotType] = mingLiInfo
	}
	mingLiInfo.MingGeProperty = propertyType
	if mingLiInfo.GetTimes() < xilianLimitCount {
		mingLiInfo.Times++
	}
	obj.updateTime = now
	obj.SetModified()
	flag = true
	return
}

//命理洗炼
func (m *PlayerMingGeDataManager) MingLiBaptize(mingGongType minggetypes.MingGongType, mingGongSubType minggetypes.MingGongAllSubType, slotTypeList []minggetypes.MingLiSlotType) (mingGongTypeMap map[minggetypes.MingGongType]bool, flag bool) {
	obj := m.GetMingGeMingLiByTypeAndSubType(mingGongType, mingGongSubType)
	if obj == nil {
		return
	}
	for _, slotType := range slotTypeList {
		if !slotType.Vaild() {
			return
		}
	}
	mingLiTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingLiTemplate(mingGongType, mingGongSubType)
	if mingLiTemplate == nil {
		return
	}
	propertyPoolList := mingLiTemplate.GetPropertyPoolList()
	if len(propertyPoolList) == 0 {
		return
	}

	rateList, _ := minggetemplate.GetMingGeTemplateService().GetMingGeDanBeiRateList(propertyPoolList)
	if len(rateList) != len(propertyPoolList) {
		return
	}
	for _, slotType := range slotTypeList {
		sucess := m.randomMingLiProperty(obj, propertyPoolList, slotType, rateList, mingLiTemplate.XilianLimitCount)
		if !sucess {
			return
		}
	}
	//判断命宫是否激活
	level := m.p.GetLevel()
	zhuanShu := m.p.GetZhuanSheng()
	mingGongTypeMap = m.CheckMingGongActivate(level, zhuanShu)
	flag = true
	return
}

func (m *PlayerMingGeDataManager) GetAllPropertyIsSame(mingGongType minggetypes.MingGongType, mingGongSubType minggetypes.MingGongAllSubType) (isSameAll bool) {
	isSameNum := int32(0)
	var propertyTypeList []minggetypes.MingGePropertyType
	obj := m.GetMingGeMingLiByTypeAndSubType(mingGongType, mingGongSubType)
	if obj == nil {
		return
	}
	mingLiMap := obj.GetMingLiMap()
	for _, mingLiInfo := range mingLiMap {
		propertyTypeList = append(propertyTypeList, mingLiInfo.GetMingGeProperty())
	}
	if len(propertyTypeList) == int(minggetypes.MingLiSlotTypeMax) {
		var lastPropertyType minggetypes.MingGePropertyType
		for index, propertyType := range propertyTypeList {
			if index == 0 {
				lastPropertyType = propertyType
				isSameNum++
				continue
			}
			if lastPropertyType != propertyType {
				break
			}
			isSameNum++
		}
		if isSameNum == int32(minggetypes.MingLiSlotTypeMax) {
			isSameAll = true
		}
	}
	return
}

//获取属性
func (m *PlayerMingGeDataManager) GetMingLiBattlePropertyMap(mingGongType minggetypes.MingGongType) (battlePropertyMap map[propertytypes.BattlePropertyType]int64) {
	mingGongSubTypeMap, ok := m.mingLiMap[mingGongType]
	if !ok {
		return
	}

	battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	for mingGongSubType, obj := range mingGongSubTypeMap {
		tempBattlePropertyMap := make(map[propertytypes.BattlePropertyType]float64)
		mingLiTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingLiTemplate(mingGongType, mingGongSubType)
		if mingLiTemplate == nil {
			continue
		}
		if len(obj.GetMingLiMap()) == 0 {
			continue
		}
		zhiDingPropertyType := mingLiTemplate.GetZhiDingPropertyType()
		isSameAll := m.GetAllPropertyIsSame(mingGongType, mingGongSubType)

		for _, mingInfo := range obj.GetMingLiMap() {
			mingGeDanBeiTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeDanBeiTemplate(mingInfo.GetMingGeProperty())
			attr := float64(mingGeDanBeiTemplate.Attr)
			times := mingInfo.GetTimes()
			attrOne := float64(mingLiTemplate.AttrOne)
			coefficientAttr1 := mingLiTemplate.GetCoefficientAttr1()
			coefficientAttr2 := mingLiTemplate.CoefficientAttr2
			value := math.Pow(float64(times), coefficientAttr1)
			normalValue := (attrOne*value + coefficientAttr2) * attr
			if mingInfo.GetMingGeProperty() == zhiDingPropertyType {
				coefficientZhiDing := int64(mingLiTemplate.CoefficientZhiDing)
				normalValue = float64(normalValue*float64(coefficientZhiDing)) / float64(common.MAX_RATE)
			}
			tempBattlePropertyMap[mingInfo.GetMingGeProperty().GetPropertyType()] += normalValue
		}
		if isSameAll {
			shouYuPercent := int64(mingLiTemplate.ShouYiPercent)
			for typ, value := range tempBattlePropertyMap {
				totalValue := float64(value*float64(shouYuPercent)) / float64(common.MAX_RATE)
				tempBattlePropertyMap[typ] = totalValue
			}
		}
		for typ, val := range tempBattlePropertyMap {
			battlePropertyMap[typ] += int64(math.Ceil(val))
		}
	}
	return
}

func (m *PlayerMingGeDataManager) IsBuchang() bool {
	return m.mingGeBuchangObj.IsBuchang()
}

func (m *PlayerMingGeDataManager) Buchang() bool {
	if m.mingGeBuchangObj.IsBuchang() {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	m.mingGeBuchangObj.buchang = 1
	m.mingGeBuchangObj.updateTime = now
	m.mingGeBuchangObj.SetModified()
	return true
}

func (m *PlayerMingGeDataManager) GetBuchangList() []int32 {
	returnList := make([]int32, 0, 8)
	for mingGongType, mingGongSubTypeMap := range m.mingLiMap {
		num := int32(0)
		for mingGongSubType, obj := range mingGongSubTypeMap {

			mingLiTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingLiTemplate(mingGongType, mingGongSubType)
			if mingLiTemplate == nil {
				continue
			}
			if len(obj.GetMingLiMap()) == 0 {
				continue
			}

			isSameAll := m.GetAllPropertyIsSame(mingGongType, mingGongSubType)

			if isSameAll {
				num += 1
			}
		}
		returnList = append(returnList, num)
	}
	return returnList
}

func (m *PlayerMingGeDataManager) GmResetBuchang() {
	now := global.GetGame().GetTimeService().Now()
	m.mingGeBuchangObj.buchang = 0
	m.mingGeBuchangObj.updateTime = now
	m.mingGeBuchangObj.SetModified()
}

func CreatePlayerMingGeDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerMingGeDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerMingGeDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerMingGeDataManager))
}
