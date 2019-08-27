package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	developfamoustypes "fgame/fgame/game/welfare/develop/famous/types"
	welfareventtypes "fgame/fgame/game/welfare/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_DEVELOP_FEED_TYPE), dispatch.HandlerFunc(handlerDevelopFeed))
}

//参与名人培养
func handlerDevelopFeed(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理参与名人培养请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSOpenActivityDevelopFeed)
	groupId := csMsg.GetGroupId()
	feedIndex := csMsg.GetFeedIndex()

	feedTimes := csMsg.GetFeedTimes()
	if feedTimes <= 0 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"groupId":   groupId,
				"feedTimes": feedTimes,
			}).Warn("welfare:参与名人培养请求，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = developFeed(tpl, groupId, feedIndex, feedTimes)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理参与名人培养请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理参与名人培养请求完成")

	return
}

//参与名人培养请求逻辑
func developFeed(pl player.Player, groupId int32, feedIndex, feedTimes int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDevelop
	subType := welfaretypes.OpenActivityDefaultSubTypeDefault

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	famousTemp := welfaretemplate.GetWelfareTemplateService().GetFamousTemplate(groupId)
	if famousTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:参与名人培养请求，名人模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	feedData := famousTemp.GetFameFeedInfo(int(feedIndex))
	if feedData == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:参与名人培养请求，喂养种类配置不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.RefreshActivityDataByGroupId(groupId)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*developfamoustypes.DevelopFameInfo)
	if !info.IfCanFeed(feedData.ItemId, feedTimes, feedData.FeedLimit) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"needItemId": feedData.ItemId,
				"timesMap":   info.FeedTimesMap,
				"limit":      feedData.FeedLimit,
				"addTimes":   feedTimes,
			}).Warn("welfare:参与名人培养请求，本物品贡献次数已达上限，无法继续贡献")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityDevelopHadFullTimes)
		return
	}

	needNum := feedTimes * feedData.ItemNum
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughItem(feedData.ItemId, needNum) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"needItemId": feedData.ItemId,
				"needNum":    needNum,
			}).Warn("welfare:参与名人培养请求，物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//奖励
	var totalItemList []*droptemplate.DropItemData
	for init := int32(1); init <= feedTimes; init++ {
		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(feedData.DropId)
		if dropData == nil {
			continue
		}
		dropData.BindType = itemtypes.ItemBindTypeUnBind
		totalItemList = append(totalItemList, dropData)
	}
	var dropItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(totalItemList) != 0 {
		dropItemList, resMap = droplogic.SeperateItemDatas(totalItemList)
	}

	if !inventoryManager.HasEnoughSlotsOfItemLevel(dropItemList) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"dropItemList": dropItemList,
			}).Warn("welfare:参与名人培养请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗物品
	useItemReason := commonlog.InventoryLogReasonOpenActivityUse
	useItemReasonText := fmt.Sprintf(useItemReason.String(), typ, subType)
	flag := inventoryManager.UseItem(feedData.ItemId, needNum, useItemReason, useItemReasonText)
	if !flag {
		panic("welfare:名人普培养消耗物品应该成功")
	}

	//增加掉落
	if len(resMap) > 0 {
		goldReason := commonlog.GoldLogReasonOpenActivityRew
		silverReason := commonlog.SilverLogReasonOpenActivityRew
		levelReason := commonlog.LevelLogReasonOpenActivityRew
		goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
		silverReasonText := fmt.Sprintf(silverReason.String(), typ, subType)
		levelReasonText := fmt.Sprintf(levelReason.String(), typ, subType)
		err = droplogic.AddRes(pl, resMap, goldReason, goldReasonText, silverReason, silverReasonText, levelReason, levelReasonText)
		if err != nil {
			return
		}
	}

	if len(dropItemList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
		flag = inventoryManager.BatchAddOfItemLevel(dropItemList, itemGetReason, itemGetReasonText)
		if !flag {
			panic("welfare:增加物品应该成功")
		}
	}

	//更新信息
	addNum := feedData.AddFavorableNum * feedTimes
	info.FavorableNum += addNum
	info.DayFavorableNum += addNum
	info.AddFeedTimes(feedData.ItemId, feedTimes)
	welfareManager.UpdateObj(obj)

	evedData := welfareventtypes.CreatePlayerFamousFeedEventData(groupId, info.FavorableNum, info.DayFavorableNum)
	gameevent.Emit(welfareventtypes.EventTypeDevelopFamousFeed, pl, evedData)

	//同步资源
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityDevelopFeed(groupId, info.FavorableNum, info.DayFavorableNum, info.FeedTimesMap, totalItemList)
	pl.SendMsg(scMsg)
	return
}
