package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	realmlogic "fgame/fgame/game/realm/logic"
	"fgame/fgame/game/realm/pbutil"
	playerrealm "fgame/fgame/game/realm/player"
	"fgame/fgame/game/realm/realm"
	scenelogic "fgame/fgame/game/scene/logic"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_REALM_NEXT_TYPE), dispatch.HandlerFunc(handleRealmNextLevel))
}

//处理下一关
func handleRealmNextLevel(s session.Session, msg interface{}) (err error) {
	log.Debug("realm:处理下一关消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = realmNextLevel(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("realm:处理下一关消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("realm:处理下一关消息完成")
	return nil

}

//下一关信息的逻辑
func realmNextLevel(pl player.Player) (err error) {
	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeTianJieTa {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	flag := manager.IfFullLevel()
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("realm:境界已达最高层")
		playerlogic.SendSystemMessage(pl, lang.RealmReachLimit)
		return
	}

	sd := s.SceneDelegate()
	sceneData, ok := sd.(*realmlogic.TianJieTaSceneData)
	if !ok {
		return
	}
	sceneData.OnSetNextLevel()

	curlevel := manager.GetTianJieTaLevel()
	nextLevel := curlevel + 1
	curSpouseId, pairFlag := realm.GetRealmRankService().IsExistPairKill(pl.GetId())
	spouseId := curSpouseId
	spl := s.GetPlayer(spouseId)
	if spl == nil {
		spouseId = 0
		if pairFlag {
			realm.GetRealmRankService().PairSpouseExit(spouseId)
		}
	}

	if pairFlag && spouseId != 0 && spl != nil { //多人副本
		s, flag := realmlogic.PlayerEnterTianJieTa(pl, spouseId, nextLevel)
		if !flag {
			panic(fmt.Errorf("realm: realmNextLevel  should be ok"))
		}
		if s == nil {
			return nil
		}
		scenelogic.PlayerEnterSingleFuBenScene(spl, s)

	} else { //单人副本
		_, flag = realmlogic.PlayerEnterTianJieTa(pl, 0, nextLevel)
		if !flag {
			panic(fmt.Errorf("realm: realmNextLevel  should be ok"))
		}
	}
	scRealmNext := pbutil.BuildSCRealmNext(nextLevel)
	pl.SendMsg(scRealmNext)
	return
}
