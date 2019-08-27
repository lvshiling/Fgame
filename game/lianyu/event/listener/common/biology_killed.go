package common

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	lianyuscene "fgame/fgame/game/lianyu/scene"
	lianyutemplate "fgame/fgame/game/lianyu/template"
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
	s := npc.GetScene()
	if s == nil {
		return
	}

	biologyTemplate := npc.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}
	if biologyTemplate.GetDropJudgeType() == scenetypes.DropJudgeTypeMaxHurtOrTeam || biologyTemplate.GetDropJudgeType() == scenetypes.DropJudgeTypeMaxHurt {
		attackId = int64(0)
		maxDamage := int64(0)
		//TODO:zrc 临时处理
		//个人数据
		for tempAttackId, damage := range npc.GetAllDamages() {
			so := s.GetSceneObject(attackId)
			if so == nil {
				continue
			}
			_, ok := so.(scene.Player)
			if !ok {
				continue
			}

			if damage > maxDamage {
				attackId = tempAttackId
				maxDamage = damage
			}
		}
	}
	pl := s.GetPlayer(attackId)
	if pl == nil {
		return
	}

	if biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeCrossLianYuBoss {
		return
	}

	sd, ok := s.SceneDelegate().(lianyuscene.LianYuSceneData)
	if !ok {
		return
	}

	//公告
	constTemplate := lianyutemplate.GetLianYuTemplateService().GetConstantTemplate(sd.GetAcitvityType())
	if constTemplate == nil {
		return
	}
	bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(constTemplate.GetBiologyTemplate().Name))
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	bossDeadContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.LianYuBossDeadNotice), bossName, playerName)
	noticelogic.NoticeNumBroadcastScene(s, []byte(bossDeadContent), 0, int32(1))
	chatlogic.BroadcastScene(s, chattypes.MsgTypeText, []byte(bossDeadContent))
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
