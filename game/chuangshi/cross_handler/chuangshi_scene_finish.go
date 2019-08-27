package cross_handler

// TODO xzk27 修改为全阵营推送
// import (
// 	"fgame/fgame/common/codec"
// 	crosspb "fgame/fgame/common/codec/pb/cross"
// 	"fgame/fgame/common/dispatch"
// 	"fgame/fgame/common/lang"
// 	"fgame/fgame/core/session"
// 	chuangshitemplate "fgame/fgame/game/chuangshi/template"
// 	emaillogic "fgame/fgame/game/email/logic"
// 	"fgame/fgame/game/player"
// 	"fgame/fgame/game/processor"
// 	gamesession "fgame/fgame/game/session"
// 	"fmt"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_CHUANGSHI_SCENE_FINISH_TYPE), dispatch.HandlerFunc(handleChuangShiSceneFinish))
// }

// //处理跨服攻城结束
// func handleChuangShiSceneFinish(s session.Session, msg interface{}) (err error) {
// 	log.Debug("chuangshi:处理跨服攻城结束")

// 	gcs := gamesession.SessionInContext(s.Context())
// 	pl := gcs.Player()
// 	tpl := pl.(player.Player)

// 	isMsg := msg.(*crosspb.ISChuangShiSceneFinish)
// 	isWin := isMsg.GetWin()
// 	err = chuangShiSceneFinish(tpl, isWin)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"isWin":    isWin,
// 				"err":      err,
// 			}).Error("chuangshi:处理跨服攻城结束,错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Debug("chuangshi:处理跨服攻城结束,完成")
// 	return nil
// }

// //攻城结束
// func chuangShiSceneFinish(pl player.Player, isWin bool) (err error) {

// 	awardTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarAwardTemp(isWin)
// 	attatchMap := awardTemp.GetRewItemMap()

// 	var contentLang lang.LangCode
// 	if isWin {
// 		contentLang = lang.ChuangShiSceneFinishWinMailContent
// 	} else {
// 		contentLang = lang.ChuangShiSceneFinishFailedMailContent
// 	}

// 	title := lang.GetLangService().ReadLang(lang.ChuangShiSceneFinishMailTitle)
// 	content := fmt.Sprintf(lang.GetLangService().ReadLang(contentLang), cityName)
// 	emaillogic.AddEmail(pl, title, content, attatchMap)

// 	return
// }
