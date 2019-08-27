package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"
	fabaotemplate "fgame/fgame/game/fabao/template"
	fabaotypes "fgame/fgame/game/fabao/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	"fmt"
)

//玩家幻化激活
func playerFaBaoUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	faBaoId := data.(int)
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBao(faBaoId)
	if faBaoTemplate == nil {
		return
	}
	if faBaoTemplate.GetTyp() != fabaotypes.FaBaoTypeSkin {
		return
	}

	// 计算战斗力
	battlePropertyMap := faBaoTemplate.GetBattleProperty()
	power := propertylogic.CulculateForce(battlePropertyMap)

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	faBaoName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(faBaoTemplate.Name))
	powerStr := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", power))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FaBaoUnrealActivateNotice), playerName, faBaoName, powerStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(fabaoeventtypes.EventTypeFaBaoUnrealActivate, event.EventListenerFunc(playerFaBaoUnrealActivate))
}
