package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/tower/pbutil"
	playertower "fgame/fgame/game/tower/player"
	"fgame/fgame/game/tower/tower"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TOWER_FLOOR_LIST_TYPE), dispatch.HandlerFunc(handlerTowerBossList))
}

//打宝塔BOSS列表请求
func handlerTowerBossList(s session.Session, msg interface{}) (err error) {
	log.Debug("towerBoss:处理打宝塔BOSS列表请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = towerBossList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("towerBoss:处理打宝塔BOSS列表请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("towerBoss:处理打宝塔BOSS列表请求完成")

	return
}

func towerBossList(pl player.Player) (err error) {
	towerManager := pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)
	remainTime := towerManager.GetRemainTime()
	bossMap := tower.GetTowerService().GetTowerBossList()
	scMsg := pbutil.BuildSCTowerFloorList(bossMap, remainTime)
	pl.SendMsg(scMsg)

	return
}
