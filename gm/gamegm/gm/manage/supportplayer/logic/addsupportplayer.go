package logic

import (
	userremote "fgame/fgame/gm/gamegm/remote/service"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type LogicAddSupportPlayerParam struct {
	ChannelId        int   //gm渠道Id
	PlatformId       int   //gm平台Id
	CenterPlatformId int   //中心服平台Id
	ServerKeyId      int32 //中心服主键ID
	PlayerId         int64 //玩家ID
}

func HandleAddSupportPlayer(rw http.ResponseWriter, req *http.Request, param *LogicAddSupportPlayerParam) error {
	remoteService := userremote.UserRemoteServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("openApi:扶持玩家设置，Remote服务为空")
		return fmt.Errorf("openApi:扶持玩家设置，Remote服务为空")
	}
	err := remoteService.PrivilegeChargeSet(param.ServerKeyId, param.PlayerId, 1)
	if err != nil {
		return err
	}
	return nil
}
