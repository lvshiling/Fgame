package api

import (
	"fgame/fgame/gm/gamegm/gm/types"
	us "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type childPrivilegeRespon struct {
	PrivilegeKey  int    `json:"key"`
	PrivilegeName string `json:"name"`
}

func handleChildPrivilege(rw http.ResponseWriter, req *http.Request) {
	log.Debug("获取下属子角色列表")
	// userId := us.GmUserIdInContext(req.Context())
	privilege := types.PrivilegeLevel(us.PrivilegeInContext(req.Context()))
	respon := make([]*childPrivilegeRespon, 0)

	if privilege.Valid() {
		childPrivilege := privilege.ChildPrivilege()
		if childPrivilege != nil && len(childPrivilege) > 0 {
			for _, value := range childPrivilege {
				item := &childPrivilegeRespon{
					PrivilegeKey:  int(value),
					PrivilegeName: value.String(),
				}
				respon = append(respon, item)
			}
		}
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
