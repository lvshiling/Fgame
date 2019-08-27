package rank

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/rank/dao"
	rankobj "fgame/fgame/game/rank/obj"
	ranktypes "fgame/fgame/game/rank/types"
	"sync"

	log "github.com/Sirupsen/logrus"
)

//排行榜接口处理
type RankService interface {
	Heartbeat()
	//启动排行
	Star() (err error)
	//获取排行map
	GetRankMap() *rankobj.RankMap
}

type rankService struct {
	rankMap *rankobj.RankMap
	//读写锁
	rwm sync.RWMutex
}

//初始化
func (rs *rankService) init(ds coredb.DBService, redisS coreredis.RedisService) (err error) {
	err = dao.Init(ds, redisS)
	if err != nil {
		return
	}

	rankTypeMap := ranktypes.RankClassTypeArea.GetRankTypeMap()
	for rankType, _ := range rankTypeMap {
		//TODO:zrc 临时处理
		if rankType != ranktypes.RankTypeForce && rankType != ranktypes.RankTypeGang {
			continue
		}

		config := ranktypes.NewAreaDefaultConfig()
		if rs.rankMap == nil {
			rs.rankMap = rankobj.NewRankMap(config.RefreshTime)
		}

		rs.rankMap.RegisterRank(rankType, config)
	}

	err = rs.rankMap.Init()
	if err != nil {
		return
	}

	return
}

//启动排行
func (rs *rankService) Star() (err error) {
	err = rs.rankMap.Star()
	if err != nil {
		return
	}
	return
}

//获取排行map
func (rs *rankService) GetRankMap() *rankobj.RankMap {
	return rs.rankMap
}

//心跳
func (rs *rankService) Heartbeat() {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()
	err := rs.rankMap.UpdateRank()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("rank:整点更新排行榜列表,错误")
		return
	}
}

//仅gm使用
func GMRankUpdate() {
	cs.rwm.Lock()
	defer cs.rwm.Unlock()
	rankClassType := ranktypes.RankClassTypeArea
	rankTypeMap := rankClassType.GetRankTypeMap()
	rankDataMap := cs.rankMap.GetRankTypeDataMap()
	for rankType, _ := range rankTypeMap {
		rankData, exist := rankDataMap[rankType]
		if !exist {
			continue
		}
		rankData.ResetRankTime()
	}
}

var (
	once sync.Once
	cs   *rankService
)

func Init(dbService coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		cs = &rankService{}
		err = cs.init(dbService, rs)
	})
	return err
}

func GetRankService() RankService {
	return cs
}
