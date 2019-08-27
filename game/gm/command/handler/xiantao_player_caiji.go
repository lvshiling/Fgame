package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	xiantaologic "fgame/fgame/game/xiantao/logic"
	"fgame/fgame/game/xiantao/pbutil"
	playerxiantao "fgame/fgame/game/xiantao/player"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeXianTaoPlayerCaiJiTimes, command.CommandHandlerFunc(handleSetXianTaoPlayerCaiJiTimes))
}

func handleSetXianTaoPlayerCaiJiTimes(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	numStr := c.Args[0]
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"num":   num,
				"error": err,
			}).Warn("gm:处理仙桃玩家采集次数任务,类型num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}
	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeXianTaoDaHui {
		return
	}

	sd := s.SceneDelegate()
	xiantaoSd, ok := sd.(xiantaologic.XianTaoSceneData)
	if !ok {
		return
	}
	xiantaoSd.SetPlayerCollectCount(pl.GetId(), int32(num))

	xianTaoManager := pl.GetPlayerDataManager(types.PlayerXianTaoDataManagerType).(*playerxiantao.PlayerXianTaoDataManager)
	xianTaoObject := xianTaoManager.GetXianTaoObject()
	pCollectCount := xiantaoSd.GetPlayerCollectCount(pl.GetId())
	scMsg := pbutil.BuildSCXiantaoPlayerAttendChange(xianTaoObject, pCollectCount)
	pl.SendMsg(scMsg)
	return
}
