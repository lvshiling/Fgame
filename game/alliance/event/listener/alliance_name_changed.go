package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	playeralliance "fgame/fgame/game/alliance/player"
	chattypes "fgame/fgame/game/chat/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

type postData struct {
	al              *alliance.Alliance
	oldAllianceName string
}

//仙盟名改变
func allianceNameChanged(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)
	oldAllianceNmae := data.(string)

	for _, member := range al.GetMemberList() {
		memberPl := player.GetOnlinePlayerManager().GetPlayerById(member.GetMemberId())
		if memberPl == nil {
			playerName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(al.GetMengzhuName()))
			oldAlName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(oldAllianceNmae))
			newAlName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(al.GetAllianceName()))

			title := lang.GetLangService().ReadLang(lang.AllianceRenameMailTitle)
			content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceRenameMailContent), playerName, oldAlName, newAlName)
			emaillogic.AddOfflineEmail(member.GetMemberId(), title, content, nil)
			continue
		}

		data := &postData{
			al:              al,
			oldAllianceName: oldAllianceNmae,
		}
		ctx := scene.WithPlayer(context.Background(), memberPl)
		allianceLevelChangedMsg := message.NewScheduleMessage(onAllianceNameChanged, ctx, data, nil)
		memberPl.Post(allianceLevelChangedMsg)
	}

	return
}

//仙盟名变化回调
func onAllianceNameChanged(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	data := result.(*postData)
	al := data.al

	allianceManager := tpl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceManager.SyncAllianceName(al.GetAllianceName())

	playerName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(al.GetMengzhuName()))
	oldAlName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(data.oldAllianceName))
	newAlName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(al.GetAllianceName()))

	title := lang.GetLangService().ReadLang(lang.AllianceRenameMailTitle)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceRenameMailContent), playerName, oldAlName, newAlName)
	emaillogic.AddEmail(tpl, title, content, nil)

	return nil
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceNameChanged, event.EventListenerFunc(allianceNameChanged))
}
