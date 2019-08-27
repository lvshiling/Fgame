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
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_FEIXIE_TYPE), dispatch.HandlerFunc(handleQuestFeiXie))
}

//处理小飞鞋
func handleQuestFeiXie(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理小飞鞋")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csQuestFeiXie := msg.(*uipb.CSQuestFeiXie)
	questId := csQuestFeiXie.GetQuestId()
	mapId := csQuestFeiXie.GetMapId()
	pos := csQuestFeiXie.GetPos()
	x := float64(pos.GetPosX())
	y := float64(pos.GetPosY())
	z := float64(pos.GetPosZ())
	transferPos := coretypes.Position{x, y, z}
	err = questXiaoFeiXie(tpl, questId, mapId, transferPos)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"posx":     x,
				"posy":     y,
				"posz":     z,
				"error":    err,
			}).Error("quest:处理小飞鞋,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"questId":  questId,
		}).Debug("quest:处理小飞鞋,完成")
	return nil
}

func questXiaoFeiXie(pl player.Player, questId int32, mapId int32, pos coretypes.Position) (err error) {
	//模板不存在
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"mapId":    mapId,
			}).Warn("quest:处理小飞鞋,任务不存在")
		playerlogic.SendSystemMessage(pl, lang.QuestNoExist)
		return
	}
	needItemId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeZhiFeiFu)
	needNum := questTemplate.GetFeiXieNum()

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	quest := manager.GetQuestById(questId)
	if quest == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"mapId":    mapId,
			}).Warn("quest:处理小飞鞋,任务不存在")
		playerlogic.SendSystemMessage(pl, lang.QuestNoExist)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	huiYuanManager := pl.GetPlayerDataManager(types.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	isHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus)
	tempHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypeInterim)
	if isHuiYuan || tempHuiYuan {
		needNum = 0
	}
	if needNum != 0 {
		if !inventoryManager.HasEnoughItem(needItemId, needNum) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"questId":  questId,
					"mapId":    mapId,
				}).Warn("quest:处理小飞鞋,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	s := scene.GetSceneService().GetWorldSceneByMapId(mapId)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"mapId":    mapId,
			}).Warn("quest:场景地图不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}
	mapTemplate := s.MapTemplate()

	mask := mapTemplate.GetMap().IsMask(pos.X, pos.Z)
	if !mask {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"mapId":    mapId,
			}).Warn("quest:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	y := mapTemplate.GetMap().GetHeight(pos.X, pos.Z)
	pos.Y = y

	if needNum != 0 {
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonQuestQuickTransfer.String(), questTemplate.TemplateId())
		flag := inventoryManager.UseItem(needItemId, needNum, commonlog.InventoryLogReasonQuestQuickTransfer, reasonText)
		if !flag {
			panic("quest:扣除消耗物品,应该成功")
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	//传送
	curScene := pl.GetScene()
	if curScene != s {
		scenelogic.PlayerEnterScene(pl, s, pos)
	} else if pos != pl.GetPos() {
		scenelogic.FixPosition(pl, pos)
	}

	scQuestFeiXie := pbutil.BuildSCQuestFeiXie(questId, mapId, pos)
	pl.SendMsg(scQuestFeiXie)
	return
}
