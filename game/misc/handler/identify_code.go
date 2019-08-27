package handler

import (
	"fgame/fgame/core/session"

	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/game/misc/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	accounttypes "fgame/fgame/login/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GET_IDENTIFY_CODE_TYPE), dispatch.HandlerFunc(handleGetIdentifyCode))

}

//获取验证码
func handleGetIdentifyCode(s session.Session, msg interface{}) error {
	log.Debug("misc:获取验证码")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csGetIndentifyCode := msg.(*uipb.CSGetIdentifyCode)
	phoneNum := csGetIndentifyCode.GetPhoneNum()
	err := getIdentifyCode(tpl, phoneNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("misc:获取验证码,错误")

		return err
	}
	log.Debug("misc:获取验证码,完成")
	return nil
}

func getIdentifyCode(pl player.Player, phoneNum string) (err error) {
	if pl.GetRealNameState() != accounttypes.RealNameStateNone {

		playerlogic.SendSystemMessage(pl, lang.MiscRealNameAuthAlready)

		return
	}
	//发送验证码
	scGetIdentifyCode := pbutil.BuildSCGetIdentifyCode()
	pl.SendMsg(scGetIdentifyCode)
	return
}
