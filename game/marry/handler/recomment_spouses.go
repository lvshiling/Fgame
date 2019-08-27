package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	playerfriend "fgame/fgame/game/friend/player"
	"fgame/fgame/game/marry/pbutil"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_RECOMMENT_SPOUSES_TYPE), dispatch.HandlerFunc(handleRecommentSpousesGet))
}

//处理配偶推荐
func handleRecommentSpousesGet(s session.Session, msg interface{}) error {
	log.Debug("marry:处理配偶推荐")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err := recommentSpousesGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("marry:处理配偶推荐,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理配偶推荐,完成")
	return nil

}

//配偶推荐
func recommentSpousesGet(pl player.Player) (err error) {
	//TODO 优化算法
	//随机3个玩家
	friendManager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	friendMap := friendManager.GetFriends()
	excludePlayers := make(map[int64]struct{})
	for _, fri := range friendManager.GetBlacks() {
		excludePlayers[fri.FriendId] = struct{}{}
	}
	scenePlayerList := make([]scene.Player, 0, 3)
	playerList := player.GetPlayerService().RecommentSpouses(pl, excludePlayers)
	for _, tempP := range playerList {
		scenePlayerList = append(scenePlayerList, tempP)
	}
	curLen := len(playerList)
	if curLen < player.RecommentSpouseNum {
		//假人
		showServerId := merge.GetMergeService().GetMergeTime() != 0
		tempList := robot.GetRobotService().CreateMarryRecommentRobot(pl, int32(player.RecommentSpouseNum-curLen), showServerId)
		for _, tempP := range tempList {
			scenePlayerList = append(scenePlayerList, tempP)
		}
	}

	friendIdList := make([]int64, 0, 8)
	for friendId, _ := range friendMap {
		friendIdList = append(friendIdList, friendId)
	}

	scMarryRecomment := pbuitl.BuildSCMarryRecomment(friendIdList, scenePlayerList)
	pl.SendMsg(scMarryRecomment)
	return
}
