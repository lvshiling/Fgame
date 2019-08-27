package api

import (
	ntservice "fgame/fgame/gm/gamegm/gm/manage/notice/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type noticeListRequest struct {
	SuccessFlag int   `json:"successFlag"`
	StartTime   int64 `json:"beginTime"`
	EndTime     int64 `json:"endTime"`
	PageIndex   int   `json:"pageIndex"`
}

type noticeListRespon struct {
	ItemArray  []*noticeListResponItem `json:"itemArray"`
	TotalCount int                     `json:"total"`
}

type noticeListResponItem struct {
	Id               int64  `json:"id"`
	ChannelId        int    `json:"channelId"`
	PlatformId       int    `json:"platformId"`
	ServerId         int    `json:"serverId"`
	Content          string `json:"content"`
	BeginTime        int64  `json:"beginTime"`
	EndTime          int64  `json:"endTime"`
	IntervalTime     int64  `json:"intervalTime"`
	UpdateTime       int64  `json:"updateTime"`
	CreateTime       int64  `json:"createTime"`
	DeleteTime       int64  `json:"deleteTime"`
	SuccessFlag      int    `json:"successFlag"`
	ErrorMsg         string `json:"errorMsg"`
	ServerName       string `json:"serverName"`
	CenterPlatformId int64  `json:"centerPlatformId"`
}

func handleNoticeList(rw http.ResponseWriter, req *http.Request) {
	form := &noticeListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("公告列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := ntservice.NoticeServiceInContext(req.Context())

	userid := gmUserService.GmUserIdInContext(req.Context())

	usservice := gmUserService.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("公告列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userPlatformList, err := usservice.GetUserCenterPlatList(userid)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("公告列表，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &noticeListRespon{}
	respon.ItemArray = make([]*noticeListResponItem, 0)

	list, err := service.GetNoticeList(form.SuccessFlag, form.StartTime, form.EndTime, form.PageIndex, userPlatformList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("公告列表，获取列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range list {
		item := &noticeListResponItem{
			Id:               value.Id,
			ChannelId:        value.ChannelId,
			PlatformId:       value.PlatformId,
			ServerId:         value.ServerId,
			Content:          value.Content,
			BeginTime:        value.BeginTime,
			EndTime:          value.EndTime,
			IntervalTime:     value.IntervalTime,
			UpdateTime:       value.UpdateTime,
			CreateTime:       value.CreateTime,
			DeleteTime:       value.DeleteTime,
			SuccessFlag:      value.SuccessFlag,
			ErrorMsg:         value.ErrorMsg,
			ServerName:       value.ServerName,
			CenterPlatformId: value.CenterPlatformId,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}
	count, err := service.GetNoticeCount(form.SuccessFlag, form.StartTime, form.EndTime, userPlatformList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("公告列表，获取列表数量异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon.TotalCount = count

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
