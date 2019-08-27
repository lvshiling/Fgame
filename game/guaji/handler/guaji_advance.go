package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/guaji/guaji"
	"fgame/fgame/game/guaji/pbutil"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GUA_JI_ADVANCE_LIST_TYPE), dispatch.HandlerFunc(handleGuaJiAdvance))
}

//处理挂机
func handleGuaJiAdvance(s session.Session, msg interface{}) (err error) {
	log.Info("guaji:处理进阶挂机")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = guaJiAdvance(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),

				"error": err,
			}).Error("guaji:处理进阶挂机,错误")

		return err
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("guaji:处理进阶挂机,完成")
	return nil
}

//挂机
func guaJiAdvance(pl player.Player) (err error) {
	if !pl.IsGuaJiPlayer() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("guaji:处理挂机,不是挂机玩家")
		playerlogic.SendSystemMessage(pl, lang.GuaJiNoGuaJiPlayer)
		return
	}
	advanceMap := make(map[guajitypes.GuaJiAdvanceType]int32)
	for typ, _ := range guajitypes.GetGuaJiAdvanceMap() {
		h := guaji.GetGuaJiAdvanceGetHandler(typ)
		if h == nil {
			continue
		}
		advanceId := h.GetAdvance(pl, typ)
		advanceMap[typ] = advanceId
	}
	scGuaJiAdvanceList := pbutil.BuildSCGuaJiAdvanceList(advanceMap)
	pl.SendMsg(scGuaJiAdvanceList)
	return
}
