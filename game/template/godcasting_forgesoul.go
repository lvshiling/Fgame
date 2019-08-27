package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fmt"
)

type GodCastingForgeSoulTemplate struct {
	*GodCastingForgeSoulTemplateVO
	firstLevelTemplate         *GodCastingForgeSoulLevelTemplate
	strengthenLevelTemplateMap map[int32]*GodCastingForgeSoulLevelTemplate //索引为等级
	teshuSkillTemp             *TeShuSkillTemplate
}

func (t *GodCastingForgeSoulTemplate) TemplateId() int {
	return t.Id
}

func (t *GodCastingForgeSoulTemplate) GetTeshuSkillTemp() *TeShuSkillTemplate {
	return t.teshuSkillTemp
}

func (t *GodCastingForgeSoulTemplate) GetLevelTemplate(level int32) *GodCastingForgeSoulLevelTemplate {
	return t.strengthenLevelTemplateMap[level]
}

func (t *GodCastingForgeSoulTemplate) GetBodyPos() inventorytypes.BodyPositionType {
	return inventorytypes.BodyPositionType(t.SubType)
}

func (t *GodCastingForgeSoulTemplate) GetSoulType() goldequiptypes.ForgeSoulType {
	return goldequiptypes.ForgeSoulType(t.Type)
}

//检查有效性
func (t *GodCastingForgeSoulTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//锻魂部位
	err = validator.MinValidate(float64(t.SubType), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.SubType)
		return template.NewTemplateFieldError("SubType", err)
	}

	//锻魂类型
	err = validator.MinValidate(float64(t.Type), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//消耗物品ID
	err = validator.MinValidate(float64(t.UseItemId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseItemId)
		return template.NewTemplateFieldError("UseItemId", err)
	}

	//锻魂等级连接相关ID
	err = validator.MinValidate(float64(t.LevelBeginId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelBeginId)
		return template.NewTemplateFieldError("LevelBeginId", err)
	}
	//关联升级表
	temp := template.GetTemplateService().Get(int(t.LevelBeginId), (*GodCastingForgeSoulLevelTemplate)(nil))
	levelTemp, _ := temp.(*GodCastingForgeSoulLevelTemplate)
	if levelTemp == nil {
		err = fmt.Errorf("GodCastingForgeSoulTemplate[%d] invalid", t.LevelBeginId)
		err = template.NewTemplateFieldError("MagicConditionParameter", err)
	}
	t.firstLevelTemplate = levelTemp

	//特殊技能表连接相关ID
	err = validator.MinValidate(float64(t.TeshuSkillId), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TeshuSkillId)
		return template.NewTemplateFieldError("TeshuSkillId", err)
	}
	//关联特殊技能表
	temp = template.GetTemplateService().Get(int(t.TeshuSkillId), (*TeShuSkillTemplate)(nil))
	teshuSkillTemp, _ := temp.(*TeShuSkillTemplate)
	if teshuSkillTemp == nil {
		err = fmt.Errorf("GodCastingForgeSoulTemplate[%d] invalid", t.TeshuSkillId)
		err = template.NewTemplateFieldError("TeshuSkillId", err)
		return
	}
	t.teshuSkillTemp = teshuSkillTemp

	return
}

//组合成需要的数据
func (t *GodCastingForgeSoulTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	return
}

//检验后组合
func (t *GodCastingForgeSoulTemplate) PatchAfterCheck() {
	t.strengthenLevelTemplateMap = make(map[int32]*GodCastingForgeSoulLevelTemplate)
	for temp := t.firstLevelTemplate; temp != nil; temp = temp.GetNextStrengthenTemplate() {
		t.strengthenLevelTemplateMap[temp.Level] = temp
	}
}

func (t *GodCastingForgeSoulTemplate) FileName() string {
	return "tb_shenzhuequip_duanhun.json"
}

func init() {
	template.Register((*GodCastingForgeSoulTemplate)(nil))
}
