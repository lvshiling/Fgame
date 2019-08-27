package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopeneventtypes "fgame/fgame/game/funcopen/event/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	advancedrewrewextendedtemplate "fgame/fgame/game/welfare/advancedrew/rew_extended/template"
	advancedrewrewextendedtypes "fgame/fgame/game/welfare/advancedrew/rew_extended/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家功能开启
func playerFuncOpen(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	funcType := data.(funcopentypes.FuncOpenType)

	//运营活动-进阶
	advancedType, ok := welfaretypes.FuncTypeToAdvancedType(funcType)
	if !ok {
		return
	}

	// 升阶奖励(永久活动随功能开启)
	advancedRewExtended(pl, advancedType)

	return
}

func advancedRewExtended(pl player.Player, advancedType welfaretypes.AdvancedType) {
	now := global.GetGame().GetTimeService().Now()
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypeRewExtended
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	rewTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range rewTimeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}

		groupTemp := groupInterface.(*advancedrewrewextendedtemplate.GroupTemplateRewExtended)
		rewType := groupTemp.GetAdvancedType()
		if advancedType != rewType {
			continue
		}
		expireTime := groupTemp.GetAdvancedRewExpireTime()
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*advancedrewrewextendedtypes.AdvancedRewExtendedInfo)
		info.RewType = advancedType
		info.ExpireTime = expireTime + now
		info.AdvancedNum = 1
		welfareManager.UpdateObj(obj)
	}
}

func init() {
	gameevent.AddEventListener(funcopeneventtypes.EventTypeFuncOpen, event.EventListenerFunc(playerFuncOpen))
}