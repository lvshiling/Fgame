package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/cross/tulong/pbutil"
	tulongscene "fgame/fgame/cross/tulong/scene"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	playerlogic "fgame/fgame/game/player/logic"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TULONG_COLLECT_TYPE), dispatch.HandlerFunc(handleTuLongCollect))
}

//处理屠龙采集
func handleTuLongCollect(s session.Session, msg interface{}) (err error) {
	log.Debug("tulong:处理屠龙采集")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	csTuLongCollect := msg.(*uipb.CSTuLongCollect)
	npcId := csTuLongCollect.GetNpcId()
	err = tuLongCollectEgg(tpl, npcId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Error("tulong:处理屠龙采集,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("tulong:处理选择屠龙")
	return nil

}

//处理屠龙采集
func tuLongCollectEgg(pl *player.Player, npcId int64) (err error) {
	s := pl.GetScene()
	if s == nil || s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossTuLong {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tulong:不在屠龙场景")
		return
	}

	//判断是否是管理层
	if pl.GetMemPos() == alliancetypes.AlliancePositionMember {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("tulong:你不是仙盟管理,无法开启龙蛋")
		playerlogic.SendSystemMessage(pl, lang.TuLongCollectNoMengZhu)
		return
	}

	//采集
	sd := s.SceneDelegate().(tulongscene.TuLongSceneData)
	// flag := sd.HasedCollectEgg(pl.GetAllianceId())
	// if flag {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"npcId":    npcId,
	// 		}).Warn("tulong:已经成功采集过龙蛋了")
	// 	playerlogic.SendSystemMessage(pl, lang.TuLongHasedCollectEgg)
	// 	return
	// }

	eggInfo, flag := sd.SmallEggShouldExist(npcId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("tulong:龙蛋应该存在的")
		playerlogic.SendSystemMessage(pl, lang.TuLongCollectEggShouldExist)
		return
	}
	eggNpc := eggInfo.GetEggNpc()

	distance := coreutils.Distance(eggNpc.GetPosition(), pl.GetPos())
	collectDistance := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCollectDistance)) / float64(1000)
	if distance > float64(collectDistance) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Warn("tulong:不在采集范围内")
		playerlogic.SendSystemMessage(pl, lang.CommonCollectNoDistance)
		return
	}

	_, flag = sd.IfCanCollectEgg(npcId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("tulong:其它玩家正在采集")
		playerlogic.SendSystemMessage(pl, lang.TuLongOtherCollect)
		return
	}
	flag = sd.CollectEgg(pl.GetId(), pl.GetAllianceId(), npcId)
	if !flag {
		panic(fmt.Errorf("tulong:采集应该成功"))
	}
	scTuLongCollect := pbutil.BuildSCTuLongCollect(npcId)
	pl.SendMsg(scTuLongCollect)
	return
}
