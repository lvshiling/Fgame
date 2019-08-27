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
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_TREE_FEED_TYPE), dispatch.HandlerFunc(handleMarryTreeFeed))
}

//处理爱情树培养信息
func handleMarryTreeFeed(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理爱情树培养消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = marryTreeFeed(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("marry:处理爱情树培养消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理爱情树培养消息完成")
	return nil
}

//处理爱情树培养信息逻辑
func marryTreeFeed(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	treeLevel := marryInfo.TreeLevel

	//未婚
	if marryInfo.Status == marrytypes.MarryStatusTypeUnmarried ||
		marryInfo.Status == marrytypes.MarryStatusTypeDivorce {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("marry:未结婚,无法培养")
		playerlogic.SendSystemMessage(pl, lang.MarryTreeFeedNoMarried)
		return
	}
	nextTreeLevel := treeLevel + 1
	treeTemplate := marrytemplate.GetMarryTemplateService().GetMarryLoveTreeTemplate(nextTreeLevel)
	if treeTemplate == nil {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("marry:已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.MarryTreeReachLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := treeTemplate.UseGold
	//进阶需要消耗绑元
	costBindGold := treeTemplate.UseBindGold
	//进阶需要消耗的银两
	costSilver := int64(treeTemplate.UseSilver)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := treeTemplate.UseItem
	useItemTemplate := treeTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = treeTemplate.ItemCount
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
	reasonGoldText := commonlog.GoldLogReasonMarryTreeFeedCost.String()
	reasonSliverText := commonlog.SilverLogReasonMarryTreeFeedCost.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonMarryTreeFeedCost, reasonGoldText, costSilver, commonlog.SilverLogReasonMarryTreeFeedCost, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("marry: marryTreeFeed Cost should be ok"))
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonMarryTreeFeedUse
		reasonText := fmt.Sprintf(inventoryReason.String(), marryInfo.TreeLevel)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("marry: marryTreeFeed use item shoud be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//同步元宝
	if costBindGold != 0 || costGold != 0 || costSilver != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//爱情树培养判断
	pro, _, sucess := marrylogic.MarryTreeFeed(marryInfo.TreeNum, marryInfo.TreeExp, treeTemplate)
	manager.TreeFeed(pro, sucess)
	if sucess {
		//同步属性
		marrylogic.MarryPropertyChanged(pl)
		spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
		//爱情树等级同步给配偶
		if spl != nil {
			scMarryTLevelChange := pbuitl.BuildSCMarryTLevelChange(pl.GetId(), marryInfo.TreeLevel)
			spl.SendMsg(scMarryTLevelChange)
		}
	}
	scMarryTreeFeed := pbuitl.BuildSCMarryTreeFeed(marryInfo.TreeLevel, marryInfo.TreeExp)
	pl.SendMsg(scMarryTreeFeed)
	return
}
