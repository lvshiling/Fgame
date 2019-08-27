package api

import (
	playerstat "fgame/fgame/gm/gamegm/gm/mglog/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type playerStatsRequest struct {
	BeginTime int64 `json:"beginTime"`
	EndTime   int64 `json:"endTime"`
	PageIndex int   `json:"pageIndex"`
}

type playerStatsRespon struct {
	ItemArray  []*playerStatsResponItem `json:"itemArray"`
	TotalCount int                      `json:"totalCount"`
}
type playerStatsResponItem struct {
	BeginTime  int64  `json:"beginTime"`
	StatType   string `json:"statType"`
	StatCount  int    `json:"statCount"`
	UpdateTime int64  `json:"updateTime"`
}

func handlePlayerStatsList(rw http.ResponseWriter, req *http.Request) {
	form := &playerStatsRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家统计信息列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := playerstat.PlayerStatsServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家统计信息列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	list, err := service.GetPlayerStats(form.BeginTime, form.EndTime, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家统计信息列表，获取玩家统计异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &playerStatsRespon{}
	for _, value := range list {
		item := &playerStatsResponItem{
			BeginTime:  value.BeginTime,
			StatType:   value.StatType,
			StatCount:  value.StatCount,
			UpdateTime: value.UpdateTime,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}
	count, err := service.GetPlayerStatsCount(form.BeginTime, form.EndTime)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家统计信息列表，获取玩家统计个数异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon.TotalCount = count

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
