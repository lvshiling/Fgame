package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/friend/friend"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_DELETE_TYPE), dispatch.HandlerFunc(handleFriendDelete))
}

//处理好友删除
func handleFriendDelete(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友删除")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFriendDelete := msg.(*uipb.CSFriendDelete)
	friendId := csFriendDelete.GetFriendId()
	err := friendDelete(tpl, friendId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理好友删除,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友删除,完成")
	return nil

}

//处理好友删除
func friendDelete(pl player.Player, friendId int64) (err error) {
	//判断是否是夫妻
	marryManager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := marryManager.GetMarryInfo()
	spouseId := marryInfo.SpouseId
	if friendId == spouseId {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:对方是您配偶,无法删除")
		playerlogic.SendSystemMessage(pl, lang.FriendIsSpouseNoDelete)
		return
	}

	err = friend.GetFriendService().DeleteFriend(pl, friendId)
	if err != nil {
		return
	}
	return
}
