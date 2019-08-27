package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//玩具套装配置
type BabyToySuitGroupTemplate struct {
	*BabyToySuitGroupTemplateVO
	suitEffectTemplateList []*BabyToySuitTemplate //套装效果
}

func (t *BabyToySuitGroupTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyToySuitGroupTemplate) GetSuitEffectSkillId(equipNum int32) (skillList []int32) {
	for _, temp := range t.suitEffectTemplateList {
		suitSkillEffectLimitNum := temp.Num
		if equipNum >= suitSkillEffectLimitNum {
			skillList = append(skillList, temp.Value1)
		}
	}

	return skillList
}

func (t *BabyToySuitGroupTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//套装技能1
	tempItemTemplate1 := template.GetTemplateService().Get(int(t.SuitId1), (*BabyToySuitTemplate)(nil))
	if tempItemTemplate1 == nil {
		err = fmt.Errorf("[%d] invalid", t.SuitId1)
		err = template.NewTemplateFieldError("suitId1", err)
		return
	}
	effectTemp1 := tempItemTemplate1.(*BabyToySuitTemplate)
	if effectTemp1 == nil {
		err = fmt.Errorf("[%d] invalid", t.SuitId1)
		err = template.NewTemplateFieldError("suitId1", err)
		return
	}

	t.suitEffectTemplateList = append(t.suitEffectTemplateList, effectTemp1)

	//套装技能2
	tempItemTemplate2 := template.GetTemplateService().Get(int(t.SuitId2), (*BabyToySuitTemplate)(nil))
	if tempItemTemplate2 == nil {
		err = fmt.Errorf("[%d] invalid", t.SuitId2)
		err = template.NewTemplateFieldError("suitId2", err)
		return
	}
	effectTemp2 := tempItemTemplate2.(*BabyToySuitTemplate)
	if effectTemp2 == nil {
		err = fmt.Errorf("[%d] invalid", t.SuitId2)
		err = template.NewTemplateFieldError("suitId2", err)
		return
	}

	t.suitEffectTemplateList = append(t.suitEffectTemplateList, effectTemp2)

	//套装技能3
	tempItemTemplate3 := template.GetTemplateService().Get(int(t.SuitId3), (*BabyToySuitTemplate)(nil))
	if tempItemTemplate3 == nil {
		err = fmt.Errorf("[%d] invalid", t.SuitId3)
		err = template.NewTemplateFieldError("suitId3", err)
		return
	}
	effectTemp3 := tempItemTemplate3.(*BabyToySuitTemplate)
	if effectTemp3 == nil {
		err = fmt.Errorf("[%d] invalid", t.SuitId3)
		err = template.NewTemplateFieldError("suitId3", err)
		return
	}

	t.suitEffectTemplateList = append(t.suitEffectTemplateList, effectTemp3)

	// //套装技能4
	// tempItemTemplate4 := template.GetTemplateService().Get(int(t.SuitId4), (*BabyToySuitTemplate)(nil))
	// if tempItemTemplate3 == nil {
	// 	err = fmt.Errorf("[%d] invalid", t.SuitId4)
	// 	err = template.NewTemplateFieldError("SuitId4", err)
	// 	return
	// }
	// effectTemp4 := tempItemTemplate4.(*BabyToySuitTemplate)
	// if effectTemp3 == nil {
	// 	err = fmt.Errorf("[%d] invalid", t.SuitId4)
	// 	err = template.NewTemplateFieldError("SuitId4", err)
	// 	return
	// }

	// t.suitEffectTemplateList = append(t.suitEffectTemplateList, effectTemp4)

	return nil
}

func (t *BabyToySuitGroupTemplate) PatchAfterCheck() {

}

func (t *BabyToySuitGroupTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//套装数量
	if err = validator.MinValidate(float64(t.MaxNum), float64(0), false); err != nil {
		err = template.NewTemplateFieldError("number", err)
		return
	}

	return nil
}

func (t *BabyToySuitGroupTemplate) FileName() string {
	return "tb_baobao_wanju_suit_group.json"
}

func init() {
	template.Register((*BabyToySuitGroupTemplate)(nil))
}
