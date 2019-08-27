package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	droptemplate "fgame/fgame/game/drop/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	gamesession "fgame/fgame/game/session"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_FIRST_CHARGE_TYPE), dispatch.HandlerFunc(handlerWelfareFistCharge))
}

//处理领取首冲
func handlerWelfareFistCharge(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理领取首冲奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = welfareFirstCharge(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理领取首冲奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理领取首冲奖励请求完成")

	return
}

//领取首冲请求逻辑
func welfareFirstCharge(pl player.Player) (err error) {
	firstTemp := welfaretemplate.GetWelfareTemplateService().GetFirstCharge(pl.GetRole(), pl.GetSex())
	if firstTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:领取首冲奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)

	//是否首充
	if !welfareManager.IsFirstCharge() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:领取首冲奖励请求，不是首充用户")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotFirstChargeUser)
		return
	}

	//是否领取
	if welfareManager.IsReceiveFirstCharge() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:领取首冲奖励请求，已领取")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityHadReceiveFirstCharge)
		return
	}

	rewItemMap := firstTemp.GetRewItemMap()
	newItemDataList := droptemplate.ConvertToItemDataList(rewItemMap, itemtypes.ItemBindTypeUnBind)
	//背包空间
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemDataList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:领取首冲奖励请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//增加物品
	itemGetReason := commonlog.InventoryLogReasonFirstCharge
	flag := inventoryManager.BatchAddOfItemLevel(newItemDataList, itemGetReason, itemGetReason.String())
	if !flag {
		panic("welfare:first charge add item should be ok")
	}

	reasonGold := commonlog.GoldLogReasonFirstCharge
	reasonSilver := commonlog.SilverLogReasonFirstCharge
	reasonLevel := commonlog.LevelLogReasonFirstCharge

	rewSilver := firstTemp.RewSilver
	rewBindGold := firstTemp.RewGoldBind
	rewGold := firstTemp.RewGold
	rewExp := int32(0)
	rewExpPoint := int32(0)
	totalRewData := propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)
	flag = propertyManager.AddRewData(totalRewData, reasonGold, reasonGold.String(), reasonSilver, reasonSilver.String(), reasonLevel, reasonLevel.String())
	if !flag {
		panic("welfare:first charge add RewData should be ok")
	}

	welfareManager.ReceiveFirstCharge()

	//公告
	itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(rewItemMap)
	if len(itemNameLinkStr) > 0 {
		plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFirstChargeRewardsNotice), plName, itemNameLinkStr)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scWelfareFirstCharge := pbutil.BuildSCOpenActivityFirstCharge(totalRewData, rewItemMap)
	pl.SendMsg(scWelfareFirstCharge)
	return
}
