package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	gemlogic "fgame/fgame/game/gem/logic"
	gemtypes "fgame/fgame/game/gem/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GEM_GAMBLE_TYPE), dispatch.HandlerFunc(handleGemGamble))
}

//处理赌石信息
func handleGemGamble(s session.Session, msg interface{}) (err error) {
	log.Debug("gem:处理赌石信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csGemGamble := msg.(*uipb.CSGemGamble)
	typ := csGemGamble.GetType()
	tenEven := csGemGamble.GetTenEven()

	err = gemGamble(tpl, gemtypes.GemGambleType(typ), tenEven)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"tenEven":  tenEven,
				"error":    err,
			}).Error("gem:处理赌石信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
			"tenEven":  tenEven,
		}).Debug("gem:处理赌石信息完成")
	return nil
}

//处理赌石信息逻辑
func gemGamble(pl player.Player, typ gemtypes.GemGambleType, tenEven bool) (err error) {
	return gemlogic.HandleGemGamble(pl, typ, tenEven)
}
