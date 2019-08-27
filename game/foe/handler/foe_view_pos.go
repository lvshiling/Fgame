package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/foe/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FOE_VIEW_POS_TYPE), dispatch.HandlerFunc(handleFoeViewPos))
}

//处理查看位置
func handleFoeViewPos(s session.Session, msg interface{}) error {
	log.Debug("foe:处理仇人查看位置")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csFoeViewPos := msg.(*uipb.CSFoeViewPos)
	foeId := csFoeViewPos.GetFoeId()
	foeName := csFoeViewPos.GetFoeName()

	err := foeViewPos(tpl, foeId, foeName)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"foeId":    foeId,
				"error":    err,
			}).Error("foe:处理仇人查看位置,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("foe:处理仇人查看位置,完成")
	return nil

}

//处理仇人查看位置
func foeViewPos(pl player.Player, foeId int64, foeName string) (err error) {
	// manager := pl.GetPlayerDataManager(playertypes.PlayerFoeDataManagerType).(*playerfoe.PlayerFoeDataManager)
	// flag := manager.IsFoe(foeId)
	// if !flag {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"foeId":    foeId,
	// 		}).Warn("foe:参数无效")
	// 	playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
	// 	return
	// }

	spl := player.GetOnlinePlayerManager().GetPlayerById(foeId)
	if spl == nil {
		scFoeViewPosNoOnline := pbutil.BuildSCFoeViewPosNoOnline(foeId, foeName)
		pl.SendMsg(scFoeViewPosNoOnline)
		return
	}

	isCross, name, mapId, pos := foeGetMap(spl)
	if !isCross {
		scFoeViewPos := pbutil.BuildSCFoeViewPos(foeId, foeName, name, mapId, pos)
		pl.SendMsg(scFoeViewPos)
	} else {
		scFoeViewCross := pbutil.BuildSCFoeViewCross(foeId, foeName, name)
		pl.SendMsg(scFoeViewCross)
	}
	return
}

func foeGetMap(pl player.Player) (isCross bool, name string, mapId int32, pos coretypes.Position) {
	if pl.IsCross() {
		isCross = true
		crossType := pl.GetCrossType()
		name = crossType.String()
	} else {
		mapId = pl.GetMapId()
		to := template.GetTemplateService().Get(int(mapId), (*gametemplate.MapTemplate)(nil))
		if to != nil {
			mapTemplate := to.(*gametemplate.MapTemplate)
			name = mapTemplate.Name
		}
		pos = pl.GetPos()
	}
	return
}
