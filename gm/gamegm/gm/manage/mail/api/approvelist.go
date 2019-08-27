package api

import (
	center "fgame/fgame/gm/gamegm/gm/center/server/service"
	mailservice "fgame/fgame/gm/gamegm/gm/manage/mail/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type approveListRequest struct {
	State     int    `json:"mailState"`
	Title     string `json:"title"`
	PageIndex int    `json:"pageIndex"`
	PlayerId  string `json:"playerId"`
}

func handleApproveList(rw http.ResponseWriter, req *http.Request) {
	form := &approveListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取审核邮件，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	mailService := mailservice.MailServiceInContext(req.Context())
	if mailService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取审核邮件，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userid := gmUserService.GmUserIdInContext(req.Context())

	usservice := gmUserService.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("邮件列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userSdkTypeList, err := usservice.GetUserCenterPlatList(userid)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("邮件列表，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := mailService.GetApproveList(form.PageIndex, form.Title, form.State, userSdkTypeList, form.PlayerId)
	if err != nil {
		log.WithFields(log.Fields{
			"userId":    userid,
			"PageIndex": form.PageIndex,
			"Title":     form.Title,
			"State":     form.State,
			"error":     err,
		}).Error("获取审核邮件，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	centerService := center.CenterServerServiceInContext(req.Context())

	respon := &mailListRespon{}
	respon.ItemArray = make([]*mailListResponItem, 0)
	for _, value := range rst {
		item := changeDbModelToRespon(value)
		cs, _ := centerService.GetCenterServer(int64(value.ServerId))
		item.ServerName = cs.ServerName
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := mailService.GetApproveCount(form.Title, form.State, userSdkTypeList, form.PlayerId)
	if err != nil {
		log.WithFields(log.Fields{
			"userId":    userid,
			"PageIndex": form.PageIndex,
			"Title":     form.Title,
			"State":     form.State,
			"error":     err,
		}).Error("获取审核邮件，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
