package api

import (
	"fgame/fgame/gm/gamegm/gm/center/user/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	"fgame/fgame/gm/gamegm/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerUserListRequest struct {
	PlatformId     int    `json:"sdkType"`
	UserId         string `json:"userId"`
	PlatformUserId string `json:"platformUserId"`
	UserName       string `json:"userName"`
	PageIndex      int    `json:"pageIndex"`
}

type centerUserListRespon struct {
	ItemArray  []*centerUserListItem `json:"itemArray"`
	TotalCount int                   `json:"total"`
}

type centerUserListItem struct {
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

func handleCenterUserList(rw http.ResponseWriter, req *http.Request) {
	form := &centerUserListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("中心用户列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rds := service.CenterUserServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("中心用户列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userCenterPlatList, err := gmUserService.GetUserSdkList(req.Context())
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("中心用户列表，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userId := utils.ConverStringToInt64(form.UserId)
	rst, err := rds.GetUserList(form.PlatformId, userId, form.UserName, form.PlatformUserId, form.PageIndex, userCenterPlatList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("中心用户列表，查询异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	count, err := rds.GetUserCount(form.PlatformId, userId, form.UserName, form.PlatformUserId, userCenterPlatList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("中心用户列表，查询异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	itemList := make([]*centerUserListItem, 0)
	for _, value := range rst {
		item := &centerUserListItem{
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

	respon := &centerUserListRespon{}
	respon.ItemArray = itemList
	respon.TotalCount = count

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
