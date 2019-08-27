package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_CS_WORLD_ENTER_SCENE_TYPE), dispatch.HandlerFunc(handlerWorldEnter))
}

//处理进入世界场景
func handlerWorldEnter(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理进入世界场景")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csWroldEnterScene := msg.(*scenepb.CSWorldEnterScene)
	mapId := csWroldEnterScene.GetWorldMapId()

	err = playerWorldEnterMap(tpl, mapId)
	if err != nil {
		return err
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"mapId":    mapId,
		}).Debug("scene:处理进入世界场景,完成")

	return nil
}

//玩家进入世界地图
func playerWorldEnterMap(pl player.Player, mapId int32) (err error) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入世界场景,场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}
	//pvp
	if pl.IsPvpBattle() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入世界场景,正在pvp")
		playerlogic.SendSystemMessage(pl, lang.PlayerInPVP)
		return
	}
	if mapTemplate.ReqLev > pl.GetLevel() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入世界场景,等级太低")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}
	bornPos := mapTemplate.GetBornPos()
	if !mapTemplate.IsWorld() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入世界场景,地图不是世界地图")
		playerlogic.SendSystemMessage(pl, lang.SceneNotWorldScene)
		return
	}

	s := scene.GetSceneService().GetWorldSceneByMapId(mapId)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入世界场景,场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	originS := pl.GetScene()
	if s == originS {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入世界场景,重复进入")
		playerlogic.SendSystemMessage(pl, lang.SceneRepeatEnter)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//消耗银两
	silverNum := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeChangeSceneCostSilver))
	if silverNum > 0 {
		if !propertyManager.HasEnoughSilver(silverNum) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入世界场景,银两不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//消耗道具
	itemMap := constant.GetConstantService().GetChangeSceneItems()
	if len(itemMap) > 0 {
		if !inventoryManager.HasEnoughItems(itemMap) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入世界场景,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//扣除银两
	if silverNum > 0 {
		reasonText := fmt.Sprintf(commonlog.SilverLogReasonChangeScene.String(), mapId)
		flag := propertyManager.CostSilver(silverNum, commonlog.SilverLogReasonChangeScene, reasonText)
		if !flag {
			panic(fmt.Errorf("scene:切换世界地图花费银两应该成功"))
		}
		propertylogic.SnapChangedProperty(pl)
	}
	if len(itemMap) > 0 {
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonChangeScene.String(), mapId)
		flag := inventoryManager.BatchRemove(itemMap, commonlog.InventoryLogReasonChangeScene, reasonText)
		if !flag {
			panic(fmt.Errorf("scene:切换世界地图消耗物品应该成功"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	scenelogic.PlayerEnterScene(pl, s, bornPos)
	return nil
}
