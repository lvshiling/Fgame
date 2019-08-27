package rank

import (
	fgamedb "fgame/fgame/core/db"
	fgameredis "fgame/fgame/core/redis"
	"fgame/fgame/core/runner"
	"fgame/fgame/rank/rank/rank"
	"time"
)

// type rankModule struct {
// 	r runner.Runner
// }

// func (m *rankModule) InitTemplate() (err error) {

// 	return
// }
// func (m *rankModule) Init() (err error) {
// 	ds := global.GetGame().GetDB()
// 	rs := global.GetGame().GetRedisService()

// 	err = rank.Init(ds, rs)
// 	if err != nil {
// 		return
// 	}
// 	m.r = runner.NewRunner(time.Minute)
// 	m.r.AddTask(rank.GetRankService())

// 	return
// }

// func (m *rankModule) Start() error {
// 	err := rank.GetRankService().Star()
// 	if err != nil {
// 		return err
// 	}
// 	m.r.Start()
// 	return nil
// }

// func (m *rankModule) Stop() {
// 	m.r.Stop()
// }

// func (m *rankModule) String() string {
// 	return "rank"
// }

// var (
// 	m = &rankModule{}
// )

// func init() {
// 	module.Register(m)
// }

var (
	r runner.GoRunner
)

func Init(db fgamedb.DBService, rs fgameredis.RedisService) (err error) {
	err = rank.Init(db, rs)
	if err != nil {
		return
	}

	r = runner.NewGoRunner("rank", rank.GetRankService().Heartbeat, time.Minute)
	return
}

func Start() (err error) {
	err = rank.GetRankService().Star()
	if err != nil {
		return err
	}
	r.Start()
	return
}

func Stop() {
	r.Stop()
}
