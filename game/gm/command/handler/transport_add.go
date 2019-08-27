package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	transportationlogic "fgame/fgame/game/transportation/logic"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	transportationtemplate "fgame/fgame/game/transportation/template"
	"fgame/fgame/game/transportation/transpotation"
	transportationtypes "fgame/fgame/game/transportation/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeTransportAdd, command.CommandHandlerFunc(handleTransportAdd))
}

func handleTransportAdd(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:生成镖车")
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	typ := c.Args[0]

	typInt, err := strconv.ParseInt(typ, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"error":  err,
				"typInt": typInt,
			}).Warn("gm:生成镖车,typ不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	ttyp := transportationtypes.TransportationType(typInt)
	if !ttyp.Valid() {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"error":  err,
				"typInt": typInt,
			}).Warn("gm:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	err = transportAdd(pl, ttyp)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Error("gm:生成镖车,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:生成镖车,完成")
	return
}

func transportAdd(p scene.Player, typ transportationtypes.TransportationType) (err error) {
	pl := p.(player.Player)
	moveTemp := transportationtemplate.GetTransportationTemplateService().GetTransportationMoveTemplateFirst()
	s := scene.GetSceneService().GetWorldSceneByMapId(moveTemp.MapId)
	if s == nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"error":  err,
				"typInt": typ,
			}).Warn("gm:场景不存在")
		return
	}

	var biaoChe *biaochenpc.BiaocheNPC
	pos := moveTemp.GetPosition()
	if typ == transportationtypes.TransportationTypeAlliance {
		_, biaoChe, err = alliance.GetAllianceService().AddTransportation(pl.GetId())
	} else {
		biaoChe, err = transpotation.GetTransportService().AddPersonalTransportation(pl.GetId(), pl.GetName(), pl.GetAllianceId(), typ)
	}
	if err != nil {
		return
	}
	transportationlogic.AddBiaoChe(pl, biaoChe, s, pos)
	return

}
