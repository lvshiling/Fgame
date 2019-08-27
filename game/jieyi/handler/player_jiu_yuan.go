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
	"fgame/fgame/game/jieyi/pbutil"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_JIU_YUAN_TYPE), dispatch.HandlerFunc(handlePlayerJiuYuan))
}

func handlePlayerJiuYuan(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理玩家救援请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	p := gcs.Player()
	pl := p.(player.Player)

	csMsg := msg.(*uipb.CSJieYiJiuYuan)
	mapId := csMsg.GetMapId()
	pos := csMsg.GetPos()
	position := coretypes.Position{
		X: float64(pos.GetPosX()),
		Y: float64(pos.GetPosY()),
		Z: float64(pos.GetPosZ()),
	}

	err = playerJiuYuan(pl, mapId, position)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
				"err":      err,
			}).Error("jieyi: 处理玩家救援请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("jieyi: 处理玩家救援请求消息,成功")

	return
}

func playerJiuYuan(pl player.Player, mapId int32, pos coretypes.Position) (err error) {
	s := scene.GetSceneService().GetSceneByMapId(mapId)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("jieyi: 场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	if s.MapTemplate().IsChuansong == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("jieyi: 该地图不支持传送")
		playerlogic.SendSystemMessage(pl, lang.JieYiMemberNotChuanSong)
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	huiyuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItemId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeZhiFeiFu))
	needItemNum := s.MapTemplate().FeixieCount
	isCostItem := huiyuanManager.GetHuiYuanType() == huiyuantypes.HuiYuanTypeCommon && needItemNum > 0
	if isCostItem {
		if !inventoryManager.HasEnoughItem(needItemId, needItemNum) {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"needItemId":  needItemId,
					"needItemNum": needItemNum,
				}).Warn("jieyi: 结义救援，传送物品不足")
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
		useReason := commonlog.InventoryLogReasonJieYiJiuYuanUse
		flag := inventoryManager.UseItem(needItemId, needItemNum, useReason, useReason.String())
		if !flag {
			panic(fmt.Errorf("jieyi:结义救援，消耗物品应该成功"))
		}

		inventorylogic.SnapInventoryChanged(pl)
	}

	scMsg := pbutil.BuildSCJieYiJiuYuan()
	pl.SendMsg(scMsg)

	return
}
