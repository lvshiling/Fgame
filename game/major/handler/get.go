package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/major/pbutil"
	playermajor "fgame/fgame/game/major/player"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MAJOR_NUM_TYPE), dispatch.HandlerFunc(handleMajorNum))
}

//处理双修次数信息
func handleMajorNum(s session.Session, msg interface{}) (err error) {
	log.Debug("major:处理双修次数消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSMajorNum)
	majorInt := csMsg.GetMajorType()

	majorType := majortypes.MajorType(majorInt)
	if !majorType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"majorType": majorType,
			}).Warn("major:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = majorNum(tpl, majorType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("major:处理双修次数消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("major:处理双修次数消息完成")
	return nil
}

//处理双修次数信息逻辑
func majorNum(pl player.Player, majorType majortypes.MajorType) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMajorDataManagerType).(*playermajor.PlayerMajorDataManager)
	num := manager.GetMajorNum(majorType)
	scMajorNum := pbutil.BuildSCMajorNum(int32(majorType), num)
	pl.SendMsg(scMajorNum)
	return
}
