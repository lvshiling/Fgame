package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	"fgame/fgame/cross/player/player"
	shenmologic "fgame/fgame/cross/shenmo/logic"
	"fgame/fgame/cross/shenmo/pbutil"
	"fgame/fgame/cross/shenmo/shenmo"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	playerlogic "fgame/fgame/game/player/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	shenmotemplate "fgame/fgame/game/shenmo/template"
	"fmt"
)

//玩家被杀
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo, ok := target.(scene.BattleObject)
	if !ok {
		return
	}
	pl, ok := bo.(*player.Player)
	if !ok {
		return
	}
	attackId, ok := data.(int64)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossShenMo {
		return
	}
	attackPl := pl.GetScene().GetPlayer(attackId)
	if attackPl == nil {
		return
	}

	//cd中
	flag, cdMill := pl.IfCanKilledInActivity(activitytypes.ActivityTypeShenMoWar)
	if !flag {
		cdSecond := cdMill / int32(common.SECOND)
		playerlogic.SendSystemMessage(attackPl, lang.PlayerKilledRewardCd, fmt.Sprintf("%d", cdSecond))
		return
	}

	killNum := pl.GetShenMoKillNum()
	shenMoTitleTemplate := shenmotemplate.GetShenMoTemplateService().GetShenMoTitleTemplateByKillNum(killNum)
	if shenMoTitleTemplate == nil {
		return
	}

	pl.KilledInActivity(activitytypes.ActivityTypeShenMoWar)
	curKillNum := attackPl.GetShenMoKillNum()
	attackPl.SetShenMoKillNum(curKillNum + 1)

	scShenMoKillNumChanged := pbutil.BuildSCShenMoKillNumChanged(curKillNum + 1)
	attackPl.SendMsg(scShenMoKillNumChanged)

	//增加仙盟积分
	allianceId := attackPl.GetAllianceId()
	if allianceId != 0 && shenMoTitleTemplate.GiveJiFen != 0 {
		serverId := attackPl.GetServerId()
		allianceName := attackPl.GetAllianceName()
		shenmo.GetShenMoService().AddJiFenNum(serverId, allianceId, allianceName, shenMoTitleTemplate.GiveJiFen)
		shenmologic.JiFenChangedAllianceBroadcast(s, allianceId)
	}

	//推送本服
	isPlayerGongXunAdd := pbutil.BuildISPlayerGongXunAdd(shenMoTitleTemplate.GiveGongXun)
	attackPl.SendMsg(isPlayerGongXunAdd)

	isShenMoKillNumChanged := pbutil.BuildISShenMoKillNumChanged(curKillNum + 1)
	attackPl.SendMsg(isShenMoKillNumChanged)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
