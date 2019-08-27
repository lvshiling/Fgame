package api

import (
	"net/http"

	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	mailservice "fgame/fgame/gm/gamegm/gm/manage/mail/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type approveMailRequest struct {
	Id     int64  `json:"id"`
	State  int    `json:"mailState"`
	Reason string `json:"approveReason"`
}

func handleApproveMail(rw http.ResponseWriter, req *http.Request) {
	form := &approveMailRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("审核邮件，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	mailService := mailservice.MailServiceInContext(req.Context())
	if mailService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("审核邮件，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	mailinfo, err := mailService.GetMailInfo(form.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("审核邮件异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !mailinfo.IfCanApprove() {
		codeErr := gmError.GetError(gmError.ErrorCodeEmailChecked)
		errhttp.ResponseWithErrorMessage(rw, codeErr, codeErr.Error())
		return
	}

	userid := gmUserService.GmUserIdInContext(req.Context())
	err = mailService.ApproveMail(form.Id, userid, form.State, form.Reason)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("审核邮件异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if form.State == 3 {
		rr := gmhttp.NewSuccessResult(nil)
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}

	remoteService := userremote.UserRemoteServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("审核邮件，Remote服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if mailinfo.MailType == 1 {
		err = remoteService.SendPlayerCompensate(int32(mailinfo.ServerId), mailinfo.Playerlist, mailinfo.Title, mailinfo.Content, mailinfo.Proplist, mailinfo.BindFlag)
		if err != nil {
			log.WithFields(log.Fields{
				"ServerId":   mailinfo.ServerId,
				"Playerlist": mailinfo.Playerlist,
				"Title":      mailinfo.Title,
				"Content":    mailinfo.Content,
				"Proplist":   mailinfo.Proplist,
				"error":      err,
			}).Error("审核邮件异常,remote发送失败")
			// rw.WriteHeader(http.StatusInternalServerError)
			// return
			codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
			errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
			return
		}
	} else {
		err = remoteService.SendServerCompensate(int32(mailinfo.ServerId), mailinfo.Title, mailinfo.Content, mailinfo.Proplist, int32(mailinfo.MinLevel), mailinfo.RoleStartTime, mailinfo.BindFlag)
		if err != nil {
			log.WithFields(log.Fields{
				"ServerId":      mailinfo.ServerId,
				"MinLevel":      mailinfo.MinLevel,
				"RoleStartTime": mailinfo.RoleStartTime,
				"Title":         mailinfo.Title,
				"Content":       mailinfo.Content,
				"Proplist":      mailinfo.Proplist,
				"error":         err,
			}).Error("审核邮件异常,remote发送失败")
			// rr := gmhttp.NewFailedResultWithMsg(1000, err.Error())
			// httputils.WriteJSON(rw, http.StatusOK, rr)
			// return
			codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
			errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
			return
		}
	}
	log.Debug("remote发送完成")
	err = mailService.UpdateSendFlag(form.Id, 1)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("审核邮件异常，更新发送状态失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
