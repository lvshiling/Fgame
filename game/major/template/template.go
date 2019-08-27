package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	majortypes "fgame/fgame/game/major/types"
	gametemplate "fgame/fgame/game/template"
	"strconv"
	"sync"
)

type MajorTemplate interface {
	TemplateId() int
	GetMapId() int32
	GetBossId() int32
	GetRewItemMap() map[int32]int32
	GetMajorType() majortypes.MajorType
	GetSaodangNeedGold() int32
	GetSaodangItemMap(saoDangNum int32) map[int32]int32
	GetSaodangRewardItemMap(saoDangNum int32) map[int32]int32
	GetSaodangRewardDropArr() []int32
	GetRawSilver() int64
	GetRawBindGold() int32
	GetRawGold() int32
	GetRawExp() int64
	GetRawExpPoint() int64
}

//双修接口处理
type MajorTemplateService interface {
	//获取双修配置
	GetMajorTemplate(majorType majortypes.MajorType, fubenId int32) MajorTemplate
	//获取双修最大默认次数
	GetMajorDefaultMaxNum(majorType majortypes.MajorType) int32
	//双修邀请cd时间
	GetInvitePairCdTime(majorType majortypes.MajorType) int64
	//获取双修cd时间
	GetInvitePairCdTimeStr() string
}

type majorTemplateService struct {
	//双修副本
	majorTemplate *gametemplate.XianFuDoubleRepairTemplate
	// 夫妻副本
	fuqiFuBenMap map[int32]*gametemplate.MarryFuBenTemplate
	cdTime       string
}

//初始化
func (ms *majorTemplateService) init() (err error) {

	//双修副本
	majorTemplateMap := template.GetTemplateService().GetAll((*gametemplate.XianFuDoubleRepairTemplate)(nil))
	for _, templateObject := range majorTemplateMap {
		ms.majorTemplate, _ = templateObject.(*gametemplate.XianFuDoubleRepairTemplate)
		break
	}

	//夫妻副本
	ms.fuqiFuBenMap = make(map[int32]*gametemplate.MarryFuBenTemplate)
	fuqiTemplateMap := template.GetTemplateService().GetAll((*gametemplate.MarryFuBenTemplate)(nil))
	for _, templateObject := range fuqiTemplateMap {
		fuqiTemp, _ := templateObject.(*gametemplate.MarryFuBenTemplate)
		ms.fuqiFuBenMap[int32(fuqiTemp.Id)] = fuqiTemp
	}

	cdTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMajorCdTime) / 1000
	ms.cdTime = strconv.Itoa(int(cdTime))

	return nil
}

func (rs *majorTemplateService) GetMajorTemplate(majorType majortypes.MajorType, fubenId int32) MajorTemplate {
	switch majorType {
	case majortypes.MajorTypeShuangXiu:
		return rs.majorTemplate
	case majortypes.MajorTypeFuQi:
		return rs.fuqiFuBenMap[fubenId]
	}

	return nil
}

func (rs *majorTemplateService) GetMajorDefaultMaxNum(majorType majortypes.MajorType) int32 {
	switch majorType {
	case majortypes.MajorTypeShuangXiu:
		return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMajorDefaultMaxNum)
	case majortypes.MajorTypeFuQi:
		return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCoupleDefaultMaxNum)
	}

	return 0
}

func (rs *majorTemplateService) GetInvitePairCdTimeStr() string {
	return rs.cdTime
}

func (rs *majorTemplateService) GetInvitePairCdTime(majorType majortypes.MajorType) int64 {
	switch majorType {
	case majortypes.MajorTypeShuangXiu:
		return int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMajorCdTime))
	case majortypes.MajorTypeFuQi:
		return int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCoupleFuBenCdTime))
	}

	return 0
}

var (
	once sync.Once
	cs   *majorTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &majorTemplateService{}
		err = cs.init()
	})
	return err
}

func GetMajorTemplateService() MajorTemplateService {
	return cs
}
