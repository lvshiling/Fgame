package handler

import (
	"fgame/fgame/core/session"
	"fgame/fgame/login/model"
	accounttypes "fgame/fgame/login/types"

	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/game/global"
	"fgame/fgame/game/misc/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_REAL_NAME_AUTH_TYPE), dispatch.HandlerFunc(handleRealNameAuth))
}

//实名认证
func handleRealNameAuth(s session.Session, msg interface{}) error {
	log.Debug("misc:实名认证")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csRealNameAuth := msg.(*uipb.CSRealNameAuth)
	code := csRealNameAuth.GetCode()
	realName := csRealNameAuth.GetName()
	phoneNum := csRealNameAuth.GetPhoneNum()
	card := csRealNameAuth.GetCardNum()
	err := realNameAuth(tpl, phoneNum, code, realName, card)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"phone":    phoneNum,
				"code":     code,
				"card":     card,
				"name":     realName,
				"error":    err,
			}).Error("misc:实名认证,错误")

		return err
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"phone":    phoneNum,
			"code":     code,
			"card":     card,
			"name":     realName,
		}).Debug("misc:实名认证,完成")

	return nil
}

func realNameAuth(pl player.Player, phoneNum string, code string, name string, card string) (err error) {
	if pl.GetRealNameState() != accounttypes.RealNameStateNone {

		playerlogic.SendSystemMessage(pl, lang.MiscRealNameAuthAlready)

		return
	}
	//修改状态
	m := &model.User{}
	err = global.GetGame().GetDB().DB().Find(m, "id=?", pl.GetUserId()).Error
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	state := accounttypes.RealNameStateUp18
	m.RealNameState = int32(state)
	m.UpdateTime = now
	err = global.GetGame().GetDB().DB().Save(m).Error
	if err != nil {
		return
	}
	pl.RealNameAuth(state)
	scRealNameAuth := pbutil.BuildSCRealNameAuth(state)
	pl.SendMsg(scRealNameAuth)
	return
}
