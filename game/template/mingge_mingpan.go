package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/mingge/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//命格合成配置
type MingGeMingPanTemplate struct {
	*MingGeMingPanTemplateVO
	mingPanType         types.MingGeAllSubType
	needItemMap         map[int32]int32
	battlePropertyMap   map[propertytypes.BattlePropertyType]int64
	nextMingPanTemplate *MingGeMingPanTemplate
}

func (mt *MingGeMingPanTemplate) TemplateId() int {
	return mt.Id
}

func (mt *MingGeMingPanTemplate) GetMingPanType() types.MingGeAllSubType {
	return mt.mingPanType
}

func (mt *MingGeMingPanTemplate) GetNextMingPanTemplate() *MingGeMingPanTemplate {
	return mt.nextMingPanTemplate
}

func (mt *MingGeMingPanTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return mt.battlePropertyMap
}

func (mt *MingGeMingPanTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	mt.mingPanType = types.MingGeAllSubType(mt.SubType)
	if !mt.mingPanType.Valid() {
		err = fmt.Errorf("[%d] invalid", mt.SubType)
		err = template.NewTemplateFieldError("SubType", err)
		return
	}

	if mt.UseItemId != 0 {
		mt.needItemMap = make(map[int32]int32)
		to := template.GetTemplateService().Get(int(mt.UseItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.UseItemId)
			err = template.NewTemplateFieldError("UseItemId", err)
			return
		}

		err = validator.MinValidate(float64(mt.UseItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mt.UseItemCount)
			err = template.NewTemplateFieldError("UseItemCount", err)
			return
		}
		mt.needItemMap[mt.UseItemId] = mt.UseItemCount
	}

	//属性
	mt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	mt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(mt.Hp)
	mt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(mt.Attack)
	mt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(mt.Defence)

	return nil
}

func (mt *MingGeMingPanTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证Number
	err = validator.MinValidate(float64(mt.Number), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Number)
		err = template.NewTemplateFieldError("Number", err)
		return
	}

	//验证 star
	err = validator.MinValidate(float64(mt.Star), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Star)
		err = template.NewTemplateFieldError("Star", err)
		return
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(mt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(mt.TimesMin), float64(0), true, float64(mt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mt.TimesMax), float64(mt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(mt.AddMin), float64(0), true, float64(mt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(mt.AddMax), float64(mt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 ZhufuMax
	err = validator.MinValidate(float64(mt.ZhuFuMax), float64(mt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.ZhuFuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 rate
	err = validator.RangeValidate(float64(mt.Rate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Rate)
		err = template.NewTemplateFieldError("Rate", err)
		return
	}

	if mt.NextId != 0 {
		to := template.GetTemplateService().Get(int(mt.NextId), (*MingGeMingPanTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		mt.nextMingPanTemplate = to.(*MingGeMingPanTemplate)
		if mt.nextMingPanTemplate.Number < mt.Number {
			err = fmt.Errorf("[%d] invalid", mt.Number)
			err = template.NewTemplateFieldError("Number", err)
			return
		}

		if mt.nextMingPanTemplate.Number == mt.Number {
			if mt.nextMingPanTemplate.Star-mt.Star != 1 {
				err = fmt.Errorf("[%d] invalid", mt.Star)
				err = template.NewTemplateFieldError("Star", err)
				return
			}
		}

	}

	return nil
}

func (mt *MingGeMingPanTemplate) PatchAfterCheck() {

}

func (mt *MingGeMingPanTemplate) FileName() string {
	return "tb_mingge_mingpan.json"
}

func init() {
	template.Register((*MingGeMingPanTemplate)(nil))
}
