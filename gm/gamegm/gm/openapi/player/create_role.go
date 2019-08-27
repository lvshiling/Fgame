package player

import (
	"net/http"

	gmerror "fgame/fgame/gm/gamegm/error"
	gmerrutil "fgame/fgame/gm/gamegm/error/utils"
	paramapi "fgame/fgame/gm/gamegm/gm/openapi/param"

	plservice "fgame/fgame/gm/gamegm/gm/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	remotemodel "fgame/fgame/gm/gamegm/remote/model"
	remote "fgame/fgame/gm/gamegm/remote/service"
	gmutils "fgame/fgame/gm/gamegm/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type openApiPlayerCreateRoleRespon struct {
	UserId string `json:"userId"`
}

func handlePlayerAddRole(rw http.ResponseWriter, req *http.Request) {
	log.Debug("api创建角色,handleAddPoolPlayer...")
	postData := paramapi.ApiDataInContext(req.Context())
	if postData == nil {
		err := gmerror.GetError(gmerror.ErrorCodeOpenApiParam)
		gmerrutil.ResponseWithError(rw, err)
		return
	}
	serverId, err := gmutils.ConverStringToInt32Error(postData["serverId"])
	if err != nil {
		gmerrutil.ResponseWithError(rw, gmerror.GetError(gmerror.ErrorCodeOpenApiParam))
		return
	}
	name := postData["name"]
	password := postData["password"]
	userId := postData["userId"]
	if len(name) == 0 || len(password) == 0 || len(userId) == 0 {
		gmerrutil.ResponseWithError(rw, gmerror.GetError(gmerror.ErrorCodeOpenApiParamNotEmpty))
		return
	}
	platformInfo := plservice.ApiPlatformInfoInContext(req.Context())

	rmcs := remote.CenterServiceInContext(req.Context())
	//发送去中心服登录
	centerRequest := &remotemodel.CenterGmLoginRequest{
		UserId:   userId,
		SdkType:  int32(platformInfo.SdkType),
		Name:     name,
		Password: password,
	}
	respon, err := rmcs.GMLogin(centerRequest)
	if err != nil {
		log.WithFields(log.Fields{
			"userId":   userId,
			"sdkType":  int32(platformInfo.SdkType),
			"name":     name,
			"password": password,
			"error":    err,
		}).Error("api创建角色:中心登录异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rmus := remote.UserRemoteServiceInContext(req.Context())
	err = rmus.CreateRole(serverId, respon.UserId, int32(platformInfo.SdkType))
	if err != nil {
		log.WithFields(log.Fields{
			"serverId": serverId,
			"sdkType":  int32(platformInfo.SdkType),
			"UserId":   respon.UserId,
			"error":    err,
		}).Error("api创建角色:创建角色异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	apiRepson := &openApiPlayerCreateRoleRespon{
		UserId: userId,
	}
	rr := gmhttp.NewSuccessResult(apiRepson)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
