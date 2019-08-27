package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//龙宫探宝常量配置
type LongGongConstantTemplate struct {
	*LongGongConstantTemplateVO
	longGongMapTemplate *MapTemplate
	bossBiology         *BiologyTemplate   //boss生物模板
	collectBiology      *BiologyTemplate   //采集点生物模板
	pos                 coretypes.Position //位置
}

func (t *LongGongConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *LongGongConstantTemplate) GetMapTemplate() *MapTemplate {
	return t.longGongMapTemplate
}

func (t *LongGongConstantTemplate) GetPos() coretypes.Position {
	return t.pos
}

func (t *LongGongConstantTemplate) GetBossBiology() *BiologyTemplate {
	return t.bossBiology
}

func (t *LongGongConstantTemplate) GetCollectBiology() *BiologyTemplate {
	return t.collectBiology
}

func (t *LongGongConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	tempMapTeamplate := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if tempMapTeamplate == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	t.longGongMapTemplate = tempMapTeamplate.(*MapTemplate)

	t.pos = coretypes.Position{
		X: t.PosX,
		Y: t.PosY,
		Z: t.PosZ,
	}

	//boss生物
	bossTemplate := template.GetTemplateService().Get(int(t.BossId), (*BiologyTemplate)(nil))
	if bossTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.BossId)
		return template.NewTemplateFieldError("BossId", err)
	}
	t.bossBiology = bossTemplate.(*BiologyTemplate)

	//collect生物
	collectTemplate := template.GetTemplateService().Get(int(t.BossBeKillCaiJiBiologyId), (*BiologyTemplate)(nil))
	if collectTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.BossBeKillCaiJiBiologyId)
		return template.NewTemplateFieldError("BossBeKillCaiJiBiologyId", err)
	}
	t.collectBiology = collectTemplate.(*BiologyTemplate)

	return nil
}

func (t *LongGongConstantTemplate) PatchAfterCheck() {

}

func (t *LongGongConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	err = validator.MinValidate(float64(t.BossNeedCaiJiCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BossNeedCaiJiCount)
		return template.NewTemplateFieldError("BossNeedCaiJiCount", err)
	}

	err = validator.MinValidate(float64(t.BossBeKillCaiJiPersonalCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BossBeKillCaiJiPersonalCount)
		return template.NewTemplateFieldError("BossBeKillCaiJiPersonalCount", err)
	}
	//boss生物验证
	if t.bossBiology.GetBiologyScriptType() != scenetypes.BiologyScriptTypeLongGongBoss {
		err = fmt.Errorf("[%d] invalid", t.BossId)
		return template.NewTemplateFieldError("BossId", err)
	}

	//财宝采集物验证
	if t.collectBiology.GetBiologyScriptType() != scenetypes.BiologyScriptTypeLongGongTreasure {
		err = fmt.Errorf("[%d] invalid", t.BossBeKillCaiJiBiologyId)
		return template.NewTemplateFieldError("BossBeKillCaiJiBiologyId", err)
	}

	//验证系统采集珍珠间隔时间
	err = validator.MinValidate(float64(t.XuJiaCaiJiAddTime), float64(1000), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XuJiaCaiJiAddTime)
		return template.NewTemplateFieldError("XuJiaCaiJiAddTime", err)
	}

	err = validator.MinValidate(float64(t.XuJiaCaiJiAddCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.XuJiaCaiJiAddCount)
		return template.NewTemplateFieldError("XuJiaCaiJiAddCount", err)
	}

	//验证地图类型
	if t.longGongMapTemplate.GetMapType() != scenetypes.SceneTypeLongGong {
		err = fmt.Errorf("mapid scene type invalid")
		return template.NewTemplateFieldError("mapid", err)
	}

	//验证pos
	if !t.longGongMapTemplate.GetMap().IsMask(t.pos.X, t.pos.Z) {
		err = fmt.Errorf("boss born pos invalid")
		return template.NewTemplateFieldError("BossBornPos", err)
	}

	return nil
}

func (t *LongGongConstantTemplate) FileName() string {
	return "tb_longgong_constant.json"
}

func init() {
	template.Register((*LongGongConstantTemplate)(nil))
}
