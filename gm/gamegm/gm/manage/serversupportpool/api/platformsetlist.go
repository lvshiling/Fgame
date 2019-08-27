package api

import (
	serversupp "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type platformSupportPoolSetRequest struct {
	PageIndex        int32 `json:"pageIndex"`
	CenterPlatformId int64 `json:"centerPlatformId"`
}

type platformSupportPoolSetRespon struct {
	ItemArray  []*platformSupportPoolSetResponItem `json:"itemArray"`
	TotalCount int32                               `json:"total"`
}

type platformSupportPoolSetResponItem struct {
	Id               int64 `json:"id"`
	CenterPlatformId int64 `json:"centerPlatformId"`
	Gold             int32 `json:"gold"`
	Percent          int32 `json:"percent"`
}

func handlePlatformSupportPoolSetList(rw http.ResponseWriter, req *http.Request) {
	form := &platformSupportPoolSetRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取平台扶植池列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	serverPoolService := serversupp.ServerSupportPoolInContext(req.Context())
	if serverPoolService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取平台扶植池列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userid := gmUserService.GmUserIdInContext(req.Context())

	usservice := gmUserService.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("获取平台扶植池列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := serverPoolService.GetPlatformSupportPoolList(form.CenterPlatformId, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"userId":    userid,
			"PageIndex": form.PageIndex,
			"error":     err,
		}).Error("获取平台扶植池列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &platformSupportPoolSetRespon{}
	respon.ItemArray = make([]*platformSupportPoolSetResponItem, 0)
	for _, value := range rst {
		item := &platformSupportPoolSetResponItem{
			Id:               value.Id,
			CenterPlatformId: value.CenterPlatformId,
			Gold:             value.SupportGold,
			Percent:          value.SupportRate,
		}

		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := serverPoolService.GetPlatformSupportPoolCount(form.CenterPlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"userId":    userid,
			"PageIndex": form.PageIndex,
			"error":     err,
		}).Error("获取平台扶植池列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
