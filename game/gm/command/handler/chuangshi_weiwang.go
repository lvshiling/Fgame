package handler

// import (
// 	"fgame/fgame/common/lang"
// 	chuangshilogic "fgame/fgame/game/chuangshi/logic"
// 	playerchuangshi "fgame/fgame/game/chuangshi/player"
// 	"fgame/fgame/game/gm/command"
// 	gmcommandtypes "fgame/fgame/game/gm/command/types"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	playertypes "fgame/fgame/game/player/types"
// 	"fgame/fgame/game/scene/scene"
// 	"strconv"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	command.Register(gmcommandtypes.CommandTypeChuangShiWeiWang, command.CommandHandlerFunc(handleChuangShiWeiWang))
// 	command.Register(gmcommandtypes.CommandTypeChuangShiJifen, command.CommandHandlerFunc(handleChuangShiJifen))
// }

// func handleChuangShiWeiWang(p scene.Player, c *command.Command) (err error) {
// 	pl := p.(player.Player)
// 	if len(c.Args) < 1 {
// 		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
// 		return
// 	}

// 	weiWangStr := c.Args[0]
// 	weiWang, err := strconv.ParseInt(weiWangStr, 10, 64)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"id":         pl.GetId(),
// 				"weiWangStr": weiWangStr,
// 				"error":      err,
// 			}).Warn("gm:处理创世威望,类型weiWang不是数字")
// 		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
// 		err = nil
// 		return
// 	}

// 	chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangShiManager.GMSetWeiWang(weiWang)

// 	chuangshilogic.SendPlayerChuangShiInfo(pl)
// 	return
// }

// func handleChuangShiJifen(p scene.Player, c *command.Command) (err error) {
// 	pl := p.(player.Player)
// 	if len(c.Args) < 1 {
// 		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
// 		return
// 	}

// 	jifenStr := c.Args[0]
// 	jifen, err := strconv.ParseInt(jifenStr, 10, 64)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"id":       pl.GetId(),
// 				"jifenStr": jifenStr,
// 				"error":    err,
// 			}).Warn("gm:处理创世积分,类型jifen不是数字")
// 		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
// 		err = nil
// 		return
// 	}

// 	chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangShiManager.GMSetJifen(jifen)

// 	chuangshilogic.SendPlayerChuangShiInfo(pl)
// 	return
// }
