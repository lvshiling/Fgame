package middleware

import (
	"fgame/fgame/gm/gamegm/gm/types"
	gmservice "fgame/fgame/gm/gamegm/gm/user/service"
	gmutils "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/xozrc/pkg/httputils"
)

func PrivilegeHandlerMiddleware(privilegeType types.PrivilegeType) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {

		ds := gmservice.GmUserServiceInContext(req.Context())
		if ds == nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		privilegeLevelInt := gmservice.PrivilegeInContext(req.Context())
		privilegeLevel := types.PrivilegeLevel(privilegeLevelInt)
		if privilegeLevel.Privilege()&int64(privilegeType) == 0 {
			log.WithFields(
				log.Fields{
					"角色": privilegeLevel.String(),
					"权限": privilegeType.String(),
				}).Warn("权限校验,没有权限")
			//TODO 判断是否有权限
			r := gmutils.NewFailedResultWithMsg(ErrorCodeNoPrivilege, errorMap[ErrorCodeNoPrivilege])
			httputils.WriteJSON(rw, http.StatusOK, r)
			return
		}
		hf.ServeHTTP(rw, req)
	})
}

func PrivilegeHandler(privilegeType types.PrivilegeType, next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		PrivilegeHandlerMiddleware(privilegeType).ServeHTTP(rw, req, next)
	})
}
