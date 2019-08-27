package check_enter

import (
	"fgame/fgame/game/foe/foe"
	"fgame/fgame/game/foe/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	xiantaologic "fgame/fgame/game/xiantao/logic"
	xiantaotypes "fgame/fgame/game/xiantao/types"
)

func init() {
	foe.RegisterFoeNoticeHandler(scenetypes.SceneTypeXianTaoDaHui, foe.FoeNoticeHandlerFunc(foeHandler))
}

func foeHandler(spl, sFoePl scene.Player, sceneType scenetypes.SceneType) (err error) {
	pl, ok := spl.(player.Player)
	if !ok {
		return
	}
	foePl, ok := sFoePl.(player.Player)
	if !ok {
		return
	}

	//掉落、获得
	dropNumMap, gainNumMap := xiantaologic.XianTaoDrop(pl, foePl)
	for typ := xiantaotypes.XianTaoTypeMin; typ <= xiantaotypes.XianTaoTypeMax; typ++ {
		_, dOK := dropNumMap[typ]
		if !dOK {
			dropNumMap[typ] = 0
		}
		_, gOK := gainNumMap[typ]
		if !gOK {
			gainNumMap[typ] = 0
		}
	}
	//被杀者推送 仇人信息推送
	beKillScMsg := pbutil.BuildSCFoeNoticeXianTao(foePl.GetId(), foePl.GetName(), foePl.GetRole(), foePl.GetSex(), int32(sceneType), dropNumMap)
	pl.SendMsg(beKillScMsg)
	//击杀者推送
	killerScMsg := pbutil.BuildSCFoeKillNoticeXianTao(pl.GetId(), pl.GetName(), pl.GetRole(), pl.GetSex(), int32(sceneType), gainNumMap)
	foePl.SendMsg(killerScMsg)
	return
}
