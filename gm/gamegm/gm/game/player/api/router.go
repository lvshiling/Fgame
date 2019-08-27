package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	userPath = "/player"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(userPath).Subrouter()

	//以下需要加入权限
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handlePlayerList)))
	sr.Path("/playermail").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handlePlayerEMailList)))
	sr.Path("/mongolog").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handlePlayerMongoLogList)))
	sr.Path("/itemchange").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handlePlayerItemChangeList)))
	sr.Path("/levelcount").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handlePlayerLevelStaticList)))
	sr.Path("/levelcountexport").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handlePlayerLevelStaticListExport)))
	sr.Path("/questcount").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handlePlayerQuestStaticList)))
	sr.Path("/questcountexport").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handlePlayerQuestStaticExportList)))
	//聊天监控
	sr.Path("/fengjinlist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleFengJinList)))
	sr.Path("/jinyanlist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleJinYanList)))
	sr.Path("/forbid").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleForbid)))
	sr.Path("/forbidchat").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleForbidChat)))
	sr.Path("/unforbid").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleUnForbid)))
	sr.Path("/unforbidchat").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleUnForbidChat)))
	sr.Path("/unIgnorechat").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleUnIgnore)))
	sr.Path("/ignorechat").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleIgnoreChat)))
	sr.Path("/ignoreList").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleIgnoreList)))
	sr.Path("/playerinfo").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handlePlayerInfo)))
	sr.Path("/kickout").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleKickOut)))
}

//基础方法全部先写到路由这边

func changeInt64ToString(p_id int64) string {
	return strconv.FormatInt(p_id, 10)
}

func changeStringToInt64(p_id string) int64 {
	rst, _ := strconv.ParseInt(p_id, 10, 64)
	return rst
}
