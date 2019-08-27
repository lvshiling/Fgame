package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/dan/dan"
	"fgame/fgame/game/dan/pbutil"
	playerdan "fgame/fgame/game/dan/player"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DAN_ALCHEMYSTART_TYPE), dispatch.HandlerFunc(handleAlchemyStart))
}

//处理开始炼丹信息
func handleAlchemyStart(s session.Session, msg interface{}) (err error) {
	log.Debug("dan:处理开始炼丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAlchemyStart := msg.(*uipb.CSAlchemyStart)
	kindId := csAlchemyStart.GetKindId()
	num := csAlchemyStart.GetNum()

	err = alchemyStart(tpl, int(kindId), num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"kindId":   kindId,
				"num":      num,
				"error":    err,
			}).Error("dan:处理开始炼丹信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"kindId":   kindId,
			"num":      num,
		}).Debug("dan:处理开始炼丹信息完成")
	return nil
}

//处理开始炼丹的逻辑
func alchemyStart(pl player.Player, kindId int, num int32) (err error) {
	danManager := pl.GetPlayerDataManager(types.PlayerDanDataManagerType).(*playerdan.PlayerDanDataManager)
	//校验数据有效性
	alchemyTemplate := dan.GetDanService().GetAlchemy(kindId)
	if alchemyTemplate == nil || num < 1 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"kindId":   kindId,
			"num":      num,
		}).Warn("dan:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag := danManager.IsCanAlchemy(kindId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"kindId":   kindId,
			"num":      num,
		}).Warn("dan:当前正在炼丹或有丹药未领取,无法炼丹")
		playerlogic.SendSystemMessage(pl, lang.DanStillNotGet)
		return
	}

	//判断炼丹材料是否够
	items := make(map[int32]int32)
	for itemId, itemNum := range alchemyTemplate.GetAllAlchemy() {
		items[int32(itemId)] = itemNum * num
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(items) != 0 {
		flag := inventoryManager.HasEnoughItems(items)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"kindId":   kindId,
				"num":      num,
			}).Warn("dan:当前材料不足，无法炼丹")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//炼丹使用物品
	inventoryReason := commonlog.InventoryLogReasonAchemyStart
	reasonText := fmt.Sprintf(inventoryReason.String(), alchemyTemplate.SynthetiseId)
	if len(items) != 0 {
		flag := inventoryManager.BatchRemove(items, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("dan: alchemyStart use item should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	danManager.SetAlchemyStart(kindId, num)
	alchemyList := danManager.GetAlchemyInfo()
	scAlchemyStart := pbuitl.BuildSCAlchemyStart(alchemyList, kindId)
	pl.SendMsg(scAlchemyStart)
	return
}
