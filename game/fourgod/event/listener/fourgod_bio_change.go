package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	fourgodeventtypes "fgame/fgame/game/fourgod/event/types"
	fourgodlogic "fgame/fgame/game/fourgod/logic"
	"fgame/fgame/game/fourgod/pbutil"
	fourgodscene "fgame/fgame/game/fourgod/scene"
	funcopentypes "fgame/fgame/game/funcopen/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//四神遗迹副本生物改变
func fourGodBioChange(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(fourgodscene.FourGodWarSceneData)
	if !ok {
		return
	}

	s := sd.GetScene()
	npc := data.(scene.NPC)
	npcId := npc.GetId()

	allPlayers := s.GetAllPlayers()
	scFourGodBioBroadcast := pbuitl.BuildSCFourGodBioBroadcast(npcId, npc)
	fourgodlogic.BroadcastMsgInScene(allPlayers, scFourGodBioBroadcast)

	//Boss刷新
	if npc.GetBiologyTemplate().GetBiologyScriptType() == scenetypes.BiologyScriptTypeFourGodBoss && !npc.IsDead() {
		bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(npc.GetBiologyTemplate().Name))
		args := []int64{int64(chattypes.ChatLinkTypeOpenView), funcopentypes.FuncOpenTypeFourGod, 0}
		joinLink := coreutils.FormatLink(chattypes.ButtonTypeToGoNow, args)
		bossBornContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FourGodBossBorn), bossName)
		bossRefreshContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FourGodBossRefresh), bossName, joinLink)
		//TODO 跑马灯
		noticelogic.NoticeNumBroadcast([]byte(bossBornContent), 0, int32(1))
		//系统频道
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(bossRefreshContent))
	}
	return
}

func init() {
	gameevent.AddEventListener(fourgodeventtypes.EventTypeFourGodBioChange, event.EventListenerFunc(fourGodBioChange))
}
