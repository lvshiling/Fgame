package api

import (
	cmlog "fgame/fgame/common/log"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type goldTypeChangeRequest struct {
	GoldType int `json:"goldType"` //类型，1增加2减少0全部
}

func handleGoldChangeType(rw http.ResponseWriter, req *http.Request) {
	form := &goldTypeChangeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取元宝配置列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := getGoldChangeType(0)

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}

func getGoldChangeType(p_goldType int) map[int32]string {
	respon := make(map[int32]string)
	goldMap := cmlog.GoldLogReasonMap
	goldReason := cmlog.GoldLogReasonRewardMap
	if p_goldType == 1 {
		for key, value := range goldMap {
			if reasonType, ok := goldReason[key]; ok {
				if reasonType {
					respon[int32(key)] = value
				}
			}
		}
	}
	if p_goldType == 2 {
		for key, value := range goldMap {
			reasonType, ok := goldReason[key]
			if !ok {
				respon[int32(key)] = value
				continue
			}
			if ok {
				if !reasonType {
					respon[int32(key)] = value
				}
			}
		}
	}
	if p_goldType == 0 {
		for key, value := range goldMap {
			respon[int32(key)] = value
		}
	}
	return respon
}
