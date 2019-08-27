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

type ignoreListRequest struct {
	PageIndex  int    `json:"pageIndex"`
	PlayerName string `json:"playerName"`
	PlatformId int32  `json:"centerPlatformId"`
	ServerId   int32  `json:"centerServerId"`
	Reason     string `json:"reason"`
}

type ignoreListRespon struct {
	ItemArray  []*ignoreListResponItem `json:"itemArray"`
	TotalCount int                     `json:"total"`
}

type ignoreListResponItem struct {
	Id                string `json:"id"`
	UserId            int64  `json:"playerId"`
	PlatformId        int32  `json:"centerPlatformId"`
	ServerId          int64  `json:"centerServerId"`
	Name              string `json:"playerName"`
	IgnoreChat        int32  `json:"ignoreChat"`
	IgnoreChatText    string `json:"ignoreChatText"`
	IgnoreChatTime    int64  `json:"ignoreChatTime"`
	IgnoreChatEndTime int64  `json:"ignoreChatEndTime"`
	IgnoreChatName    string `json:"ignoreChatName"`
}

func handleIgnoreList(rw http.ResponseWriter, req *http.Request) {
	form := &ignoreListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := playerservice.PlayerServiceInContext(req.Context())
	rsp := &ignoreListRespon{}
	rsp.ItemArray = make([]*ignoreListResponItem, 0)

	centerService := monitor.CenterServerServiceInContext(req.Context())
	serverid := centerService.GetCenterServerDBId(form.PlatformId, form.ServerId)

	if serverid < 1 {
		log.WithFields(log.Fields{
			"serverid": serverid,
		}).Error("获取游戏玩家列表，获得服务器id为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetPlayerIgnoreList(gmdb.GameDbLink(serverid), form.ServerId, form.PlayerName, form.Reason, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家组列表异常,组")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &ignoreListResponItem{
			Id:                changeInt64ToString(value.Id),
			UserId:            value.UserId,
			ServerId:          value.ServerId,
			Name:              value.Name,
			IgnoreChat:        int32(value.IgnoreChat),
			IgnoreChatText:    value.IgnoreChatText,
			IgnoreChatTime:    value.IgnoreChatTime,
			IgnoreChatEndTime: value.IgnoreChatEndTime,
			IgnoreChatName:    value.IgnoreChatName,
			PlatformId:        form.PlatformId,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	count, err := service.GetPlayerIgnoreCount(gmdb.GameDbLink(serverid), form.ServerId, form.PlayerName, form.Reason)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家组列表异常")
	}
	rsp.TotalCount = count
	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
