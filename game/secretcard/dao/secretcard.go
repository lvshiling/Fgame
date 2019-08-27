package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	secretcardentity "fgame/fgame/game/secretcard/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "secretcard"
)

type SecretCardDao interface {
	//查询玩家天机牌信息
	GetSecretCardEntity(playerId int64) (*secretcardentity.PlayerSecretCardEntity, error)
}

type secretCardDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

//查询玩家天机牌信息
func (scd *secretCardDao) GetSecretCardEntity(playerId int64) (cardEntity *secretcardentity.PlayerSecretCardEntity, err error) {
	cardEntity = &secretcardentity.PlayerSecretCardEntity{}
	err = dao.ds.DB().First(cardEntity, "playerId=?", playerId).Error
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
	dao  *secretCardDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &secretCardDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetSecretCardDao() SecretCardDao {
	return dao
}
