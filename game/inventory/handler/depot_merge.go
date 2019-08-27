package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DEPOT_MERGE_TYPE), dispatch.HandlerFunc(handleDepotMerge))
}

//处理仓库合并
func handleDepotMerge(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理合并仓库")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = depotMerge(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理合并仓库,错误")

		return
	}
	log.Debug("inventory:处理合并仓库,完成")
	return
}

//合并
func depotMerge(pl player.Player) (err error) {
	// TODO:xzk:整理CD 10秒  您刚整理过背包，x秒方可再次操作
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	inventoryManager.MergeDepot()

	scDepotMerge := pbutil.BuildSCDepotMerge(inventoryManager.GetDepotAll())
	pl.SendMsg(scDepotMerge)
	return
}
