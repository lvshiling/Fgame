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
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	soulruinslogic "fgame/fgame/game/soulruins/logic"
	"fgame/fgame/game/soulruins/pbutil"
	playersoulruins "fgame/fgame/game/soulruins/player"
	"fgame/fgame/game/soulruins/soulruins"
	soulruinstypes "fgame/fgame/game/soulruins/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOULRUINS_NEXT_LEVEL_TYPE), dispatch.HandlerFunc(handleSoulRuinsNextLevel))
}

//处理下一关
func handleSoulRuinsNextLevel(s session.Session, msg interface{}) (err error) {
	log.Debug("soulruins:处理下一关消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = soulRuinsNextLevel(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("soulruins:处理下一关消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soulruins:处理下一关消息完成")
	return nil

}

//下一关信息的逻辑
func soulRuinsNextLevel(pl player.Player) (err error) {
	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeFuBenSoulRuins {
		return
	}
	sd := s.SceneDelegate()
	sceneData, ok := sd.(*soulruinslogic.SoulRuinsSceneData)
	if !ok {
		return
	}
	soulRuinsTemplate := sceneData.GetSoulRuinsTemplate()
	chapter := soulRuinsTemplate.Chapter
	typ := soulruinstypes.SoulRuinsType(soulRuinsTemplate.Type)
	level := soulRuinsTemplate.Level

	nextChapter, nextTyp, nextLevel, flag := soulruins.GetSoulRuinsService().GetSoulRuinsNextLevel(chapter, typ, level)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  nextChapter,
				"typ":      nextTyp,
				"level":    nextLevel,
			}).Warn("soulruins:当前关卡已是最高关卡,不存在下一关")
		playerlogic.SendSystemMessage(pl, lang.SoulRuinsNoExistNextLevel)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	//判断挑战次数是否足够
	flag = manager.HasEnoughChallengeNum(1)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"level":    level,
			}).Warn("soulruins:挑战次数不足")
		playerlogic.SendSystemMessage(pl, lang.SoulRuinsChallengeNumNotEnough)
		return
	}

	//判断前置通关
	flag = manager.IfPreSoulRuinsPassed(nextChapter, nextTyp, nextLevel)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"level":    level,
			}).Warn("soulruins:请先通关前置关卡")
		playerlogic.SendSystemMessage(pl, lang.SoulRuinsNotReachPreLevel)
		return
	}

	//消耗挑战次数
	manager.UseChallengeNum(1)
	flag = soulruinslogic.PlayerEnterSoulRuins(pl, nextChapter, nextTyp, nextLevel)
	if !flag {
		panic(fmt.Errorf("soulruins: soulRuinsNextLevel should be ok"))
	}

	numObj := manager.GetSoulRuinsNum()
	scSoulRuinsChallenge := pbutil.BuildSCSoulRuinsChallenge(numObj)
	pl.SendMsg(scSoulRuinsChallenge)
	return
}
