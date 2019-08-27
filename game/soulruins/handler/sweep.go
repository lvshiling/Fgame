package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	soulruinseventtypes "fgame/fgame/game/soulruins/event/types"
	soulruinslogic "fgame/fgame/game/soulruins/logic"
	"fgame/fgame/game/soulruins/pbutil"
	playersoulruins "fgame/fgame/game/soulruins/player"
	"fgame/fgame/game/soulruins/soulruins"
	soulruinstypes "fgame/fgame/game/soulruins/types"
	weeklogic "fgame/fgame/game/week/logic"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOULRUINS_SWEEP_TYPE), dispatch.HandlerFunc(handleSoulRuinsSweep))
}

//处理帝陵遗迹扫荡信息
func handleSoulRuinsSweep(s session.Session, msg interface{}) (err error) {
	log.Debug("soulruins:处理获取帝陵遗迹扫荡消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulRuinsSweep := msg.(*uipb.CSSoulRuinsSweep)
	chapterInfo := csSoulRuinsSweep.GetChapterInfo()
	chapter := chapterInfo.GetChapter()
	typ := chapterInfo.GetTyp()
	level := csSoulRuinsSweep.GetLevel()
	num := csSoulRuinsSweep.GetNum()

	err = soulRuinsSweep(tpl, chapter, soulruinstypes.SoulRuinsType(typ), level, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"level":    level,
				"num":      num,
				"error":    err,
			}).Error("soulruins:处理获取帝陵遗迹扫荡消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soulruins:处理获取帝陵遗迹扫荡消息完成")
	return nil

}

//获取帝陵遗迹扫荡界面信息的逻辑
func soulRuinsSweep(pl player.Player, chapter int32, typ soulruinstypes.SoulRuinsType, level int32, num int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	flag := manager.IsValid(chapter, typ, level)
	if !flag || num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"level":    level,
				"num":      num,
			}).Warn("soulruins:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//挑战次数判断
	flag = manager.HasEnoughChallengeNum(num)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"level":    level,
				"num":      num,
			}).Warn("soulruins:挑战次数不足,无法进行扫荡")
		playerlogic.SendSystemMessage(pl, lang.SoulRuinsChallengeNumNotEnough)
		return
	}

	soulRuinsTemplate := soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(chapter, typ, level)
	needItems := make(map[int32]int32)
	//免费扫荡
	if !weeklogic.IsSeniorWeek(pl) {
		needItemMap := soulRuinsTemplate.GetSweepNeedItemMap()
		for itemId, itemNum := range needItemMap {
			needItems[itemId] = itemNum * num
		}
	}

	//是否已通关
	flag = manager.IfSoulRuinsExist(chapter, typ, level)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"level":    level,
				"num":      num,
			}).Warn("soulruins:关卡挑战成功后,才能扫荡")
		playerlogic.SendSystemMessage(pl, lang.SoulRuinsNotPassedToSweep)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//判断物品
	if len(needItems) > 0 {
		if !inventoryManager.HasEnoughItems(needItems) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"chapter":  chapter,
					"typ":      typ,
					"level":    level,
					"num":      num,
				}).Warn("soulruins:当前扫荡券不足,无法进行扫荡")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//扫荡奖励判断
	dropItemList, rewData, isReturn, err := soulruinslogic.GiveSoulRuinsSweepReward(pl, chapter, typ, level, num)
	if err != nil || isReturn {
		return
	}

	//消耗物品
	if len(needItems) != 0 {
		inventoryLogReason := commonlog.InventoryLogReasonSoulRuinsSweep
		reasonText := fmt.Sprintf(inventoryLogReason.String(), chapter, typ, level, num)
		flag = inventoryManager.BatchRemove(needItems, inventoryLogReason, reasonText)
		if !flag {
			panic(fmt.Errorf("soulruins: soulRuinsSweep BatchRemove should be ok"))
		}
	}

	//同步物品
	inventorylogic.SnapInventoryChanged(pl)

	//消耗挑战次数
	manager.UseChallengeNum(num)
	//发送事件
	soulRuinsId := int32(soulRuinsTemplate.TemplateId())
	eventData := soulruinseventtypes.CreateSoulRuinsFinishEventData(soulRuinsId, num)
	gameevent.Emit(soulruinseventtypes.EventTypeSoulruinsFinish, pl, eventData)
	gameevent.Emit(soulruinseventtypes.EventTypeSoulruinsSweep, pl, num)

	numObj := manager.GetSoulRuinsNum()
	scSoulRuinsChallenge := pbutil.BuildSCSoulRuinsSweep(numObj, rewData, dropItemList)
	pl.SendMsg(scSoulRuinsChallenge)
	return
}
