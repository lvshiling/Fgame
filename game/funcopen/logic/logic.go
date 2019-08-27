package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/additionsys/additionsys"
	"fgame/fgame/game/common/common"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/funcopen/funcopen"
	"fgame/fgame/game/funcopen/pbutil"
	playerfuncopen "fgame/fgame/game/funcopen/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerquest "fgame/fgame/game/quest/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

const (
	funcOpenResTypeActiveSucceed int32 = iota //激活成功
	funcOpenResTypeActiveFailed               //激活失败
)

//处理功能开启手动激活
func HandleManualActive(pl player.Player, moduleId int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeMingRiKaiQi) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": moduleId,
			}).Warn("funcopen:处理功能开启手动激活，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	moduleType := funcopentypes.FuncOpenType(moduleId)
	if !moduleType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"moduleId": moduleId,
				"error":    err,
			}).Error("funcopen:处理功能开启手动激活,错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	funcOpenTemplate := funcopen.GetFuncOpenService().GetFuncOpenTemplate(moduleType)
	if funcOpenTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"moduleId": moduleId,
				"error":    err,
			}).Error("funcopen:处理功能开启手动激活,错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if funcOpenTemplate.GetFuncOpenCheckType() != funcopentypes.FuncOpenCheckTypeManual {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"moduleId": moduleId,
				"error":    err,
			}).Error("funcopen:处理功能开启,不是手动激活功能")
		return
	}

	activeType, err := CheckFuncOpenByType(pl, moduleType)
	if err != nil {
		return
	}

	//激活结果前端表现用的
	activeResult := funcOpenResTypeActiveSucceed
	if !activeType.Valid() {
		activeResult = funcOpenResTypeActiveFailed
	}

	updateList := make([]funcopentypes.FuncOpenType, 0, 1)
	updateList = append(updateList, activeType)
	scMsg := pbutil.BuildSCFuncOpenManualActive(updateList, activeResult)
	pl.SendMsg(scMsg)
	return
}

func CheckFuncOpen(pl player.Player) (activeFuncOpens []funcopentypes.FuncOpenType, err error) {

	//TODO 优化
Loop:
	for _, funcOpenTemplate := range funcopen.GetFuncOpenService().GetAll() {
		//是否自动开启的
		if funcOpenTemplate.GetFuncOpenCheckType() != funcopentypes.FuncOpenCheckTypeAuto {
			continue
		}

		activeType, err := CheckFuncOpenByType(pl, funcOpenTemplate.GetFuncOpenType())
		if err != nil {
			panic(fmt.Errorf("funcopen:检测功能应该成功"))
		}

		if !activeType.Valid() {
			continue
		}

		activeFuncOpens = append(activeFuncOpens, activeType)
		goto Loop
	}

	return
}

//功能开启检测（单个）
func CheckFuncOpenByType(pl player.Player, moduleType funcopentypes.FuncOpenType) (activeType funcopentypes.FuncOpenType, err error) {

	funcOpenManager := pl.GetPlayerDataManager(playertypes.PlayerFuncOpenDataManagerType).(*playerfuncopen.PlayerFuncOpenDataManager)
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	funcOpenTemplate := funcopen.GetFuncOpenService().GetFuncOpenTemplate(moduleType)
	if funcOpenTemplate == nil {
		return
	}
	if funcOpenManager.IsOpen(funcOpenTemplate.GetFuncOpenType()) {
		return
	}
	//没有条件
	noCondition := true
	oneCondition := false
	//有开启任务
	if funcOpenTemplate.OpenedQuestId != 0 {
		if questManager.IsCommit(funcOpenTemplate.OpenedQuestId) {
			oneCondition = true
		}
		noCondition = false
	}
	if funcOpenTemplate.OpenedLevel != 0 || funcOpenTemplate.OpenedZhuanShu != 0 {
		//检查等级
		if funcOpenTemplate.OpenedLevel <= pl.GetLevel() && funcOpenTemplate.OpenedZhuanShu <= pl.GetZhuanSheng() {
			oneCondition = true
		}
		noCondition = false
	}
	//检查物品
	if len(funcOpenTemplate.GetOpenedItems()) != 0 {
		if inventoryManager.HasEnoughItems(funcOpenTemplate.GetOpenedItems()) {
			oneCondition = true
		}
		noCondition = false
	}

	//自动开启时间
	if funcOpenTemplate.OpenedTime != 0 {
		if pl.GetTotalOnlineTime() > funcOpenTemplate.OpenedTime {
			oneCondition = true
		}
		noCondition = false
	}

	//不满足一个条件
	if !noCondition && !oneCondition {
		return
	}

	//建号天数限制条件
	if funcOpenTemplate.JianHaoDay > 0 {
		jianhaoTime := pl.GetCreateTime()
		now := global.GetGame().GetTimeService().Now()
		diffDay, _ := timeutils.DiffDay(now, jianhaoTime)
		if diffDay < funcOpenTemplate.JianHaoDay {
			return
		}
	}

	//系统阶别条件
	needSysType := funcOpenTemplate.GetNeedSysType()
	if needSysType.Valid() {
		if additionsys.GetSystemAdvancedNum(pl, needSysType) < funcOpenTemplate.NeedSysNum {
			return
		}
	}

	//充值条件
	if funcOpenTemplate.OpenedShouchong > 0 {
		if pl.GetChargeGoldNum() < int64(funcOpenTemplate.OpenedShouchong) {
			return
		}
	}

	//开服时间限制
	if funcOpenTemplate.KaifuTime > 0 {
		serverTime := global.GetGame().GetServerTime()
		openDayTime, _ := timeutils.BeginOfDayOfTime(timeutils.MillisecondToTime(serverTime))
		now := global.GetGame().GetTimeService().Now()
		limitKaiFuTime := openDayTime + int64(funcOpenTemplate.KaifuTime-1)*int64(common.DAY)
		if limitKaiFuTime > now {
			return
		}
	}

	allCommit := true
	//检查前置功能id
	for _, parentFuncOpen := range funcOpenTemplate.GetParentFuncOpens() {
		if !funcOpenManager.IsOpen(parentFuncOpen) {
			allCommit = false
			break
		}
	}
	if !allCommit {
		return
	}
	flag := funcOpenManager.AddOpenFunc(funcOpenTemplate.GetFuncOpenType())
	if !flag {
		panic(fmt.Errorf("funcopen:添加功能应该成功"))
	}

	//开启奖励
	if len(funcOpenTemplate.GetRewItems()) > 0 {
		if !inventoryManager.HasEnoughSlots(funcOpenTemplate.GetRewItems()) {
			// 发邮件
			title := lang.GetLangService().ReadLang(lang.EmailFuncOpenRewTitle)
			content := lang.GetLangService().ReadLang(lang.EmailFuncOpenRewContent)
			emaillogic.AddEmail(pl, title, content, funcOpenTemplate.GetRewItems())
		} else {
			itemGetReason := commonlog.InventoryLogReasonFuncopenRew
			reasonText := fmt.Sprintf(itemGetReason.String(), funcOpenTemplate.FuncId)
			flag := inventoryManager.BatchAdd(funcOpenTemplate.GetRewItems(), itemGetReason, reasonText)
			if !flag {
				panic(fmt.Errorf("funcopen:添加功能开启奖励应该成功"))
			}

			inventorylogic.SnapInventoryChanged(pl)
		}
	}
	if funcOpenTemplate.RewGold > 0 || funcOpenTemplate.RewBindgold > 0 || funcOpenTemplate.RewYinliang > 0 {
		goldReason := commonlog.GoldLogReasonFuncopenRew
		goldReasonText := fmt.Sprintf(goldReason.String(), funcOpenTemplate.FuncId)
		silverReason := commonlog.SilverLogReasonActivityTickRew
		silverReasonText := fmt.Sprintf(silverReason.String(), funcOpenTemplate.FuncId)
		flag := propertyManager.AddMoney(funcOpenTemplate.RewBindgold, funcOpenTemplate.RewGold, goldReason, goldReasonText, funcOpenTemplate.RewYinliang, silverReason, silverReasonText)
		if !flag {
			panic(fmt.Errorf("funcopen:添加功能开启奖励元宝:%d,绑元:%d,银两:%d 应该成功", funcOpenTemplate.RewGold, funcOpenTemplate.RewBindgold, funcOpenTemplate.RewYinliang))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	//开启邮件通知
	isMail := false
	if len(funcOpenTemplate.MailDes) > 0 {
		isMail = true
	}
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByFuncType(funcOpenTemplate.GetFuncOpenType())
	if len(timeTempList) > 0 {
		isMail = false
	}
	if isMail {
		emaillogic.AddEmail(pl, funcOpenTemplate.MailTitle, funcOpenTemplate.MailDes, funcOpenTemplate.GetMailRewItems())
	}

	activeType = funcOpenTemplate.GetFuncOpenType()

	return
}
