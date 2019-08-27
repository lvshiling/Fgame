package tick

import (
	"fgame/fgame/gm/gamegm/db"
	"fgame/fgame/gm/gamegm/gm/center/model"
	"fgame/fgame/pkg/timeutils"
	"time"

	log "github.com/Sirupsen/logrus"

	poolmodel "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/model"

	"github.com/jinzhu/gorm"
)

type IGmTickService interface {
	StartTick()
	StopTick()
}

type gmTickService struct {
	centerdb    db.DBService
	gmdb        db.DBService
	done        chan struct{}
	nextAddTime int64
}

func (m *gmTickService) StartTick() {
	log.Debug("开始定时...")
	beginNow, _ := timeutils.BeginOfDayOfTime(time.Now())
	m.nextAddTime = beginNow + 24*60*60*1000 + 10*1000
	go func() {
	loop:
		for {
			select {
			case <-time.After(time.Second * 1):
				{
					err := m.tickSharePool()
					if err != nil {
						log.WithFields(log.Fields{
							"error": err,
						}).Error("tick连接池异常")
					}
				}
			case <-m.done:
				{
					break loop
				}
			}
		}
	}()
}

func (m *gmTickService) StopTick() {
	m.done <- struct{}{}
}

const (
	orderstSql = `SELECT sdkType,serverId,SUM(gold) AS gold
	FROM
		t_order
	WHERE
		updateTime >= ? and updateTime < ? and status IN (1,2)
	GROUP BY
		sdkType,serverId`
)

func (m *gmTickService) tickSharePool() error {
	now := timeutils.TimeToMillisecond(time.Now())
	if m.nextAddTime > now {
		return nil
	}
	m.nextAddTime = m.nextAddTime + 24*60*60*1000
	end, _ := timeutils.BeginOfDayOfTime(time.Now())
	start := end - 24*60*60*1000
	skdTypeList := make([]*gmSdkType, 0)
	exdb := m.gmdb.DB().Where("deleteTime = 0").Find(&skdTypeList)
	if exdb.Error != nil {
		return exdb.Error
	}
	sdkMap := make(map[int]int)
	for _, value := range skdTypeList {
		sdkMap[value.SdkType] = value.CenterPlatformId
	}

	orderList := make([]*centerOrderSt, 0)
	exdb = m.centerdb.DB().Raw(orderstSql, start, end).Scan(&orderList)
	if exdb.Error != nil {
		return exdb.Error
	}
	for _, value := range orderList {
		if platformid, ok := sdkMap[value.SdkType]; ok {
			serverinfo := &model.CenterServer{}
			serdb := m.centerdb.DB().Where("platform=? and serverId = ? and serverType=0", platformid, value.ServerId).First(serverinfo)
			if serdb.Error != nil && serdb.Error != gorm.ErrRecordNotFound {
				return serdb.Error
			}
			if serdb.Error != nil && serdb.Error == gorm.ErrRecordNotFound {
				continue
			}
			serverid := serverinfo.Id
			poolinfo := &poolmodel.ServerSupportPool{}
			gmPooldb := m.gmdb.DB().Where("serverId=? and deleteTime = 0", serverid).First(poolinfo)
			if gmPooldb.Error != nil && gmPooldb.Error != gorm.ErrRecordNotFound {
				return gmPooldb.Error
			}
			poolinfo.CurGold = poolinfo.CurGold + value.Gold
			if poolinfo.CenterPlatformId == 0 {
				poolinfo.CenterPlatformId = int64(platformid)
				poolinfo.ServerId = int(serverid)
				poolinfo.CreateTime = timeutils.TimeToMillisecond(time.Now())
			}
			svdb := m.gmdb.DB().Save(poolinfo)
			if svdb.Error != nil {
				return svdb.Error
			}
		}
	}

	return nil
}

func NewTickService(p_center db.DBService, p_gm db.DBService) IGmTickService {
	rst := &gmTickService{
		centerdb: p_center,
		gmdb:     p_gm,
	}
	return rst
}

func (m *gmTickService) getdb(p_dblink db.GameDbLink) db.DBService {
	return db.GetDb(p_dblink)
}

type centerOrderSt struct {
	SdkType  int `gorm:"column:sdkType"`
	ServerId int `gorm:"column:serverId"`
	Gold     int `gorm:"column:gold"`
}

type gmSdkType struct {
	SdkType          int `gorm:"column:sdkType"`
	CenterPlatformId int `gorm:"column:centerPlatformId"`
}

func (m *gmSdkType) TableName() string {
	return "t_platform"
}
