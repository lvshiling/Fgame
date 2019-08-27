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
	"fgame/fgame/game/juexue/juexue"
	"fgame/fgame/game/juexue/pbutil"
	playerjx "fgame/fgame/game/juexue/player"
	jxtypes "fgame/fgame/game/juexue/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JUEXUE_UPGRADE_TYPE), dispatch.HandlerFunc(handleJueXueUpgrade))
}

//处理绝学升级信息
func handleJueXueUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("juexue:处理绝学升级信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csJueXueUpgrade := msg.(*uipb.CSJueXueUpgrade)
	typ := csJueXueUpgrade.GetTyp()

	err = juexueUpgrade(tpl, jxtypes.JueXueType(typ))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Error("juexue:处理绝学升级信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("juexue:处理绝学升级完成")
	return nil
}

//绝学升级的逻辑
func juexueUpgrade(pl player.Player, typ jxtypes.JueXueType) (err error) {
	jxManager := pl.GetPlayerDataManager(types.PlayerJueXueDataManagerType).(*playerjx.PlayerJueXueDataManager)
	flag := jxManager.IfJueXueExist(typ)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:未激活的绝学,无法升级")
		playerlogic.SendSystemMessage(pl, lang.JueXueNotActiveNotUpgrade)
		return
	}

	flag = jxManager.IfCanUpgrade(typ)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:绝学已达最高级")
		playerlogic.SendSystemMessage(pl, lang.JueXueReacheFullUpgrade)
		return
	}

	_, curLevel := jxManager.GetJueXueLevelByTyp(typ)
	jueXueTemplate := juexue.GetJueXueService().GetJueXueByTypeAndLevel(typ, jxtypes.JueXueStageTypeAorU, curLevel+1)
	if jueXueTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:绝学强化已达到最高级")
		playerlogic.SendSystemMessage(pl, lang.JueXueReacheFullUpgrade)
		return
	}
	useItem := jueXueTemplate.GetUseItempTemplate()
	if useItem != nil {
		needItem := jueXueTemplate.NeedItemId
		needCount := jueXueTemplate.NeedItemNum
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItem(needItem, needCount)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("juexue:道具不足，无法升级")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonJueXueUpgrade.String()
		flag = inventoryManager.UseItem(needItem, needCount, commonlog.InventoryLogReasonJueXueUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("juexue: juexueUpgrade use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = jxManager.Upgrade(typ)
	if !flag {
		panic(fmt.Errorf("juexue: juexueUpgrade should be ok"))
	}
	juexueId := int32(jueXueTemplate.TemplateId())
	scJueXueUpgrade := pbutil.BuildSCJueXueUpgrade(juexueId)
	pl.SendMsg(scJueXueUpgrade)
	return
}
