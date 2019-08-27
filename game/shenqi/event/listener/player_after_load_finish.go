package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shenqi/pbutil"
	playershenqi "fgame/fgame/game/shenqi/player"
	shenqitemplate "fgame/fgame/game/shenqi/template"
	shenqitypes "fgame/fgame/game/shenqi/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//加载完成后
func playerShenQiAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeShenQi) {
		return
	}
	manager := p.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	debrisMap := manager.GetShenQiDebrisMap()
	smeltMap := manager.GetShenQiSmeltMap()
	qiLingMap := manager.GetShenQiQiLingMap()
	shenQiOjb := manager.GetShenQiOjb()
	scMsg := pbutil.BuildSCShenQiInfoGet(qiLingMap, debrisMap, smeltMap, shenQiOjb.LingQiNum)
	p.SendMsg(scMsg)

	for typ := shenqitypes.MinShenQiType; typ <= shenqitypes.MaxShenQiType; typ++ {
		//加载技能
		shenQiLevel := manager.GetShenQiDebrisMinLevelByShenQi(typ)
		skillId := shenqitemplate.GetShenQiTemplateService().GetShenQiSkillId(typ, shenQiLevel)
		if skillId == 0 {
			continue
		}
		skilllogic.TempSkillChange(p, 0, skillId)
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerShenQiAfterLoadFinish))
}
