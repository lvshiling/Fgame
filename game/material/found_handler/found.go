package found_handler

import (
	"fgame/fgame/game/found/found"
	foundtypes "fgame/fgame/game/found/types"
	materialplayer "fgame/fgame/game/material/player"
	materialtemplate "fgame/fgame/game/material/template"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialMount, found.FoundObjDataHandlerFunc(getMountFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialWing, found.FoundObjDataHandlerFunc(getWingFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialShenFa, found.FoundObjDataHandlerFunc(getShenFaFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialFaBao, found.FoundObjDataHandlerFunc(getFaBaoFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialXianTi, found.FoundObjDataHandlerFunc(getXianTiFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialTianMo, found.FoundObjDataHandlerFunc(getTianMoFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialLingTong, found.FoundObjDataHandlerFunc(getLingTongFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialLingTongWeapon, found.FoundObjDataHandlerFunc(getLingTongWeaponFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMaterialLingTongLingYu, found.FoundObjDataHandlerFunc(getLingTongLingYuFoundParam))
}

func getWingFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, materialtypes.MaterialTypeWing)
}

func getMountFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, materialtypes.MaterialTypeMount)
}

func getShenFaFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, materialtypes.MaterialTypeShenfa)
}

func getFaBaoFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, materialtypes.MaterialTypeFabao)
}

func getXianTiFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, materialtypes.MaterialTypeXianti)
}

func getTianMoFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, materialtypes.MaterialTypeTianMo)
}

func getLingTongFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, materialtypes.MaterialTypeLingTong)
}

func getLingTongWeaponFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, materialtypes.MaterialTypeLingBing)
}

func getLingTongLingYuFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, materialtypes.MaterialTypeLingYu)
}

func getParam(pl player.Player, typ materialtypes.MaterialType) (resLevel int32, maxTimes int32, group int32) {
	materialManager := pl.GetPlayerDataManager(playertypes.PlayerMaterialDataManagerType).(*materialplayer.PlayerMaterialDataManager)
	obj := materialManager.GetPlayerMaterialInfo(typ)
	if obj == nil {
		return
	}
	group = obj.GetGroup()
	if group == 0 {
		group = 1
	}

	resLevel = pl.GetLevel()
	maxTimes = materialtemplate.GetMaterialTemplateService().GetFreePlayTimes(typ)

	return
}
