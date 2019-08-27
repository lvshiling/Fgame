package handler

// import (
// 	"fgame/fgame/common/codec"
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/game/chuangshi/pbutil"
// 	crosslogic "fgame/fgame/game/cross/logic"
// 	crosstypes "fgame/fgame/game/cross/types"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"
// 	"fmt"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_ENTER_CITY_TYPE), dispatch.HandlerFunc(handlerEnterCity))
// }

// //前往城池请求
// func handlerEnterCity(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理前往城池请求")

// 	pl := gamesession.SessionInContext(s.Context()).Player()
// 	tpl := pl.(player.Player)
// 	cs := msg.(*uipb.CSChuangShiEnterCity)
// 	cityId := cs.GetCityId()

// 	err = enterCity(tpl, cityId)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": tpl.GetId(),
// 				"err":      err,
// 			}).Error("chuangshi:处理前往城池请求，错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": tpl.GetId(),
// 		}).Debug("chuangshi:处理前往城池请求完成")

// 	return
// }

// func enterCity(pl player.Player, cityId int64) (err error) {
// 	playerlogic.CheckCanEnterScene(pl)

// 	argCityId := fmt.Sprintf("%d", cityId)
// 	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeChuangShi, argCityId)

// 	scMsg := pbutil.BuildSCChuangShiEnterCity(cityId)
// 	pl.SendMsg(scMsg)
// 	return
// }
