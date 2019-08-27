package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	playertypes "fgame/fgame/game/player/types"
)

//元神金装重铸配置
type GoldEquipChongzhuTemplate struct {
	*GoldEquipChongzhuTemplateVO
}

func (t *GoldEquipChongzhuTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldEquipChongzhuTemplate) Patch() (err error) {

	return nil
}

func (t *GoldEquipChongzhuTemplate) PatchAfterCheck() {

}

func (t *GoldEquipChongzhuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//角色
	role := playertypes.RoleType(t.Profession)
	if !role.Valid() {
		err = template.NewTemplateFieldError("Profession", err)
		return
	}

	//性别
	sex := playertypes.SexType(t.Gender)
	if !sex.Valid() {
		err = template.NewTemplateFieldError("Gender", err)
		return
	}
	//质量
	if err = validator.MinValidate(float64(t.Quality), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("Quality", err)
		return
	}

	//转生数
	if err = validator.MinValidate(float64(t.Level), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	return nil
}

func (t *GoldEquipChongzhuTemplate) FileName() string {
	return "tb_gold_equip_chongzhu.json"
}

func init() {
	template.Register((*GoldEquipChongzhuTemplate)(nil))
}
