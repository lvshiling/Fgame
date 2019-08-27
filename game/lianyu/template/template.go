package template

import (
	"fgame/fgame/core/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/lianyu/types"
	lianyutypes "fgame/fgame/game/lianyu/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//无间炼狱常量接口处理
type LianYuTemplateService interface {
	//获取常量配置
	GetConstantTemplate(activityType activitytypes.ActivityType) *gametemplate.LianYuConstantTemplate
	//出生位置配置
	GetBornPosTemplate(bornType lianyutypes.LianYuBornType) *gametemplate.LianYuPosTemplate
}

type lianYuTemplateService struct {
	//无间炼狱
	constantTempMap map[activitytypes.ActivityType]*gametemplate.LianYuConstantTemplate
	//出生配置
	lianYuPosMap map[types.LianYuBornType]map[lianyutypes.LianYuPosType]*gametemplate.LianYuPosTemplate
}

//初始化
func (rs *lianYuTemplateService) init() (err error) {

	rs.constantTempMap = make(map[activitytypes.ActivityType]*gametemplate.LianYuConstantTemplate)
	constantTemplateMap := template.GetTemplateService().GetAll((*gametemplate.LianYuConstantTemplate)(nil))
	for _, to := range constantTemplateMap {
		constTemp := to.(*gametemplate.LianYuConstantTemplate)
		rs.constantTempMap[constTemp.GetActivityType()] = constTemp
	}

	rs.lianYuPosMap = make(map[types.LianYuBornType]map[lianyutypes.LianYuPosType]*gametemplate.LianYuPosTemplate)
	lianYuTemplateMap := template.GetTemplateService().GetAll((*gametemplate.LianYuPosTemplate)(nil))
	for _, templateObject := range lianYuTemplateMap {
		lianYuPosTemplate, _ := templateObject.(*gametemplate.LianYuPosTemplate)

		bornType := lianYuPosTemplate.GetBornType()
		posType := lianYuPosTemplate.GetPosType()

		lianYuPosTypeMap, exist := rs.lianYuPosMap[bornType]
		if !exist {
			lianYuPosTypeMap = make(map[lianyutypes.LianYuPosType]*gametemplate.LianYuPosTemplate)
			rs.lianYuPosMap[bornType] = lianYuPosTypeMap
		}
		lianYuPosTypeMap[posType] = lianYuPosTemplate
	}

	return nil
}

func (rs *lianYuTemplateService) GetConstantTemplate(activityType activitytypes.ActivityType) *gametemplate.LianYuConstantTemplate {
	temp, ok := rs.constantTempMap[activityType]
	if !ok {
		return nil
	}
	return temp
}

func (rs *lianYuTemplateService) GetBornPosTemplate(bornType lianyutypes.LianYuBornType) *gametemplate.LianYuPosTemplate {
	lianYuPosTypeMap, exist := rs.lianYuPosMap[bornType]
	if !exist {
		return nil
	}
	lianYuTempalte, exist := lianYuPosTypeMap[lianyutypes.LianYuPosTypeMin]
	if !exist {
		return nil
	}
	return lianYuTempalte
}

var (
	once sync.Once
	cs   *lianYuTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &lianYuTemplateService{}
		err = cs.init()
	})
	return err
}

func GetLianYuTemplateService() LianYuTemplateService {
	return cs
}
