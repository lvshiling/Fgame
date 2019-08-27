package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	"fgame/fgame/game/common/common"
	emaillogic "fgame/fgame/game/email/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_DONGFANG_TYPE), dispatch.HandlerFunc(handleBabyDongFang))
}

//处理洞房
func handleBabyDongFang(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理洞房消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSBabyDongFang)
	isUseItem := csMsg.GetIsUseItem()

	err = babyDongFang(tpl, isUseItem)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baby:处理洞房消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baby:处理洞房消息完成")
	return nil

}

//洞房界面逻辑
func babyDongFang(pl player.Player, isUseItem bool) (err error) {
	// 配偶是否在线
	marryManager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	if !marryManager.IsTrueMarry() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("baby:处理洞房消息,没有伴侣")
		playerlogic.SendSystemMessage(pl, lang.BabyPregnantNotMarry)
		return
	}

	spl := player.GetOnlinePlayerManager().GetPlayerById(marryManager.GetSpouseId())
	if spl == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("baby:处理洞房消息,伴侣不在线")
		playerlogic.SendSystemMessage(pl, lang.BabyCoupleNotOnline)
		return
	}

	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)

	//洞房模板
	nextBabyNum := babyManager.GetBabyNum() + 1
	nextDongFangTemplate := babytemplate.GetBabyTemplateService().GetBabyDongFangTemplate(nextBabyNum)
	if nextDongFangTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"nextBabyNum": nextBabyNum,
			}).Warn("baby:处理洞房消息,洞房模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//宝宝数量
	if !babyManager.IsCanAddBaby() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("baby:处理洞房消息,当前宝宝个数已达上限，超生可提高宝宝数量")
		playerlogic.SendSystemMessage(pl, lang.BabyMaxNum)
		return
	}

	babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	useItemMap := make(map[int32]int32)
	useItemMap[nextDongFangTemplate.PregnantItem] += nextDongFangTemplate.PregnantCount
	if isUseItem {
		useItemMap[babyConstantTemplate.RiverItem] += 1
	}
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughItems(useItemMap) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"useItemMap": useItemMap,
			}).Warn("baby:处理洞房消息, 物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	itemUseReason := commonlog.InventoryLogReasonBabyDongFangUse
	flag := inventoryManager.BatchRemove(useItemMap, itemUseReason, itemUseReason.String())
	if !flag {
		panic(fmt.Errorf("baby: 洞房消耗物品应该成功"))
	}

	//洞房
	rate := nextDongFangTemplate.PregnantRate
	if isUseItem {
		itemTemp := item.GetItemService().GetItem(int(babyConstantTemplate.RiverItem))
		rate += itemTemp.TypeFlag1
	}

	failItemMap := make(map[int32]int32)
	isHit := mathutils.RandomHit(common.MAX_RATE, int(rate))
	var babyObj *playerbaby.PlayerBabyObject
	if isHit {
		babyObj = babyManager.AddBabyByDongFang()
	} else {
		failItemMap = nextDongFangTemplate.GetFailReturnItemMap()
		if !inventoryManager.HasEnoughSlots(failItemMap) {
			title := lang.GetLangService().ReadLang(lang.EmailInventorySlotNoEnough)
			content := lang.GetLangService().ReadLang(lang.BabyFailReturnMailContent)
			emaillogic.AddEmail(pl, title, content, failItemMap)
		} else {
			itemGetReason := commonlog.InventoryLogReasonBabyDongFangReturn
			flag := inventoryManager.BatchAdd(failItemMap, itemGetReason, itemGetReason.String())
			if !flag {
				panic("baby:洞房失败返还物品添加应该成功")
			}
		}
	}
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCBabyDongFang(isUseItem, isHit, failItemMap, babyObj)
	pl.SendMsg(scMsg)
	return
}
