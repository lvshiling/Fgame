package found_handler

import (
	"fgame/fgame/game/found/found"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	teamtypes "fgame/fgame/game/team/types"
	teamcopytemplate "fgame/fgame/game/teamcopy/template"
)

func init() {
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialTeamSilver, found.FoundObjDataHandlerFunc(getSilverFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialTeamLingTong, found.FoundObjDataHandlerFunc(getLingTongFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialTeamEquip, found.FoundObjDataHandlerFunc(getEquipFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialTeamStrengthen, found.FoundObjDataHandlerFunc(getStrengthenFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialTeamWeapon, found.FoundObjDataHandlerFunc(getWeaponFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialTeamXueMo, found.FoundObjDataHandlerFunc(getXueMoFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialTeamUpStar, found.FoundObjDataHandlerFunc(getUpStarFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialTeamXingChen, found.FoundObjDataHandlerFunc(getXingChenFoundParam))
}

func getSilverFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, teamtypes.TeamPurposeTypeFuBenSilver)
}

func getLingTongFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, teamtypes.TeamPurposeTypeFuBenLingTong)
}

func getEquipFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, teamtypes.TeamPurposeTypeFuBenZhuangShengEquip)
}

func getStrengthenFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, teamtypes.TeamPurposeTypeFuBenStrength)
}

func getWeaponFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, teamtypes.TeamPurposeTypeFuBenWeapon)
}

func getXueMoFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, teamtypes.TeamPurposeTypeFuBenXueMo)
}

func getUpStarFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, teamtypes.TeamPurposeTypeFuBenUpstar)
}

func getXingChenFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, teamtypes.TeamPurposeTypeFuBenXingChen)
}

func getParam(pl player.Player, typ teamtypes.TeamPurposeType) (resLevel int32, maxTimes int32, group int32) {
	group = int32(1)
	resLevel = pl.GetLevel()
	maxTimes = teamcopytemplate.GetTeamCopyTemplateService().GetTeamCopyRewardNumber(typ)
	return
}
