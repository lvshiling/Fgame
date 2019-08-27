package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_REMOVE_BLACK_TYPE), dispatch.HandlerFunc(handleFriendRemoveBlack))
}

//处理移除黑名单
func handleFriendRemoveBlack(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友移除黑名单")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	CSFriendRemoveBlack := msg.(*uipb.CSFriendRemoveBlack)
	friendId := CSFriendRemoveBlack.GetFriendId()
	err := friendRemoveBlack(tpl, friendId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理好友移除黑名单,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友移除黑名单,完成")
	return nil

}

//处理好友移除黑名单
func friendRemoveBlack(pl player.Player, friendId int64) (err error) {
	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	isFriend := friendManager.IsFriend(friendId)
	//判断是否可以移除黑名单
	flag := friendManager.ShouldRemoveBlack(friendId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:对方不在黑名单内")
		playerlogic.SendSystemMessage(pl, lang.FriendNotInBlack)
		return
	}

	friendManager.RemoveBlack(friendId)
	scFriendRemoveBlack := pbutil.BuildSCFriendRemoveBlack(friendId, isFriend)
	pl.SendMsg(scFriendRemoveBlack)
	return
}
