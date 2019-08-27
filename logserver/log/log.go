package log

import (
	coremongo "fgame/fgame/core/mongo"
)

type LogService interface {
	AddLog(m LogMsg) (err error)
}

type logService struct {
	mongoService coremongo.MongoService
}

func (l *logService) AddLog(m LogMsg) (err error) {
	dbName := l.mongoService.GetDatabase()
	collectionName := m.LogName()
	c := l.mongoService.Session().DB(dbName).C(collectionName)
	err = c.Insert(m)
	if err != nil {
		return
	}

	return nil
}

func newLogService(mongoService coremongo.MongoService) LogService {
	ls := &logService{
		mongoService: mongoService,
	}
	return ls
}

var (
	s LogService
)

func GetLogService() LogService {
	return s
}

func Init(mongoService coremongo.MongoService) (err error) {
	ls := newLogService(mongoService)
	s = ls
	return nil
}
