package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	xueduneventtypes "fgame/fgame/game/xuedun/event/types"
	playerxuedun "fgame/fgame/game/xuedun/player"
	xueduntemplate "fgame/fgame/game/xuedun/template"
	"fmt"
	"math"
)

//玩家血盾阶数改变
func playerXueDunNumberChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	xueDunInfo := manager.GetXueDunInfo()
	number := xueDunInfo.GetNumber()
	star := xueDunInfo.GetStar()

	xueDunTemplate := xueduntemplate.GetXueDunTemplateService().GetXueDunNumber(number, star)
	if xueDunTemplate == nil {
		return
	}
	numberName := fmt.Sprintf("%s", xueDunTemplate.Name)
	diffPower := int32(math.Ceil(float64(xueDunTemplate.Note) / float64(common.MAX_RATE) * float64(100.0)))
	power := fmt.Sprintf("%d%s", diffPower, "%")

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	shenfaName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(numberName))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, coreutils.FormatNoticeStr(power))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.XueDunUpgradeNotice), playerName, shenfaName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(xueduneventtypes.EventTypeXueDunNumberChanged, event.EventListenerFunc(playerXueDunNumberChanged))
}
