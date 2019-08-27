package api

import (
	"fmt"
	"net/http"

	gmdb "fgame/fgame/gm/gamegm/db"
	alliservice "fgame/fgame/gm/gamegm/gm/game/alliance/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type allianceRequest struct {
	ServerId     int    `json:"serverId"`
	AllianceName string `json:"allianceName"`
	PageIndex    int    `json:"pageIndex"`
	OrderColumn  int    `json:"ordercol"`
	OrderType    int    `json:"ordertype"`
}

type allianceRespon struct {
	ItemArray  []*allianceResponItem `json:"itemArray"`
	TotalCount int                   `json:"total"`
}

type allianceResponItem struct {
	Id            string `json:"id"`
	AllianceName  string `json:"allianceName"`
	AllianceLevel int    `json:"allianceLevel"`
	TotalForce    int    `json:"totalForce"`
	CreateTime    int    `json:"createTime"`
	PlayerCount   int    `json:"playerCount"`
	Notice        string `json:"notice"`
}

func handleAllianceList(rw http.ResponseWriter, req *http.Request) {
	form := &allianceRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("仙盟工会列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := alliservice.AllianceServiceInContext(req.Context())
	centerService := monitor.CenterServerServiceInContext(req.Context())

	acServerId, err := centerService.GetServerId(int64(form.ServerId))
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("仙盟工会列表，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetAllianceList(gmdb.GameDbLink(form.ServerId), acServerId, form.AllianceName, form.PageIndex, form.OrderColumn, form.OrderType)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("仙盟工会列表，获取失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &allianceRespon{}
	respon.ItemArray = make([]*allianceResponItem, 0)
	for _, value := range rst {
		item := &allianceResponItem{
			Id:            fmt.Sprintf("%d", value.Id),
			AllianceName:  value.AllianceName,
			AllianceLevel: value.AllianceLevel,
			TotalForce:    value.TotalForce,
			CreateTime:    value.CreateTime,
			PlayerCount:   value.PlayerCount,
			Notice:        value.Notice,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := service.GetAllianceCount(gmdb.GameDbLink(form.ServerId), acServerId, form.AllianceName)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("仙盟工会列表异常")
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
