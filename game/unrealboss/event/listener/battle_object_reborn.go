package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/funcopen/funcopen"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/unrealboss/pbutil"
)

//幻境boss重生
func battleObjectReborn(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}

	//校验怪物类型
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeUnrealBoss {
		return
	}

	funcOpenTemp := funcopen.GetFuncOpenService().GetFuncOpenTemplate(funcopentypes.FuncOpenTypeBossHuanJing)
	scMsg := pbutil.BuildSCUnrealBossInfoBroadcast(n)
	playerList := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, pl := range playerList {
		if pl.GetLevel() < funcOpenTemp.OpenedLevel {
			continue
		}
		pl.SendMsg(scMsg)
	}

	// bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(n.GetBiologyTemplate().Name))
	// mapName := coreutils.FormatColor(chattypes.ColorTypeBossMap, coreutils.FormatNoticeStr(n.GetScene().MapTemplate().Name))
	// bossId := int64(n.GetBiologyTemplate().Id)
	// args := []int64{int64(chattypes.ChatLinkTypeOpenView), funcopentypes.FuncOpenTypeWorldBoss, bossId}
	// joinLink := coreutils.FormatLink( chattypes.ButtonTypeToKill, args)

	// //系统广播
	// format := lang.GetLangService().ReadLang(lang.WorldBossReborn)
	// content := fmt.Sprintf(format, bossName, mapName, joinLink)
	// chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	// //跑马灯
	// noticeFormat := lang.GetLangService().ReadLang(lang.WorldBossRebornNotice)
	// noticeContent := fmt.Sprintf(noticeFormat, bossName, mapName)
	// noticelogic.NoticeNumBroadcast([]byte(noticeContent), 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectReborn, event.EventListenerFunc(battleObjectReborn))
}