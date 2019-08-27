package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	crosslogic "fgame/fgame/game/cross/logic"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_ATTEND_SUCCESS_TYPE), dispatch.HandlerFunc(handleArenapvpAttendSuccess))
}

//处理跨服参加pvp成功
func handleArenapvpAttendSuccess(s session.Session, msg interface{}) (err error) {
	log.Debug("pvp:处理跨服参加pvp")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenapvpAttendSuccess(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("pvp:处理跨服参加pvp,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("pvp:处理跨服参加pvp,完成")
	return nil
}

//参加pvp
func arenapvpAttendSuccess(pl player.Player) (err error) {
	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpInfo := arenapvpManager.GetPlayerArenapvpObj()
	if !arenapvpInfo.IfBuyTicket() {
		constantTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpConstantTemplate()
		needBindGold := int64(constantTemp.RuchangUseBindgold)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		if !propertyManager.HasEnoughGold(needBindGold, true) {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"needBindGold": needBindGold,
				}).Warn("arenapvp:比武大会购买门票,元宝不足")

			crosslogic.PlayerExitCross(pl)
			return
		}

		useReason := commonlog.GoldLogReasonArenapvpBuyTicket
		flag := propertyManager.CostGold(needBindGold, true, useReason, useReason.String())
		if !flag {
			crosslogic.PlayerExitCross(pl)
			return
		}
		propertylogic.SnapChangedProperty(pl)

		flag = arenapvpManager.BuyArenapvpTicket()
		if !flag {
			crosslogic.PlayerExitCross(pl)
			return
		}
	}

	arenapvpManager.EnterArenapvp()
	return
}
