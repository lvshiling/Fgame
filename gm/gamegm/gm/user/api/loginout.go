package api

import (
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleLoginOut(rw http.ResponseWriter, req *http.Request) {
	service := gmUserService.LoginServiceInContext(req.Context())
	if service == nil {
		log.Error("用户登陆，获取登陆服务异常")
	}
	gmUserId := gmUserService.GmUserIdInContext(req.Context())

	err := service.LoginOut(gmUserId)
	if err != nil {
		log.WithFields(log.Fields{
			"playid": gmUserId,
			"error":  err,
		}).Error("用户登出异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
