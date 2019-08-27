package template

import (
	"fgame/fgame/core/template"
	"fmt"
)

func init() {
	template.Register((*ShengShouBossTemplate)(nil))
}

type ShengShouBossTemplate struct {
	*ShengShouBossTemplateVO

	bossTemp *BiologyTemplate
}

func (t *ShengShouBossTemplate) TemplateId() int {
	return t.Id
}

func (t *ShengShouBossTemplate) FileName() string {
	return "tb_shengshou_boss.json"
}

//组合成需要的数据
func (t *ShengShouBossTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检查有效性
func (t *ShengShouBossTemplate) Check() (err error) {
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

	return nil
}

//检验后组合
func (t *ShengShouBossTemplate) PatchAfterCheck() {
}

func (t *ShengShouBossTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.bossTemp
}

func (t *ShengShouBossTemplate) GetBiologyId() int32 {
	return t.BiologyId
}
func (t *ShengShouBossTemplate) GetMapId() int32 {
	return t.MapId
}

func (t *ShengShouBossTemplate) GetRecForce() int64 {
	return t.RecForce
}
