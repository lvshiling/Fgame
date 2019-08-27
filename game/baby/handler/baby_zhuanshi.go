package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/core/utils"
	babylogic "fgame/fgame/game/baby/logic"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	"fgame/fgame/game/common/common"
	emaillogic "fgame/fgame/game/email/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_ZHUAN_SHI_TYPE), dispatch.HandlerFunc(handleBabyZhuanShi))
}

//处理宝宝转世
func handleBabyZhuanShi(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理宝宝转世消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSBabyZhuanShi)
	babyId := csMsg.GetBabyId()

	err = babyZhuanShi(tpl, babyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baby:处理宝宝转世消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baby:处理宝宝转世消息完成")
	return nil

}

//宝宝转世界面逻辑
func babyZhuanShi(pl player.Player, babyId int64) (err error) {

	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	baby := babyManager.GetBabyInfo(babyId)
	if baby == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"babyId":   babyId,
			}).Warn("baby:处理宝宝转世消息,宝宝不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//读书返还
	returnItemMap := map[int32]int32{}
	babyLearnTemp := babytemplate.GetBabyTemplateService().GetBabyLearnTemplate(baby.GetLearnLevel())
	if babyLearnTemp != nil {
		returnItemMap = utils.MergeMap(returnItemMap, babyLearnTemp.GetReturnItemMap())
	}

	//天赋洗练返还
	if baby.GetRefreshCostNum() > 0 {
		babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
		returnNum := math.Ceil(float64(baby.GetRefreshCostNum()) * float64(babyConstantTemplate.ZsTianFuReturnRate) / float64(common.MAX_RATE))
		returnItemMap[babyConstantTemplate.XiLianItemId] += int32(returnNum)
	}

	//宝宝转世
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlots(returnItemMap) {
		title := lang.GetLangService().ReadLang(lang.EmailInventorySlotNoEnough)
		content := lang.GetLangService().ReadLang(lang.BabyFailReturnMailContent)
		emaillogic.AddEmail(pl, title, content, returnItemMap)
	} else {
		itemGetReason := commonlog.InventoryLogReasonBabyZhuanShiReturn
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), baby.GetLearnLevel(), baby.GetRefreshCostNum())
		flag := inventoryManager.BatchAdd(returnItemMap, itemGetReason, itemGetReasonText)
		if !flag {
			panic("baby:宝宝转世失败返还物品添加应该成功")
		}
	}
	babyManager.BabyZhuanShi(babyId)
	inventorylogic.SnapInventoryChanged(pl)
	babylogic.BabyPropertyChanged(pl)

	scMsg := pbutil.BuildSCBabyZhuanShi(babyId, returnItemMap, babyManager.GetBabyNum())
	pl.SendMsg(scMsg)
	return
}
