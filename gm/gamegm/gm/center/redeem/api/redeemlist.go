package api

import (
	"fgame/fgame/gm/gamegm/gm/center/redeem/pbmodel"
	"fgame/fgame/gm/gamegm/gm/center/redeem/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type redeemRequest struct {
	Name      string `json:"name"`
	SdkType   int    `json:"sdkType"`
	PageIndex int    `json:"pageIndex"`
}

type redeemRespon struct {
	ItemArray  []*pbmodel.RedeemInfo `json:"itemArray"`
	TotalCount int                   `json:"total"`
}

func handleRedeemList(rw http.ResponseWriter, req *http.Request) {
	form := &redeemRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取兑换码列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rds := service.RedeemServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取兑换码列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userid := gmUserService.GmUserIdInContext(req.Context())

	usservice := gmUserService.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("获取兑换码列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userCenterPlatList, err := usservice.GetUserCenterPlatList(userid)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("获取兑换码列表，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := rds.GetRedeemList(form.Name, form.SdkType, form.PageIndex, userCenterPlatList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("获取兑换码列表，获取一场")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &redeemRespon{}
	respon.ItemArray = rst
	count, err := rds.GetRedeemCount(form.Name, form.SdkType, userCenterPlatList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userid,
		}).Error("获取兑换码列表，获取一场")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon.TotalCount = count

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
