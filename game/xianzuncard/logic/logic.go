package logic

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	gametemplate "fgame/fgame/game/template"
	xianfutypes "fgame/fgame/game/xianfu/types"
	playerxianzuncard "fgame/fgame/game/xianzuncard/player"
	xianzuncardtemplate "fgame/fgame/game/xianzuncard/template"
	xianzuncardtypes "fgame/fgame/game/xianzuncard/types"
)

//变更战翼 坐骑属性
func PropertyChanged(pl player.Player) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeWing.Mask() | playerpropertytypes.PlayerPropertyEffectorTypeMount.Mask())
	// propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeMount.Mask())
	return
}

//仙府副本次数特权
func XianFuFreeTimes(pl player.Player, xianFuType xianfutypes.XianfuType) int32 {
	// 是否白银仙尊
	xianZunManager := pl.GetPlayerDataManager(playertypes.PlayerXianZunCardManagerType).(*playerxianzuncard.PlayerXianZunCardDataManager)
	if !xianZunManager.IsActivite(xianzuncardtypes.XianZunCardTypeSliver) {
		return 0
	}

	// 取模板配置
	xianZunTemp := xianzuncardtemplate.GetXianZunCardTemplateService().GetXianZunCardTemplate(xianzuncardtypes.XianZunCardTypeSliver)
	if xianZunTemp == nil {
		return 0
	}

	return xianZunTemp.GetXianFuFreeTimes(xianFuType)
}

//珍稀BOSS特权
func ZhenXiBossFreeTimes(pl player.Player) int32 {
	xianZunManager := pl.GetPlayerDataManager(playertypes.PlayerXianZunCardManagerType).(*playerxianzuncard.PlayerXianZunCardDataManager)
	if !xianZunManager.IsActivite(xianzuncardtypes.JieYiDaoJuTypeDiamond) {
		return 0
	}

	// 取模板配置
	xianZunTemp := xianzuncardtemplate.GetXianZunCardTemplateService().GetXianZunCardTemplate(xianzuncardtypes.JieYiDaoJuTypeDiamond)
	if xianZunTemp == nil {
		return 0
	}

	return xianZunTemp.TianJieBossFreeAdd
}

//3v3积分加成
func ArenaExtralPercent(pl player.Player) (baseMaxPercent int32, baseAddPercent int32) {
	xianZunManager := pl.GetPlayerDataManager(playertypes.PlayerXianZunCardManagerType).(*playerxianzuncard.PlayerXianZunCardDataManager)
	if !xianZunManager.IsActivite(xianzuncardtypes.JieYiDaoJuTypeDiamond) {
		return
	}

	// 取模板配置
	xianZunTemp := xianzuncardtemplate.GetXianZunCardTemplateService().GetXianZunCardTemplate(xianzuncardtypes.JieYiDaoJuTypeDiamond)
	if xianZunTemp == nil {
		return
	}

	return xianZunTemp.JiFenMaxAddPercent, xianZunTemp.JiFenAddPercent
}

//野外怪经验加成
func MonsterExpExtralPercent(pl player.Player, monsterId int32) (expPercent int32) {
	s := pl.GetScene()
	if s == nil {
		return
	}

	xianZunManager := pl.GetPlayerDataManager(playertypes.PlayerXianZunCardManagerType).(*playerxianzuncard.PlayerXianZunCardDataManager)
	if !xianZunManager.IsActivite(xianzuncardtypes.JieYiDaoJuTypeGold) {
		return
	}

	to := template.GetTemplateService().Get(int(monsterId), (*gametemplate.BiologyTemplate)(nil))
	if to == nil {
		return
	}
	biologyTemp := to.(*gametemplate.BiologyTemplate)

	// 取模板配置
	xianZunTemp := xianzuncardtemplate.GetXianZunCardTemplateService().GetXianZunCardTemplate(xianzuncardtypes.JieYiDaoJuTypeGold)
	if xianZunTemp == nil {
		return
	}

	// 判断生物类型是否享受经验加成
	typList := xianZunTemp.GetExpBiologyList()
	for _, typ := range typList {
		if typ == biologyTemp.GetBiologySetType() {
			return xianZunTemp.ExpBiologyAddPercent
		}
	}

	return 0
}
