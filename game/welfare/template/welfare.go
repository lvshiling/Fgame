package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	groupcollectenum "fgame/fgame/game/welfare/group/collect/enum"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/mathutils"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"sync"
)

type WelfareTemplateService interface {
	//首冲配置
	GetFirstCharge(role playertypes.RoleType, sex playertypes.SexType) *gametemplate.FirstChargeTemplate
	//礼包配置
	GetCodeGift(code int32) *gametemplate.GiftCodeTemplate
	//活动配置
	GetOpenActivityTemplate(tempId int32) *gametemplate.OpenserverActivityTemplate
	GetOpenActivityTemplateByGroup(groupId int32) map[int32]*gametemplate.OpenserverActivityTemplate
	GetOpenActivityGroupTemplateInterface(groupId int32) GroupTemplateI

	//七日冲刺活动最大过期时间
	GetRankMaxEndTime(openTime int64) (int64, []*gametemplate.OpenserverTimeTemplate)
	//活动时间内配置
	GetOpenActivityTimeTemplate(groupId int32) *gametemplate.OpenserverTimeTemplate
	GetOpenActivityTimeTemplateById(id int32) *gametemplate.OpenserverTimeTemplate
	GetOpenActivityTimeTemplateByType(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) []*gametemplate.OpenserverTimeTemplate
	GetOpenActivityTemplateByFuncType(funcType funcopentypes.FuncOpenType) []*gametemplate.OpenserverTimeTemplate
	//排行榜活动时间配置
	GetRankActivityTimeTemplate() map[welfaretypes.OpenActivityRankSubType][]*gametemplate.OpenserverTimeTemplate
	//所有活动时间配置
	GetAllActivityTimeTemplate() map[int32]*gametemplate.OpenserverTimeTemplate
	//抽奖配置
	GetLuckDrewTemplate(groupId int32) *gametemplate.LuckyDrewTemplate
	//根据抽奖等级配置
	GetLuckDrewTemplateByArg(groupId int32, level int32) *gametemplate.LuckyDrewTemplate
	//折扣配置
	GetDiscountTemplate(discountId int32) *gametemplate.DiscountTemplate
	GetDiscountTemplateByDayGroup(dayGroup int32) []*gametemplate.DiscountTemplate
	// 转生大礼包组配置
	GetDiscountZhuanShengGroupTemplate(groupId int32) *ZhuanSengGiftGroupTemplate
	//特殊掉落配置（活动怪）
	GetGroupBiologyDropId(groupId, dropNum int32) int32
	//元宝拉霸配置
	GetGoldLabaTemplate(groupId, times int32) *gametemplate.GoldLaBaTemplate
	GetRandomLaBaTemplate(groupId int32) *gametemplate.GoldLaBaTemplate
	//次数奖励配置
	GetTimesRewTemplate(id int32) *gametemplate.TimesRewTemplate
	GetTimesRewTemplateByGorup(groupId int32) []*gametemplate.TimesRewTemplate
	//炼制配置
	GetMadeTemplate(groupId, level int32) *gametemplate.MadeTemplate
	// 打折礼包组配置
	GetDiscountBargainShopGroupTemplate(groupId int32) *BargainShopGroupTemplate
	// 名人普配置
	GetFamousTemplate(groupId int32) *gametemplate.FamousTemplate
	// 循环活动-当前循环日
	GetCurActivityXunHuanDay(openServerTime int64) int32
	GetRandomActivityXunHuanArrGroup(openServerTime int64) int32
	GetActivityXunHuanTemplate(openServerTime int64, arrIndex, cycDay int32) *gametemplate.OpenActivityXunHuanTemplate
	// 摸金配置
	GetCollectPokerTemplate(groupId int32, pokerType groupcollectenum.PokerType) *gametemplate.ChouJiangPokerTemplate
	// 转生礼包折扣配置
	GetZhuanShengCircleBargainTemplate(class, buyNum int32) *gametemplate.CircleBargainTemplate
	//合服循环活动
	GetActivityMergeXunHuanTemplate(openTime, mergeTime int64) *gametemplate.OpenActivityMergeXunHuanTemplate
}

type welfareTemplateService struct {
	//首冲奖励配置
	firstChargeMap map[playertypes.RoleType]map[playertypes.SexType]*gametemplate.FirstChargeTemplate
	//开服活动奖励配置
	openActivityByIdMap    map[int32]*gametemplate.OpenserverActivityTemplate
	openActivityByGroupMap map[int32]map[int32]*gametemplate.OpenserverActivityTemplate
	groupTempMap           map[int32]GroupTemplateI
	//开服活动时间配置
	openActivityTimeMap        map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType][]*gametemplate.OpenserverTimeTemplate
	openActivityTimeByGroupMap map[int32]*gametemplate.OpenserverTimeTemplate
	openActivityTimeIdMap      map[int32]*gametemplate.OpenserverTimeTemplate
	openActivityTimeFuncIdMap  map[funcopentypes.FuncOpenType][]*gametemplate.OpenserverTimeTemplate
	//礼包配置
	giftMap map[int32]*gametemplate.GiftCodeTemplate
	//幸运抽奖配置
	luckyDrewMap map[int32]map[int32]*gametemplate.LuckyDrewTemplate
	//折扣配置
	discountMap        map[int32]*gametemplate.DiscountTemplate
	discountMapOfGroup map[int32][]*gametemplate.DiscountTemplate
	//转生大礼包
	zhuanShengGroupMap map[int32]*ZhuanSengGiftGroupTemplate
	//活动怪掉落配置
	groupBiologyDropMap map[int32][]*gametemplate.TeShuDropTemplate
	//拉霸配置
	goldLabaMap map[int32]map[int32]*gametemplate.GoldLaBaTemplate
	//次数奖励配置
	timesRewMap        map[int32]*gametemplate.TimesRewTemplate
	timesRewMapOfGroup map[int32][]*gametemplate.TimesRewTemplate
	//炼制配置
	madeGroupMap map[int32][]*gametemplate.MadeTemplate
	//转生大礼包
	bargainGroupMap map[int32]*BargainShopGroupTemplate
	//名人普
	fameMap map[int32]*gametemplate.FamousTemplate
	//循环活动
	xunhuanMap         map[int32]map[int32]map[int32]*gametemplate.OpenActivityXunHuanTemplate
	xunhuanTypeMap     map[int32]*gametemplate.OpenActivityXunHuanTimeTemplate
	circleGroupListMap map[int32][]int32
	hadChooseGroup     int32
	lastChooseGroup    int32
	//摸金-卡牌收集
	pokerTempMap map[int32]map[groupcollectenum.PokerType]*gametemplate.ChouJiangPokerTemplate
	//转生礼包购买折扣
	zhuanshengBargainMap map[int32][]*gametemplate.CircleBargainTemplate
	//合服循环活动
	mergeXunHuanMap    map[int32]*gametemplate.OpenActivityMergeXunHuanTemplate
	maxMergeXunHuanDay int32
}

func (st *welfareTemplateService) init() (err error) {
	// 首冲奖励配置
	st.firstChargeMap = make(map[playertypes.RoleType]map[playertypes.SexType]*gametemplate.FirstChargeTemplate)
	firstTempMap := template.GetTemplateService().GetAll((*gametemplate.FirstChargeTemplate)(nil))
	for _, temp := range firstTempMap {
		firstTemp, _ := temp.(*gametemplate.FirstChargeTemplate)
		firstChargeSexMap, ok := st.firstChargeMap[firstTemp.GetRole()]
		if !ok {
			firstChargeSexMap = make(map[playertypes.SexType]*gametemplate.FirstChargeTemplate)
			st.firstChargeMap[firstTemp.GetRole()] = firstChargeSexMap
		}
		firstChargeSexMap[firstTemp.GetSex()] = firstTemp
	}

	//礼包配置
	st.giftMap = make(map[int32]*gametemplate.GiftCodeTemplate)
	giftTempMap := template.GetTemplateService().GetAll((*gametemplate.GiftCodeTemplate)(nil))
	for _, temp := range giftTempMap {
		giftTemp, _ := temp.(*gametemplate.GiftCodeTemplate)
		st.giftMap[int32(temp.TemplateId())] = giftTemp
	}

	//开服活动时间配置
	st.openActivityTimeMap = make(map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType][]*gametemplate.OpenserverTimeTemplate)
	st.openActivityTimeIdMap = make(map[int32]*gametemplate.OpenserverTimeTemplate)
	st.openActivityTimeByGroupMap = make(map[int32]*gametemplate.OpenserverTimeTemplate)
	st.openActivityTimeFuncIdMap = make(map[funcopentypes.FuncOpenType][]*gametemplate.OpenserverTimeTemplate)
	openTimeTempMap := template.GetTemplateService().GetAll((*gametemplate.OpenserverTimeTemplate)(nil))
	for _, temp := range openTimeTempMap {
		openTimeTemp, _ := temp.(*gametemplate.OpenserverTimeTemplate)
		_, ok := st.openActivityTimeByGroupMap[openTimeTemp.Group]
		if ok {
			return fmt.Errorf("openserver_time repeat gourpId:%d", openTimeTemp.Group)
		}
		st.openActivityTimeByGroupMap[openTimeTemp.Group] = openTimeTemp
		st.openActivityTimeIdMap[int32(openTimeTemp.Id)] = openTimeTemp
		st.openActivityTimeFuncIdMap[openTimeTemp.GetOpenFuncType()] = append(st.openActivityTimeFuncIdMap[openTimeTemp.GetOpenFuncType()], openTimeTemp)

		typ := openTimeTemp.GetOpenType()
		subType := openTimeTemp.GetOpenSubType()
		openSubMap, ok := st.openActivityTimeMap[typ]
		if !ok {
			openSubMap = make(map[welfaretypes.OpenActivitySubType][]*gametemplate.OpenserverTimeTemplate)
			st.openActivityTimeMap[typ] = openSubMap
		}
		openSubMap[subType] = append(openSubMap[subType], openTimeTemp)

	}

	// 开服活动奖励配置
	st.openActivityByIdMap = make(map[int32]*gametemplate.OpenserverActivityTemplate)
	st.openActivityByGroupMap = make(map[int32]map[int32]*gametemplate.OpenserverActivityTemplate)
	st.groupTempMap = make(map[int32]GroupTemplateI)
	openTempMap := template.GetTemplateService().GetAll((*gametemplate.OpenserverActivityTemplate)(nil))
	for _, temp := range openTempMap {
		openTemp, _ := temp.(*gametemplate.OpenserverActivityTemplate)
		//group
		groupMap, ok := st.openActivityByGroupMap[openTemp.Group]
		if !ok {
			groupMap = make(map[int32]*gametemplate.OpenserverActivityTemplate)
			st.openActivityByGroupMap[openTemp.Group] = groupMap
		}
		groupMap[int32(openTemp.TemplateId())] = openTemp

		// timeTemp
		timeTemp, ok := st.openActivityTimeByGroupMap[openTemp.Group]
		if !ok {
			return fmt.Errorf("活动没有时间配置,group:%d", openTemp.Group)
		}

		//tempId
		st.openActivityByIdMap[int32(openTemp.TemplateId())] = openTemp

		// groupTempMap
		groupInterface, ok := st.groupTempMap[openTemp.Group]
		if !ok {
			base := CreateGroupTemplateBase(timeTemp)
			groupInterface = CreateGroupTemplateI(timeTemp.GetOpenType(), timeTemp.GetOpenSubType(), base)
			if groupInterface == nil {
				continue
			}
			st.groupTempMap[openTemp.Group] = groupInterface
		}
		groupInterface.AddOpenTemp(openTemp)
	}

	// 幸运转盘
	st.luckyDrewMap = make(map[int32]map[int32]*gametemplate.LuckyDrewTemplate)
	drewTempMap := template.GetTemplateService().GetAll((*gametemplate.LuckyDrewTemplate)(nil))
	for _, temp := range drewTempMap {
		drewTemp, _ := temp.(*gametemplate.LuckyDrewTemplate)
		tempM, ok := st.luckyDrewMap[drewTemp.ChouJiangId]
		if !ok {
			tempM = make(map[int32]*gametemplate.LuckyDrewTemplate)
			st.luckyDrewMap[drewTemp.ChouJiangId] = tempM
		}
		tempM[drewTemp.Level] = drewTemp
	}

	// 折扣配置
	st.discountMap = make(map[int32]*gametemplate.DiscountTemplate)
	st.discountMapOfGroup = make(map[int32][]*gametemplate.DiscountTemplate)
	discountTempMap := template.GetTemplateService().GetAll((*gametemplate.DiscountTemplate)(nil))
	for _, temp := range discountTempMap {
		discountTemp, _ := temp.(*gametemplate.DiscountTemplate)
		st.discountMap[int32(discountTemp.TemplateId())] = discountTemp

		//group
		st.discountMapOfGroup[discountTemp.DayGroup] = append(st.discountMapOfGroup[discountTemp.DayGroup], discountTemp)
	}

	// 转生大礼包配置
	st.zhuanShengGroupMap = make(map[int32]*ZhuanSengGiftGroupTemplate)
	zhuanShengTempMap := template.GetTemplateService().GetAll((*gametemplate.ZhuanShengGiftTemplate)(nil))
	for _, temp := range zhuanShengTempMap {
		zhuanShengTemp, _ := temp.(*gametemplate.ZhuanShengGiftTemplate)

		zhuanShengGroup, ok := st.zhuanShengGroupMap[zhuanShengTemp.Group]
		if !ok {
			zhuanShengGroup = CreateZhuanShengGroupTemplate()
			st.zhuanShengGroupMap[zhuanShengTemp.Group] = zhuanShengGroup
		}
		zhuanShengGroup.AddTemplate(zhuanShengTemp)
	}

	//活动怪物掉落配置
	st.groupBiologyDropMap = make(map[int32][]*gametemplate.TeShuDropTemplate)
	dropTempMap := template.GetTemplateService().GetAll((*gametemplate.TeShuDropTemplate)(nil))
	for _, temp := range dropTempMap {
		dropTemp := temp.(*gametemplate.TeShuDropTemplate)
		st.groupBiologyDropMap[dropTemp.GroupId] = append(st.groupBiologyDropMap[dropTemp.GroupId], dropTemp)
	}

	//元宝拉霸配置
	st.goldLabaMap = make(map[int32]map[int32]*gametemplate.GoldLaBaTemplate)
	labaTempMap := template.GetTemplateService().GetAll((*gametemplate.GoldLaBaTemplate)(nil))
	for _, temp := range labaTempMap {
		labaTemp := temp.(*gametemplate.GoldLaBaTemplate)
		timesMap, ok := st.goldLabaMap[labaTemp.GroupId]
		if !ok {
			timesMap = make(map[int32]*gametemplate.GoldLaBaTemplate)
			st.goldLabaMap[labaTemp.GroupId] = timesMap
		}
		timesMap[labaTemp.Times] = labaTemp
	}

	//次数奖励配置
	st.timesRewMap = make(map[int32]*gametemplate.TimesRewTemplate)
	st.timesRewMapOfGroup = make(map[int32][]*gametemplate.TimesRewTemplate)
	timesRewTempMap := template.GetTemplateService().GetAll((*gametemplate.TimesRewTemplate)(nil))
	for _, temp := range timesRewTempMap {
		timesRewTemp := temp.(*gametemplate.TimesRewTemplate)
		st.timesRewMap[int32(timesRewTemp.Id)] = timesRewTemp

		//group
		st.timesRewMapOfGroup[timesRewTemp.Group] = append(st.timesRewMapOfGroup[timesRewTemp.Group], timesRewTemp)
	}

	//炼制配置
	st.madeGroupMap = make(map[int32][]*gametemplate.MadeTemplate)
	madeTempMap := template.GetTemplateService().GetAll((*gametemplate.MadeTemplate)(nil))
	for _, temp := range madeTempMap {
		madeTemp := temp.(*gametemplate.MadeTemplate)
		st.madeGroupMap[madeTemp.GroupId] = append(st.madeGroupMap[madeTemp.GroupId], madeTemp)
	}

	// 打折礼包配置
	st.bargainGroupMap = make(map[int32]*BargainShopGroupTemplate)
	bargainShopTempMap := template.GetTemplateService().GetAll((*gametemplate.BargainShopTemplate)(nil))
	for _, temp := range bargainShopTempMap {
		bargainShopTemp, _ := temp.(*gametemplate.BargainShopTemplate)

		shopGroup, ok := st.bargainGroupMap[bargainShopTemp.Group]
		if !ok {
			shopGroup = CreateBargainGroupTemplate()
			st.bargainGroupMap[bargainShopTemp.Group] = shopGroup
		}
		shopGroup.AddTemplate(bargainShopTemp)
	}
	// 名人普配置
	st.fameMap = make(map[int32]*gametemplate.FamousTemplate)
	fameTempMap := template.GetTemplateService().GetAll((*gametemplate.FamousTemplate)(nil))
	for _, temp := range fameTempMap {
		fameTemp, _ := temp.(*gametemplate.FamousTemplate)
		st.fameMap[fameTemp.GroupId] = fameTemp
	}

	//循环活动时间类型配置
	st.xunhuanTypeMap = make(map[int32]*gametemplate.OpenActivityXunHuanTimeTemplate)
	xunhuanTimeTempMap := template.GetTemplateService().GetAll((*gametemplate.OpenActivityXunHuanTimeTemplate)(nil))
	for _, temp := range xunhuanTimeTempMap {
		timeTypeTemp, _ := temp.(*gametemplate.OpenActivityXunHuanTimeTemplate)
		st.xunhuanTypeMap[timeTypeTemp.TimeType] = timeTypeTemp
	}
	// 循环活动配置
	st.xunhuanMap = make(map[int32]map[int32]map[int32]*gametemplate.OpenActivityXunHuanTemplate)
	st.circleGroupListMap = make(map[int32][]int32)
	xunhuanTempMap := template.GetTemplateService().GetAll((*gametemplate.OpenActivityXunHuanTemplate)(nil))
	for _, temp := range xunhuanTempMap {
		xunhuanTemp, _ := temp.(*gametemplate.OpenActivityXunHuanTemplate)
		subMap, ok := st.xunhuanMap[xunhuanTemp.TimeType]
		if !ok {
			subMap = make(map[int32]map[int32]*gametemplate.OpenActivityXunHuanTemplate)
			st.xunhuanMap[xunhuanTemp.TimeType] = subMap
		}
		subOfSubMap, ok := subMap[xunhuanTemp.ArrIndex]
		if !ok {
			subOfSubMap = make(map[int32]*gametemplate.OpenActivityXunHuanTemplate)
			subMap[xunhuanTemp.ArrIndex] = subOfSubMap
		}
		subOfSubMap[xunhuanTemp.CycleDay] = xunhuanTemp

		//校验groupId
		for _, groupId := range xunhuanTemp.GetGroupIdList() {
			timeTemp, ok := st.openActivityTimeByGroupMap[groupId]
			if !ok {
				return fmt.Errorf("循环活动配置错误，[%d] invalid", groupId)
			}

			if !timeTemp.IsXunHuan() {
				return fmt.Errorf("不是循环活动，[%d] invalid", groupId)
			}
		}
		//校验timeType
		_, ok = st.xunhuanTypeMap[xunhuanTemp.TimeType]
		if !ok {
			return fmt.Errorf("循环活动时间类型配置错误，[%d] invalid", xunhuanTemp.TimeType)
		}

		//随机组列表
		circleGroupList := st.circleGroupListMap[xunhuanTemp.TimeType]
		if utils.ContainInt32(circleGroupList, xunhuanTemp.ArrIndex) {
			continue
		}
		st.circleGroupListMap[xunhuanTemp.TimeType] = append(st.circleGroupListMap[xunhuanTemp.TimeType], xunhuanTemp.ArrIndex)
	}

	// 循环天数校验
	for _, subMap := range st.xunhuanMap {
		for _, subOfSubMap := range subMap {
			curXunHuanDayNum := len(subOfSubMap)
			for _, subOfSubMap := range subMap {
				if curXunHuanDayNum != len(subOfSubMap) {
					return fmt.Errorf("配置的循环天数不一致")
				}
			}
		}
	}

	// 摸金-卡牌模板
	st.pokerTempMap = make(map[int32]map[groupcollectenum.PokerType]*gametemplate.ChouJiangPokerTemplate)
	pokerTempMap := template.GetTemplateService().GetAll((*gametemplate.ChouJiangPokerTemplate)(nil))
	for _, temp := range pokerTempMap {
		pokerTemp, _ := temp.(*gametemplate.ChouJiangPokerTemplate)

		subMap, ok := st.pokerTempMap[pokerTemp.GroupId]
		if !ok {
			subMap = make(map[groupcollectenum.PokerType]*gametemplate.ChouJiangPokerTemplate)
			st.pokerTempMap[pokerTemp.GroupId] = subMap
		}
		_, ok = subMap[pokerTemp.GetPokerType()]
		if ok {
			return fmt.Errorf("卡牌奖励只有一条，重复配置;pokerType:%d", pokerTemp.GetPokerType())
		}

		subMap[pokerTemp.GetPokerType()] = pokerTemp
	}

	// 转生礼包折扣配置
	st.zhuanshengBargainMap = make(map[int32][]*gametemplate.CircleBargainTemplate)
	circleBargainTempMap := template.GetTemplateService().GetAll((*gametemplate.CircleBargainTemplate)(nil))
	for _, temp := range circleBargainTempMap {
		bargainTemp, _ := temp.(*gametemplate.CircleBargainTemplate)
		st.zhuanshengBargainMap[bargainTemp.Type] = append(st.zhuanshengBargainMap[bargainTemp.Type], bargainTemp)
	}

	// 合服循环活动配置
	st.mergeXunHuanMap = make(map[int32]*gametemplate.OpenActivityMergeXunHuanTemplate)
	mergeXunhuanTempMap := template.GetTemplateService().GetAll((*gametemplate.OpenActivityMergeXunHuanTemplate)(nil))
	for _, temp := range mergeXunhuanTempMap {
		xunhuanTemp, _ := temp.(*gametemplate.OpenActivityMergeXunHuanTemplate)
		st.mergeXunHuanMap[xunhuanTemp.CycleDay] = xunhuanTemp

		//最大循环日
		if st.maxMergeXunHuanDay < xunhuanTemp.CycleDay {
			st.maxMergeXunHuanDay = xunhuanTemp.CycleDay
		}

		//校验groupId
		for _, groupId := range xunhuanTemp.GetGroupIdList() {
			timeTemp, ok := st.openActivityTimeByGroupMap[groupId]
			if !ok {
				return fmt.Errorf("循环活动配置错误，[%d] invalid", groupId)
			}

			if !timeTemp.IsMergeXunHuan() {
				return fmt.Errorf("不是合服循环活动，[%d] invalid", groupId)
			}
		}
	}

	//校验合服循环活动时间不超过最大循环日
	constantTempMap := template.GetTemplateService().GetAll((*gametemplate.ConstantTemplate)(nil))
	for _, to := range constantTempMap {
		constantTemp, _ := to.(*gametemplate.ConstantTemplate)
		if constantTemp.GetConstantType() != constanttypes.ConstantTypeMergeXunHuanKeepDay {
			continue
		}
		if int32(constantTemp.Value) > st.maxMergeXunHuanDay {
			return fmt.Errorf("合服循环活动持续时间超过最大循环日")
		}
	}

	//有需要特殊处理的在这里(需要放在最后面)
	for _, groupInterface := range st.groupTempMap {
		err = groupInterface.Init()
		if err != nil {
			return
		}
	}

	return
}

func (st *welfareTemplateService) GetFirstCharge(role playertypes.RoleType, sex playertypes.SexType) *gametemplate.FirstChargeTemplate {
	firstChargeSexMap, ok := st.firstChargeMap[role]
	if !ok {
		return nil
	}
	firstCharge, ok := firstChargeSexMap[sex]
	if !ok {
		return nil
	}

	return firstCharge
}

func (st *welfareTemplateService) GetCodeGift(code int32) *gametemplate.GiftCodeTemplate {
	temp, ok := st.giftMap[code]
	if !ok {
		return nil
	}
	return temp
}

func (st *welfareTemplateService) GetOpenActivityTemplate(tempId int32) *gametemplate.OpenserverActivityTemplate {
	temp, ok := st.openActivityByIdMap[tempId]
	if !ok {
		return nil
	}

	return temp
}

func (st *welfareTemplateService) GetOpenActivityTemplateByGroup(groupId int32) map[int32]*gametemplate.OpenserverActivityTemplate {
	tempList, ok := st.openActivityByGroupMap[groupId]
	if !ok {
		return nil
	}

	return tempList
}

func (st *welfareTemplateService) GetOpenActivityGroupTemplateInterface(groupId int32) GroupTemplateI {
	groupTemp, ok := st.groupTempMap[groupId]
	if !ok {
		return nil
	}

	return groupTemp
}

func (st *welfareTemplateService) GetDiscountZhuanShengGroupTemplate(groupId int32) *ZhuanSengGiftGroupTemplate {
	groupTemp, ok := st.zhuanShengGroupMap[groupId]
	if !ok {
		return nil
	}

	return groupTemp
}

func (st *welfareTemplateService) GetDiscountBargainShopGroupTemplate(groupId int32) *BargainShopGroupTemplate {
	groupTemp, ok := st.bargainGroupMap[groupId]
	if !ok {
		return nil
	}

	return groupTemp
}

func (st *welfareTemplateService) GetFamousTemplate(groupId int32) *gametemplate.FamousTemplate {
	temp, ok := st.fameMap[groupId]
	if !ok {
		return nil
	}

	return temp
}

func (st *welfareTemplateService) GetCurActivityXunHuanDay(openServerTime int64) int32 {
	cycle := int32(st.getXunHuanDayNum())
	diffDay := st.getOpenDiff(openServerTime)
	cycleDay := diffDay % cycle
	return cycleDay + 1
}

// 随机循环活动组
func (st *welfareTemplateService) GetRandomActivityXunHuanArrGroup(openServerTime int64) int32 {
	timeTypeTemp := st.GetActivityXunHuanTimeTemplate(openServerTime)
	if timeTypeTemp == nil {
		return 0
	}
	cycleGroupList, ok := st.circleGroupListMap[timeTypeTemp.TimeType]
	if !ok {
		return 0
	}

	// 循环
	var indexList []int
	for index, _ := range cycleGroupList {
		mask := int32(1 << uint(index))
		if st.hadChooseGroup&mask != 0 {
			continue
		}

		indexList = append(indexList, index)
	}

	// 重置
	if len(indexList) == 0 {
		st.hadChooseGroup = 0

		// 第一个和上一轮最后一个不同
		for index, group := range cycleGroupList {
			if st.lastChooseGroup == group {
				continue
			}

			indexList = append(indexList, index)
		}
	}

	randomHit := mathutils.RandomRange(0, len(indexList))
	arrGroupIndex := indexList[randomHit]
	chooseGroup := cycleGroupList[arrGroupIndex]
	st.hadChooseGroup += 1 << uint(arrGroupIndex)

	if len(indexList) == 1 {
		st.lastChooseGroup = chooseGroup
	}

	return chooseGroup
}

func (st *welfareTemplateService) getXunHuanDayNum() int32 {
	for _, subMap := range st.xunhuanMap {
		for _, subOfSubMap := range subMap {
			return int32(len(subOfSubMap))
		}
	}

	return 0
}

func (st *welfareTemplateService) GetActivityXunHuanTemplate(openServerTime int64, arrIndex, cycDay int32) *gametemplate.OpenActivityXunHuanTemplate {
	timeTypeTemp := st.GetActivityXunHuanTimeTemplate(openServerTime)
	if timeTypeTemp == nil {
		return nil
	}

	subMap, ok := st.xunhuanMap[timeTypeTemp.TimeType]
	if !ok {
		return nil
	}
	subOfSubMap, ok := subMap[arrIndex]
	if !ok {
		return nil
	}
	temp, ok := subOfSubMap[cycDay]
	if !ok {
		return nil
	}

	return temp

}

func (st *welfareTemplateService) GetActivityXunHuanTimeTemplate(openServerTime int64) *gametemplate.OpenActivityXunHuanTimeTemplate {
	openDay := st.getOpenDiff(openServerTime) + 1
	for _, timeTypeTemp := range st.xunhuanTypeMap {
		if !timeTypeTemp.IsOnRange(openDay) {
			continue
		}

		return timeTypeTemp
	}

	return nil
}

func (st *welfareTemplateService) GetCollectPokerTemplate(groupId int32, pokerType groupcollectenum.PokerType) *gametemplate.ChouJiangPokerTemplate {
	subMap, ok := st.pokerTempMap[groupId]
	if !ok {
		return nil
	}
	temp, ok := subMap[pokerType]
	if !ok {
		return nil
	}

	return temp
}

func (st *welfareTemplateService) GetZhuanShengCircleBargainTemplate(class, buyNum int32) *gametemplate.CircleBargainTemplate {
	tempList, ok := st.zhuanshengBargainMap[class]
	if !ok {
		return nil
	}

	for _, temp := range tempList {
		if buyNum < temp.ItemCountMin || buyNum > temp.ItemCountMax {
			continue
		}

		return temp
	}

	return nil
}

func (st *welfareTemplateService) GetRankMaxEndTime(openTime int64) (int64, []*gametemplate.OpenserverTimeTemplate) {
	openTimeSubMap, ok := st.openActivityTimeMap[welfaretypes.OpenActivityTypeRank]
	if !ok {
		return -1, nil
	}
	now := global.GetGame().GetTimeService().Now()
	maxEndTime := int64(0)
	var newTempList []*gametemplate.OpenserverTimeTemplate
	for _, tempList := range openTimeSubMap {
		for _, temp := range tempList {

			if temp.IsRankCharge() || temp.IsRankCost() || temp.IsRankCharm() || temp.IsRankMarryDevelop() {
				continue
			}

			endTime, _ := temp.GetEndTime(now, openTime)
			if maxEndTime > endTime {
				continue
			}

			maxEndTime = endTime
		}
		if len(tempList) > 0 {
			newTempList = append(newTempList, tempList...)
		}
	}

	return maxEndTime, newTempList
}

func (st *welfareTemplateService) GetLuckDrewTemplate(groupId int32) *gametemplate.LuckyDrewTemplate {
	return st.GetLuckDrewTemplateByArg(groupId, 0)
}

func (st *welfareTemplateService) GetLuckDrewTemplateByArg(groupId int32, level int32) *gametemplate.LuckyDrewTemplate {
	subMap, ok := st.luckyDrewMap[groupId]
	if !ok {
		return nil
	}

	temp, ok := subMap[level]
	if !ok {
		return nil
	}

	return temp
}

func (st *welfareTemplateService) GetDiscountTemplate(discountId int32) *gametemplate.DiscountTemplate {
	temp, ok := st.discountMap[discountId]
	if !ok {
		return nil
	}

	return temp
}

func (st *welfareTemplateService) GetDiscountTemplateByDayGroup(dayGroup int32) []*gametemplate.DiscountTemplate {
	tempList, ok := st.discountMapOfGroup[dayGroup]
	if !ok {
		return nil
	}

	return tempList
}

func (st *welfareTemplateService) GetGroupBiologyDropId(groupId, dropNum int32) int32 {
	dropTempList := st.groupBiologyDropMap[groupId]
	for _, teshuDropTemp := range dropTempList {
		if dropNum < teshuDropTemp.MinCount || dropNum > teshuDropTemp.MaxCount {
			continue
		}
		return teshuDropTemp.DropId
	}

	return 0
}

func (st *welfareTemplateService) GetRankActivityTimeTemplate() map[welfaretypes.OpenActivityRankSubType][]*gametemplate.OpenserverTimeTemplate {
	rankTimeTmepMap := make(map[welfaretypes.OpenActivityRankSubType][]*gametemplate.OpenserverTimeTemplate)

	for rankType := welfaretypes.MinOpenActivityRankSubType; rankType <= welfaretypes.MaxOpenActivityRankSubType; rankType++ {
		tempList := st.GetOpenActivityTimeTemplateByType(welfaretypes.OpenActivityTypeRank, rankType)
		if len(tempList) > 0 {
			rankTimeTmepMap[rankType] = tempList
		}
	}

	return rankTimeTmepMap
}

func (st *welfareTemplateService) GetAllActivityTimeTemplate() map[int32]*gametemplate.OpenserverTimeTemplate {
	return st.openActivityTimeByGroupMap
}

func (st *welfareTemplateService) GetOpenActivityTimeTemplateByType(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) []*gametemplate.OpenserverTimeTemplate {
	openTimeSubMap, ok := st.openActivityTimeMap[typ]
	if !ok {
		return nil
	}
	timeTempList, ok := openTimeSubMap[subType]
	if !ok {
		return nil
	}

	return timeTempList
}

func (st *welfareTemplateService) GetOpenActivityTemplateByFuncType(funcType funcopentypes.FuncOpenType) []*gametemplate.OpenserverTimeTemplate {
	timeTempList, ok := st.openActivityTimeFuncIdMap[funcType]
	if !ok {
		return nil
	}

	return timeTempList
}

func (st *welfareTemplateService) GetOpenActivityTimeTemplate(groupId int32) *gametemplate.OpenserverTimeTemplate {
	timeTemp, ok := st.openActivityTimeByGroupMap[groupId]
	if !ok {
		return nil
	}

	return timeTemp
}

func (st *welfareTemplateService) GetOpenActivityTimeTemplateById(id int32) *gametemplate.OpenserverTimeTemplate {
	timeTemp, ok := st.openActivityTimeIdMap[id]
	if !ok {
		return nil
	}

	return timeTemp
}

func (st *welfareTemplateService) GetGoldLabaTemplate(groupId, times int32) *gametemplate.GoldLaBaTemplate {
	tempMap, ok := st.goldLabaMap[groupId]
	if !ok {
		return nil
	}
	temp, ok := tempMap[times]
	if !ok {
		return nil
	}
	return temp
}

func (st *welfareTemplateService) GetRandomLaBaTemplate(groupId int32) *gametemplate.GoldLaBaTemplate {
	randomTimes := int32(mathutils.RandomRange(1, len(st.goldLabaMap)+1))
	return st.GetGoldLabaTemplate(groupId, randomTimes)
}

func (st *welfareTemplateService) GetTimesRewTemplate(id int32) *gametemplate.TimesRewTemplate {
	temp, ok := st.timesRewMap[id]
	if !ok {
		return nil
	}
	return temp
}

func (st *welfareTemplateService) GetTimesRewTemplateByGorup(groupId int32) []*gametemplate.TimesRewTemplate {
	tempList, ok := st.timesRewMapOfGroup[groupId]
	if !ok {
		return nil
	}
	return tempList
}

func (st *welfareTemplateService) GetMadeTemplate(groupId, level int32) *gametemplate.MadeTemplate {
	tempList, ok := st.madeGroupMap[groupId]
	if !ok {
		return nil
	}

	for _, temp := range tempList {
		if level > temp.LevelMax {
			continue
		}
		if level < temp.LevelMin {
			continue
		}

		return temp
	}
	return nil
}

func (st *welfareTemplateService) GetActivityMergeXunHuanTemplate(openTime, mergeTime int64) *gametemplate.OpenActivityMergeXunHuanTemplate {
	now := global.GetGame().GetTimeService().Now()

	// 活动结束
	mergeDiff, _ := timeutils.DiffDay(now, mergeTime)
	mergeXunHuanKeepDay := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMergeXunHuanKeepDay)
	if mergeDiff >= mergeXunHuanKeepDay {
		return nil
	}

	// 计算开始日
	openDay := st.getOpenDiff(openTime) + 1
	cycleDay := openDay % st.maxMergeXunHuanDay
	if cycleDay == 0 {
		cycleDay = st.maxMergeXunHuanDay
	}

	temp, ok := st.mergeXunHuanMap[cycleDay]
	if !ok {
		return nil
	}

	return temp
}

func (st *welfareTemplateService) getOpenDiff(openTime int64) int32 {
	now := global.GetGame().GetTimeService().Now()
	openDiff, _ := timeutils.DiffDay(now, openTime)
	return openDiff
}

var (
	once sync.Once
	st   *welfareTemplateService
)

func Init() (err error) {
	once.Do(func() {
		st = &welfareTemplateService{}
		err = st.init()
	})

	return
}

func GetWelfareTemplateService() WelfareTemplateService {
	return st
}
