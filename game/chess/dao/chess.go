package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	chessentity "fgame/fgame/game/chess/entity"
	"fgame/fgame/game/global"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "chess"
)

type ChessDao interface {
	//获取玩家棋局信息
	GetPlayerChessEntityList(playerId int64) (entityList []*chessentity.PlayerChessEntity, err error)
	//获取棋局日志
	GetChessLogEntityList() (logEntityList []*chessentity.ChessLogEntity, err error)
}

type chessDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *chessDao) GetPlayerChessEntityList(playerId int64) (entityList []*chessentity.PlayerChessEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chessDao) GetChessLogEntityList() (logEntityList []*chessentity.ChessLogEntity, err error) {
	err = dao.ds.DB().Order("updateTime ASC").Find(&logEntityList, "serverId = ? AND deleteTime=0", global.GetGame().GetServerIndex()).Error
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
	dao  *chessDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &chessDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetChessDao() ChessDao {
	return dao
}
