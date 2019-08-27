package found_handler

import (
	"fgame/fgame/game/found/found"
	foundtypes "fgame/fgame/game/found/types"
	majortemplate "fgame/fgame/game/major/template"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
)

func init() {
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMajorShuangXiu, found.FoundObjDataHandlerFunc(getShuangXiuFoundParam))
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeMajorFuQi, found.FoundObjDataHandlerFunc(getFuQiFoundParam))
}

func getShuangXiuFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, majortypes.MajorTypeShuangXiu)
}

func getFuQiFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, majortypes.MajorTypeFuQi)
}

func getParam(pl player.Player, typ majortypes.MajorType) (resLevel int32, maxTimes int32, group int32) {
	group = int32(1)
	resLevel = pl.GetLevel()
	maxTimes = majortemplate.GetMajorTemplateService().GetMajorDefaultMaxNum(typ)
	return
}
