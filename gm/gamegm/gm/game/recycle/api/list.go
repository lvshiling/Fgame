package api

import (
	"net/http"

	gmdb "fgame/fgame/gm/gamegm/db"
	recycleservice "fgame/fgame/gm/gamegm/gm/game/recycle/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type recycleRequest struct {
	ServerId int `json:"serverId"`
}

type recycleRespon struct {
	ItemArray []*recycleResponItem `json:"itemArray"`
}

type recycleResponItem struct {
	Id                int64 `json:"id"`
	ServerId          int32 `json:"serverId"`
	RecycleGold       int64 `json:"recycleGold"`
	RecycleTime       int64 `json:"recycleTime"`
	CustomRecycleGold int64 `json:"customRecycleGold"`
}

func handleRecycleList(rw http.ResponseWriter, req *http.Request) {
	form := &recycleRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取回收元宝列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := recycleservice.RecycleServiceInContext(req.Context())
	centerService := monitor.CenterServerServiceInContext(req.Context())

	acServerId, err := centerService.GetServerId(int64(form.ServerId))
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("获取回收元宝列表，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetServerRecycleModel(gmdb.GameDbLink(form.ServerId), acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("获取回收元宝列表，获取失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &recycleRespon{}
	respon.ItemArray = make([]*recycleResponItem, 0)

	item := &recycleResponItem{
		Id:                rst.Id,
		ServerId:          rst.ServerId,
		RecycleGold:       rst.RecycleGold,
		RecycleTime:       rst.RecycleTime,
		CustomRecycleGold: rst.CustomRecycleGold,
	}
	respon.ItemArray = append(respon.ItemArray, item)

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
