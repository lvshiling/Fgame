package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	playerfriend "fgame/fgame/game/friend/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/pbutil"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_PLAYER_BASIC_INFO_BATCH_GET_TYPE), dispatch.HandlerFunc(handlePlayerBasicInfoBatchGet))
}

//处理玩家信息批量获取
func handlePlayerBasicInfoBatchGet(s session.Session, msg interface{}) (err error) {
	log.Debug("player:处理玩家信息批量获取")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(player.Player)
	csPlayerBasicInfoBatchGet := msg.(*uipb.CSPlayerBasicInfoBatchGet)
	getPlayerIdList := csPlayerBasicInfoBatchGet.GetPlayerIdList()
	err = playerBasicInfoBatchGet(pl, getPlayerIdList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"getPlayerIdList": getPlayerIdList,
				"error":           err,
			}).Error("player:处理玩家信息批量获取,创建失败")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":        pl.GetId(),
			"getPlayerIdList": getPlayerIdList,
		}).Debug("player:处理玩家信息批量获取,创建成功")
	return
}

func playerBasicInfoBatchGet(pl player.Player, getPlayerIdList []int64) (err error) {
	infoList, err := player.GetPlayerService().BatchGetPlayerInfo(getPlayerIdList)
	if err != nil {
		return
	}

	friendManager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	revBlackMap := friendManager.GetRevBlacks()
	for _, info := range infoList {
		_, exist := revBlackMap[info.PlayerId]
		if !exist {
			continue
		}
		info.OnlineState = types.PlayerOnlineStateOffline
	}

	scPlayerBasicInfoBatchGet := pbutil.BuildSCPlayerBasicInfoBatchGet(infoList)
	pl.SendMsg(scPlayerBasicInfoBatchGet)
	return
}
