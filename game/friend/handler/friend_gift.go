package handler

import (
	"context"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	message "fgame/fgame/common/message"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/friend/friend"
	"fgame/fgame/game/friend/pbutil"
	playerfriend "fgame/fgame/game/friend/player"
	friendtypes "fgame/fgame/game/friend/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_GIFT_TYPE), dispatch.HandlerFunc(handleFriendGift))
}

//处理好友礼物
func handleFriendGift(s session.Session, msg interface{}) (err error) {
	log.Debug("friend:处理好友礼物")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFriendGift := msg.(*uipb.CSFriendGift)
	friendId := csFriendGift.GetFriendId()
	itemId := csFriendGift.GetItemId()
	num := csFriendGift.GetNum()
	if num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友礼物,参数不对")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	msgId := csFriendGift.GetMsgId()
	auto := csFriendGift.GetAuto()
	msgContent := csFriendGift.GetMsgContent()
	friendGift(tpl, friendId, itemId, num, msgId, auto, msgContent)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理好友礼物,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理好友礼物,完成")
	return nil

}

//处理好友礼物
func friendGift(pl player.Player, friendId int64, itemId int32, num int32, msgId int32, auto int32, msgContent string) (err error) {
	if friendId == pl.GetId() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:不能给自己赠送")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	itemTempalte := item.GetItemService().GetItem(int(itemId))
	if itemTempalte == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"friendId": friendId,
			}).Warn("friend:物品不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if itemTempalte.GetItemType() != itemtypes.ItemTypeXianHua &&
		itemTempalte.GetItemType() != itemtypes.ItemTypeBiaoBai {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"friendId": friendId,
			}).Warn("friend:物品不是鲜花")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//亲密度
	pointNum := itemTempalte.TypeFlag1 * num
	//魅力值
	charmNum := itemTempalte.TypeFlag2 * num
	//表白经验
	developExp := itemTempalte.TypeFlag3 * num

	marryManager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	if itemTempalte.GetItemType() == itemtypes.ItemTypeBiaoBai && marryManager.GetSpouseId() != friendId {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"friendId": friendId,
			}).Warn("marry:处理赠送表白礼物失败，对方不是您的伴侣")
		playerlogic.SendSystemMessage(pl, lang.MarryNotCouple)
		return
	}

	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	//判断是否可以加为好友
	isFriend := friendManager.IsFriend(friendId)
	if !isFriend {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:对方不是您好友")
		playerlogic.SendSystemMessage(pl, lang.FriendIsNotFriend)
		return
	}
	//判断是否在线
	fri := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	//拉黑显示不在线
	flag := friendManager.IsBlacked(friendId)
	if (fri == nil && itemTempalte.GetItemType() == itemtypes.ItemTypeBiaoBai) || flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"friendId": friendId,
			}).Warn("friend:处理好友礼物,用户不在线")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoOnline)
		return
	}
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//giftId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeGiftId)

	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	numOfGifts := inventoryManager.NumOfItems(itemId)
	costGiftNum := num
	//扣除物品
	if numOfGifts < num {
		//不自动购买
		if auto == 0 {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"friendId": friendId,
				}).Warn("friend:处理好友礼物,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		costGiftNum = numOfGifts
		//计算需要购买的
		needBuy := num - numOfGifts
		// shopTemplate := shop.GetShopService().GetShopTemplateByItem(itemId)
		// if shopTemplate == nil {
		// 	log.WithFields(
		// 		log.Fields{
		// 			"playerId": pl.GetId(),
		// 			"friendId": friendId,
		// 		}).Warn("friend:处理好友礼物,物品不足")
		// 	playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		// 	return
		// }
		// needBindGold := int64(0)
		// needGold := int64(0)
		// needSilver := int64(0)
		// switch shopTemplate.GetShopConsumeType() {
		// case shoptypes.ShopConsumeTypeBindGold:
		// 	needBindGold = int64(shopTemplate.ConsumeData1 * needBuy)
		// 	break
		// case shoptypes.ShopConsumeTypeGold:
		// 	needGold = int64(shopTemplate.ConsumeData1 * needBuy)
		// 	break
		// case shoptypes.ShopConsumeTypeSliver:
		// 	needSilver = int64(shopTemplate.ConsumeData1) * int64(needBuy)
		// 	break
		// }

		needBindGold := int64(0)
		needGold := int64(0)
		needSilver := int64(0)
		if needBuy > 0 {
			if !shop.GetShopService().ShopIsSellItem(itemId) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"friendId": friendId,
				}).Warn("friend:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, itemId, needBuy)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"friendId": friendId,
				}).Warn("friend:购买鲜花物品失败")
				playerlogic.SendSystemMessage(pl, lang.ShopFlowerAutoBuyItemFail)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			needGold += shopNeedGold
			needBindGold += shopNeedBindGold
			needSilver += shopNeedSilver
		}

		if !propertyManager.HasEnoughCost(needBindGold, needGold, needSilver) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"friendId": friendId,
				}).Warn("friend:处理好友礼物,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		//更新自动购买每日限购次数
		if len(shopIdMap) != 0 {
			shoplogic.ShopDayCountChanged(pl, shopIdMap)
		}

		goldReason := commonlog.GoldLogReasonGiftBuy
		goldReasonText := fmt.Sprintf(goldReason.String(), itemId, needBuy)
		silverReason := commonlog.SilverLogReasonGiftBuy
		silverReasonText := fmt.Sprintf(silverReason.String(), itemId, needBuy)

		flag := propertyManager.Cost(needBindGold, needGold, goldReason, goldReasonText, needSilver, silverReason, silverReasonText)
		if !flag {
			panic(fmt.Errorf("friend:购买好友礼物,应该成功"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	if costGiftNum != 0 {
		inventoryReason := commonlog.InventoryLogReasonGift
		flag := inventoryManager.UseItem(itemId, costGiftNum, inventoryReason, inventoryReason.String())
		if !flag {
			panic(fmt.Errorf("friend:扣除礼物物品,应该成功"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	if itemTempalte.GetItemType() == itemtypes.ItemTypeBiaoBai {
		beforeDevelopExp := marryManager.GetMarryDevelopExp()
		marryManager.AddDevelopExp(developExp)
		//表白经验后台日志
		curDevelopExp := marryManager.GetMarryDevelopExp()
		reason := commonlog.MarryLogReasonDevelopExpByItem
		reasonText := fmt.Sprintf(commonlog.MarryLogReasonDevelopExpByItem.String(), itemId, num)
		logEventData := marryeventtypes.CreatePlayerDevelopExpLogEventData(beforeDevelopExp, curDevelopExp, reason, reasonText)
		gameevent.Emit(marryeventtypes.EventTypePlayerMarryDevelopExpLog, pl, logEventData)
		gameevent.Emit(marryeventtypes.EventTypePlayerMarryBiaoBai, pl, num)

		//添加表白日志记录
		if itemTempalte.GetQualityType() >= itemtypes.ItemQualityTypePurple {
			logData := friendtypes.NewAllMarryDevelopData(pl.GetId(), fri.GetId(), pl.GetName(), fri.GetName(), itemId, num, charmNum, developExp, msgContent)
			friend.GetFriendService().AddMarryDevelopLog(logData)
			friendManager.AddMarryDevelopSendLog(logData)
			ctx := scene.WithPlayer(context.Background(), fri)
			fri.Post(message.NewScheduleMessage(friendAddMarryDevelopRecvLog, ctx, logData, nil))
		}
		scFriendGiftRecv := pbutil.BuildSCFriendGiftRecv(pl.GetId(), itemId, num, msgId, msgContent)
		fri.SendMsg(scFriendGiftRecv)
	}

	scFriendGift := pbutil.BuildSCFriendGift(friendId, itemId, num, msgId, auto)
	pl.SendMsg(scFriendGift)
	friendGiftEventData := friendeventtypes.CreateFriendGiftEventData(pl.GetId(), friendId, itemId, num, pointNum, charmNum)
	gameevent.Emit(friendeventtypes.EventTypeFriendGift, pl, friendGiftEventData)
	return
}

func friendAddMarryDevelopRecvLog(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	logData := result.(*friendtypes.MarryDevelopLogData)
	friendManager := tpl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	friendManager.AddMarryDevelopRecvLog(logData)
	return nil
}
