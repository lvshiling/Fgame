package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/emperor/emperor"
	emperorlogic "fgame/fgame/game/emperor/logic"
	"fgame/fgame/game/emperor/pbutil"
	playeremperor "fgame/fgame/game/emperor/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EMPEROR_WORSHIP_TYPE), dispatch.HandlerFunc(handleEmperorWorship))
}

//处理膜拜信息
func handleEmperorWorship(s session.Session, msg interface{}) (err error) {
	log.Debug("emperor:处理膜拜信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = emperorWorship(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("emperor:处理膜拜信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("emperor:处理膜拜信息完成")
	return nil
}

//处理膜拜界面信息逻辑
func emperorWorship(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerEmperorDataManagerType).(*playeremperor.PlayerEmperorDataManager)
	playerId := pl.GetId()
	emperorId, _ := emperor.GetEmperorService().GetEmperorIdAndRobTime()
	//没有帝王不能膜拜
	if emperorId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("emperor:帝王无主")
		playerlogic.SendSystemMessage(pl, lang.EmperorWorshipReachLimit)
		return
	}

	//膜拜是否到达购买上限
	flag := manager.IfWorshipReachLimit()
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("emperor:今日膜拜次数已达上限")
		playerlogic.SendSystemMessage(pl, lang.EmperorWorshipReachLimit)
		return
	}

	if playerId == emperorId {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("emperor:不可以对自己进行膜拜")
		playerlogic.SendSystemMessage(pl, lang.EmperorNoWorshipMyself)
		return
	}
	//增加膜拜次数
	curNum := manager.AddWorshipNum()
	//玩家膜拜奖励
	emperorlogic.GiveWorshipReward(pl)
	storage := emperor.GetEmperorService().GetEmperorStorage()
	scEmperorWorship := pbuitl.BuildSCEmperorWorship(curNum, storage)
	pl.SendMsg(scEmperorWorship)
	return
}
