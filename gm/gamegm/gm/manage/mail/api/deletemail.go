package api

import (
	"net/http"

	mailservice "fgame/fgame/gm/gamegm/gm/manage/mail/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type deleteMailRequest struct {
	Id int64 `json:"id"`
}

func handleDeleteMail(rw http.ResponseWriter, req *http.Request) {
	form := &deleteMailRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除邮件，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	mailService := mailservice.MailServiceInContext(req.Context())
	if mailService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除邮件，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	info, err := mailService.GetMailInfo(form.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除邮件，获取异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	//下面应该要提示
	if info.MailState > 1 {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除邮件，已经审核，不能修改")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	// userid := gmUserService.GmUserIdInContext(req.Context())
	err = mailService.DeleteInfo(form.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除邮件异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
