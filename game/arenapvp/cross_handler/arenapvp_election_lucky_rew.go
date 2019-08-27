package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_ELECTION_LUCKY_REW_TYPE), dispatch.HandlerFunc(handleArenapvpElectionLuckyRew))
}

//处理跨服pvp海选幸运奖励
func handleArenapvpElectionLuckyRew(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenapvpElectionLuckyRew(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("pvp:处理跨服pvp海选幸运奖励,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("pvp:处理跨服pvp海选幸运奖励,完成")
	return nil

}

//pvp海选幸运奖励
func arenapvpElectionLuckyRew(pl player.Player) (err error) {
	pvpConstantTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpConstantTemplate()
	rewMap := pvpConstantTemp.GetLuckyItemMap()

	rewNameStr := chatlogic.FormatMailKeyWordNoticeStr(lang.GetLangService().ReadLang(lang.ArenapvpLuckyRewNameText))
	title := lang.GetLangService().ReadLang(lang.ArenapvpLuckyRewMailTitle)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpLuckyRewMailContent), rewNameStr)
	emaillogic.AddEmail(pl, title, content, rewMap)
	return
}
