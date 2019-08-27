package logic

import (
	"fgame/fgame/cross/arenapvp/pbutil"
	"fgame/fgame/game/scene/scene"
)

//复活
func Reborn(pl scene.Player) {
	if pl.IsRobot() {
		reliveTime := pl.GetArenapvpReliveTimes() + 1
		//扣除复活次数
		pl.SetArenapvpReliveTimes(reliveTime)
	} else {
		isMsg := pbutil.BuildISArenapvpRelive()
		pl.SendMsg(isMsg)
	}
}

//重置对战状态
func ResetBattleInfo(pl scene.Player) {
	if pl.IsRobot() {
		//重置复活次数复活次数
		pl.SetArenapvpReliveTimes(0)
	} else {
		isMsg := pbutil.BuildISArenapvpResetReliveTimes()
		pl.SendMsg(isMsg)
	}

	pl.ResetHP()
}

// //机器人退出
// func RobotExit(pl scene.RobotPlayer) {
// 	arenapvp.GetArenapvpService().ArenapvpMemeberExit(pl)
// 	robotlogic.RemoveRobot(pl)
// }
