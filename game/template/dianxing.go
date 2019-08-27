package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//升级配置
type DianXingTemplate struct {
	*DianXingTemplateVO
	useItemTemplate      *ItemTemplate                                      //升级物品
	nextDianXingTemplate *DianXingTemplate                                  //下一级
	battlePropertyMap    map[propertytypes.BattlePropertyType]int64         //属性
	externalPercentMap   map[playerpropertytypes.PropertyEffectorType]int32 //影响模块百分比
}

func (mclt *DianXingTemplate) TemplateId() int {
	return mclt.Id
}

func (mclt *DianXingTemplate) GetUseItemTemplate() *ItemTemplate {
	return mclt.useItemTemplate
}

func (mclt *DianXingTemplate) GetNextTemplate() *DianXingTemplate {
	return mclt.nextDianXingTemplate
}

func (mclt *DianXingTemplate) GetIsClear() bool {
	return mclt.IsClear != 0
}

func (mclt *DianXingTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return mclt.battlePropertyMap
}

func (mclt *DianXingTemplate) GetExternalPercentMap() map[playerpropertytypes.PropertyEffectorType]int32 {
	return mclt.externalPercentMap
}

func (mclt *DianXingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//下一级
	if mclt.NextId != 0 {
		tempNextDianXingTemplate := template.GetTemplateService().Get(int(mclt.NextId), (*DianXingTemplate)(nil))
		if tempNextDianXingTemplate == nil {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		mclt.nextDianXingTemplate = tempNextDianXingTemplate.(*DianXingTemplate)
	}

	//验证 UseItem
	if mclt.UseItem != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(mclt.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", mclt.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		mclt.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(mclt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mclt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	//属性
	mclt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	mclt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(mclt.Hp)
	mclt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(mclt.Attack)
	mclt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(mclt.Defence)
	//影响模块百分比
	mclt.externalPercentMap = make(map[playerpropertytypes.PropertyEffectorType]int32)
	mclt.externalPercentMap[playerpropertytypes.PlayerPropertyEffectorTypeMount] = mclt.MountPercent
	mclt.externalPercentMap[playerpropertytypes.PlayerPropertyEffectorTypeWing] = mclt.WingPercent
	mclt.externalPercentMap[playerpropertytypes.PlayerPropertyEffectorTypeLingyu] = mclt.FieldPercent
	mclt.externalPercentMap[playerpropertytypes.PlayerPropertyEffectorTypeShenfa] = mclt.ShenFaPercent
	mclt.externalPercentMap[playerpropertytypes.PlayerPropertyEffectorTypeBodyShield] = mclt.BodyShieldPercent
	mclt.externalPercentMap[playerpropertytypes.PlayerPropertyEffectorTypeAnqi] = mclt.AnQiPercent
	mclt.externalPercentMap[playerpropertytypes.PlayerPropertyEffectorTypeFaBao] = mclt.FaBaoPercent
	mclt.externalPercentMap[playerpropertytypes.PlayerPropertyEffectorTypeXianTi] = mclt.XianTiPercent

	return nil
}

func (mclt *DianXingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(mclt.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证星谱
	err = validator.MinValidate(float64(mclt.XingPuType), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.XingPuType)
		err = template.NewTemplateFieldError("XingPuType", err)
		return
	}

	//验证 next_id
	if mclt.NextId != 0 {
		diff := mclt.NextId - int32(mclt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(mclt.NextId), (*DianXingTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		dianXingTemplate := to.(*DianXingTemplate)
		diffType := dianXingTemplate.XingPuType - mclt.XingPuType
		if diffType == 0 {
			diffLevel := dianXingTemplate.Level - mclt.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", mclt.Level)
				return template.NewTemplateFieldError("Level", err)
			}
		} else if diffType != 1 {
			err = fmt.Errorf("[%d] invalid", mclt.XingPuType)
			return template.NewTemplateFieldError("XingPuType", err)
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(mclt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_silver
	err = validator.MinValidate(float64(mclt.UseSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.UseSilver)
		err = template.NewTemplateFieldError("UseSilver", err)
		return
	}

	//验证use_xingchen
	err = validator.MinValidate(float64(mclt.UseXingChen), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.UseXingChen)
		err = template.NewTemplateFieldError("UseXingChen", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(mclt.TimesMin), float64(0), true, float64(mclt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mclt.TimesMax), float64(mclt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(mclt.AddMin), float64(0), true, float64(mclt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(mclt.AddMax), float64(mclt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(mclt.ZhufuMax), float64(mclt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 hp
	err = validator.MinValidate(float64(mclt.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(mclt.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(mclt.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	//验证 is_clear
	err = validator.MinValidate(float64(mclt.IsClear), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.IsClear)
		err = template.NewTemplateFieldError("IsClear", err)
		return
	}

	//验证MountPercent
	err = validator.RangeValidate(float64(mclt.MountPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.MountPercent)
		err = template.NewTemplateFieldError("MountPercent", err)
		return
	}

	//验证WingPercent
	err = validator.RangeValidate(float64(mclt.WingPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.WingPercent)
		err = template.NewTemplateFieldError("WingPercent", err)
		return
	}

	//验证FieldPercent
	err = validator.RangeValidate(float64(mclt.FieldPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.FieldPercent)
		err = template.NewTemplateFieldError("FieldPercent", err)
		return
	}

	//验证ShenFaPercent
	err = validator.RangeValidate(float64(mclt.ShenFaPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.ShenFaPercent)
		err = template.NewTemplateFieldError("ShenFaPercent", err)
		return
	}

	//验证BodyShieldPercent
	err = validator.RangeValidate(float64(mclt.BodyShieldPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.BodyShieldPercent)
		err = template.NewTemplateFieldError("BodyShieldPercent", err)
		return
	}

	//验证AnQiPercent
	err = validator.RangeValidate(float64(mclt.AnQiPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.AnQiPercent)
		err = template.NewTemplateFieldError("AnQiPercent", err)
		return
	}

	//验证FaBaoPercent
	err = validator.RangeValidate(float64(mclt.FaBaoPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.FaBaoPercent)
		err = template.NewTemplateFieldError("FaBaoPercent", err)
		return
	}

	//验证XianTiPercent
	err = validator.RangeValidate(float64(mclt.XianTiPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.XianTiPercent)
		err = template.NewTemplateFieldError("XianTiPercent", err)
		return
	}

	return nil
}
func (mclt *DianXingTemplate) PatchAfterCheck() {

}
func (mclt *DianXingTemplate) FileName() string {
	return "tb_dianxing.json"
}

func init() {
	template.Register((*DianXingTemplate)(nil))
}
