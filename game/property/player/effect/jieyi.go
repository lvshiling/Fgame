package effect

import (
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/jieyi/jieyi"
	playerjieyi "fgame/fgame/game/jieyi/player"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	"math"

	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeJieYi, JieYiPropertyEffect)
}

// 结义作用器
func JieYiPropertyEffect(pl player.Player, prop *propertycommon.SystemPropertySegment) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeJieYi) {
		return
	}

	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	if !jieYiManager.IsJieYi() {
		return
	}

	hp := int64(0)
	attack := int64(0)
	defence := int64(0)

	info := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	if info == nil {
		return
	}
	jieYiId := info.GetJieYiId()
	memberList := jieyi.GetJieYiService().GetJieYiMemberList(jieYiId)

	token := info.GetTokenType()
	tokenLev := info.GetTokenLev()
	tokenLevTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiTokenLevelTemplate(token, tokenLev)

	nameTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiNameTemplate(jieYiManager.GetPlayerJieYiObj().GetNameLev())

	// 威名属性加成
	if nameTemp != nil {
		hp += nameTemp.Hp
		attack += nameTemp.Attack
		defence += nameTemp.Defence
	}

	// 自身信物属性加成
	if info.HasToken() {
		if tokenLevTemp != nil {
			hp += tokenLevTemp.Hp
			attack += tokenLevTemp.Attack
			defence += tokenLevTemp.Defence
		}
	}

	// 兄弟信物属性加成
	for _, member := range memberList {
		if !member.HasToken() {
			continue
		}
		if member.GetPlayerId() == pl.GetId() {
			continue
		}

		token = member.GetTokenType()
		tokenLev = member.GetTokenLev()
		tokenTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiTokenTemplate(token)
		tokenLevTemp = jieyitemplate.GetJieYiTemplateService().GetJieYiTokenLevelTemplate(token, tokenLev)

		mHp := int64(0)
		mAttack := int64(0)
		mDefence := int64(0)
		share := int64(0)

		if tokenTemp != nil {
			share = int64(tokenTemp.SharePercent)
		}

		if tokenLevTemp != nil {
			mHp += tokenLevTemp.Hp
			mAttack += tokenLevTemp.Attack
			mDefence += tokenLevTemp.Defence
		}

		hp += int64(math.Ceil(float64(mHp) * float64(share) / float64(common.MAX_RATE)))
		attack += int64(math.Ceil(float64(mAttack) * float64(share) / float64(common.MAX_RATE)))
		defence += int64(math.Ceil(float64(mDefence) * float64(share) / float64(common.MAX_RATE)))
	}

	prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, hp)
	prop.SetBase(propertytypes.BattlePropertyTypeAttack, attack)
	prop.SetBase(propertytypes.BattlePropertyTypeDefend, defence)

}
