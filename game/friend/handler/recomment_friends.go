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
	processor.Register(codec.MessageType(uipb.MessageType_CS_RECOMMENT_FRIENDS_GET_TYPE), dispatch.HandlerFunc(handleRecommentFriendsGet))
}

//处理好友推荐
func handleRecommentFriendsGet(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友推荐")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err := recommentFriendsGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理好友推荐,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友推荐,完成")
	return nil
}

//好友推荐
func recommentFriendsGet(pl player.Player) (err error) {
	//TODO 优化算法
	//随机6个玩家
	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	excludePlayers := friendManager.GetExcludeForAddAll()
	infoList := player.GetPlayerService().RecommentPlayersExclude(excludePlayers)
	scRecommentFriendsGet := pbutil.BuildSCRecommentFriendsGet(infoList)
	pl.SendMsg(scRecommentFriendsGet)
	return
}
