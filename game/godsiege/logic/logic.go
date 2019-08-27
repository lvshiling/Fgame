package logic

import (
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/godsiege/pbutil"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	"fgame/fgame/game/player"
	gametemplate "fgame/fgame/game/template"
)

//神兽攻城-麒麟
func PlayerEnterGodSiegeQiLinScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	//进入跨服
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeGodSiegeQiLin)
	flag = true
	return
}

//神兽攻城-火凤
func PlayerEnterGodSiegeHuoFengScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	//进入跨服
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeGodSiegeHuoFeng)
	flag = true
	return
}

//神兽攻城-毒龙
func PlayerEnterGodSiegeDuLongScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	//进入跨服
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeGodSiegeDuLong)
	flag = true
	return
}

//发送参加神兽攻城
func GodSiegeAttendSend(pl player.Player, crossType crosstypes.CrossType) {
	godType := godsiegetypes.GodSiegeTypeQiLin
	switch crossType {
	case crosstypes.CrossTypeGodSiegeQiLin:
		{
			break
		}
	case crosstypes.CrossTypeGodSiegeHuoFeng:
		{
			godType = godsiegetypes.GodSiegeTypeHuoFeng
			break
		}
	case crosstypes.CrossTypeGodSiegeDuLong:
		{
			godType = godsiegetypes.GodSiegeTypeDuLong
			break
		}
	case crosstypes.CrossTypeDenseWat:
		{
			godType = godsiegetypes.GodSiegeTypeDenseWat
		}
	default:
		return
	}
	siGodSiegeAttend := pbutil.BuildSIGodSiegeAttend(int32(godType))
	pl.SendCrossMsg(siGodSiegeAttend)
}

//发送玩家取消排队
func GodSiegeCancleLineUpSend(pl player.Player, godType int32) {
	siGodSiegeCancleLineUp := pbutil.BuildSIGodSiegeCancleLineUp(godType)
	pl.SendCrossMsg(siGodSiegeCancleLineUp)
}

func GodSiegeLineUpSuccess(pl player.Player, godType int32) {
	siGodSiegeLineUpSuccess := pbutil.BuildSIGodSiegeLineUpSuccess(godType)
	pl.SendCrossMsg(siGodSiegeLineUpSuccess)
}

func GodSiegeFinishLineUpCancle(pl player.Player, godType int32) {
	siGodSiegeFinishLineUpCancle := pbutil.BuildSIGodSiegeFinishLineUpCancle(godType)
	pl.SendCrossMsg(siGodSiegeFinishLineUpCancle)
}

func BroadGodSiegeLineUpChanged(godType int32, pos int32, lineList []int64) {
	for index, playerId := range lineList {
		if int32(index) < pos {
			continue
		}
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		scGodSiegeLineUpChanged := pbutil.BuildSCGodSiegeLineUp(godType, int32(index))
		pl.SendMsg(scGodSiegeLineUpChanged)
	}
}

func BroadGodSiegeFinishToLineUp(godType int32, lineList []int64) {
	for _, playerId := range lineList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		pl.GodSiegeCancleLineUp()
		scGodSiegeFinishToLineUp := pbutil.BuildSCGodSiegeFinishToLineUp(godType)
		pl.SendMsg(scGodSiegeFinishToLineUp)
	}
}
