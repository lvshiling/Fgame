package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"

	types "fgame/fgame/gm/gamegm/gm/types"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
)

type userListRequest struct {
	PageIndex int    `form:"pageIndex" json:"pageIndex"`
	UserName  string `form:"userName" json:"userName"`
	Privilege int    `form:"privilege" json:"privilege"`
}

type userListRespon struct {
	ItemArray  []*userListResponItem `json:"itemArray"`
	TotalCount int                   `json:"total"`
}

type userListResponItem struct {
	UserID        int64  `json:"userId"`
	UserName      string `json:"userName"`
	PrivilegeName string `json:"privilegeName"`
	Privilege     int    `json:"privilegeid"`
	PlatformId    int64  `json:"platformId"`
	ChannelId     int64  `json:"channelId"`
}

func handleUserList(rw http.ResponseWriter, req *http.Request) {
	form := &userListRequest{}
	// postbytte, posterr := ioutil.ReadAll(req.Body)
	// log.WithFields(log.Fields{
	// 	"postdata": string(postbytte),
	// 	"error":    posterr,
	// }).Debug("post数据")

	// rw.WriteHeader(http.StatusOK)
	// return
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取用户列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmUserService.GetGmUserServiceInstance()
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取用户列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userid := gmUserService.GmUserIdInContext(req.Context())
	userInfo, err := service.GetUserInfo(userid)
	if err != nil || userInfo == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取用户信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	privilege := types.PrivilegeLevel(gmUserService.PrivilegeInContext(req.Context()))
	if !privilege.Valid() {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	channelId := 0
	platformId := 0
	if privilege.HasChannel() {
		channelId = int(userInfo.ChannelID)
	}
	if privilege.HasPlatform() {
		platformId = int(userInfo.PlatformId)
	}
	childArray := make([]int, 0)
	childRst := privilege.ChildPrivilege()
	for _, value := range childRst {
		childArray = append(childArray, int(value))
	}

	rst, err := service.GetUserList(form.UserName, form.Privilege, form.PageIndex, channelId, platformId, childArray)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"userName":  form.UserName,
			"privilege": form.Privilege,
			"index":     form.PageIndex,
		}).Error("获取用户列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &userListRespon{}
	respon.ItemArray = make([]*userListResponItem, 0)

	for _, value := range rst {
		item := &userListResponItem{
			UserID:     value.UserId,
			UserName:   value.UserName,
			ChannelId:  value.ChannelID,
			PlatformId: value.PlatformId,
		}
		prig := types.PrivilegeLevel(value.PrivilegeLevel)
		item.PrivilegeName = prig.String()
		item.Privilege = value.PrivilegeLevel
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := service.GetUserCount(form.UserName, form.Privilege, channelId, platformId, childArray)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"userName":  form.UserName,
			"privilege": form.Privilege,
			"index":     form.PageIndex,
		}).Error("获取用户列表个数，执行异常")
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
