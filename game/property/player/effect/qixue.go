package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playerqixue "fgame/fgame/game/qixue/player"
	qixuetemplate "fgame/fgame/game/qixue/template"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeQiXue, qiXuePropertyEffect)
}

//泣血枪作用器
func qiXuePropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeQiXueQiang) {
		return
	}

	qiXueManager := p.GetPlayerDataManager(playertypes.PlayerQiXueDataManagerType).(*playerqixue.PlayerQiXueDataManager)
	qiXueInfo := qiXueManager.GetQiXueInfo()
	qiXueTemplate := qixuetemplate.GetQiXueTemplateService().GetQiXueTemplateByLevel(qiXueInfo.GetLevel(), qiXueInfo.GetStar())
	//戮仙刃系统默认不开启
	if qiXueTemplate == nil {
		return
	}

	//泣血枪属性
	for key, val := range qiXueTemplate.GetBattleProperty() {
		val += prop.GetBase(key)
		prop.SetBase(key, val)
	}
}
