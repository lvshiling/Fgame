package api

import (
	"net/http"
	"strconv"

	"fgame/fgame/gm/gamegm/constant"
	mongoservice "fgame/fgame/gm/gamegm/mglog/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type getMongoLogRequest struct {
	TableName  string `json:"tableName"`
	BeginTime  int64  `json:"beginTime"`
	EndTime    int64  `json:"endTime"`
	Platform   int32  `json:"platformId"`
	ServerType int32  `json:"serverType"`
	ServerId   int32  `json:"serverId"`
	PageIndex  int    `json:"pageIndex"`
	PlayerId   string `json:"playerId"`
	AllianceId string `json:"allianceId"`
}

type getMongoLogRespon struct {
	ItemArray  interface{} `json:"itemArray"`
	TotalCount int         `json:"totalCount"`
}

func handleGetMongoLog(rw http.ResponseWriter, req *http.Request) {
	form := &getMongoLogRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("查询日志，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := mongoservice.MgLogServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("查询日志，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	searchPlayerId, _ := strconv.ParseInt(form.PlayerId, 10, 64)
	searchAllianceId, _ := strconv.ParseInt(form.AllianceId, 10, 64)

	rst, err := service.GetLogMsg(form.TableName, form.BeginTime, form.EndTime, form.Platform, form.ServerType, form.ServerId, searchPlayerId, searchAllianceId, form.PageIndex, constant.DefaultPageSize)
	if err != nil {
		log.WithFields(log.Fields{
			"tableName":  form.TableName,
			"beginTime":  form.BeginTime,
			"endTime":    form.EndTime,
			"platform":   form.Platform,
			"serverType": form.ServerType,
			"serverId":   form.ServerId,
			"pageindex":  form.PageIndex,
			"error":      err,
		}).Error("查询日志，执行查询异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	totalCount, err := service.GetLogMsgCount(form.TableName, form.BeginTime, form.EndTime, form.Platform, form.ServerType, form.ServerId, searchPlayerId, searchAllianceId)
	if err != nil {
		log.WithFields(log.Fields{
			"tableName":  form.TableName,
			"beginTime":  form.BeginTime,
			"endTime":    form.EndTime,
			"platform":   form.Platform,
			"serverType": form.ServerType,
			"serverId":   form.ServerId,
			"pageindex":  form.PageIndex,
			"error":      err,
		}).Error("查询日志，执行查询总数异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &getMongoLogRespon{}
	respon.ItemArray = rst
	respon.TotalCount = totalCount
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
