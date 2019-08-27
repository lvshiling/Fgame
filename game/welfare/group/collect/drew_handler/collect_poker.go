package drew_handler

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	drewchargedrewrelatehandler "fgame/fgame/game/welfare/drew/charge_drew/relate_handler"
	drewcommontypes "fgame/fgame/game/welfare/drew/common/types"
	groupcollectenum "fgame/fgame/game/welfare/group/collect/enum"
	groupcollecttypes "fgame/fgame/game/welfare/group/collect/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

func init() {
	drewchargedrewrelatehandler.RegistRelateHandler(welfaretypes.OpenActivityTypeGroup, welfaretypes.OpenActivityGroupSubTypeCollectPoker, drewchargedrewrelatehandler.RelateHandlerFunc(collectPoker))
}

// 摸金卡牌收集
func collectPoker(pl player.Player, relateGroupId int32, msg *uipb.SCOpenActivityChargeDrewAttend) {
	typ := welfaretypes.OpenActivityTypeGroup
	subType := welfaretypes.OpenActivityGroupSubTypeCollectPoker

	attendNum := drewcommontypes.LuckyDrewAttendType(msg.GetDrewType()).GetAttendNum()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	if !welfarelogic.IsOnActivityTime(relateGroupId) {
		return
	}

	relateObj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, relateGroupId)
	info := relateObj.GetActivityData().(*groupcollecttypes.CollectRewInfo)

	var pokerList []int32
	var rewPokerTypeList []int32
	for attendNum > 0 {
		poker := info.AddPoker()
		isFinish, pokerType := info.CheckCollect()
		if isFinish {
			addFinishPokerRew(pl, relateGroupId, pokerType)
			info.AddRewRecord(pokerType)
			rewPokerTypeList = append(rewPokerTypeList, int32(pokerType))
		}

		attendNum -= 1
		pokerList = append(pokerList, poker)
	}
	welfareManager.UpdateObj(relateObj)

	pbutil.BuildSCOpenActivityChargeDrewAttendWithDrewPokerInfo(msg, pokerList, rewPokerTypeList)
}

// 奖励
func addFinishPokerRew(pl player.Player, groupId int32, pokerType groupcollectenum.PokerType) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	pokerTemplate := welfaretemplate.GetWelfareTemplateService().GetCollectPokerTemplate(groupId, pokerType)
	if pokerTemplate == nil {
		return
	}

	rewGold := int64(pokerTemplate.RawGold)
	rewBindGold := int64(pokerTemplate.RawBindGold)
	rewSilver := int64(pokerTemplate.RawYinliang)

	silverReason := commonlog.SilverLogReasonGroupCollectRew
	goldReason := commonlog.GoldLogReasonGroupCollectRew
	silverReasonText := fmt.Sprintf(silverReason.String(), pokerType)
	goldReasonText := fmt.Sprintf(goldReason.String(), pokerType)
	flag := propertyManager.AddMoney(rewBindGold, rewGold, goldReason, goldReasonText, rewSilver, silverReason, silverReasonText)
	if !flag {
		panic("welfare:添加摸金卡牌收集奖励应该成功")
	}
}
