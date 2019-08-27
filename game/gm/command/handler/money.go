package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/feedbackfee/pbutil"
	playerfeedbackfee "fgame/fgame/game/feedbackfee/player"
	feedbackfeetypes "fgame/fgame/game/feedbackfee/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMoney, command.CommandHandlerFunc(handleMoney))
}

func handleMoney(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	moneyStr := c.Args[0]
	money, err := strconv.ParseInt(moneyStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"money": moneyStr,
				"error": err,
			}).Warn("gm:处理设置库存,money不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	err = setMoney(pl, int32(money))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"money": moneyStr,
				"error": err,
			}).Error("gm:处理设置库存,错误")

		return
	}

	return
}

func setMoney(pl player.Player, money int32) error {
	money *= feedbackfeetypes.MoneyYuan
	m := pl.GetPlayerDataManager(playertypes.PlayerFeedbackFeeDataManagerType).(*playerfeedbackfee.PlayerFeedbackFeeManager)
	m.GMSetMoney(money)
	feeInfo := m.GetFeedbackFeeInfo()
	recordObj := m.GetCurrentRecord()
	xianJinflag := center.GetCenterService().IsXianJinOpen()
	//发送通知
	scFeedbackFeeInfo := pbutil.BuildSCFeedbackFeeInfo(xianJinflag, feeInfo, recordObj)
	pl.SendMsg(scFeedbackFeeInfo)
	return nil
}
