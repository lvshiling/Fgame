package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	friendlogic "fgame/fgame/game/friend/logic"
	"fgame/fgame/game/friend/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_FEEDBACK_ASK_FOR_REPLY_TYPE), dispatch.HandlerFunc(handleFriendFeedbackAskForReply))
}

//处理好友索取
func handleFriendFeedbackAskForReply(s session.Session, msg interface{}) (err error) {
	log.Debug("friend:处理好友索取")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSFriendFeedbackAskForReply)
	isAgree := csMsg.GetIsAgree()
	itemId := csMsg.GetItemId()
	friendId := csMsg.GetFriendId()

	err = friendFeedbackAskForReply(tpl, isAgree, friendId, itemId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理好友索取,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友索取,完成")
	return nil
}

//处理好友索取
func friendFeedbackAskForReply(pl player.Player, isAgree bool, friendId int64, itemId int32) (err error) {
	//判断是否在线
	fri := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if fri == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友索取,用户不在线")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoOnline)
		return
	}

	//索取应答推送
	friendlogic.GiveItemForFriend(pl, fri, itemId, isAgree)

	scMsg := pbutil.BuildSCFriendFeedbackAskForReply(isAgree)
	pl.SendMsg(scMsg)
	return
}
