package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fgame/fgame/gm/gamegm/constant"
	mongoservice "fgame/fgame/gm/gamegm/mglog/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	"fgame/fgame/gm/gamegm/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type getMongoChatLogRequest struct {
	TableName   string `json:"tableName"`
	BeginTime   int64  `json:"beginTime"`
	EndTime     int64  `json:"endTime"`
	Platform    int32  `json:"platformId"`
	ServerType  int32  `json:"serverType"`
	ServerId    int32  `json:"serverId"`
	PageIndex   int    `json:"pageIndex"`
	PlayerId    string `json:"playerId"`
	ChatContent string `json:"chatContent"`
	ChatType    int    `json:"chatType"`
}

type getMongoChatLogResponItem struct {
	LogTime        int64  `json:"logTime"`
	Platform       int32  `json:"platform"`
	ServerType     int32  `json:"serverType"`
	SdkType        int32  `json:"sdkType"`
	DeviceType     int32  `json:"deviceType"`
	ServerId       int32  `json:"serverId"`
	UserId         int64  `json:"userId"`
	PlayerId       int64  `json:"playerId"`
	PlayerIdString string `json:"playerIdString"`
	Ip             string `json:"ip"`
	Name           string `json:"name"`
	Role           int32  `json:"role"`
	Sex            int32  `json:"sex"`
	Level          int32  `json:"level"`
	Vip            int32  `json:"vip"`
	Channel        int32  `json:"channel"`

	//接收者id
	RecvId int64 `json:"recvId"`

	//接收者名字
	RecvName string `json:"recvName"`

	//消息类型(0:文本,1:表情,2:语音)
	MsgType int32 `json:"msgType"`

	//内容
	Content []byte `json:"content"`

	//文本内容
	Text string `json:"text"`
}

type getMongoChatLogRespon struct {
	ItemArray  []*getMongoChatLogResponItem `json:"itemArray"`
	TotalCount int                          `json:"totalCount"`
}

func handleGetMongoChatLog(rw http.ResponseWriter, req *http.Request) {
	form := &getMongoChatLogRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("查询日志，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Debug("查询聊天日志，查询内容：", form.ChatContent)
	service := mongoservice.MgLogServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("查询日志，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if form.Platform == 0 {
		rr := gmhttp.NewFailedResultWithMsg(114, fmt.Sprintf("平台不能为空"))
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}

	searchPlayerId, _ := strconv.ParseInt(form.PlayerId, 10, 64)

	rst, err := service.GetChatLogMsg(form.TableName, form.BeginTime, form.EndTime, form.Platform, form.ServerType, form.ServerId, searchPlayerId, form.PageIndex, constant.DefaultPageSize, form.ChatContent, form.ChatType)
	if err != nil {
		log.WithFields(log.Fields{
			"tableName":  form.TableName,
			"beginTime":  form.BeginTime,
			"endTime":    form.EndTime,
			"platform":   form.Platform,
			"serverType": form.ServerType,
			"serverId":   form.ServerId,
			"pageindex":  form.PageIndex,
			"error":      err,
		}).Error("查询日志，执行查询异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	totalCount, err := service.GetChatLogMsgCount(form.TableName, form.BeginTime, form.EndTime, form.Platform, form.ServerType, form.ServerId, searchPlayerId, form.ChatContent, form.ChatType)
	if err != nil {
		log.WithFields(log.Fields{
			"tableName":  form.TableName,
			"beginTime":  form.BeginTime,
			"endTime":    form.EndTime,
			"platform":   form.Platform,
			"serverType": form.ServerType,
			"serverId":   form.ServerId,
			"pageindex":  form.PageIndex,
			"error":      err,
		}).Error("查询日志，执行查询总数异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	finnalValue := make([]*getMongoChatLogResponItem, 0)
	rstValue, err := json.Marshal(rst)
	if err != nil {
		log.Error("json转换失败：", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(rstValue, &finnalValue)
	if err != nil {
		log.Error("json转换失败：", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, value := range finnalValue {
		value.PlayerIdString = utils.ConverInt64ToString(value.PlayerId)
	}
	respon := &getMongoChatLogRespon{}
	respon.ItemArray = finnalValue
	respon.TotalCount = totalCount
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
