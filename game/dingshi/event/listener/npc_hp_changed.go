package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	dingshitemplate "fgame/fgame/game/dingshi/template"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
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

	n, ok := target.(scene.NPC)
	if !ok {
		return
	}
	s := n.GetScene()
	if s == nil {
		return
	}

	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeDingShiBoss {
		return
	}

	dingShiBossTemplate := dingshitemplate.GetDingShiTemplateService().GetDingShiBossTemplateByBiologyId(int32(n.GetBiologyTemplate().TemplateId()))
	if dingShiBossTemplate == nil {
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

	percent, flag := dingShiBossTemplate.GetBossThreshold(oldPercent, newPercent)
	if !flag {
		return
	}
	clientMapType := int64(s.MapTemplate().Type)
	mapId := int64(s.MapId())
	bossName := coreutils.FormatColor(chattypes.ColorTypeBoss, coreutils.FormatNoticeStr(n.GetBiologyTemplate().Name))

	args := []int64{int64(chattypes.ChatPlayerEnterScene), funcopentypes.FuncOpenTypeDingShiBoss, clientMapType, mapId}
	joinLink := coreutils.FormatLink(chattypes.ButtonTypeToGoNow, args)

	bossBornBlood := fmt.Sprintf(lang.GetLangService().ReadLang(lang.DingShiBossBlood), bossName, percent, joinLink)
	bossBornBloodChat := fmt.Sprintf(lang.GetLangService().ReadLang(lang.DingShiBossBlood), bossName, percent, joinLink)
	//TODO 跑马灯
	noticelogic.NoticeNumBroadcast([]byte(bossBornBlood), 0, int32(1))
	//系统频道
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(bossBornBloodChat))

	return nil
}

func init() {
	gameevent.AddEventListener(npceventtypes.EventTypeNPCHPChanged, event.EventListenerFunc(npcHPChanged))
}
