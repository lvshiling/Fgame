package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/realm/pbutil"
	"fgame/fgame/game/realm/realm"
	realmtypes "fgame/fgame/game/realm/types"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_REALM_PAIR_DEAL_TYPE), dispatch.HandlerFunc(handleRealmPairDeal))
}

//处理夫妻对战决策信息
func handleRealmPairDeal(s session.Session, msg interface{}) (err error) {
	log.Debug("realm:处理夫妻对战决策消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csRealmPairDeal := msg.(*uipb.CSRealmPairDeal)
	result := csRealmPairDeal.GetResult()
	err = realmDeal(tpl, realmtypes.RealmResultType(result))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"result":   result,
				"error":    err,
			}).Error("realm:处理夫妻对战决策消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("realm:处理夫妻对战决策消息完成")
	return nil
}

//处理夫妻对战决策信息逻辑
func realmDeal(pl player.Player, result realmtypes.RealmResultType) (err error) {
	if result == realmtypes.RealmResultTypeOk {
		if !playerlogic.CheckCanEnterScene(pl) {
			return
		}
		realmInvite := realm.GetRealmRankService().GetRealmInvite(pl.GetId())
		if realmInvite != nil {
			peerPlayerId := realmInvite.PlayerId
			teamDataObj := team.GetTeamService().GetTeamByPlayerId(peerPlayerId)
			if teamDataObj != nil && teamDataObj.IsMatch() {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
				}).Warn("realm:对方当前正在3v3匹配,邀请取消")
				playerlogic.SendSystemMessage(pl, lang.RealmPeerIn3v3Match)
				return
			}
		}
	}

	inviteName, codeResult := realm.GetRealmRankService().PairInviteDeal(pl, result)
	scRealmPair := pbutil.BuildSCRealmPairDeal(int32(codeResult))
	pl.SendMsg(scRealmPair)
	if codeResult == realmtypes.RealmPairCodeTypeCancle {
		scRealmPairPushCancle := pbutil.BuildSCRealmPairPushCancle(inviteName)
		pl.SendMsg(scRealmPairPushCancle)
	}
	return
}
