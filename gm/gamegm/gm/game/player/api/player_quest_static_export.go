package api

import (
	"bytes"
	gmdb "fgame/fgame/gm/gamegm/db"
	playmodel "fgame/fgame/gm/gamegm/gm/game/player/model"
	playerservice "fgame/fgame/gm/gamegm/gm/game/player/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	"net/http"
	"strconv"

	tempservice "fgame/fgame/gm/gamegm/gm/template"

	log "github.com/Sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"github.com/xozrc/pkg/httputils"
)

func handlePlayerQuestStaticExportList(rw http.ResponseWriter, req *http.Request) {
	form := &playerQuestStaticRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家任务统计，解析异常")
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
		}).Error("获取游戏玩家任务统计，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &playerQuestStaticRespon{}
	respon.ItemArray = make([]*playerQuestStaticResponItem, 0)
	rst, err := service.GetPlayerQuestPlayerCountStatic(gmdb.GameDbLink(form.ServerId), acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("获取游戏玩家任务统计，获取统计结果异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	playerQuestMap := make(map[int]*playmodel.QueryPlayerQuestStatic)
	for _, value := range rst {
		playerQuestMap[int(value.QuestId)] = value
	}

	questArray := tempservice.GetGmTemplateService().GetMainQuest()

	for _, value := range questArray {
		questPlayerCount := int32(0)
		playerQuest, exists := playerQuestMap[value.Id]
		if exists {
			questPlayerCount = playerQuest.PlayerCount
		}
		item := &playerQuestStaticResponItem{
			QuestID:     int32(value.Id),
			PlayerCount: questPlayerCount,
			QuestName:   value.Name,
			Objectives:  value.Objectives,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	heads := []string{"任务ID", "任务名", "任务文本", "玩家数"}
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("sheet")
	row := sheet.AddRow()
	for _, value := range heads {
		row.AddCell().Value = value
	}
	for _, rowValue := range respon.ItemArray {
		newRow := sheet.AddRow()
		newRow.AddCell().Value = strconv.Itoa(int(rowValue.QuestID))
		newRow.AddCell().Value = rowValue.QuestName
		newRow.AddCell().Value = rowValue.Objectives
		countCell := newRow.AddCell()
		countCell.SetInt(int(rowValue.PlayerCount))
	}
	var buffer bytes.Buffer
	if err := file.Write(&buffer); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家任务统计，写入缓存失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	// r := bytes.NewReader(buffer.Bytes())

	rw.Header().Add("Content-Disposition", "attachment")
	rw.Header().Add("Content-Type", "application/vnd.ms-excel")
	rw.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	rw.Write(buffer.Bytes())
}
