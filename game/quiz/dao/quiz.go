package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	quizentity "fgame/fgame/game/quiz/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "quiz"
)

type QuizDao interface {
	//获取仙尊答题信息
	GetQuizEntity(serverId int32) (quizEntity *quizentity.QuizEntity, err error)
}

type quizDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *quizDao) GetQuizEntity(serverId int32) (quizEntity *quizentity.QuizEntity, err error) {
	quizEntity = &quizentity.QuizEntity{}
	err = dao.ds.DB().First(quizEntity, "serverId=? AND deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

var (
	once sync.Once
	dao  *quizDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &quizDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetQuizDao() QuizDao {
	return dao
}
