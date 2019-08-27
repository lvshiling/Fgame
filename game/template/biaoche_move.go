package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fmt"
)

//镖车配置
type BiaocheMoveTemplate struct {
	*BiaocheMoveTemplateVO
	nextTem *BiaocheMoveTemplate
	pos     coretypes.Position
}

func (t *BiaocheMoveTemplate) TemplateId() int {
	return t.Id
}

func (t *BiaocheMoveTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//目标地
	t.pos = coretypes.Position{
		X: t.PosX,
		Y: t.PosY,
		Z: t.PosZ,
	}

	return nil
}

func (t *BiaocheMoveTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//mapId
	mto := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if mto == nil {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return err
	}
	mapTem, ok := mto.(*MapTemplate)
	if !ok {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return
	}

	//x、y、z
	if !mapTem.GetMap().IsMask(t.PosX, t.PosZ) {
		err = fmt.Errorf("[%.2f] [%2.f] invalid", t.PosX, t.PosZ)
		return template.NewTemplateFieldError("pos", err)
	}

	//nextId
	if t.NextId != 0 {
		tem := template.GetTemplateService().Get(int(t.NextId), (*BiaocheMoveTemplate)(nil))
		if tem == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return err
		}
		t.nextTem = tem.(*BiaocheMoveTemplate)
	}

	return nil
}

func (t *BiaocheMoveTemplate) PatchAfterCheck() {
}

func (t *BiaocheMoveTemplate) FileName() string {
	return "tb_biaoche_move.json"
}

func (t *BiaocheMoveTemplate) GetPosition() coretypes.Position {
	return t.pos
}

func (t *BiaocheMoveTemplate) GetNextTemp() *BiaocheMoveTemplate {
	return t.nextTem
}

func init() {
	template.Register((*BiaocheMoveTemplate)(nil))
}
