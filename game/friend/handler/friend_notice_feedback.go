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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_NOTICE_FEEDBACK_TYPE), dispatch.HandlerFunc(handleFriendNoticeFeedback))
}

//处理赞赏好友
func handleFriendNoticeFeedback(s session.Session, msg interface{}) (err error) {
	log.Debug("friend:处理赞赏好友")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSFriendNoticeFeedback)
	friendId := csMsg.GetFriendId()
	nType := csMsg.GetNoticeType()
	fType := csMsg.GetFeedbackType()
	condition := csMsg.GetCondition()

	noticeType := friendtypes.FriendNoticeType(nType)
	if !noticeType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"NoticeType": nType,
			}).Warn("friend:处理赞赏好友,提醒类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	feedbackType := friendtypes.FriendFeedbackType(fType)
	if !feedbackType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"NoticeType": nType,
			}).Warn("friend:处理赞赏好友,赞赏类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = friendNoticeFeedback(tpl, friendId, noticeType, feedbackType, condition)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理赞赏好友,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理赞赏好友,完成")
	return nil
}

//处理赞赏好友
func friendNoticeFeedback(pl player.Player, friendId int64, noticeType friendtypes.FriendNoticeType, feedbackType friendtypes.FriendFeedbackType, condition int32) (err error) {
	//判断是否在线
	fri := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if fri == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理赞赏好友,用户不在线")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoOnline)
		return
	}

	scMsg := pbutil.BuildSCFriendNoticeFeedback(feedbackType, condition)
	pl.SendMsg(scMsg)

	friendlogic.FriendNoticeFeedbackNotice(fri, pl, noticeType, feedbackType, condition)
	return
}
