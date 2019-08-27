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

type jinyanListRequest struct {
	PageIndex  int    `json:"pageIndex"`
	PlayerName string `json:"playerName"`
	PlatformId int32  `json:"centerPlatformId"`
	ServerId   int32  `json:"centerServerId"`
	Reason     string `json:"reason"`
	ForbidTime int64  `json:"forbidTime"`
}

type jinyanListRespon struct {
	ItemArray  []*jinyanListResponItem `json:"itemArray"`
	TotalCount int                     `json:"total"`
}

type jinyanListResponItem struct {
	Id                string `json:"id"`
	UserId            int64  `json:"playerId"`
	PlatformId        int32  `json:"centerPlatformId"`
	ServerId          int64  `json:"centerServerId"`
	Name              string `json:"playerName"`
	ForbidChat        int32  `json:"forbidChat"`
	ForbidChatText    string `json:"forbidChatText"`
	ForbidChatTime    int64  `json:"forbidChatTime"`
	ForbidChatEndTime int64  `json:"forbidChatEndTime"`
	ForbidChatName    string `json:"forbidChatName"`
}

func handleJinYanList(rw http.ResponseWriter, req *http.Request) {
	form := &jinyanListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := playerservice.PlayerServiceInContext(req.Context())
	rsp := &jinyanListRespon{}
	rsp.ItemArray = make([]*jinyanListResponItem, 0)

	centerService := monitor.CenterServerServiceInContext(req.Context())
	serverid := centerService.GetCenterServerDBId(form.PlatformId, form.ServerId)

	if serverid < 1 {
		log.WithFields(log.Fields{
			"serverid": serverid,
		}).Error("获取游戏玩家列表，获得服务器id为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetPlayerJinYanList(gmdb.GameDbLink(serverid), form.ServerId, form.PlayerName, form.Reason, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家组列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &jinyanListResponItem{
			Id:                changeInt64ToString(value.Id),
			UserId:            value.UserId,
			ServerId:          value.ServerId,
			Name:              value.Name,
			ForbidChat:        int32(value.ForbidChat),
			ForbidChatName:    value.ForbidChatName,
			ForbidChatText:    value.ForbidChatText,
			ForbidChatTime:    value.ForbidChatTime,
			PlatformId:        form.PlatformId,
			ForbidChatEndTime: value.ForbidChatEndTime,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	count, err := service.GetPlayerJinYanCount(gmdb.GameDbLink(serverid), form.ServerId, form.PlayerName, form.Reason)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家组列表异常")
	}
	rsp.TotalCount = count
	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
