package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_RING_FEED_TYPE), dispatch.HandlerFunc(handleMarryRingFeed))
}

//处理婚戒培养信息
func handleMarryRingFeed(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理婚戒培养消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = marryRingFeed(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("marry:处理婚戒培养消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理婚戒培养消息完成")
	return nil
}

//处理婚戒培养信息逻辑
func marryRingFeed(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	ringLevel := marryInfo.RingLevel
	spouseId := marryInfo.SpouseId
	sringLevel := int32(0)

	//未婚
	if marryInfo.Status == marrytypes.MarryStatusTypeUnmarried ||
		marryInfo.Status == marrytypes.MarryStatusTypeDivorce {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("marry:未结婚,无法培养")
		playerlogic.SendSystemMessage(pl, lang.MarryRingFeedNoMarried)
		return
	}
	nextRingLevel := ringLevel + 1
	ringTemplate := marrytemplate.GetMarryTemplateService().GetMarryRingTemplate(marryInfo.Ring, nextRingLevel)
	if ringTemplate == nil {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("marry:已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.MarryRingReachLimit)
		return
	}

	//获取配偶婚戒等级
	marryObj := marry.GetMarryService().GetMarry(pl.GetId())
	if marryObj == nil {
		return
	}
	if marryObj.PlayerId == pl.GetId() {
		sringLevel = marryObj.PlayerRingLevel
	} else {
		sringLevel = marryObj.SpouseRingLevel
	}
	if spouseId != 0 {
		//等级差
		levelDiff := marrytemplate.GetMarryTemplateService().GetMarryConstRingLevelDiff()
		diffLevel := ringLevel - sringLevel
		if diffLevel >= levelDiff || diffLevel <= -levelDiff {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("marry:夫妻戒指等级差不满足,无法提升")
			playerlogic.SendSystemMessage(pl, lang.MarryRingFeedLevelNoEnough)
			return
		}
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := ringTemplate.UseGold
	//进阶需要消耗绑元
	costBindGold := ringTemplate.UseBindGold
	//进阶需要消耗的银两
	costSilver := int64(ringTemplate.UseSilver)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := ringTemplate.UseItem
	useItemTemplate := ringTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = ringTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("marry:物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("marry:银两不足,无法培养")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//是否足够绑元
	if costBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("marry:元宝不足,无法培养")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costBindGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("marry:元宝不足,无法培养")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonMarryRingFeedCost.String()
	reasonSliverText := commonlog.SilverLogReasonMarryRingFeedCost.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonMarryRingFeedCost, reasonGoldText, costSilver, commonlog.SilverLogReasonMarryRingFeedCost, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("marry: marryRingFeed Cost should be ok"))
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonMarryRingFeedUse
		reasonText := fmt.Sprintf(inventoryReason.String(), marryInfo.RingLevel)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("marry: marryRingFeed use item shoud be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//同步元宝
	if costBindGold != 0 || costGold != 0 || costSilver != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//婚戒培养判断
	pro, _, sucess := marrylogic.MarryRingFeed(pl, marryInfo.RingNum, marryInfo.RingExp, ringTemplate)
	manager.RingFeed(pro, sucess)
	if sucess {
		//同步属性
		marrylogic.MarryPropertyChanged(pl)
		spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
		//婚戒等级同步给配偶
		if spl != nil {
			scMarryRLevelChange := pbuitl.BuildSCMarryRLevelChange(pl.GetId(), marryInfo.RingLevel)
			spl.SendMsg(scMarryRLevelChange)
		}
	}

	scMarryRingFeed := pbuitl.BuildSCMarryRingFeed(marryInfo.RingLevel, marryInfo.RingExp)
	pl.SendMsg(scMarryRingFeed)
	return
}
