package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	registerentity "fgame/fgame/game/register/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "register"
)

type RegisterDao interface {
	//获取设置
	GetRegisterSetting(serverId int32) (*registerentity.RegisterSettingEntity, error)
}

type registerDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *registerDao) GetRegisterSetting(serverId int32) (registerSetttingEntity *registerentity.RegisterSettingEntity, err error) {
	registerSetttingEntity = &registerentity.RegisterSettingEntity{}
	err = dao.ds.DB().First(registerSetttingEntity, "serverId=?", serverId).Error
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
	dao  *registerDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &registerDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetRegisterDao() RegisterDao {
	return dao
}
