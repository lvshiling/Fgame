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

const (
	minLevelNum = int32(1)
)

type playerLevelStaticRequest struct {
	ServerId int `json:"serverId"`
}

type playerLevelStaticRespon struct {
	ItemArray []*playerLevelStaticResponItem `json:"itemArray"`
}

type playerLevelStaticResponItem struct {
	Level       int32 `json:"level"`
	PlayerCount int32 `json:"playerCount"`
	TotalCount  int32 `json:"totalCount"`
	Rate        int32 `json:"rate"`
}

func handlePlayerLevelStaticList(rw http.ResponseWriter, req *http.Request) {
	form := &playerLevelStaticRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家等级统计，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := playerservice.PlayerServiceInContext(req.Context())
	centerService := monitor.CenterServerServiceInContext(req.Context())

	acServerId, err := centerService.GetServerId(int64(form.ServerId))
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("获取游戏玩家等级统计，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &playerLevelStaticRespon{}
	respon.ItemArray = make([]*playerLevelStaticResponItem, 0)
	rst, err := service.GetPlayerLevelStatic(gmdb.GameDbLink(form.ServerId), acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("获取游戏玩家等级统计，获取统计结果异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	totalCount := int32(0)
	tempLevelIndex := minLevelNum
	for _, value := range rst {
		if tempLevelIndex != value.Level {
			for itemIndex := tempLevelIndex; itemIndex < value.Level; itemIndex++ {
				myItem := &playerLevelStaticResponItem{
					Level:       itemIndex,
					PlayerCount: 0,
				}
				respon.ItemArray = append(respon.ItemArray, myItem)
			}
		}
		item := &playerLevelStaticResponItem{
			Level:       value.Level,
			PlayerCount: value.PlayerCount,
		}
		totalCount += value.PlayerCount
		respon.ItemArray = append(respon.ItemArray, item)
		tempLevelIndex = value.Level + 1
	}
	for _, value := range respon.ItemArray {
		value.TotalCount = totalCount
		if totalCount != 0 {
			rate := value.PlayerCount * 10000 / totalCount
			value.Rate = rate
		}
	}
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
