package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/dan/pbutil"
	playerdan "fgame/fgame/game/dan/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DAN_UPGRADE_TYPE), dispatch.HandlerFunc(handleDanUpgrade))
}

//处理食丹升级
func handleDanUpgrade(s session.Session, msg interface{}) error {
	log.Debug("dan:处理食丹升级消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err := danUpgrade(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("dan:处理食丹升级消息,错误")
		return err
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("dan:处理食丹升级消息完成")

	return nil
}

//处理食丹升级的逻辑
func danUpgrade(pl player.Player) (err error) {
	danManager := pl.GetPlayerDataManager(types.PlayerDanDataManagerType).(*playerdan.PlayerDanDataManager)
	flag := danManager.CheckFullLevel()
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("dan:当前食丹等级已达最高级，无法再升级")
		playerlogic.SendSystemMessage(pl, lang.DanLevelReachedLimit)
		return
	}
	//校验已食用的丹药数
	flag = danManager.IfEatEnough()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("dan:当前食用丹药进度未满，无法升级")
		playerlogic.SendSystemMessage(pl, lang.DanUseIsNotEnough)
		return
	}

	flag = danManager.Upgrade()
	if !flag {
		panic("dan:食丹升级应该是成功的")
	}

	danInfo := danManager.GetDanInfo()
	scDanUpgrade := pbuitl.BuildSCDanUpgrade(danInfo.LevelId)
	pl.SendMsg(scDanUpgrade)
	return
}
