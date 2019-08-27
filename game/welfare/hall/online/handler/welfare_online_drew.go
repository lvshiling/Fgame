package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	hallonlinetemplate "fgame/fgame/game/welfare/hall/online/template"
	hallonlinetypes "fgame/fgame/game/welfare/hall/online/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_WELFARE_ONLINE_DREW_TYPE), dispatch.HandlerFunc(handlerWelfareOnlineDrew))
}

//处理在线抽奖
func handlerWelfareOnlineDrew(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理在线抽奖请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityWelfareOnlineDrew := msg.(*uipb.CSOpenActivityWelfareOnlineDrew)
	typ := csOpenActivityWelfareOnlineDrew.GetTyp()
	groupId := csOpenActivityWelfareOnlineDrew.GetGroupId()

	drewType := hallonlinetypes.OnlineDrewType(typ)
	if !drewType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("welfare:在线抽奖请求，类型错误")

		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = welfareOnlineDrew(tpl, drewType, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理在线抽奖请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理在线抽奖请求完成")

	return
}

//在线抽奖请求逻辑
func welfareOnlineDrew(pl player.Player, drewType hallonlinetypes.OnlineDrewType, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeWelfare
	subType := welfaretypes.OpenActivityWelfareSubTypeOnline
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	err = welfareManager.RefreshActivityDataByGroupId(groupId)
	if err != nil {
		return
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	//元宝抽奖条件
	if drewType == hallonlinetypes.OnlineDrewTypeGold {
		if !welfareManager.IsFirstCharge() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"drewType": drewType,
				}).Warn("welfare:在线抽奖请求，需要首充")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityNotFirstChargeUser)
			return
		}
	}

	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*hallonlinetypes.WelfareOnlineInfo)

	//次数判断
	onlineTime := pl.GetTodayOnlineTime()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:在线抽奖请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	groupTemp := groupInterface.(*hallonlinetemplate.GroupTemplateWelfareOnline)
	curMaxTimes := groupTemp.GetOpenActivityWelfareOnlineDrewTimes(onlineTime)
	if info.DrawTimes >= curMaxTimes {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"onlineTime": onlineTime,
				"drewType":   drewType,
			}).Warn("welfare:在线抽奖请求，次数不足")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotEnoughDrewTimes)
		return
	}

	dropId := int32(0)
	if drewType == hallonlinetypes.OnlineDrewTypeGold {
		dropId = constant.GetConstantService().GetConstant(constanttypes.ConstantTypeOnlinDrewGold)
	} else {
		dropId = constant.GetConstantService().GetConstant(constanttypes.ConstantTypeOnlinDrewSilver)
	}
	dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
	dropData.BindType = itemtypes.ItemBindTypeUnBind
	if dropData != nil {
		goldReason := commonlog.GoldLogReasonOpenActivityRew
		silverReason := commonlog.SilverLogReasonOpenActivityRew
		itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
		levelReason := commonlog.LevelLogReasonOpenActivityRew
		goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
		silverReasonText := fmt.Sprintf(silverReason.String(), typ, subType)
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
		levelReasonText := fmt.Sprintf(levelReason.String(), typ, subType)
		flag, err := droplogic.AddItem(pl, dropData, goldReason, goldReasonText, silverReason, silverReasonText, itemGetReason, itemGetReasonText, levelReason, levelReasonText)
		if err != nil {
			return err
		}
		if !flag {
			panic("welfare:在线抽奖应该成功")
		}
	}

	//更新
	info.DrawTimes += 1
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	itemId := dropData.GetItemId()
	num := dropData.GetNum()
	scMsg := pbutil.BuildSCOpenActivityWelfareOnlineDrew(itemId, num)
	pl.SendMsg(scMsg)
	return
}
