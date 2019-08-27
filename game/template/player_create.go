package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/core/utils"
	playertypes "fgame/fgame/game/player/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

type PlayerCreateTemplate struct {
	*PlayerCreateTemplateVO
	sex            playertypes.SexType
	role           playertypes.RoleType
	skillNormalMap map[int32]int32 //普通技能
	jumpSkill      *SkillTemplate  //跳跃技能
	bornPos        coretypes.Position
}

func (t *PlayerCreateTemplate) GetBornPos() coretypes.Position {
	return t.bornPos
}

func (t *PlayerCreateTemplate) GetSkillNormalMap() map[int32]int32 {
	return t.skillNormalMap
}

func (t *PlayerCreateTemplate) GetJumpSkill() *SkillTemplate {
	return t.jumpSkill
}

func (t *PlayerCreateTemplate) GetSex() playertypes.SexType {
	return t.sex
}
func (t *PlayerCreateTemplate) GetRole() playertypes.RoleType {
	return t.role
}

func (t *PlayerCreateTemplate) TemplateId() int {
	return t.Id
}

func (t *PlayerCreateTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.sex = playertypes.SexType(t.Gender)
	t.role = playertypes.RoleType(t.Zhiye)

	//检验技能是否存在
	t.skillNormalMap = make(map[int32]int32)
	if t.Spell != "" {
		spellArr, errInfo := utils.SplitAsIntArray(t.Spell)
		if errInfo != nil {
			return errInfo
		}

		for _, spell := range spellArr {
			tempSkillTemplate := template.GetTemplateService().Get(int(spell), (*SkillTemplate)(nil))
			if tempSkillTemplate == nil {
				err = fmt.Errorf("spell [%s] 无效", t.Spell)
				return
			}
			skillTemplate, _ := tempSkillTemplate.(*SkillTemplate)

			firstType := skillTemplate.GetSkillFirstType()
			if firstType != skilltypes.SkillFirstTypeNormal {
				err = fmt.Errorf("spell [%s] 应该是普通技能", t.Spell)
				return
			}
			t.skillNormalMap[spell] = skillTemplate.Lev

		}
	}

	//检验初始地图
	mto := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if mto == nil {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return err
	}
	_, ok := mto.(*MapTemplate)
	if !ok {
		err = fmt.Errorf("mapId [%d] no exist", t.MapId)
		return
	}
	//检验时装
	fashionTemplate := template.GetTemplateService().Get(int(t.FashionId), (*FashionTemplate)(nil))
	if fashionTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.FashionId)
		err = template.NewTemplateFieldError("fashionId", err)
		return
	}

	//检验时装
	weaponTemplate := template.GetTemplateService().Get(int(t.WeaponId), (*WeaponTemplate)(nil))
	if weaponTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.WeaponId)
		err = template.NewTemplateFieldError("WeaponId", err)
		return
	}

	tempJumpSkillTemplate := template.GetTemplateService().Get(int(t.JumpId), (*SkillTemplate)(nil))
	if tempJumpSkillTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.JumpId)
		err = template.NewTemplateFieldError("JumpId", err)
		return
	}

	t.jumpSkill = tempJumpSkillTemplate.(*SkillTemplate)

	return nil
}

func (t *PlayerCreateTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	if !t.sex.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Gender)
		err = template.NewTemplateFieldError("gender", err)
		return
	}
	if !t.role.Valid() {
		err = fmt.Errorf("zhiye [%d] invalid", t.Zhiye)
		err = template.NewTemplateFieldError("zhiye", err)
		return
	}
	firstType := t.jumpSkill.GetSkillFirstType()
	if firstType != skilltypes.SkillFirstTypeJump {
		err = fmt.Errorf("spell [%d] [%d] 应该是跳跃技能", firstType, t.JumpId)
		err = template.NewTemplateFieldError("JumpId", err)
		return
	}

	t.bornPos = coretypes.Position{
		X: t.PositionX,
		Z: t.PositionY,
	}
	tempMapTemplate := template.GetTemplateService().Get(int(t.MapId), (*MapTemplate)(nil))
	if tempMapTemplate == nil {
		err = fmt.Errorf("spell [%d] [%d] 应该是跳跃技能", t.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}

	mapTemplate := tempMapTemplate.(*MapTemplate)
	t.bornPos.Y = mapTemplate.GetMap().GetHeight(t.bornPos.X, t.bornPos.Z)
	if !mapTemplate.GetMap().IsMask(t.bornPos.X, t.bornPos.Z) {
		err = fmt.Errorf("[%.2f] [%.2f] 位置无效", t.bornPos.X, t.bornPos.Z)
		err = template.NewTemplateFieldError("Position", err)
		return
	}

	return nil
}
func (pct *PlayerCreateTemplate) PatchAfterCheck() {

}
func (pct *PlayerCreateTemplate) FileName() string {
	return "tb_playcreat_info.json"
}

func init() {
	template.Register((*PlayerCreateTemplate)(nil))
}
