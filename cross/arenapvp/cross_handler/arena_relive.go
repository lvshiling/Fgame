package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	arenapvpscene "fgame/fgame/cross/arenapvp/scene"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_ARENAPVP_RELIVE_TYPE), dispatch.HandlerFunc(handleArenapvpRelive))
}

//竞技场复活
func handleArenapvpRelive(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:竞技场复活")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	siMsg := msg.(*crosspb.SIArenapvpRelive)
	err = arenapvpRelive(tpl, siMsg.GetSuccess())
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arenapvp:竞技场复活,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arenapvp:竞技场复活,完成")
	return nil
}

//竞技场获胜
func arenapvpRelive(pl *player.Player, suceess bool) (err error) {
	if suceess {
		pos := pl.GetScene().MapTemplate().RandomPosition()
		_, ok := pl.GetScene().SceneDelegate().(arenapvpscene.ArenapvpBattleSceneData)
		if ok {
			pos = pl.GetPos()
		}
		pl.Reborn(pos)
	} else {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Info("arenapvp:pvp竞技场复活,失败")
		//TODO 是否需要退出场景
	}
	return
}
