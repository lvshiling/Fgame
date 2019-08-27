package api

import (
	supportplayer "fgame/fgame/gm/gamegm/gm/manage/supportplayer/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type supportPlayerLogRequest struct {
	ChannelId  int    `json:"channelId"`
	PlatformId int    `json:"platformId"`
	ServerId   int    `json:"serverId"`
	PlayerName string `json:"playerName"`
	PlayerId   string `json:"playerId"`
	PageIndex  int    `json:"pageIndex"`
}

type supportPlayerLogRespon struct {
	ItemArray  []*supportPlayerLogResponItem `json:"itemArray"`
	TotalCount int                           `json:"total"`
}

type supportPlayerLogResponItem struct {
	Id               int64  `json:"id"`
	PlayerId         string `json:"playerId"`
	PlayerName       string `json:"playerName"`
	ChannelId        int    `json:"channelId"`
	PlatformId       int    `json:"platformId"`
	CenterPlatformId int    `json:"centerPlatformId"`
	ServerId         int    `json:"serverId"`
	ServerName       string `json:"serverName"`
	Gold             int    `json:"gold"`
	ChargeTime       int64  `json:"chargeTime"`
	UserName         string `json:"userName"`
	Reason           string `json:"reason"`
	UpdateTime       int64  `json:"updateTime"`
	CreateTime       int64  `json:"createTime"`
	DeleteTime       int64  `json:"deleteTime"`
}

func handleSupportPlayerLogList(rw http.ResponseWriter, req *http.Request) {
	log.Debug("扶持玩家日志列表")
	form := &supportPlayerLogRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("扶持玩家日志列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userid := gmUserService.GmUserIdInContext(req.Context())

	usservice := gmUserService.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("扶持玩家日志列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userPlatformList, err := usservice.GetUserCenterPlatList(userid)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("扶持玩家日志列表，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := supportplayer.SupportPlayerServiceInContext(req.Context())

	respon := &supportPlayerLogRespon{}
	respon.ItemArray = make([]*supportPlayerLogResponItem, 0)

	playerId := changeStringToInt64(form.PlayerId)
	logList, err := service.GetChargeLogList(form.ChannelId, form.PlatformId, form.ServerId, form.PlayerName, playerId, form.PageIndex, userPlatformList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("扶持玩家日志列表，获取列表数据异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range logList {
		item := &supportPlayerLogResponItem{
			Id:               value.Id,
			PlayerId:         changeInt64ToString(value.PlayerId),
			PlayerName:       value.PlayerName,
			ChannelId:        value.ChannelId,
			PlatformId:       value.PlatformId,
			CenterPlatformId: value.CenterPlatformId,
			ServerId:         value.ServerId,
			ServerName:       value.ServerName,
			Gold:             value.Gold,
			ChargeTime:       value.ChargeTime,
			UserName:         value.UserName,
			Reason:           value.Reason,
			UpdateTime:       value.UpdateTime,
			CreateTime:       value.CreateTime,
			DeleteTime:       value.DeleteTime,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}
	count, err := service.GetChargeLogCount(form.ChannelId, form.PlatformId, form.ServerId, form.PlayerName, playerId, userPlatformList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("扶持玩家日志列表，获取列表个数数据异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon.TotalCount = count

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
