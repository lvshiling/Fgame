package logic

import (
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/shenmo/pbutil"
	"fgame/fgame/game/shenmo/shenmo"
	gametemplate "fgame/fgame/game/template"
)

//神魔战场
func PlayerEnterShenMoScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	timeTemp, _ := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
	if timeTemp == nil {
		return
	}
	endTime, _ := timeTemp.GetEndTime(now)
	shenMoEndTime := pl.GetShenMoEndTime()
	if endTime != shenMoEndTime {
		pl.SetShenMoEndTime(endTime)
	}
	//进入跨服
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeShenMoWar)
	flag = true
	return
}

//发送参加神魔战场
func ShenMoAttendSend(pl player.Player) {
	siShenMoAttend := pbutil.BuildSIShenMoAttend()
	pl.SendCrossMsg(siShenMoAttend)
}

//发送玩家取消排队
func ShenMoCancleLineUpSend(pl player.Player) {
	siShenMoCancleLineUp := pbutil.BuildSIShenMoCancleLineUp()
	pl.SendCrossMsg(siShenMoCancleLineUp)
}

func ShenMoLineUpSuccess(pl player.Player) {
	siShenMoLineUpSuccess := pbutil.BuildSIShenMoLineUpSuccess()
	pl.SendCrossMsg(siShenMoLineUpSuccess)
}

func ShenMoFinishLineUpCancle(pl player.Player) {
	siShenMoFinishLineUpCancle := pbutil.BuildSIShenMoFinishLineUpCancle()
	pl.SendCrossMsg(siShenMoFinishLineUpCancle)
}

func BroadShenMoLineUpChanged(pos int32, lineList []int64) {
	for index, playerId := range lineList {
		if int32(index) < pos {
			continue
		}
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		scShenMoLineUpChanged := pbutil.BuildSCShenMoLineUp(int32(index))
		pl.SendMsg(scShenMoLineUpChanged)
	}
}

func BroadShenMoFinishToLineUpCancle(lineList []int64) {
	for _, playerId := range lineList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		pl.ShenMoLineUp(false)
		scShenMoFinishToLineUp := pbutil.BuildSCShenMoFinishToLineUp()
		pl.SendMsg(scShenMoFinishToLineUp)
	}
}

// 添加功勋
func AddGongXun(pl player.Player, addGongXun int32) {
	curNum := pl.GetShenMoGongXunNum()
	totalNum := curNum + addGongXun
	pl.SetShenMoGongXunNum(totalNum)
}

func JiFenChangedAllianceBroadcast(s scene.Scene, allianceId int64) {
	if s == nil {
		return
	}

	jiFenNum := shenmo.GetShenMoService().GetJiFenNum(allianceId)
	if jiFenNum == 0 {
		return
	}

	scShenMoJiFenNumChanged := pbutil.BuildSCShenMoJiFenNumChanged(jiFenNum)
	allPlayers := s.GetAllPlayers()
	for _, pl := range allPlayers {
		if pl == nil {
			continue
		}
		if pl.GetAllianceId() != allianceId {
			continue
		}
		pl.SendMsg(scShenMoJiFenNumChanged)
	}
}
