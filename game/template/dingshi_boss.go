package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	"fmt"
)

func init() {
	template.Register((*DingShiBossTemplate)(nil))
}

type DingShiBossTemplate struct {
	*DingShiBossTemplateVO

	bossTemp    *BiologyTemplate
	gongGaoList []int32
}

func (t *DingShiBossTemplate) TemplateId() int {
	return t.Id
}

func (t *DingShiBossTemplate) FileName() string {
	return "tb_dingshi_boss.json"
}

func (t *DingShiBossTemplate) GetBossThreshold(oldPercent int32, newPercent int32) (percent int32, flag bool) {

	for _, threshold := range t.gongGaoList {
		if !flag {
			percent = threshold
		}
		if threshold >= newPercent && threshold < oldPercent {
			flag = true
			if threshold <= percent {
				percent = threshold
			}
		}
	}
	return
}

//组合成需要的数据
func (t *DingShiBossTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.gongGaoList, err = coreutils.SplitAsIntArray(t.Gonggao)
	if err != nil {
		return
	}
	return nil
}

//检查有效性
func (t *DingShiBossTemplate) Check() (err error) {
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

	previousGongGao := int32(100)
	for _, gongGao := range t.gongGaoList {
		err = validator.MaxValidate(float64(gongGao), float64(previousGongGao), false)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.Gonggao)
			return template.NewTemplateFieldError("gonggao", err)
		}
		previousGongGao = gongGao
	}

	err = validator.MinValidate(float64(previousGongGao), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Gonggao)
		return template.NewTemplateFieldError("gonggao", err)
	}
	return nil
}

//检验后组合
func (t *DingShiBossTemplate) PatchAfterCheck() {
}

func (t *DingShiBossTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.bossTemp
}
