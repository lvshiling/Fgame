package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/tower/pbutil"
	"fgame/fgame/game/tower/tower"
	"fmt"
)

//boss死亡
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}

	attackId, ok := data.(int64)
	if !ok {
		return
	}

	//校验怪物类型
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeTowerBoss {
		return
	}
	//TODO xzk:通过场景获取
	attackPl := player.GetOnlinePlayerManager().GetPlayerById(attackId)
	if attackPl == nil {
		return
	}

	s := n.GetScene()
	if s == nil {
		return
	}

	sd := s.SceneDelegate().(tower.TowerSceneData)

	floor := int32(sd.GetTowerTemplate().TemplateId())
	bossName := coreutils.FormatNoticeStr(n.GetBiologyTemplate().Name)
	playerName := coreutils.FormatNoticeStr(attackPl.GetName())
	scMsg := pbutil.BuildSCTowerBossInfoDeadNotice(n, playerName, bossName, floor)
	for _, s := range scene.GetSceneService().GetAllTowerScene() {
		for _, pl := range s.GetAllPlayers() {
			pl.SendMsg(scMsg)
		}
	}

	//系统广播
	colorPlayerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, playerName)
	colorBossName := coreutils.FormatColor(chattypes.ColorTypeBoss, bossName)
	format := lang.GetLangService().ReadLang(lang.TowerBossDead)
	content := fmt.Sprintf(format, colorPlayerName, colorBossName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	//跑马灯
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
