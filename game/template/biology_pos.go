package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fmt"
)

func init() {
	template.Register((*BiologyPosTemplate)(nil))
}

type BiologyPosTemplate struct {
	*BiologyPosTemplateVO
	pos      coretypes.Position //位置
	boTemp   *BiologyTemplate
	nextTemp *BiologyPosTemplate
}

func (t *BiologyPosTemplate) TemplateId() int {
	return t.Id
}

func (t *BiologyPosTemplate) GetNextTemp() *BiologyPosTemplate {
	return t.nextTemp
}

func (t *BiologyPosTemplate) GetPos() coretypes.Position {
	return t.pos
}

func (t *BiologyPosTemplate) GetBiologyTemp() *BiologyTemplate {
	return t.boTemp
}

func (t *BiologyPosTemplate) FileName() string {
	return "tb_biology_pos.json"
}

//组合成需要的数据
func (t *BiologyPosTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//怪物id
	if t.BiologyId > 0 {
		bilogyTemp := template.GetTemplateService().Get(int(t.BiologyId), (*BiologyTemplate)(nil))
		if bilogyTemp == nil {
			err = fmt.Errorf("BiologyId [%d] no exist", t.BiologyId)
			return err
		}
		bo, ok := bilogyTemp.(*BiologyTemplate)
		if !ok {
			err = fmt.Errorf("BiologyId [%d] no exist", t.BiologyId)
			return
		}

		t.boTemp = bo
	}

	//
	t.pos = coretypes.Position{
		X: t.PosX,
		Y: t.PosY,
		Z: t.PosZ,
	}

	//
	if t.NextId > 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*BiologyPosTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("BiologyId [%d] no exist", t.NextId)
			return err
		}
		t.nextTemp = to.(*BiologyPosTemplate)
	}

	return nil
}

//检查有效性
func (t *BiologyPosTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	to := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	mapTemplate := to.(*MapTemplate)
	mask := mapTemplate.GetMap().IsMask(t.pos.X, t.pos.Z)
	if !mask {
		err = fmt.Errorf("pos[%s] invalid", t.pos.String())
		err = template.NewTemplateFieldError("pos", err)
		return
	}
	y := mapTemplate.GetMap().GetHeight(t.pos.X, t.pos.Z)
	t.pos.Y = y

	return nil
}

//检验后组合
func (t *BiologyPosTemplate) PatchAfterCheck() {
}
