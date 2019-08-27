package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/godsiege/types"
	godsiegegametypes "fgame/fgame/game/godsiege/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//神兽攻城常量接口处理
type GodSiegeTemplateService interface {
	//获取常量配置
	GetConstantTemplate() *gametemplate.GongChengConstantTemplate
	//玩家出生出生配置
	GetPlayerPosTemplate(mapId int32, posType godsiegegametypes.GodSiegePosType) *gametemplate.GongChengPosTemplate
	//boss出生配置
	GetBossPosTemplate(mapId int32) *gametemplate.GongChengPosTemplate
}

type godSiegeTemplateService struct {
	//神兽攻城
	constantTemplate *gametemplate.GongChengConstantTemplate
	//出生配置
	godSiegePosMap map[int32]map[types.GodSiegeBornType]map[godsiegegametypes.GodSiegePosType]*gametemplate.GongChengPosTemplate
}

//初始化
func (rs *godSiegeTemplateService) init() (err error) {
	rs.godSiegePosMap = make(map[int32]map[types.GodSiegeBornType]map[godsiegegametypes.GodSiegePosType]*gametemplate.GongChengPosTemplate)
	constantTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GongChengConstantTemplate)(nil))
	for _, templateObject := range constantTemplateMap {
		rs.constantTemplate, _ = templateObject.(*gametemplate.GongChengConstantTemplate)
		break
	}

	godSiegeTemplateMap := template.GetTemplateService().GetAll((*gametemplate.GongChengPosTemplate)(nil))
	for _, templateObject := range godSiegeTemplateMap {
		godSiegePosTemplate, _ := templateObject.(*gametemplate.GongChengPosTemplate)

		bornType := godSiegePosTemplate.GetBornType()
		posType := godSiegePosTemplate.GetPosType()
		mapId := godSiegePosTemplate.GetMapId()

		godSiegePosTypeMap, exist := rs.godSiegePosMap[mapId]
		if !exist {
			godSiegePosTypeMap = make(map[types.GodSiegeBornType]map[godsiegegametypes.GodSiegePosType]*gametemplate.GongChengPosTemplate)
			rs.godSiegePosMap[mapId] = godSiegePosTypeMap
		}
		godSiegePosBornMap, exist := godSiegePosTypeMap[bornType]
		if !exist {
			godSiegePosBornMap = make(map[godsiegegametypes.GodSiegePosType]*gametemplate.GongChengPosTemplate)
			godSiegePosTypeMap[bornType] = godSiegePosBornMap
		}
		godSiegePosBornMap[posType] = godSiegePosTemplate
	}

	return nil
}

func (rs *godSiegeTemplateService) GetConstantTemplate() *gametemplate.GongChengConstantTemplate {
	return rs.constantTemplate
}

func (rs *godSiegeTemplateService) GetPlayerPosTemplate(mapId int32, posType godsiegegametypes.GodSiegePosType) *gametemplate.GongChengPosTemplate {
	godSiegeBornMap, exist := rs.godSiegePosMap[mapId]
	if !exist {
		return nil
	}
	godSiegePosTypeMap, exist := godSiegeBornMap[types.GodSiegeBornTypePlayer]
	if !exist {
		return nil
	}
	godSiegeTempalte, exist := godSiegePosTypeMap[posType]
	if !exist {
		return nil
	}
	return godSiegeTempalte
}

func (rs *godSiegeTemplateService) GetBossPosTemplate(mapId int32) *gametemplate.GongChengPosTemplate {
	godSiegeBornMap, exist := rs.godSiegePosMap[mapId]
	if !exist {
		return nil
	}
	godSiegePosTypeMap, exist := godSiegeBornMap[types.GodSiegeBornTypeBoss]
	if !exist {
		return nil
	}
	godSiegeTempalte, exist := godSiegePosTypeMap[godsiegegametypes.GodSiegePosTypeMin]
	if !exist {
		return nil
	}
	return godSiegeTempalte
}

var (
	once sync.Once
	cs   *godSiegeTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &godSiegeTemplateService{}
		err = cs.init()
	})
	return err
}

func GetGodSiegeTemplateService() GodSiegeTemplateService {
	return cs
}
