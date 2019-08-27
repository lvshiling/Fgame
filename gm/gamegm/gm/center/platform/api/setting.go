package api

import (
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	"fgame/fgame/gm/gamegm/gm/center/platform/types"

	"github.com/xozrc/pkg/httputils"
)

type centerPlatformSettingRespon struct {
	ItemArray []*types.SettingObject `json:"itemArray"`
}

func handleCenterPlatformSetting(rw http.ResponseWriter, req *http.Request) {
	respon := &centerPlatformSettingRespon{}
	settingArray := types.GetPlatformSetting()
	respon.ItemArray = settingArray
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
