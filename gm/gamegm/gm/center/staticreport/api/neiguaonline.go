package api

import (
	"fmt"
	"net/http"
	"sort"

	stservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	server "fgame/fgame/gm/gamegm/gm/center/server/service"

	"fgame/fgame/gm/gamegm/constant"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type neiGuaOnLineStaticRequest struct {
	PlatformId int   `json:"platformId"`
	StartTime  int64 `json:"startTime"`
	EndTime    int64 `json:"endTime"`
}

type neiGuaOnLineStaticRespon struct {
	ItemArray   []*neiGuaOnLineStaticResponItem `json:"itemArray"`
	ServerArray []*neiGuaOnLineStaticServer     `json:"serverArray"`
}

type neiGuaOnLineStaticResponItem struct {
	Date        int64          `json:"datestr"`
	MinuteIndex int64          `json:"minuteindex"`
	OnLineMap   map[string]int `json:"onlineMap"`
}

type neiGuaOnLineStaticServer struct {
	ServerId   string `json:"serverId"`
	ServerName string `json:"serverName"`
}

func handleNeiGuaOnLineStatic(rw http.ResponseWriter, req *http.Request) {
	form := &neiGuaOnLineStaticRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取在线列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	serverService := server.CenterServerServiceInContext(req.Context())
	if serverService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取在线列表，中心服务获取为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rpservice := stservice.StaticReportServiceInContext(req.Context())
	rst, err := rpservice.GetNeiGuaOnLineStatic(form.StartTime, form.EndTime, form.PlatformId, constant.DefaultTimeZone)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取在线列表，获取数据异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &neiGuaOnLineStaticRespon{}
	respon.ItemArray = make([]*neiGuaOnLineStaticResponItem, 0)
	tempMap := make(map[int]map[int]*neiGuaOnLineStaticResponItem)

	for _, value := range rst {
		if value.Id == nil {
			continue
		}
		if _, ok := tempMap[value.Id.Date]; !ok {
			tempMap[value.Id.Date] = make(map[int]*neiGuaOnLineStaticResponItem)
		}
		dateMap := tempMap[value.Id.Date]
		if _, ok := dateMap[value.Id.MinuteIndex]; !ok {
			singleItem := &neiGuaOnLineStaticResponItem{}
			singleItem.Date = int64(value.Id.Date*86400000) - int64(constant.DefaultTimeZone*60*60*1000)
			singleItem.MinuteIndex = singleItem.Date + int64(value.Id.MinuteIndex*300000)
			singleItem.OnLineMap = make(map[string]int)
			dateMap[value.Id.MinuteIndex] = singleItem
		}
		item := dateMap[value.Id.MinuteIndex]
		serverKey := fmt.Sprintf("s_%d_%d", value.Id.ServerId, value.Id.Platform)
		item.OnLineMap[serverKey] = value.MaxPlayer
	}
	for _, value := range tempMap {
		for _, rpItem := range value {
			respon.ItemArray = append(respon.ItemArray, rpItem)
		}
	}
	sort.Sort(neiGuaOnLineStaticSlice(respon.ItemArray))

	serverList, err := serverService.GetCenterServerListByPlatform(form.PlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取在线列表，获取服务列表数据异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, value := range serverList {
		serverItem := &neiGuaOnLineStaticServer{
			ServerId:   fmt.Sprintf("s_%d_%d", value.ServerId, value.Platform),
			ServerName: value.ServerName,
		}
		respon.ServerArray = append(respon.ServerArray, serverItem)
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}

type neiGuaOnLineStaticSlice []*neiGuaOnLineStaticResponItem //用于排序用的
func (a neiGuaOnLineStaticSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a neiGuaOnLineStaticSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}

func (a neiGuaOnLineStaticSlice) Less(i, j int) bool { //
	if a[j].Date < a[i].Date {
		return true
	}
	if a[j].Date > a[i].Date {
		return false
	}
	return a[j].MinuteIndex < a[i].MinuteIndex
}
