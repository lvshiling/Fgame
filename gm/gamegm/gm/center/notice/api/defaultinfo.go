package api

import (
	"net/http"

	gmloginNotice "fgame/fgame/gm/gamegm/gm/center/notice/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type defaultNoticeRespon struct {
	ID         int64  `json:"id"`
	Content    string `json:"content"`
	PlatformId int    `json:"platformId"`
	CreateTime int64  `json:"createTime"`
}

func handleDefaultNoticeInfo(rw http.ResponseWriter, req *http.Request) {

	service := gmloginNotice.LoginNoticeServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{}).Error("获取中心默认公告，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetDefaultNotice()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心默认公告，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &defaultNoticeRespon{
		ID:         rst.Id,
		Content:    rst.Content,
		PlatformId: rst.PlatformId,
		CreateTime: rst.CreateTime,
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
