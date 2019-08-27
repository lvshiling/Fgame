package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	commonlog "fgame/fgame/common/log"
// 	"fgame/fgame/core/session"
// 	chuangshilogic "fgame/fgame/game/chuangshi/logic"
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
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_POSITION_ADVANCE_TYPE), dispatch.HandlerFunc(handlePositionAdvance))
// }

// func handlePositionAdvance(s session.Session, msg interface{}) (err error) {
// 	log.Debug("开始处理创世之战升职消息")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	err = positionAdvance(tpl)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": tpl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理创世之战升职消息,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": tpl.GetId(),
// 		}).Debug("chuangshi:处理创世之战升职消息,成功")

// 	return
// }

// func positionAdvance(pl player.Player) (err error) {
// 	chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	info := chuangShiManager.GetPlayerChuangShiGuanZhiInfo()
// 	level := info.GetLevel() + 1

// 	guanZhiTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiGuanZhiTemplate(level)
// 	if guanZhiTemp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"level":    level,
// 			}).Warnln("chuangshi:处理创世之战升职消息,已经到达最高职位")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiPositionLevelAlreadyTop)
// 		return
// 	}

// 	if info.GetWeiWang() < guanZhiTemp.UseWeiWang {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"level":    level,
// 			}).Warnln("chuangshi:处理创世之战升职消息,威望不足")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiWeiWangNotEnough)
// 		return
// 	}

// 	// 消耗物品
// 	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 	useItem := guanZhiTemp.GetUseItemMap()
// 	if len(useItem) != 0 {
// 		if !inventoryManager.HasEnoughItems(useItem) {
// 			log.WithFields(
// 				log.Fields{
// 					"playerId": pl.GetId(),
// 					"level":    level,
// 				}).Warnln("chuangshi:处理创世之战升职消息,所需物品不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}
// 	}

// 	Reason := commonlog.InventoryLogReasonChuangShiGuanZhiUse
// 	ReasonText := fmt.Sprintf(Reason.String(), level)
// 	flag := inventoryManager.BatchRemove(useItem, Reason, ReasonText)
// 	if !flag {
// 		panic("chuangshi: 消耗物品应该成功")
// 	}

// 	success := chuangshilogic.ChuangShiGuanZhiAdvance(info.GetTimes(), guanZhiTemp)

// 	// 推送变化
// 	inventorylogic.SnapInventoryChanged(pl)

// 	// 刷新数据
// 	chuangShiManager.UpdatePlayerGuanZhiData(success, guanZhiTemp.UseWeiWang)

// 	scMsg := pbutil.BuildSCChuangShiPositionAdvance(success, level, info.GetWeiWang())
// 	pl.SendMsg(scMsg)
// 	return
// }
