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
	playerjuexue "fgame/fgame/game/juexue/player"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_JUEXUE_ACTIVE_TYPE), dispatch.HandlerFunc(handleJueXueActive))
}

//处理绝学激活信息
func handleJueXueActive(s session.Session, msg interface{}) (err error) {
	log.Debug("juexue:处理绝学激活信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csJueXueActive := msg.(*uipb.CSJueXueActive)
	typ := csJueXueActive.GetTyp()

	err = juexueActive(tpl, jxtypes.JueXueType(typ))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Error("juexue:处理绝学激活信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Debug("juexue:处理绝学激活信息完成")
	return nil
}

//处理绝学激活信息逻辑
func juexueActive(pl player.Player, typ jxtypes.JueXueType) (err error) {
	jxManager := pl.GetPlayerDataManager(types.PlayerJueXueDataManagerType).(*playerjuexue.PlayerJueXueDataManager)
	if !typ.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag := jxManager.IfJueXueExist(typ)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:该绝学已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.JueXueRepeatActive)
		return
	}

	jueXueTemplate := juexue.GetJueXueService().GetJueXueByTypeAndLevel(typ, jxtypes.JueXueStageTypeAorU, 1)
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
			}).Warn("juexue:道具不足，无法激活")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		inventoryReason := commonlog.InventoryLogReasonJueXueActive
		reasonText := fmt.Sprintf(inventoryReason.String(), typ)
		flag = inventoryManager.UseItem(needItem, needCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("juexue: juexueActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = jxManager.JueXueActive(typ)
	if !flag {
		panic(fmt.Errorf("juexue: JueXueActive should be ok"))
	}

	juexueId := int32(jueXueTemplate.TemplateId())
	scJueXueActive := pbutil.BuildSCJueXueActive(juexueId)
	pl.SendMsg(scJueXueActive)
	return
}
