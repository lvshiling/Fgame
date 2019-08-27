package service

import (
	"context"
	"encoding/json"
	mongo "fgame/fgame/core/mongo"
	"fgame/fgame/gm/gamegm/gm/center/staticreport/model"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"net/http"
	"time"

	mgmodel "fgame/fgame/logserver/model"

	"github.com/codegangsta/negroni"
	"gopkg.in/mgo.v2/bson"
)

type IStaticReportService interface {
	GetOnLineStatic(p_start int64, p_end int64, p_platformId int, p_timeZone int) ([]*model.OnLineStatic, error)
	GetNeiGuaOnLineStatic(p_start int64, p_end int64, p_platformId int, p_timeZone int) ([]*model.OnLineStatic, error)
	GetrecycleStatic(p_start int64, p_end int64, p_platformId int, p_timeZone int) ([]*model.RecycleStatic, error)
	GetLastOnLineStatic(p_lastTime int64, p_platformList []int64) ([]*model.LastOnLineStatic, error)
	GetOnLinePlayerNum(p_start int64, p_end int64, p_platformId int64, p_serverId int) (int, error)
	GetOnLinePlayerNumGroupSdk(p_start int64, p_end int64, p_sdkType int, p_sdkTypeList []int) ([]*model.OnLinePlayerStaticSdk, error)
	GetGoldChangeStatic(p_start int64, p_end int64, p_platform int, p_serverid int, p_sdkType int, p_playerList []int64, p_sdkTypeList []int) ([]*model.PlayerGoldChange, error)
	GetNewBangYuanChangeStatic(p_start int64, p_end int64, p_platform int, p_serverid int, p_sdkType int, p_playerList []int64, p_sdkTypeList []int, p_goldType int) ([]*model.PlayerGoldChange, error)
	GetNewGoldChangeStatic(p_start int64, p_end int64, p_platform int, p_serverid int, p_sdkType int, p_playerList []int64, p_sdkTypeList []int, p_goldType int) ([]*model.PlayerGoldChange, error)
	GetOnLinePlayerNumGroupServer(p_start int64, p_end int64) ([]*model.OnLinePlayerStaticServer, error)
	GetOnLineServerFirstTime() (int64, error)

	GetOnLineStaticDaily(p_start int64, p_end int64) ([]*model.OnLinePlayerStaticDaily, error)
	GetServerLoginPlayerNum(p_start int64, p_end int64) ([]*model.ServerPlayerLoginStatic, error)
}

type staticReportService struct {
	mongoService mongo.MongoService
	mongoConfig  *mongo.MongoConfig
}

func (m *staticReportService) GetOnLineStatic(p_start int64, p_end int64, p_platformId int, p_timeZone int) ([]*model.OnLineStatic, error) {
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("system_online")
	if c == nil {
		return nil, nil
	}
	timeZoneMilSecond := p_timeZone * 60 * 60 * 1000
	projectMatch := bson.M{
		"$project": bson.M{
			"logtime":    "$logtime",
			"onlinenum":  "$onlinenum",
			"serverid":   "$serverid",
			"platform":   "$platform",
			"servertype": "$servertype",
			"minuteindex": bson.M{
				"$trunc": bson.M{
					"$divide": []interface{}{bson.M{"$mod": []interface{}{bson.M{"$add": []interface{}{"$logtime", timeZoneMilSecond}}, 86400000}}, 300000},
				},
			},
			"datestr": bson.M{
				"$trunc": bson.M{
					"$divide": []interface{}{bson.M{"$add": []interface{}{"$logtime", timeZoneMilSecond}}, 86400000},
				},
			},
		},
	}
	whereMatch := bson.M{
		"$match": bson.M{
			"logtime": bson.M{
				"$gte": p_start,
				"$lt":  p_end,
			},
			"platform":   p_platformId,
			"servertype": 0,
		},
	}
	groupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"serverid":    "$serverid",
				"minuteindex": "$minuteindex",
				"datestr":     "$datestr",
				"platform":    "$platform",
			},
			"maxplayer": bson.M{
				"$max": "$onlinenum",
			},
		},
	}
	if groupMatch != nil {

	}
	operations := []bson.M{projectMatch, whereMatch, groupMatch}
	pipe := c.Pipe(operations)
	// rst := make([]*model.OneLineStatic, 0)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	rst := make([]*model.OnLineStatic, 0)
	err = json.Unmarshal(rstByte, &rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (m *staticReportService) GetNeiGuaOnLineStatic(p_start int64, p_end int64, p_platformId int, p_timeZone int) ([]*model.OnLineStatic, error) {
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("system_neigua_online")
	if c == nil {
		return nil, nil
	}
	timeZoneMilSecond := p_timeZone * 60 * 60 * 1000
	projectMatch := bson.M{
		"$project": bson.M{
			"logtime":    "$logtime",
			"onlinenum":  "$onlinenum",
			"serverid":   "$serverid",
			"platform":   "$platform",
			"servertype": "$servertype",
			"minuteindex": bson.M{
				"$trunc": bson.M{
					"$divide": []interface{}{bson.M{"$mod": []interface{}{bson.M{"$add": []interface{}{"$logtime", timeZoneMilSecond}}, 86400000}}, 300000},
				},
			},
			"datestr": bson.M{
				"$trunc": bson.M{
					"$divide": []interface{}{bson.M{"$add": []interface{}{"$logtime", timeZoneMilSecond}}, 86400000},
				},
			},
		},
	}
	whereMatch := bson.M{
		"$match": bson.M{
			"logtime": bson.M{
				"$gte": p_start,
				"$lt":  p_end,
			},
			"platform":   p_platformId,
			"servertype": 0,
		},
	}
	groupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"serverid":    "$serverid",
				"minuteindex": "$minuteindex",
				"datestr":     "$datestr",
				"platform":    "$platform",
			},
			"maxplayer": bson.M{
				"$max": "$onlinenum",
			},
		},
	}
	if groupMatch != nil {

	}
	operations := []bson.M{projectMatch, whereMatch, groupMatch}
	pipe := c.Pipe(operations)
	// rst := make([]*model.OneLineStatic, 0)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	rst := make([]*model.OnLineStatic, 0)
	err = json.Unmarshal(rstByte, &rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (m *staticReportService) GetrecycleStatic(p_start int64, p_end int64, p_platformId int, p_timeZone int) ([]*model.RecycleStatic, error) {
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("system_trade_recycle")
	if c == nil {
		return nil, nil
	}
	timeZoneMilSecond := p_timeZone * 60 * 60 * 1000
	projectMatch := bson.M{
		"$project": bson.M{
			"logtime":     "$logtime",
			"recyclegold": "$recyclegold",
			"serverid":    "$serverid",
			"platform":    "$platform",
			"servertype":  "$servertype",
			"minuteindex": bson.M{
				"$trunc": bson.M{
					"$divide": []interface{}{bson.M{"$mod": []interface{}{bson.M{"$add": []interface{}{"$logtime", timeZoneMilSecond}}, 86400000}}, 300000},
				},
			},
			"datestr": bson.M{
				"$trunc": bson.M{
					"$divide": []interface{}{bson.M{"$add": []interface{}{"$logtime", timeZoneMilSecond}}, 86400000},
				},
			},
		},
	}
	whereMatch := bson.M{
		"$match": bson.M{
			"logtime": bson.M{
				"$gte": p_start,
				"$lt":  p_end,
			},
			"platform":   p_platformId,
			"servertype": 0,
		},
	}
	groupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"serverid":    "$serverid",
				"minuteindex": "$minuteindex",
				"datestr":     "$datestr",
				"platform":    "$platform",
			},
			"recyclegold": bson.M{
				"$max": "$recyclegold",
			},
		},
	}
	if groupMatch != nil {

	}
	operations := []bson.M{projectMatch, whereMatch, groupMatch}
	pipe := c.Pipe(operations)
	// rst := make([]*model.OneLineStatic, 0)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	rst := make([]*model.RecycleStatic, 0)
	err = json.Unmarshal(rstByte, &rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (m *staticReportService) GetLastOnLineStatic(p_lastTime int64, p_platformList []int64) ([]*model.LastOnLineStatic, error) {
	rst := make([]*model.LastOnLineStatic, 0)
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("system_online")
	if c == nil {
		return rst, nil
	}

	projectMatch := bson.M{
		"$project": bson.M{
			"logtime":    "$logtime",
			"onlinenum":  "$onlinenum",
			"serverid":   "$serverid",
			"platform":   "$platform",
			"servertype": "$servertype",
		},
	}

	now := timeutils.TimeToMillisecond(time.Now())
	tempWhere := bson.M{
		"logtime": bson.M{
			"$gte": p_lastTime,
			"$lt":  now,
		},
		"servertype": 0,
	}

	if len(p_platformList) > 0 {
		tempWhere["platform"] = bson.M{"$in": p_platformList}
	}

	whereMatch := bson.M{
		"$match": tempWhere,
	}

	groupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"serverid": "$serverid",
				"platform": "$platform",
			},
			"lastplayer": bson.M{
				"$last": "$onlinenum",
			},
		},
	}
	operations := []bson.M{projectMatch, whereMatch, groupMatch}
	pipe := c.Pipe(operations)
	// rst := make([]*model.OneLineStatic, 0)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rstByte, &rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (m *staticReportService) GetOnLinePlayerNum(p_start int64, p_end int64, p_platformId int64, p_serverId int) (int, error) {
	rst := make([]*model.OnLinePlayerStatic, 0)
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("player_login")
	if c == nil {
		return 0, nil
	}

	projectMatch := bson.M{
		"$project": bson.M{
			"playerid":   "$playerid",
			"logtime":    "$logtime",
			"platform":   "$platform",
			"serverid":   "$serverid",
			"servertype": "$servertype",
		},
	}

	tempWhere := bson.M{
		"logtime": bson.M{
			"$gte": p_start,
			"$lt":  p_end,
		},
		"platform":   p_platformId,
		"serverid":   p_serverId,
		"servertype": 0,
	}

	whereMatch := bson.M{
		"$match": tempWhere,
	}

	distinctGroupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"playerid": "$playerid",
			},
			"count": bson.M{
				"$sum": 1,
			},
		},
	}
	countGroupMath := bson.M{
		"$group": bson.M{
			"_id": 1,
			"count": bson.M{
				"$sum": 1,
			},
		},
	}
	operations := []bson.M{projectMatch, whereMatch, whereMatch, distinctGroupMatch, countGroupMath}
	pipe := c.Pipe(operations)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return 0, err
	}
	rstByte, err := json.Marshal(results)
	fmt.Println(string(rstByte))
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(rstByte, &rst)
	if err != nil {
		return 0, err
	}
	if len(rst) > 0 {
		return rst[0].Count, nil
	}
	return 0, nil
}

func (m *staticReportService) GetOnLinePlayerNumGroupSdk(p_start int64, p_end int64, p_sdkType int, p_sdkTypeList []int) ([]*model.OnLinePlayerStaticSdk, error) {
	rst := make([]*model.OnLinePlayerStaticSdk, 0)
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("player_login")
	if c == nil {
		return rst, nil
	}

	projectMatch := bson.M{
		"$project": bson.M{
			"playerid":   "$playerid",
			"logtime":    "$logtime",
			"platform":   "$platform",
			"serverid":   "$serverid",
			"servertype": "$servertype",
			"sdktype":    "$sdktype",
		},
	}

	tempWhere := bson.M{
		"logtime": bson.M{
			"$gte": p_start,
			"$lt":  p_end,
		},
		"servertype": 0,
	}
	if p_sdkType > 0 {
		tempWhere["sdktype"] = p_sdkType
	} else {
		if len(p_sdkTypeList) > 0 {
			tempWhere["sdktype"] = bson.M{"$in": p_sdkTypeList}
		}
	}

	whereMatch := bson.M{
		"$match": tempWhere,
	}

	distinctGroupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"playerid": "$playerid",
				"platform": "$sdktype",
			},
			"count": bson.M{
				"$sum": 1,
			},
		},
	}
	countGroupMath := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"platform": "$_id.platform",
			},
			"sdkcount": bson.M{
				"$sum": 1,
			},
		},
	}
	operations := []bson.M{projectMatch, whereMatch, whereMatch, distinctGroupMatch, countGroupMath}
	pipe := c.Pipe(operations)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	fmt.Println(string(rstByte))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rstByte, &rst)
	return rst, nil
}

func (m *staticReportService) GetGoldChangeStatic(p_start int64, p_end int64, p_platform int, p_serverid int, p_sdkType int, p_playerList []int64, p_sdkTypeList []int) ([]*model.PlayerGoldChange, error) {
	rst := make([]*model.PlayerGoldChange, 0)
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("player_glod_changed")
	if c == nil {
		return rst, nil
	}

	projectMatch := bson.M{
		"$project": bson.M{
			"playerid":   "$playerid",
			"logtime":    "$logtime",
			"platform":   "$platform",
			"serverid":   "$serverid",
			"servertype": "$servertype",
			"sdktype":    "$sdktype",
			"reason":     "$reason",
			"changednum": "$changednum",
		},
	}

	tempWhere := bson.M{
		"logtime": bson.M{
			"$gte": p_start,
			"$lt":  p_end,
		},
		"servertype": 0,
	}
	if p_sdkType > 0 {
		tempWhere["sdktype"] = p_sdkType
	} else {
		if len(p_sdkTypeList) > 0 {
			tempWhere["sdktype"] = bson.M{"$in": p_sdkTypeList}
		}
	}
	if len(p_playerList) > 0 {
		tempWhere["playerid"] = bson.M{"$in": p_playerList}
	}

	if p_platform > 0 {
		tempWhere["platform"] = p_platform
	}
	if p_serverid > 0 {
		tempWhere["serverid"] = p_serverid
	}

	whereMatch := bson.M{
		"$match": tempWhere,
	}

	distinctGroupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"playerid": "$playerid",
				"platform": "$platform",
				"serverid": "$serverid",
				"reason":   "$reason",
			},
			"changednum": bson.M{
				"$sum": "$changednum",
			},
		},
	}
	countGroupMath := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"platform": "$_id.platform",
				"serverid": "$_id.serverid",
				"reason":   "$_id.reason",
			},
			"playercount": bson.M{
				"$sum": 1,
			},
			"changednum": bson.M{
				"$sum": "$changednum",
			},
		},
	}
	operations := []bson.M{projectMatch, whereMatch, whereMatch, distinctGroupMatch, countGroupMath}
	pipe := c.Pipe(operations)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	fmt.Println(string(rstByte))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rstByte, &rst)
	return rst, nil
}

func (m *staticReportService) GetNewBangYuanChangeStatic(p_start int64, p_end int64, p_platform int, p_serverid int, p_sdkType int, p_playerList []int64, p_sdkTypeList []int, p_goldType int) ([]*model.PlayerGoldChange, error) {
	rst := make([]*model.PlayerGoldChange, 0)
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("player_new_bind_glod_changed")
	if c == nil {
		return rst, nil
	}

	projectMatch := bson.M{
		"$project": bson.M{
			"playerid":   "$playerid",
			"logtime":    "$logtime",
			"platform":   "$platform",
			"serverid":   "$serverid",
			"servertype": "$servertype",
			"sdktype":    "$sdktype",
			"reason":     "$reason",
			"changednum": "$changednum",
		},
	}

	tempWhere := bson.M{
		"logtime": bson.M{
			"$gte": p_start,
			"$lt":  p_end,
		},
		"servertype": 0,
	}
	if p_sdkType > 0 {
		tempWhere["sdktype"] = p_sdkType
	} else {
		if len(p_sdkTypeList) > 0 {
			tempWhere["sdktype"] = bson.M{"$in": p_sdkTypeList}
		}
	}
	if len(p_playerList) > 0 {
		tempWhere["playerid"] = bson.M{"$in": p_playerList}
	}
	if p_goldType == 1 {
		tempWhere["changednum"] = bson.M{"$gte": 0}
	}
	if p_goldType == 2 {
		tempWhere["changednum"] = bson.M{"$lt": 0}
	}

	if p_platform > 0 {
		tempWhere["platform"] = p_platform
	}
	if p_serverid > 0 {
		tempWhere["serverid"] = p_serverid
	}

	whereMatch := bson.M{
		"$match": tempWhere,
	}

	distinctGroupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"playerid": "$playerid",
				"platform": "$platform",
				"serverid": "$serverid",
				"reason":   "$reason",
			},
			"changednum": bson.M{
				"$sum": "$changednum",
			},
		},
	}
	countGroupMath := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"platform": "$_id.platform",
				"serverid": "$_id.serverid",
				"reason":   "$_id.reason",
			},
			"playercount": bson.M{
				"$sum": 1,
			},
			"changednum": bson.M{
				"$sum": "$changednum",
			},
		},
	}
	operations := []bson.M{projectMatch, whereMatch, whereMatch, distinctGroupMatch, countGroupMath}
	pipe := c.Pipe(operations)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	fmt.Println(string(rstByte))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rstByte, &rst)
	return rst, nil
}

func (m *staticReportService) GetNewGoldChangeStatic(p_start int64, p_end int64, p_platform int, p_serverid int, p_sdkType int, p_playerList []int64, p_sdkTypeList []int, p_goldType int) ([]*model.PlayerGoldChange, error) {
	rst := make([]*model.PlayerGoldChange, 0)
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("player_new_glod_changed")
	if c == nil {
		return rst, nil
	}

	projectMatch := bson.M{
		"$project": bson.M{
			"playerid":   "$playerid",
			"logtime":    "$logtime",
			"platform":   "$platform",
			"serverid":   "$serverid",
			"servertype": "$servertype",
			"sdktype":    "$sdktype",
			"reason":     "$reason",
			"changednum": "$changednum",
		},
	}

	tempWhere := bson.M{
		"logtime": bson.M{
			"$gte": p_start,
			"$lt":  p_end,
		},
		"servertype": 0,
		"changednum": bson.M{
			"$gte": 0,
		},
	}
	if p_sdkType > 0 {
		tempWhere["sdktype"] = p_sdkType
	} else {
		if len(p_sdkTypeList) > 0 {
			tempWhere["sdktype"] = bson.M{"$in": p_sdkTypeList}
		}
	}
	if len(p_playerList) > 0 {
		tempWhere["playerid"] = bson.M{"$in": p_playerList}
	}

	if p_platform > 0 {
		tempWhere["platform"] = p_platform
	}
	if p_serverid > 0 {
		tempWhere["serverid"] = p_serverid
	}

	if p_goldType == 1 {
		tempWhere["changednum"] = bson.M{"$gte": 0}
	}
	if p_goldType == 2 {
		tempWhere["changednum"] = bson.M{"$lt": 0}
	}

	whereMatch := bson.M{
		"$match": tempWhere,
	}

	distinctGroupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"playerid": "$playerid",
				"platform": "$platform",
				"serverid": "$serverid",
				"reason":   "$reason",
			},
			"changednum": bson.M{
				"$sum": "$changednum",
			},
		},
	}
	countGroupMath := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"platform": "$_id.platform",
				"serverid": "$_id.serverid",
				"reason":   "$_id.reason",
			},
			"playercount": bson.M{
				"$sum": 1,
			},
			"changednum": bson.M{
				"$sum": "$changednum",
			},
		},
	}
	operations := []bson.M{projectMatch, whereMatch, whereMatch, distinctGroupMatch, countGroupMath}
	pipe := c.Pipe(operations)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	fmt.Println(string(rstByte))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rstByte, &rst)
	return rst, nil
}

func (m *staticReportService) GetOnLinePlayerNumGroupServer(p_start int64, p_end int64) ([]*model.OnLinePlayerStaticServer, error) {
	rst := make([]*model.OnLinePlayerStaticServer, 0)
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("player_login")
	if c == nil {
		return rst, nil
	}

	projectMatch := bson.M{
		"$project": bson.M{
			"playerid":   "$playerid",
			"logtime":    "$logtime",
			"platform":   "$platform",
			"serverid":   "$serverid",
			"servertype": "$servertype",
			"sdktype":    "$sdktype",
		},
	}

	tempWhere := bson.M{
		"logtime": bson.M{
			"$gte": p_start,
			"$lt":  p_end,
		},
		"servertype": 0,
	}
	whereMatch := bson.M{
		"$match": tempWhere,
	}

	distinctGroupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"playerid": "$playerid",
				"platform": "$platform",
				"serverid": "$serverid",
			},
			"maxlogintime": bson.M{
				"$max": "$logtime",
			},
		},
	}
	operations := []bson.M{projectMatch, whereMatch, whereMatch, distinctGroupMatch}
	pipe := c.Pipe(operations)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	fmt.Println(string(rstByte))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rstByte, &rst)
	return rst, nil
}

func (m *staticReportService) GetOnLineServerFirstTime() (int64, error) {
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("player_login")
	firstLoginInfo := &mgmodel.PlayerLogin{}
	condition := bson.M{}
	err := c.Find(condition).Sort("+logtime").Skip(0).Limit(1).One(firstLoginInfo)
	if err != nil {
		return 0, err
	}
	return firstLoginInfo.LogTime, nil
}

func (m *staticReportService) GetOnLineStaticDaily(p_start int64, p_end int64) ([]*model.OnLinePlayerStaticDaily, error) {
	rst := make([]*model.OnLinePlayerStaticDaily, 0)
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("system_online")
	if c == nil {
		return rst, nil
	}

	projectMatch := bson.M{
		"$project": bson.M{
			"logtime":    "$logtime",
			"onlinenum":  "$onlinenum",
			"serverid":   "$serverid",
			"platform":   "$platform",
			"servertype": "$servertype",
		},
	}

	tempWhere := bson.M{
		"logtime": bson.M{
			"$gte": p_start,
			"$lt":  p_end,
		},
		"servertype": 0,
	}

	whereMatch := bson.M{
		"$match": tempWhere,
	}

	groupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"serverid": "$serverid",
				"platform": "$platform",
			},
			"maxplayer": bson.M{
				"$max": "$onlinenum",
			},
		},
	}
	operations := []bson.M{projectMatch, whereMatch, groupMatch}
	pipe := c.Pipe(operations)
	// rst := make([]*model.OneLineStatic, 0)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rstByte, &rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (m *staticReportService) GetServerLoginPlayerNum(p_start int64, p_end int64) ([]*model.ServerPlayerLoginStatic, error) {
	rst := make([]*model.ServerPlayerLoginStatic, 0)

	c := m.mongoService.Session().DB(m.mongoConfig.Database).C("player_login")
	if c == nil {
		return rst, nil
	}

	projectMatch := bson.M{
		"$project": bson.M{
			"playerid":   "$playerid",
			"logtime":    "$logtime",
			"platform":   "$platform",
			"serverid":   "$serverid",
			"servertype": "$servertype",
		},
	}

	tempWhere := bson.M{
		"logtime": bson.M{
			"$gte": p_start,
			"$lt":  p_end,
		},
		"servertype": 0,
	}

	whereMatch := bson.M{
		"$match": tempWhere,
	}

	distinctGroupMatch := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"playerid": "$playerid",
				"serverid": "$serverid",
				"platform": "$platform",
			},
			"count": bson.M{
				"$sum": 1,
			},
		},
	}
	countGroupMath := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"serverid": "$_id.serverid",
				"platform": "$_id.platform",
			},
			"totalplayer": bson.M{
				"$sum": 1,
			},
		},
	}
	operations := []bson.M{projectMatch, whereMatch, whereMatch, distinctGroupMatch, countGroupMath}
	pipe := c.Pipe(operations)
	results := []bson.M{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	rstByte, err := json.Marshal(results)
	fmt.Println(string(rstByte))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rstByte, &rst)
	if err != nil {
		return nil, err
	}
	if len(rst) > 0 {
		return rst, nil
	}
	return rst, nil
}

func NewReportStatic(p_mogo mongo.MongoService, p_config *mongo.MongoConfig) IStaticReportService {
	rst := &staticReportService{
		mongoService: p_mogo,
		mongoConfig:  p_config,
	}
	return rst
}

type contextKey string

const (
	staticReportServiceKey = contextKey("StaticReportService")
)

func WithStaticReportService(ctx context.Context, ls IStaticReportService) context.Context {
	return context.WithValue(ctx, staticReportServiceKey, ls)
}

func StaticReportServiceInContext(ctx context.Context) IStaticReportService {
	us, ok := ctx.Value(staticReportServiceKey).(IStaticReportService)
	if !ok {
		return nil
	}
	return us
}

func SetupStaticReportServiceHandler(ls IStaticReportService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithStaticReportService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
