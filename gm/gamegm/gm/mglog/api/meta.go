package api

import (
	"net/http"

	mglog "fgame/fgame/gm/gamegm/gm/mglog/metadata"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type getLogMetaRequest struct {
	TableName string `json:"tableName"`
	LogType   int    `json:"logType"`
}

func handleGetLogMeta(rw http.ResponseWriter, req *http.Request) {
	form := &getLogMetaRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取消息元数据，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logType := mglog.MsgType(form.LogType)
	// if logType == mglog.MsgTypePlayerLog {
	// 	respon := mglog.GetPlayerMsgItemInfo(form.TableName)
	// 	rr := gmhttp.NewSuccessResult(respon)
	// 	httputils.WriteJSON(rw, http.StatusOK, rr)
	// 	return
	// }
	// if logType == mglog.MsgTypeAllianceLog {
	// 	respon := mglog.GetAllianceMsgItemInfo(form.TableName)
	// 	rr := gmhttp.NewSuccessResult(respon)
	// 	httputils.WriteJSON(rw, http.StatusOK, rr)
	// 	return
	// }

	respon := mglog.GetMsgItemInfo(logType, form.TableName)
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
	return
}
