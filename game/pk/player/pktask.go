package player

// import (
// 	"fgame/fgame/game/player"
// 	playertypes "fgame/fgame/game/player/types"
// 	"time"
// )

// const (
// 	pkTaskTime = time.Second * 5
// )

// type pkTask struct {
// 	p player.Player
// }

// func (t *pkTask) Run() {
// 	manager := t.p.GetPlayerDataManager(playertypes.PlayerPkDataManagerType).(*PlayerPkDataManager)
// 	manager.refresh(true)
// }

// func (t *pkTask) ElapseTime() time.Duration {
// 	return pkTaskTime
// }

// func newPkTask(p player.Player) *pkTask {
// 	t := &pkTask{
// 		p: p,
// 	}
// 	return t
// }
