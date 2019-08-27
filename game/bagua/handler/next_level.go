package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/bagua/bagua"
	bagualogic "fgame/fgame/game/bagua/logic"
	"fgame/fgame/game/bagua/pbutil"
	playerbagua "fgame/fgame/game/bagua/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BAGUA_NEXT_TYPE), dispatch.HandlerFunc(handleBaGuaNextLevel))
}

//处理下一关
func handleBaGuaNextLevel(s session.Session, msg interface{}) (err error) {
	log.Debug("bagua:处理下一关消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = baGuaNextLevel(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("bagua:处理下一关消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("bagua:处理下一关消息完成")
	return nil

}

//下一关信息的逻辑
func baGuaNextLevel(pl player.Player) (err error) {
	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeBaGuaMiJing {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
	flag := manager.IfFullLevel()
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("bagua:八卦秘境已达最高层")
		playerlogic.SendSystemMessage(pl, lang.BaGuaReachLimit)
		return
	}

	sd := s.SceneDelegate()
	sceneData, ok := sd.(*bagualogic.BaGuaSceneData)
	if !ok {
		return
	}
	sceneData.OnSetNextLevel()

	curlevel := manager.GetLevel()
	nextLevel := curlevel + 1
	curSpouseId, pairFlag := bagua.GetBaGuaService().IsExistPairKill(pl.GetId())
	spouseId := curSpouseId
	spl := s.GetPlayer(spouseId)
	if spl == nil {
		spouseId = 0
		if pairFlag {
			bagua.GetBaGuaService().PairSpouseExit(spouseId)
		}
	}

	if pairFlag && spouseId != 0 && spl != nil { //多人副本
		s, flag := bagualogic.PlayerEnterBaGua(pl, spouseId, nextLevel)
		if !flag {
			panic(fmt.Errorf("bagua: baGuaNextLevel  should be ok"))
		}
		if s == nil {
			return nil
		}
		scenelogic.PlayerEnterSingleFuBenScene(spl, s)

	} else { //单人副本
		_, flag = bagualogic.PlayerEnterBaGua(pl, 0, nextLevel)
		if !flag {
			panic(fmt.Errorf("bagua: baGuaNextLevel  should be ok"))
		}
	}
	scBaGuaNext := pbutil.BuildSCBaGuaNext(nextLevel)
	pl.SendMsg(scBaGuaNext)
	return
}
