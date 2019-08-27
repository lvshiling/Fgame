package api

import (
	"net/http"
	"strings"

	"fgame/fgame/login/login"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type verifyForm struct {
	Token string `form:"token" json:"token"`
}

type VerifyResponse struct {
	PlayerId int64 `json:"playerId"`
}

func handleVerify(rw http.ResponseWriter, req *http.Request) {
	log.Debug("验证")
	loginForm := &verifyForm{}
	if err := httputils.Bind(req, loginForm); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.WithFields(log.Fields{
			"error": err,
		}).Error("验证,解析失败")
		return
	}
	token := loginForm.Token
	token = strings.TrimSpace(token)

	ls := login.LoginServiceInContext(req.Context())

	pId, err := ls.Verify(token)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"token": token,
			"error": err,
		}).Error("验证,失败")
		return
	}

	lr := &VerifyResponse{}
	lr.PlayerId = pId

	rr := RestResult{}
	rr.ErrorCode = 0
	rr.Result = lr
	httputils.WriteJSON(rw, http.StatusOK, rr)
	log.WithFields(
		log.Fields{
			"token": token,
		},
	).Debug("验证")
}
