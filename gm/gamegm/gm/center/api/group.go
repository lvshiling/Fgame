package api

import (
	gmcenter "fgame/fgame/gm/gamegm/gm/center/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type groupRequest struct {
	PlatformId int64 `form:"centerPlatformId" json:"centerPlatformId"`
}

type groupRespon struct {
	ItemArray []*groupResponItem `json:"itemArray"`
}

type groupResponItem struct {
	GroupId    int64 `json:"groupId"`
	Name       int64 `json:"groupOtherId"`
	PlatformId int64 `json:"centerPlatformId"`
}

func handleGroup(rw http.ResponseWriter, req *http.Request) {
	form := &groupRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心组列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := gmcenter.CenterPlatformServiceInContext(req.Context())
	rsp := &groupRespon{}
	rsp.ItemArray = make([]*groupResponItem, 0)

	rst, err := service.GetGroupByPlatForm(form.PlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心组列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &groupResponItem{
			GroupId:    value.Id,
			Name:       value.GroupId,
			PlatformId: value.Platform,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
