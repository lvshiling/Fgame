package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/feedbackfee/feedbackfee"
	"fgame/fgame/game/feedbackfee/pbutil"
	playerfeedbackfee "fgame/fgame/game/feedbackfee/player"
	feedbackfeetemplate "fgame/fgame/game/feedbackfee/template"
	feedbackfeetypes "fgame/fgame/game/feedbackfee/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func HandlePlayerExchange(pl player.Player, money int32) {
	//TODO 验证档次
	moneyTemplate := feedbackfeetemplate.GetFeedbackfeeTemplateService().GetExchangeTemplate(money)
	if moneyTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("feedbackfee:处理获取兑换请求,正在兑换中")
		playerlogic.SendSystemMessage(pl, lang.FeebackMoneyMoneyWrong)
		return
	}

	money *= feedbackfeetypes.MoneyYuan
	playerFeedbackFeeDataManager := pl.GetPlayerDataManager(playertypes.PlayerFeedbackFeeDataManagerType).(*playerfeedbackfee.PlayerFeedbackFeeManager)
	playerPropertyDataManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//判断是否正在兑换
	currentRecordList := playerFeedbackFeeDataManager.GetCurrentRecordList()
	if len(currentRecordList) > 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("feedbackfee:处理获取兑换请求,正在兑换中")
		playerlogic.SendSystemMessage(pl, lang.FeebackExchange)
		return
	}
	//判断是否可以兑换
	if !playerFeedbackFeeDataManager.IfCanExchange(money) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"money":    money,
			}).Warn("feedbackfee:处理获取兑换请求,余额不足或者超过限制")
		playerlogic.SendSystemMessage(pl, lang.FeebackMoneyNoEnoughOrLimit)
		return
	}
	xianJinflag := center.GetCenterService().IsXianJinOpen()
	var record *playerfeedbackfee.PlayerFeedbackRecordObject
	if xianJinflag {
		record = playerFeedbackFeeDataManager.Exchange(money)
		if record == nil {
			panic(fmt.Errorf("feedbackfee:应该兑换成功,玩家[%d],钱[%d]", pl.GetId(), money))
		}
	} else {
		record = playerFeedbackFeeDataManager.ExchangeGold(money, int64(moneyTemplate.Gold))
		if record == nil {
			panic(fmt.Errorf("feedbackfee:应该兑换成功,玩家[%d],钱[%d]", pl.GetId(), money))
		}
		reason := commonlog.GoldLogReasonXianJinExchange
		reasonText := fmt.Sprintf(reason.String(), record.GetId(), moneyTemplate.Money)
		playerPropertyDataManager.AddGold(int64(moneyTemplate.Gold), false, reason, reasonText)
		propertylogic.SnapChangedProperty(pl)
	}

	//获取兑换码
	feedbackfee.GetFeedbackFeeService().Exchange(pl.GetId(), record.GetId(), record.GetMoney(), record.GetExpiredTime())
	//下发记录信息
	feeInfo := playerFeedbackFeeDataManager.GetFeedbackFeeInfo()
	recordObj := playerFeedbackFeeDataManager.GetCurrentRecord()

	//发送通知
	scFeedbackFeeInfo := pbutil.BuildSCFeedbackFeeInfo(xianJinflag, feeInfo, recordObj)
	pl.SendMsg(scFeedbackFeeInfo)
}

func PlayerCodeExpire(p player.Player, obj *feedbackfee.FeedbackExchangeObject) {
	m := p.GetPlayerDataManager(playertypes.PlayerFeedbackFeeDataManagerType).(*playerfeedbackfee.PlayerFeedbackFeeManager)
	playerObj := m.Expire(obj.GetExchangeId())
	feedbackfee.GetFeedbackFeeService().ExchangeNotify(obj.GetId())

	//通知用户
	if playerObj == nil {
		return
	}
	feeInfo := m.GetFeedbackFeeInfo()
	recordObj := m.GetCurrentRecord()
	xianJinflag := center.GetCenterService().IsXianJinOpen()
	//发送通知
	scFeedbackFeeInfo := pbutil.BuildSCFeedbackFeeInfo(xianJinflag, feeInfo, recordObj)
	p.SendMsg(scFeedbackFeeInfo)
}

func PlayerCodeGenerate(p player.Player, obj *feedbackfee.FeedbackExchangeObject) {
	m := p.GetPlayerDataManager(playertypes.PlayerFeedbackFeeDataManagerType).(*playerfeedbackfee.PlayerFeedbackFeeManager)
	playerObj := m.CodeGenerate(obj.GetExchangeId(), obj.GetCode())
	feedbackfee.GetFeedbackFeeService().FillCode(obj.GetId())

	//通知用户
	if playerObj == nil {
		return
	}
	//发送通知
	feeInfo := m.GetFeedbackFeeInfo()
	recordObj := m.GetCurrentRecord()
	xianJinflag := center.GetCenterService().IsXianJinOpen()
	//发送通知
	scFeedbackFeeInfo := pbutil.BuildSCFeedbackFeeInfo(xianJinflag, feeInfo, recordObj)
	p.SendMsg(scFeedbackFeeInfo)
}

func PlayerCodeFinish(p player.Player, obj *feedbackfee.FeedbackExchangeObject) {
	m := p.GetPlayerDataManager(playertypes.PlayerFeedbackFeeDataManagerType).(*playerfeedbackfee.PlayerFeedbackFeeManager)
	playerObj := m.CodeFinish(obj.GetExchangeId())
	feedbackfee.GetFeedbackFeeService().ExchangeNotify(obj.GetId())

	//通知用户
	if playerObj == nil {
		return
	}
	//发送通知
	feeInfo := m.GetFeedbackFeeInfo()
	recordObj := m.GetCurrentRecord()
	xianJinflag := center.GetCenterService().IsXianJinOpen()
	//发送通知
	scFeedbackFeeInfo := pbutil.BuildSCFeedbackFeeInfo(xianJinflag, feeInfo, recordObj)
	p.SendMsg(scFeedbackFeeInfo)
}

//添加逆付费
func AddMoney(pl player.Player, money int32, reason commonlog.FeedbackLogReason, reasonText string) {
	m := pl.GetPlayerDataManager(playertypes.PlayerFeedbackFeeDataManagerType).(*playerfeedbackfee.PlayerFeedbackFeeManager)
	m.AddMoney(money, reason, reasonText)

	feeInfo := m.GetFeedbackFeeInfo()
	recordObj := m.GetCurrentRecord()
	xianJinflag := center.GetCenterService().IsXianJinOpen()
	//发送通知
	scFeedbackFeeInfo := pbutil.BuildSCFeedbackFeeInfo(xianJinflag, feeInfo, recordObj)
	pl.SendMsg(scFeedbackFeeInfo)
}
