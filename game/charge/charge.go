package charge

import (
	"fgame/fgame/game/charge/charge"
	"fgame/fgame/game/charge/dao"
	chargetemplate "fgame/fgame/game/charge/template"
	"fgame/fgame/game/global"

	//注册管理器
	_ "fgame/fgame/game/charge/event/listener"
	_ "fgame/fgame/game/charge/guaji"
	_ "fgame/fgame/game/charge/handler"
	_ "fgame/fgame/game/charge/player"
	_ "fgame/fgame/game/charge/use"
)

func Init(cfg *charge.ChargeConfig) (err error) {
	if err = chargetemplate.Init(); err != nil {
		return
	}
	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	err = charge.Init(cfg)
	if err != nil {
		return
	}
	return
}
