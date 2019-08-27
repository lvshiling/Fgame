package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/gem/gem"
	"fgame/fgame/game/gem/pbutil"
	playergem "fgame/fgame/game/gem/player"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_GEM_MINE_ACTIVE_TYPE), dispatch.HandlerFunc(handleGemMineActive))
}

//处理矿工激活信息
func handleGemMineActive(s session.Session, msg interface{}) (err error) {
	log.Debug("gem:处理矿工激活信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csGemMineActive := msg.(*uipb.CSGemMineActive)
	level := csGemMineActive.GetLevel()

	err = gemMineActive(tpl, level)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"level":    level,
				"error":    err,
			}).Error("gem:处理矿工激活信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
		}).Debug("gem:处理矿工激活信息完成")
	return nil
}

//处理矿工激活信息逻辑
func gemMineActive(pl player.Player, level int32) (err error) {
	gemManager := pl.GetPlayerDataManager(types.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)
	flag := gemManager.IsMineValid(level)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
		}).Warn("gem:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	mineTemplate := gem.GetGemService().GetMineTemplateByLevel(level)
	needYinLiang := int64(mineTemplate.NeedYinLiang)
	needGold := int64(mineTemplate.NeedGold)
	useItem := mineTemplate.GetUseItemTemplate()

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//判断银两
	if needYinLiang != 0 {
		flag = propertyManager.HasEnoughSilver(needYinLiang)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"level":    level,
			}).Warn("gem:银两不足,无法激活")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//判断元宝
	if needGold != 0 {
		flag = propertyManager.HasEnoughGold(int64(needGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"level":    level,
			}).Warn("gem:元宝不足，无法激活")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//判断物品
	needItem := int32(0)
	needNum := int32(0)
	if useItem != nil {
		needItem = mineTemplate.ItemId1
		needNum = mineTemplate.ItemCount1
	}
	flag = inventoryManager.HasEnoughItem(needItem, needNum)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
		}).Warn("gem:道具不足，无法激活")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//消耗银两
	if needYinLiang != 0 {
		reasonText := commonlog.SilverLogReasonGemMineActive.String()
		flag = propertyManager.CostSilver(needYinLiang, commonlog.SilverLogReasonGemMineActive, reasonText)
		if !flag {
			panic(fmt.Errorf("gem: gemMineActive CostSilver  should be ok"))
		}
	}

	//消耗元宝
	if needGold != 0 {
		reasonText := commonlog.GoldLogReasonGemMineActive.String()
		flag = propertyManager.CostGold(needGold, false, commonlog.GoldLogReasonGemMineActive, reasonText)
		if !flag {
			panic(fmt.Errorf("gem: gemMineActive CostGold  should be ok"))
		}
	}
	//同步银两、元宝
	if needGold != 0 || needYinLiang != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//消耗物品
	if needItem != 0 {
		reasonText := commonlog.InventoryLogReasonGemMineActive.String()
		flag = inventoryManager.UseItem(needItem, needNum, commonlog.InventoryLogReasonGemMineActive, reasonText)
		if !flag {
			panic(fmt.Errorf("gem: gemMineActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = gemManager.MineActive(level)
	if !flag {
		panic(fmt.Errorf("gem: gemMineActive MineActive should be ok"))
	}
	mine := gemManager.GetMine()
	scGemMineActive := pbutil.BuildSCGemMineActive(mine)
	pl.SendMsg(scGemMineActive)
	return
}
