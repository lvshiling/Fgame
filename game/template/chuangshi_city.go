package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coretypes "fgame/fgame/core/types"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fmt"
)

type ChuangShiCityTemplate struct {
	*ChuangShiCityTemplateVO
	campType     chuangshitypes.ChuangShiCampType
	cityType     chuangshitypes.ChuangShiCityType
	constantTemp *ChuangShiCityConstantTemplate
}

func (t *ChuangShiCityTemplate) GetBornPos(campType chuangshitypes.ChuangShiCampType) coretypes.Position {
	pos := t.constantTemp.GetPos1()
	if t.campType != campType {
		pos = t.constantTemp.GetPos2()
	}
	return pos
}

func (t *ChuangShiCityTemplate) IfCanEnter(now int64) bool {
	if t.constantTemp.IsJinRu == 0 {
		return false
	}

	begin := t.constantTemp.GetBeginTime(now)
	end := t.constantTemp.GetEndTime(now)

	if now < begin || now > end {
		return false
	}

	return true
}

func (t *ChuangShiCityTemplate) GetMapId() int32 {
	return t.constantTemp.MapId
}

func (t *ChuangShiCityTemplate) GetCamp() chuangshitypes.ChuangShiCampType {
	return t.campType
}

func (t *ChuangShiCityTemplate) GetCityType() chuangshitypes.ChuangShiCityType {
	return t.cityType
}

func (t *ChuangShiCityTemplate) TemplateId() int {
	return t.Id
}

func (t *ChuangShiCityTemplate) FileName() string {
	return "tb_chuangshi_city.json"
}

func (t *ChuangShiCityTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	to := template.GetTemplateService().Get(int(t.CityConstantId), (*ChuangShiCityConstantTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.CityConstantId)
		return template.NewTemplateFieldError("CityConstantId", err)
	}
	t.constantTemp = to.(*ChuangShiCityConstantTemplate)

	//
	t.campType = chuangshitypes.ChuangShiCampType(t.Camp)
	if !t.campType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Camp)
		return template.NewTemplateFieldError("Camp", err)
	}
	//
	t.cityType = chuangshitypes.ChuangShiCityType(t.Type)
	if !t.cityType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}
	return
}

func (t *ChuangShiCityTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	err = validator.MinValidate(float64(t.PlayerRewJifen), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.PlayerRewJifen)
		return template.NewTemplateFieldError("PlayerRewJifen", err)
	}

	//
	err = validator.MinValidate(float64(t.PlayerRewZuanshi), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.PlayerRewZuanshi)
		return template.NewTemplateFieldError("PlayerRewZuanshi", err)
	}
	//
	err = validator.MinValidate(float64(t.ZhenyingRewJifen), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhenyingRewJifen)
		return template.NewTemplateFieldError("ZhenyingRewJifen", err)
	}

	//
	err = validator.MinValidate(float64(t.ZhenyingRewZuanshi), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhenyingRewZuanshi)
		return template.NewTemplateFieldError("ZhenyingRewZuanshi", err)
	}

	return
}

func (t *ChuangShiCityTemplate) PatchAfterCheck() {

}

func init() {
	template.Register((*ChuangShiCityTemplate)(nil))
}
