package handler

import (
	"fgame/fgame/client/player/pbutil"
	"fgame/fgame/client/player/player"
	"fgame/fgame/client/processor"
	clientsession "fgame/fgame/client/session"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	playertypes "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_SC_ENETER_SELECT_JOB_TYPE), dispatch.HandlerFunc(handleEnterSelectJob))
}

//进入选择角色界面
func handleEnterSelectJob(s session.Session, msg interface{}) (err error) {
	log.Debug("login:进入选择角色界面")
	cs := clientsession.SessionInContext(s.Context())
	pl := cs.Player().(*player.Player)
	err = playerEnterSelectJob(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.Id(),
				"err":      err,
			}).Error("login:进入选择角色界面")
		return
	}
	log.Debug("login:进入选择角色界面完成")
	return
}

func playerEnterSelectJob(pl *player.Player) (err error) {
	flag := pl.SelectRole()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.Id(),
				"state":    pl.CurrentState(),
			}).Warn("login:进入选择角色界面,失败")
		return
	}

	role := playertypes.RoleTypeKaiTian
	sex := playertypes.SexTypeMan
	name := pl.UserName()
	csSelectJob := pbutil.BuildCSSelectJob(role, sex, name)
	pl.SendMessage(csSelectJob)
	return nil
}
