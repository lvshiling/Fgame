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
	processor.Register(codec.MessageType(uipb.MessageType_CS_JUEXUE_INSIGHT_TYPE), dispatch.HandlerFunc(handleJueXueInsight))
}

//处理绝学顿悟信息
func handleJueXueInsight(s session.Session, msg interface{}) (err error) {
	log.Debug("juexue:处理绝学顿悟信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csJueXueInsight := msg.(*uipb.CSJueXueInsight)
	typ := csJueXueInsight.GetTyp()

	err = juexueInsight(tpl, jxtypes.JueXueType(typ))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Error("juexue:处理绝学顿悟信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("juexue:处理绝学顿悟完成")
	return nil
}

//绝学顿悟的逻辑
func juexueInsight(pl player.Player, typ jxtypes.JueXueType) (err error) {
	jxManager := pl.GetPlayerDataManager(types.PlayerJueXueDataManagerType).(*playerjx.PlayerJueXueDataManager)
	flag := jxManager.IfJueXueExist(typ)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:未激活的绝学,无法顿悟")
		playerlogic.SendSystemMessage(pl, lang.JueXueNotActiveNotInsight)
		return
	}

	flag = jxManager.IfCanInsight(typ)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:条件未达成,无法顿悟")
		playerlogic.SendSystemMessage(pl, lang.JueXueInsightNotReach)
		return
	}

	state, level := jxManager.GetJueXueLevelByTyp(typ)
	curLevel := level
	if state == jxtypes.JueXueStageTypeAorU {
		curLevel = 0
	}

	jueXueTemplate := juexue.GetJueXueService().GetJueXueByTypeAndLevel(typ, jxtypes.JueXueStageTypeInsight, curLevel+1)
	if jueXueTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:顿悟达到最高级")
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
			}).Warn("juexue:道具不足，无法顿悟")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonJueXueUpgrade.String()
		flag = inventoryManager.UseItem(needItem, needCount, commonlog.InventoryLogReasonJueXueUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("juexue: juexueInsight use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = jxManager.Insight(typ)
	if !flag {
		panic(fmt.Errorf("juexue: Insight should be ok"))
	}
	juexueId := int32(jueXueTemplate.TemplateId())
	scJueXueInsight := pbutil.BuildSCJueXueInsight(juexueId)
	pl.SendMsg(scJueXueInsight)
	return
}
