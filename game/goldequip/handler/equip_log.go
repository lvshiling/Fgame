package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_LOG_TYPE), dispatch.HandlerFunc(handleGoldEquipLog))
}

//处理日志元神金装
func handleGoldEquipLog(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理元神金装日志")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = goldEquipLog(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("goldequip:处理元神金装日志,错误")

		return err
	}
	log.Debug("goldequip:处理元神金装日志,完成")
	return nil
}

//日志
func goldEquipLog(pl player.Player) (err error) {
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	logList := goldequipManager.GetLogList()
	scMsg := pbutil.BuildSCGoldEquipLog(logList)
	pl.SendMsg(scMsg)
	return
}
