package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	fourgodtemplate "fgame/fgame/game/fourgod/template"
	noticelogic "fgame/fgame/game/notice/logic"
	npceventtypes "fgame/fgame/game/npc/event/types"
	"fgame/fgame/game/npc/npc"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
	"math"
)

//四神boss血量变化
func npcHPChanged(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}

	if n.GetScene().MapTemplate().GetMapType() != scenetypes.SceneTypeFourGodWar {
		return
	}
	fourGodTemplate := fourgodtemplate.GetFourGodTemplateService().GetFourGodConstTemplate()
	if fourGodTemplate == nil {
		return
	}
	if n.GetBiologyTemplate().TemplateId() != int(fourGodTemplate.BossId) {
		return
	}
	eventData, ok := data.(*npc.NPCHPChangedEventData)
	oldHp := eventData.GetOldHP()
	newHp := eventData.GetNewHP()
	if oldHp <= newHp {
		return
	}
	maxHp := n.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	oldPercent := int32(math.Floor(float64(oldHp) / float64(maxHp) * float64(100)))
	newPercent := int32(math.Floor(float64(newHp) / float64(maxHp) * float64(100)))

	percent, flag := fourgodtemplate.GetFourGodTemplateService().GetBossThreshold(oldPercent, newPercent)
	if !flag {
		return
	}

	//TODO 公告
	bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(n.GetBiologyTemplate().Name))
	bossBornBlood := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FourGodBossBlood), bossName, percent)
	bossBornBloodChat := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FourGodBossBloodChat), bossName, percent)
	//TODO 跑马灯
	noticelogic.NoticeNumBroadcast([]byte(bossBornBlood), 0, int32(1))
	//系统频道
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(bossBornBloodChat))

	return nil
}

func init() {
	gameevent.AddEventListener(npceventtypes.EventTypeNPCHPChanged, event.EventListenerFunc(npcHPChanged))
}
