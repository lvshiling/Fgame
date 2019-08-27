package api

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	playerservice "fgame/fgame/gm/gamegm/gm/game/player/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type fengJinListRequest struct {
	PageIndex  int    `json:"pageIndex"`
	PlayerName string `json:"playerName"`
	PlatformId int32  `json:"centerPlatformId"`
	ServerId   int32  `json:"centerServerId"`
	Reason     string `json:"reason"`
}

type fengJinListRespon struct {
	ItemArray  []*fengJinListResponItem `json:"itemArray"`
	TotalCount int                      `json:"total"`
}

type fengJinListResponItem struct {
	Id            string `json:"id"`
	UserId        int64  `json:"playerId"`
	PlatformId    int32  `json:"centerPlatformId"`
	ServerId      int64  `json:"centerServerId"`
	Name          string `json:"playerName"`
	Forbid        int32  `json:"forbid"`
	ForbidText    string `json:"forbidText"`
	ForbidTime    int64  `json:"forbidTime"`
	ForbidEndTime int64  `json:"forbidEndTime"`
	ForbidName    string `json:"forbidName"`
}

func handleFengJinList(rw http.ResponseWriter, req *http.Request) {
	form := &fengJinListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := playerservice.PlayerServiceInContext(req.Context())
	rsp := &fengJinListRespon{}
	rsp.ItemArray = make([]*fengJinListResponItem, 0)

	centerService := monitor.CenterServerServiceInContext(req.Context())
	serverid := centerService.GetCenterServerDBId(form.PlatformId, form.ServerId)

	if serverid < 1 {
		log.WithFields(log.Fields{
			"serverid": serverid,
		}).Error("获取游戏玩家列表，获得服务器id为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetPlayerFengJinList(gmdb.GameDbLink(serverid), form.ServerId, form.PlayerName, form.Reason, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家组列表异常,组")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &fengJinListResponItem{
			Id:            changeInt64ToString(value.Id),
			UserId:        value.UserId,
			ServerId:      value.ServerId,
			Name:          value.Name,
			Forbid:        int32(value.Forbid),
			ForbidText:    value.ForbidText,
			ForbidTime:    value.ForbidTime,
			ForbidName:    value.ForbidName,
			PlatformId:    form.PlatformId,
			ForbidEndTime: value.ForbidEndTime,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	count, err := service.GetPlayerFengJinCount(gmdb.GameDbLink(serverid), form.ServerId, form.PlayerName, form.Reason)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家组列表异常")
	}
	rsp.TotalCount = count
	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
