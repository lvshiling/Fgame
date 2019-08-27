package service

import (
	"context"
	"encoding/json"
	mongo "fgame/fgame/core/mongo"
	mglog "fgame/fgame/logserver/log"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"gopkg.in/mgo.v2/bson"
)

type IMgLogService interface {
	//获得日志的json字符串数据
	GetLogMsg(p_msgName string, p_beginTime int64, p_endTime int64, p_platformid int32, p_serverType int32, p_serverid int32, p_playerid int64, p_allianceId int64, p_pageIndex int, p_pageSize int) (result interface{}, err error)
	GetLogMsgCount(p_msgName string, p_beginTime int64, p_endTime int64, p_platformid int32, p_serverType int32, p_serverid int32, p_playerid int64, p_allianceId int64) (recordCount int, err error)
	GetChatLogMsg(p_msgName string, p_beginTime int64, p_endTime int64, p_platformid int32, p_serverType int32, p_serverid int32, p_playerid int64, p_pageIndex int, p_pageSize int, p_chatContent string, p_chatType int) (result interface{}, err error)
	GetChatLogMsgCount(p_msgName string, p_beginTime int64, p_endTime int64, p_platformid int32, p_serverType int32, p_serverid int32, p_playerid int64, p_chatContent string, p_chatType int) (recordCount int, err error)
}

type mgLogService struct {
	mongoService mongo.MongoService
	mongoConfig  *mongo.MongoConfig
}

func (m *mgLogService) GetLogMsg(p_msgName string, p_beginTime int64, p_endTime int64, p_platformid int32, p_serverType int32, p_serverid int32, p_playerid int64, p_allianceId int64, p_pageIndex int, p_pageSize int) (result interface{}, err error) {

	rst := mglog.GetLogMsgList(p_msgName)
	if rst == nil {
		return
	}
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C(p_msgName)
	condition := bson.M{}
	datecondition := bson.M{}
	if p_beginTime > 0 {
		datecondition["$gte"] = p_beginTime
	}
	if p_endTime > 0 {
		datecondition["$lte"] = p_endTime
	}
	if len(datecondition) > 0 {
		condition["logtime"] = datecondition
	}
	if p_platformid > 0 {
		condition["platform"] = p_platformid
	}
	if p_serverType > -1 {
		condition["servertype"] = p_serverType
	}
	if p_serverid > 0 {
		condition["serverid"] = p_serverid
	}
	if p_playerid > 0 {
		condition["playerid"] = p_playerid
	}
	if p_allianceId > 0 {
		condition["allianceid"] = p_allianceId
	}

	limit := p_pageSize
	offect := (p_pageIndex - 1) * p_pageSize
	if offect < 0 {
		offect = 0
	}

	err = c.Find(condition).Sort("-logtime").Skip(offect).Limit(limit).All(rst)
	if err != nil {
		return
	}
	result = rst
	return
}

func (m *mgLogService) GetLogMsgCount(p_msgName string, p_beginTime int64, p_endTime int64, p_platformid int32, p_serverType int32, p_serverid int32, p_playerid int64, p_allianceId int64) (recordCount int, err error) {
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C(p_msgName)
	condition := bson.M{}
	datecondition := bson.M{}
	if p_beginTime > 0 {
		datecondition["$gte"] = p_beginTime
	}
	if p_endTime > 0 {
		datecondition["$lte"] = p_endTime
	}
	if len(datecondition) > 0 {
		condition["logtime"] = datecondition
	}
	if p_platformid > 0 {
		condition["platform"] = p_platformid
	}
	if p_serverType > -1 {
		condition["servertype"] = p_serverType
	}
	if p_serverid > 0 {
		condition["serverid"] = p_serverid
	}
	if p_playerid > 0 {
		condition["playerid"] = p_playerid
	}
	if p_allianceId > 0 {
		condition["allianceid"] = p_allianceId
	}

	recordCount, err = c.Find(condition).Count()
	if err != nil {
		return
	}
	return
}

func (m *mgLogService) GetChatLogMsg1(p_msgName string, p_beginTime int64, p_endTime int64, p_platformid int32, p_serverType int32, p_serverid int32, p_playerid int64, p_pageIndex int, p_pageSize int, p_chatContent string, p_chatType int) (result interface{}, err error) {
	// rst := mglog.GetLogMsgList(p_msgName)
	// if rst == nil {
	// 	return
	// }
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C(p_msgName)
	// addFileds := bson.M{"$addFields": bson.M{"convercontent": bson.M{"$toString": "$content"}}}
	projectMatch := bson.M{
		"$project": bson.M{
			"playerid":   "$playerid",
			"logtime":    "$logtime",
			"platform":   "$platform",
			"serverid":   "$serverid",
			"servertype": "$servertype",
			"content":    "$content",
			// "convercontent": "1",
			"convercontent": bson.M{"$strLenBytes": "$content"},
			// "convercontent": bson.M{"$convert": bson.M{"input": "hello", "to": "bool"}},
		},
	}
	limitRows := bson.M{"logtime": bson.M{
		"$gte": p_beginTime,
		"$lt":  p_endTime,
	}}
	whereMatch := bson.M{
		"$match": limitRows,
	}
	operations := []bson.M{projectMatch, whereMatch}
	pipe := c.Pipe(operations)
	results := []bson.M{}
	err = pipe.All(&results)
	if err != nil {
		fmt.Println("error in method:", err)
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	fmt.Println("返回结果")
	fmt.Println(string(rstByte))
	return nil, nil
}

func (m *mgLogService) GetChatLogMsg(p_msgName string, p_beginTime int64, p_endTime int64, p_platformid int32, p_serverType int32, p_serverid int32, p_playerid int64, p_pageIndex int, p_pageSize int, p_chatContent string, p_chatType int) (result interface{}, err error) {
	rst := mglog.GetLogMsgList(p_msgName)
	if rst == nil {
		return
	}
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C(p_msgName)
	condition := bson.M{}
	datecondition := bson.M{}
	if p_beginTime > 0 {
		datecondition["$gte"] = p_beginTime
	}
	if p_endTime > 0 {
		datecondition["$lte"] = p_endTime
	}
	if len(datecondition) > 0 {
		condition["logtime"] = datecondition
	}
	if p_platformid > 0 {
		condition["platform"] = p_platformid
	}
	if p_serverType > -1 {
		condition["servertype"] = p_serverType
	}
	if p_serverid > 0 {
		condition["serverid"] = p_serverid
	}
	if p_playerid > 0 {
		condition["playerid"] = p_playerid
	}
	if len(p_chatContent) > 0 {
		condition["text"] = bson.M{"$regex": p_chatContent, "$options": "$i"}
		// condition["$expr"] = bson.M{bson.M{"$toString", "$content"}: bson.M{"$regex": p_chatContent, "$options": "$i"}}
	}

	if p_chatType > -1 {
		condition["channel"] = p_chatType
	}

	limit := p_pageSize
	offect := (p_pageIndex - 1) * p_pageSize
	if offect < 0 {
		offect = 0
	}

	err = c.Find(condition).Sort("-logtime").Skip(offect).Limit(limit).All(rst)
	if err != nil {
		return
	}
	result = rst

	return
}

func (m *mgLogService) GetChatLogMsgCount(p_msgName string, p_beginTime int64, p_endTime int64, p_platformid int32, p_serverType int32, p_serverid int32, p_playerid int64, p_chatContent string, p_chatType int) (recordCount int, err error) {
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C(p_msgName)
	condition := bson.M{}
	datecondition := bson.M{}
	if p_beginTime > 0 {
		datecondition["$gte"] = p_beginTime
	}
	if p_endTime > 0 {
		datecondition["$lte"] = p_endTime
	}
	if len(datecondition) > 0 {
		condition["logtime"] = datecondition
	}
	if p_platformid > 0 {
		condition["platform"] = p_platformid
	}
	if p_serverType > -1 {
		condition["servertype"] = p_serverType
	}
	if p_serverid > 0 {
		condition["serverid"] = p_serverid
	}
	if p_playerid > 0 {
		condition["playerid"] = p_playerid
	}
	if len(p_chatContent) > 0 {
		condition["text"] = bson.M{"$regex": p_chatContent, "$options": "$i"}
		// condition["$expr"] = bson.M{bson.M{"$toString", "$content"}: bson.M{"$regex": p_chatContent, "$options": "$i"}}
	}

	if p_chatType > -1 {
		condition["channel"] = p_chatType
	}

	recordCount, err = c.Find(condition).Count()
	if err != nil {
		return
	}
	return
}

func NewMgLogService(p_service mongo.MongoService, p_config *mongo.MongoConfig) IMgLogService {
	rst := &mgLogService{
		mongoService: p_service,
		mongoConfig:  p_config,
	}
	return rst
}

type contextKey string

const (
	mgLogServiceKey = contextKey("MgLogService")
)

func WithMgLogService(ctx context.Context, ls IMgLogService) context.Context {
	return context.WithValue(ctx, mgLogServiceKey, ls)
}

func MgLogServiceInContext(ctx context.Context) IMgLogService {
	us, ok := ctx.Value(mgLogServiceKey).(IMgLogService)
	if !ok {
		return nil
	}
	return us
}

func SetupMgLogServiceHandler(ls IMgLogService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithMgLogService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
