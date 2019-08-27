package server

import (
	"fgame/fgame/client/session"
	uipb "fgame/fgame/common/codec/pb/ui"

	log "github.com/Sirupsen/logrus"
)

func GetServers(s session.Session) (err error) {
	log.WithFields(
		log.Fields{}).Debug("gm:获取服务器列表")
	csServerList := &uipb.CSServerList{}
	s.Send(csServerList)
	return nil
}
