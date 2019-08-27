package api

import (
	"fgame/fgame/gm/gamegm/gm/center/redeem/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type redeemCodeRequest struct {
	Id int `json:"id"`
}

func handleRedeemCode(rw http.ResponseWriter, req *http.Request) {
	form := &redeemCodeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("生成兑换码，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rds := service.RedeemServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("生成兑换码，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rds.NewRedeemCode(form.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("生成兑换码，添加异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
