package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeCharge, command.CommandHandlerFunc(handleCharge))
}

func handleCharge(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理充值")
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	numStr := c.Args[0]
	chargeId, _ := strconv.ParseInt(numStr, 10, 64)

	err = chargeDeal(pl, int32(chargeId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理充值,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理充值,完成")
	return
}

func chargeDeal(p scene.Player, chargeId int32) (err error) {
	// pl := p.(player.Player)
	// chargeTmep := chargetemplate.GetChargeTemplateService().GetChargeTemplate(chargeId)
	// if chargeTmep == nil {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"id": pl.GetId(),
	// 		}).Warn("gm:处理充值,模板不存在")

	// 	playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
	// 	return
	// }

	// chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
	// chargeManager.AddCharge(chargeId)

	// propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	// reason := commonlog.GoldLogReasonGM
	// propertyManager.AddGold(int64(chargeTmep.Gold), false, reason, reason.String())
	// propertylogic.SnapChangedProperty(pl)

	// scCharge := chargepbutil.BuildSCCharge(chargeId)
	// pl.SendMsg(scCharge)

	// welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	// isFirst := welfareManager.IsFirstCharge()
	// isReceive := welfareManager.IsReceiveFirstCharge()
	// firstChagreNotice := pbutil.BuildSCOpenActivityFirstChargeNotice(isFirst, isReceive)
	// pl.SendMsg(firstChagreNotice)
	return
}
