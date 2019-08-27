package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/center/center"
	gamecentertypes "fgame/fgame/game/center/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/marry/types"
	marrytypes "fgame/fgame/game/marry/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"sync"
)

type MarryTemplateService interface {
	//获取婚戒培养配置
	GetMarryRingTemplate(ringType types.MarryRingType, level int32) *gametemplate.MarryRingTemplate
	//获取爱情树培养配置
	GetMarryLoveTreeTemplate(id int32) *gametemplate.MarryTreeTemplate
	//获取婚宴档次配置
	GetMarryBanquetTeamplate(banquetType types.MarryBanquetType, banquetSubType types.MarryBanquetSubType) *gametemplate.MarryBanquetTemplate
	//获取婚宴档次配置
	GetMarryBanquetTemplateByHouTai(houTai types.MarryHoutaiType, banquetType types.MarryBanquetType, banquetSubType types.MarryBanquetSubType) *gametemplate.MarryBanquetTemplate
	//获取贺礼档次配置
	GetMarryGiftTeamplate(id int32) *gametemplate.MarryGiftTemplate
	//获取婚车路径配置
	GetMarryMoveTeamplate(id int32) *gametemplate.MarryMoveTemplate
	//婚车开始位置
	GetMarryMoveFirstTeamplate() *gametemplate.MarryMoveTemplate
	//获取结婚常量模板
	GetMarryConstTempalte() *gametemplate.MarryTemplate
	//结婚亲密度要求
	GetMarryConstIntimacy() (int32, bool)
	//协议离婚亲密度剩余百分比
	GetMarryDivorceLeftIntimacy() float64
	//婚戒等级最大等级差
	GetMarryConstRingLevelDiff() int32
	//婚宴最早一场开始时间
	GetMarryFisrtWedTime(now int64) (int64, error)
	//系统一天能举办的婚宴次数
	GetMarryConstWedNum() int32
	//一次婚礼持续时间(毫秒)
	GetMarryDurationTime() int64
	//获取清场时间
	GetMarryQingChangTime() int64
	//预定消耗
	GetMarryGradeCost(grade types.MarryBanquetSubTypeWed, hunCheGrade types.MarryBanquetSubTypeHunChe, sugarGrade types.MarryBanquetSubTypeSugar) (costBindGold int32, costGold int32, costSilver int64)
	//获取婚宴结束奖励
	GetMarryEndRewMap(grade types.MarryBanquetSubTypeWed, hunCheGrade types.MarryBanquetSubTypeHunChe, sugarGrade types.MarryBanquetSubTypeSugar) map[int32]int32

	//获得游车赠送模板
	GetMarryPreGiftTemplate(preType types.MarryPreGiftType) *gametemplate.MarryTuiSongTemplate
	//表白系统配置
	GetMarryDeveopTemplate(level int32) *gametemplate.MarryDevelopTemplate

	//获得纪念模板
	GetMarryJiNianTemplate(subType types.MarryBanquetSubTypeWed) *gametemplate.MarryJiNianTemplate
	GetAllMarryJiNianTemplate() map[types.MarryBanquetSubTypeWed]*gametemplate.MarryJiNianTemplate
	//获得信物
	GetMarryXinWuGroupTemplate(id int32) *gametemplate.MarryXinWuSuitGroupTemplate
	GetMarryXinWuItem(suitId int32, posId int32) *gametemplate.MarryXinWuTemplate
	//获得结婚结婚纪念赠送的物品
	GetMarryJiNianSjItems() map[int32]int32

	// SetHouTaiType(houtaiTp types.MarryHoutaiType)
	GetHouTaiType() types.MarryHoutaiType

	GetBanquetMap1() map[types.MarryBanquetType]map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate

	GetBanquetMap2() map[types.MarryBanquetType]map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate
}

type marryTemplateService struct {
	//结婚常量配置
	marryConstTemplate *gametemplate.MarryTemplate
	//婚戒map
	ringMap map[types.MarryRingType]map[int32]*gametemplate.MarryRingTemplate
	//爱情树map
	treeMap map[int32]*gametemplate.MarryTreeTemplate
	//婚宴map
	//TODO xubin 能不能把这几个婚宴map合成一个map
	banquetMap1 map[types.MarryBanquetType]map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate //现在版的
	banquetMap2 map[types.MarryBanquetType]map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate //廉价版的
	banquetMap3 map[types.MarryBanquetType]map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate //贵价版的
	//贺礼档次map
	giftMap map[int32]*gametemplate.MarryGiftTemplate
	//婚车移动map
	moveMap map[int32]*gametemplate.MarryMoveTemplate
	//第一场婚宴预定时间
	wedBeginTime int64

	//结婚纪念玩家赠送
	preGiftMap map[types.MarryPreGiftType]*gametemplate.MarryTuiSongTemplate
	//表白系统配置
	marryDevelopMap map[int32]*gametemplate.MarryDevelopTemplate

	//结婚纪念配置
	jinianMap map[types.MarryBanquetSubTypeWed]*gametemplate.MarryJiNianTemplate
	//结婚信物配置
	xinWuGroupMap map[int32]*gametemplate.MarryXinWuSuitGroupTemplate

	// houtaiType types.MarryHoutaiType
}

//初始化
func (mts *marryTemplateService) init() (err error) {
	mts.ringMap = make(map[types.MarryRingType]map[int32]*gametemplate.MarryRingTemplate)
	mts.treeMap = make(map[int32]*gametemplate.MarryTreeTemplate)
	mts.banquetMap1 = make(map[types.MarryBanquetType]map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate)
	mts.banquetMap2 = make(map[types.MarryBanquetType]map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate)
	mts.banquetMap3 = make(map[types.MarryBanquetType]map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate)
	mts.giftMap = make(map[int32]*gametemplate.MarryGiftTemplate)
	mts.moveMap = make(map[int32]*gametemplate.MarryMoveTemplate)
	mts.preGiftMap = make(map[types.MarryPreGiftType]*gametemplate.MarryTuiSongTemplate)
	mts.marryDevelopMap = make(map[int32]*gametemplate.MarryDevelopTemplate)
	mts.jinianMap = make(map[types.MarryBanquetSubTypeWed]*gametemplate.MarryJiNianTemplate)
	mts.xinWuGroupMap = make(map[int32]*gametemplate.MarryXinWuSuitGroupTemplate)

	// mts.houtaiType = types.MarryHoutaiTypeCommon //一开始初始化为正常版

	ringMap := template.GetTemplateService().GetAll((*gametemplate.MarryRingTemplate)(nil))
	for _, templateObject := range ringMap {
		ringTemplate, _ := templateObject.(*gametemplate.MarryRingTemplate)

		typ := ringTemplate.GetRingType()
		ringTypeMap, exist := mts.ringMap[typ]
		if !exist {
			ringTypeMap = make(map[int32]*gametemplate.MarryRingTemplate)
			mts.ringMap[typ] = ringTypeMap
		}
		ringTypeMap[ringTemplate.Level] = ringTemplate
	}

	treeMap := template.GetTemplateService().GetAll((*gametemplate.MarryTreeTemplate)(nil))
	for _, templateObject := range treeMap {
		treeTemplate, _ := templateObject.(*gametemplate.MarryTreeTemplate)
		mts.treeMap[int32(treeTemplate.TemplateId())] = treeTemplate
	}

	banquetMap := template.GetTemplateService().GetAll((*gametemplate.MarryBanquetTemplate)(nil))
	for _, templateObject := range banquetMap {
		banquetTemplate, _ := templateObject.(*gametemplate.MarryBanquetTemplate)

		typ := banquetTemplate.GetBanquetType()
		subType := banquetTemplate.GetBanquetSubType()
		houtaiTp := banquetTemplate.GetHoutaiType()
		if houtaiTp == types.MarryHoutaiTypeCommon {
			banquetTypeMap, ok := mts.banquetMap1[typ]
			if !ok {
				banquetTypeMap = make(map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate)
				mts.banquetMap1[typ] = banquetTypeMap
			}
			banquetTypeMap[subType] = banquetTemplate
		}
		if houtaiTp == types.MarryHoutaiTypeCheep {
			banquetTypeMap, ok := mts.banquetMap2[typ]
			if !ok {
				banquetTypeMap = make(map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate)
				mts.banquetMap2[typ] = banquetTypeMap
			}
			banquetTypeMap[subType] = banquetTemplate
		}
		if houtaiTp == types.MarryHoutaiTypeExp {
			banquetTypeMap, ok := mts.banquetMap3[typ]
			if !ok {
				banquetTypeMap = make(map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate)
				mts.banquetMap3[typ] = banquetTypeMap
			}
			banquetTypeMap[subType] = banquetTemplate
		}
	}

	giftMap := template.GetTemplateService().GetAll((*gametemplate.MarryGiftTemplate)(nil))
	for _, templateObject := range giftMap {
		giftTemplate, _ := templateObject.(*gametemplate.MarryGiftTemplate)
		mts.giftMap[int32(giftTemplate.TemplateId())] = giftTemplate
	}

	moveMap := template.GetTemplateService().GetAll((*gametemplate.MarryMoveTemplate)(nil))
	for _, templateObject := range moveMap {
		moveTemplate, _ := templateObject.(*gametemplate.MarryMoveTemplate)
		mts.moveMap[int32(moveTemplate.TemplateId())] = moveTemplate
	}

	marryMap := template.GetTemplateService().GetAll((*gametemplate.MarryTemplate)(nil))
	for _, templateObject := range marryMap {
		mts.marryConstTemplate, _ = templateObject.(*gametemplate.MarryTemplate)
		break
	}

	preGiftMap := template.GetTemplateService().GetAll((*gametemplate.MarryTuiSongTemplate)(nil))
	for _, templateObject := range preGiftMap {
		item := templateObject.(*gametemplate.MarryTuiSongTemplate)
		mts.preGiftMap[item.PreGiftType] = item
	}
	// 表白培养
	developMap := template.GetTemplateService().GetAll((*gametemplate.MarryDevelopTemplate)(nil))
	for _, templateObject := range developMap {
		developTemplate, _ := templateObject.(*gametemplate.MarryDevelopTemplate)
		mts.marryDevelopMap[developTemplate.Level] = developTemplate
	}

	//纪念配置
	jinianMap := template.GetTemplateService().GetAll((*gametemplate.MarryJiNianTemplate)(nil))
	for _, templateObject := range jinianMap {
		jinianTemplate, _ := templateObject.(*gametemplate.MarryJiNianTemplate)
		mts.jinianMap[jinianTemplate.GetMarrySubType()] = jinianTemplate
	}

	//信物配置
	xinWuMap := template.GetTemplateService().GetAll((*gametemplate.MarryXinWuSuitGroupTemplate)(nil))
	for _, templateObject := range xinWuMap {
		xinwuTemplate, _ := templateObject.(*gametemplate.MarryXinWuSuitGroupTemplate)
		mts.xinWuGroupMap[int32(xinwuTemplate.Id)] = xinwuTemplate
	}

	//wedBeginTime := strconv.Itoa(int(mts.marryConstTemplate.MarryFirstTime))
	//转字符串
	mts.wedBeginTime, err = timeutils.ParseDayOfHHMM(mts.marryConstTemplate.MarryFirstTime)
	if err != nil {
		return fmt.Errorf("marry: ParseDayOfHHMM 应该是ok的")
	}

	for wedType, _ := range types.GetMarryWedTypeMap() {
		banquetTypeMap, exist := mts.banquetMap1[types.MarryBanquetTypeWed]
		if !exist {
			return fmt.Errorf("marry: 婚宴档次%d配置应该存在", wedType)
		}
		_, exist = banquetTypeMap[wedType]
		if !exist {
			return fmt.Errorf("marry: 婚宴档次%d配置应该存在", wedType)
		}

		banquetTypeMap, exist = mts.banquetMap2[types.MarryBanquetTypeWed]
		if !exist {
			return fmt.Errorf("marry: 婚宴档次%d配置应该存在", wedType)
		}
		_, exist = banquetTypeMap[wedType]
		if !exist {
			return fmt.Errorf("marry: 婚宴档次%d配置应该存在", wedType)
		}

		banquetTypeMap, exist = mts.banquetMap3[types.MarryBanquetTypeWed]
		if !exist {
			return fmt.Errorf("marry: 婚宴档次%d配置应该存在", wedType)
		}
		_, exist = banquetTypeMap[wedType]
		if !exist {
			return fmt.Errorf("marry: 婚宴档次%d配置应该存在", wedType)
		}
	}

	for hunCheType, _ := range types.GetMarryHunCheTypeMap() {
		banquetTypeMap, exist := mts.banquetMap1[types.MarryBanquetTypeHunChe]
		if !exist {
			return fmt.Errorf("marry: 婚车档次%d配置应该存在", hunCheType)
		}
		_, exist = banquetTypeMap[hunCheType]
		if !exist {
			return fmt.Errorf("marry: 婚车档次%d配置应该存在", hunCheType)
		}

		banquetTypeMap, exist = mts.banquetMap2[types.MarryBanquetTypeHunChe]
		if !exist {
			return fmt.Errorf("marry: 婚车档次%d配置应该存在", hunCheType)
		}
		_, exist = banquetTypeMap[hunCheType]
		if !exist {
			return fmt.Errorf("marry: 婚车档次%d配置应该存在", hunCheType)
		}

		banquetTypeMap, exist = mts.banquetMap3[types.MarryBanquetTypeHunChe]
		if !exist {
			return fmt.Errorf("marry: 婚车档次%d配置应该存在", hunCheType)
		}
		_, exist = banquetTypeMap[hunCheType]
		if !exist {
			return fmt.Errorf("marry: 婚车档次%d配置应该存在", hunCheType)
		}
	}

	for sugarType, _ := range types.GetMarrySugarTypeMap() {
		banquetTypeMap, exist := mts.banquetMap1[types.MarryBanquetTypeSugar]
		if !exist {
			return fmt.Errorf("marry: 喜糖档次%d配置应该存在", sugarType)
		}
		_, exist = banquetTypeMap[sugarType]
		if !exist {
			return fmt.Errorf("marry: 喜糖档次%d配置应该存在", sugarType)
		}

		banquetTypeMap, exist = mts.banquetMap2[types.MarryBanquetTypeSugar]
		if !exist {
			return fmt.Errorf("marry: 喜糖档次%d配置应该存在", sugarType)
		}
		_, exist = banquetTypeMap[sugarType]
		if !exist {
			return fmt.Errorf("marry: 喜糖档次%d配置应该存在", sugarType)
		}

		banquetTypeMap, exist = mts.banquetMap3[types.MarryBanquetTypeSugar]
		if !exist {
			return fmt.Errorf("marry: 喜糖档次%d配置应该存在", sugarType)
		}
		_, exist = banquetTypeMap[sugarType]
		if !exist {
			return fmt.Errorf("marry: 喜糖档次%d配置应该存在", sugarType)
		}
	}

	for ringType, _ := range types.GetMarryRingTypeMap() {
		banquetTypeMap, exist := mts.banquetMap1[types.MarryBanquetTypeRing]
		if !exist {
			return fmt.Errorf("marry: 婚戒配置%d配置应该存在", ringType)
		}
		_, exist = banquetTypeMap[ringType]
		if !exist {
			return fmt.Errorf("marry: 婚戒配置%d配置应该存在", ringType)
		}

		banquetTypeMap, exist = mts.banquetMap2[types.MarryBanquetTypeRing]
		if !exist {
			return fmt.Errorf("marry: 婚戒配置%d配置应该存在", ringType)
		}
		_, exist = banquetTypeMap[ringType]
		if !exist {
			return fmt.Errorf("marry: 婚戒配置%d配置应该存在", ringType)
		}

		banquetTypeMap, exist = mts.banquetMap3[types.MarryBanquetTypeRing]
		if !exist {
			return fmt.Errorf("marry: 婚戒配置%d配置应该存在", ringType)
		}
		_, exist = banquetTypeMap[ringType]
		if !exist {
			return fmt.Errorf("marry: 婚戒配置%d配置应该存在", ringType)
		}
	}

	//婚车开始位置校验
	_, exist := mts.moveMap[1]
	if !exist {
		return fmt.Errorf("marry: 婚车开始地点应该存在的")
	}

	//校验数据
	err = mts.checkData()
	if err != nil {
		return
	}

	return nil
}

//校验
func (mts *marryTemplateService) checkData() (err error) {
	//校验婚戒
	wedRingMap := item.GetItemService().GetItemClassMap(itemtypes.ItemTypeWedRing)
	if wedRingMap == nil {
		return fmt.Errorf("marry:婚戒对配置应该是存在的")
	}
	for itemSubType, _ := range types.ItemWedRingMap {
		_, exist := wedRingMap[itemSubType]
		if !exist {
			return fmt.Errorf("marry:婚戒对配置应该是存在的")
		}
	}
	return
}

//婚车开始位置
func (mts *marryTemplateService) GetMarryMoveFirstTeamplate() *gametemplate.MarryMoveTemplate {
	return mts.moveMap[1]
}

//获取婚戒培养配置
func (mts *marryTemplateService) GetMarryRingTemplate(ringType types.MarryRingType, level int32) (ringTemplate *gametemplate.MarryRingTemplate) {
	ringTypeMap, exist := mts.ringMap[ringType]
	if !exist {
		return
	}
	ringTemplate, _ = ringTypeMap[level]
	return
}

//获取爱情树培养配置
func (mts *marryTemplateService) GetMarryLoveTreeTemplate(id int32) (treeTemplate *gametemplate.MarryTreeTemplate) {
	treeTemplate, exist := mts.treeMap[id]
	if !exist {
		return nil
	}
	return
}

//获取婚宴档次配置
func (mts *marryTemplateService) GetMarryBanquetTeamplate(banquetType types.MarryBanquetType,
	banquetSubType types.MarryBanquetSubType) (banquetTemplate *gametemplate.MarryBanquetTemplate) {
	return mts.GetMarryBanquetTemplateByHouTai(mts.GetHouTaiType(), banquetType, banquetSubType)
}

//获取婚宴档次配置
func (mts *marryTemplateService) GetMarryBanquetTemplateByHouTai(
	houtaiType types.MarryHoutaiType,
	banquetType types.MarryBanquetType,
	banquetSubType types.MarryBanquetSubType) (banquetTemplate *gametemplate.MarryBanquetTemplate) {
	if houtaiType == types.MarryHoutaiTypeCommon {
		banquetTypeMap, exist := mts.banquetMap1[banquetType]
		if !exist {
			return nil
		}
		banquetTemplate, exist = banquetTypeMap[banquetSubType]
		if !exist {
			return nil
		}
	}
	if houtaiType == types.MarryHoutaiTypeCheep {
		banquetTypeMap, exist := mts.banquetMap2[banquetType]
		if !exist {
			return nil
		}
		banquetTemplate, exist = banquetTypeMap[banquetSubType]
		if !exist {
			return nil
		}
	}
	if houtaiType == types.MarryHoutaiTypeExp {
		banquetTypeMap, exist := mts.banquetMap3[banquetType]
		if !exist {
			return nil
		}
		banquetTemplate, exist = banquetTypeMap[banquetSubType]
		if !exist {
			return nil
		}
	}
	return
}

//获取贺礼档次配置
func (mts *marryTemplateService) GetMarryGiftTeamplate(id int32) (giftTemplate *gametemplate.MarryGiftTemplate) {
	giftTemplate, exist := mts.giftMap[id]
	if !exist {
		return nil
	}
	return
}

//获取婚车路径配置
func (mts *marryTemplateService) GetMarryMoveTeamplate(id int32) (moveTemplate *gametemplate.MarryMoveTemplate) {
	moveTemplate, exist := mts.moveMap[id]
	if !exist {
		return nil
	}
	return
}

//获取结婚常量模板
func (mts *marryTemplateService) GetMarryConstTempalte() *gametemplate.MarryTemplate {
	return mts.marryConstTemplate
}

//获取结婚常量模板
func (mts *marryTemplateService) GetMarryDeveopTemplate(level int32) *gametemplate.MarryDevelopTemplate {
	temp, ok := mts.marryDevelopMap[level]
	if !ok {
		return nil
	}

	return temp
}

//结婚亲密度要求
func (mts *marryTemplateService) GetMarryConstIntimacy() (intimacy int32, flag bool) {
	intimacy, flag = mts.marryConstTemplate.GetQinMiDuByVersion(mts.GetHouTaiType())
	if !flag {
		return
	}
	flag = true
	return
}

//协议离婚亲密度剩余
func (mts *marryTemplateService) GetMarryDivorceLeftIntimacy() (coefficient float64) {
	return 1.0 - float64(mts.marryConstTemplate.DivorceQinmidu)/float64(common.MAX_RATE)
}

//婚戒等级最大等级差
func (mts *marryTemplateService) GetMarryConstRingLevelDiff() (levelDiff int32) {
	return mts.marryConstTemplate.RingLevelGap
}

//婚宴最早开始时间
func (mts *marryTemplateService) GetMarryFisrtWedTime(now int64) (wedFirstTime int64, err error) {
	beginDay, err := timeutils.BeginOfNow(now)
	if err != nil {
		return 0, err
	}
	return beginDay + mts.wedBeginTime, nil
}

//系统一天能举办的婚宴次数
func (mts *marryTemplateService) GetMarryConstWedNum() (wedNum int32) {
	return mts.marryConstTemplate.MarryAmount
}

//一次婚礼持续时间
func (mts *marryTemplateService) GetMarryDurationTime() int64 {
	return int64(mts.marryConstTemplate.MarryTime)
}

//清场时间
func (mts *marryTemplateService) GetMarryQingChangTime() int64 {
	return int64(mts.marryConstTemplate.QingChangTime)
}

func (mts *marryTemplateService) GetMarryGradeCost(grade types.MarryBanquetSubTypeWed,
	hunCheGrade types.MarryBanquetSubTypeHunChe,
	sugarGrade types.MarryBanquetSubTypeSugar) (costBindGold int32, costGold int32, costSilver int64) {
	banquetTemplate := mts.GetMarryBanquetTeamplate(types.MarryBanquetTypeWed, grade)
	if banquetTemplate == nil {
		return
	}
	hunTemplate := mts.GetMarryBanquetTeamplate(types.MarryBanquetTypeHunChe, hunCheGrade)
	if hunTemplate == nil {
		return
	}
	sugarTemplate := mts.GetMarryBanquetTeamplate(types.MarryBanquetTypeSugar, sugarGrade)
	if sugarTemplate == nil {
		return
	}
	costSilver = int64(banquetTemplate.UseSilver) + int64(hunTemplate.UseSilver) + int64(sugarTemplate.UseSilver)
	costGold = banquetTemplate.UseGold + hunTemplate.UseGold + sugarTemplate.UseGold
	costBindGold = banquetTemplate.UseBinggold + hunTemplate.UseBinggold + sugarTemplate.UseBinggold
	return
}

//获取婚宴结束奖励
func (mts *marryTemplateService) GetMarryEndRewMap(grade types.MarryBanquetSubTypeWed,
	hunCheGrade types.MarryBanquetSubTypeHunChe,
	sugarGrade types.MarryBanquetSubTypeSugar) (rewItemMap map[int32]int32) {
	rewItemMap = make(map[int32]int32)

	banquetTemplate := mts.GetMarryBanquetTeamplate(types.MarryBanquetTypeWed, grade)
	if banquetTemplate == nil {
		return
	}
	hunTemplate := mts.GetMarryBanquetTeamplate(types.MarryBanquetTypeHunChe, hunCheGrade)
	if hunTemplate == nil {
		return
	}
	sugarTemplate := mts.GetMarryBanquetTeamplate(types.MarryBanquetTypeSugar, sugarGrade)
	if sugarTemplate == nil {
		return
	}

	banquetMap := banquetTemplate.GetEndRewIdMap()
	for itemId, num := range banquetMap {
		rewItemMap[itemId] += num
	}

	hunCheMap := hunTemplate.GetEndRewIdMap()
	for itemId, num := range hunCheMap {
		rewItemMap[itemId] += num
	}

	sugarMap := sugarTemplate.GetEndRewIdMap()
	for itemId, num := range sugarMap {
		rewItemMap[itemId] += num
	}

	return

}

//获得游车赠送模板
func (mst *marryTemplateService) GetMarryPreGiftTemplate(preType types.MarryPreGiftType) *gametemplate.MarryTuiSongTemplate {
	return mst.preGiftMap[preType]
}

func (mst *marryTemplateService) GetMarryJiNianTemplate(subType types.MarryBanquetSubTypeWed) *gametemplate.MarryJiNianTemplate {
	return mst.jinianMap[subType]
}

func (mst *marryTemplateService) GetAllMarryJiNianTemplate() map[types.MarryBanquetSubTypeWed]*gametemplate.MarryJiNianTemplate {
	return mst.jinianMap
}

func (mst *marryTemplateService) GetMarryXinWuGroupTemplate(id int32) *gametemplate.MarryXinWuSuitGroupTemplate {
	return mst.xinWuGroupMap[id]
}

func (mst *marryTemplateService) GetMarryXinWuItem(suitId int32, posId int32) *gametemplate.MarryXinWuTemplate {
	_, exists := mst.xinWuGroupMap[suitId]
	if !exists {
		return nil
	}
	itemMap := mst.xinWuGroupMap[suitId].GetSuitItemMap()
	return itemMap[posId]
}

func (mst *marryTemplateService) GetMarryJiNianSjItems() map[int32]int32 {
	return mst.marryConstTemplate.GetJiNianSjItemMap()
}

func (mst *marryTemplateService) GetHouTaiType() types.MarryHoutaiType {
	marryKindType := center.GetCenterService().GetMarryKindType()
	houtaiType := marrytypes.MarryHoutaiTypeCommon
	switch marryKindType {
	case gamecentertypes.MarryPriceTypeCheap:
		houtaiType = marrytypes.MarryHoutaiTypeCheep
		break
	case gamecentertypes.MarryPriceTypeExp:
		houtaiType = marrytypes.MarryHoutaiTypeExp
		break
	default:
		houtaiType = marrytypes.MarryHoutaiTypeCommon
	}
	return houtaiType
}

func (mst *marryTemplateService) GetBanquetMap1() map[types.MarryBanquetType]map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate {
	return mst.banquetMap1
}

func (mst *marryTemplateService) GetBanquetMap2() map[types.MarryBanquetType]map[types.MarryBanquetSubType]*gametemplate.MarryBanquetTemplate {
	return mst.banquetMap2
}

var (
	once sync.Once
	cs   *marryTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &marryTemplateService{}
		err = cs.init()
	})
	return err
}

func GetMarryTemplateService() MarryTemplateService {
	return cs
}
