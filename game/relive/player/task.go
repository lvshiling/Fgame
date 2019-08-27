package player

//TODO 控制同步时间
// const (
// 	reliveTaskTime = time.Second * 1
// )

// type reliveTask struct {
// 	p player.Player
// }

// func (t *reliveTask) Run() {
// 	m := t.p.GetPlayerDataManager(playertypes.PlayerReliveDataManagerType).(*PlayerReliveDataManager)
// 	flag := m.Refresh()
// 	if !flag {
// 		return
// 	}
// 	//发送刷新事件

// 	gameevent.Emit(reliveeventtypes.EventTypePlayerReliveRefresh, t.p, nil)
// }

// func (t *reliveTask) ElapseTime() time.Duration {
// 	return reliveTaskTime
// }

// func CreateReliveTask(p player.Player) *reliveTask {
// 	t := &reliveTask{
// 		p: p,
// 	}
// 	return t
// }
