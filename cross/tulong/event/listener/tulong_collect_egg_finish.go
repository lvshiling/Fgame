package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/cross/chat/logic"
	noticelogic "fgame/fgame/cross/notice/logic"
	tulongeventtypes "fgame/fgame/cross/tulong/event/types"
	"fgame/fgame/cross/tulong/pbutil"
	tulongscene "fgame/fgame/cross/tulong/scene"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	tulongtemplate "fgame/fgame/game/tulong/template"
	"fmt"
)

//屠龙采集龙蛋完成
func tuLongCollectEggFinish(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(tulongscene.TuLongSceneData)
	if !ok {
		return
	}
	npcId, ok := data.(int64)
	if !ok {
		return
	}
	collectEggMap := sd.GetCollectEgg()
	collectEggInfo, exist := collectEggMap[npcId]
	if !exist {
		return
	}

	pl := sd.GetScene().GetPlayer(collectEggInfo.GetPlayerId())
	if pl == nil {
		return
	}

	scTuLongCollectStop := pbutil.BuildSCTuLongCollectFinish(npcId)
	pl.SendMsg(scTuLongCollectStop)

	//TODO 公告
	npc := collectEggInfo.GetEggNpc()
	pos := npc.GetPosition()
	serverId := pl.GetServerId()
	tuLongConstTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongConstTemplate()
	allianceNameStr := fmt.Sprintf("s%d.%s", serverId, pl.GetAllianceName())
	allianceName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(allianceNameStr))
	mapNameStr := fmt.Sprintf("%s(%.2f,%.2f)", tuLongConstTemplate.GetMapTemplate().Name, pos.X, pos.Z)
	mapName := coreutils.FormatColor(chattypes.ColorTypeBossMap, coreutils.FormatNoticeStr(mapNameStr))
	smallBossBornContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.TuLongSmallBossBorn), allianceName, mapName)
	//TODO 跑马灯
	noticelogic.NoticeNumBroadcast([]byte(smallBossBornContent), 0, int32(1))
	//系统频道
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(smallBossBornContent))
	return
}

func init() {
	gameevent.AddEventListener(tulongeventtypes.EventTypeTuLongCollectFinish, event.EventListenerFunc(tuLongCollectEggFinish))
}
