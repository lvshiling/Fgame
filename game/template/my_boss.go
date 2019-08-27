package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	mybosstypes "fgame/fgame/game/myboss/types"
	"fmt"
)

func init() {
	template.Register((*MyBossTemplate)(nil))
}

type MyBossTemplate struct {
	*MyBossTemplateVO
	bossType mybosstypes.MyBossType
	bossTemp *BiologyTemplate
}

func (t *MyBossTemplate) TemplateId() int {
	return t.Id
}

func (t *MyBossTemplate) GetBossType() mybosstypes.MyBossType {
	return t.bossType
}

func (t *MyBossTemplate) FileName() string {
	return "tb_geren_boss.json"
}

//组合成需要的数据
func (t *MyBossTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检查有效性
func (t *MyBossTemplate) Check() (err error) {
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

	// vip等级
	err = validator.MinValidate(float64(t.NeedVipLevel), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("NeedVipLevel", err)
	}

	// 免费次数
	err = validator.MinValidate(float64(t.FreeTimes), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("FreeTimes", err)
	}

	// 总次数
	err = validator.MinValidate(float64(t.TimesCount), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("TimesCount", err)
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
	t.bossType = mybosstypes.MyBossType(t.Type)
	if !t.bossType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	return nil
}

//检验后组合
func (t *MyBossTemplate) PatchAfterCheck() {
}

func (t *MyBossTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.bossTemp
}
