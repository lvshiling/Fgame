package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	onearenaeventtypes "fgame/fgame/game/onearena/event/types"
	onearenalogic "fgame/fgame/game/onearena/logic"
	playeronearena "fgame/fgame/game/onearena/player"
	onearenatemplate "fgame/fgame/game/onearena/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fmt"
)

//灵池抢夺成功
func oneArenaRobSucess(target event.EventTarget, data event.EventData) (err error) {
	eventData := data.(*onearenaeventtypes.OneArenaRobSucessEventData)

	oneArenaData := eventData.GetOneArenaData()
	peerArenaData := eventData.GetPeerOneArenaData()

	playerId := oneArenaData.GetPlayerId()
	level := oneArenaData.GetLevel()
	pos := oneArenaData.GetPos()

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}
	if peerArenaData != nil {
		err = onearenalogic.PeerRobbedRecord(peerArenaData, pl.GetName(), true)
		if err != nil {
			return
		}
	}

	manager := pl.GetPlayerDataManager(types.PlayerOneArenaDataManagerType).(*playeronearena.PlayerOneArenaDataManager)
	manager.ReplaceOneArena(level, pos)

	oneArenaTemplate := onearenatemplate.GetOneArenaTemplateService().GetOneArenaTemplateByLevel(level, pos)
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	oneArenaName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(oneArenaTemplate.Name))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OneArenaRobSucess), playerName, oneArenaName)
	//跑马登
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	//发送系统频道
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	return
}

func init() {
	gameevent.AddEventListener(onearenaeventtypes.EventTypePlayerOneArenaSucess, event.EventListenerFunc(oneArenaRobSucess))
}
