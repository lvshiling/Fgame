package api

import (
	"net/http"

	mailservice "fgame/fgame/gm/gamegm/gm/manage/mail/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type updateMailRequest struct {
	Id               int64  `json:"id"`
	MailType         int    `json:"mailType"`
	ServerId         int    `json:"serverId"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Playerlist       string `json:"playerlist"`
	Proplist         string `json:"proplist"`
	FreezTime        int    `json:"freezTime"`
	EffectDays       int    `json:"effectDays"`
	RoleStartTime    int64  `json:"roleStartTime"`
	RoleEndTime      int64  `json:"roleEndTime"`
	MinLevel         int    `json:"minLevel"`
	MaxLevel         int    `json:"maxLevel"`
	SdkType          int    `json:"sdkType"`
	CenterPlatformId int64  `json:"centerPlatformId"`
	BindFlag         int    `json:"bindFlag"`
	Remark           string `json:"remark"`
}

func handleUpdateMail(rw http.ResponseWriter, req *http.Request) {
	form := &updateMailRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("修改邮件，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	mailService := mailservice.MailServiceInContext(req.Context())
	if mailService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("修改邮件，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	info, err := mailService.GetMailInfo(form.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("修改邮件，获取异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	//下面应该要提示
	if info.SendFlag > 1 {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("修改邮件，已经发送，不能修改")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userid := gmUserService.GmUserIdInContext(req.Context())
	err = mailService.UpdateMailInfo(form.Id, form.MailType, form.ServerId, form.Title, form.Content, form.Playerlist, form.Proplist, form.FreezTime, form.EffectDays, form.RoleStartTime, form.RoleEndTime, form.MinLevel, form.MaxLevel, userid, form.SdkType, form.CenterPlatformId, form.BindFlag, form.Remark)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("修改邮件异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
