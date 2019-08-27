package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fmt"
)

type GodCastingCastingSpiritTemplate struct {
	*GodCastingCastingSpiritTemplateVO
	firstLevelTemplate         *GodCastingCastingSpiritLevelTemplate
	strengthenLevelTemplateMap map[int32]*GodCastingCastingSpiritLevelTemplate //索引为等级
}

func (t *GodCastingCastingSpiritTemplate) TemplateId() int {
	return t.Id
}

//该铸灵是否开启
func (t *GodCastingCastingSpiritTemplate) IsActive(godCastingLevel int32) bool {
	if godCastingLevel >= t.NeedShenzhuLevel {
		return true
	} else {
		return false
	}
}

func (t *GodCastingCastingSpiritTemplate) GetLevelTemplate(level int32) *GodCastingCastingSpiritLevelTemplate {
	return t.strengthenLevelTemplateMap[level]
}

func (t *GodCastingCastingSpiritTemplate) GetBodyPos() inventorytypes.BodyPositionType {
	return inventorytypes.BodyPositionType(t.SubType)
}

func (t *GodCastingCastingSpiritTemplate) GetSpiritType() goldequiptypes.SpiritType {
	return goldequiptypes.SpiritType(t.ZhulingType)
}

//检查有效性
func (t *GodCastingCastingSpiritTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//铸灵部位
	err = validator.MinValidate(float64(t.SubType), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.SubType)
		return template.NewTemplateFieldError("SubType", err)
	}

	//铸灵类型
	err = validator.MinValidate(float64(t.ZhulingType), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhulingType)
		return template.NewTemplateFieldError("ZhulingType", err)
	}

	//神铸解锁需要等级
	err = validator.MinValidate(float64(t.NeedShenzhuLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedShenzhuLevel)
		return template.NewTemplateFieldError("NeedShenzhuLevel", err)
	}

	//消耗物品ID
	err = validator.MinValidate(float64(t.UseItemId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseItemId)
		return template.NewTemplateFieldError("UseItemId", err)
	}

	//神铸等级连接相关ID
	err = validator.MinValidate(float64(t.ZhulingLevelBegin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhulingLevelBegin)
		return template.NewTemplateFieldError("ZhulingLevelBegin", err)
	}

	//关联升级表
	temp := template.GetTemplateService().Get(int(t.ZhulingLevelBegin), (*GodCastingCastingSpiritLevelTemplate)(nil))
	levelTemp, _ := temp.(*GodCastingCastingSpiritLevelTemplate)
	if levelTemp == nil {
		err = fmt.Errorf("GodCastingCastingSpiritTemplate[%d] invalid", t.ZhulingLevelBegin)
		err = template.NewTemplateFieldError("MagicConditionParameter", err)
	}
	t.firstLevelTemplate = levelTemp

	return
}

//组合成需要的数据
func (t *GodCastingCastingSpiritTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	return
}

//检验后组合
func (t *GodCastingCastingSpiritTemplate) PatchAfterCheck() {
	t.strengthenLevelTemplateMap = make(map[int32]*GodCastingCastingSpiritLevelTemplate)
	for temp := t.firstLevelTemplate; temp != nil; temp = temp.GetNextStrengthenTemplate() {
		t.strengthenLevelTemplateMap[temp.Level] = temp
	}
}

func (t *GodCastingCastingSpiritTemplate) FileName() string {
	return "tb_shenzhuequip_zhuling.json"
}

func init() {
	template.Register((*GodCastingCastingSpiritTemplate)(nil))
}
