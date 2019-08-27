package monitor

import (
	"context"
	gmdb "fgame/fgame/gm/gamegm/db"
	centermodel "fgame/fgame/gm/gamegm/monitor/model"
	"net/http"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type ICenterServer interface {
	SyncServer() error
	GetCenterServerDBId(p_platid int32, p_serverId int32) int64
	GetServerId(p_id int64) (int, error)
	GetCenterServerDbInfo(p_id int64) (*centermodel.CenterServer, error)
}

type centerServer struct {
	serverIdMap map[int32]map[int32]int64 //平台服务及数据库服务对应关系 key:中心平台id， key：中心服务serverid， value：中心服务器数据库主键ID
	db          gmdb.DBService
	rwm         sync.RWMutex
}

func (m *centerServer) SyncServer() error {
	log.Debug("同步中心服务器数据")
	rst := make([]*centermodel.CenterServer, 0)
	exdb := m.db.DB().Where("serverType=0 and deleteTime=0").Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}

	serverMap := make(map[int32]map[int32]int64)
	for _, value := range rst {
		platId := int32(value.Platform)
		serverId := int32(value.ServerId)
		id := value.Id
		if _, ok := serverMap[platId]; !ok {
			serverMap[platId] = make(map[int32]int64)
		}
		serverMap[platId][serverId] = id
	}

	m.setServerIdMap(serverMap)
	return nil
}

func (m *centerServer) GetServerId(p_id int64) (int, error) {
	model := &centermodel.CenterServer{}
	exdb := m.db.DB().Where("id = ?", p_id).First(model)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return model.ServerId, nil
}

func (m *centerServer) GetCenterServerDbInfo(p_id int64) (*centermodel.CenterServer, error) {
	model := &centermodel.CenterServer{}
	exdb := m.db.DB().Where("id = ?", p_id).First(model)
	if exdb.Error != nil {
		return nil, exdb.Error
	}
	return model, nil
}

func (m *centerServer) GetCenterServerDBId(p_platid int32, p_serverId int32) int64 {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	if value, ok := m.serverIdMap[p_platid]; ok {
		if dbid, idok := value[p_serverId]; idok {
			return dbid
		}
	}
	return 0
}

func (m *centerServer) setServerIdMap(p_map map[int32]map[int32]int64) {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	m.serverIdMap = p_map
}

func NewCenterServer(p_db gmdb.DBService) ICenterServer {
	rst := &centerServer{
		db: p_db,
	}
	return rst
}

// type contextKey string

const (
	centerServerServiceKey = contextKey("CenterServerService")
)

func WithCenterServerService(ctx context.Context, ls ICenterServer) context.Context {
	return context.WithValue(ctx, centerServerServiceKey, ls)
}

func CenterServerServiceInContext(ctx context.Context) ICenterServer {
	us, ok := ctx.Value(centerServerServiceKey).(ICenterServer)
	if !ok {
		return nil
	}
	return us
}

func SetupCenterServerServiceHandler(ls ICenterServer) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithCenterServerService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
