package api

import (
	"fgame/fgame/gm/gamegm/constant"
	playerservice "fgame/fgame/gm/gamegm/gm/game/player/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/gm/gamegm/utils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type playerItemChangeRequest struct {
	// ServerId  int32  `json:"serverId"`
	PlayerId  string `json:"playerId"`
	ItemId    int64  `json:"itemId"`
	StartTime int64  `json:"begin"`
	EndTime   int64  `json:"end"`
	PageIndex int    `json:"pageIndex"`
}

type playerItemChangeRespon struct {
	ItemArray interface{} `json:"itemArray"`
	ItemCount int         `json:"total"`
}

func handlePlayerItemChangeList(rw http.ResponseWriter, req *http.Request) {
	form := &playerItemChangeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家物品更改日志列表，解析异常")
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
	logList, err := mongoLogService.GetPlayerItemChangeLogList(form.StartTime, form.EndTime, playerId, form.ItemId, form.PageIndex, pageSize)

	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"startTime": form.StartTime,
			"endTime":   form.EndTime,
			"playerId":  playerId,
			"pageindex": form.PageIndex,
		}).Error("获取玩家物品变更日志异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	logCount, err := mongoLogService.GetPlayerItemChangeLogCount(form.StartTime, form.EndTime, playerId, form.ItemId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"startTime": form.StartTime,
			"endTime":   form.EndTime,
			"playerId":  playerId,
			"pageindex": form.PageIndex,
		}).Error("获取玩家物品变更日志异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &playerItemChangeRespon{}
	respon.ItemArray = logList
	respon.ItemCount = logCount

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
