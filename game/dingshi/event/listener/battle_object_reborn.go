package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	noticelogic "fgame/fgame/game/notice/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fmt"
)

//boss重生
func battleObjectReborn(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}

	//校验怪物类型
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeDingShiBoss {
		return
	}

	scBroadcast := pbutil.BuildSCWorldBossInfoBroadcast(n, worldbosstypes.BossTypeDingShi)
	for _, pl := range n.GetScene().GetAllPlayers() {
		pl.SendMsg(scBroadcast)
	}

	bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(n.GetBiologyTemplate().Name))
	mapName := coreutils.FormatColor(chattypes.ColorTypeBossMap, coreutils.FormatNoticeStr(n.GetScene().MapTemplate().Name))
	bossId := int64(n.GetBiologyTemplate().Id)
	args := []int64{int64(chattypes.ChatLinkTypeOpenView), funcopentypes.FuncOpenTypeDingShiBoss, bossId}
	joinLink := coreutils.FormatLink(chattypes.ButtonTypeToKill, args)

	//系统广播
	format := lang.GetLangService().ReadLang(lang.WorldBossReborn)
	content := fmt.Sprintf(format, bossName, mapName, joinLink)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	//跑马灯
	noticeFormat := lang.GetLangService().ReadLang(lang.WorldBossRebornNotice)
	noticeContent := fmt.Sprintf(noticeFormat, bossName, mapName)
	noticelogic.NoticeNumBroadcast([]byte(noticeContent), 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectReborn, event.EventListenerFunc(battleObjectReborn))
}
