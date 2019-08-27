package handler

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
)

//处理获取服务器列表
func handleServerList(s session.Session, msg interface{}) (err error) {
	log.Debug("server:获取服务器列表")
	scServerList := msg.(*uipb.SCServerList)
	for _, server := range scServerList.Servers {
		log.WithFields(
			log.Fields{
				"ip":   server.GetIp(),
				"name": server.GetName(),
			}).Debug("server:获取服务器列表完成")
	}
	log.Debug("server:获取服务器列表完成")
	return
}
