package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancelogic "fgame/fgame/game/alliance/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fmt"
)

const (
	Mail = 5
)

//斗神领域变化
func allianceDouShenLingYuChanged(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)
	newJoinList := data.([]*alliance.AllianceMemberObject)

	for _, mem := range al.GetMemberList() {
		p := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if p == nil {
			continue
		}

		//等级限制
		levelLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDouShenLevelLimit)
		if p.GetLevel() < levelLimit {
			continue
		}

		// 重新计算斗神增幅
		alliancelogic.DoushenChanged(p)

		//斗神增幅邮件提示
		for _, doushen := range newJoinList {
			if mem.GetMemberId() == doushen.GetMemberId() {
				continue
			}

			douShenName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(doushen.GetName()))
			title := lang.GetLangService().ReadLang(lang.AllianceDouShenJoinMailTitle)
			content := ""
			if doushen.GetLingyuId() >= 5 {
				content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceDouShenJoinMailContentSenior), douShenName)
			} else {
				textWithColor := coreutils.FormatColor(chattypes.ColorTypeEmailRedWord, "无法给您带来战斗力的大幅提升")
				content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceDouShenJoinMailContentJunior), douShenName, textWithColor)
			}
			emaillogic.AddEmail(p, title, content, nil)
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceDouShenLingYuChanged, event.EventListenerFunc(allianceDouShenLingYuChanged))
}
