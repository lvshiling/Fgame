package utils

import (
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fmt"
	"net/http"

	gmerror "fgame/fgame/gm/gamegm/error"

	"github.com/xozrc/pkg/httputils"
)

func ResponseWithError(rw http.ResponseWriter, err error) {
	switch terr := err.(type) {
	case gmerror.Error:
		{

			rr := gmhttp.NewFailedResultWithMsg(int(terr.Code()), terr.Error())
			httputils.WriteJSON(rw, http.StatusOK, rr)
			return
		}
	}
	rw.WriteHeader(http.StatusInternalServerError)
}

func ResponseWithErrorMessage(rw http.ResponseWriter, err error, msg string) {
	switch terr := err.(type) {
	case gmerror.Error:
		{
			errorMsg := fmt.Sprintf("%s:%s", terr.Error(), msg)
			rr := gmhttp.NewFailedResultWithMsg(int(terr.Code()), errorMsg)
			httputils.WriteJSON(rw, http.StatusOK, rr)
			return
		}
	}
	rw.WriteHeader(http.StatusInternalServerError)
}
