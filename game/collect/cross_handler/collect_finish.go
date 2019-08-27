package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	collectlogic "fgame/fgame/game/collect/logic"
	"fgame/fgame/game/collect/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_COLLECT_FINISH_TYPE), dispatch.HandlerFunc(handleCollectFinish))
}

//处理采集完成
func handleCollectFinish(s session.Session, msg interface{}) (err error) {
	log.Debug("collect:处理跨服采集完成")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isCollectFinish := msg.(*crosspb.ISCollectFinish)
	biologyId := isCollectFinish.GetBiologyId()
	err = collectFinish(tpl, biologyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Error("collect:处理跨服采集完成,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"biologyId": biologyId,
		}).Debug("tulong:处理跨服采集完成,完成")
	return nil

}

//跨服采集完成
func collectFinish(pl player.Player, biologyId int32) (err error) {
	log.Debug("collect:采集完成回跨服")
	dropItemList := collectlogic.CollectDropToInventory(pl, biologyId)
	itemMap := make(map[int32]int32)
	if len(dropItemList) > 0 {
		for _, dropItem := range dropItemList {
			itemMap[dropItem.GetItemId()] = dropItem.GetNum()
		}
	}

	siCollectFinish := pbutil.BuildSICollectFinish(itemMap)
	pl.SendCrossMsg(siCollectFinish)
	return
}
