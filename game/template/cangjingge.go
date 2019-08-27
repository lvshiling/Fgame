package template

import (
	"fgame/fgame/core/template"
	cangjinggetypes "fgame/fgame/game/cangjingge/types"
	"fmt"
)

func init() {
	template.Register((*CangJingGeTemplate)(nil))
}

type CangJingGeTemplate struct {
	*CangJingGeTemplateVO
	bossType cangjinggetypes.CangJingGeBossType
	bossTemp *BiologyTemplate
}

func (t *CangJingGeTemplate) TemplateId() int {
	return t.Id
}

func (t *CangJingGeTemplate) GetBossType() cangjinggetypes.CangJingGeBossType {
	return t.bossType
}

func (t *CangJingGeTemplate) FileName() string {
	return "tb_cangjingge.json"
}

//组合成需要的数据
func (t *CangJingGeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检查有效性
func (t *CangJingGeTemplate) Check() (err error) {
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
	t.bossType = cangjinggetypes.CangJingGeBossType(t.Type)
	if !t.bossType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	return nil
}

//检验后组合
func (t *CangJingGeTemplate) PatchAfterCheck() {
}

// func (t *CangJingGeTemplate) GetRewItemMap() map[int32]int32 {
// 	return t.rewItemMap
// }

func (t *CangJingGeTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.bossTemp
}
