package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	commonlog "fgame/fgame/common/log"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/game/chuangshi/chuangshi"
// 	"fgame/fgame/game/chuangshi/pbutil"
// 	playerchuangshi "fgame/fgame/game/chuangshi/player"
// 	chuangshitemplate "fgame/fgame/game/chuangshi/template"
// 	chuangshitypes "fgame/fgame/game/chuangshi/types"
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
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_TIANQI_SET_TYPE), dispatch.HandlerFunc(handleChuangShiCityTianQiSet))
// }

// //处理创世城池天气设置
// func handleChuangShiCityTianQiSet(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理创世城池天气设置")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	csMsg := msg.(*uipb.CSChuangShiCityTianQiSet)
// 	cityId := csMsg.GetCityId()
// 	level := csMsg.GetLevel()

// 	err = chuangShiCityTianQiSet(tpl, cityId, level)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理创世城池天气设置,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理创世城池天气设置,成功")
// 	return nil
// }

// func chuangShiCityTianQiSet(pl player.Player, cityId int64, level int32) (err error) {
// 	chuangshiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangshiInfo := chuangshiManager.GetPlayerChuangShiInfo()

// 	camp := chuangshi.GetChuangShiService().GetCamp(pl.GetId())
// 	if camp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世城池天气设置,没有阵营")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotCamp)
// 		return
// 	}

// 	city := camp.GetCityById(cityId)
// 	if city == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"cityId":   cityId,
// 			}).Warnln("chuangshi:处理创世城池天气设置,城池不存在")
// 		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	// 官职判断
// 	if !chuangshiInfo.IfShenWang() && !city.IfChengZhu(pl.GetId()) {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"cityId":   cityId,
// 			}).Warnln("chuangshi:处理创世城池天气设置,不是神王或当前城主")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiJianSheSkillSetFailed)
// 		return
// 	}

// 	chengFangTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiChengFangTemp(chuangshitypes.ChuangShiCityJianSheTypeTianQi)
// 	if chengFangTemp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世城池天气设置,城防天气设置模板不存在")
// 		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
// 		return
// 	}

// 	jianSheLevelTemp := chengFangTemp.GetJianSheLevelTemp(level)
// 	if jianSheLevelTemp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"cityId":   cityId,
// 				"Level":    level,
// 			}).Warnln("chuangshi:处理创世城池天气设置,城防天气设置建筑等级模板不存在")
// 		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
// 		return
// 	}

// 	tianQiTemp := jianSheLevelTemp.GetTianQiTemp()
// 	if tianQiTemp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世城池天气设置,城防天气设置天气模板不存在")
// 		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
// 		return
// 	}

// 	// 未激活
// 	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 	needItemMap := tianQiTemp.GetActivateItemMap()
// 	isActivate := city.IfActivateJianSheSkill(level)
// 	if !isActivate {
// 		//物品是否足够
// 		if len(needItemMap) > 0 {
// 			if !inventoryManager.HasEnoughItems(needItemMap) {
// 				log.WithFields(
// 					log.Fields{
// 						"playerId":    pl.GetId(),
// 						"needItemMap": needItemMap,
// 					}).Warnln("chuangshi:处理创世城池天气设置,物品不足")
// 				playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 				return
// 			}
// 		}
// 	}

// 	// 天气设置
// 	success := chuangshi.GetChuangShiService().CityTianQiSet(pl.GetId(), cityId, level)
// 	if !success {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"cityId":   cityId,
// 				"level":    level,
// 			}).Warnln("chuangshi:处理创世城池天气设置,设置天气失败")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiJianSheSkillSetFailed)
// 		return
// 	}

// 	if !isActivate {
// 		//扣物品
// 		if len(needItemMap) > 0 {
// 			itemReason := commonlog.InventoryLogReasonChuangShiJianSheActivateSkillUse
// 			itemReasonText := fmt.Sprintf(itemReason.String(), level)
// 			flag := inventoryManager.BatchRemove(needItemMap, itemReason, itemReasonText)
// 			if !flag {
// 				panic(fmt.Errorf("chuangshi:城池天气设置消耗物品应该成功"))
// 			}
// 			//同步背包
// 			inventorylogic.SnapInventoryChanged(pl)
// 		}
// 	}

// 	scMsg := pbutil.BuildSCChuangShiCityTianQiSet(cityId, level, int32(city.OrignalCamp), int32(city.CityType), city.Index)
// 	pl.SendMsg(scMsg)
// 	return
// }
