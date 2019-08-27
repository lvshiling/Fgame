package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/activity/activity"
	"fgame/fgame/game/activity/pbutil"
	playeractivity "fgame/fgame/game/activity/player"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	team "fgame/fgame/game/team/team"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//活动进入请求逻辑
func HandleActiveAttend(pl player.Player, activeType activitytypes.ActivityType, args string) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activeType)
	if activityTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"activeType": activeType,
			}).Warn("activity:进入活动请求，活动不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	activityManager := pl.GetPlayerDataManager(types.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	activityManager.RefreshData()

	//每天限制参与次数
	if flag := activityManager.IsHaveTimes(activeType); !flag {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"activeType": activeType,
			}).Warn("activity:进入活动请求，活动参与次数不足")
		playerlogic.SendSystemMessage(pl, lang.ActivityNoTimes)
		return
	}

	//是否满足开启条件
	if !pl.IsFuncOpen(activityTemplate.GetFuncOpenType()) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"activeType": activeType,
			}).Warn("activity:进入活动请求，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	//判断活动是否在活动时间内
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	flag, err := activityTemplate.IsAtActivityTime(now, openTime, mergeTime)
	if err != nil {
		return
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"activeType": activeType,
			}).Warn("activity:进入活动请求，活动未开启")
		playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
		return
	}
	//是否满足等级条件
	if !activityTemplate.IsReacheLevel(pl.GetLevel()) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"activeType": activeType.String(),
			}).Warn("activity:进入活动请求，不满足等级条件")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//TODO 统一处理
	//若有队伍，则直接从队伍离队进入
	teamObj := team.GetTeamService().GetTeamByPlayerId(pl.GetId())
	if teamObj != nil {
		team.GetTeamService().LeaveTeam(pl)
	}

	//获取处理器
	h := activity.GetActivityHandler(activityTemplate.GetActivityType())
	if h == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"activeType": activeType.String(),
			}).Warn("activity:进入活动请求，活动处理器未找到")
		return
	}
	//获取结束时间
	timeTemplate, _ := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
	endTime, _ := timeTemplate.GetEndTime(now)
	//进入活动
	pl.EnterActivity(activeType, endTime)

	flag, err = h.Attend(pl, activityTemplate, args)
	if err != nil {
		return
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"activeType": activeType.String(),
			}).Warn("activity:进入活动请求，活动不存在")
		return
	}

	//更新玩家挑战记录
	activityManager.AttendActivity(activeType)

	scActivityAttend := pbutil.BuildSCActivityAttend(activeType, flag)
	pl.SendMsg(scActivityAttend)
	return
}

//是否活动定时奖励
func IsAddTickRew(pl player.Player, acType activitytypes.ActivityType, firstTime, rewTime int64) bool {
	activityManager := pl.GetPlayerDataManager(playertypes.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	now := global.GetGame().GetTimeService().Now()

	//判断时间间隔
	if activityManager.GetPreRewTime(acType) == 0 {
		enterTime := activityManager.GetEnterTime(acType)
		diffFirst := now - enterTime
		if diffFirst > firstTime {
			return true
		}
	} else {
		preRewTime := activityManager.GetPreRewTime(acType)
		diffRew := now - preRewTime
		isRew := diffRew > rewTime
		if isRew {
			return true
		}
	}
	return false
}

//添加活动定时奖励
func AddActivityTickRew(pl player.Player, acType activitytypes.ActivityType, gold, bindGold, exp, expPoint, silver int32, itemMap map[int32]int32) {
	//添加奖励
	reasonGold := commonlog.GoldLogReasonActivityTickRew
	reasonGoldText := fmt.Sprintf(reasonGold.String(), acType)
	reasonSilver := commonlog.SilverLogReasonActivityTickRew
	reasonSilverText := fmt.Sprintf(reasonSilver.String(), acType)
	reasonLevel := commonlog.LevelLogReasonActivityTickRew
	reasonLevelText := fmt.Sprintf(reasonLevel.String(), acType)
	totalRewData := propertytypes.CreateRewData(exp, expPoint, silver, gold, bindGold)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSilverText, reasonLevel, reasonLevelText)
	if !flag {
		panic("activity:activity ranking rewards add RewData should be ok")
	}

	//添加物品
	if len(itemMap) > 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		flag := inventoryManager.HasEnoughSlots(itemMap)
		if !flag {
			mailName := lang.GetLangService().ReadLang(lang.ActivityTickRewEmailTitle)
			mailContent := lang.GetLangService().ReadLang(lang.ActivityTickRewEmailContent)
			emaillogic.AddEmail(pl, mailName, mailContent, itemMap)
		} else {
			addItemReason := commonlog.InventoryLogReasonActivityTickRew
			flag := inventoryManager.BatchAdd(itemMap, addItemReason, addItemReason.String())
			if !flag {
				panic("activity:activity tick rewards add item should be ok")
			}
		}
		inventorylogic.SnapInventoryChanged(pl)
	}
	propertylogic.SnapChangedProperty(pl)

	activityManager := pl.GetPlayerDataManager(playertypes.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	activityManager.UpdateLastRewTime(acType)

	return
}

//是否活动定时积分奖励
func IsAddTickPointRew(pl player.Player, acType activitytypes.ActivityType, firstTime, rewTime int64) bool {
	activityManager := pl.GetPlayerDataManager(playertypes.PlayerActivityDataManagerType).(*playeractivity.PlayerActivityDataManager)
	now := global.GetGame().GetTimeService().Now()

	//判断时间间隔
	preRewTime := activityManager.GetPreRewPointTime(acType)
	if preRewTime == 0 {
		enterTime := activityManager.GetEnterTime(acType)
		diffFirst := now - enterTime
		if diffFirst > firstTime {
			return true
		}
	} else {
		diffRew := now - preRewTime
		isRew := diffRew > rewTime
		if isRew {
			return true
		}
	}
	return false
}

func IfActivityTime(activityType activitytypes.ActivityType) (flag bool) {
	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activityType)
	if activityTemplate == nil {
		return
	}

	//判断活动是否在活动时间内
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	flag, _ = activityTemplate.IsAtActivityTime(now, openTime, mergeTime)
	return
}
