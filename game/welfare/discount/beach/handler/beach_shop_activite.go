package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	discountbeachtemplate "fgame/fgame/game/welfare/discount/beach/template"
	discountbeachtypes "fgame/fgame/game/welfare/discount/beach/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_BEACH_SHOP_ACTIVITE_TYPE), dispatch.HandlerFunc(handleBeachShopActivite))
}

func handleBeachShopActivite(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare: 处理沙滩商店激活消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityBeachShopActivate)
	groupId := csMsg.GetGroupId()

	err = beachShopActivite(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare: 处理沙滩商店激活消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare：处理沙滩商店激活消息,成功")

	return
}

func beachShopActivite(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeBeach

	// 校验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*discountbeachtypes.BeachShopInfo)
	// 判断是否已经激活
	if info.IsActivited() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare: 该沙滩商店已经激活")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityDiscountBeachShopActivite)
		return
	}

	groupTempI := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupTempI == nil {
		return
	}
	openTemp := groupTempI.(*discountbeachtemplate.GroupTemplateDiscountBeachShop)
	itemMap := openTemp.GetAvtiviteItemMap()

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(itemMap) != 0 {
		if !inventoryManager.HasEnoughItems(itemMap) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("welfare: 物品不足，无法激活")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	reason := commonlog.InventoryLogReasonBeachShopActivite
	reasonText := fmt.Sprintf(reason.String(), openTemp.GetActivityName())
	flag := inventoryManager.BatchRemove(itemMap, reason, reasonText)
	if !flag {
		panic("welfare: 激活沙滩商店消耗物品应该成功")
	}

	// 更新数据
	info.ActiviteSuccess()
	welfareManager.UpdateObj(obj)

	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityBeachShopActivite(groupId)
	pl.SendMsg(scMsg)

	return
}
