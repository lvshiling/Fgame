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

type playerListRequest struct {
	PageIndex   int    `json:"pageIndex"`
	PlayerName  string `json:"playerName"`
	ServerId    int    `json:"serverId"`
	OrderColumn int    `json:"ordercol"`
	OrderType   int    `json:"ordertype"`
}

type playerListRespon struct {
	ItemArray  []*playerListResponItem `json:"itemArray"`
	TotalCount int                     `json:"total"`
}

type playerListResponItem struct {
	Id                       string `json:"id"`
	UserId                   int64  `json:"userId"`
	ServerId                 int64  `json:"serverId"`
	Name                     string `json:"name"`
	Role                     int    `json:"role"`
	Sex                      int    `json:"sex"`
	LastLoginTime            int64  `json:"lastLoginTime"`
	LastLogoutTime           int64  `json:"lastLogoutTime"`
	OnlineTime               int64  `json:"onlineTime"`
	OfflineTime              int64  `json:"offlineTime"`
	TotalOnlineTime          int64  `json:"totalOnlineTime"`
	TodayOnlineTime          int64  `json:"todayOnlineTime"`
	UpdateTime               int64  `json:"updateTime"`
	CreateTime               int64  `json:"createTime"`
	DeleteTime               int64  `json:"deleteTime"`
	Forbid                   int    `json:"forbid"`
	Level                    int    `json:"level"`
	ZhuanSheng               int    `json:"zhuanSheng"`
	Silver                   int    `json:"silver"`
	Gold                     int    `json:"gold"`
	BindGold                 int    `json:"bindGold"`
	Yuanshi                  int    `json:"yuanshi"`
	AllianceName             string `json:"allianceName"`
	SpouseName               string `json:"spouseName"`
	Charm                    int    `json:"charm"`
	Power                    int    `json:"power"`
	TotalChargeMoney         int    `json:"totalChargeMoney"`
	TotalChargeGold          int    `json:"totalChargeGold"`
	TotalPrivilegeChargeGold int    `json:"totalPrivilegeChargeGold"`
	Ip                       string `json:"ip"`
	OriginServerId           int32  `json:"originServerId"`
	TodayChargeMoney         int64  `json:"todayChargeMoney"`
	YesterdayChargeMoney     int64  `json:"yesterdayChargeMoney"`
	SdkType                  int32  `json:"sdkType"`
}

func handlePlayerList(rw http.ResponseWriter, req *http.Request) {
	form := &playerListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家列表，解析异常")
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
		}).Error("获取游戏玩家列表，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rsp := &playerListRespon{}
	rsp.ItemArray = make([]*playerListResponItem, 0)

	rst, err := service.GetPlayerList(gmdb.GameDbLink(form.ServerId), acServerId, form.PageIndex, form.PlayerName, form.OrderColumn, form.OrderType)
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
		item := &playerListResponItem{
			Id:                       changeInt64ToString(value.Id),
			UserId:                   value.UserId,
			ServerId:                 value.ServerId,
			Name:                     value.Name,
			Role:                     value.Role,
			Sex:                      value.Sex,
			LastLoginTime:            value.LastLoginTime,
			LastLogoutTime:           value.LastLogoutTime,
			OnlineTime:               value.OnlineTime,
			OfflineTime:              value.OfflineTime,
			TotalOnlineTime:          value.TotalOnlineTime,
			TodayOnlineTime:          value.TodayOnlineTime,
			UpdateTime:               value.UpdateTime,
			CreateTime:               value.CreateTime,
			DeleteTime:               value.DeleteTime,
			Forbid:                   value.Forbid,
			Level:                    value.Level,
			ZhuanSheng:               value.ZhuanSheng,
			Silver:                   value.Silver,
			Gold:                     value.Gold,
			BindGold:                 value.BindGold,
			Yuanshi:                  value.Yuanshi,
			AllianceName:             value.AllianceName,
			SpouseName:               value.SpouseName,
			Charm:                    value.Charm,
			Power:                    value.Power,
			TotalChargeMoney:         value.TotalChargeMoney,
			TotalChargeGold:          value.TotalChargeGold,
			TotalPrivilegeChargeGold: value.TotalPrivilegeChargeGold,
			Ip:                       value.Ip,
			OriginServerId:           value.OriginServerId,
			TodayChargeMoney:         value.TodayChargeMoney,
			YesterdayChargeMoney:     value.YesterdayChargeMoney,
			SdkType:                  value.SdkType,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	count, err := service.GetPlayerCount(gmdb.GameDbLink(form.ServerId), acServerId, form.PlayerName)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家组列表异常")
	}
	rsp.TotalCount = count
	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
