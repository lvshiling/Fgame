package api

import (
	"fgame/fgame/gm/gamegm/gm/center/redeem/pbmodel"
	"fgame/fgame/gm/gamegm/gm/center/redeem/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type redeemCodeListRequest struct {
	Id int `json:"id"`
}

type redeemCodeListRespon struct {
	ItemArray  []*pbmodel.RedeemCodeInfo `json:"itemArray"`
	TotalCount int                       `json:"total"`
}

func handleRedeemCodeList(rw http.ResponseWriter, req *http.Request) {
	form := &redeemCodeListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("兑换码列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rds := service.RedeemServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("兑换码列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := rds.GetRedeemCodeList(form.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("兑换码列表，添加异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	count, err := rds.GetRedeemCodeCount(form.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("兑换码列表，添加异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &redeemCodeListRespon{}
	respon.ItemArray = rst
	respon.TotalCount = count

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
