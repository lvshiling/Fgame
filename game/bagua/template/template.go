package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//八卦秘境接口处理
type BaGuaTemplateService interface {
	//获取八卦秘境配置通过等级
	GetBaGuaTemplateByLevel(level int32) *gametemplate.BaGuaMiJingTemplate
	//夫妻助战邀请cd时间
	GetInvitePairCdTime() int64
	//通过层数获取八卦秘境补偿模板
	GetBaGuaMiJingBuChangTemplateByLevel(level int32) *gametemplate.BaGuaMiJingBuChangTemplate
}

type baGuaTemplateService struct {
	//八卦秘境
	baGuaTaTemplateMap map[int32]*gametemplate.BaGuaMiJingTemplate
	//八卦秘境补偿
	baGuaMiJingBuChangTemplateMap map[int32]*gametemplate.BaGuaMiJingBuChangTemplate
}

//初始化
func (rs *baGuaTemplateService) init() (err error) {
	rs.baGuaTaTemplateMap = make(map[int32]*gametemplate.BaGuaMiJingTemplate)
	rs.baGuaMiJingBuChangTemplateMap = make(map[int32]*gametemplate.BaGuaMiJingBuChangTemplate)

	//赋值 baGuaTaTemplateMap
	tjtTemplateMap := template.GetTemplateService().GetAll((*gametemplate.BaGuaMiJingTemplate)(nil))
	for _, templateObject := range tjtTemplateMap {
		baGuaTemplate, _ := templateObject.(*gametemplate.BaGuaMiJingTemplate)
		rs.baGuaTaTemplateMap[baGuaTemplate.Level] = baGuaTemplate
	}

	buchangTempMap := template.GetTemplateService().GetAll((*gametemplate.BaGuaMiJingBuChangTemplate)(nil))
	for _, tempObj := range buchangTempMap {
		buchangTemp, _ := tempObj.(*gametemplate.BaGuaMiJingBuChangTemplate)
		rs.baGuaMiJingBuChangTemplateMap[buchangTemp.Level] = buchangTemp
	}

	return nil
}

//获取八卦秘境配置通过等级
func (rs *baGuaTemplateService) GetBaGuaTemplateByLevel(level int32) *gametemplate.BaGuaMiJingTemplate {
	to, ok := rs.baGuaTaTemplateMap[level]
	if !ok {
		return nil
	}
	return to
}

//通过层数获取八卦秘境补偿模板
func (rs *baGuaTemplateService) GetBaGuaMiJingBuChangTemplateByLevel(level int32) *gametemplate.BaGuaMiJingBuChangTemplate {
	to, ok := rs.baGuaMiJingBuChangTemplateMap[level]
	if !ok {
		return nil
	}
	return to
}

//八卦秘境助战按钮邀请成功的冷却时间(毫秒)
func (rs *baGuaTemplateService) GetInvitePairCdTime() int64 {
	return int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBaGuaMiJingInviteCd))
}

var (
	once sync.Once
	cs   *baGuaTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &baGuaTemplateService{}
		err = cs.init()
	})
	return err
}

func GetBaGuaTemplateService() BaGuaTemplateService {
	return cs
}
