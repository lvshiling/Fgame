package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/friend/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_GIFT_FEEDBACK_TYPE), dispatch.HandlerFunc(handleFriendGiftFeedback))
}

//处理好友礼物反馈
func handleFriendGiftFeedback(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友礼物反馈")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFriendGiftFeedback := msg.(*uipb.CSFriendGiftFeedback)
	friendId := csFriendGiftFeedback.GetFriendId()

	err := friendGiftFeedback(tpl, friendId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理好友礼物反馈,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友礼物反馈,完成")
	return nil

}

//处理好友礼物反馈
func friendGiftFeedback(pl player.Player, friendId int64) (err error) {
	//判断是否在线
	fri := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if fri == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友礼物,用户不在线")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoOnline)
		return
	}

	scFriendGiftFeedback := pbutil.BuildSCFriendGiftFeedback(friendId)
	scFriendGiftFeedbackRecv := pbutil.BuildSCFriendGiftFeedbackRecv(pl.GetId())
	pl.SendMsg(scFriendGiftFeedback)
	fri.SendMsg(scFriendGiftFeedbackRecv)
	return
}
