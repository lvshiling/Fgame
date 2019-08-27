package api

import (
	"net/http"

	gmchatSet "fgame/fgame/gm/gamegm/gm/center/chatset/service"
	us "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type chatSetListPlatformRequest struct {
	PageIndex  int `form:"pageIndex" json:"pageIndex"`
	PlatformId int `form:"centerPlatformId" json:"centerPlatformId"`
}

type chatSetListPlatformRespon struct {
	ItemArray  []*chatSetListPlatformResponItem `json:"itemArray"`
	TotalCount int                              `json:"total"`
}

type chatSetListPlatformResponItem struct {
	ChatSetId        int `form:"chatSetId" json:"chatSetId"`
	PlatformId       int `form:"centerPlatformId" json:"centerPlatformId"`
	WorldVip         int `form:"worldVip" json:"worldVip"`
	WorldPlayerLevel int `form:"worldPlayerLevel" json:"worldPlayerLevel"`
	PChatVip         int `form:"pChatVip" json:"pChatVip"`
	PChatPlayerLevel int `form:"pChatPlayerLevel" json:"pChatPlayerLevel"`
	GuildVip         int `form:"guildVip" json:"guildVip"`
	GuildPlayerLevel int `form:"guildPlayerLevel" json:"guildPlayerLevel"`
	TeamVip          int `form:"teamVip" json:"teamVip"`
	TeamPlayerLevel  int `form:"teamPlayerLevel" json:"teamPlayerLevel"`
}

func handleChatSetListPlatform(rw http.ResponseWriter, req *http.Request) {
	form := &chatSetListPlatformRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取聊天配置列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmchatSet.ChatSetServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取聊天配置列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userId := us.GmUserIdInContext(req.Context())

	usservice := us.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("获取平台列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userSdkTypeList, err := usservice.GetUserCenterPlatList(userId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": userId,
		}).Error("获取聊天配置列表，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetChatSetListPlatform(form.PlatformId, form.PageIndex, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"index": form.PageIndex,
		}).Error("获取聊天配置列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &chatSetListPlatformRespon{}
	respon.ItemArray = make([]*chatSetListPlatformResponItem, 0)

	for _, value := range rst {
		item := &chatSetListPlatformResponItem{
			ChatSetId:        value.Id,
			PlatformId:       value.PlatformId,
			WorldVip:         value.WorldVip,
			WorldPlayerLevel: value.WorldPlayerLevel,
			PChatVip:         value.PChatVip,
			PChatPlayerLevel: value.PChatPlayerLevel,
			GuildVip:         value.GuildVip,
			GuildPlayerLevel: value.GuildPlayerLevel,
			TeamVip:          value.TeamVip,
			TeamPlayerLevel:  value.TeamPlayerLevel,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := service.GetChatSetCountPlatform(form.PlatformId, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":      err,
			"PlatformId": form.PlatformId,
			"index":      form.PageIndex,
		}).Error("获取聊天配置列表，执行异常")
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
