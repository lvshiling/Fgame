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
	playeralliance "fgame/fgame/game/alliance/player"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_SCENE_GET_REWARD_TYPE), dispatch.HandlerFunc(handleAllianceSceneGetReward))
}

//处理仙盟领取奖励
func handleAllianceSceneGetReward(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟领取奖励")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAllianceSceneGetReward := msg.(*uipb.CSAllianceSceneGetReward)
	door := csAllianceSceneGetReward.GetDoor()
	err = allianceSceneGetReward(tpl, door)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理仙盟领取奖励,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟领取奖励,完成")
	return nil

}

//处理仙盟领取奖励
func allianceSceneGetReward(pl player.Player, door int32) (err error) {
	//判断是否在仙盟
	//判断是否门已经破了
	//判断是否已经领取奖励过了
	s := pl.GetScene()
	sd := s.SceneDelegate()
	switch ssd := sd.(type) {
	case alliancescene.AllianceSceneData:
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
		currentDoor := ssd.GetCurrentDoor()
		if door >= currentDoor {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("alliance:处理仙盟领取奖励,城门还没攻破")
			playerlogic.SendSystemMessage(pl, lang.AllianceSceneDoorNotBroke)
			return
		}
		doorRewardTemplate := alliancetemplate.GetAllianceTemplateService().GetDoorRewardTemplate(door)
		if doorRewardTemplate == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("alliance:处理仙盟领取奖励,城门奖励不存在")
			playerlogic.SendSystemMessage(pl, lang.AllianceSceneDoorRewardNoExist)
			return
		}

		if !allianceManager.IsEnoughWarPoint(doorRewardTemplate.NeedJiFen) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("alliance:处理仙盟领取奖励,城战积分不足")
			playerlogic.SendSystemMessage(pl, lang.AllianceScenePointNotEnough)
			return
		}

		allianceSceneData := alliance.GetAllianceService().GetAllianceSceneData()
		allianceId := pl.GetAllianceId()
		defendAllianceId := allianceSceneData.GetFirstDefendAllianceId()
		if allianceId == defendAllianceId {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("alliance:处理仙盟领取奖励,属于城战最开始守方")
			playerlogic.SendSystemMessage(pl, lang.AllianceSceneBelongFirstDefend)
			return
		}
		if !allianceManager.IfCanGetReward(door) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("alliance:处理仙盟领取奖励,奖励已经领取")
			playerlogic.SendSystemMessage(pl, lang.AllianceSceneDoorRewardAlreadyGet)
			return
		}
		//发放奖励
		itemMap := doorRewardTemplate.GetItemMap()
		if len(itemMap) > 0 {
			if !inventoryManager.HasEnoughSlots(itemMap) {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
					}).Warn("alliance:处理仙盟领取奖励,背包不足")
				playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
				return
			}
			reason := commonlog.InventoryLogReasonAllianceSceneDoorReward
			reasonText := fmt.Sprintf(reason.String(), allianceId, door)
			inventoryManager.BatchAdd(itemMap, reason, reasonText)
		}

		//添加属性
		rewData := doorRewardTemplate.GetRewData()
		goldReason := commonlog.GoldLogReasonAllianceSceneDoorReward
		goldReasonText := fmt.Sprintf(goldReason.String(), allianceId, door)
		silverReason := commonlog.SilverLogReasonAllianceSceneDoorReward
		silverReasonText := fmt.Sprintf(silverReason.String(), allianceId, door)
		levelReason := commonlog.LevelLogReasonAllianceSceneDoorReward
		levelReasonText := fmt.Sprintf(silverReason.String(), allianceId, door)
		propertyManager.AddRewData(rewData, goldReason, goldReasonText, silverReason, silverReasonText, levelReason, levelReasonText)

		flag := allianceManager.GetReward(door)
		if !flag {
			panic(fmt.Errorf("领取奖励应该成功"))
		}
		//同步属性和背包
		propertylogic.SnapChangedProperty(pl)
		inventorylogic.SnapInventoryChanged(pl)
		//获取奖励
		scAllianceSceneGetReward := pbutil.BuildSCAllianceSceneGetReward(door, itemMap, rewData)
		pl.SendMsg(scAllianceSceneGetReward)
		break
	default:
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理仙盟领取奖励,不在城战")
		playerlogic.SendSystemMessage(pl, lang.AllianceSceneNotInAllianceScene)

		break
	}

	return
}
