package mongo

import (
	"gopkg.in/mgo.v2"
)

type MongoConfig struct {
	Addrs     []string `json:"addrs"`
	FailFast  bool     `json:"failFast"`
	PoolLimit int      `json:"poolLimit"`
	Database  string   `json:"database"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
}

type MongoService interface {
	Session() *mgo.Session
	GetDatabase() string
}

type mongoService struct {
	session  *mgo.Session
	database string
}

func (ms *mongoService) Session() *mgo.Session {
	return ms.session
}

func (ms *mongoService) GetDatabase() string {
	return ms.database
}

func NewMongoService(mc *MongoConfig) (MongoService, error) {
	info := &mgo.DialInfo{}

	info.Addrs = mc.Addrs
	info.FailFast = mc.FailFast
	info.Username = mc.Username
	info.Password = mc.Password
	info.PoolLimit = mc.PoolLimit
	info.Database = mc.Database
	s, err := mgo.DialWithInfo(info)
	if err != nil {
		return nil, err
	}
	ms := &mongoService{}
	ms.session = s
	ms.database = mc.Database
	return ms, nil
}
