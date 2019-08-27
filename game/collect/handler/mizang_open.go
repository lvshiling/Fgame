package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	collectnpc "fgame/fgame/game/collect/npc"
	"fgame/fgame/game/collect/pbutil"
	collecttypes "fgame/fgame/game/collect/types"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_MIZANG_OPEN_TYPE), dispatch.HandlerFunc(handleSceneMiZangOpen))
}

//处理采集密藏
func handleSceneMiZangOpen(s session.Session, msg interface{}) (err error) {
	log.Debug("collect:处理采集密藏")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSSceneCollectMiZangOpen)
	npcId := csMsg.GetNpcId()
	openType := collecttypes.MiZangOpenType(csMsg.GetType())

	if !openType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
				"openType": openType,
			}).Warn("collect:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = miZangOpen(tpl, openType, npcId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
				"openType": openType,
				"error":    err,
			}).Error("collect:处理采集密藏,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
			"openType": openType,
		}).Info("collect:处理采集密藏完成")
	return nil
}

//处理采集密藏逻辑
func miZangOpen(pl player.Player, openType collecttypes.MiZangOpenType, npcId int64) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("collect:处理采集消息,场景为空")
		return
	}

	// npc是否存在
	so := s.GetSceneObject(npcId)
	if so == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,生物不存在")
		playerlogic.SendSystemMessage(pl, lang.CollectMiZangDisappear)
		return
	}

	n, ok := so.(*collectnpc.CollecMiZangNPC)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,不是密藏")
		playerlogic.SendSystemMessage(pl, lang.CollectNotCollectNPC)
		return
	}

	//是否采集完成
	if !n.IfMiZangCanCollect(pl) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,没有采集密藏")
		playerlogic.SendSystemMessage(pl, lang.CollectMiZangNotCollect)
		return
	}

	miZangTemplate := n.GetMiZangTemplate()
	miZangOpen := miZangTemplate.GetMiZang(openType)
	if miZangOpen == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,没有采集密藏")
		playerlogic.SendSystemMessage(pl, lang.CollectMiZangOpenTypeWrong)
		return
	} 
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItemMap := miZangOpen.GetItemMap()
	if !inventoryManager.HasEnoughItems(needItemMap) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}
	reason := commonlog.InventoryLogReasonMiZangCost
	reasonText := fmt.Sprintf(reason.String(), openType.String())
	flag := inventoryManager.BatchRemove(needItemMap, reason, reasonText)
	if !flag {
		panic(fmt.Errorf("collect:物品移除应该成功"))
	}

	dropList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(miZangOpen.GetDropList())
	if len(dropList) != 0 {
		if !inventoryManager.HasEnoughSlotsOfItemLevel(dropList) {
			title := lang.GetLangService().ReadLang(lang.CollectMiZangTitle)
			content := lang.GetLangService().ReadLang(lang.CollectMiZangContent)
			now := global.GetGame().GetTimeService().Now()
			emaillogic.AddEmailItemLevel(pl, title, content, now, dropList)
		} else {
			reason := commonlog.InventoryLogReasonMiZangGet
			reasonText := fmt.Sprintf(reason.String(), openType.String())
			inventoryManager.BatchAddOfItemLevel(dropList, reason, reasonText)
		}
	}

	//打开密藏
	flag = n.MiZangCollectFinish(pl)
	if !flag {
		panic("mizang:打开密藏应该成功")
	}

	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCSceneCollectMiZangOpen(int32(openType), npcId, dropList)
	pl.SendMsg(scMsg)
	return
}
