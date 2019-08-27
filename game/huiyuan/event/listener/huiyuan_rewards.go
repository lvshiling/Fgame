package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	huiyuaneventtypes "fgame/fgame/game/huiyuan/event/types"
	huiyuantemplate "fgame/fgame/game/huiyuan/template"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//会员未领取奖励
func playerHuiYuanRewards(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	huiyuanRewardsData := data.(*huiyuaneventtypes.HuiYuanRewardsEventData)
	huiyuanType := huiyuanRewardsData.GetHuiYuanType()
	dayNum := huiyuanRewardsData.GetRewardsDayNum()
	buyTime := huiyuanRewardsData.GetBuyTime()

	maxValidDay := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEmailSaveDay)
	if dayNum > maxValidDay {
		dayNum = maxValidDay
	}
	houtaiType := center.GetCenterService().GetZhiZunType()
	temp := huiyuantemplate.GetHuiYuanTemplateService().GetHuiYuanTemplate(houtaiType, huiyuanType)
	now := global.GetGame().GetTimeService().Now()
	beginTimeOfDay, err := timeutils.BeginOfNow(now)
	if err != nil {
		return
	}

	for index := int32(1); index <= dayNum; index++ {
		createTime := beginTimeOfDay - int64(common.DAY)*int64(index)
		isSame, err := timeutils.IsSameDay(createTime, buyTime)
		if err != nil {
			return err
		}

		var rewItemMap map[int32]int32
		if isSame {
			rewItemMap = temp.GetEmailFirstRewItemMap()
		} else {
			rewItemMap = temp.GetEmailRewItemMap()
		}

		title := lang.GetLangService().ReadLang(lang.EmailOpenActivityHuiYuanTitle)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityHuiYuanContent))
		emaillogic.AddEmailDefinTime(pl, title, content, createTime, rewItemMap)
	}

	return
}

func init() {
	gameevent.AddEventListener(huiyuaneventtypes.EventTypeHuiYuanRewards, event.EventListenerFunc(playerHuiYuanRewards))
}
