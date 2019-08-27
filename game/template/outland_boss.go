package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	outlandbosstypes "fgame/fgame/game/outlandboss/types"
	"fmt"
)

func init() {
	template.Register((*OutlandBossTemplate)(nil))
}

type OutlandBossTemplate struct {
	*OutlandBossTemplateVO
	bossTemp *BiologyTemplate
}

func (t *OutlandBossTemplate) TemplateId() int {
	return t.Id
}

func (t *OutlandBossTemplate) FileName() string {
	return "tb_waiyu_boss.json"
}

//组合成需要的数据
func (t *OutlandBossTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检查有效性
func (t *OutlandBossTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//地图
	mto := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if mto == nil {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return err
	}
	_, ok := mto.(*MapTemplate)
	if !ok {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return
	}

	//怪物id
	bilogyTemp := template.GetTemplateService().Get(int(t.BiologyId), (*BiologyTemplate)(nil))
	if bilogyTemp == nil {
		err = fmt.Errorf("[%d] invalid", t.BiologyId)
		return template.NewTemplateFieldError("BiologyId", err)
	}
	bossTemp := bilogyTemp.(*BiologyTemplate)
	if bossTemp == nil {
		err = fmt.Errorf("[%d] invalid", t.BiologyId)
		return template.NewTemplateFieldError("BiologyId", err)
	}
	t.bossTemp = bossTemp

	//类型
	bossType := outlandbosstypes.OutlandBossType(t.Type)
	if !bossType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//假掉落记录出现概率
	err = validator.MinValidate(float64(t.Rate), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Rate", err)
	}

	return nil
}

//检验后组合
func (t *OutlandBossTemplate) PatchAfterCheck() {
}

func (t *OutlandBossTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.bossTemp
}
