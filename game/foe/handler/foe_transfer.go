package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/foe/pbutil"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FOE_TRANSFER_TYPE), dispatch.HandlerFunc(handleFoeTransfer))
}

//处理确定追踪
func handleFoeTransfer(s session.Session, msg interface{}) error {
	log.Debug("foe:处理仇人确定追踪")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFoeTransfer := msg.(*uipb.CSFoeTransfer)
	foeId := csFoeTransfer.GetFoeId()
	err := foeTransfer(tpl, foeId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"foeId":    foeId,
				"error":    err,
			}).Error("foe:处理仇人确定追踪,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("foe:处理仇人确定追踪,完成")
	return nil

}

//处理仇人确定追踪
func foeTransfer(pl player.Player, foeId int64) (err error) {
	// manager := pl.GetPlayerDataManager(playertypes.PlayerFoeDataManagerType).(*playerfoe.PlayerFoeDataManager)
	// flag := manager.IsFoe(foeId)
	// if !flag {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"foeId":    foeId,
	// 		}).Warn("foe:参数无效")
	// 	playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
	// 	return
	// }

	if pl.IsCross() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("foe:跨服地图无法直接传送,请自行前往")
		playerlogic.SendSystemMessage(pl, lang.PlayerTrackInCross)
		return
	}

	foePl := player.GetOnlinePlayerManager().GetPlayerById(foeId)
	if foePl == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"foeId":    foeId,
			}).Warn("foe:参数无效")
		playerlogic.SendSystemMessage(pl, lang.PlayerTrackNoOnline)
		return
	}

	if foePl.IsCross() {
		if foePl.GetCrossType() != crosstypes.CrossTypeWorldboss {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"foeId":    foeId,
				}).Warn("foe:跨服地图无法直接传送,请自行前往")
			playerlogic.SendSystemMessage(pl, lang.PlayerTrackInCross)
			return
		}
	}

	itemId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTrackItem)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	flag := inventoryManager.HasEnoughItem(itemId, 1)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"foeId":    foeId,
		}).Warn("foe:道具不足,无法追踪")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	if foePl.IsCross() {
		if !crosslogic.PlayerTracEnterCross(pl, foePl.GetCrossType(), crosstypes.CrossBehaviorTypeTrack, fmt.Sprintf("%d", foePl.GetId())) {
			return
		}
	} else {
		foeS := foePl.GetScene()
		if foeS == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"foeId":    foeId,
				}).Warn("foe:玩家不在场景")
			playerlogic.SendSystemMessage(pl, lang.PlayerNoInScene)
			return
		}
		if foeS.MapTemplate().IsFuBen() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"foeId":    foeId,
				}).Warn("foe:副本地图无法直接传送")
			playerlogic.SendSystemMessage(pl, lang.PlayerInFuBen)
			return
		}

		foePos := foePl.GetPos()
		if pl.GetScene() != foeS {
			if !scenelogic.PlayerTrackEnterScene(pl, foeS, foePos) {
				return
			}
		} else {
			scenelogic.FixPosition(pl, foePos)
		}
	}

	inventoryReason := commonlog.InventoryLogReasonFoeTrackUse
	reasonText := inventoryReason.String()
	flag = inventoryManager.UseItem(itemId, 1, inventoryReason, reasonText)
	if !flag {
		panic(fmt.Errorf("foe: foeViewPos use item should be ok"))
	}
	inventorylogic.SnapInventoryChanged(pl)

	scFoeTransfer := pbutil.BuildSCFoeTransfer(foeId)
	pl.SendMsg(scFoeTransfer)
	return
}
