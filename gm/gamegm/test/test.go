package test

import (
	gmdb "fgame/fgame/gm/gamegm/db"
)

func NewQiPaiDB() gmdb.DBService {
	conf := &gmdb.DbConfig{
		Debug:       true,
		Dialect:     "mysql",
		User:        "root",
		Password:    "zrc881002",
		Host:        "127.0.0.1:3316",
		DBName:      "gamegm",
		ParseTime:   true,
		Charset:     "utf8",
		MaxIdle:     100,
		MaxActive:   50,
		MaxLifeTime: 240,
	}
	rst, _ := gmdb.NewDBService(conf)
	return rst
}
