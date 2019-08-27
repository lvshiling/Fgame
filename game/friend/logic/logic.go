package logic

import (
	"context"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/common/message"
	chatpbutil "fgame/fgame/game/chat/pbutil"
	chattypes "fgame/fgame/game/chat/types"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	friendentity "fgame/fgame/game/friend/entity"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/friend/pbutil"
	playerfriend "fgame/fgame/game/friend/player"
	friendtemplate "fgame/fgame/game/friend/template"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/idutil"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//广播鲜花消息
func BroadcastMsg(pl player.Player, itemId int32, itemCount int32, playerIdList []int64, playerNameList []string) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	scFriendGiveFlowerLight := pbutil.BuildSCFriendGiveFlowerLight(itemId, itemCount, playerIdList, playerNameList)
	s.BroadcastMsg(scFriendGiveFlowerLight)
}

//好友度变化推送
func FrinedPointChanged(pl player.Player, friendId int64, point int32) (err error) {
	scFriendPointChange := pbutil.BuildSCFriendPointChange(friendId, point)
	pl.SendMsg(scFriendPointChange)
	return
}

func ReverseFriendLog(playerId int64, friendId int64, oprType friendtypes.FriendLogType) (err error) {
	//写离线日志
	id, err := idutil.GetId()
	if err != nil {
		return err
	}
	now := global.GetGame().GetTimeService().Now()
	friendLogEntity := &friendentity.PlayerFriendLogEntity{
		Id:         id,
		PlayerId:   playerId,
		FriendId:   friendId,
		Type:       int32(oprType),
		UpdateTime: now,
		CreateTime: now,
		DeleteTime: 0,
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(friendLogEntity)
	return nil
}

// 推送消息
func BroadcastFriendNotice(pl player.Player, noticeType friendtypes.FriendNoticeType, condition int32, args string) {
	scMsg := pbutil.BuildSCFriendNoticeBroadcase(noticeType, pl.GetId(), pl.GetName(), pl.GetRole(), pl.GetSex(), condition, args)
	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	frList := friendManager.GetFriends()
	for frId, _ := range frList {
		frPl := player.GetOnlinePlayerManager().GetPlayerById(frId)
		if frPl == nil {
			continue
		}
		frPl.SendMsg(scMsg)
	}
}

type FriendAskForReplyData struct {
	friendId     int64
	friendName   string
	isAgree      bool
	itemId       int32
	num          int32
	level        int32
	bind         itemtypes.ItemBindType
	propertyData inventorytypes.ItemPropertyData
}

// 赠送好友物品
func GiveItemForFriend(pl, fri player.Player, itemId int32, isAgree bool) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	it := inventoryManager.GetBageLastItem(itemId)

	if it == nil || it.IsEmpty() {
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		isAgree = false
	}

	// 好友
	data := &FriendAskForReplyData{
		friendId:   pl.GetId(),
		friendName: pl.GetName(),
		isAgree:    isAgree,
	}
	if isAgree {
		data.itemId = it.ItemId
		data.num = it.Num
		data.level = it.Level
		data.bind = it.BindType
		data.propertyData = it.PropertyData

		//移除
		useReason := commonlog.InventoryLogReasonFriendGiveUse
		flag, _ := inventoryManager.RemoveIndex(inventorytypes.BagTypePrim, it.Index, it.Num, useReason, useReason.String())
		if !flag {
			panic("friend:好友索取，移除物品应该成功")
		}
		inventorylogic.SnapInventoryChanged(pl)

		//赠送日志
		reason := commonlog.FriendLogReasonAgreeGive
		reasonText := fmt.Sprintf(reason.String(), fri.GetId(), data.itemId, data.num)
		eventData := friendeventtypes.CreateFriendGiveEventData(reason, reasonText)
		gameevent.Emit(friendeventtypes.EventTypeFriendGiveLog, pl, eventData)
	}

	//发送个人信息
	noticeConstantTemp := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeConstanTemplate()
	msgContent := ""
	if isAgree {
		msgContent = noticeConstantTemp.SongliTongyi
	} else {
		msgContent = noticeConstantTemp.SongliJujue
	}
	chatRecv := chatpbutil.BuildSCChatRecvWithCliArgs(pl.GetId(), pl.GetName(), chattypes.ChannelTypePerson, fri.GetId(), chattypes.MsgTypeText, []byte(msgContent), "")
	fri.SendMsg(chatRecv)

	ctx := scene.WithPlayer(context.Background(), fri)
	msg := message.NewScheduleMessage(onFriendGetItem, ctx, data, nil)
	fri.Post(msg)
}

func onFriendGetItem(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)
	data := result.(*FriendAskForReplyData)

	//发送个人信息
	friend := player.GetOnlinePlayerManager().GetPlayerById(data.friendId)
	if friend != nil {
		noticeConstantTemp := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeConstanTemplate()
		msgContent := ""
		if data.isAgree {
			msgContent = noticeConstantTemp.ShouliTongyi
		} else {
			msgContent = noticeConstantTemp.ShouliJujue
		}
		chatRecv := chatpbutil.BuildSCChatRecvWithCliArgs(pl.GetId(), pl.GetName(), chattypes.ChannelTypePerson, data.friendId, chattypes.MsgTypeText, []byte(msgContent), "")
		friend.SendMsg(chatRecv)
	}

	if data.isAgree {
		now := global.GetGame().GetTimeService().Now()
		itemData := inventorylogic.ConverToGoldEquipItemData(data.itemId, data.num, data.level, data.bind, data.propertyData)
		itemList := []*droptemplate.DropItemData{itemData}
		title := lang.GetLangService().ReadLang(lang.FriendAskForGiveTitle)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FriendAskForGiveContent), data.friendName)
		emaillogic.AddEmailItemLevel(pl, title, content, now, itemList)
	}

	scMsg := pbutil.BuildSCFriendFeedbackAskForReplyNotice(data.isAgree, data.friendName)
	pl.SendMsg(scMsg)

	return nil
}

type FriendFeedbackData struct {
	feedbackPlId int64
	feedbackName string
	noticeType   friendtypes.FriendNoticeType
	feedbackType friendtypes.FriendFeedbackType
	condition    int32
}

// 好友反馈推送
func FriendNoticeFeedbackNotice(fri, pl player.Player, noticeType friendtypes.FriendNoticeType, feedbackType friendtypes.FriendFeedbackType, condition int32) {

	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	if friendManager.IsLimitAdmire(fri.GetId()) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:处理赠送好友错误，已达次数上限")
		return
	}

	//玩家
	flag := addCongratulateRew(pl, noticeType, feedbackType, condition)
	if !flag {
		return
	}
	friendManager.AddAdmireTimes(fri.GetId())

	//好友
	data := &FriendFeedbackData{
		feedbackPlId: pl.GetId(),
		feedbackName: pl.GetName(),
		noticeType:   noticeType,
		feedbackType: feedbackType,
		condition:    condition,
	}
	ctx := scene.WithPlayer(context.Background(), fri)
	msg := message.NewScheduleMessage(onFriendFeedback, ctx, data, nil)
	fri.Post(msg)
}

func onFriendFeedback(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	data := result.(*FriendFeedbackData)
	pl := p.(player.Player)

	flag := addFeedbackRew(pl, data.noticeType, data.feedbackType, data.condition)
	if !flag {
		return nil
	}

	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	friendManager.AddFriendFeedback(data.feedbackPlId, data.feedbackName, data.condition, data.noticeType, data.feedbackType)

	frScMsg := pbutil.BuildSCFriendNoticeFeedbackNotice(data.feedbackPlId, data.feedbackName, data.noticeType, data.feedbackType, data.condition)
	pl.SendMsg(frScMsg)
	return nil
}

// 被祝贺奖励
func addFeedbackRew(pl player.Player, noticeType friendtypes.FriendNoticeType, feedbackType friendtypes.FriendFeedbackType, condition int32) bool {
	noticeTemp := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeTemplateByCondition(noticeType, condition)
	if noticeTemp == nil {
		return false
	}
	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	if friendManager.IsLimitCongraluteRew() {
		return false
	}

	addExp := int64(0)
	addExpPoint := int64(0)
	switch feedbackType {
	case friendtypes.FriendFeedbackTypeEgg:
		{
			addExp = int64(noticeTemp.JiDanRewardExp)
			addExpPoint = int64(noticeTemp.JiDanRewardExpPoint)
		}
	case friendtypes.FriendFeedbackTypeFlower:
		{
			addExp = int64(noticeTemp.XianHuaRewardExp)
			addExpPoint = int64(noticeTemp.XianHuaRewardExpPoint)
		}
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	addReason := commonlog.LevelLogReasonFriendNoticeRew
	addReasonString := fmt.Sprintf(addReason.String(), feedbackType)
	if addExp > 0 {
		propertyManager.AddExp(addExp, addReason, addReasonString)
	}
	if addExpPoint > 0 {
		propertyManager.AddExpPoint(addExpPoint, addReason, addReasonString)
	}
	friendManager.AddCongratulateTimes()
	propertylogic.SnapChangedProperty(pl)

	return true
}

//祝贺奖励
func addCongratulateRew(pl player.Player, noticeType friendtypes.FriendNoticeType, feedbackType friendtypes.FriendFeedbackType, condition int32) bool {
	noticeTemp := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeTemplateByCondition(noticeType, condition)
	if noticeTemp == nil {
		return false
	}

	// 背包
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlots(noticeTemp.GetRewItemMap()) {
		//邮件
		emaillogic.AddEmail(pl, "背包空间不足", "赞赏好友奖励邮件", noticeTemp.GetRewItemMap())
	} else {
		addReason := commonlog.InventoryLogReasonFriendNoticeFeedback
		addReasonString := fmt.Sprintf(addReason.String(), noticeType, condition)
		flag := inventoryManager.BatchAdd(noticeTemp.GetRewItemMap(), addReason, addReasonString)
		if !flag {
			panic("friend:添加赞赏好友奖励失败")
		}

		inventorylogic.SnapInventoryChanged(pl)
	}

	return true
}
