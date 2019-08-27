package api

import (
	stservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"
	"fgame/fgame/gm/gamegm/gm/types"
	us "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type lastOnLineRespoon struct {
	OnLineNum int `json:"onLineNum"`
}

const (
	beforeMinitue = 10 * int64(time.Minute/time.Millisecond)
)

func handleLastOnLineStatic(rw http.ResponseWriter, req *http.Request) {

	respon := &lastOnLineRespoon{}
	userId := us.GmUserIdInContext(req.Context())

	usservice := us.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("获取在线数量，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userPrivilege := types.PrivilegeLevel(us.PrivilegeInContext(req.Context()))
	if userPrivilege == types.PrivilegeLevelKeFu ||
		userPrivilege == types.PrivilegeLevelMinitor ||
		userPrivilege == types.PrivilegeLevelCommonKeFu {
		rr := gmhttp.NewSuccessResult(respon)
		httputils.WriteJSON(rw, http.StatusOK, rr)
	}

	userCenterPlatList, err := usservice.GetUserCenterPlatList(userId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userId,
		}).Error("获取在线数量，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	lastTime := timeutils.TimeToMillisecond(time.Now()) - beforeMinitue
	rpservice := stservice.StaticReportServiceInContext(req.Context())
	rst, err := rpservice.GetLastOnLineStatic(lastTime, userCenterPlatList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取在线列表，获取数据异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	totalCount := 0
	if len(rst) > 0 {
		for _, value := range rst {
			totalCount += value.MaxPlayer
		}
	}
	respon.OnLineNum = totalCount

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
