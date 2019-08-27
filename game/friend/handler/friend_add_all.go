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
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_ADD_ALL_TYPE), dispatch.HandlerFunc(handleFriendAddAll))
}

//处理好友一键添加
func handleFriendAddAll(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友一键添加")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err := friendAddAll(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理好友一键添加,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友一键添加,完成")
	return nil

}

//处理好友一键添加
func friendAddAll(pl player.Player) (err error) {
	pl.GetLastKillTime()
	manager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	if !manager.IfCanAddAll() {
		leftTimeMs := manager.AddAllLeftTime()
		leftTime := timeutils.MillisecondToSecondCeil(leftTimeMs)
		leftTimeStr := fmt.Sprintf("%d", leftTime)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("friend:一键添加,cd")
		playerlogic.SendSystemMessage(pl, lang.MiscExitKaSiCd, leftTimeStr)
		return
	}
	//判断好友数量
	numFriends := manager.NumOfFriend()
	maxFriends := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFriendLimit)
	if numFriends >= int(maxFriends) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("friend:处理好友添加邀请,好友已经达到上限")
		playerlogic.SendSystemMessage(pl, lang.FriendAlreadyFull)
		return
	}

	excludePlayers := manager.GetExcludeForAddAll()
	infoList := player.GetPlayerService().RecommentPlayersExclude(excludePlayers)
	if len(infoList) != 0 {
		gameevent.Emit(friendeventtypes.EventTypeFriendAddAll, infoList, pl)
	}

	cdTime := manager.InviteAddAllCdTime()
	scFriendAddAll := pbutil.BuildSCFriendAddAll(cdTime)
	pl.SendMsg(scFriendAddAll)
	return
}
