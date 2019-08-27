package common

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	godsiegetemplate "fgame/fgame/game/godsiege/template"
	noticelogic "fgame/fgame/game/notice/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//boss被杀
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo, ok := target.(scene.BattleObject)
	if !ok {
		return
	}
	npc, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	attackId, ok := data.(int64)
	if !ok {
		return
	}

	pl := npc.GetScene().GetPlayer(attackId)
	if pl == nil {
		return
	}

	biologyTemplate := npc.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}
	if biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeGodSiegeBoss {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossGodSiege {
		return
	}
	mapId := pl.GetScene().MapId()

	//公告
	constTemplate := godsiegetemplate.GetGodSiegeTemplateService().GetConstantTemplate()
	bossTemplate := constTemplate.GetBiologyTemplate(mapId)
	bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(bossTemplate.Name))
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	bossDeadContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.GodSiegeBossDeadNotice), bossName, playerName)
	noticelogic.NoticeNumBroadcastScene(pl.GetScene(), []byte(bossDeadContent), 0, int32(1))
	chatlogic.BroadcastScene(pl.GetScene(), chattypes.MsgTypeText, []byte(bossDeadContent))
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
