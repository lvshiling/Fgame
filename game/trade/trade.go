package trade

import (
	"fgame/fgame/core/runner"
	"fgame/fgame/game/global"
	"fgame/fgame/game/trade/dao"

	_ "fgame/fgame/game/trade/event/listener"
	_ "fgame/fgame/game/trade/handler"
	_ "fgame/fgame/game/trade/player"
	"fgame/fgame/game/trade/template"
	"fgame/fgame/game/trade/trade"
	"time"
)

const (
	tradeTimer = time.Second * 5
)

func Init(cfg *trade.TradeOptions) (err error) {
	err = template.Init()
	if err != nil {
		return
	}
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = trade.Init(cfg)
	if err != nil {
		return
	}
	r = runner.NewGoRunner("trade", trade.GetTradeService().Heartbeat, tradeTimer)
	return
}

var (
	r runner.GoRunner
)

func Start() {
	trade.GetTradeService().Start()
	r.Start()
	return
}

func Stop() {
	r.Stop()
	trade.GetTradeService().Stop()
	return
}
