package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	bagualogic "fgame/fgame/game/bagua/logic"
	"fgame/fgame/game/bagua/pbutil"
	playerbagua "fgame/fgame/game/bagua/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BAGUA_TOKILL_TYPE), dispatch.HandlerFunc(handleBaGuaToKill))
}

//处理前往击杀信息
func handleBaGuaToKill(s session.Session, msg interface{}) (err error) {
	log.Debug("bagua:处理获取前往击杀消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = baGuaToKill(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("bagua:处理获取前往击杀消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("bagua:处理获取前往击杀消息完成")
	return nil

}

//获取前往击杀界面信息的逻辑
func baGuaToKill(pl player.Player) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeBaGuaMiJing {
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

	curlevel := manager.GetLevel()
	nextLevel := curlevel + 1
	//调用场景接口
	_, flag = bagualogic.PlayerEnterBaGua(pl, 0, nextLevel)
	if !flag {
		panic(fmt.Errorf("bagua: baGuaToKill  should be ok"))
	}
	scBaGuaToKill := pbutil.BuildSCBaGuaToKill(nextLevel)
	pl.SendMsg(scBaGuaToKill)
	return
}
