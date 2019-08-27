package template

import (
	"fgame/fgame/core/template"
	"fmt"
)

func init() {
	template.Register((*ZhenXiBossTemplate)(nil))
}

type ZhenXiBossTemplate struct {
	*ZhenXiBossTemplateVO
	mapTemplate *MapTemplate
	bossTemp    *BiologyTemplate
}

func (t *ZhenXiBossTemplate) TemplateId() int {
	return t.Id
}

func (t *ZhenXiBossTemplate) FileName() string {
	return "tb_zhenxi_boss.json"
}

//组合成需要的数据
func (t *ZhenXiBossTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检查有效性
func (t *ZhenXiBossTemplate) Check() (err error) {
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
	mapTemplate, ok := mto.(*MapTemplate)
	if !ok {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return
	}
	t.mapTemplate = mapTemplate
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
func (t *ZhenXiBossTemplate) PatchAfterCheck() {
}

func (t *ZhenXiBossTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.bossTemp
}

func (t *ZhenXiBossTemplate) GetMapTemplate() *MapTemplate {
	return t.mapTemplate
}

func (t *ZhenXiBossTemplate) GetBiologyId() int32 {
	return t.BiologyId
}
func (t *ZhenXiBossTemplate) GetMapId() int32 {
	return t.MapId
}

func (t *ZhenXiBossTemplate) GetRecForce() int64 {
	return int64(t.RecForce)
}
