package service

import (
	"context"
	"fgame/fgame/core/mongo"
	mglog "fgame/fgame/logserver/log"
	"net/http"

	"github.com/codegangsta/negroni"
	"gopkg.in/mgo.v2/bson"
)

const (
	playerItemChangeLogType = "player_item_changed"
)

type IPlayerMongoLogService interface {
	GetPlayerItemChangeLogList(p_beginTime int64, p_endTime int64, p_playerid int64, p_itemId int64, p_pageIndex int, p_pageSize int) (result interface{}, err error)
	GetPlayerItemChangeLogCount(p_beginTime int64, p_endTime int64, p_playerid int64, p_itemId int64) (recordCount int, err error)

	GetPlayerMongoLogList(p_subject string, p_beginTime int64, p_endTime int64, p_playerid int64, p_pageIndex int, p_pageSize int) (result interface{}, err error)
	GetPlayerMongoLogCount(p_subject string, p_beginTime int64, p_endTime int64, p_playerid int64) (recordCount int, err error)
}

type playerLogService struct {
	mongoService mongo.MongoService
	mongoConfig  *mongo.MongoConfig
}

func (m *playerLogService) GetPlayerItemChangeLogList(p_beginTime int64, p_endTime int64, p_playerid int64, p_itemId int64, p_pageIndex int, p_pageSize int) (result interface{}, err error) {
	rst := mglog.GetLogMsgList(playerItemChangeLogType)
	if rst == nil {
		return
	}
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C(playerItemChangeLogType)
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
	// if p_platformid > 0 {
	// 	condition["platform"] = p_platformid
	// }
	condition["servertype"] = 0

	// if p_serverid > 0 {
	// 	condition["serverid"] = p_serverid
	// }
	if p_playerid > 0 {
		condition["playerid"] = p_playerid
	}
	if p_itemId > 0 {
		condition["changeditemid"] = p_itemId
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

func (m *playerLogService) GetPlayerItemChangeLogCount(p_beginTime int64, p_endTime int64, p_playerid int64, p_itemId int64) (recordCount int, err error) {
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C(playerItemChangeLogType)
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
	// if p_platformid > 0 {
	// 	condition["platform"] = p_platformid
	// }

	condition["servertype"] = 0

	// if p_serverid > 0 {
	// 	condition["serverid"] = p_serverid
	// }
	if p_playerid > 0 {
		condition["playerid"] = p_playerid
	}
	if p_itemId > 0 {
		condition["changeditemid"] = p_itemId
	}

	recordCount, err = c.Find(condition).Count()
	if err != nil {
		return
	}
	return
}

func (m *playerLogService) GetPlayerMongoLogList(p_subject string, p_beginTime int64, p_endTime int64, p_playerid int64, p_pageIndex int, p_pageSize int) (result interface{}, err error) {
	rst := mglog.GetLogMsgList(p_subject)
	if rst == nil {
		return
	}
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C(p_subject)
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
	// if p_platformid > 0 {
	// 	condition["platform"] = p_platformid
	// }
	condition["servertype"] = 0

	// if p_serverid > 0 {
	// 	condition["serverid"] = p_serverid
	// }
	if p_playerid > 0 {
		condition["playerid"] = p_playerid
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

func (m *playerLogService) GetPlayerMongoLogCount(p_subject string, p_beginTime int64, p_endTime int64, p_playerid int64) (recordCount int, err error) {
	c := m.mongoService.Session().DB(m.mongoConfig.Database).C(p_subject)
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

	condition["servertype"] = 0

	if p_playerid > 0 {
		condition["playerid"] = p_playerid
	}

	//限定1000
	recordCount, err = c.Find(condition).Count()
	if err != nil {
		return
	}
	return
}

func NewPlayerMongoLogService(p_service mongo.MongoService, p_config *mongo.MongoConfig) IPlayerMongoLogService {
	rst := &playerLogService{
		mongoService: p_service,
		mongoConfig:  p_config,
	}
	return rst
}

const (
	playMongoLogServiceKey = contextKey("playerMongoLogService")
)

func WithPlayerMongoLogService(ctx context.Context, ls IPlayerMongoLogService) context.Context {
	return context.WithValue(ctx, playMongoLogServiceKey, ls)
}

func PlayerMongoLogServiceInContext(ctx context.Context) IPlayerMongoLogService {
	us, ok := ctx.Value(playMongoLogServiceKey).(IPlayerMongoLogService)
	if !ok {
		return nil
	}
	return us
}

func SetupPlayerMongoLogServiceHandler(ls IPlayerMongoLogService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithPlayerMongoLogService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
