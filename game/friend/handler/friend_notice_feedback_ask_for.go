package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	friendlogic "fgame/fgame/game/friend/logic"
	"fgame/fgame/game/friend/pbutil"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_FEEDBACK_ASK_FOR_TYPE), dispatch.HandlerFunc(handleFriendFeedbackAskFor))
}

//处理向好友索取
func handleFriendFeedbackAskFor(s session.Session, msg interface{}) (err error) {
	log.Debug("friend:处理向好友索取")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSFriendFeedbackAskFor)
	friendId := csMsg.GetFriendId()
	itemId := csMsg.GetItemId()
	condition := csMsg.GetCondition()

	err = friendFeedbackAskFor(tpl, friendId, itemId, condition)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理向好友索取,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理向好友索取,完成")
	return nil
}

//处理向好友索取
func friendFeedbackAskFor(pl player.Player, friendId int64, itemId, condition int32) (err error) {
	//判断是否在线
	fri := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if fri == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理向好友索取,用户不在线")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoOnline)
		return
	}

	scMsg := pbutil.BuildSCFriendFeedbackAskFor()
	pl.SendMsg(scMsg)

	//送花
	friendlogic.FriendNoticeFeedbackNotice(fri, pl, friendtypes.FriendNoticeTypeKillBoss, friendtypes.FriendFeedbackTypeFlower, condition)

	//索取推送
	friScMsg := pbutil.BuildSCFriendFeedbackAskForNotice(pl.GetId(), pl.GetName(), pl.GetAllianceName(), int32(pl.GetRole()), int32(pl.GetSex()), itemId)
	fri.SendMsg(friScMsg)
	return
}
