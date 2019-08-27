package logic

import (
	gametemplate "fgame/fgame/game/template"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	xiuxianbooktemplate "fgame/fgame/game/welfare/xiuxianbook/template"
	xiuxianbooktypes "fgame/fgame/game/welfare/xiuxianbook/types"
	"fgame/fgame/game/welfare/xiuxianbook/xiuxianbook"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

// 计算等级
func CountLevel(obj *playerwelfare.PlayerOpenActivityObject) (level int32, err error) {
	pl := obj.GetPlayer()
	// groupId := obj.GetGroupId()
	info, _ := obj.GetActivityData().(*xiuxianbooktypes.XiuxianBookInfo)
	// if !ok {
	// 	err = fmt.Errorf("welfare:修仙典籍，物品对象不对")
	// 	return
	// }
	//计算等级
	countLevelHandler := xiuxianbook.GetCountLevelHandler(obj.GetActivityType(), obj.GetActivitySubType())
	level, err = countLevelHandler.CountLevel(pl)
	level = info.GetHighestLevelInHistory(level)
	// log.WithFields(
	// 	log.Fields{
	// 		"playerId": pl.GetId(),
	// 		"groupId":  groupId,
	// 		"subtype":  obj.GetActivitySubType(),
	// 		"level":    level,
	// 	}).Info("welfare:计算等级")
	return
}

// 最近一档可领取奖励
func CountRecentCanReceiveRewLevel(obj *playerwelfare.PlayerOpenActivityObject) (recentNeedlevel int32, err error) {
	//计算等级
	curLevel, err := CountLevel(obj)
	if err != nil {
		return 0, err
	}

	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		err = fmt.Errorf("welfare:修仙典籍获取最近一档可领取奖励错误，模板不存在")
		return 0, err
	}
	groupTemp := groupInterface.(*xiuxianbooktemplate.GroupTemplateXiuxianBook)
	tempList := groupTemp.GetGroupXiuxianBookList()
	beforeLevel := int32(0)
	//计算最近一档可领取奖励
	for i := 0; i < len(tempList); i++ {
		needLevel := tempList[i].Value1
		if curLevel < needLevel {
			recentNeedlevel = beforeLevel
			return recentNeedlevel, nil
		}
		beforeLevel = needLevel
	}
	return 0, nil
}

// 是否能领取奖励
func IsCanReceiceRew(obj *playerwelfare.PlayerOpenActivityObject, openTemp *gametemplate.OpenserverActivityTemplate) (flag bool, curLevel int32, needLevel int32) {
	info, _ := obj.GetActivityData().(*xiuxianbooktypes.XiuxianBookInfo)
	// if !ok {
	// 	flag = false
	// 	return
	// }
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	needLevel = openTemp.Value1
	needChargeNum := openTemp.Value2
	curLevel, err := CountLevel(obj)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"err":      err,
			}).Warn("welfare:计算等级出错")
		flag = false
		return
	}
	return info.IsCanReceiveReward(needLevel, curLevel, needChargeNum), curLevel, needLevel
}

// 可领取奖励
func GetCanRewList(obj *playerwelfare.PlayerOpenActivityObject) (level int32, maxLevel int32, err error) {
	info := obj.GetActivityData().(*xiuxianbooktypes.XiuxianBookInfo)
	curLevel, err := CountLevel(obj)
	if err != nil {
		return 0, 0, err
	}
	// groupId := obj.GetGroupId()
	// groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	// if groupInterface == nil {
	// 	err = fmt.Errorf("welfare:修仙典籍获取可领取奖励，模板不存在")
	// 	return 0, 0, err
	// }
	// groupTemp := groupInterface.(*xiuxianbooktemplate.GroupTemplateXiuxianBook)
	// canReceiveList := groupTemp.GetCanReceiveList(curLevel, info.ChargeNum)
	// level = info.GetMinCanReceiveLevel(canReceiveList)
	level = info.FirstTimeRewRecord
	return level, curLevel, nil
}
