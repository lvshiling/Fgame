package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	questentity "fgame/fgame/game/quest/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "quest"
)

type QuestDao interface {
	//玩家任务列表
	GetQuestList(playerId int64) ([]*questentity.PlayerQuestEntity, error)
	//玩家屠魔次数
	GetTuMoQuestNum(playerId int64) (*questentity.PlayerTuMoEntity, error)
	//玩家日环次数
	//GetDailyQuestNum(playerId int64) (*questentity.PlayerDailyEntity, error)
	GetDailyQuestNum(playerId int64) ([]*questentity.PlayerDailyEntity, error)
	//玩家活跃度跨5点
	GetLivenessCrossFive(playerId int64) (*questentity.PlayerLivenessCrossFiveEntity, error)
	//玩家获取开服目标
	GetKaiFuMuBiaoList(playerId int64) ([]*questentity.PlayerKaiFuMuBiaoEntity, error)
	//玩家任务模块跨天
	GetQuestCrossDay(playerId int64) (*questentity.PlayerQuestCrossDayEntity, error)
	//玩家获取奇遇
	GetQiYuList(playerId int64) ([]*questentity.PlayerQiYuEntity, error)
}

type questDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *questDao) GetQuestList(playerId int64) (questList []*questentity.PlayerQuestEntity, err error) {
	err = dao.ds.DB().Find(&questList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *questDao) GetTuMoQuestNum(playerId int64) (tuMoEntity *questentity.PlayerTuMoEntity, err error) {
	tuMoEntity = &questentity.PlayerTuMoEntity{}
	err = dao.ds.DB().First(tuMoEntity, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

// func (dao *questDao) GetDailyQuestNum(playerId int64) (dailyEntity *questentity.PlayerDailyEntity, err error) {
// 	dailyEntity = &questentity.PlayerDailyEntity{}
// 	err = dao.ds.DB().First(dailyEntity, "playerId=? ", playerId).Error
// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			return
// 		}
// 		return nil, nil
// 	}
// 	return
// }

func (dao *questDao) GetDailyQuestNum(playerId int64) (dailyEntityList []*questentity.PlayerDailyEntity, err error) {
	err = dao.ds.DB().Find(&dailyEntityList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *questDao) GetLivenessCrossFive(playerId int64) (crossFiveEntity *questentity.PlayerLivenessCrossFiveEntity, err error) {
	crossFiveEntity = &questentity.PlayerLivenessCrossFiveEntity{}
	err = dao.ds.DB().First(crossFiveEntity, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *questDao) GetKaiFuMuBiaoList(playerId int64) (kaiFuMuBiaoList []*questentity.PlayerKaiFuMuBiaoEntity, err error) {
	err = dao.ds.DB().Find(&kaiFuMuBiaoList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *questDao) GetQiYuList(playerId int64) (qiyuList []*questentity.PlayerQiYuEntity, err error) {
	err = dao.ds.DB().Find(&qiyuList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *questDao) GetQuestCrossDay(playerId int64) (crossDayEntity *questentity.PlayerQuestCrossDayEntity, err error) {
	crossDayEntity = &questentity.PlayerQuestCrossDayEntity{}
	err = dao.ds.DB().First(crossDayEntity, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

var (
	once sync.Once
	dao  *questDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &questDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetQuestDao() QuestDao {
	return dao
}
