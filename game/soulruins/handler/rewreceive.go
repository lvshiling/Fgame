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
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/soulruins/pbutil"
	playersoulruins "fgame/fgame/game/soulruins/player"
	"fgame/fgame/game/soulruins/soulruins"
	soulruinstypes "fgame/fgame/game/soulruins/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOULRUINS_REWRECEIVE_TYPE), dispatch.HandlerFunc(handleSoulRuinsRewReceive))
}

//处理帝陵遗迹星级奖励信息
func handleSoulRuinsRewReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("soulruins:处理获取帝陵遗迹星级奖励消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulRuinsRewReceive := msg.(*uipb.CSSoulRuinsRewReceive)
	chapterInfo := csSoulRuinsRewReceive.GetChapterInfo()
	chapter := chapterInfo.GetChapter()
	typ := chapterInfo.GetTyp()

	err = soulRuinsRewReceive(tpl, chapter, soulruinstypes.SoulRuinsType(typ))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"error":    err,
			}).Error("soulruins:处理获取帝陵遗迹星级奖励消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"chapter":  chapter,
			"typ":      typ,
		}).Debug("soulruins:处理获取帝陵遗迹星级奖励消息完成")
	return nil

}

//获取帝陵遗迹星级奖励界面信息的逻辑
func soulRuinsRewReceive(pl player.Player, chapter int32, typ soulruinstypes.SoulRuinsType) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	flag := manager.IsChapterAndTypValid(chapter, typ)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
			}).Warn("soulruins:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//是否已领取过
	flag = manager.IfSoulRuinsRewReceived(chapter, typ)
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
			}).Warn("soulruins:该章节星级奖励已领取过")
		playerlogic.SendSystemMessage(pl, lang.SoulRuinsRewChapterRepeatReceive)
		return
	}

	//章节星级数是否达标
	flag = manager.IfStarsReachStandard(chapter, typ)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
			}).Warn("soulruins:您当前星数不满足领取条件")
		playerlogic.SendSystemMessage(pl, lang.SoulRuinsChapterStarNotEnough)
		return
	}

	starTemplate := soulruins.GetSoulRuinsService().GetSoulRuinsStarTemplate(chapter, typ)
	rewData := starTemplate.GetRewData()
	rewItemMap := starTemplate.GetRewItemMap()
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//判断背包是否足够
	if len(rewItemMap) != 0 {
		flag = inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
			}).Warn("soulruins:背包空间不足,清理后再来")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
		inventoryLogReason := commonlog.InventoryLogReasonSoulRuinsRewChapter
		reasonText := fmt.Sprintf(inventoryLogReason.String(), chapter, typ)
		flag = inventoryManager.BatchAdd(rewItemMap, commonlog.InventoryLogReasonSoulRuinsRewChapter, reasonText)
		if !flag {
			panic(fmt.Errorf("soulruins: soulRuinsRewReceive BatchAdd should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//章节星级奖励领取记录
	flag = manager.RewReceiveChapter(chapter, typ)
	if !flag {
		panic(fmt.Errorf("soulruins: soulRuinsRewReceive RewReceiveChapter should be ok"))
	}

	//添加奖励属性
	if rewData != nil {
		goldLog := commonlog.GoldLogReasonSoulRuinsRewChapter
		silverLog := commonlog.SilverLogReasonSoulRuinsRewChapter
		levelReason := commonlog.LevelLogReasonSoulRuinsRewChapter
		goldReasonText := fmt.Sprintf(goldLog.String(), chapter, typ)
		silverReasonText := fmt.Sprintf(silverLog.String(), chapter, typ)
		reasonText := fmt.Sprintf(levelReason.String(), chapter, typ)
		flag = propertyManager.AddRewData(rewData, goldLog, goldReasonText, silverLog, silverReasonText, levelReason, reasonText)
		if !flag {
			panic(fmt.Errorf("soulruins: soulRuinsRewReceive AddRewData should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	rewChapterObj := manager.GetSoulRuinsRewChapter(chapter, typ)
	scSoulRuinsRewReceive := pbutil.BuildSCSoulRuinsRewReceive(rewChapterObj)
	pl.SendMsg(scSoulRuinsRewReceive)
	return
}
