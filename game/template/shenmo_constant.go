package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/scene/types"
	"fmt"
)

//神魔常量配置
type ShenMoConstantTemplate struct {
	*ShenMoConstantTemplateVO
}

func (t *ShenMoConstantTemplate) IsShenMoCollect(biologyId int32) bool {
	if t.WuZiBiologyId != biologyId && t.DaQiBiologyId != biologyId {
		return false
	}

	return true
}

func (t *ShenMoConstantTemplate) GetCollectPoint(biologyId int32) (addGongXunNum, addJiFenNum int32) {
	if t.WuZiBiologyId == biologyId {
		addGongXunNum = t.WuZiGetGongXun
		addJiFenNum = t.WuZiGeiJiFen
	}

	if t.DaQiBiologyId == biologyId {
		addGongXunNum = t.DaQiGetGongXun
		addJiFenNum = t.DaQiGetJiFen
	}

	return
}

func (t *ShenMoConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *ShenMoConstantTemplate) PatchAfterCheck() {

}

func (t *ShenMoConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShenMoConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	wuZiTemplate := template.GetTemplateService().Get(int(t.WuZiBiologyId), (*BiologyTemplate)(nil))
	if wuZiTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.WuZiBiologyId)
		err = template.NewTemplateFieldError("WuZiBiologyId", err)
		return
	}
	wuZiBiology := wuZiTemplate.(*BiologyTemplate)
	if wuZiBiology.GetBiologyScriptType() != types.BiologyScriptTypeGeneralCollect {
		err = fmt.Errorf("[%d] invalid", t.WuZiBiologyId)
		err = template.NewTemplateFieldError("WuZiBiologyId", err)
		return
	}

	daQiTemplate := template.GetTemplateService().Get(int(t.DaQiBiologyId), (*BiologyTemplate)(nil))
	if wuZiTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.DaQiBiologyId)
		err = template.NewTemplateFieldError("DaQiBiologyId", err)
		return
	}
	daQiBiology := daQiTemplate.(*BiologyTemplate)
	if daQiBiology.GetBiologyScriptType() != types.BiologyScriptTypeGeneralCollect {
		err = fmt.Errorf("[%d] invalid", t.DaQiBiologyId)
		err = template.NewTemplateFieldError("DaQiBiologyId", err)
		return
	}

	//验证 wuzi_get_gongxun
	err = validator.MinValidate(float64(t.WuZiGetGongXun), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.WuZiGetGongXun)
		err = template.NewTemplateFieldError("WuZiGetGongXun", err)
		return
	}

	//验证 wuzi_gei_jifen
	err = validator.MinValidate(float64(t.WuZiGeiJiFen), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.WuZiGeiJiFen)
		err = template.NewTemplateFieldError("WuZiGeiJiFen", err)
		return
	}

	//验证 daqi_get_gongxun
	err = validator.MinValidate(float64(t.DaQiGetGongXun), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DaQiGetGongXun)
		err = template.NewTemplateFieldError("DaQiGetGongXun", err)
		return
	}

	//验证 daqi_get_jifen
	err = validator.MinValidate(float64(t.DaQiGetJiFen), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DaQiGetJiFen)
		err = template.NewTemplateFieldError("DaQiGetJiFen", err)
		return
	}

	//验证 player_limit_count
	err = validator.MinValidate(float64(t.PlayerLimitCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.PlayerLimitCount)
		err = template.NewTemplateFieldError("PlayerLimitCount", err)
		return
	}

	//验证 reborn_buff
	buffTemplate := template.GetTemplateService().Get(int(t.RebornBuff), (*BuffTemplate)(nil))
	if buffTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.RebornBuff)
		err = template.NewTemplateFieldError("RebornBuff", err)
		return
	}

	//验证 map_id
	to := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}

	mapTemplate := to.(*MapTemplate)
	if mapTemplate.GetMapType() != types.SceneTypeCrossShenMo {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}

	return nil
}

func (t *ShenMoConstantTemplate) FileName() string {
	return "tb_shenmo_constant.json"
}

func init() {
	template.Register((*ShenMoConstantTemplate)(nil))
}
