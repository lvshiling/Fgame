package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/core/utils"
	onearenatypes "fgame/fgame/game/onearena/types"
	scenetypes "fgame/fgame/game/scene/types"

	"fmt"
)

//灵池争夺配置
type OneArenaTemplate struct {
	*OneArenaTemplateVO
	arenaType       onearenatypes.OneArenaLevelType
	mapTemplate     *MapTemplate
	biologyTemplate *BiologyTemplate
	dropList        []int32
	pos             coretypes.Position //位置
}

func (oat *OneArenaTemplate) TemplateId() int {
	return oat.Id
}

func (oat *OneArenaTemplate) GetArenaType() onearenatypes.OneArenaLevelType {
	return oat.arenaType
}

func (oat *OneArenaTemplate) GetMapTemplate() *MapTemplate {
	return oat.mapTemplate
}

func (oat *OneArenaTemplate) GetBiologyTemplate() *BiologyTemplate {
	return oat.biologyTemplate
}

func (oat *OneArenaTemplate) GetDropList() []int32 {
	return oat.dropList
}

func (oat *OneArenaTemplate) GetPos() coretypes.Position {
	return oat.pos
}

func (oat *OneArenaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(oat.FileName(), oat.TemplateId(), err)
			return
		}
	}()

	oat.arenaType = onearenatypes.OneArenaLevelType(oat.Level)
	if !oat.arenaType.Vaild() {
		err = fmt.Errorf("[%d] invalid", oat.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	oat.pos = coretypes.Position{
		X: oat.PosX,
		Y: oat.PosY,
		Z: oat.PosZ,
	}

	to := template.GetTemplateService().Get(int(oat.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", oat.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}
	oat.mapTemplate = to.(*MapTemplate)

	to = template.GetTemplateService().Get(int(oat.BiologyId), (*BiologyTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", oat.BiologyId)
		err = template.NewTemplateFieldError("BiologyId", err)
		return
	}
	oat.biologyTemplate = to.(*BiologyTemplate)

	oat.dropList = make([]int32, 0, 8)
	if oat.DropId != "" {
		dropIdArr, err := utils.SplitAsIntArray(oat.DropId)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", oat.DropId)
			err = template.NewTemplateFieldError("DropId", err)
			return err
		}

		for _, dropId := range dropIdArr {
			oat.dropList = append(oat.dropList, dropId)
		}
	}

	return nil
}

func (oat *OneArenaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(oat.FileName(), oat.TemplateId(), err)
			return
		}
	}()

	if oat.NextId != 0 {
		to := template.GetTemplateService().Get(int(oat.NextId), (*OneArenaTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", oat.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		oneArenaTemplate := to.(*OneArenaTemplate)
		if oneArenaTemplate.Level != oat.Level {
			err = fmt.Errorf("[%d] invalid", oat.Level)
			err = template.NewTemplateFieldError("Level", err)
			return
		}
		diffPos := oneArenaTemplate.PosId - oat.PosId
		if diffPos != 1 {
			err = fmt.Errorf("[%d] invalid", oat.PosId)
			err = template.NewTemplateFieldError("PosId", err)
			return
		}
	}

	err = validator.MinValidate(float64(oat.RefreshTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", oat.RefreshTime)
		err = template.NewTemplateFieldError("RefreshTime", err)
		return
	}

	err = validator.MinValidate(float64(oat.CoolTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", oat.CoolTime)
		err = template.NewTemplateFieldError("CoolTime", err)
		return
	}

	if oat.mapTemplate.GetMapType() != scenetypes.SceneTypeLingChiFighting {
		err = fmt.Errorf("[%d] invalid", oat.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	mask := oat.mapTemplate.GetMap().IsMask(oat.pos.X, oat.pos.Z)
	if !mask {
		err = fmt.Errorf("pos[%s] invalid", oat.pos.String())
		err = template.NewTemplateFieldError("pos", err)
		return
	}
	y := oat.mapTemplate.GetMap().GetHeight(oat.pos.X, oat.pos.Z)
	oat.pos.Y = y

	if oat.biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeOneArenaGuardian {
		err = fmt.Errorf("[%d] invalid", oat.BiologyId)
		return template.NewTemplateFieldError("BiologyId", err)
	}

	return nil
}
func (oat *OneArenaTemplate) PatchAfterCheck() {

}
func (oat *OneArenaTemplate) FileName() string {
	return "tb_onearena.json"
}

func init() {
	template.Register((*OneArenaTemplate)(nil))
}
