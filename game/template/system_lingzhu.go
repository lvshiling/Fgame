package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fmt"
)

type SystemLingZhuTemplate struct {
	*SystemLingZhuTemplateVO
	lingZhuType                additionsystypes.LingZhuType
	firstLevelTemplate         *SystemLingZhuUpLevelTemplate
	strengthenLevelTemplateMap map[int32]*SystemLingZhuUpLevelTemplate //索引为等级
}

func (t *SystemLingZhuTemplate) TemplateId() int {
	return t.Id
}

func (t *SystemLingZhuTemplate) GetLevelTemplate(level int32) *SystemLingZhuUpLevelTemplate {
	return t.strengthenLevelTemplateMap[level]
}

func (t *SystemLingZhuTemplate) GetLingZhuType() additionsystypes.LingZhuType {
	return t.lingZhuType
}

//检查有效性
func (t *SystemLingZhuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//灵童ID
	err = validator.MinValidate(float64(t.LingtongId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingtongId)
		return template.NewTemplateFieldError("LingtongId", err)
	}

	//灵珠类型
	err = validator.MinValidate(float64(t.Type), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	t.lingZhuType = additionsystypes.LingZhuType(t.Type)
	if !t.lingZhuType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//消耗物品ID
	err = validator.MinValidate(float64(t.UseItemId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseItemId)
		return template.NewTemplateFieldError("UseItemId", err)
	}

	//升级起始Id
	err = validator.MinValidate(float64(t.LevelBegin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelBegin)
		return template.NewTemplateFieldError("LevelBegin", err)
	}

	//关联升级表
	temp := template.GetTemplateService().Get(int(t.LevelBegin), (*SystemLingZhuUpLevelTemplate)(nil))
	levelTemp, _ := temp.(*SystemLingZhuUpLevelTemplate)
	if levelTemp == nil {
		err = fmt.Errorf("SystemLingZhuUpLevelTemplate[%d] invalid", t.LevelBegin)
		err = template.NewTemplateFieldError("MagicConditionParameter", err)
	}
	t.firstLevelTemplate = levelTemp

	return
}

//组合成需要的数据
func (t *SystemLingZhuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	return
}

//检验后组合
func (t *SystemLingZhuTemplate) PatchAfterCheck() {
	t.strengthenLevelTemplateMap = make(map[int32]*SystemLingZhuUpLevelTemplate)
	for temp := t.firstLevelTemplate; temp != nil; temp = temp.GetNextStrengthenTemplate() {
		t.strengthenLevelTemplateMap[temp.Level] = temp
	}
}

func (t *SystemLingZhuTemplate) FileName() string {
	return "tb_lingtong_lingzhu.json"
}

func init() {
	template.Register((*SystemLingZhuTemplate)(nil))
}
