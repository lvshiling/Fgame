package template

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//充值配置
type ChargeTemplateService interface {
	//获取充值模板
	GetChargeTemplate(chargeId int32) *gametemplate.ChargeTemplate
	GetChargeTemplateByType(typ logintypes.SDKType, subType int32) *gametemplate.ChargeTemplate
	GetChargeTemplateByGold(typ logintypes.SDKType, gold int32) *gametemplate.ChargeTemplate
	//获取渠道模板
	GetQuDaoTemplateByType(typ logintypes.SDKType) *gametemplate.QuDaoTemplate
}

type chargeTemplateService struct {
	//充值配置
	chargeMap     map[int32]*gametemplate.ChargeTemplate
	chargeTypeMap map[logintypes.SDKType]map[int32]*gametemplate.ChargeTemplate
	//渠道配置
	quDaoMap map[logintypes.SDKType]*gametemplate.QuDaoTemplate
}

//初始化
func (cs *chargeTemplateService) init() error {
	cs.chargeMap = make(map[int32]*gametemplate.ChargeTemplate)
	cs.chargeTypeMap = make(map[logintypes.SDKType]map[int32]*gametemplate.ChargeTemplate)
	cs.quDaoMap = make(map[logintypes.SDKType]*gametemplate.QuDaoTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ChargeTemplate)(nil))
	for _, temp := range templateMap {
		chargeTemplate, _ := temp.(*gametemplate.ChargeTemplate)
		cs.chargeMap[int32(chargeTemplate.TemplateId())] = chargeTemplate

		//type
		subMap, ok := cs.chargeTypeMap[chargeTemplate.GetType()]
		if !ok {
			subMap = make(map[int32]*gametemplate.ChargeTemplate)
			cs.chargeTypeMap[chargeTemplate.GetType()] = subMap
		}

		subMap[chargeTemplate.SubType] = chargeTemplate
	}

	quDaoTemplateMap := template.GetTemplateService().GetAll((*gametemplate.QuDaoTemplate)(nil))
	for _, temp := range quDaoTemplateMap {
		quDaoTemp, _ := temp.(*gametemplate.QuDaoTemplate)
		cs.quDaoMap[quDaoTemp.GetType()] = quDaoTemp
	}

	return nil
}

func (cs *chargeTemplateService) GetChargeTemplate(chargeId int32) *gametemplate.ChargeTemplate {
	chargeTemp, ok := cs.chargeMap[chargeId]
	if !ok {
		return nil
	}
	return chargeTemp
}

func (cs *chargeTemplateService) GetQuDaoTemplateByType(typ logintypes.SDKType) *gametemplate.QuDaoTemplate {
	quDaoTemp, ok := cs.quDaoMap[typ]
	if !ok {
		return nil
	}

	return quDaoTemp
}

func (cs *chargeTemplateService) GetChargeTemplateByType(typ logintypes.SDKType, subType int32) *gametemplate.ChargeTemplate {
	subMap, ok := cs.chargeTypeMap[typ]
	if !ok {
		return nil
	}

	chargeTemp, ok := subMap[subType]
	if !ok {
		return nil
	}
	return chargeTemp
}
func (cs *chargeTemplateService) GetChargeTemplateByGold(typ logintypes.SDKType, gold int32) *gametemplate.ChargeTemplate {
	subMap, ok := cs.chargeTypeMap[typ]
	if !ok {
		return nil
	}

	for _, chargeTemplate := range subMap {
		if chargeTemplate.Gold == gold {
			return chargeTemplate
		}
	}

	return nil
}

var (
	once sync.Once
	cs   *chargeTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &chargeTemplateService{}
		err = cs.init()
	})
	return err
}

func GetChargeTemplateService() ChargeTemplateService {
	return cs
}
