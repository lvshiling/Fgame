package logic

import (
	bufftemplate "fgame/fgame/game/buff/template"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func RemoveBuffByAction(bo scene.BattleObject, removeType scenetypes.BuffRemoveType) {
	for _, b := range bo.GetBuffs() {
		buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(b.GetBuffId())
		if buffTemplate == nil {
			continue
		}
		if buffTemplate.TypeRemove&removeType.Mask() != 0 {
			bo.RemoveBuff(b.GetBuffId())
		}
	}
	return
}

func TouchBuffByAction(bo scene.BattleObject, touchType scenetypes.BuffTouchType) {
	for _, b := range bo.GetBuffs() {
		buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(b.GetBuffId())
		if buffTemplate == nil {
			continue
		}
		if buffTemplate.GetTouchType() == touchType {
			bo.TouchBuff(b.GetBuffId())

		}
	}
	return
}

//更新属性
func UpdateBattleProperty(bo scene.BattleObject) {
	bo.UpdateBuffProperty()
	bo.Calculate()
	return
}
