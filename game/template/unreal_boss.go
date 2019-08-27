package template

import (
	"fgame/fgame/core/template"
	unrealbosstypes "fgame/fgame/game/unrealboss/types"
	"fmt"
)

func init() {
	template.Register((*UnrealBossTemplate)(nil))
}

type UnrealBossTemplate struct {
	*UnrealBossTemplateVO
	bossTemp *BiologyTemplate
}

func (t *UnrealBossTemplate) TemplateId() int {
	return t.Id
}

func (t *UnrealBossTemplate) FileName() string {
	return "tb_huanjing_boss.json"
}

//组合成需要的数据
func (t *UnrealBossTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检查有效性
func (t *UnrealBossTemplate) Check() (err error) {
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
	bossType := unrealbosstypes.UnrealBossType(t.Type)
	if !bossType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	return nil
}

//检验后组合
func (t *UnrealBossTemplate) PatchAfterCheck() {
}

func (t *UnrealBossTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.bossTemp
}
