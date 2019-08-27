package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DbConfig struct {
	Debug       bool   `json:"debug"`
	Dialect     string `json:"dialect"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	DBName      string `json:"dbName"`
	ParseTime   bool   `json:"parseTime"`
	Charset     string `json:"charset"`
	MaxIdle     int    `json:"maxIdle"`
	MaxActive   int    `json:"maxActive"`
	MaxLifeTime int64  `json:"maxLifeTime"`
}

type DBService interface {
	DB() *gorm.DB
}

type dbService struct {
	db     *gorm.DB
	config *DbConfig
}

func (rs *dbService) DB() *gorm.DB {
	return rs.db
}

func NewDBService(config *DbConfig) (DBService, error) {
	dbArgs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%v", config.User, config.Password, config.Host, config.DBName, config.Charset, config.ParseTime)
	s, err := gorm.Open(config.Dialect, dbArgs)
	if err != nil {
		return nil, err
	}
	s.DB().SetMaxIdleConns(config.MaxIdle)
	s.DB().SetMaxOpenConns(config.MaxActive)
	s.DB().SetConnMaxLifetime(time.Duration(int64(time.Second) * config.MaxLifeTime))
	s.LogMode(config.Debug)
	s.Set("gorm:table_options", "ENGINE=InnoDB")
	if s.Error != nil {
		return nil, s.Error
	}
	return &dbService{
		db:     s,
		config: config,
	}, nil
}
