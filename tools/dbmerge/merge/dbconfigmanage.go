package merge

import (
	fdb "fgame/fgame/core/db"
	model "fgame/fgame/tools/dbmerge/model"
	"fmt"
	"sync"
)

type IDbConfigManage interface {
	RegisterDbConfigInfo(p_platformId int64, p_serverId int, p_info *model.DBConfigInfo)
	GetDbConfigInfo(p_platformId int64, p_serverId int) *model.DBConfigInfo
	GetDbService(p_platformId int64, p_serverId int) (fdb.DBService, error)
}

type dbConfigManage struct {
	rwm       sync.RWMutex
	configMap map[int64]map[int]*model.DBConfigInfo
	dbMap     map[int64]map[int]fdb.DBService
}

func (m *dbConfigManage) RegisterDbConfigInfo(p_platformId int64, p_serverId int, p_info *model.DBConfigInfo) {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	if m.configMap == nil {
		m.configMap = make(map[int64]map[int]*model.DBConfigInfo)
		m.dbMap = make(map[int64]map[int]fdb.DBService)
	}
	if _, ok := m.configMap[p_platformId]; !ok {
		m.configMap[p_platformId] = make(map[int]*model.DBConfigInfo)
	}
	m.configMap[p_platformId][p_serverId] = p_info
}

func (m *dbConfigManage) GetDbConfigInfo(p_platformId int64, p_serverId int) *model.DBConfigInfo {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	return m.getDbConfigInfo(p_platformId, p_serverId)
}

func (m *dbConfigManage) getDbConfigInfo(p_platformId int64, p_serverId int) *model.DBConfigInfo {
	if platvalue, platok := m.configMap[p_platformId]; platok {
		if info, serverok := platvalue[p_serverId]; serverok {
			return info
		}
	}
	return nil
}

func (m *dbConfigManage) GetDbService(p_platformId int64, p_serverId int) (fdb.DBService, error) {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	if _, ok := m.dbMap[p_platformId]; !ok {
		m.dbMap[p_platformId] = make(map[int]fdb.DBService)
	}
	if dbvalue, ok := m.dbMap[p_platformId][p_serverId]; ok {
		return dbvalue, nil
	}

	toDB := m.getDbConfigInfo(p_platformId, p_serverId)
	if toDB == nil {
		return nil, fmt.Errorf("db not exist")
	}

	dbconfig := &fdb.DbConfig{
		Debug:       false,
		Dialect:     "mysql",
		User:        toDB.UserName,
		Password:    toDB.PassWord,
		Host:        fmt.Sprintf("%s:%d", toDB.Host, toDB.Port),
		DBName:      toDB.DBName,
		ParseTime:   true,
		Charset:     "utf8mb4",
		MaxIdle:     50,
		MaxActive:   100,
		MaxLifeTime: 200,
	}

	db, err := fdb.NewDBService(dbconfig)
	if err != nil {
		return nil, err
	}
	m.dbMap[p_platformId][p_serverId] = db
	return db, nil
}

func NewDbConfigManage() IDbConfigManage {
	rst := &dbConfigManage{}
	return rst
}
