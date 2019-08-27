package api

import (
	"net/http"
	"strconv"

	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"
	gameplayerservice "fgame/fgame/gm/gamegm/gm/game/player/service"

	"fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	serverdailymodel "fgame/fgame/gm/gamegm/gm/manage/serverdaily/model"
	serverdaily "fgame/fgame/gm/gamegm/gm/manage/serverdaily/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type serverDailyReportRequest struct {
	CenterPlatformId int   `json:"centerPlatformId"`
	Begin            int64 `json:"startTime"`
	End              int64 `json:"endTime"`
}

type serverDailyReportRespon struct {
	ItemArray  []*serverDailyReportResponServer `json:"itemArray"`
	DailyArray []int64                          `json:"dailyArray"`
}

type serverDailyReportResponServer struct {
	ServerId       int32                                  `json:"serverId"`
	ServerName     string                                 `json:"serverName"`
	ServerChildStr string                                 `json:"serverChildStr"`
	FirstPower     int32                                  `json:"firstPower"`
	SecondPower    int32                                  `json:"secondPower"`
	ThirdPower     int32                                  `json:"thirePower"`
	DailData       map[int64]*serverDailyReportResponItem `json:"dailyData"`
}

type serverDailyReportResponItem struct {
	MaxOnLine   int32 `json:"maxOnLine"`
	LoginNum    int32 `json:"loginNum"`
	OrderAmount int32 `json:"orderAmount"`
}

func handleServerDailyReport(rw http.ResponseWriter, req *http.Request) {
	form := &serverDailyReportRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleServerDailyReport，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	centerServerService := centerserver.CenterServerServiceInContext(req.Context()) //中心服
	serverList, err := centerServerService.GetCenterMainServerListByPlatform(form.CenterPlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":            err,
			"centerPlatformId": form.CenterPlatformId,
		}).Error("handleServerDailyReport,获取中心服务器列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	gamePlayerService := gameplayerservice.PlayerServiceInContext(req.Context())
	dailyServerService := serverdaily.ServerDailyStatInContext(req.Context())

	dailyServerStat, err := dailyServerService.GetPlatformStatListByDate(form.CenterPlatformId, form.Begin, form.End)
	if err != nil {
		log.WithFields(log.Fields{
			"error":            err,
			"centerPlatformId": form.CenterPlatformId,
			"beginTime":        form.Begin,
			"endTime":          form.End,
		}).Error("handleServerDailyReport,获取每日统计异常列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	dailyServerMap := make(map[int64]map[int32]*serverdailymodel.ServerDailyStats)
	for _, value := range dailyServerStat {
		_, exists := dailyServerMap[value.CurDate]
		if !exists {
			dailyServerMap[value.CurDate] = make(map[int32]*serverdailymodel.ServerDailyStats)
		}
		dailyServerMap[value.CurDate][value.ServerId] = value
	}
	respon := &serverDailyReportRespon{}
	respon.ItemArray = make([]*serverDailyReportResponServer, 0)
	respon.DailyArray = make([]int64, 0)
	for timeValue := form.Begin; timeValue <= form.End; timeValue = timeValue + constant.DailyMillSecond {
		respon.DailyArray = append(respon.DailyArray, timeValue)
	}

	for _, serverValue := range serverList {
		if serverValue.ServerType != 0 {
			continue
		}
		serverItem := &serverDailyReportResponServer{}
		serverItem.ServerId = int32(serverValue.ServerId)
		serverItem.ServerName = serverValue.ServerName
		serverItem.DailData = make(map[int64]*serverDailyReportResponItem)
		for sgTimeValue := form.Begin; sgTimeValue <= form.End; sgTimeValue = sgTimeValue + constant.DailyMillSecond {
			item := &serverDailyReportResponItem{}
			dailyDateItem, exists := dailyServerMap[sgTimeValue]
			if exists {
				serverDailyItem, exists := dailyDateItem[int32(serverValue.ServerId)]
				if exists {
					item.LoginNum = serverDailyItem.LoginNum
					item.MaxOnLine = serverDailyItem.MaxOnlineNum
					item.OrderAmount = int32(serverDailyItem.OrderMoney)
				}
			}
			serverItem.DailData[sgTimeValue] = item
		}
		//获取前三
		threePower, err := gamePlayerService.GetPlayerTopThreePower(gmdb.GameDbLink(serverValue.Id), serverValue.ServerId)
		if err != nil {
			log.WithFields(log.Fields{
				"error":            err,
				"centerPlatformId": form.CenterPlatformId,
				"ServerKeyId":      serverValue.Id,
				"ServerId":         serverValue.ServerId,
			}).Error("handleServerDailyReport,获取服务器前三战力异常")
		}
		if threePower != nil {
			if len(threePower) > 0 {
				serverItem.FirstPower = threePower[0].Power
			}
			if len(threePower) > 1 {
				serverItem.SecondPower = threePower[1].Power
			}
			if len(threePower) > 2 {
				serverItem.ThirdPower = threePower[2].Power
			}
		}
		//获取合服的子服
		serverChildArray, err := centerServerService.GetCenterMergeRecord(int(serverValue.Platform), serverValue.ServerId)
		if err != nil {
			log.WithFields(log.Fields{
				"error":            err,
				"centerPlatformId": form.CenterPlatformId,
				"ServerId":         serverValue.ServerId,
			}).Error("handleServerDailyReport,获取合服服务器列表异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		serverStr := ""
		for index, childValue := range serverChildArray {
			if index > 0 {
				serverStr += ","
			}
			serverStr += strconv.Itoa(childValue.FromServerID)
		}
		serverItem.ServerChildStr = serverStr
		respon.ItemArray = append(respon.ItemArray, serverItem)
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
