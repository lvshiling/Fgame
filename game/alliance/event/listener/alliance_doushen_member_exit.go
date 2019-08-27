package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancelogic "fgame/fgame/game/alliance/logic"
	chattypes "fgame/fgame/game/chat/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fmt"
)

//斗神成员退出仙盟
func douShenMemberExit(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)
	exitMem := data.(*alliance.AllianceMemberObject)

	pl := player.GetOnlinePlayerManager().GetPlayerById(exitMem.GetMemberId())
	if pl == nil {
		return
	}

	for _, mem := range al.GetMemberList() {
		p := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if p == nil {
			continue
		}

		alliancelogic.DoushenChanged(p)

		//斗神殿成员退出提示
		if exitMem.GetLingyuId() == 0 {
			continue
		}

		douShenName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(exitMem.GetName()))
		textWithColor := coreutils.FormatColor(chattypes.ColorTypeEmailRedWord, "战斗力大幅受损")
		title := lang.GetLangService().ReadLang(lang.AllianceDouShenExitMailTitle)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceDouShenExitMailContent), douShenName, textWithColor)
		emaillogic.AddEmail(p, title, content, nil)
	}

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceDouShenMemberExit, event.EventListenerFunc(douShenMemberExit))
}
