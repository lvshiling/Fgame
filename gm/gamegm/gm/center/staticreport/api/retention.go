package api

import (
	centerservice "fgame/fgame/gm/gamegm/gm/center/server/service"
	stservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type retentionRequest struct {
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	ServerId  int   `json:"serverId"`
}

type retentionRespon struct {
	ItemArray []*retentionResponServerItem `json:"itemArray"`
}

type retentionResponServerItem struct {
	OnLineDate int64 `json:"onLineDate"`
	Num0       int64 `json:"num0"`
	Num1       int64 `json:"num1"`
	Num2       int64 `json:"num2"`
	Num3       int64 `json:"num3"`
	Num4       int64 `json:"num4"`
	Num5       int64 `json:"num5"`
	Num6       int64 `json:"num6"`
	Num7       int64 `json:"num7"`
	Num14      int64 `json:"num14"`
	Num30      int64 `json:"num30"`
}

//统计留存率
func handleRetentionStatic(rw http.ResponseWriter, req *http.Request) {
	form := &retentionRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("统计留存率，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	serverService := centerservice.CenterServerServiceInContext(req.Context())
	serverInfo, err := serverService.GetCenterServer(int64(form.ServerId))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("统计留存率，获取服务器信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	playerStService := stservice.PlayerStaticInContext(req.Context())
	//查询出所有的在线人数
	list, err := playerStService.GetServerOnLine(int(serverInfo.Platform), serverInfo.ServerId, form.StartTime, form.EndTime)
	if err != nil {
		log.WithFields(log.Fields{
			"centerServerId": form.ServerId,
			"StartTime":      form.StartTime,
			"EndTime":        form.EndTime,
			"error":          err,
		}).Error("统计留存率，获取时间段在线人数异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &retentionRespon{}
	respon.ItemArray = make([]*retentionResponServerItem, 0)
	for _, value := range list {
		item := &retentionResponServerItem{
			OnLineDate: value.OnLineDate,
			Num0:       value.Num0,
			Num1:       value.Num1,
			Num2:       value.Num2,
			Num3:       value.Num3,
			Num4:       value.Num4,
			Num5:       value.Num5,
			Num6:       value.Num6,
			Num7:       value.Num7,
			Num14:      value.Num14,
			Num30:      value.Num30,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
