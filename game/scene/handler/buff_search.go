package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	bufftemplate "fgame/fgame/game/buff/template"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BUFF_SEARCH_TYPE), dispatch.HandlerFunc(handleBuffSearch))
}

//处理buff列表包
func handleBuffSearch(s session.Session, msg interface{}) error {
	log.Debugln("scene:处理对象buff列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	csBuffSearch := msg.(*uipb.CSBuffSearch)
	buffId := csBuffSearch.GetBuffId()
	buffSearch(tpl, buffId)

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("scene:处理对象buff列表,完成")
	return nil
}

func buffSearch(pl scene.Player, buffId int32) {
	result := false
	buffTemp := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if buffTemp != nil {
		b := pl.GetBuff(buffTemp.Group)
		if b != nil {
			result = true
		}
	}
	scBuffSearch := pbutil.BuildSCBuffSearch(buffId, result)
	pl.SendMsg(scBuffSearch)
}
