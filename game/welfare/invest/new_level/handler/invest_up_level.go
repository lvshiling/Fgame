package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	investnewleveltemplate "fgame/fgame/game/welfare/invest/new_level/template"
	investnewleveltypes "fgame/fgame/game/welfare/invest/new_level/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_INVEST_UP_LEVEL_TYPE), dispatch.HandlerFunc(handleUpLevelInvest))
}

func handleUpLevelInvest(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理升级新等级投资计划请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityInvestUpLevel)
	typ := csMsg.GetTyp()
	groupId := csMsg.GetGroupId()

	investNewLevelType := investnewleveltypes.InvestNewLevelType(typ)
	if !investNewLevelType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("welfare:处理升级新等级投资计划请求，类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = upLevelInvest(tpl, investNewLevelType, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理升级新等级投资计划请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理升级新等级投资计划请求完成")

	return
}

func upLevelInvest(pl player.Player, investLevelType investnewleveltypes.InvestNewLevelType, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeInvest
	subType := welfaretypes.OpenActivityInvestSubTypeNewLevel

	//检查活动
	flag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !flag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:升级新等级投资计划请求，不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"investLevelType": investLevelType,
				"groupId":         groupId,
			}).Warn("welfare:升级投资计划请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	activityObj := welfareManager.GetOpenActivity(groupId)
	if activityObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:升级新等级投资计划请求，活动对象不存在")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityInvestLevelObjectNotExist)
		return
	}

	info := activityObj.GetActivityData().(*investnewleveltypes.InvestNewLevelInfo)
	//是否能够升级
	if !info.IsCanUpLevel(investLevelType) {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"groupId":         groupId,
				"investLevelType": investLevelType,
			}).Warn("welfare:购买投资计划请求，已购买")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityInvestLevelNotCanUpLev)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	groupTemp := groupInterface.(*investnewleveltemplate.GroupTemplateInvestNewLevel)

	//计算所需元宝数量
	needGoldNum := groupTemp.GetInvestLevelNeedGold(investLevelType)
	lastInvestLevelType, _ := info.GetInvestNewLevelType()
	lastNeedGoldNum := groupTemp.GetInvestLevelNeedGold(lastInvestLevelType)
	curNeedGoldNum := needGoldNum - lastNeedGoldNum

	//判断元宝是否足够
	if !propertyManager.HasEnoughGold(int64(curNeedGoldNum), false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": curNeedGoldNum,
			}).Warn("welfare:购买投资计划请求，当前元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗日志
	goldReason, ok := investLevelType.ConvertToInvesetNewLevelUpgradeCostLogType()
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"investLevelType": investLevelType,
			}).Warn("welfare:购买投资计划请求，升级日志类型应该存在")
		return
	}
	goldReasonText := goldReason.String()
	flag = propertyManager.CostGold(int64(curNeedGoldNum), false, goldReason, goldReasonText)
	if !flag {
		panic(fmt.Errorf("welfare: 升级等级投资计划消耗元宝应该成功"))
	}

	//升级成功计算需要补发的元宝
	var emailGoldNum int32
	var emailBindGoldNum int32
	rewRecordList, ok := info.InvestBuyInfoMap[lastInvestLevelType]
	if !ok {
		return
	}
	for _, level := range rewRecordList {
		openTemp := groupTemp.GetInvestLevelRewTempByArg(investLevelType, level)
		lastOpenTemp := groupTemp.GetInvestLevelRewTempByArg(lastInvestLevelType, level)
		emailGoldNum += (openTemp.RewGold - lastOpenTemp.RewGold)
		emailBindGoldNum += (openTemp.RewGoldBind - lastOpenTemp.RewGoldBind)
	}

	goldMap := make(map[int32]int32)
	if emailGoldNum != 0 {
		goldMap[int32(constanttypes.GoldItem)] = emailGoldNum
	}
	if emailBindGoldNum != 0 {
		goldMap[int32(constanttypes.BindGoldItem)] = emailBindGoldNum
	}

	// 补发元宝通过邮件
	if len(goldMap) != 0 {
		title := groupTemp.GetActivityName()
		typString, _ := investLevelType.ConvertToInvesetNewLevelEmailType()
		lastTypString, _ := lastInvestLevelType.ConvertToInvesetNewLevelEmailType()
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailUpLevelInvest), lastTypString, typString)
		emaillogic.AddEmail(pl, title, content, goldMap)
	}

	info.UpdateInvestLevelType(investLevelType)
	welfareManager.UpdateObj(activityObj)

	//推送属性变化
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCOpenActivityInvestUpLevel()
	pl.SendMsg(scMsg)

	return
}
