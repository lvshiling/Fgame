package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	feedbackfeeentity "fgame/fgame/game/feedbackfee/entity"
	feedbackfeetypes "fgame/fgame/game/feedbackfee/types"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "feedbackfee"
)

type FeedbackFeeDao interface {
	//逆付费信息
	GetFeedbackFeeEntity(playerId int64) (feedbackfeeEntity *feedbackfeeentity.PlayerFeedbackFeeEntity, err error)
	//获取正在进行中的记录 特殊处理获取状态小于status
	GetFeedbackRecordList(playerId int64, status int32) (recordList []*feedbackfeeentity.PlayerFeedbackRecordEntity, err error)
	//获取兑换记录
	GetUnfinishFeedbackExchangeList(serverId int32) (exchangeList []*feedbackfeeentity.FeedbackExchangeEntity, err error)
}

type feedbackfeeDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *feedbackfeeDao) GetFeedbackFeeEntity(playerId int64) (feedbackfeeEntity *feedbackfeeentity.PlayerFeedbackFeeEntity, err error) {
	feedbackfeeEntity = &feedbackfeeentity.PlayerFeedbackFeeEntity{}
	err = dao.ds.DB().First(feedbackfeeEntity, "playerId=? and deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *feedbackfeeDao) GetFeedbackRecordList(playerId int64, status int32) (recordList []*feedbackfeeentity.PlayerFeedbackRecordEntity, err error) {
	recordList = make([]*feedbackfeeentity.PlayerFeedbackRecordEntity, 0, 1)
	err = dao.ds.DB().Find(&recordList, "playerId=? and status<?  and deleteTime=0", playerId, status).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *feedbackfeeDao) GetUnfinishFeedbackExchangeList(serverId int32) (exchangeList []*feedbackfeeentity.FeedbackExchangeEntity, err error) {
	exchangeList = make([]*feedbackfeeentity.FeedbackExchangeEntity, 0, 8)
	err = dao.ds.DB().Find(&exchangeList, "serverId=? and status!=?  and deleteTime=0", serverId, int32(feedbackfeetypes.FeedbackExchangeStatusNotify)).Error
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
	dao  *feedbackfeeDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &feedbackfeeDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetFeedbackFeeDao() FeedbackFeeDao {
	return dao
}
