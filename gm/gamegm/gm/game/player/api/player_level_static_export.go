package api

import (
	"bytes"
	gmdb "fgame/fgame/gm/gamegm/db"
	playerservice "fgame/fgame/gm/gamegm/gm/game/player/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"github.com/xozrc/pkg/httputils"
)

func handlePlayerLevelStaticListExport(rw http.ResponseWriter, req *http.Request) {
	form := &playerLevelStaticRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家等级统计导出，解析异常")
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
		}).Error("获取游戏玩家等级统计导出，获取服务id异常")
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
		}).Error("获取游戏玩家等级统计导出，获取统计结果异常")
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

	heads := []string{"等级", "玩家数", "玩家总占比", "玩家总数"}
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("sheet")
	row := sheet.AddRow()
	for _, value := range heads {
		row.AddCell().Value = value
	}
	for _, rowValue := range respon.ItemArray {
		newRow := sheet.AddRow()
		newRow.AddCell().Value = strconv.Itoa(int(rowValue.Level))
		newRow.AddCell().Value = strconv.Itoa(int(rowValue.PlayerCount))
		newRow.AddCell().Value = strconv.Itoa(int(rowValue.Rate))
		countNewRow := newRow.AddCell()
		countNewRow.SetInt(int(rowValue.TotalCount))
	}
	var buffer bytes.Buffer
	if err := file.Write(&buffer); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家等级统计导出，写入缓存失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	// r := bytes.NewReader(buffer.Bytes())

	rw.Header().Add("Content-Disposition", "attachment")
	rw.Header().Add("Content-Type", "application/vnd.ms-excel")
	rw.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	rw.Write(buffer.Bytes())
}
