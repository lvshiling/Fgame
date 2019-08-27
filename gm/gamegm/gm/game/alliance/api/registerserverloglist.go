package api

import (
	"net/http"

	gmdb "fgame/fgame/gm/gamegm/db"
	alliservice "fgame/fgame/gm/gamegm/gm/game/alliance/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type registerServerLogRequest struct {
	ServerId  int32 `json:"serverId"`
	PageIndex int   `json:"pageIndex"`
}

type registerServerLogRespon struct {
	ItemArray  []*registerServerLogResponItem `json:"itemArray"`
	TotalCount int                            `json:"total"`
}

type registerServerLogResponItem struct {
	Id         int64 `json:"id"`
	ServerId   int   `json:"serverId"`
	Open       int   `json:"open"`
	UpdateTime int64 `json:"updateTime"`
	CreateTime int64 `json:"createTime"`
	DeleteTime int64 `json:"deleteTime"`
}

func handleRegisterServerLog(rw http.ResponseWriter, req *http.Request) {
	form := &registerServerLogRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("服务器启用状态列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := alliservice.AllianceServiceInContext(req.Context())
	if service == nil {
		log.Error("服务器启用状态列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	centerService := monitor.CenterServerServiceInContext(req.Context())

	acServerId, err := centerService.GetServerId(int64(form.ServerId))
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("服务器启用状态列表，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetServerRegisterLogList(gmdb.GameDbLink(form.ServerId), acServerId, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("服务器启用状态列表，获取失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rsp := &registerServerLogRespon{}
	rsp.ItemArray = make([]*registerServerLogResponItem, 0)
	for _, value := range rst {
		item := &registerServerLogResponItem{
			Id:         value.Id,
			ServerId:   value.ServerId,
			Open:       value.Open,
			UpdateTime: value.UpdateTime,
			CreateTime: value.CreateTime,
			DeleteTime: value.DeleteTime,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	count, err := service.GetServerRegisterLogCount(gmdb.GameDbLink(form.ServerId), acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("服务器启用状态列表，获取失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rsp.TotalCount = count

	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
