package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_MEMBER_RESCUE_TYPE), dispatch.HandlerFunc(handlerAllianceRescue))
}

//仙盟救援
func handlerAllianceRescue(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟救援")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSAllianceMemberRescue)

	mapId := csMsg.GetMapId()
	pos := csMsg.GetPos()
	position := coretypes.Position{
		X: float64(pos.GetPosX()),
		Y: float64(pos.GetPosY()),
		Z: float64(pos.GetPosZ()),
	}

	err = allianceRescue(tpl, mapId, position)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("alliance:处理仙盟救援，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("alliance:处理仙盟救援完成")

	return
}

func allianceRescue(pl player.Player, mapId int32, pos coretypes.Position) (err error) {
	//玩家跳转场景
	s := scene.GetSceneService().GetSceneByMapId(mapId)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance: 场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	if s.MapTemplate().IsChuansong == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance: 该地图不支持传送")
		playerlogic.SendSystemMessage(pl, lang.AllianceMemberNotRescue)
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needNum := s.MapTemplate().FeixieCount
	needItemId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeZhiFeiFu))
	isCostItem := huiyuanManager.GetHuiYuanType() == huiyuantypes.HuiYuanTypeCommon && needNum > 0
	if isCostItem {
		if !inventoryManager.HasEnoughItem(needItemId, needNum) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"needNum":  needNum,
				}).Warn("alliance:仙盟救援，传送物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	playerScene := pl.GetScene()
	pos.Y = s.MapTemplate().GetMap().GetHeight(pos.X, pos.Z)
	if playerScene == s {
		scenelogic.FixPosition(pl, pos)
	} else {
		if !scenelogic.PlayerTrackEnterScene(pl, s, pos) {
			return
		}
	}

	if isCostItem {
		useReason := commonlog.InventoryLogReasonDeliverUse
		flag := inventoryManager.UseItem(needItemId, needNum, useReason, useReason.String())
		if !flag {
			panic(fmt.Errorf("alliance:仙盟救援，消耗物品应该成功"))
		}

		inventorylogic.SnapInventoryChanged(pl)
	}

	scMsg := pbutil.BuildSCAllianceMemberRescue()
	pl.SendMsg(scMsg)
	return
}
