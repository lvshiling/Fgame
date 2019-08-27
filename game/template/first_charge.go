package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	playertypes "fgame/fgame/game/player/types"

	"fmt"
)

//首冲配置
type FirstChargeTemplate struct {
	*FirstChargeTemplateVO
	role       playertypes.RoleType
	sex        playertypes.SexType
	rewItemMap map[int32]int32
}

func (t *FirstChargeTemplate) TemplateId() int {
	return t.Id
}

func (t *FirstChargeTemplate) GetSex() playertypes.SexType {
	return t.sex
}

func (t *FirstChargeTemplate) GetRole() playertypes.RoleType {
	return t.role
}

func (t *FirstChargeTemplate) GetRewItemMap() map[int32]int32 {
	return t.rewItemMap
}

func (t *FirstChargeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.rewItemMap = make(map[int32]int32)
	//验证 rew_item_id
	rewItemIdList, err := utils.SplitAsIntArray(t.RewItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RewItemId)
		return template.NewTemplateFieldError("RewItemId", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(t.RewItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.RewItemCount)
		return template.NewTemplateFieldError("RewItemCount", err)
	}
	if len(rewItemCountList) != len(rewItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.RewItemId, t.RewItemCount)
		err = template.NewTemplateFieldError("RewItemId or RewItemCount", err)
		return err
	}
	if len(rewItemIdList) > 0 {
		//组合数据
		for index, itemId := range rewItemIdList {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.RewItemId)
				return template.NewTemplateFieldError("RewItemId", err)
			}

			err = validator.MinValidate(float64(rewItemCountList[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("RewItemCount", err)
			}

			t.rewItemMap[itemId] = rewItemCountList[index]
		}
	}

	return nil
}

func (t *FirstChargeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//角色
	role := playertypes.RoleType(t.Profession)
	if !role.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Profession)
		return template.NewTemplateFieldError("Profession", err)
	}
	t.role = role

	//性别
	sex := playertypes.SexType(t.Gender)
	if !sex.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Gender)
		return template.NewTemplateFieldError("Gender", err)
	}
	t.sex = sex

	//验证 rew_silver
	err = validator.MinValidate(float64(t.RewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewSilver)
		err = template.NewTemplateFieldError("RewSilver", err)
		return
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(t.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewGold)
		err = template.NewTemplateFieldError("RewGold", err)
		return
	}

	//验证 rew_bind_gold
	err = validator.MinValidate(float64(t.RewGoldBind), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.RewGoldBind)
		err = template.NewTemplateFieldError("RewGoldBind", err)
		return
	}

	return nil
}
func (t *FirstChargeTemplate) PatchAfterCheck() {

}
func (t *FirstChargeTemplate) FileName() string {
	return "tb_firstcharge.json"
}

func init() {
	template.Register((*FirstChargeTemplate)(nil))
}
