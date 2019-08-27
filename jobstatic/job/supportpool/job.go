package supportpool

import (
	"fgame/fgame/core/db"
	"fgame/fgame/core/redis"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	poolmodel "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/model"

	"fgame/fgame/gm/gamegm/gm/lockerkey"

	log "github.com/Sirupsen/logrus"
)

const (
	beforeOrderTime = int64(time.Minute/time.Millisecond) * 10 //每次统计的订单的时间为10分钟前
)

type SupportPoolJob struct {
	ds     db.DBService
	rs     redis.RedisService
	centDs db.DBService
}

/*******接口开始*************/

func (m *SupportPoolJob) GetId() string {
	return "supportPoolJob"
}

func (m *SupportPoolJob) Run() error {
	log.Debug("作业SupportPoolJob开始运行...")
	poolList := make([]*poolmodel.ServerSupportPool, 0)
	exdb := m.ds.DB().Where("deleteTime=0").Find(&poolList)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}

	now := timeutils.TimeToMillisecond(time.Now()) - beforeOrderTime

	for _, value := range poolList {
		err := m.addOrderServerPool(value.Id, now)
		if err != nil {
			return err
		}
	}
	return nil
}

var (
	orderSql = `SELECT SUM(A.gold) AS orderGold FROM t_order A 
	INNER JOIN t_server B
	ON A.serverId = B.serverId
	WHERE B.id=? and A.sdkType = ? and A.updateTime >= ? and A.updateTime < ? and A.status IN (1,2)`
)

func (m *SupportPoolJob) addOrderServerPool(poolId int64, lastTime int64) error {
	lockerKey := lockerkey.GetServerSupportRedisLocker()
	conn := m.rs.Pool().Get()
	lockRst, err := redis.LockDefault(conn, lockerKey)
	if err != nil {
		return err
	}
	if !lockRst {
		return fmt.Errorf("lock fail")
	}
	defer func() {
		redis.Unlock(conn, lockerKey)
		conn.Close()
	}()
	poolInfo := &poolmodel.ServerSupportPool{}
	exdb := m.ds.DB().Where("id=?", poolId).First(poolInfo)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	if poolInfo.Id < 1 {
		return nil
	}
	if poolInfo.CurOrderTime == 0 { //如果是旧数据，未叠加过
		poolInfo.CurOrderTime = lastTime
		exdb = m.ds.DB().Save(poolInfo)
		if exdb.Error != nil {
			return exdb.Error
		}
		return nil
	}

	serverKeyId := poolInfo.ServerId
	beginTime := poolInfo.CurOrderTime
	endTime := lastTime
	sdkType := poolInfo.SdkType
	orderCount := &OrderGoldCount{}
	odExdb := m.centDs.DB().Raw(orderSql, serverKeyId, sdkType, beginTime, endTime).Scan(orderCount)
	if odExdb.Error != nil && odExdb.Error != gorm.ErrRecordNotFound {
		return odExdb.Error
	}
	poolInfo.CurOrderTime = endTime
	addOrderCount := orderCount.OrderGold * poolInfo.OrderGoldPer / 100
	poolInfo.CurGold = poolInfo.CurGold + int(addOrderCount)
	poolInfo.OrderGold = poolInfo.OrderGold + addOrderCount

	exdb = m.ds.DB().Save(poolInfo)
	if exdb.Error != nil {
		return exdb.Error
	}

	return nil
}

func (m *SupportPoolJob) GetTickSecond() int64 {
	return 10 * 60
	// return 10 * 60
}

/*******接口结束*************/

func NewSupportPoolJob(ds db.DBService, rs redis.RedisService, centerDs db.DBService) *SupportPoolJob {
	rst := &SupportPoolJob{
		ds:     ds,
		rs:     rs,
		centDs: centerDs,
	}
	return rst
}
