package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	playertypes "fgame/fgame/game/player/types"
	synthesistypes "fgame/fgame/game/synthesis/types"
	"fmt"
)

//合成等级配置
type SynthesisTemplate struct {
	*SynthesisTemplateVO
	needRole      playertypes.RoleType
	needGender    playertypes.SexType
	synthesisType synthesistypes.SynthesisType
	synthesisMap  map[int32]int32
}

func (t *SynthesisTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.synthesisMap = make(map[int32]int32)
	if t.NeedItemId1 != 0 {
		t.synthesisMap[t.NeedItemId1] = t.NeedItemCount1
	}
	if t.NeedItemId2 != 0 {
		t.synthesisMap[t.NeedItemId2] = t.NeedItemCount2
	}
	if t.NeedItemId3 != 0 {
		t.synthesisMap[t.NeedItemId3] = t.NeedItemCount3
	}
	if t.NeedItemId4 != 0 {
		t.synthesisMap[t.NeedItemId4] = t.NeedItemCount4
	}
	if t.NeedItemId5 != 0 {
		t.synthesisMap[t.NeedItemId5] = t.NeedItemCount5
	}
	return nil
}

//Check 校验数据
func (t *SynthesisTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//子标签id
	t.synthesisType = synthesistypes.SynthesisType(t.FirstTabId)
	if !t.synthesisType.Valid() {
		return fmt.Errorf("FirstTabId [%d] invalid", t.FirstTabId)
	}

	//角色
	if t.NeedProfession > 0 {
		role := playertypes.RoleType(t.NeedProfession)
		if !role.Valid() {
			err = template.NewTemplateFieldError("NeedProfession", err)
			return
		}
		t.needRole = role
	}

	//性别
	if t.NeedGender > 0 {
		sex := playertypes.SexType(t.NeedGender)
		if !sex.Valid() {
			err = template.NewTemplateFieldError("NeedGender", err)
			return
		}
		t.needGender = sex
	}

	//验证合成后的物品ID
	sysnthesisTmp := template.GetTemplateService().Get(int(t.ItemId), (*ItemTemplate)(nil))
	if sysnthesisTmp == nil {
		err = fmt.Errorf("[%d] invalid", t.ItemId)
		return template.NewTemplateFieldError("ItemId", err)
	}
	//验证合成后的数量
	if err = validator.MinValidate(float64(t.ItemCount), float64(1), true); err != nil {
		err = template.NewTemplateFieldError("ItemCount", err)
		return
	}

	// 验证合成所需物品 和 合成所需物品的数量
	if len(t.synthesisMap) == 0 {
		err = fmt.Errorf("need item should be more than 0")
		err = template.NewTemplateFieldError("needItem", err)
		return
	}
	for needItemID, needItemCount := range t.synthesisMap {
		itemTmp := template.GetTemplateService().Get(int(needItemID), (*ItemTemplate)(nil))
		if itemTmp == nil {
			err = fmt.Errorf("[%d] invalid", needItemID)
			return template.NewTemplateFieldError("NeedItemId", err)
		}

		err = validator.MinValidate(float64(needItemCount), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
	}

	//验证合成所需银两
	if err = validator.MinValidate(float64(t.NeedSilver), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("NeedSilver", err)
		return
	}
	//验证合成所需元宝
	if err = validator.MinValidate(float64(t.NeedGold), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("NeedGold", err)
		return
	}
	//验证合成所需绑定元宝
	if err = validator.MinValidate(float64(t.NeedBindGold), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("NeedBindGold", err)
		return
	}
	//验证合成成功率
	if err = validator.MinValidate(float64(t.SuccessRate), float64(1), true); err != nil {
		err = template.NewTemplateFieldError("SuccessRate", err)
		return
	}
	//验证合成最大数量限制
	if err = validator.MinValidate(float64(t.MaxCount), float64(1), true); err != nil {
		err = template.NewTemplateFieldError("MaxCount", err)
		return
	}

	return nil
}

func (t *SynthesisTemplate) PatchAfterCheck() {

}

func (t *SynthesisTemplate) TemplateId() int {
	return t.Id
}

func (t *SynthesisTemplate) FileName() string {
	return "tb_item_synthesis.json"
}

func (t *SynthesisTemplate) GetSynthesisMap() map[int32]int32 {
	return t.synthesisMap
}

func (t *SynthesisTemplate) GetSynthesisType() synthesistypes.SynthesisType {
	return t.synthesisType
}

func (t *SynthesisTemplate) GetNeedRole() playertypes.RoleType {
	return t.needRole
}

func (t *SynthesisTemplate) GetNeedGender() playertypes.SexType {
	return t.needGender
}

func init() {
	template.Register((*SynthesisTemplate)(nil))
}
