package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	managePath = "/manage"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(managePath).Subrouter()

	//以下需要加入权限
	sr.Path("/addmail").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeMailApply, http.HandlerFunc(handleAddMail)))
	sr.Path("/updatemail").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeMailApply, http.HandlerFunc(handleUpdateMail)))
	sr.Path("/deletemail").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeMailApply, http.HandlerFunc(handleDeleteMail)))
	sr.Path("/applymaillist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeMailApply, http.HandlerFunc(handleApplyList)))
	//审核
	sr.Path("/approvemail").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeMailApprove, http.HandlerFunc(handleApproveMail)))
	sr.Path("/approvemailarray").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeMailApprove, http.HandlerFunc(handleApproveMailArray)))
	sr.Path("/sendmail").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeMailApprove, http.HandlerFunc(handleSendMail)))
	//客服等也需要查看,列表改为查询
	sr.Path("/approvemaillist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeMailApply, http.HandlerFunc(handleApproveList)))
}

//基础方法全部先写到路由这边

func changeInt64ToString(p_id int64) string {
	return strconv.FormatInt(p_id, 10)
}

func changeStringToInt64(p_id string) int64 {
	rst, _ := strconv.ParseInt(p_id, 10, 64)
	return rst
}
