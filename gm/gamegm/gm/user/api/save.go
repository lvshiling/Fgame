package api

import (
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmUserModel "fgame/fgame/gm/gamegm/gm/user/model"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	gmerr "fgame/fgame/gm/gamegm/error"
	"fgame/fgame/gm/gamegm/gm/types"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type saveUserRequest struct {
	UserId         int64  `form:"userid" json:"userId"`
	Psd            string `form:"password" json:"password"`
	UserName       string `form:"userName" json:"userName"`
	PrivilegeLevel int    `form:"privilegeid" json:"privilegeid"`
	ChannelID      int64  `form:"channelId" json:"channelId"`
	PlatformID     int64  `form:"platformId" json:"platformId"`
}

func handleSaveUser(rw http.ResponseWriter, req *http.Request) {
	form := &saveUserRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("用户登陆，解析异常")
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
	privilegeLevel := types.PrivilegeLevel(int32(form.PrivilegeLevel))
	if privilegeLevel.HasChannel() && form.ChannelID < 1 {
		err = gmerr.GetError(gmerr.ErrorCodeMissChannel)
		errhttp.ResponseWithError(rw, err)
		return
	}
	if privilegeLevel.HasPlatform() && form.PlatformID < int64(1) {
		err = gmerr.GetError(gmerr.ErrorCodeMissPlatform)
		errhttp.ResponseWithError(rw, err)
		return
	}

	userInfo := &gmUserModel.DBGmUserInfo{}
	userInfo.UserId = form.UserId
	userInfo.Psd = form.Psd
	userInfo.PrivilegeLevel = form.PrivilegeLevel
	userInfo.UserName = form.UserName
	userInfo.ChannelID = form.ChannelID
	userInfo.PlatformId = form.PlatformID

	if form.UserId == 0 {
		err := service.AddUser(userInfo)
		if err != nil {
			errhttp.ResponseWithError(rw, err)
			return
		}
	} else {
		err := service.UpdateUser(userInfo)
		if err != nil {
			errhttp.ResponseWithError(rw, err)
			return
		}
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
