package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/common/common"
	"fmt"
)

type PortalTemplate struct {
	*PortalTemplateVO
	pos         coretypes.Position
	mapTemplate *MapTemplate
}

func (pt *PortalTemplate) GetPosition() coretypes.Position {
	return pt.pos
}

func (pt *PortalTemplate) TemplateId() int {
	return pt.Id
}

func (pt *PortalTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(pt.FileName(), pt.TemplateId(), err)
			return
		}
	}()
	tempMapTemplate := template.GetTemplateService().Get(int(pt.MapId), (*MapTemplate)(nil))
	if tempMapTemplate == nil {
		err = fmt.Errorf("%d invalid", pt.MapId)
		err = template.NewTemplateFieldError("mapId", err)
		return
	}
	pt.mapTemplate = tempMapTemplate.(*MapTemplate)
	pt.pos = coretypes.Position{
		X: float64(pt.PosX) / common.MILL_METER,
		Y: float64(pt.PosY) / common.MILL_METER,
		Z: float64(pt.PosZ) / common.MILL_METER,
	}
	return nil
}

func (pt *PortalTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(pt.FileName(), pt.TemplateId(), err)
			return
		}
	}()
	mask := pt.mapTemplate.GetMap().IsMask(pt.pos.X, pt.pos.Z)
	if !mask {
		err = fmt.Errorf("pos[%s] invalid", pt.pos.String())
		err = template.NewTemplateFieldError("pos", err)
	}
	y := pt.mapTemplate.GetMap().GetHeight(pt.pos.X, pt.pos.Z)
	pt.pos.Y = y

	return nil
}
func (pt *PortalTemplate) PatchAfterCheck() {

}

func (pt *PortalTemplate) FileName() string {
	return "tb_portal.json"
}

func init() {
	template.Register((*PortalTemplate)(nil))
}
