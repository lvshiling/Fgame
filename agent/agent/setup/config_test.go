package setup_test

import (
	. "fgame/fgame/agent/agent/setup"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/center/center"
	chargecharge "fgame/fgame/game/charge/charge"
	couponcoupon "fgame/fgame/game/coupon/coupon"
	gamelog "fgame/fgame/game/log/log"
	gameserver "fgame/fgame/game/server"
	"fmt"
	"testing"
)

const (
	tpl = "./config/game.gotmpl"
)

func TestSetupConfig(t *testing.T) {
	options := &gameserver.GameServerOptions{}
	options.Debug = true
	options.Game = &gameserver.GameOptions{}
	options.Game.Redis = &coreredis.RedisConfig{}
	options.Game.Db = &coredb.DbConfig{}
	options.Game.Center = &center.CenterConfig{}
	options.Game.Log = &gamelog.LogConfig{}
	options.Game.Remote = &gameserver.RemoteServerOptions{}
	options.Game.Register = &gameserver.RegisterServerOptions{}
	options.Game.Charge = &chargecharge.ChargeConfig{}
	options.Game.Coupon = &couponcoupon.CouponConfig{}

	options.Server = &gameserver.ServerOptions{}
	content, err := SetupConfig(tpl, options)
	if err != nil {
		t.Errorf("生成配置错误,[%s]", err.Error())
		return
	}
	fmt.Println(content)
}
