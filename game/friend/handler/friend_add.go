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
	"fgame/fgame/game/friend/friend"
	playerfriend "fgame/fgame/game/friend/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_ADD_TYPE), dispatch.HandlerFunc(handleFriendAdd))
}

//处理好友添加
func handleFriendAdd(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友添加")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFriendAdd := msg.(*uipb.CSFriendAdd)
	agree := csFriendAdd.GetAgree()
	friendId := csFriendAdd.GetFriendId()
	_, err := friendAdd(tpl, agree, friendId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"agree":    agree,
				"error":    err,
			}).Error("friend:处理好友添加,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友添加,完成")
	return nil

}

//处理好友添加
func friendAdd(pl player.Player, agree bool, friendId int64) (flag bool, err error) {
	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	flag = friendManager.HasedInvite(friendId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"agree":    agree,
				"friendId": friendId,
			}).Warn("friend:不能加自己为好友")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if friendId == pl.GetId() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"agree":    agree,
				"friendId": friendId,
			}).Warn("friend:不能加自己为好友")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//不同意
	if !agree {
		friendManager.RemoveFriendInvite(friendId)
		gameevent.Emit(friendeventtypes.EventTypeFriendAddRefuse, pl, friendId)
		return
	}

	//判断是否可以加为好友
	flag = friendManager.ShouldAddFriend(friendId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友添加,已经是好友")
		playerlogic.SendSystemMessage(pl, lang.FriendAlreadyFriend)
		return
	}

	//TODO 判断用户是否存在
	playerInfo, err := player.GetPlayerService().GetPlayerInfo(friendId)
	if err != nil {
		return false, err
	}
	if playerInfo == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友添加,用户不存在")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoExist)
		return
	}

	flag = friendManager.IsBlack(friendId)
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友添加,对方被你拉黑了")
		playerlogic.SendSystemMessage(pl, lang.FriendIsBlack)
		return
	}
	friendManager.RemoveFriendInvite(friendId)

	//判断好友数量
	numFriends := friendManager.NumOfFriend()
	maxFriends := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFriendLimit)
	if numFriends >= int(maxFriends) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友添加,好友已经达到上限")
		playerlogic.SendSystemMessage(pl, lang.FriendAlreadyFull)
		return
	}

	_, err = friend.GetFriendService().AddFriend(pl, friendId)
	if err != nil {
		return
	}
	return
}
