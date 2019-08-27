package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	fourgodtypes "fgame/fgame/game/fourgod/types"
	scenetypes "fgame/fgame/game/scene/types"

	"fmt"
)

//四神遗迹特殊怪配置
type FourGodSpecialTemplate struct {
	*FourGodSpecialTemplateVO
	typ fourgodtypes.FourGodSpecialPosType
	//地图位置
	pos coretypes.Position
	//四神遗迹地图
	fourGodWarMap *MapTemplate
}

func (fgst *FourGodSpecialTemplate) TemplateId() int {
	return fgst.Id
}

func (fgst *FourGodSpecialTemplate) GetTyp() fourgodtypes.FourGodSpecialPosType {
	return fgst.typ
}

func (fgst *FourGodSpecialTemplate) GetPos() coretypes.Position {
	return fgst.pos
}

func (fgst *FourGodSpecialTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(fgst.FileName(), fgst.TemplateId(), err)
			return
		}
	}()

	//type
	fgst.typ = fourgodtypes.FourGodSpecialPosType(fgst.Type)
	if !fgst.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", fgst.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//map_id
	to := template.GetTemplateService().Get(int(fgst.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", fgst.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	fgst.fourGodWarMap = to.(*MapTemplate)

	fgst.pos = coretypes.Position{
		X: fgst.PosX,
		Y: fgst.PosY,
		Z: fgst.PosZ,
	}

	return nil
}

func (fgst *FourGodSpecialTemplate) PatchAfterCheck() {

}

func (fgst *FourGodSpecialTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(fgst.FileName(), fgst.TemplateId(), err)
			return
		}
	}()

	mask := fgst.fourGodWarMap.GetMap().IsMask(fgst.pos.X, fgst.pos.Z)
	if !mask {
		err = fmt.Errorf("pos[%s] invalid", fgst.pos.String())
		err = template.NewTemplateFieldError("pos", err)
		return
	}
	y := fgst.fourGodWarMap.GetMap().GetHeight(fgst.pos.X, fgst.pos.Z)
	fgst.pos.Y = y

	//验证类型
	if fgst.fourGodWarMap.GetMapType() != scenetypes.SceneTypeFourGodWar {
		err = fmt.Errorf("[%d] invalid", fgst.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	return nil
}

func (fgst *FourGodSpecialTemplate) FileName() string {
	return "tb_four_special.json"
}

func init() {
	template.Register((*FourGodSpecialTemplate)(nil))
}
