package template

import (
	"fgame/fgame/core/template"
	"fmt"
)

func init() {
	template.Register((*GuideReplicaMoJianTemplate)(nil))
}

type GuideReplicaMoJianTemplate struct {
	*GuideReplicaMoJianTemplateVO
	bossTemp     *BiologyTemplate
	retinueTemp  *BiologyTemplate
	wuDiBuffTemp *BuffTemplate
}

func (t *GuideReplicaMoJianTemplate) TemplateId() int {
	return t.Id
}

func (t *GuideReplicaMoJianTemplate) GetBossTemplate() *BiologyTemplate {
	return t.bossTemp
}

func (t *GuideReplicaMoJianTemplate) GetRetinueTemplate() *BiologyTemplate {
	return t.retinueTemp
}

func (t *GuideReplicaMoJianTemplate) GetWuDiBuffTemplate() *BuffTemplate {
	return t.wuDiBuffTemp
}

func (t *GuideReplicaMoJianTemplate) FileName() string {
	return "tb_daily_mojian.json"
}

//组合成需要的数据
func (t *GuideReplicaMoJianTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

//检查有效性
func (t *GuideReplicaMoJianTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//boss怪物id
	bossTempeVO := template.GetTemplateService().Get(int(t.BossId), (*BiologyTemplate)(nil))
	if bossTempeVO == nil {
		err = fmt.Errorf("BiologyId [%d] no exist", t.BossId)
		return err
	}
	t.bossTemp = bossTempeVO.(*BiologyTemplate)

	//随从怪物id
	retinueTempeVO := template.GetTemplateService().Get(int(t.RetinueId), (*BiologyTemplate)(nil))
	if retinueTempeVO == nil {
		err = fmt.Errorf("BiologyId2 [%d] no exist", t.RetinueId)
		return err
	}
	t.retinueTemp = retinueTempeVO.(*BiologyTemplate)

	buffTemplateVO := template.GetTemplateService().Get(int(t.BuffId), (*BuffTemplate)(nil))
	if buffTemplateVO == nil {
		err = fmt.Errorf("[%d] invalid", t.BuffId)
		err = template.NewTemplateFieldError("BuffId", err)
		return
	}
	t.wuDiBuffTemp = buffTemplateVO.(*BuffTemplate)

	return nil
}

//检验后组合
func (t *GuideReplicaMoJianTemplate) PatchAfterCheck() {
}
