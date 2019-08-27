package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/utils"
	"fmt"
)

func init() {
	template.Register((*GuideReplicaRescureTemplate)(nil))
}

type GuideReplicaRescureTemplate struct {
	*GuideReplicaRescureTemplateVO
	dropIdList    []int32
	herbsBuffTemp *BuffTemplate
}

func (t *GuideReplicaRescureTemplate) TemplateId() int {
	return t.Id
}

func (t *GuideReplicaRescureTemplate) GetDropId() []int32 {
	return t.dropIdList
}

func (t *GuideReplicaRescureTemplate) GetHerbsBuffTemplate() *BuffTemplate {
	return t.herbsBuffTemp
}

func (t *GuideReplicaRescureTemplate) FileName() string {
	return "tb_daily_jiuyuan.json"
}

//组合成需要的数据
func (t *GuideReplicaRescureTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	t.dropIdList, err = utils.SplitAsIntArray(t.DropId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.DropId)
		return template.NewTemplateFieldError("DropId", err)
	}

	return nil
}

//检查有效性
func (t *GuideReplicaRescureTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//怪物id
	if t.RescureBiologyId > 0 {
		bilogyTemp := template.GetTemplateService().Get(int(t.RescureBiologyId), (*BiologyTemplate)(nil))
		if bilogyTemp == nil {
			err = fmt.Errorf("BiologyId [%d] no exist", t.RescureBiologyId)
			return err
		}
		_, ok := bilogyTemp.(*BiologyTemplate)
		if !ok {
			err = fmt.Errorf("[%s] invalid", t.RescureBiologyId)
			return template.NewTemplateFieldError("RescureBiologyId", err)
		}
	}
	//怪物id
	if t.CollectBiologyId > 0 {
		bilogyTemp := template.GetTemplateService().Get(int(t.CollectBiologyId), (*BiologyTemplate)(nil))
		if bilogyTemp == nil {
			err = fmt.Errorf("BiologyId [%d] no exist", t.CollectBiologyId)
			return err
		}
		_, ok := bilogyTemp.(*BiologyTemplate)
		if !ok {
			err = fmt.Errorf("BiologyId [%d] no exist", t.CollectBiologyId)
			return
		}
	}

	buffTemplateVO := template.GetTemplateService().Get(int(t.BuffId), (*BuffTemplate)(nil))
	if buffTemplateVO == nil {
		err = fmt.Errorf("[%d] invalid", t.BuffId)
		err = template.NewTemplateFieldError("BuffId", err)
		return
	}
	t.herbsBuffTemp = buffTemplateVO.(*BuffTemplate)

	return nil
}

//检验后组合
func (t *GuideReplicaRescureTemplate) PatchAfterCheck() {
}
