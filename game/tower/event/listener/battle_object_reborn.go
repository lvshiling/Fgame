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
	"fgame/fgame/game/tower/pbutil"
	"fgame/fgame/game/tower/tower"
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
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeTowerBoss {
		return
	}
	s := n.GetScene()
	if s == nil {
		return
	}

	sd := s.SceneDelegate().(tower.TowerSceneData)

	//场景类型是打宝塔的所有玩家
	floor := int32(sd.GetTowerTemplate().TemplateId())
	bossName := coreutils.FormatNoticeStr(n.GetBiologyTemplate().Name)
	mapName := coreutils.FormatNoticeStr(n.GetScene().MapTemplate().Name)
	scMsg := pbutil.BuildSCTowerBossInfoReBornNotice(n, bossName, mapName, floor)
	for _, s := range scene.GetSceneService().GetAllTowerScene() {
		for _, pl := range s.GetAllPlayers() {
			pl.SendMsg(scMsg)
		}
	}

	args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(funcopentypes.FuncOpenTypeDaBaoTower), int64(floor)}
	link := coreutils.FormatLink(chattypes.ButtonTypeToKill, args)
	//系统广播
	colorBossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(bossName))
	colorMapName := coreutils.FormatColor(chattypes.ColorTypeBossMap, coreutils.FormatNoticeStr(mapName))
	format := lang.GetLangService().ReadLang(lang.TowerBossReborn)
	content := fmt.Sprintf(format, colorBossName, colorMapName, link)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	//跑马灯
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectReborn, event.EventListenerFunc(battleObjectReborn))
}
