package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/tulong/types"
	"fmt"
)

//屠龙出生点配置
type TuLongPosTemplate struct {
	*TuLongPosTemplateVO
	posType types.TuLongPosType
	pos     coretypes.Position //位置
}

func (tl *TuLongPosTemplate) TemplateId() int {
	return tl.Id
}

func (tl *TuLongPosTemplate) GetPos() coretypes.Position {
	return tl.pos
}

func (tl *TuLongPosTemplate) GetPosType() types.TuLongPosType {
	return tl.posType
}

func (tl *TuLongPosTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tl.FileName(), tl.TemplateId(), err)
			return
		}
	}()

	tl.pos = coretypes.Position{
		X: tl.PosX,
		Y: tl.PosY,
		Z: tl.PosZ,
	}

	return nil
}

func (tl *TuLongPosTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tl.FileName(), tl.TemplateId(), err)
			return
		}
	}()

	tl.posType = types.TuLongPosType(tl.Type)
	if !tl.posType.Valid() {
		err = fmt.Errorf("[%d] invalid", tl.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	to := template.GetTemplateService().Get(int(tl.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", tl.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	mapTemplate := to.(*MapTemplate)

	if mapTemplate.GetMapType() != scenetypes.SceneTypeCrossTuLong {
		err = fmt.Errorf("[%d] invalid", tl.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	mask := mapTemplate.GetMap().IsMask(tl.pos.X, tl.pos.Z)
	if !mask {
		err = fmt.Errorf("pos[%s] invalid", tl.pos.String())
		err = template.NewTemplateFieldError("pos", err)
		return
	}
	y := mapTemplate.GetMap().GetHeight(tl.pos.X, tl.pos.Z)
	tl.pos.Y = y

	return nil
}
func (tl *TuLongPosTemplate) PatchAfterCheck() {

}
func (tl *TuLongPosTemplate) FileName() string {
	return "tb_tulong_pos.json"
}

func init() {
	template.Register((*TuLongPosTemplate)(nil))
}
