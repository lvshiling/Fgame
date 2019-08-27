package api

import (
	gmcenter "fgame/fgame/gm/gamegm/gm/center/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleAllServerById(rw http.ResponseWriter, req *http.Request) {
	form := &serverRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心组列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := gmcenter.CenterPlatformServiceInContext(req.Context())
	rsp := &serverRespon{}
	rsp.ItemArray = make([]*serverResponItem, 0)

	rst, err := service.GetAllServerByPlatForm(form.PlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心组列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &serverResponItem{
			ID:         value.Id,
			ServerId:   value.ServerId,
			ServerType: value.ServerType,
			Name:       value.ServerName,
			PlatformId: value.Platform,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
