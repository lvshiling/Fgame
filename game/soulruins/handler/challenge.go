package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	soulruinslogic "fgame/fgame/game/soulruins/logic"
	soulruinstypes "fgame/fgame/game/soulruins/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOULRUINS_CHALLENGE_TYPE), dispatch.HandlerFunc(handleSoulRuinsChallenge))
}

//处理帝陵遗迹挑战信息
func handleSoulRuinsChallenge(s session.Session, msg interface{}) (err error) {
	log.Debug("soulruins:处理获取帝陵遗迹挑战消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulRuinsChallenge := msg.(*uipb.CSSoulRuinsChallenge)
	chapterInfo := csSoulRuinsChallenge.GetChapterInfo()
	chapter := chapterInfo.GetChapter()
	typ := chapterInfo.GetTyp()
	level := csSoulRuinsChallenge.GetLevel()

	err = soulRuinsChallenge(tpl, chapter, soulruinstypes.SoulRuinsType(typ), level)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"chapter":  chapter,
				"typ":      typ,
				"level":    level,
				"error":    err,
			}).Error("soulruins:处理获取帝陵遗迹挑战消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"chapter":  chapter,
			"typ":      typ,
			"level":    level,
		}).Debug("soulruins:处理获取帝陵遗迹挑战消息完成")
	return nil

}

//获取帝陵遗迹挑战界面信息的逻辑
func soulRuinsChallenge(pl player.Player, chapter int32, typ soulruinstypes.SoulRuinsType, level int32) (err error) {
	err = soulruinslogic.SoulRuinsChallenge(pl, chapter, typ, level)
	return
}
