package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/friend/pbutil"
	playerfriend "fgame/fgame/game/friend/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_INVITE_TYPE), dispatch.HandlerFunc(handleFriendInvite))
}

//处理好友添加邀请邀请
func handleFriendInvite(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友添加邀请")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFriendInvite := msg.(*uipb.CSFriendInvite)
	friendId := csFriendInvite.GetFriendId()
	_, err := friendInvite(tpl, friendId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理好友添加邀请,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友添加邀请,完成")
	return nil
}

//处理好友添加邀请
func friendInvite(pl player.Player, friendId int64) (flag bool, err error) {
	if friendId == pl.GetId() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:不能加自己为好友")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)

	flag = friendManager.InviteFrequent(friendId)
	inviteTime := int64(0)
	if flag {
		playerlogic.SendSystemMessage(pl, lang.CommonOperFrequent)
		return
	} else {
		inviteTime = friendManager.InviteTime(friendId)
	}

	//判断是否可以加为好友
	flag = friendManager.ShouldAddFriend(friendId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友添加邀请,已经是好友")
		playerlogic.SendSystemMessage(pl, lang.FriendAlreadyFriend)
		return
	}

	//判断对方是否在线
	spl := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if spl == nil {
		//TODO 加了假人做特殊处理
		scFriendInvite := pbutil.BuildSCFriendInvite(friendId, inviteTime)
		pl.SendMsg(scFriendInvite)
		return
	}

	flag = friendManager.IsBlack(friendId)
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友添加邀请,对方在您的黑名单里")
		playerlogic.SendSystemMessage(pl, lang.FriendIsBlack)
		return
	}

	//判断好友数量
	numFriends := friendManager.NumOfFriend()
	maxFriends := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFriendLimit)
	if numFriends >= int(maxFriends) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友添加邀请,好友已经达到上限")
		playerlogic.SendSystemMessage(pl, lang.FriendAlreadyFull)
		return
	}

	scFriendInvite := pbutil.BuildSCFriendInvite(friendId, inviteTime)
	pl.SendMsg(scFriendInvite)

	//对方把你拉黑
	flag = friendManager.IsBlacked(friendId)
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Info("friend:处理好友添加邀请,对方已经把你拉黑")
		return
	}

	gameevent.Emit(friendeventtypes.EventTypeFriendInvite, pl, friendId)
	return
}
