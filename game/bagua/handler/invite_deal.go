package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/bagua/bagua"
	"fgame/fgame/game/bagua/pbutil"
	baguatypes "fgame/fgame/game/bagua/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BAGUA_PAIR_DEAL_TYPE), dispatch.HandlerFunc(handleBaGuaPairDeal))
}

//处理夫妻对战决策信息
func handleBaGuaPairDeal(s session.Session, msg interface{}) (err error) {
	log.Debug("bagua:处理夫妻对战决策消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csBaGuaPairDeal := msg.(*uipb.CSBaGuaPairDeal)
	result := csBaGuaPairDeal.GetResult()
	err = baGuaPairDeal(tpl, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"result":   result,
				"error":    err,
			}).Error("bagua:处理夫妻对战决策消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("bagua:处理夫妻对战决策消息完成")
	return nil
}

//处理夫妻对战决策信息逻辑
func baGuaPairDeal(pl player.Player, result bool) (err error) {
	if result {
		if !playerlogic.CheckCanEnterScene(pl) {
			return
		}
		baGuaInvite := bagua.GetBaGuaService().GetBaGuaInvite(pl.GetId())
		if baGuaInvite != nil {
			peerPlayerId := baGuaInvite.PlayerId
			teamDataObj := team.GetTeamService().GetTeamByPlayerId(peerPlayerId)
			if teamDataObj != nil && teamDataObj.IsMatch() {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
				}).Warn("bagua:对方当前正在3v3匹配,邀请取消")
				playerlogic.SendSystemMessage(pl, lang.BaGuaPeerIn3v3Match)
				return
			}
		}
	}

	inviteName, codeResult := bagua.GetBaGuaService().PairInviteDeal(pl, result)
	scBaGuaPairDeal := pbutil.BuildSCBaGuaPairDeal(int32(codeResult))
	pl.SendMsg(scBaGuaPairDeal)
	if codeResult == baguatypes.BaGuaPairCodeTypeCancle {
		scBaGuaPairPushCancle := pbutil.BuildSCBaGuaPairPushCancle(inviteName)
		pl.SendMsg(scBaGuaPairPushCancle)
	}
	return
}
