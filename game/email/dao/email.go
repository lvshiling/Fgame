package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	emailentity "fgame/fgame/game/email/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "email"
)

type EmailDao interface {
	//玩家邮件列表
	GetEmails(playerId int64) ([]*emailentity.PlayerEmailEntity, error)
}

type emailDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *emailDao) GetEmails(playerId int64) (emails []*emailentity.PlayerEmailEntity, err error) {
	err = dao.ds.DB().Order("createTime ASC").Find(&emails, "playerId = ? AND deleteTime = 0 ", playerId).Error
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
	dao  *emailDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) error {
	once.Do(func() {
		dao = &emailDao{
			ds: ds,
			rs: rs,
		}
	})

	return nil
}

func GetEmailDao() EmailDao {
	return dao
}
