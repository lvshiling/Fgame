package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_BLACK_TYPE), dispatch.HandlerFunc(handleFriendBlack))
}

//处理黑名单
func handleFriendBlack(s session.Session, msg interface{}) error {
	log.Debug("friend:处理好友黑名单")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFriendBlack := msg.(*uipb.CSFriendBlack)
	friendId := csFriendBlack.GetFriendId()
	err := friendBlack(tpl, friendId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理好友黑名单,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友黑名单,完成")
	return nil

}

//处理好友添加
func friendBlack(pl player.Player, friendId int64) (err error) {
	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	//判断是否可以加入黑名单
	flag := friendManager.ShouldAddBlack(friendId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:对方已经被你拉黑过了")
		playerlogic.SendSystemMessage(pl, lang.FriendIsBlack)
		return
	}
	numBlacks := friendManager.NumOfBlack()
	maxBlacks := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBlackLimit)

	if numBlacks >= int(maxBlacks) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友黑名单,黑名单已经上限")
		playerlogic.SendSystemMessage(pl, lang.FriednBlackAlreadyFull)
		return
	}

	playerInfo, err := player.GetPlayerService().GetPlayerInfo(friendId)
	if err != nil {
		return err
	}
	if playerInfo == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理拉黑,用户不存在")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoExist)
		return
	}

	friendManager.AddBlack(friendId)
	scFriendBlack := pbutil.BuildSCFriendBlack(friendId)
	pl.SendMsg(scFriendBlack)
	return
}
