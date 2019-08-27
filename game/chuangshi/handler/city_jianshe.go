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
// 	"fgame/fgame/game/item/item"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	playertypes "fgame/fgame/game/player/types"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"
// 	"fmt"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_CITY_JIANSHE_TYPE), dispatch.HandlerFunc(handleChuangShiCityJianShe))
// }

// //处理创世城池建设
// func handleChuangShiCityJianShe(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理创世城池建设")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	csMsg := msg.(*uipb.CSChuangShiCityJianShe)

// 	// campTypeInt := csMsg.GetOrignalCampType()
// 	// cityTypeInt := csMsg.GetCityType()
// 	// index := csMsg.GetIndex()
// 	cityId := csMsg.GetCityId()
// 	buildTypeInt := csMsg.GetBuildType()
// 	num := csMsg.GetNum()

// 	// campType := chuangshitypes.ChuangShiCampType(campTypeInt)
// 	// if !campType.Valid() {
// 	// 	log.WithFields(
// 	// 		log.Fields{
// 	// 			"playerId": pl.GetId(),
// 	// 			"campType": campType,
// 	// 		}).Warnln("chuangshi:处理创世城池建设,阵营类型参数错误")
// 	// 	playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
// 	// 	return
// 	// }

// 	// cityType := chuangshitypes.ChuangShiCityType(cityTypeInt)
// 	// if !cityType.Valid() {
// 	// 	log.WithFields(
// 	// 		log.Fields{
// 	// 			"playerId": pl.GetId(),
// 	// 			"cityType": cityType,
// 	// 		}).Warnln("chuangshi:处理创世城池建设,城市类型参数错误")
// 	// 	playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
// 	// 	return
// 	// }

// 	buildType := chuangshitypes.ChuangShiCityJianSheType(buildTypeInt)
// 	if !buildType.Valid() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":  pl.GetId(),
// 				"buildType": buildType,
// 			}).Warnln("chuangshi:处理创世城池建设,建设类型参数错误")
// 		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	if num < 0 {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"num":      num,
// 			}).Warnln("chuangshi:处理创世城池建设,建设数量错误")
// 		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	err = chuangShiCityJianShe(tpl, cityId, buildType, num)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"error":    err,
// 			}).Error("chuangshi:处理创世城池建设,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理创世城池建设,成功")
// 	return nil
// }

// func chuangShiCityJianShe(pl player.Player, cityId int64, buildType chuangshitypes.ChuangShiCityJianSheType, num int32) (err error) {
// 	jianSheObj := chuangshi.GetChuangShiService().GetChengFangJianShe(pl.GetId())
// 	if jianSheObj != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":  pl.GetId(),
// 				"buildType": buildType,
// 			}).Warnln("chuangshi:处理创世城池建设,城防建设正在建设")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiJianSheProgressing)
// 		return
// 	}

// 	chengFangTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiChengFangTemp(buildType)
// 	if chengFangTemp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":  pl.GetId(),
// 				"buildType": buildType,
// 			}).Warnln("chuangshi:处理创世城池建设,城防建设模板不存在")
// 		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
// 		return
// 	}

// 	camp := chuangshi.GetChuangShiService().GetCamp(pl.GetId())
// 	if camp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世城池建设,没有阵营")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotCamp)
// 		return
// 	}

// 	city := camp.GetCityById(cityId)
// 	if city == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"cityId":   cityId,
// 			}).Warnln("chuangshi:处理创世城池建设,城池不存在")
// 		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	jianShe := city.GetJianShe(buildType)
// 	jianSheLevelTemp := chengFangTemp.GetJianSheLevelTemp(jianShe.Level)
// 	if jianSheLevelTemp == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":      pl.GetId(),
// 				"cityId":        cityId,
// 				"buildType":     buildType,
// 				"jianShe.Level": jianShe.Level,
// 			}).Warnln("chuangshi:处理创世城池建设,城防建设等级模板不存在")
// 		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
// 		return
// 	}

// 	if jianSheLevelTemp.GetNextTemp() == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":      pl.GetId(),
// 				"buildType":     buildType,
// 				"jianShe.Level": jianShe.Level,
// 			}).Warnln("chuangshi:处理创世城池建设,城防建设已经满级")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiJianSheFullLevel)
// 		return
// 	}

// 	//物品是否足够
// 	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 	needItemId := chengFangTemp.LevelItemId
// 	if !inventoryManager.HasEnoughItem(needItemId, num) {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":   pl.GetId(),
// 				"needItemId": needItemId,
// 				"num":        num,
// 			}).Warnln("chuangshi:处理创世城池建设,物品不足")
// 		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 		return
// 	}

// 	// 建设
// 	chuangshi.GetChuangShiService().ChengFangJianShe(pl.GetId(), cityId, buildType, num)

// 	//扣物品
// 	itemReason := commonlog.InventoryLogReasonChuangShiChengFangJianShe
// 	itemReasonText := fmt.Sprintf(itemReason.String(), num)
// 	flag := inventoryManager.UseItem(needItemId, num, itemReason, itemReasonText)
// 	if !flag {
// 		panic(fmt.Errorf("chuangshi:城池建设消耗物品应该成功"))
// 	}

// 	//加积分
// 	chuangshiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	itemTemp := item.GetItemService().GetItem(int(needItemId))
// 	addJifen := itemTemp.TypeFlag2 * num
// 	if addJifen > 0 {
// 		flag := chuangshiManager.AddJiFen(addJifen)
// 		if !flag {
// 			panic(fmt.Errorf("chuangshi:城池建设获取创世积分应该成功"))
// 		}
// 	}

// 	//同步背包
// 	inventorylogic.SnapInventoryChanged(pl)

// 	scMsg := pbutil.BuildSCChuangShiCityJianShe()
// 	pl.SendMsg(scMsg)
// 	return
// }
