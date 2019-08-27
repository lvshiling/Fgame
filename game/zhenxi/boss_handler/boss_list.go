package boss_handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shareboss/shareboss"
	worldbosspbutil "fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
	zhenxipbutil "fgame/fgame/game/zhenxi/pbutil"
	playerzhenxi "fgame/fgame/game/zhenxi/player"
)

func init() {
	worldboss.RegistBossListHandler(worldbosstypes.BossTypeZhenXi, worldboss.BossListHandlerFunc(zhenXiBossList))
}

func zhenXiBossList(pl player.Player, typ worldbosstypes.BossType) {

	bossList := shareboss.GetShareBossService().GetShareBossList(typ)
	scShareBossList := worldbosspbutil.BuildSCWorldBossListShareBoss(bossList, int32(typ))
	pl.SendMsg(scShareBossList)

	zhenXinManager := pl.GetPlayerDataManager(playertypes.PlayerZhenXiDataManagerType).(*playerzhenxi.PlayerZhenXiDataManager)
	obj := zhenXinManager.GetPlayerZhenXiObject()
	scPlayerZhenXiBossInfo := zhenxipbutil.BuildSCPlayerZhenXiBossInfo(obj.GetEnterTimes())
	pl.SendMsg(scPlayerZhenXiBossInfo)
}
