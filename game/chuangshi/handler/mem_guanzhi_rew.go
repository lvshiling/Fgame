package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	commonlog "fgame/fgame/common/log"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/game/chuangshi/pbutil"
// 	playerchuangshi "fgame/fgame/game/chuangshi/player"
// 	chuangshitemplate "fgame/fgame/game/chuangshi/template"
// 	inventorylogic "fgame/fgame/game/inventory/logic"
// 	playerinventory "fgame/fgame/game/inventory/player"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	playertypes "fgame/fgame/game/player/types"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"
// 	"fmt"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_GUANZHI_REW_TYPE), dispatch.HandlerFunc(handleGuanZhiRew))
// }

// func handleGuanZhiRew(s session.Session, msg interface{}) (err error) {
// 	log.Debug("开始处理创世之战领取官职奖励消息")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)
// 	csMsg := msg.(*uipb.CSChuangShiGuanZhiRew)
// 	level := csMsg.GetRewLevel()

// 	err = guanZhiRewGet(tpl, level)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": tpl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理创世之战领取官职奖励消息,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": tpl.GetId(),
// 		}).Debug("chuangshi:处理创世之战领取官职奖励消息,成功")

// 	return
// }

// func guanZhiRewGet(pl player.Player, rewLevel int32) (err error) {

// 	guanZhiTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiGuanZhiTemplate(rewLevel)
// 	if guanZhiTemp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"rewLevel": rewLevel,
// 			}).Warnln("chuangshi:处理创世之战升职消息,模板不存在")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiTemplateNotExist)
// 		return
// 	}

// 	chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	info := chuangShiManager.GetPlayerChuangShiGuanZhiInfo()
// 	if !info.IsCanReceive(rewLevel) {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"rewLevel": rewLevel,
// 			}).Warnln("chuangshi:处理创世之战升职消息,已经领取过官职奖励")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiAlreadyReceive)
// 		return
// 	}

// 	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

// 	// 获得物品
// 	getItemMap := guanZhiTemp.GetReceiveItemMap()
// 	if len(getItemMap) != 0 {
// 		if !inventoryManager.HasEnoughSlots(getItemMap) {
// 			log.WithFields(
// 				log.Fields{
// 					"playerId": pl.GetId(),
// 					"rewLevel": rewLevel,
// 				}).Warnln("chuangshi:处理创世之战升职消息,背包空间不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnoughSlot)
// 			return
// 		}

// 		Reason := commonlog.InventoryLogReasonChuangShiGuanZhiGet
// 		ReasonText := fmt.Sprintf(Reason.String(), rewLevel)
// 		flag := inventoryManager.BatchAdd(getItemMap, Reason, ReasonText)
// 		if !flag {
// 			panic("chuangshi: 官职奖励添加物品应该成功")
// 		}

// 		// 推送变化
// 		inventorylogic.SnapInventoryChanged(pl)
// 	}

// 	flag := chuangShiManager.ReceiveGuanZhiRew(rewLevel)
// 	if flag {
// 		panic(fmt.Errorf("chuangshi:领取官职奖励应该成功，rewLevel：%d", rewLevel))
// 	}

// 	scMsg := pbutil.BuildSCChuangShiGuanZhiRew(rewLevel)
// 	pl.SendMsg(scMsg)
// 	return
// }
