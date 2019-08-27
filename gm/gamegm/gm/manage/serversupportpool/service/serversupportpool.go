package model

import (
	"context"
	"fgame/fgame/gm/gamegm/common"
	"fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	gmerr "fgame/fgame/gm/gamegm/error"
	smodel "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	"fgame/fgame/core/redis"
	platformmodel "fgame/fgame/gm/gamegm/gm/platform/model"

	centerservermodel "fgame/fgame/gm/gamegm/gm/center/model"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
)

const (
	beforeOrderTime = int64(time.Minute/time.Millisecond) * 5 //每次统计的订单的时间为10分钟前
)

type IServerSupportPool interface {
	AddServerSupportPool(p_serverId int, p_gold int, p_sdkType int, p_centerPlatformId int64, p_percent int32) error
	UpdateServerSupportPool(p_id int, p_gold int, percent int32) error
	DeleteServerSupportPool(p_id int) error

	GetServerSupportPoolList(p_pageindex int, p_serverId int, p_centerPlatformId int, p_platformList []int64) ([]*smodel.ServerSupportPool, error)
	GetServerSupportPoolCount(p_serverId int, p_centerPlatformId int, p_platformList []int64) (int, error)
	GetServerSupportPoolInfo(p_serverId int32) (*smodel.ServerSupportPool, error)
	ReduceServerGold(p_serverId int32, p_gold int) error

	AddOrderPoolAmount() error
	Heartbeat()

	AddPlatformSupportPoolSet(p_centerPlatformId int64, p_supportGold int32, p_supportRate int32) error
	UpdatePlatformSupportPoolSet(p_id int32, p_supportGold int32, p_supportRate int32) error
	DeletePlatformSupportPoolSet(p_id int32) error
	GetPlatformSupportPoolList(p_centerPlatformId int64, p_index int32) ([]*smodel.PlatformSupportPoolSetInfo, error)
	GetPlatformSupportPoolCount(p_centerPlatformId int64) (int32, error)
	GetPlatformSupportPoolSet(p_centerPlatformId int64) (*smodel.PlatformSupportPoolSetInfo, error)

	FillAllServerPoolSet() error
}

type serverSupportPool struct {
	db     gmdb.DBService
	centDs gmdb.DBService
	rs     redis.RedisService
	rwm    sync.RWMutex
}

func (m *serverSupportPool) existsServerId(p_serverId int) (bool, error) {
	rst := 0
	errdb := m.db.DB().Table("t_server_support_pool").Where("serverId = ? and deleteTime=0", p_serverId).Count(&rst)
	if errdb.Error != nil {
		return false, errdb.Error
	}
	if rst > 0 {
		return true, nil
	}
	return false, nil
}

func (m *serverSupportPool) AddServerSupportPool(p_serverId int, p_gold int, p_sdkType int, p_centerPlatformId int64, p_percent int32) error {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	exflag, err := m.existsServerId(p_serverId)
	if err != nil {
		return err
	}
	if exflag {
		return gmerr.GetError(gmerr.ErrorCodeServerSupportPoolExists)
	}

	now := timeutils.TimeToMillisecond(time.Now())
	info := &smodel.ServerSupportPool{
		ServerId:         p_serverId,
		BeginGold:        p_gold,
		CurGold:          p_gold,
		CreateTime:       now,
		SdkType:          p_sdkType,
		CenterPlatformId: p_centerPlatformId,
		OrderGoldPer:     p_percent,
		OrderGold:        0,
		BeginOrderTime:   now,
		CurOrderTime:     now,
	}
	errdb := m.db.DB().Save(info)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *serverSupportPool) UpdateServerSupportPool(p_id int, p_gold int, percent int32) error {
	//加锁
	m.rwm.Lock()
	defer m.rwm.Unlock()

	dberr := m.db.DB().Table("t_server_support_pool").Where("id=?", p_id).Updates(map[string]interface{}{"curGold": p_gold, "orderGoldPer": percent})
	if dberr.Error != nil {
		return dberr.Error
	}
	return nil
}

func (m *serverSupportPool) DeleteServerSupportPool(p_id int) error {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	now := timeutils.TimeToMillisecond(time.Now())
	dberr := m.db.DB().Table("t_server_support_pool").Where("id=?", p_id).Update("deleteTime", now)
	if dberr.Error != nil {
		return dberr.Error
	}
	return nil
}

func (m *serverSupportPool) GetServerSupportPoolList(p_pageindex int, p_serverId int, p_centerPlatformId int, p_platformList []int64) ([]*smodel.ServerSupportPool, error) {
	rst := make([]*smodel.ServerSupportPool, 0)
	offset := (p_pageindex - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	where := "deleteTime=0"
	if p_serverId > 0 {
		where += fmt.Sprintf(" and serverId=%d", p_serverId)
	}
	if len(p_platformList) > 0 {
		where += fmt.Sprintf(" and centerPlatformId IN (%s)", common.CombinInt64Array(p_platformList))
	}
	if p_centerPlatformId > 0 {
		where += fmt.Sprintf(" and centerPlatformId = %d", p_centerPlatformId)
	}
	dberr := m.db.DB().Where(where).Offset(offset).Limit(limit).Find(&rst)
	if dberr.Error != nil {
		return nil, dberr.Error
	}
	return rst, nil
}

func (m *serverSupportPool) GetServerSupportPoolCount(p_serverId int, p_centerPlatformId int, p_platformList []int64) (int, error) {
	rst := 0
	where := "deleteTime=0"
	if p_serverId > 0 {
		where += fmt.Sprintf(" and serverId=%d", p_serverId)
	}
	if len(p_platformList) > 0 {
		where += fmt.Sprintf(" and centerPlatformId IN (%s)", common.CombinInt64Array(p_platformList))
	}
	if p_centerPlatformId > 0 {
		where += fmt.Sprintf(" and centerPlatformId = %d", p_centerPlatformId)
	}
	dberr := m.db.DB().Table("t_server_support_pool").Where(where).Count(&rst)
	if dberr.Error != nil {
		return 0, dberr.Error
	}
	return rst, nil
}

func (m *serverSupportPool) GetServerSupportPoolInfo(p_serverId int32) (*smodel.ServerSupportPool, error) {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	rst := &smodel.ServerSupportPool{}
	dberr := m.db.DB().Where("deleteTime=0 and serverId=?", p_serverId).First(rst)
	if dberr.Error != nil {
		if dberr.Error == gorm.ErrRecordNotFound { //不存在扶持池
			err := m.fillServerSupportPool(p_serverId)
			if err != nil {
				return nil, err
			}
			dberr = m.db.DB().Where("deleteTime=0 and serverId=?", p_serverId).First(rst)
			if dberr.Error != nil && dberr.Error != gorm.ErrRecordNotFound {
				return nil, dberr.Error
			}
			return rst, nil
		}
		return nil, dberr.Error
	}
	return rst, nil
}

func (m *serverSupportPool) ReduceServerGold(p_serverId int32, p_gold int) error {
	//加锁
	m.rwm.Lock()
	defer m.rwm.Unlock()

	rst := &smodel.ServerSupportPool{}
	dberr := m.db.DB().Where("deleteTime=0 and serverId=?", p_serverId).First(rst)
	if dberr.Error != nil {
		return dberr.Error
	}
	rst.DelGold = rst.DelGold + p_gold
	rst.CurGold = rst.CurGold - p_gold
	dberr = m.db.DB().Save(rst)
	if dberr.Error != nil {
		return dberr.Error
	}
	return nil
}

var (
	orderSql = `SELECT SUM(A.gold) AS orderGold FROM t_order A 
	INNER JOIN t_server B
	ON A.serverId = B.serverId
	WHERE B.id=? and A.sdkType IN (?) and A.updateTime >= ? and A.updateTime < ? and A.status IN (2)`
)

func (m *serverSupportPool) AddOrderPoolAmount() error {
	poolList := make([]*smodel.ServerSupportPool, 0)
	exdb := m.db.DB().Where("deleteTime=0").Find(&poolList)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	now := timeutils.TimeToMillisecond(time.Now()) - beforeOrderTime
	orderMap := make(map[int64]*smodel.OrderPoolAddInfo)
	allSkdMap := make(map[int64][]int)
	allPlatformArray := make([]*platformmodel.PlatformInfo, 0)
	exdb = m.db.DB().Where("deleteTime =0").Find(&allPlatformArray)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	for _, platformInfo := range allPlatformArray {
		allSkdMap[platformInfo.CenterPlatformID] = append(allSkdMap[platformInfo.CenterPlatformID], platformInfo.SdkType)
	}

	for _, value := range poolList {
		if value.CurOrderTime == 0 {
			orderMap[value.Id] = &smodel.OrderPoolAddInfo{}
			continue
		}
		sdkList, exists := allSkdMap[value.CenterPlatformId]
		if !exists {
			orderMap[value.Id] = &smodel.OrderPoolAddInfo{}
			continue
		}
		orderCountInfo, err := m.getOrderAmount(value.ServerId, sdkList, value.CurOrderTime, now)
		if err != nil {
			return err
		}
		orderMap[value.Id] = orderCountInfo
	}

	m.rwm.Lock()
	defer m.rwm.Unlock()
	err := m.addOrderServerPoolByOrderMap(orderMap, now)
	if err != nil {
		return err
	}
	return nil
}

func (m *serverSupportPool) addOrderServerPoolByOrderMap(orderMap map[int64]*smodel.OrderPoolAddInfo, lastTime int64) error {
	for poolId, orderCount := range orderMap {
		poolInfo := &smodel.ServerSupportPool{}
		exdb := m.db.DB().Where("id=?", poolId).First(poolInfo)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			return exdb.Error
		}
		if poolInfo == nil || poolInfo.Id < 1 {
			continue
		}
		if poolInfo.CurOrderTime == 0 { //如果是旧数据，未叠加过
			poolInfo.CurOrderTime = lastTime
			exdb = m.db.DB().Save(poolInfo)
			if exdb.Error != nil {
				return exdb.Error
			}
			continue
		}
		poolInfo.CurOrderTime = lastTime
		addOrderCount := orderCount.OrderGold * poolInfo.OrderGoldPer / 100
		poolInfo.CurGold = poolInfo.CurGold + int(addOrderCount)
		poolInfo.OrderGold = poolInfo.OrderGold + addOrderCount

		exdb = m.db.DB().Save(poolInfo)
		if exdb.Error != nil {
			return exdb.Error
		}
	}
	return nil
}

func (m *serverSupportPool) getOrderAmount(serverKeyId int, sdkType []int, beginTime int64, endTime int64) (*smodel.OrderPoolAddInfo, error) {
	orderCount := &smodel.OrderPoolAddInfo{}
	odExdb := m.centDs.DB().Raw(orderSql, serverKeyId, sdkType, beginTime, endTime).Scan(orderCount)
	if odExdb.Error != nil && odExdb.Error != gorm.ErrRecordNotFound {
		return nil, odExdb.Error
	}
	return orderCount, nil
}

func (m *serverSupportPool) AddPlatformSupportPoolSet(p_centerPlatformId int64, p_supportGold int32, p_supportRate int32) error {
	now := timeutils.TimeToMillisecond(time.Now())
	info := &smodel.PlatformSupportPoolSetInfo{
		CenterPlatformId: p_centerPlatformId,
		SupportGold:      p_supportGold,
		SupportRate:      p_supportRate,
		CreateTime:       now,
	}
	exdb := m.db.DB().Save(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *serverSupportPool) UpdatePlatformSupportPoolSet(p_id int32, p_supportGold int32, p_supportRate int32) error {
	now := timeutils.TimeToMillisecond(time.Now())
	exdb := m.db.DB().Table("t_platform_supportpool_set").Where("id=?", p_id).Updates(map[string]interface{}{"updateTime": now, "supportGold": p_supportGold, "supportRate": p_supportRate})
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *serverSupportPool) DeletePlatformSupportPoolSet(p_id int32) error {
	now := timeutils.TimeToMillisecond(time.Now())
	exdb := m.db.DB().Table("t_platform_supportpool_set").Where("id=?", p_id).Updates(map[string]interface{}{"deleteTime": now})
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *serverSupportPool) GetPlatformSupportPoolList(p_centerPlatformId int64, p_pageindex int32) ([]*smodel.PlatformSupportPoolSetInfo, error) {
	rst := make([]*smodel.PlatformSupportPoolSetInfo, 0)
	offset := (p_pageindex - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	where := "deleteTime = 0"
	if p_centerPlatformId > 0 {
		where += fmt.Sprintf(" and centerPlatformId=%d", p_centerPlatformId)
	}
	exdb := m.db.DB().Where(where).Offset(offset).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *serverSupportPool) GetPlatformSupportPoolCount(p_centerPlatformId int64) (int32, error) {
	totalCount := int32(0)
	where := "deleteTime = 0"
	if p_centerPlatformId > 0 {
		where += fmt.Sprintf(" and centerPlatformId=%d", p_centerPlatformId)
	}
	exdb := m.db.DB().Table("t_platform_supportpool_set").Where(where).Count(&totalCount)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return int32(0), exdb.Error
	}
	return totalCount, nil
}

func (m *serverSupportPool) GetPlatformSupportPoolSet(p_centerPlatformId int64) (*smodel.PlatformSupportPoolSetInfo, error) {
	info := &smodel.PlatformSupportPoolSetInfo{}
	exdb := m.db.DB().Where("centerPlatformId=? and deleteTime=0", p_centerPlatformId).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return info, nil
}

func (m *serverSupportPool) Heartbeat() {
	log.Debug("serverSupportPool Heartbeat running")
	err := m.AddOrderPoolAmount()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("serverSupportPool Heartbeat error")
	}
}

func (m *serverSupportPool) FillAllServerPoolSet() error {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	now := timeutils.TimeToMillisecond(time.Now())

	allPlatformPool := make([]*smodel.PlatformSupportPoolSetInfo, 0)
	exdb := m.db.DB().Where("deleteTime = 0").Find(&allPlatformPool)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	if len(allPlatformPool) == 0 {
		return nil
	}

	allCenterServer := make([]*centerservermodel.CenterServer, 0)
	exdb = m.centDs.DB().Where("serverType=0 and deleteTime=0").Find(&allCenterServer)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}

	allServerPool := make([]*smodel.ServerSupportPool, 0)
	exdb = m.db.DB().Where("deleteTime=0").Find(&allServerPool)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}

	allPlatformInfo := make([]*centerservermodel.CenterPlatformInfo, 0)
	exdb = m.db.DB().Where("deleteTime=0").Find(&allPlatformInfo)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}

	allPlatformPoolMap := make(map[int64]*smodel.PlatformSupportPoolSetInfo)
	for _, value := range allPlatformPool {
		allPlatformPoolMap[value.CenterPlatformId] = value
	}
	allServerPoolMap := make(map[int64]*smodel.ServerSupportPool)
	for _, value := range allServerPool {
		allServerPoolMap[int64(value.ServerId)] = value
	}

	for _, value := range allCenterServer {
		_, exists := allServerPoolMap[value.Id]
		if exists {
			continue
		}
		platformSet, exists := allPlatformPoolMap[value.Platform]
		if !exists {
			continue
		}
		item := &smodel.ServerSupportPool{
			ServerId:         int(value.Id),
			BeginGold:        int(platformSet.SupportGold),
			CurGold:          int(platformSet.SupportGold),
			CreateTime:       now,
			CenterPlatformId: platformSet.CenterPlatformId,
			OrderGoldPer:     platformSet.SupportRate,
			BeginOrderTime:   now,
			CurOrderTime:     now,
		}
		for _, plValue := range allPlatformInfo {
			if platformSet.CenterPlatformId == plValue.PlatformId {
				item.SdkType = plValue.SkdType
			}
		}

		exdb = m.db.DB().Save(item)
		if exdb.Error != nil {
			return exdb.Error
		}
	}

	return nil
}

func (m *serverSupportPool) fillServerSupportPool(serverId int32) error {

	serverInfo := &centerservermodel.CenterServer{}
	exdb := m.centDs.DB().Where("serverType=0 and deleteTime=0 and id = ?", serverId).First(&serverInfo)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	if serverInfo == nil || serverInfo.Id == 0 {
		return nil
	}
	supportPlatformPool := &smodel.PlatformSupportPoolSetInfo{}
	exdb = m.db.DB().Where("centerPlatformId=? and deleteTime=0", serverInfo.Platform).First(supportPlatformPool)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	if supportPlatformPool == nil || supportPlatformPool.Id == 0 {
		return nil
	}
	platformInfo := &centerservermodel.CenterPlatformInfo{}
	exdb = m.centDs.DB().Where("id=? and deleteTime=0", serverInfo.Platform).First(&platformInfo)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	if platformInfo == nil || platformInfo.PlatformId == 0 {
		return nil
	}
	now := timeutils.TimeToMillisecond(time.Now())
	item := &smodel.ServerSupportPool{
		ServerId:         int(serverInfo.Id),
		BeginGold:        int(supportPlatformPool.SupportGold),
		CurGold:          int(supportPlatformPool.SupportGold),
		CreateTime:       now,
		CenterPlatformId: serverInfo.Platform,
		OrderGoldPer:     supportPlatformPool.SupportRate,
		BeginOrderTime:   now,
		CurOrderTime:     now,
		SdkType:          platformInfo.SkdType,
	}
	exdb = m.db.DB().Save(item)
	if exdb.Error != nil {
		return exdb.Error
	}

	return nil
}

func NewServerSupportPool(p_db gmdb.DBService, p_rs redis.RedisService, p_centds gmdb.DBService) IServerSupportPool {
	rst := &serverSupportPool{
		db:     p_db,
		rs:     p_rs,
		centDs: p_centds,
	}
	return rst
}

type contextKey string

const (
	serverSupportPoolKey = contextKey("ServerSupportPool")
)

func WithServerSupportPool(ctx context.Context, ls IServerSupportPool) context.Context {
	return context.WithValue(ctx, serverSupportPoolKey, ls)
}

func ServerSupportPoolInContext(ctx context.Context) IServerSupportPool {
	us, ok := ctx.Value(serverSupportPoolKey).(IServerSupportPool)
	if !ok {
		return nil
	}
	return us
}

func SetupServerSupportPoolHandler(ls IServerSupportPool) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithServerSupportPool(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
