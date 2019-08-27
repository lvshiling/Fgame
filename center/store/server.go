package store

import (
	"fgame/fgame/center/types"
	coredb "fgame/fgame/core/db"

	"github.com/jinzhu/gorm"
)

type ServerEntity struct {
	Id                int32  `gorm:"primary_key;column:id"`
	ServerType        int32  `gorm:"column:serverType"`
	ServerId          int32  `gorm:"column:serverId"`
	Platform          int32  `gorm:"column:platform"`
	Name              string `gorm:"column:name"`
	StartTime         int64  `gorm:"column:startTime"`
	ServerIp          string `gorm:"column:serverIp"`
	ServerPort        string `gorm:"column:serverPort"`
	ServerRemoteIp    string `gorm:"column:serverRemoteIp"`
	ServerRemotePort  string `gorm:"column:serverRemotePort"`
	ServerDBIp        string `gorm:"column:serverDBIp"`
	ServerDBPort      string `gorm:"column:serverDBPort"`
	ServerDBName      string `gorm:"column:serverDBName"`
	ServerDBUser      string `gorm:"column:serverDBUser"`
	ServerDBPassword  string `gorm:"column:serverDBPassword"`
	ServerTag         int32  `gorm:"column:serverTag"`
	ServerStatus      int32  `gorm:"column:serverStatus"`
	ParentServerId    int32  `gorm:"column:parentServerId"`
	PreShow           int32  `gorm:"column:preShow"`
	PingTaiFuServerId int32  `gorm:"column:pingTaiFuServerId"`
	UpdateTime        int64  `gorm:"column:updateTime"`
	CreateTime        int64  `gorm:"column:createTime"`
	DeleteTime        int64  `gorm:"column:deleteTime"`
}

func (e *ServerEntity) TableName() string {
	return "t_server"
}

type MergeRecordEntity struct {
	Id            int32 `gorm:"primary_key;column:id"`
	FromServerId  int32 `gorm:"column:fromServerId"`
	ToServerId    int32 `gorm:"column:toServerId"`
	FinalServerId int32 `gorm:"column:finalServerId"`
	Platform      int32 `gorm:"column:platform"`
	MergeTime     int64 `gorm:"column:mergeTime"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (e *MergeRecordEntity) TableName() string {
	return "t_merge_record"
}

type ServerStore interface {
	//获取所有服务器
	GetAll() ([]*ServerEntity, error)
	GetAllMergeList() ([]*MergeRecordEntity, error)
	GetServerList(platform int32, serverType types.GameServerType) ([]*ServerEntity, error)
	GetMergeList(platform int32) ([]*MergeRecordEntity, error)
	Save(e *ServerEntity) (err error)
}

var (
	dbName = "server"
)

type serverStore struct {
	db coredb.DBService
}

func (s *serverStore) GetAll() (eList []*ServerEntity, err error) {
	err = s.db.DB().Find(&eList, "deleteTime=0").Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return
	}
	return
}

func (s *serverStore) GetServerList(platform int32, serverType types.GameServerType) (eList []*ServerEntity, err error) {
	err = s.db.DB().Find(&eList, "platform=? and serverType=? and deleteTime=0", platform, int32(serverType)).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return
	}
	return
}

func (s *serverStore) GetAllMergeList() (eList []*MergeRecordEntity, err error) {
	err = s.db.DB().Find(&eList, "deleteTime=0").Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return
	}
	return
}

func (s *serverStore) GetMergeList(platform int32) (eList []*MergeRecordEntity, err error) {
	err = s.db.DB().Find(&eList, "platform=? and  deleteTime=0", platform).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return
	}
	return
}

func (s *serverStore) Save(e *ServerEntity) (err error) {
	err = s.db.DB().Save(e).Error
	if err != nil {
		return
	}
	return
}

func NewServerStore(db coredb.DBService) ServerStore {
	s := &serverStore{
		db: db,
	}
	return s
}
