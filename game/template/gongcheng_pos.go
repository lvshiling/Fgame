package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/godsiege/types"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//神兽攻城出生配置
type GongChengPosTemplate struct {
	*GongChengPosTemplateVO
	bornType types.GodSiegeBornType
	posType  godsiegetypes.GodSiegePosType
	pos      coretypes.Position //位置
}

func (t *GongChengPosTemplate) TemplateId() int {
	return t.Id
}

func (t *GongChengPosTemplate) GetPos() coretypes.Position {
	return t.pos
}

func (t *GongChengPosTemplate) GetBornType() types.GodSiegeBornType {
	return t.bornType
}

func (t *GongChengPosTemplate) GetPosType() godsiegetypes.GodSiegePosType {
	return t.posType
}

func (t *GongChengPosTemplate) GetMapId() int32 {
	return t.MapId
}

func (t *GongChengPosTemplate) Patch() (err error) {
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

func (t *GongChengPosTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.bornType = types.GodSiegeBornType(t.Type)
	if !t.bornType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	t.posType = godsiegetypes.GodSiegePosType(t.Pos)
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

	if mapTemplate.GetMapType() != scenetypes.SceneTypeCrossGodSiege &&
		mapTemplate.GetMapType() != scenetypes.SceneTypeCrossDenseWat {
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
func (t *GongChengPosTemplate) PatchAfterCheck() {

}
func (t *GongChengPosTemplate) FileName() string {
	return "tb_gongcheng_pos.json"
}

func init() {
	template.Register((*GongChengPosTemplate)(nil))
}
