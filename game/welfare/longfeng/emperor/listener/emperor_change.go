package listener

import (
	commonlang "fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	emperoreventtypes "fgame/fgame/game/emperor/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	longfengemperortypes "fgame/fgame/game/welfare/longfeng/emperor/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家抢龙椅
func emperorChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeLongFeng
	subType := welfaretypes.OpenActivityDefaultSubTypeDefault

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*longfengemperortypes.LongFengInfo)
		info.RobTimes += 1
		welfareManager.UpdateObj(obj)

		//抢夺成功奖励
		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		openTmep := groupInterface.GetFirstOpenTemp()
		title := commonlang.GetLangService().ReadLang(commonlang.EmailOpenActivityLongFengRobTitle)
		content := commonlang.GetLangService().ReadLang(commonlang.EmailOpenActivityLongFengRobContent)
		now := global.GetGame().GetTimeService().Now()
		emaillogic.AddEmailItemLevel(pl, title, content, now, welfarelogic.ConvertToItemDataWithWelfareData(openTmep.GetEmailRewItemDataList(), openTmep.GetExpireType(), openTmep.GetExpireTime()))
	}

	return
}

func init() {
	gameevent.AddEventListener(emperoreventtypes.EmperorEventTypeRobed, event.EventListenerFunc(emperorChanged))
}
