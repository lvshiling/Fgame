package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	shenmologic "fgame/fgame/game/shenmo/logic"
	"fgame/fgame/game/shenmo/pbutil"
	"fgame/fgame/game/shenmo/shenmo"
	shenmotemplate "fgame/fgame/game/shenmo/template"
	"fmt"
)

//玩家被杀
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo, ok := target.(scene.BattleObject)
	if !ok {
		return
	}
	pl, ok := bo.(player.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	attackId, ok := data.(int64)
	if !ok {
		return
	}

	// TODO xzk:类型可能修改
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossShenMo {
		return
	}
	spl := s.GetPlayer(attackId)
	if spl == nil {
		return
	}
	attackPl := spl.(player.Player)

	//被杀CD中
	flag, cdMill := pl.IfCanKilledInActivity(activitytypes.ActivityTypeShenMoWar)
	if !flag {
		cdSecond := cdMill / int32(common.SECOND)
		playerlogic.SendSystemMessage(attackPl, lang.PlayerKilledRewardCd, fmt.Sprintf("%d", cdSecond))
		return
	}

	//被杀玩家 被杀次数相关配置
	killNum := pl.GetShenMoKillNum()
	shenMoTitleTemplate := shenmotemplate.GetShenMoTemplateService().GetShenMoTitleTemplateByKillNum(killNum)
	if shenMoTitleTemplate == nil {
		return
	}

	// 被击杀记录
	pl.KilledInActivity(activitytypes.ActivityTypeShenMoWar)

	// 击杀信息
	curKillNum := attackPl.GetShenMoKillNum() + 1
	attackPl.SetShenMoKillNum(curKillNum)
	killNumScMsg := pbutil.BuildSCShenMoKillNumChanged(curKillNum)
	attackPl.SendMsg(killNumScMsg)
	//增加功勋
	shenmologic.AddGongXun(attackPl, shenMoTitleTemplate.GiveGongXun)

	//增加仙盟积分
	allianceId := spl.GetAllianceId()
	if allianceId != 0 && shenMoTitleTemplate.GiveJiFen != 0 {
		// serverId := spl.GetServerId()
		allianceName := spl.GetAllianceName()
		shenmo.GetShenMoService().AddLocalJiFenNum(allianceId, allianceName, shenMoTitleTemplate.GiveJiFen)
		shenmologic.JiFenChangedAllianceBroadcast(s, allianceId)
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
