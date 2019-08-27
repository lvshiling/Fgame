package api

import (
	"fgame/fgame/gm/gamegm/constant"
	playerservice "fgame/fgame/gm/gamegm/gm/game/player/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	"fgame/fgame/gm/gamegm/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type playerMongoLogRequest struct {
	TableName string `json:"tableName"`
	PlayerId  string `json:"playerId"`
	StartTime int64  `json:"begin"`
	EndTime   int64  `json:"end"`
	PageIndex int    `json:"pageIndex"`
}

type playerMongoLogRespon struct {
	ItemArray interface{} `json:"itemArray"`
	ItemCount int         `json:"total"`
}

func handlePlayerMongoLogList(rw http.ResponseWriter, req *http.Request) {
	form := &playerMongoLogRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家金币变更改日志列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	playerId := utils.ConverStringToInt64(form.PlayerId)
	// centerService := monitor.CenterServerServiceInContext(req.Context())
	// serverinfo, err := centerService.GetCenterServerDbInfo(int64(form.ServerId))
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"dbid":  form.ServerId,
	// 		"error": err,
	// 	}).Error("获取玩家物品更改日志列表，获取服务id异常")
	// 	rw.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// platformId := int32(serverinfo.Platform)
	// serverId := int32(serverinfo.ServerId)

	pageSize := constant.DefaultPageSize
	mongoLogService := playerservice.PlayerMongoLogServiceInContext(req.Context())
	logList, err := mongoLogService.GetPlayerMongoLogList(form.TableName, form.StartTime, form.EndTime, playerId, form.PageIndex, pageSize)

	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"startTime": form.StartTime,
			"endTime":   form.EndTime,
			"playerId":  playerId,
			"pageindex": form.PageIndex,
			"TableName": form.TableName,
		}).Error("获取玩家金币变更日志异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	logCount, err := mongoLogService.GetPlayerMongoLogCount(form.TableName, form.StartTime, form.EndTime, playerId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"startTime": form.StartTime,
			"endTime":   form.EndTime,
			"TableName": form.TableName,
			"pageindex": form.PageIndex,
		}).Error("获取玩家金币变更日志异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &playerMongoLogRespon{}
	respon.ItemArray = logList
	respon.ItemCount = logCount

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
