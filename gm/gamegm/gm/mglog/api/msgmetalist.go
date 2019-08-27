package api

import (
	"net/http"

	mglog "fgame/fgame/gm/gamegm/gm/mglog/metadata"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type getLogMsgMetaListRequest struct {
	LogType int `json:"logType"`
}

func handleGetLogMsgMetaList(rw http.ResponseWriter, req *http.Request) {
	form := &getLogMsgMetaListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取消息元数据列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logType := mglog.MsgType(form.LogType)
	respon := mglog.GetMstList(logType)
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
	return
}
