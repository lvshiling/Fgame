package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fmt"
)

/*场景部怪配置*/
type SceneTemplate struct {
	*SceneTemplateVO
	pos     coretypes.Position
	biology *BiologyTemplate
}

func (st *SceneTemplate) GetPos() coretypes.Position {
	return st.pos
}

func (st *SceneTemplate) GetBiology() *BiologyTemplate {
	return st.biology
}

//特殊处理
func (st *SceneTemplate) TemplateId() int {
	return st.Idx
}

func (st *SceneTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()
	st.pos = coretypes.Position{
		X: st.PosX,
		Y: st.PosY,
		Z: st.PosZ,
	}

	to := template.GetTemplateService().Get(st.TempId, (*BiologyTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("%d invalid", st.TempId)
		err = template.NewTemplateFieldError("tempId", err)
		return
	}
	biology, _ := to.(*BiologyTemplate)
	st.biology = biology

	return nil
}
func (st *SceneTemplate) PatchAfterCheck() {

}
func (st *SceneTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()
	// if st.biology.ScriptType != 0 {
	//检查位置
	tempMapTemplate := template.GetTemplateService().Get(int(st.SceneID), (*MapTemplate)(nil))
	if tempMapTemplate == nil {
		err = fmt.Errorf("%d invalid", st.SceneID)
		err = template.NewTemplateFieldError("SceneID", err)
		return
	}

	mapTemplate := tempMapTemplate.(*MapTemplate)

	if !mapTemplate.GetMap().IsMask(st.pos.X, st.pos.Z) {
		err = template.NewTemplateFieldError("Pos", err)
		return
	}

	return nil
}

func (st *SceneTemplate) FileName() string {
	return "tb_scene.json"
}

func init() {
	template.Register((*SceneTemplate)(nil))
}
