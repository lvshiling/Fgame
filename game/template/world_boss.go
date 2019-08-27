package template

import (
	"fgame/fgame/core/template"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fmt"
)

func init() {
	template.Register((*WorldBossTemplate)(nil))
}

type WorldBossTemplate struct {
	*WorldBossTemplateVO
	bossType worldbosstypes.WorldBossType
	bossTemp *BiologyTemplate
}

func (t *WorldBossTemplate) TemplateId() int {
	return t.Id
}

func (t *WorldBossTemplate) GetBossType() worldbosstypes.WorldBossType {
	return t.bossType
}

func (t *WorldBossTemplate) FileName() string {
	return "tb_world_boss.json"
}

//组合成需要的数据
func (t *WorldBossTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检查有效性
func (t *WorldBossTemplate) Check() (err error) {
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
	t.bossType = worldbosstypes.WorldBossType(t.Type)
	if !t.bossType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	return nil
}

//检验后组合
func (t *WorldBossTemplate) PatchAfterCheck() {
}

// func (t *WorldBossTemplate) GetRewItemMap() map[int32]int32 {
// 	return t.rewItemMap
// }

func (t *WorldBossTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.bossTemp
}
func (t *WorldBossTemplate) GetBiologyId() int32 {
	return t.BiologyId
}
func (t *WorldBossTemplate) GetMapId() int32 {
	return t.MapId
}

func (t *WorldBossTemplate) GetRecForce() int64 {
	return t.RecForce
}
