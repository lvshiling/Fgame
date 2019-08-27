package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/lianyu/types"
	lianyutypes "fgame/fgame/game/lianyu/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//无间炼狱出生配置
type LianYuPosTemplate struct {
	*LianYuPosTemplateVO
	bornType types.LianYuBornType
	posType  lianyutypes.LianYuPosType
	pos      coretypes.Position //位置
}

func (t *LianYuPosTemplate) TemplateId() int {
	return t.Id
}

func (t *LianYuPosTemplate) GetPos() coretypes.Position {
	return t.pos
}

func (t *LianYuPosTemplate) GetBornType() types.LianYuBornType {
	return t.bornType
}

func (t *LianYuPosTemplate) GetPosType() lianyutypes.LianYuPosType {
	return t.posType
}

func (t *LianYuPosTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.pos = coretypes.Position{
		X: t.PosX,
		Y: t.PosY,
		Z: t.PosZ,
	}

	return nil
}

func (t *LianYuPosTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.bornType = types.LianYuBornType(t.Type)
	if !t.bornType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	t.posType = lianyutypes.LianYuPosType(t.Pos)
	if !t.posType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Pos)
		return template.NewTemplateFieldError("Pos", err)
	}

	to := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	mapTemplate := to.(*MapTemplate)

	if mapTemplate.GetMapType() != scenetypes.SceneTypeCrossLianYu && mapTemplate.GetMapType() != scenetypes.SceneTypeLocalLianYu {
		err = fmt.Errorf("[%d] invalid", t.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

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
func (t *LianYuPosTemplate) PatchAfterCheck() {

}
func (t *LianYuPosTemplate) FileName() string {
	return "tb_lianyu_pos.json"
}

func init() {
	template.Register((*LianYuPosTemplate)(nil))
}
