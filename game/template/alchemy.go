package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//炼丹配置
type AlchemyTemplate struct {
	*AlchemyTemplateVO
	alchemyMap map[int]int32
}

func (at *AlchemyTemplate) TemplateId() int {
	return at.Id
}

func (at *AlchemyTemplate) GetAllAlchemy() map[int]int32 {
	return at.alchemyMap
}

func (at *AlchemyTemplate) GetAchemyTime() int32 {
	return at.Time
}

func (at *AlchemyTemplate) GetAchemyaMoney() int32 {
	return at.AccelerateMoney
}

func (at *AlchemyTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(at.FileName(), at.TemplateId(), err)
			return
		}
	}()

	at.alchemyMap = make(map[int]int32)

	//验证NeedItemId1
	needItem1Template := template.GetTemplateService().Get(int(at.NeedItemId1), (*ItemTemplate)(nil))
	if needItem1Template == nil {
		err = fmt.Errorf("[%d] invalid", at.NeedItemId1)
		return template.NewTemplateFieldError("NeedItemId1", err)
	}

	//验证NeedItemNum1
	err = validator.MinValidate(float64(at.NeedItemNum1), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("NeedItemNum1", err)
		return
	}

	//验证NeedItemId2

	needItem2Template := template.GetTemplateService().Get(int(at.NeedItemId2), (*ItemTemplate)(nil))
	if needItem2Template == nil {
		err = fmt.Errorf("[%d] invalid", at.NeedItemId2)
		return template.NewTemplateFieldError("NeedItemId2", err)
	}

	//验证NeedItemNum2
	err = validator.MinValidate(float64(at.NeedItemNum2), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("NeedItemNum2", err)
		return
	}

	//验证NeedItemId3
	needItem3Template := template.GetTemplateService().Get(int(at.NeedItemId3), (*ItemTemplate)(nil))
	if needItem3Template == nil {
		err = fmt.Errorf("[%d] invalid", at.NeedItemId3)
		return template.NewTemplateFieldError("NeedItemId3", err)
	}

	//验证NeedItemNum3
	err = validator.MinValidate(float64(at.NeedItemNum3), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("NeedItemNum3", err)
		return
	}
	at.alchemyMap[int(at.NeedItemId1)] = int32(at.NeedItemNum1)
	at.alchemyMap[int(at.NeedItemId2)] = int32(at.NeedItemNum2)
	at.alchemyMap[int(at.NeedItemId3)] = int32(at.NeedItemNum3)

	return nil
}

func (at *AlchemyTemplate) PatchAfterCheck() {

}

func (at *AlchemyTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(at.FileName(), at.TemplateId(), err)
			return
		}
	}()

	//验证SynthetiseId
	synthetiseItemTemplate := template.GetTemplateService().Get(int(at.SynthetiseId), (*ItemTemplate)(nil))
	if synthetiseItemTemplate == nil {
		err = fmt.Errorf("[%d] invalid", at.SynthetiseId)
		return template.NewTemplateFieldError("SynthetiseId", err)
	}

	//验证SynthetiseNum
	err = validator.MinValidate(float64(at.SynthetiseNum), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("SynthetiseNum", err)
		return
	}

	//验证Time
	err = validator.MinValidate(float64(at.Time), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("Time", err)
		return
	}

	//验证Time
	err = validator.MinValidate(float64(at.AccelerateMoney), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("AccelerateMoney", err)
		return
	}

	return nil
}

func (at *AlchemyTemplate) FileName() string {
	return "tb_alchemy.json"
}

func init() {
	template.Register((*AlchemyTemplate)(nil))
}
