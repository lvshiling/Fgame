package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/tower/pbutil"
	playertower "fgame/fgame/game/tower/player"
	towertemplate "fgame/fgame/game/tower/template"
	towertypes "fgame/fgame/game/tower/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func HandleOperateTower(pl player.Player, operateType towertypes.TowerOperationType) (err error) {

	towerManager := pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)
	switch operateType {
	case towertypes.TowerOperationTypeBegin:
		{
			remainTime := towerManager.GetRemainTime()
			if remainTime <= 0 {
				log.WithFields(
					log.Fields{
						"playerId":    pl.GetId(),
						"operateType": operateType,
					}).Warn("chess:处理获取打宝塔操作请求,没有打宝时间")
				playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
				return
			}
			towerManager.StartDaBao()

			// 打宝光效
			buffId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDaBaoBuff))
			scenelogic.AddBuff(pl, buffId, pl.GetId(), common.MAX_RATE)
		}
	case towertypes.TowerOperationTypeEnd:
		{
			towerManager.EndDaBao()

			// 移除光效
			buffId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDaBaoBuff))
			scenelogic.RemoveBuff(pl, buffId)
		}
	}

	remainTime := towerManager.GetRemainTime()
	scMsg := pbutil.BuildSCTowerDaBao(remainTime)
	pl.SendMsg(scMsg)
	return
}

const (
	defaultFloor = int32(1)
)

func CheckIfCanEnterTower(pl player.Player, floor int32) (flag bool) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	towerTemplate := towertemplate.GetTowerTemplateService().GetTowerTemplate(floor)
	if towerTemplate == nil {
		return false
	}

	// 使用道具
	itemId := towerTemplate.ZhifeiItem
	num := towerTemplate.ZhifeiItemCount
	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if floor != defaultFloor && huiyuanManager.GetHuiYuanType() == huiyuantypes.HuiYuanTypeCommon {
		if !inventoryManager.HasEnoughItem(itemId, num) {
			return false
		}
	}

	return true
}

func HandleEnterTower(pl player.Player, floor int32) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	towerTemplate := towertemplate.GetTowerTemplateService().GetTowerTemplate(floor)
	if towerTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"floor":    floor,
			}).Warn("tower:处理进入打宝塔消息,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// 使用道具
	itemId := towerTemplate.ZhifeiItem
	num := towerTemplate.ZhifeiItemCount
	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if floor != defaultFloor && huiyuanManager.GetHuiYuanType() == huiyuantypes.HuiYuanTypeCommon {
		if !inventoryManager.HasEnoughItem(itemId, num) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"floor":    floor,
				}).Warn("tower:处理进入打宝塔消息,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		itemUseReason := commonlog.InventoryLogReasonEnterTower
		flag := inventoryManager.UseItem(itemId, num, itemUseReason, itemUseReason.String())
		if !flag {
			panic(fmt.Errorf("tower: 进入打宝塔消耗物品应该成功"))
		}
		// 同步
		inventorylogic.SnapInventoryChanged(pl)
	}

	s := scene.GetSceneService().GetTowerSceneByMapId(towerTemplate.MapId)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"floor":    floor,
			}).Warn("tower:处理进入打宝塔消息,场景不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
	}
	pos := s.MapTemplate().GetBornPos()
	scenelogic.PlayerEnterScene(pl, s, pos)
	return
}
