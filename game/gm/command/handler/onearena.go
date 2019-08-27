package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	onearenalogic "fgame/fgame/game/onearena/logic"
	"fgame/fgame/game/onearena/onearena"
	"fgame/fgame/game/onearena/pbutil"
	playeronearena "fgame/fgame/game/onearena/player"
	onearenatypes "fgame/fgame/game/onearena/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeOneArena, command.CommandHandlerFunc(handleOneArena))
}

//进入灵池争夺
func handleOneArena(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理灵池争夺")
	pl := p.(player.Player)
	if len(c.Args) < 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	posStr := c.Args[1]
	levelt, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelt,
				"error": err,
			}).Warn("gm:处理灵池争夺,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	post, err := strconv.ParseInt(posStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"pos":   post,
				"error": err,
			}).Warn("gm:处理灵池争夺,pos不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeLingChiFighting {
		return
	}

	level := onearenatypes.OneArenaLevelType(levelt)
	pos := int32(post)
	manager := pl.GetPlayerDataManager(types.PlayerOneArenaDataManagerType).(*playeronearena.PlayerOneArenaDataManager)
	flag := manager.IsValid(level, pos)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
			"pos":      pos,
		}).Warn("onearena:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	oneArenaObj := manager.GetOneArena()
	curLevel := oneArenaObj.Level
	if level == curLevel {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
			"pos":      pos,
		}).Warn("onearena:不能同级抢夺")
		playerlogic.SendSystemMessage(pl, lang.OneArenaRobSameLevel)
		return
	}

	if level < curLevel {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
			"pos":      pos,
		}).Warn("onearena:不能抢夺低级灵池")
		playerlogic.SendSystemMessage(pl, lang.OneArenaRobLowLevel)
		return
	}
	if level > curLevel && level-curLevel != 1 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
			"pos":      pos,
		}).Warn("onearena:不能越级抢夺")
		playerlogic.SendSystemMessage(pl, lang.OneArenaGoGrab)
		return
	}

	flag = manager.IsRobCoolTime(level, pos)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"level":    level,
			"pos":      pos,
		}).Warn("onearena:您当前处于抢夺冷却时间中,请等冷却时间结束后再来!")
		playerlogic.SendSystemMessage(pl, lang.OneArenaCoolTime)
		return
	}

	ownerId, ownerName, codeResult := onearena.GetOneArenaService().OneArenaRob(oneArenaObj, level, pos)
	if codeResult != onearenatypes.OneArenaRobCodeTypeIsRobbing {
		//传送玩家
		onearenalogic.PlayerEnterOneArena(pl, ownerId, ownerName, level, pos)
	}
	scOneArenaRob := pbutil.BuildSCOneArenaRob(int32(codeResult))
	pl.SendMsg(scOneArenaRob)
	return
}
