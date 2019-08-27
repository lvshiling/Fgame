package xianfu

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/constant/constant"
	gametemplate "fgame/fgame/game/template"
	xianfutypes "fgame/fgame/game/xianfu/types"
	"sync"
)

type XianFuTemplate interface {
	TemplateId() int
	GetMapTemplate() *gametemplate.MapTemplate
	GetXianFuType() xianfutypes.XianfuType
	GetBossId() int32
	GetUpgradeTime() int64
	GetUpgradeGold() int32
	GetUpgradeBindGold() int32
	GetUpgradeYinliang() int64
	GetUpgradeItemId() int32
	GetUpgradeItemNum() int32
	GetSpeedUpNeedGold() float64
	GetRawExp() int64
	GetRawExpPoint() int64
	GetRawGold() int32
	GetRawBindGold() int32
	GetRawSilver() int64
	GetNeedItemId() int32
	GetNeedItemCount() int32
	GetNextId() int32
	GetFree() int32
	GetSaodangNeedGold() int32
	GetSaodangItemMap(saoDangNum int32) map[int32]int32
	GetSaodangRewardItemMap(saoDangNum int32) map[int32]int32
	GetSaodangRewardDropArr() []int32
	GetChallengeRewardsItemMap() map[int32]int32
	GetGroupLimit() int32
}

//模板配置获取接口
type XianfuTemplateService interface {
	GetXianfu(xianfuId int32, xfType xianfutypes.XianfuType) XianFuTemplate
	GetFirstXianfu(xianfutypes.XianfuType) XianFuTemplate
	//仙府基础挑战次数
	GetBasicPlayTimes(xfType xianfutypes.XianfuType) int32
	//仙府免费挑战次数
	GetFreePlayTimes(typ xianfutypes.XianfuType, xianfuId int32) int32
}

type xianfuTemplateService struct {
	xianfuSilverMap map[int32]XianFuTemplate
	xianfuExpMap    map[int32]XianFuTemplate
}

//初始化仙府配置
func (s *xianfuTemplateService) init() (err error) {
	//银两副本
	s.xianfuSilverMap = make(map[int32]XianFuTemplate)

	silverTempMap := template.GetTemplateService().GetAll((*gametemplate.XianFuSilverTemplate)(nil))
	for _, temp := range silverTempMap {
		silverTemp, _ := temp.(*gametemplate.XianFuSilverTemplate)
		s.xianfuSilverMap[int32(silverTemp.TemplateId())] = silverTemp
	}

	//经验副本
	s.xianfuExpMap = make(map[int32]XianFuTemplate)
	expTempMap := template.GetTemplateService().GetAll((*gametemplate.XianFuExpTemplate)(nil))
	for _, temp := range expTempMap {
		expTemp, _ := temp.(*gametemplate.XianFuExpTemplate)
		s.xianfuExpMap[int32(expTemp.TemplateId())] = expTemp
	}

	return
}

//获取初始秘境仙府配置
func (s *xianfuTemplateService) GetFirstXianfu(xfType xianfutypes.XianfuType) XianFuTemplate {
	return getXianfuMap(xfType)[int32(1)]
}

//获取秘境仙府配置
func (s *xianfuTemplateService) GetXianfu(xianfuId int32, xfType xianfutypes.XianfuType) XianFuTemplate {
	return getXianfuMap(xfType)[xianfuId]
}

//按类型获取秘境仙府配置Map
func getXianfuMap(xfType xianfutypes.XianfuType) map[int32]XianFuTemplate {
	switch xfType {
	case xianfutypes.XianfuTypeSilver:
		return s.xianfuSilverMap
	case xianfutypes.XianfuTypeExp:
		return s.xianfuExpMap
	}
	return nil
}

//获取基础次数
func (s *xianfuTemplateService) GetBasicPlayTimes(xfType xianfutypes.XianfuType) int32 {
	constantType := xfType.GetChallengeNumConstantType()
	maxChallengeTimes := constant.GetConstantService().GetConstant(constantType)

	return maxChallengeTimes
}

//仙府免费挑战次数
func (s *xianfuTemplateService) GetFreePlayTimes(typ xianfutypes.XianfuType, xianfuId int32) int32 {
	temp := s.GetXianfu(xianfuId, typ)
	if temp == nil {
		return 0
	}

	return temp.GetFree()
}

var (
	once sync.Once
	s    *xianfuTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &xianfuTemplateService{}
		err = s.init()
	})
	return
}

func GetXianfuTemplateService() XianfuTemplateService {
	return s
}
