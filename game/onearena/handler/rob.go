package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	onearenalogic "fgame/fgame/game/onearena/logic"
	"fgame/fgame/game/onearena/onearena"
	"fgame/fgame/game/onearena/pbutil"
	playeronearena "fgame/fgame/game/onearena/player"
	onearenatypes "fgame/fgame/game/onearena/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ONE_ARENA_ROB_TYPE), dispatch.HandlerFunc(handleOneArenaRob))
}

//处理灵池争夺信息
func handleOneArenaRob(s session.Session, msg interface{}) (err error) {
	log.Debug("onearena:处理获取灵池争夺消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csOneArenaRob := msg.(*uipb.CSOneArenaRob)
	level := csOneArenaRob.GetLevel()
	pos := csOneArenaRob.GetPos()
	err = oneArenaRob(tpl, onearenatypes.OneArenaLevelType(level), pos)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("onearena:处理获取灵池争夺消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("onearena:处理获取灵池争夺消息完成")
	return nil
}

//处理灵池争夺信息逻辑
func oneArenaRob(pl player.Player, level onearenatypes.OneArenaLevelType, pos int32) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeLingChiFighting {
		return
	}

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
