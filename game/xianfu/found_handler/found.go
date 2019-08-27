package found_handler

import (
	"fgame/fgame/game/found/found"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	xianfuplayer "fgame/fgame/game/xianfu/player"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"
)

func init() {
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeExp, found.FoundObjDataHandlerFunc(getExpFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeSilver, found.FoundObjDataHandlerFunc(getSilverFoundParam))
}

func getExpFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, xianfutypes.XianfuTypeExp)
}

func getSilverFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, xianfutypes.XianfuTypeSilver)
}

func getParam(pl player.Player, typ xianfutypes.XianfuType) (resLevel int32, maxTimes int32, group int32) {
	xianfuManager := pl.GetPlayerDataManager(playertypes.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	group = xianfuManager.GetGroup(typ)
	if group == 0 {
		group = 1
	}
	resLevel = xianfuManager.GetXianfuId(typ)
	maxTimes = xianfutemplate.GetXianfuTemplateService().GetFreePlayTimes(typ, resLevel)
	return
}
