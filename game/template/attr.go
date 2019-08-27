package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"

	"fgame/fgame/game/property/types"
)

// [
//   {
//     "id": 1,
//     "hp": 10,
//     "attack": 10000,
//     "defence": 10000,
//     "critical": 10000,
//     "tough": 10000,
//     "block": 10000,
//     "break": 10000,
//     "hit": 10000,
//     "dodge": 10000,
//     "hunyuan_att": 10000,
//     "hunyuan_def": 10000,
//     "bingdong_res": 10000,
//     "pojia_res": 10000,
//     "kuilei_res": 10000,
//     "kujie_res": 10000,
//     "shiming_res": 10000,
//     "hunmi_res": 10000,
//     "xuruo_res": 10000,
//     "jiaoxie_res": 10000,
//     "zhongdu_res": 10000
//   }
// ]

type AttrTemplate struct {
	*AttrTemplateVO
	battlePropertyMap        map[types.BattlePropertyType]int64
	battlePropertyPercentMap map[types.BattlePropertyType]int64
}

func (at *AttrTemplate) GetAllBattleProperty() map[types.BattlePropertyType]int64 {
	return at.battlePropertyMap
}

func (at *AttrTemplate) GetBattlePropertyPercent() map[types.BattlePropertyType]int64 {
	return at.battlePropertyPercentMap
}

func (at *AttrTemplate) TemplateId() int {
	return at.Id
}

func (at *AttrTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(at.FileName(), at.TemplateId(), err)
			return
		}
	}()
	at.battlePropertyMap = make(map[types.BattlePropertyType]int64)
	at.battlePropertyMap[types.BattlePropertyTypeMaxHP] = int64(at.Hp)
	at.battlePropertyMap[types.BattlePropertyTypeAttack] = int64(at.Attack)
	at.battlePropertyMap[types.BattlePropertyTypeDefend] = int64(at.Defence)
	at.battlePropertyMap[types.BattlePropertyTypeCrit] = int64(at.Critical)
	at.battlePropertyMap[types.BattlePropertyTypeTough] = int64(at.Tough)
	at.battlePropertyMap[types.BattlePropertyTypeAbnormality] = int64(at.Abnormality)
	at.battlePropertyMap[types.BattlePropertyTypeBlock] = int64(at.Block)
	at.battlePropertyMap[types.BattlePropertyTypeDodge] = int64(at.Dodge)
	at.battlePropertyMap[types.BattlePropertyTypeHit] = int64(at.Hit)
	at.battlePropertyMap[types.BattlePropertyTypeHuanYunAttack] = int64(at.HunyuanAtt)
	at.battlePropertyMap[types.BattlePropertyTypeHuanYunDef] = int64(at.HunyuanDef)
	at.battlePropertyMap[types.BattlePropertyTypeBingDongRes] = int64(at.BingdongRes)
	at.battlePropertyMap[types.BattlePropertyTypePoJiaRes] = int64(at.PojiaRes)
	at.battlePropertyMap[types.BattlePropertyTypeKuiLeiRes] = int64(at.KuileiRes)
	at.battlePropertyMap[types.BattlePropertyTypeKuJieRes] = int64(at.KujieRes)
	at.battlePropertyMap[types.BattlePropertyTypeShiMingRes] = int64(at.ShimingRes)
	at.battlePropertyMap[types.BattlePropertyTypeHunMiRes] = int64(at.HunmiRes)
	at.battlePropertyMap[types.BattlePropertyTypeXuRuoRes] = int64(at.XuruoRes)
	at.battlePropertyMap[types.BattlePropertyTypeJiaoXieRes] = int64(at.JiaoxieRes)
	at.battlePropertyMap[types.BattlePropertyTypeZhongDuRes] = int64(at.ZhongduRes)
	at.battlePropertyMap[types.BattlePropertyTypeMoveSpeed] = int64(at.MoveSpeed)
	at.battlePropertyMap[types.BattlePropertyTypeFanTan] = int64(at.FantanAdd)
	at.battlePropertyMap[types.BattlePropertyTypeFanTanPercent] = int64(at.FantanPercent)
	at.battlePropertyMap[types.BattlePropertyTypeCritRatePercent] = int64(at.CritRatePercent)
	at.battlePropertyMap[types.BattlePropertyTypeCritHarmPercent] = int64(at.CritHarmPercent)
	at.battlePropertyMap[types.BattlePropertyTypeHitRatePercent] = int64(at.HitRatePercent)
	at.battlePropertyMap[types.BattlePropertyTypeDodgeRatePercent] = int64(at.DodgeRatePercent)
	at.battlePropertyMap[types.BattlePropertyTypeSpellCdPercent] = int64(at.SpellCdPercent)
	at.battlePropertyMap[types.BattlePropertyTypeAddExp] = int64(at.AddExp)
	at.battlePropertyMap[types.BattlePropertyTypeBlockRatePercent] = int64(at.BlockRatePercent)
	at.battlePropertyMap[types.BattlePropertyTypeZhuoShaoRes] = int64(at.ZhuoshaoRes)
	at.battlePropertyMap[types.BattlePropertyTypeJianSuRes] = int64(at.JiansuRes)
	at.battlePropertyMap[types.BattlePropertyTypeDingShenRes] = int64(at.DingshenRes)

	at.battlePropertyPercentMap = make(map[types.BattlePropertyType]int64)
	at.battlePropertyPercentMap[types.BattlePropertyTypeMaxHP] = int64(at.HpPercent)
	at.battlePropertyPercentMap[types.BattlePropertyTypeAttack] = int64(at.AttackPercent)
	at.battlePropertyPercentMap[types.BattlePropertyTypeDefend] = int64(at.DefPercent)
	return nil
}

func (at *AttrTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(at.FileName(), at.TemplateId(), err)
			return
		}
	}()
	for key, value := range at.battlePropertyMap {
		//验证Time
		err = validator.MinValidate(float64(value), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", value)
			return template.NewTemplateFieldError(key.String(), err)
		}
	}

	for key, value := range at.battlePropertyPercentMap {
		err = validator.MinValidate(float64(value), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", value)
			return template.NewTemplateFieldError(key.String(), err)
		}
	}
	return nil
}

func (at *AttrTemplate) PatchAfterCheck() {

}

func (at *AttrTemplate) FileName() string {
	return "tb_attr.json"
}

func init() {
	template.Register((*AttrTemplate)(nil))
}
