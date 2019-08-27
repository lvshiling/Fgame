package db

import (
	"fmt"
	"sync"
)

type GameDbLink int32

type IDbManager interface {
	SetDb(p_serverid GameDbLink, p_db DBService)
	GetDB(p_serverid GameDbLink) DBService
	AddDbConfig(p_serverid GameDbLink, p_config *DbConfig)
	InitCenterDB(centerDB DBService)
	GetFinnalServerId(serverId int32) int32
}

type DBManager struct {
	_dbMap       map[GameDbLink]DBService
	_dbConfigMap map[GameDbLink]*DbConfig
	_rwm         sync.RWMutex
	_db          DBService
}

func (m *DBManager) InitCenterDB(centerDB DBService) {
	m._db = centerDB
}

func (m *DBManager) SetDb(p_serverid GameDbLink, p_db DBService) {
	m._rwm.Lock()
	defer m._rwm.Unlock()
	m._dbMap[p_serverid] = p_db
}

func (m *DBManager) GetDB(p_serverid GameDbLink) DBService {
	m._rwm.Lock()
	defer m._rwm.Unlock()
	if value, ok := m._dbMap[p_serverid]; ok {
		return value
	}
	if configvalue, configok := m._dbConfigMap[p_serverid]; configok {
		gamedb, gameerr := NewDBService(configvalue)
		if gameerr != nil {
			return nil
		}
		m._dbMap[p_serverid] = gamedb
		return gamedb
	}
	return nil
}

func (m *DBManager) AddDbConfig(p_serverid GameDbLink, p_config *DbConfig) {
	m._rwm.Lock()
	defer m._rwm.Unlock()
	m._dbConfigMap[p_serverid] = p_config
}

func (m *DBManager) GetFinnalServerId(serverId int32) int32 {
	link := GameDbLink(serverId)
	finnalLink := m.getFinnalServerId(link)
	return int32(finnalLink)
}

const (
	finnalServerSql = `SELECT
	IFNULL(C.id,A.id) AS id
FROM
	t_server A
	LEFT JOIN t_merge_record B
	ON A.platform = B.platform AND A.serverId = B.fromServerId AND B.deleteTime = 0
	LEFT JOIN t_server C
	ON B.platform = C.platform AND B.finalServerId = C.serverId AND C.deleteTime = 0 AND C.serverType = 0
WHERE
	A.id=?`
)

type serverInfo struct {
	FinnalId int `gorm:"column:id"`
}

func (m *DBManager) getFinnalServerId(link GameDbLink) GameDbLink {
	info := &serverInfo{}
	errdb := m._db.DB().Raw(finnalServerSql, int(link)).Scan(info)
	if errdb.Error != nil {
		return GameDbLink(0)
	}
	return GameDbLink(info.FinnalId)
}

func NewDBManager() IDbManager {
	rst := &DBManager{}
	rst._dbMap = make(map[GameDbLink]DBService)
	rst._dbConfigMap = make(map[GameDbLink]*DbConfig)
	return rst
}

var (
	_dbManager = NewDBManager()
)

func GetDb(p_serviceid GameDbLink) DBService {
	finnalLink := _dbManager.GetFinnalServerId(int32(p_serviceid))
	fmt.Println("新服务器ID:", finnalLink)
	return _dbManager.GetDB(GameDbLink(finnalLink))
}

func SetDb(p_serverid GameDbLink, p_db DBService) {
	_dbManager.SetDb(p_serverid, p_db)
}

func AddDbConfig(p_serverid GameDbLink, p_config *DbConfig) {
	_dbManager.AddDbConfig(p_serverid, p_config)
}

func GetFinnalServerId(serverId int32) int32 {
	return _dbManager.GetFinnalServerId(serverId)
}

func RegisterDB(p_serverid GameDbLink, p_host string, p_port string, p_dbName string, p_userName string, p_password string) error {
	dbhost := fmt.Sprintf("%s:%s", p_host, p_port)
	dbconfig := &DbConfig{
		Debug:       true,
		Dialect:     "mysql",
		User:        p_userName,
		Password:    p_password,
		Host:        dbhost,
		DBName:      p_dbName,
		ParseTime:   true,
		Charset:     "utf8mb4",
		MaxIdle:     50,
		MaxActive:   100,
		MaxLifeTime: 240,
	}
	_dbManager.AddDbConfig(p_serverid, dbconfig)
	gamedb, gameerr := NewDBService(dbconfig)
	if gameerr != nil {
		return gameerr
	}
	SetDb(p_serverid, gamedb)
	return nil
}

func InitDBManager(centerDB DBService) {
	_dbManager.InitCenterDB(centerDB)
}
