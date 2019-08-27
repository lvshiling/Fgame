package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/chuangshi/chuangshi"
	chuangshilogic "fgame/fgame/game/chuangshi/logic"
	"fgame/fgame/game/chuangshi/pbutil"
	playerchuangshi "fgame/fgame/game/chuangshi/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	loadYuGao(pl)

	// loadShenWangBaoMing(pl)

	// loadShenWangVote(pl)

	// loadChengFangJianShe(pl)

	//
	sendPlayerInfo(pl)
	return
}

// 加载创世预告
func loadYuGao(pl player.Player) {
	playerChuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
	isJoin := playerChuangShiManager.IsJoin()
	num := chuangshi.GetChuangShiService().GetBaoMingChuangShiPlayerNum()

	scMsg := pbutil.BuildSCChuangShiYuGaoInfo(isJoin, num)
	pl.SendMsg(scMsg)
}

// // 同步报名信息
// func loadShenWangBaoMing(pl player.Player) {
// 	chuangshilogic.ShenWangSignUpResult(pl)
// }

// //同步投票信息
// func loadShenWangVote(pl player.Player) {
// 	chuangshilogic.ShenWangVoteResult(pl)
// }

// //同步建设信息
// func loadChengFangJianShe(pl player.Player) {
// 	chuangshilogic.ChengFangJianSheResult(pl)
// }

//发送玩家信息
func sendPlayerInfo(pl player.Player) {
	chuangshilogic.SendPlayerChuangShiInfo(pl)
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
 