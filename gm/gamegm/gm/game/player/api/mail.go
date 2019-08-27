package api

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	playerservice "fgame/fgame/gm/gamegm/gm/game/player/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	"fgame/fgame/gm/gamegm/common"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type playerEmailRequest struct {
	ServerId  int    `json:"serverId"`
	PageIndex int    `json:"pageIndex"`
	PlayerId  string `json:"playerId"`
	Begin     int64  `json:"begin"`
	End       int64  `json:"end"`
}

type playerEmailRespon struct {
	ItemArray  []*playerEmailResponItem `json:"itemArray"`
	TotalCount int                      `json:"total"`
}

type playerEmailResponItem struct {
	Id              int64  `json:"id"`
	PlayerId        int64  `json:"playerId"`
	IsRead          int64  `json:"isRead"`
	IsGetAttachment int64  `json:"isGetAttachment"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	AttachementInfo string `json:"attachementInfo"`
	UpdateTime      int64  `json:"updateTime"`
	CreateTime      int64  `json:"createTime"`
	DeleteTime      int64  `json:"deleteTime"`
}

func handlePlayerEMailList(rw http.ResponseWriter, req *http.Request) {
	form := &playerEmailRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家邮件列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := playerservice.PlayerServiceInContext(req.Context())

	rsp := &playerEmailRespon{}
	rsp.ItemArray = make([]*playerEmailResponItem, 0)
	playerId := common.ConverStringToInt64(form.PlayerId)
	rst, err := service.GetPlayerMailList(gmdb.GameDbLink(form.ServerId), playerId, form.Begin, form.End, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家组列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if rst == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家组列表异常，db数据库为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &playerEmailResponItem{
			Id:              value.Id,
			PlayerId:        value.PlayerId,
			IsRead:          value.IsRead,
			IsGetAttachment: value.IsGetAttachment,
			Title:           value.Title,
			Content:         value.Content,
			AttachementInfo: value.AttachementInfo,
			UpdateTime:      value.UpdateTime,
			CreateTime:      value.CreateTime,
			DeleteTime:      value.DeleteTime,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	count, err := service.GetPlayerMailCount(gmdb.GameDbLink(form.ServerId), playerId, form.Begin, form.End)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家组列表异常")
	}
	rsp.TotalCount = count
	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
