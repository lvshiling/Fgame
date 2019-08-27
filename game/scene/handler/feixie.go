package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coretypes "fgame/fgame/core/types"
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
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_PLAYER_FEIXIE_TRANSFER_TYPE), dispatch.HandlerFunc(handlerPlayerFeiXieTransfer))
}

//玩家飞鞋传送
func handlerPlayerFeiXieTransfer(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理玩家飞鞋传送")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSPlayerFeiXieTransfer)

	mapId := csMsg.GetMapId()
	pos := csMsg.GetPos()
	position := coretypes.Position{
		X: float64(pos.GetPosX()),
		Y: float64(pos.GetPosY()),
		Z: float64(pos.GetPosZ()),
	}

	err = playerFeiXieTransfer(tpl, mapId, position)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("scene:处理玩家飞鞋传送，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("scene:处理玩家飞鞋传送完成")

	return
}

func playerFeiXieTransfer(pl player.Player, mapId int32, pos coretypes.Position) (err error) {
	//玩家跳转场景
	s := scene.GetSceneService().GetSceneByMapId(mapId)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("scene: 场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	if s.MapTemplate().IsChuansong == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("scene: 场景地图不支持传送")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoTransfer)
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	playerScene := pl.GetScene()
	mask := s.MapTemplate().GetMap().IsMask(pos.X, pos.Z)
	if !mask {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needNum := s.MapTemplate().FeixieCount
	if huiyuanManager.GetHuiYuanType() == huiyuantypes.HuiYuanTypeCommon && needNum > 0 {
		needItemId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeZhiFeiFu))
		if !inventoryManager.HasEnoughItem(needItemId, needNum) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"needNum":  needNum,
				}).Warn("scene:玩家飞鞋传送，传送物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		useReason := commonlog.InventoryLogReasonPlayerTransfer
		flag := inventoryManager.UseItem(needItemId, needNum, useReason, useReason.String())
		if !flag {
			panic(fmt.Errorf("scene:玩家飞鞋传送，消耗物品应该成功"))
		}

		inventorylogic.SnapInventoryChanged(pl)
	}

	if playerScene == s {
		scenelogic.FixPosition(pl, pos)
	} else {
		scenelogic.PlayerEnterScene(pl, s, pos)
	}

	scMsg := pbutil.BuildSCPlayerFeiXieTransfer()
	pl.SendMsg(scMsg)
	return
}
