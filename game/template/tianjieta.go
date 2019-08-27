package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"
	scenetypes "fgame/fgame/game/scene/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

//天劫塔配置
type TianJieTaTemplate struct {
	*TianJieTaTemplateVO
	battleAttrTemplate *AttrTemplate          //阶别属性
	rewItemMap         map[int32]int32        //奖励物品
	rewData            *propertytypes.RewData //奖励属性
	mapTemplate        *MapTemplate           //副本地图
}

func (tjtt *TianJieTaTemplate) TemplateId() int {
	return tjtt.Id
}

func (tjtt *TianJieTaTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return tjtt.battleAttrTemplate
}

func (tjtt *TianJieTaTemplate) GetRewItemMap() map[int32]int32 {
	return tjtt.rewItemMap
}

func (tjtt *TianJieTaTemplate) GetRewData() *propertytypes.RewData {
	return tjtt.rewData
}

func (tjtt *TianJieTaTemplate) GetMapTemplate() *MapTemplate {
	return tjtt.mapTemplate
}

func (tjtt *TianJieTaTemplate) GetRewQinMiDu() int32 {
	return tjtt.RewQinMiDu
}

func (tjtt *TianJieTaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tjtt.FileName(), tjtt.TemplateId(), err)
			return
		}
	}()

	//验证 Attr
	to := template.GetTemplateService().Get(int(tjtt.Attr), (*AttrTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", tjtt.Attr)
		return template.NewTemplateFieldError("Attr", err)
	}
	attrTemplate, _ := to.(*AttrTemplate)
	tjtt.battleAttrTemplate = attrTemplate
	//验证 map_id
	tempMapTemplate := template.GetTemplateService().Get(int(tjtt.MapId), (*MapTemplate)(nil))
	if tempMapTemplate == nil {
		err = fmt.Errorf("[%d] invalid", tjtt.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	tjtt.mapTemplate = tempMapTemplate.(*MapTemplate)

	//验证  rew_item
	if tjtt.RewItem != 0 {
		tjtt.rewItemMap = make(map[int32]int32)
		to = template.GetTemplateService().Get(int(tjtt.RewItem), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", tjtt.RewItem)
			return template.NewTemplateFieldError("RewItem", err)
		}

		//验证 rew_count
		err = validator.MinValidate(float64(tjtt.RewCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("RewCount", err)
			return
		}

		tjtt.rewItemMap[tjtt.RewItem] = tjtt.RewCount
	}

	//验证 rew_exp
	err = validator.MinValidate(float64(tjtt.RewExp), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewExp", err)
		return
	}

	//验证 rew_exp_point
	err = validator.MinValidate(float64(tjtt.RewExpPoint), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewExpPoint", err)
		return
	}

	//验证 rew_silver
	err = validator.MinValidate(float64(tjtt.RewSilver), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(tjtt.RewGold), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewGold", err)
		return
	}

	//验证 rew_bind_gold
	err = validator.MinValidate(float64(tjtt.RewBindGold), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewBindGold", err)
		return
	}

	if tjtt.RewExp > 0 || tjtt.RewExpPoint > 0 || tjtt.RewSilver > 0 || tjtt.RewGold > 0 || tjtt.RewBindGold > 0 {
		tjtt.rewData = propertytypes.CreateRewData(tjtt.RewExp, tjtt.RewExpPoint, tjtt.RewSilver, tjtt.RewGold, tjtt.RewBindGold)
	}

	return nil
}

func (tjtt *TianJieTaTemplate) PatchAfterCheck() {

}

func (tjtt *TianJieTaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tjtt.FileName(), tjtt.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if tjtt.NextId != 0 {
		to := template.GetTemplateService().Get(int(tjtt.NextId), (*TianJieTaTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", tjtt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		tjtto := to.(*TianJieTaTemplate)
		//验证level
		diffLevel := tjtto.Level - tjtt.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", tjtto.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证 level
	err = validator.MinValidate(float64(tjtt.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tjtt.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 skill_id
	tempSkillTemplate := template.GetTemplateService().Get(int(tjtt.SkillId), (*SkillTemplate)(nil))
	if tempSkillTemplate == nil {
		err = fmt.Errorf("[%d] invalid", tjtt.SkillId)
		return template.NewTemplateFieldError("SkillId", err)
	}
	skilltyp := tempSkillTemplate.(*SkillTemplate).GetSkillFirstType()
	if skilltyp != skilltypes.SkillFirstTypeRealm {
		err = fmt.Errorf("[%d] invalid", tjtt.SkillId)
		return template.NewTemplateFieldError("SkillId", err)
	}

	//map_id 类型
	mapType := tjtt.mapTemplate.GetMapType()
	if mapType != scenetypes.SceneTypeTianJieTa {
		err = fmt.Errorf("[%d]  invalid", tjtt.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}

	//验证 rew_qinmidu
	err = validator.MinValidate(float64(tjtt.RewQinMiDu), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tjtt.RewQinMiDu)
		return template.NewTemplateFieldError("RewQinMiDu", err)
	}

	return nil
}

func (tjtt *TianJieTaTemplate) FileName() string {
	return "tb_tianjieta.json"
}

func init() {
	template.Register((*TianJieTaTemplate)(nil))
}
