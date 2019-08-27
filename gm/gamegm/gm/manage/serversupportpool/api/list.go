package api

import (
	center "fgame/fgame/gm/gamegm/gm/center/server/service"
	serversupp "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type serverSupportPoolRequest struct {
	PageIndex        int `json:"pageIndex"`
	ServerId         int `json:"serverId"`
	CenterPlatformId int `json:"centerPlatformId"`
}

type serverSupportPoolRespon struct {
	ItemArray  []*serverSupportPoolResponItem `json:"itemArray"`
	TotalCount int                            `json:"total"`
}

type serverSupportPoolResponItem struct {
	Id         int64  `json:"id"`
	ServerId   int    `json:"serverId"`
	ServerName string `json:"serverName"`
	BeginGold  int    `json:"beginGold"`
	CurGold    int    `json:"curGold"`
	DelGold    int    `json:"delGold"`
	Percent    int32  `json:"percent"`
}

func handleServerSupportPoolList(rw http.ResponseWriter, req *http.Request) {
	form := &serverSupportPoolRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取扶植池列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	serverPoolService := serversupp.ServerSupportPoolInContext(req.Context())
	if serverPoolService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取扶植池列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = serverPoolService.FillAllServerPoolSet()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取扶植池列表，填充缺失的服务异常")
		return
	}
	userid := gmUserService.GmUserIdInContext(req.Context())

	usservice := gmUserService.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("获取扶植池列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userPlatformList, err := usservice.GetUserCenterPlatList(userid)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("获取扶植池列表，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rst, err := serverPoolService.GetServerSupportPoolList(form.PageIndex, form.ServerId, form.CenterPlatformId, userPlatformList)
	if err != nil {
		log.WithFields(log.Fields{
			"userId":    userid,
			"PageIndex": form.PageIndex,
			"ServerId":  form.ServerId,
			"error":     err,
		}).Error("获取扶植池列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	centerService := center.CenterServerServiceInContext(req.Context())

	respon := &serverSupportPoolRespon{}
	respon.ItemArray = make([]*serverSupportPoolResponItem, 0)
	for _, value := range rst {
		item := &serverSupportPoolResponItem{
			Id:        value.Id,
			ServerId:  value.ServerId,
			BeginGold: value.BeginGold,
			CurGold:   value.CurGold,
			DelGold:   value.DelGold,
			Percent:   value.OrderGoldPer,
		}
		cs, _ := centerService.GetCenterServer(int64(value.ServerId))
		item.ServerName = cs.ServerName
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := serverPoolService.GetServerSupportPoolCount(form.ServerId, form.CenterPlatformId, userPlatformList)
	if err != nil {
		log.WithFields(log.Fields{
			"userId":    userid,
			"PageIndex": form.PageIndex,
			"ServerId":  form.ServerId,
			"error":     err,
		}).Error("获取扶植池列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
