package api

import (
	mailservice "fgame/fgame/gm/gamegm/gm/manage/mail/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	center "fgame/fgame/gm/gamegm/gm/center/server/service"
	dbmodel "fgame/fgame/gm/gamegm/gm/manage/mail/model"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type applyListRequest struct {
	State     int    `json:"mailState"`
	Title     string `json:"title"`
	PlayerId  string `json:"playerId"`
	PageIndex int    `json:"pageIndex"`
}

type mailListRespon struct {
	ItemArray  []*mailListResponItem `json:"itemArray"`
	TotalCount int                   `json:"total"`
}

type mailListResponItem struct {
	Id               int64  `json:"id"`
	MailType         int    `json:"mailType"`
	ServerId         int    `json:"serverId"`
	ServerName       string `json:"serverName"`
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
	UpdateTime       int64  `json:"updateTime"`
	CreateTime       int64  `json:"createTime"`
	DeleteTime       int64  `json:"deleteTime"`
	MailUser         int64  `json:"mailUser"`
	MailTime         int64  `json:"mailTime"`
	MailState        int    `json:"mailState"`
	ApproveUser      int64  `json:"approveUser"`
	ApproveTime      int64  `json:"approveTime"`
	ApproveReason    string `json:"approveReason"`
	SendFlag         int    `json:"sendFlag"`
	SdkType          int    `json:"sdkType"`
	CenterPlatformId int64  `json:"centerPlatformId"`
	BindFlag         int    `json:"bindFlag"`
	Remark           string `json:"remark"`
}

func handleApplyList(rw http.ResponseWriter, req *http.Request) {
	form := &applyListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("邮件列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Debug("邮件列表：", form.PlayerId)
	mailService := mailservice.MailServiceInContext(req.Context())
	if mailService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取邮件，服务为空")
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

	rst, err := mailService.GetApplyList(userid, form.PageIndex, form.Title, form.State, userSdkTypeList, false, form.PlayerId)
	if err != nil {
		log.WithFields(log.Fields{
			"userId":    userid,
			"PageIndex": form.PageIndex,
			"Title":     form.Title,
			"State":     form.State,
			"error":     err,
		}).Error("获取邮件，执行异常")
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

	count, err := mailService.GetApplyCount(userid, form.Title, form.State, userSdkTypeList, false, form.PlayerId)
	if err != nil {
		log.WithFields(log.Fields{
			"userId":    userid,
			"PageIndex": form.PageIndex,
			"Title":     form.Title,
			"State":     form.State,
			"error":     err,
		}).Error("获取邮件个数，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}

func changeDbModelToRespon(p_info *dbmodel.MailApply) *mailListResponItem {
	rst := &mailListResponItem{}
	rst.Id = p_info.Id
	rst.MailType = p_info.MailType
	rst.ServerId = p_info.ServerId
	rst.Title = p_info.Title
	rst.Content = p_info.Content
	rst.Playerlist = p_info.Playerlist
	rst.Proplist = p_info.Proplist
	rst.FreezTime = p_info.FreezTime
	rst.EffectDays = p_info.EffectDays
	rst.RoleStartTime = p_info.RoleStartTime
	rst.RoleEndTime = p_info.RoleEndTime
	rst.MinLevel = p_info.MinLevel
	rst.MaxLevel = p_info.MaxLevel
	rst.UpdateTime = p_info.UpdateTime
	rst.CreateTime = p_info.CreateTime
	rst.DeleteTime = p_info.DeleteTime
	rst.MailUser = p_info.MailUser
	rst.MailTime = p_info.MailTime
	rst.MailState = p_info.MailState
	rst.ApproveUser = p_info.ApproveUser
	rst.ApproveTime = p_info.ApproveTime
	rst.ApproveReason = p_info.ApproveReason
	rst.SendFlag = p_info.SendFlag
	rst.SdkType = p_info.SdkType
	rst.CenterPlatformId = p_info.CenterPlatformId
	rst.BindFlag = p_info.BindFlag
	rst.Remark = p_info.Remark
	return rst
}
