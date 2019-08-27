package api

import (
	"net/http"

	gmloginNotice "fgame/fgame/gm/gamegm/gm/center/notice/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type loginNoticeListRequest struct {
	PageIndex int `form:"pageIndex" json:"pageIndex"`
}

type loginNoticeListRespon struct {
	ItemArray  []*loginNoticeListResponItem `json:"itemArray"`
	TotalCount int                          `json:"total"`
}

type loginNoticeListResponItem struct {
	ID         int64  `json:"id"`
	Content    string `json:"content"`
	PlatformId int    `json:"platformId"`
	CreateTime int64  `json:"createTime"`
}

func handleLoginNoticeList(rw http.ResponseWriter, req *http.Request) {
	form := &loginNoticeListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心公告列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmloginNotice.LoginNoticeServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心公告列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetLoginNoticeList(form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"index": form.PageIndex,
		}).Error("获取中心公告列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &loginNoticeListRespon{}
	respon.ItemArray = make([]*loginNoticeListResponItem, 0)

	for _, value := range rst {
		item := &loginNoticeListResponItem{
			ID:         value.Id,
			Content:    value.Content,
			PlatformId: value.PlatformId,
			CreateTime: value.CreateTime,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := service.GetLoginNoticeCount()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"index": form.PageIndex,
		}).Error("获取中心公告列表，执行异常")
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
