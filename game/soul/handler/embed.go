package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	soullogic "fgame/fgame/game/soul/logic"
	soultypes "fgame/fgame/game/soul/types"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_SOUL_EMBED_TYPE), dispatch.HandlerFunc(handleSoulEmbed))

}

//处理帝魂镶嵌信息
func handleSoulEmbed(s session.Session, msg interface{}) (err error) {
	log.Debug("soul:处理帝魂镶嵌信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulEmbed := msg.(*uipb.CSSoulEmbed)
	soulTag := csSoulEmbed.GetSoulTag()

	err = soulEmbed(tpl, soultypes.SoulType(soulTag))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
				"error":    err,
			}).Error("soul:处理帝魂镶嵌信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soul:处理帝魂镶嵌信息完成")
	return nil

}

//帝魂镶嵌逻辑
func soulEmbed(pl player.Player, soulTag soultypes.SoulType) (err error) {
	return soullogic.HandleSoulEmbed(pl, soulTag)
}
