package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	alliancescene "fgame/fgame/game/alliance/scene"
	alliancetemplate "fgame/fgame/game/alliance/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_SCENE_CALL_TYPE), dispatch.HandlerFunc(handleCallGuard))
}

//处理仙盟召唤守卫
func handleCallGuard(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟召唤守卫")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAllianceSceneCall := msg.(*uipb.CSAllianceSceneCall)
	guardId := csAllianceSceneCall.GetGuardId()
	err = allianceCallGuard(tpl, guardId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),

				"error": err,
			}).Error("alliance:处理仙盟召唤守卫,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟召唤守卫,完成")
	return nil

}

//城战召唤守卫
func allianceCallGuard(pl player.Player, guardId int32) (err error) {
	s := pl.GetScene()
	sd := s.SceneDelegate()

	switch ssd := sd.(type) {
	case alliancescene.AllianceSceneData:
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		guardTemplate := alliancetemplate.GetAllianceTemplateService().GetGuardTemplate(guardId)
		if guardTemplate == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    s.MapId(),
					"guardId":  guardId,
				}).Warn("alliance:处理仙盟召唤守卫,守卫不存在")
			playerlogic.SendSystemMessage(pl, lang.AllianceSceneGuardNoExist)
			return
		}

		allianceId := pl.GetAllianceId()
		//判断是否是最初守方
		if allianceId != ssd.GetCurrentDefendAllianceId() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    s.MapId(),
					"guardId":  guardId,
				}).Warn("alliance:处理仙盟召唤守卫,不属于守方")
			playerlogic.SendSystemMessage(pl, lang.AllianceSceneNotBelongFirstDefend)
			return
		}

		//判断消耗
		needGold := guardTemplate.NeedGold
		if needGold > 0 {
			if !propertyManager.HasEnoughGold(int64(needGold), false) {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"mapId":    s.MapId(),
						"guardId":  guardId,
					}).Warn("alliance:处理仙盟召唤守卫,元宝不足")
				playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
				return
			}
		}
		needBindGold := guardTemplate.NeedBindGold
		needAllGold := needGold + needBindGold
		if needBindGold > 0 {
			if !propertyManager.HasEnoughGold(int64(needAllGold), true) {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"mapId":    s.MapId(),
						"guardId":  guardId,
					}).Warn("alliance:处理仙盟召唤守卫,绑元不足")
				playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
				return
			}
		}
		needSilver := int64(guardTemplate.NeedSilver)
		if needSilver > 0 {
			if !propertyManager.HasEnoughSilver(needSilver) {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"mapId":    s.MapId(),
						"guardId":  guardId,
					}).Warn("alliance:处理仙盟召唤守卫,银两不足")
				playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
				return
			}
		}
		needItemMap := guardTemplate.GetItemMap()
		if len(needItemMap) > 0 {
			if !inventoryManager.HasEnoughItems(needItemMap) {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"mapId":    s.MapId(),
						"guardId":  guardId,
					}).Warn("alliance:处理仙盟召唤守卫,物品不足")
				playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
				return
			}
		}

		flag := alliance.GetAllianceService().GetAllianceSceneData().CallGuard(guardId)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    s.MapId(),
					"guardId":  guardId,
				}).Warn("alliance:处理仙盟召唤守卫,已经召唤过")
			playerlogic.SendSystemMessage(pl, lang.AllianceSceneGuardCalled)
			return
		}
		// flag := ssd.IfGuardCanCall(guardId)
		// if !flag {
		// 	log.WithFields(
		// 		log.Fields{
		// 			"playerId": pl.GetId(),
		// 			"mapId":    s.MapId(),
		// 			"guardId":  guardId,
		// 		}).Warn("alliance:处理仙盟召唤守卫,已经召唤过")
		// 	playerlogic.SendSystemMessage(pl, lang.AllianceSceneGuardCalled)
		// 	return
		// }

		// flag = ssd.CallGuard(guardId)
		// if !flag {
		// 	panic(fmt.Errorf("召唤守卫应该成功"))
		// }

		if needGold > 0 {
			reason := commonlog.GoldLogReasonAllianceCallGuard
			reasonText := fmt.Sprintf(reason.String(), allianceId, s.MapId(), guardId)
			flag := propertyManager.CostGold(int64(needGold), false, reason, reasonText)
			if !flag {
				panic(fmt.Errorf("花费元宝应该成功"))
			}
		}

		if needBindGold > 0 {
			reason := commonlog.GoldLogReasonAllianceCallGuard
			reasonText := fmt.Sprintf(reason.String(), allianceId, s.MapId(), guardId)
			flag := propertyManager.CostGold(int64(needBindGold), true, reason, reasonText)
			if !flag {
				panic(fmt.Errorf("花费绑元元宝应该成功"))
			}
		}

		if needSilver > 0 {
			reason := commonlog.SilverLogReasonAllianceCallGuard
			reasonText := fmt.Sprintf(reason.String(), allianceId, s.MapId(), guardId)
			flag := propertyManager.CostSilver(needSilver, reason, reasonText)
			if !flag {
				panic(fmt.Errorf("花费银两应该成功"))
			}
		}

		if len(needItemMap) > 0 {
			reason := commonlog.InventoryLogReasonAllianceCallGuard
			reasonText := fmt.Sprintf(reason.String(), allianceId, s.MapId(), guardId)
			flag := inventoryManager.BatchRemove(needItemMap, reason, reasonText)
			if !flag {
				panic(fmt.Errorf("扣除物品应该成功"))
			}
		}
		//同步属性和背包
		propertylogic.SnapChangedProperty(pl)
		inventorylogic.SnapInventoryChanged(pl)
		scAllianceSceneCall := pbutil.BuildSCAllianceSceneCall(guardId)
		pl.SendMsg(scAllianceSceneCall)
		break
	default:
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理仙盟召唤守卫,不在城战")
		playerlogic.SendSystemMessage(pl, lang.AllianceSceneNotInAllianceScene)

		break
	}

	return
}
