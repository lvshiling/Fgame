package api

import (
	"fgame/fgame/gm/gamegm/gm/center/redeem/pbmodel"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	"net/http"

	cenplatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	"fgame/fgame/gm/gamegm/gm/center/redeem/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleRedeemAdd(rw http.ResponseWriter, req *http.Request) {
	form := &pbmodel.RedeemInfo{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加兑换码，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rds := service.RedeemServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加兑换码，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(form.SdkTypes) == 0 {
		userid := gmUserService.GmUserIdInContext(req.Context())

		usservice := gmUserService.GmUserServiceInContext(req.Context())
		if usservice == nil {
			log.WithFields(log.Fields{}).Error("添加兑换码，用户服务为空")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		userSdkList, err := usservice.GetUserSdkTypeList(userid)
		if err != nil {
			log.WithFields(log.Fields{
				"error":  err,
				"userId": userid,
			}).Error("添加兑换码，获取权限中心平台列表异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		form.SdkTypes = userSdkList
	}

	censervice := cenplatform.CenterPlatformServiceInContext(req.Context())
	centerplatformList, err := censervice.GetPlatformIdBySdkType(form.SdkTypes)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加兑换码，获取sdk中心服id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rds.AddRedeem(form, centerplatformList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加兑换码，添加异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
