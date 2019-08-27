package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	"fgame/fgame/game/common/common"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	houseventtypes "fgame/fgame/game/house/event/types"
	playerhouse "fgame/fgame/game/house/player"
	housetemplate "fgame/fgame/game/house/template"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家房子跨天前处理
func playerHouseBeforeCrossDay(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	houseObj, ok := data.(*playerhouse.PlayerHouseObject)
	if !ok {
		return
	}

	if houseObj.IsBroken() {
		return
	}

	houseTemp := housetemplate.GetHouseTemplateService().GetHouseTemplate(houseObj.GetHouseIndex(), houseObj.GetHouseType(), houseObj.GetHouseLevel())
	if houseTemp == nil {
		return
	}

	//计算可领取天数
	now := global.GetGame().GetTimeService().Now()
	preDay := now - int64(common.DAY)
	lastRefreshTime := houseObj.GetRefreshUpdateTime()
	diffDay, _ := timeutils.DiffDay(preDay, lastRefreshTime) //计算两个时间相差几天
	if !houseObj.IsRent() {
		diffDay += 1
	}
	if diffDay <= 0 {
		return
	}

	beginDay, err := timeutils.BeginOfNow(now) // 当天的零点的时间戳
	if err != nil {
		return
	}
	for initTimes := int32(1); initTimes <= diffDay; initTimes++ {
		createTime := beginDay - int64(initTimes)*int64(common.DAY)
		title := lang.GetLangService().ReadLang(lang.HouseRentEmailTitle)
		houseNum := houseTemp.HouseIndex + 1
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.HouseRentEmailContent), houseNum, houseTemp.Rent)
		emaillogic.AddEmailDefinTime(pl, title, content, createTime, houseTemp.GetRentItemMap())
	}

	return
}

func init() {
	gameevent.AddEventListener(houseventtypes.EventTypeHouseBeforeCrossDay, event.EventListenerFunc(playerHouseBeforeCrossDay))
}
