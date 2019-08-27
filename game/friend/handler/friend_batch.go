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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_BATCH_TYPE), dispatch.HandlerFunc(handleFriendBatch))
}

//处理好友批量决策
func handleFriendBatch(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友批量决策")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFriendBatch := msg.(*uipb.CSFriendBatch)
	agree := csFriendBatch.GetAgree()
	_, err := friendBatch(tpl, agree)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"agree":    agree,
				"error":    err,
			}).Error("friend:处理好友批量决策,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友批量决策,完成")
	return nil

}

//处理好友批量决策
func friendBatch(pl player.Player, agree bool) (flag bool, err error) {
	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	inviteMap := friendManager.GetFriendInviteMap()
	inviteLen := len(inviteMap)
	if inviteLen == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"agree":    agree,
			}).Warn("friend:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//不同意
	if !agree {
		for friendId, _ := range inviteMap {
			friendManager.RemoveFriendInvite(friendId)
			gameevent.Emit(friendeventtypes.EventTypeFriendBatchRefuse, pl, friendId)
		}
	} else {
		//判断好友数量
		numFriends := friendManager.NumOfFriend()
		maxFriends := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFriendLimit)
		if numFriends+inviteLen >= int(maxFriends) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"agree":    agree,
				}).Warn("friend:处理好友批量决策,全部同意好友数会达上限")
			playerlogic.SendSystemMessage(pl, lang.FriendBatchAgreeFull)
			return
		}

		for friendId, _ := range inviteMap {
			flag = friendManager.IsBlack(friendId)
			friendManager.RemoveFriendInvite(friendId)
			if flag {
				gameevent.Emit(friendeventtypes.EventTypeFriendBatchRefuse, pl, friendId)
			} else {
				friend.GetFriendService().AddFriend(pl, friendId)
			}
		}
	}
	scFrinedBatch := pbutil.BuildSCFrinedBatch(agree)
	pl.SendMsg(scFrinedBatch)
	return
}
