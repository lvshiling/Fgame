package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	advancedfeedbacktypes "fgame/fgame/game/welfare/advanced/feedback/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeAdvanced, welfaretypes.OpenActivityAdvancedSubTypeFeedback, reddot.HandlerFunc(handleRedDotAdvanced))
}

//升阶返利红点（废弃）
func handleRedDotAdvanced(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	//进阶返利
	info := obj.GetActivityData().(*advancedfeedbacktypes.AdvancedInfo)
	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	maxRewNum := int32(0)
	for _, temp := range tempList {
		//对应的充值日
		cycDay := temp.Value1
		if cycDay != info.AdvancedDay {
			continue
		}

		if temp.Value2 > info.DanNum {
			continue
		}
		maxRewNum += 1
	}

	if len(info.RewRecord) >= int(maxRewNum) {
		return
	}

	isNotice = true
	return
}
