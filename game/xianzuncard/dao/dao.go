package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/xianzuncard/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "xianzuncard"
)

type XianZunCard interface {
	// 查询玩家仙尊特权卡信息
	GetPlayerXianZunCardEntityList(playerId int64) ([]*entity.PlayerXianZunCardEntity, error)
}

type xianZunCard struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *xianZunCard) GetPlayerXianZunCardEntityList(playerId int64) (entityList []*entity.PlayerXianZunCardEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *xianZunCard
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &xianZunCard{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetXianZunCard() XianZunCard {
	return dao
}
