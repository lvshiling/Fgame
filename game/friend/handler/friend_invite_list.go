package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/friend/pbutil"
	playerfriend "fgame/fgame/game/friend/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_INVITE_LIST_TYPE), dispatch.HandlerFunc(handleFriendInviteList))
}

//处理好友邀请列表邀请
func handleFriendInviteList(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友邀请列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err := friendInviteList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理好友邀请列表,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友邀请列表,完成")
	return nil
}

//处理好友邀请列表
func friendInviteList(pl player.Player) (err error) {
	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	inviteMap := friendManager.GetFriendInviteMap()
	scFriendInviteList := pbutil.BuildSCFriendInviteList(inviteMap)
	pl.SendMsg(scFriendInviteList)
	return
}
