package middleware

import (
	gmuserservice "fgame/fgame/gm/gamegm/gm/user/service"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
)

//这个为了让前端模板一样少改点，没节操了，直接改服务器了
func AuthHandlerMiddleware() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		//TODO 处理不需要认证的
		if strings.HasPrefix(req.URL.Path, "/api/gm/user/login") || strings.HasPrefix(req.URL.Path, "/websocket") {
			hf.ServeHTTP(rw, req)
			return
		}
		if !strings.HasPrefix(req.URL.Path, "/api/gm") {
			hf.ServeHTTP(rw, req)
			return
		}
		ds := gmuserservice.LoginServiceInContext(req.Context())
		if ds == nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		ah := req.Header.Get("Authorization")
		// log.Debug("token:" + ah)
		// contentTyep := req.Header.Get("Content-Type")
		// log.Debug("auth_Content-Type:" + contentTyep)
		// if len(ah) > 7 {
		// 	log.Debug("ah[7:]:", ah[7:])
		// }
		if ah != "" {
			// Should be a bearer token
			if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
				id, privilege, err := ds.VerifyToken(ah[7:])
				if err != nil {
					log.WithFields(
						log.Fields{
							"token": ah,
							"error": err,
						}).Warn("认证失败")
					// rr := gmhttp.NewFailedResultWithMsg(50008, "非法Token")
					// httputils.WriteJSON(rw, http.StatusOK, rr)
					rw.WriteHeader(http.StatusUnauthorized)
					return
				}
				if id == 0 {
					log.WithFields(
						log.Fields{
							"token": ah,
						}).Warn("认证失败,id为空")
					rw.WriteHeader(http.StatusUnauthorized)
					// rr := gmhttp.NewFailedResultWithMsg(50014, "Token已过期")
					// httputils.WriteJSON(rw, http.StatusOK, rr)
					return
				}
				ctx := gmuserservice.WithGmUserId(req.Context(), id)
				ctx = gmuserservice.WithPrivilege(ctx, privilege)
				nreq := req.WithContext(ctx)
				hf.ServeHTTP(rw, nreq)
				return
			}
		}
		// log.Debug("认证这里失败")
		// rw.WriteHeader(http.StatusUnauthorized)
		rw.WriteHeader(http.StatusUnauthorized)
		// rr := gmhttp.NewFailedResultWithMsg(50008, "非法Token")
		// httputils.WriteJSON(rw, http.StatusOK, rr)
	})
}
