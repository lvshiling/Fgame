package api

import (
	"fgame/fgame/gm/gamegm/gm/center/user/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/gm/gamegm/utils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerNeiGuaUserListRequest struct {
	PlatformId     int    `json:"sdkType"`
	UserId         string `json:"userId"`
	PlatformUserId string `json:"platformUserId"`
	UserName       string `json:"userName"`
	PageIndex      int    `json:"pageIndex"`
}

type centerNeiGuaUserListRespon struct {
	ItemArray  []*centerNeiGuaUserListItem `json:"itemArray"`
	TotalCount int                         `json:"total"`
}

type centerNeiGuaUserListItem struct {
	Id             int    `json:"id"`
	Platform       int    `json:"platform"`
	PlatformUserId string `json:"platformUserId"`
	Name           string `json:"name"`
	PhoneNum       string `json:"phoneNum"`
	IdCard         string `json:"idCard"`
	RealName       string `json:"realName"`
	RealNameState  int    `json:"realNameState"`
	UpdateTime     int64  `json:"updateTime"`
	CreateTime     int64  `json:"createTime"`
	DeleteTime     int64  `json:"deleteTime"`
	Gm             int    `json:"gm"`
}

func handleCenterNeiGuaUserList(rw http.ResponseWriter, req *http.Request) {
	form := &centerNeiGuaUserListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterNeiGuaUserList，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rds := service.CenterUserServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterNeiGuaUserList，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userCenterPlatList, err := gmUserService.GetUserSdkList(req.Context())
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterNeiGuaUserList，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userId := utils.ConverStringToInt64(form.UserId)
	rst, err := rds.GetNeiGuaUserList(form.PlatformId, userId, form.UserName, form.PlatformUserId, form.PageIndex, userCenterPlatList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterNeiGuaUserList，查询异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	count, err := rds.GetNeiGuaUserCount(form.PlatformId, userId, form.UserName, form.PlatformUserId, userCenterPlatList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterNeiGuaUserList，查询异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	itemList := make([]*centerNeiGuaUserListItem, 0)
	for _, value := range rst {
		item := &centerNeiGuaUserListItem{
			Id:             value.Id,
			Platform:       value.Platform,
			PlatformUserId: value.PlatformUserId,
			Name:           value.Name,
			PhoneNum:       value.PhoneNum,
			IdCard:         value.IdCard,
			RealName:       value.RealName,
			RealNameState:  value.RealNameState,
			UpdateTime:     value.UpdateTime,
			CreateTime:     value.CreateTime,
			DeleteTime:     value.DeleteTime,
			Gm:             value.Gm,
		}
		itemList = append(itemList, item)
	}

	respon := &centerNeiGuaUserListRespon{}
	respon.ItemArray = itemList
	respon.TotalCount = count

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
