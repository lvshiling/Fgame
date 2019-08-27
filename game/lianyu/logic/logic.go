package logic

import (
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/lianyu/pbutil"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
)

func PlayerEnterLianYuScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	//进入跨服
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeLianYu)
	flag = true
	return
}

//发送参加无间炼狱
func LianYuAttendSend(pl player.Player) {
	siLianYuAttend := pbutil.BuildSILianYuAttend()
	pl.SendCrossMsg(siLianYuAttend)
}

//发送玩家取消排队
func LianYuCancleLineUpSend(pl player.Player) {
	siLianYuCancleLineUp := pbutil.BuildSILianYuCancleLineUp()
	pl.SendCrossMsg(siLianYuCancleLineUp)
}

func LianYuLineUpSuccess(pl player.Player) {
	siLianYuLineUpSuccess := pbutil.BuildSILianYuLineUpSuccess()
	pl.SendCrossMsg(siLianYuLineUpSuccess)
}

func LianYuFinishLineUpCancle(pl player.Player) {
	siLianYuFinishLineUpCancle := pbutil.BuildSILianYuFinishLineUpCancle()
	pl.SendCrossMsg(siLianYuFinishLineUpCancle)
}

// 排队人数变化
func BroadLianYuLineUpChanged(pos int32, lineList []int64) {
	for index, playerId := range lineList {
		if int32(index) < pos {
			continue
		}
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		scLianYuLineUpChanged := pbutil.BuildSCLianYuLineUpChanged(int32(index))
		pl.SendMsg(scLianYuLineUpChanged)
	}
}

func BroadLineYuFinishToLineUp(lineList []int64) {
	for _, playerId := range lineList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		scLianYuFinishToLineUp := pbutil.BuildSCLianYuFinishToLineUp()
		pl.SendMsg(scLianYuFinishToLineUp)
	}
}
