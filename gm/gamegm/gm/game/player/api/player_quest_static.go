package api

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	playmodel "fgame/fgame/gm/gamegm/gm/game/player/model"
	playerservice "fgame/fgame/gm/gamegm/gm/game/player/service"
	tempservice "fgame/fgame/gm/gamegm/gm/template"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type playerQuestStaticRequest struct {
	ServerId int `json:"serverId"`
}

type playerQuestStaticRespon struct {
	ItemArray []*playerQuestStaticResponItem `json:"itemArray"`
}

type playerQuestStaticResponItem struct {
	QuestID     int32  `json:"questId"`
	QuestName   string `json:"questName"`
	Objectives  string `json:"objectives"`
	PlayerCount int32  `json:"playerCount"`
}

func handlePlayerQuestStaticList(rw http.ResponseWriter, req *http.Request) {
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

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
