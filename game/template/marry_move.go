package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//婚车移动配置
type MarryMoveTemplate struct {
	*MarryMoveTemplateVO
	nextTem *MarryMoveTemplate
	pos     coretypes.Position //位置
}

func (mmt *MarryMoveTemplate) TemplateId() int {
	return mmt.Id
}

func (mmt *MarryMoveTemplate) GetPos() coretypes.Position {
	return mmt.pos
}

func (mmt *MarryMoveTemplate) GetNextTemp() *MarryMoveTemplate {
	return mmt.nextTem
}

func (mmt *MarryMoveTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mmt.FileName(), mmt.TemplateId(), err)
			return
		}
	}()

	mmt.pos = coretypes.Position{
		X: mmt.PosX,
		Y: mmt.PosY,
		Z: mmt.PosZ,
	}

	if mmt.NextId != 0 {
		diff := mmt.NextId - int32(mmt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", mmt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		to := template.GetTemplateService().Get(int(mmt.NextId), (*MarryMoveTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mmt.MapId)
			return template.NewTemplateFieldError("MapId", err)
		}
		mmt.nextTem = to.(*MarryMoveTemplate)
	}

	return nil
}

func (mmt *MarryMoveTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mmt.FileName(), mmt.TemplateId(), err)
			return
		}
	}()

	to := template.GetTemplateService().Get(int(mmt.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", mmt.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	mapTemplate := to.(*MapTemplate)

	if mapTemplate.GetMapType() != scenetypes.SceneTypeWorld {
		err = fmt.Errorf("[%d] invalid", mmt.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	mask := mapTemplate.GetMap().IsMask(mmt.pos.X, mmt.pos.Z)
	if !mask {
		err = fmt.Errorf("pos[%s] invalid", mmt.pos.String())
		err = template.NewTemplateFieldError("pos", err)
		return
	}
	y := mapTemplate.GetMap().GetHeight(mmt.pos.X, mmt.pos.Z)
	mmt.pos.Y = y

	return nil
}
func (mmt *MarryMoveTemplate) PatchAfterCheck() {

}
func (mmt *MarryMoveTemplate) FileName() string {
	return "tb_marry_move.json"
}

func init() {
	template.Register((*MarryMoveTemplate)(nil))
}
