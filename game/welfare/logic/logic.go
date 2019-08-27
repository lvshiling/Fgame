package logic

import (
	"context"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/common/message"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	ranktypes "fgame/fgame/game/rank/types"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	advancedrewblesscrittemplate "fgame/fgame/game/welfare/advancedrew/bless_crit/template"
	alliancecheertypes "fgame/fgame/game/welfare/alliance/cheer/types"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	drewbaokucrittemplate "fgame/fgame/game/welfare/drew/baoku_crit/template"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfareranktemplate "fgame/fgame/game/welfare/rank/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	"fgame/fgame/pkg/mathutils"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

//奖励领取
func AddOpenActivityRewards(pl player.Player, openTemp *gametemplate.OpenserverActivityTemplate) (totalRewData *propertytypes.RewData, totalItemMap map[int32]int32, flag bool) {
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(openTemp.Group)
	if timeTemp == nil {
		return
	}
	typ := timeTemp.GetOpenType()
	subType := timeTemp.GetOpenSubType()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	newItemDataList := ConvertToItemDataWithWelfareData(openTemp.GetRewItemDataList(), openTemp.GetExpireType(), openTemp.GetExpireTime())
	//背包空间
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemDataList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:领取活动奖励请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//增加物品
	itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
	itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
	flag = inventoryManager.BatchAddOfItemLevel(newItemDataList, itemGetReason, itemGetReasonText)
	if !flag {
		panic("welfare:welfare rewards add item should be ok")
	}

	reasonGold := commonlog.GoldLogReasonOpenActivityRew
	reasonSilver := commonlog.SilverLogReasonOpenActivityRew
	reasonLevel := commonlog.LevelLogReasonOpenActivityRew
	reasonGoldText := fmt.Sprintf(reasonGold.String(), typ, subType)
	reasonSilverText := fmt.Sprintf(reasonSilver.String(), typ, subType)
	reasonLevelText := fmt.Sprintf(reasonLevel.String(), typ, subType)

	rewSilver := openTemp.RewSilver
	rewBindGold := openTemp.RewGoldBind
	rewGold := openTemp.RewGold
	rewExp := int32(0)
	rewExpPoint := int32(0)
	totalRewData = propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)
	flag = propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSilverText, reasonLevel, reasonLevelText)
	if !flag {
		panic("welfare:welfare rewards add RewData should be ok")
	}
	totalItemMap = openTemp.GetEmailRewItemMap()
	return
}

//排行榜奖励邮件
func AddRankRewards(rankList []*ranktypes.RankingInfo, groupId int32, endTime int64) {
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	for _, info := range rankList {
		playerId := info.GetPlayerId()
		playerName := info.GetPlayerName()
		ranking := info.GetRanking()
		rankNum := info.GetRankNum()

		emailContentType := welfaretypes.EmailContentTypeDefault
		var spouseEmailContentType welfaretypes.EmailContentType
		//表白榜配偶奖励
		isSpouseGetRew := false
		if timeTemp.IsRankMarryDevelop() {
			emailContentType = welfaretypes.EmailContentTypeMarryDevelop
			spouseEmailContentType = welfaretypes.EmailContentTypeMarryDevelopSpouse
			isSpouseGetRew = true
		}

		//魅力榜配偶奖励
		if timeTemp.IsRankCharm() {
			emailContentType = welfaretypes.EmailContentTypeCharm
			spouseEmailContentType = welfaretypes.EmailContentTypeCharmSpouse
			isSpouseGetRew = true
		}

		//次数
		// if timeTemp.IsRankNunber() {
		// 	emailContentType = welfaretypes.EmailContentTypeBoatRaceForce
		// }

		if isSpouseGetRew {
			spouseId := marry.GetMarryService().GetMarrySpouseId(playerId)
			if spouseId != 0 {
				plSpouse := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
				if plSpouse != nil {
					ctx := scene.WithPlayer(context.Background(), plSpouse)
					data := welfaretypes.NewRankEmailData(endTime, ranking, rankNum, groupId, spouseEmailContentType)
					plSpouse.Post(message.NewScheduleMessage(sendRankRewards, ctx, data, nil))
				} else {
					sendRankMail(spouseId, groupId, ranking, rankNum, endTime, spouseEmailContentType)
				}
			}
		}

		//自己排行奖励
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl != nil {
			ctx := scene.WithPlayer(context.Background(), pl)
			data := welfaretypes.NewRankEmailData(endTime, ranking, rankNum, groupId, emailContentType)
			pl.Post(message.NewScheduleMessage(sendRankRewards, ctx, data, nil))
		} else {
			sendRankMail(playerId, groupId, ranking, rankNum, endTime, emailContentType)
		}

		//全服传音(排行第一)
		//目前配置写死了只有第一名会发送消息，需要改变此条件找策划添加字段
		rankFirst := int32(1)
		if timeTemp.IsChuanYin() && ranking == rankFirst {
			playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, fmt.Sprintf("%s", playerName))
			activityName := coreutils.FormatColor(chattypes.ColorTypeModuleName, fmt.Sprintf("%s", timeTemp.Name))
			groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
			if groupInterface == nil {
				return
			}
			groupTemp := groupInterface.(*welfareranktemplate.GroupTemplateRank)
			isOnLevel, openTemp := groupTemp.GetRankRewardsOpenTemp(ranking, int32(rankNum))
			itemStr := ""
			if isOnLevel {
				rewMap := openTemp.GetRewItemMap()
				for itemId, itemNum := range rewMap {
					itemTemp := item.GetItemService().GetItem(int(itemId))
					itemName := coreutils.FormatColor(itemTemp.GetQualityType().GetColor(), fmt.Sprintf("%s", itemTemp.FormateItemNameOfNum(itemNum)))
					linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
					nameLink := coreutils.FormatLink(itemName, linkArgs)
					itemStr = itemStr + "【" + nameLink + "】"
				}

				contentStr := timeTemp.ChuanyinText
				content := fmt.Sprintf(contentStr, playerName, activityName, itemStr)
				chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
			}
		}
	}
}

//排行榜奖励邮件回调
func sendRankRewards(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)

	data := result.(*welfaretypes.RankEmailData)
	ranking := data.Ranking
	rankNum := data.RankNum
	endTime := data.EndTime
	groupId := data.GroupId
	contentType := data.ContentType

	sendRankMail(tpl.GetId(), groupId, ranking, rankNum, endTime, contentType)
	return nil
}

func sendRankMail(playerId int64, groupId, ranking int32, rankNum, endTime int64, emailContentType welfaretypes.EmailContentType) {
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*welfareranktemplate.GroupTemplateRank)
	isOnLevel, openTemp := groupTemp.GetRankRewardsOpenTemp(ranking, int32(rankNum))
	if openTemp == nil {
		return
	}

	newItemDataList := ConvertToItemDataWithWelfareData(openTemp.GetEmailRewItemDataList(), openTemp.GetExpireType(), openTemp.GetExpireTime())
	title := lang.GetLangService().ReadLang(lang.EmailOpenActivityRankTitle)
	rankingText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", ranking))
	rankName := chatlogic.FormatMailKeyWordNoticeStr(openTemp.Label)
	content := ""
	langCode := emailContentType.ConvertToRankEmailContentLangCodeType(isOnLevel)
	switch emailContentType {
	case welfaretypes.EmailContentTypeMarryDevelop,
		welfaretypes.EmailContentTypeMarryDevelopSpouse,
		welfaretypes.EmailContentTypeCharm,
		welfaretypes.EmailContentTypeCharmSpouse:
		content = fmt.Sprintf(lang.GetLangService().ReadLang(langCode), rankName, rankNum, rankingText)
		break
	// case welfaretypes.EmailContentTypeBoatRaceForce:
	// 	title = openTemp.Label
	// 	rankNumText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", rankNum))
	// 	content = fmt.Sprintf(lang.GetLangService().ReadLang(langCode), rankName, rankNumText, rankingText)
	// 	break
	default:
		content = fmt.Sprintf(lang.GetLangService().ReadLang(langCode), rankName, rankingText)
	}
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		emaillogic.AddOfflineEmailItemLevel(playerId, title, content, endTime, newItemDataList)
	} else {
		emaillogic.AddEmailItemLevel(pl, title, content, endTime, newItemDataList)
	}
}

//奖励公告
func RewardsItemNoticeStr(rewMap map[int32]int32) string {
	itemNameLinkStr := ""
	for itemId, num := range rewMap {
		itemTemplate := item.GetItemService().GetItem(int(itemId))
		itemName := coreutils.FormatColor(itemTemplate.GetQualityType().GetColor(), coreutils.FormatNoticeStrUnderline(itemTemplate.FormateItemNameOfNum(num)))
		linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
		itemNameLink := coreutils.FormatLink(itemName, linkArgs)

		if len(itemNameLinkStr) == 0 {
			itemNameLinkStr += itemNameLink
		} else {
			itemNameLinkStr += ", " + itemNameLink
		}
	}
	return itemNameLinkStr
}

// 计算登录天数
func CountWelfareLoginDay(createTime int64) int32 {
	now := global.GetGame().GetTimeService().Now()
	diffDay, _ := timeutils.DiffDay(now, createTime)
	if diffDay < 0 {
		return 0
	}

	return diffDay + 1
}

// 每日充值-当前充值类型
func CountCycleDay(groupId int32) (cycleDay int32) {
	now := global.GetGame().GetTimeService().Now()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return -1
	}
	maxDay := groupInterface.GetMaxValue1()
	cycle := maxDay + 1
	openServerTime := welfare.GetWelfareService().GetServerStartTime() //global.GetGame().GetServerTime()
	diffDay, _ := timeutils.DiffDay(now, openServerTime)
	cycleDay = diffDay % cycle
	return
}

// 计算当前活动日
func CountCurActivityDay(groupId int32) int32 {
	now := global.GetGame().GetTimeService().Now()
	startTime, _ := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	if startTime == 0 {
		return 0
	}
	if now < startTime {
		return 0
	}

	diffDay, _ := timeutils.DiffDay(now, startTime)
	return diffDay
}

// 检查活动Id
func CheckGroupId(pl player.Player, typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType, groupId int32) bool {
	//活动时间判断
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
				"groupId":  groupId,
			}).Warn("welfare:运营活动,时间模板不存在")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityArgumentInvalid, fmt.Sprintf("%d", groupId))
		return false
	}

	// 循环活动校验
	// if timeTemp.IsXunHuan() && !IsOnXunHuanActivity(groupId) {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"typ":      typ,
	// 			"subType":  subType,
	// 			"groupId":  groupId,
	// 		}).Warn("welfare:运营活动,不是循环活动时间")
	// 	playerlogic.SendSystemMessage(pl, lang.OpenActivityArgumentInvalid, fmt.Sprintf("%d", groupId))
	// 	return false
	// }

	// // 合服循环活动校验
	// if timeTemp.IsMergeXunHuan() && !IsOnMergeXunHuanActivity(groupId) {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"typ":      typ,
	// 			"subType":  subType,
	// 			"groupId":  groupId,
	// 		}).Warn("welfare:运营活动,不是合服循环活动时间")
	// 	playerlogic.SendSystemMessage(pl, lang.OpenActivityArgumentInvalid, fmt.Sprintf("%d", groupId))
	// 	return false
	// }

	if !IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
				"groupId":  groupId,
			}).Warn("welfare:运营活动,不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return false
	}

	if timeTemp.GetOpenType() != typ || timeTemp.GetOpenSubType() != subType {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
				"groupId":  groupId,
			}).Warn("welfare:运营活动,活动类型不一致")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityArgumentInvalid, fmt.Sprintf("%d", groupId))
		return false
	}

	return true
}

//运营活动：物品绑定
func ConvertToItemData(itemMap map[int32]int32, expireType inventorytypes.NewItemLimitTimeType, expireTime int64) []*droptemplate.DropItemData {
	var newItemDataList []*droptemplate.DropItemData
	for itemId, num := range itemMap {
		level := int32(0)
		bind := itemtypes.ItemBindTypeUnBind
		itemGetTime := global.GetGame().GetTimeService().Now()

		newData := droptemplate.CreateItemDataWithExpire(itemId, num, level, bind, expireType, expireTime, itemGetTime)
		newItemDataList = append(newItemDataList, newData)
	}

	return newItemDataList
}

//部分时效部分非时效
func ConvertToItemDataWithWelfareData(welfareItemDataList []*gametemplate.WelfareItemData, expireType inventorytypes.NewItemLimitTimeType, expireTime int64) []*droptemplate.DropItemData {
	var newItemDataList []*droptemplate.DropItemData
	for _, data := range welfareItemDataList {
		itemId := data.ItemId
		num := data.Num
		level := int32(0)
		bind := itemtypes.ItemBindTypeUnBind
		itemGetTime := global.GetGame().GetTimeService().Now()
		initExpireType := expireType
		if data.ExpireFlag == 0 {
			initExpireType = inventorytypes.NewItemLimitTimeTypeNone
		}

		newData := droptemplate.CreateItemDataWithExpire(itemId, num, level, bind, initExpireType, expireTime, itemGetTime)
		newItemDataList = append(newItemDataList, newData)
	}

	return newItemDataList
}

// 是否活动时间
func IsOnActivityTime(groupId int32) bool {
	now := global.GetGame().GetTimeService().Now()
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		return false
	}

	if timeTemp.GetOpenTimeType() == welfaretypes.OpenTimeTypeOpenActivityNoMerge {
		mergeTime := welfare.GetWelfareService().GetServerMergeTime()
		//合服 不走了
		if mergeTime != 0 {
			return false
		}
	}

	// 循环活动
	if timeTemp.IsXunHuan() && !IsOnXunHuanActivity(groupId) {
		return false
	}

	// 合服循环活动
	if timeTemp.IsMergeXunHuan() && !IsOnMergeXunHuanActivity(groupId) {
		return false
	}

	if timeTemp.GetOpenTimeType() == welfaretypes.OpenTimeTypeNotTimeliness {
		return true
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	//是否处于时间段
	if startTime == 0 || endTime == 0 {
		return false
	}
	if now < startTime || now > endTime {
		return false
	}

	return true
}

// 升阶祝福暴击11-3
func IsCanAdvancedBlessCrit(advancedType welfaretypes.AdvancedType) (isDouble bool, addTimes int32) {
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypeBlessCrit
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !IsOnActivityTime(groupId) {
			continue
		}

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*advancedrewblesscrittemplate.GroupTemplateAdvancedBlessCrit)
		if advancedType != groupTemp.GetAdvancedType() {
			continue
		}

		doubleRate := groupTemp.GetCritRate()
		addTimes := groupTemp.GetExtralAddTimes()
		isDouble = mathutils.RandomHit(common.MAX_RATE, int(doubleRate))
		return isDouble, addTimes
	}

	return
}

// 装备宝库暴击7-6
func IsCanDrewBaoKuCrit() (isDouble bool, luckyPointCritNum int32, attendPointCritNum int32) {
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeBaoKuCrit
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !IsOnActivityTime(groupId) {
			continue
		}

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}

		groupTemp := groupInterface.(*drewbaokucrittemplate.GroupTemplateDrewBaoKuCrit)
		isDouble = mathutils.RandomHit(common.MAX_RATE, int(groupTemp.GetCritRate()))
		luckyPointCritNum = groupTemp.GetLuckyPointCritNum()
		attendPointCritNum = groupTemp.GetAttendPointCritNum()
		return
	}

	return
}

//活动开启邮件
func SendOpenNoticeMail(pl player.Player, group int32) {
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(group)
	if !IsOnActivityTime(group) {
		return
	}

	if len(timeTemp.MailDes) <= 0 {
		return
	}

	if !pl.IsFuncOpen(timeTemp.GetOpenFuncType()) {
		return
	}

	//是否邮件
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	if welfareManager.IsOpenMailRecord(group) {
		return
	}

	//发邮件
	emaillogic.AddEmail(pl, timeTemp.MailTitle, timeTemp.MailDes, timeTemp.GetMailRewItems())

	// 邮件记录
	welfareManager.AddOpenMailRecord(group)
}

// 成功购买礼包
func BuyZhuanShengGift(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject, giftBuyMap map[int32]int32) (newItemMap map[int32]int32, success bool) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	info, ok := obj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:类型转换失败")
		return
	}

	giftGroupTemp := welfaretemplate.GetWelfareTemplateService().GetDiscountZhuanShengGroupTemplate(obj.GetGroupId())
	newItemMap = make(map[int32]int32)
	for giftType, buyNum := range giftBuyMap {
		discountTemp := giftGroupTemp.GetDiscountZhuanShengTemplateByType(pl.GetRole(), pl.GetSex(), giftType)

		//判断背包空间
		for itemId, num := range discountTemp.GetItemMap() {
			_, ok := newItemMap[itemId]
			if ok {
				newItemMap[itemId] += num * buyNum
			} else {
				newItemMap[itemId] = num * buyNum
			}
		}
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(obj.GetGroupId())
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:模板不存在")
		return
	}
	firsOpenTemp := groupInterface.GetFirstOpenTemp()
	itemDataList := ConvertToItemData(newItemMap, firsOpenTemp.GetExpireType(), firsOpenTemp.GetExpireTime())
	if !inventoryManager.HasEnoughSlotsOfItemLevel(itemDataList) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"newItemMap": newItemMap,
			}).Warn("welfare:背包空间不足，请清理后再购买")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//添加物品
	itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
	itemReasonText := fmt.Sprintf(itemGetReason.String(), obj.GetActivityType(), obj.GetActivitySubType())
	flag := inventoryManager.BatchAddOfItemLevel(itemDataList, itemGetReason, itemReasonText)
	if !flag {
		panic("welfare: buy discount add item should be ok")
	}

	//更新信息
	for giftType, buyNum := range giftBuyMap {
		_, ok = info.BuyRecord[giftType]
		if !ok {
			info.BuyRecord[giftType] = buyNum
		} else {
			info.BuyRecord[giftType] += buyNum
		}
	}
	welfareManager.UpdateObj(obj)
	success = true
	return
}

// 额外元宝
func AddExtralRewGold(pl player.Player, chargeGold int32, returnType welfaretypes.ChargeReturnType, returnRatio int32) {
	extralNum := int64(math.Ceil(float64(chargeGold*returnRatio) / float64(common.MAX_RATE)))
	if extralNum > 0 {
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		addGodlReason := commonlog.GoldLogReasonOpenActivityChargeReturn
		switch returnType {
		case welfaretypes.ChargeReturnTypeBindGold:
			{
				propertyManager.AddGold(extralNum, true, addGodlReason, addGodlReason.String())
			}
		case welfaretypes.ChargeReturnTypeGold:
			{
				propertyManager.AddGold(extralNum, false, addGodlReason, addGodlReason.String())
			}
		}
		propertylogic.SnapChangedProperty(pl)
	}

	return
}

//城战助威邮件
func AllianceCheerEndMail(pl player.Player, groupId int32) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		return
	}

	endTime := obj.GetEndTime()
	info := obj.GetActivityData().(*alliancecheertypes.AllianceCheerInfo)
	if info.CheerGoldNum < 1 {
		info.IsEmail = true
	}

	if !info.IsEmail {
		rewMap := make(map[int32]int32)
		rewMap[constanttypes.GoldItem] = info.CheerGoldNum

		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
		title := timeTemp.Name
		acName := chatlogic.FormatMailKeyWordNoticeStr(timeTemp.Name)
		cheerGoldText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), info.CheerGoldNum))
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityAllianceCheerMailContent), acName, cheerGoldText, cheerGoldText)
		emaillogic.AddEmailDefinTime(pl, title, econtent, endTime, rewMap)
	}

	info.IsEmail = true
	welfareManager.UpdateObj(obj)
}

// 获取循环活动
func GetXunHuanGroupList() (groupIdList []int32) {
	// 是否循环活动
	if !welfare.GetWelfareService().IsOnXunHuan() {
		return
	}

	//循环活动-刷新跨天
	openServerTime := welfare.GetWelfareService().GetServerStartTime()
	arrGroup, curDay := welfare.GetWelfareService().GetXunHuanInfo()
	xunhuanTemp := welfaretemplate.GetWelfareTemplateService().GetActivityXunHuanTemplate(openServerTime, arrGroup, curDay)
	if xunhuanTemp == nil {
		return
	}
	groupIdList = xunhuanTemp.GetGroupIdList()
	return
}

func IsOnXunHuanActivity(groupId int32) bool {
	groupIdList := GetXunHuanGroupList()
	if !coreutils.ContainInt32(groupIdList, groupId) {
		return false
	}

	return true
}

func IsOnMergeXunHuanActivity(groupId int32) (isCircle bool) {
	openTime := welfare.GetWelfareService().GetServerStartTime()
	mergeTime := welfare.GetWelfareService().GetServerMergeTime()
	mergeXunHuanTemp := welfaretemplate.GetWelfareTemplateService().GetActivityMergeXunHuanTemplate(openTime, mergeTime)
	if mergeXunHuanTemp == nil {
		return
	}

	groupIdList := mergeXunHuanTemp.GetGroupIdList()
	if !coreutils.ContainInt32(groupIdList, groupId) {
		return
	}

	isCircle = true
	return
}
