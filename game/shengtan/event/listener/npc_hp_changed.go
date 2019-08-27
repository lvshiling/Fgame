package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	npceventtypes "fgame/fgame/game/npc/event/types"
	"fgame/fgame/game/npc/npc"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/shengtan/pbutil"
	shengtanscene "fgame/fgame/game/shengtan/scene"
	shengtantemplate "fgame/fgame/game/shengtan/template"
	"fmt"
	"math"
)

//圣坛
func npcHPChanged(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}

	s := n.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeAllianceShengTan {
		return
	}
	sd, ok := s.SceneDelegate().(shengtanscene.ShengTanSceneData)
	if !ok {
		return
	}
	shengTanTemplate := shengtantemplate.GetShengTanTemplateService().GetShengTanTemplate()

	if n.GetBiologyTemplate().TemplateId() != int(shengTanTemplate.ShengtanId) {
		return
	}

	eventData, ok := data.(*npc.NPCHPChangedEventData)
	oldHp := eventData.GetOldHP()
	newHp := eventData.GetNewHP()

	//广播血量变化
	scShengTanSceneBossHpChanged := pbutil.BuildSCShengTanSceneBossHpChanged(newHp)
	s.BroadcastMsg(scShengTanSceneBossHpChanged)
	if oldHp <= newHp {
		return
	}

	maxHp := n.GetMaxHP()
	oldPercent := int32(math.Floor(float64(oldHp) / float64(maxHp) * float64(common.MAX_RATE)))
	newPercent := int32(math.Floor(float64(newHp) / float64(maxHp) * float64(common.MAX_RATE)))
	threshold := shengTanTemplate.GonggaoHpPercent
	if oldPercent/threshold == newPercent/threshold {
		return
	}
	// 发公告
	noticeArgs := []int64{int64(chattypes.ChatLinkTypeOpenView), funcopentypes.FuncOpenTypeAllianceAltar, 0}
	joinLink := coreutils.FormatLink(chattypes.ButtonTypeToGoNow, noticeArgs)
	noticeContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ShengTanHpChangedNotice), fmt.Sprintf("%d", newPercent/100), joinLink)
	chatlogic.SystemBroadcastAllianceId(sd.GetAllianceId(), chattypes.MsgTypeText, []byte(noticeContent))

	return nil
}

func init() {
	gameevent.AddEventListener(npceventtypes.EventTypeNPCHPChanged, event.EventListenerFunc(npcHPChanged))
}
