package openapi

import (
	"fgame/fgame/gm/gamegm/gm/openapi/pool"

	"fgame/fgame/gm/gamegm/gm/openapi/player"

	"github.com/gorilla/mux"
)

const (
	openApiPath = "/openapi"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(openApiPath).Subrouter()
	pool.Router(sr)
	player.Router(sr)
}
