package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"

	"fgame/fgame/game/property/types"
)

type ForceMouldTemplate struct {
	*ForceMouldTemplateVO
	forcePropertyMap map[types.BattlePropertyType]int64
}

func (at *ForceMouldTemplate) GetAllForceProperty() map[types.BattlePropertyType]int64 {
	return at.forcePropertyMap
}

func (at *ForceMouldTemplate) TemplateId() int {
	return at.Id
}

func (at *ForceMouldTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(at.FileName(), at.TemplateId(), err)
			return
		}
	}()
	at.forcePropertyMap = make(map[types.BattlePropertyType]int64)
	at.forcePropertyMap[types.BattlePropertyTypeMaxHP] = int64(at.Hp)
	at.forcePropertyMap[types.BattlePropertyTypeAttack] = int64(at.Attack)
	at.forcePropertyMap[types.BattlePropertyTypeDefend] = int64(at.Defence)
	at.forcePropertyMap[types.BattlePropertyTypeCrit] = int64(at.Critical)
	at.forcePropertyMap[types.BattlePropertyTypeTough] = int64(at.Tough)
	at.forcePropertyMap[types.BattlePropertyTypeAbnormality] = int64(at.Abnormality)
	at.forcePropertyMap[types.BattlePropertyTypeBlock] = int64(at.Block)
	at.forcePropertyMap[types.BattlePropertyTypeDodge] = int64(at.Dodge)
	at.forcePropertyMap[types.BattlePropertyTypeHit] = int64(at.Hit)
	at.forcePropertyMap[types.BattlePropertyTypeHuanYunAttack] = int64(at.HunyuanAtt)
	at.forcePropertyMap[types.BattlePropertyTypeHuanYunDef] = int64(at.HunyuanDef)
	at.forcePropertyMap[types.BattlePropertyTypeBingDongRes] = int64(at.BingdongRes)
	at.forcePropertyMap[types.BattlePropertyTypePoJiaRes] = int64(at.PojiaRes)
	at.forcePropertyMap[types.BattlePropertyTypeKuiLeiRes] = int64(at.KuileiRes)
	at.forcePropertyMap[types.BattlePropertyTypeKuJieRes] = int64(at.KujieRes)
	at.forcePropertyMap[types.BattlePropertyTypeShiMingRes] = int64(at.ShimingRes)
	at.forcePropertyMap[types.BattlePropertyTypeHunMiRes] = int64(at.HunmiRes)
	at.forcePropertyMap[types.BattlePropertyTypeXuRuoRes] = int64(at.XuruoRes)
	at.forcePropertyMap[types.BattlePropertyTypeJiaoXieRes] = int64(at.JiaoxieRes)
	at.forcePropertyMap[types.BattlePropertyTypeZhongDuRes] = int64(at.ZhongduRes)
	at.forcePropertyMap[types.BattlePropertyTypeMoveSpeed] = int64(at.SpeedMove)
	return nil
}

func (at *ForceMouldTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(at.FileName(), at.TemplateId(), err)
			return
		}
	}()
	for key, value := range at.forcePropertyMap {
		//验证Time
		err = validator.MinValidate(float64(value), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", value)
			return template.NewTemplateFieldError(key.String(), err)
		}
	}
	return nil
}

func (at *ForceMouldTemplate) PatchAfterCheck() {

}

func (at *ForceMouldTemplate) FileName() string {
	return "tb_force_mould.json"
}

func init() {
	template.Register((*ForceMouldTemplate)(nil))
}
