package serverdailystat

import (
	"fgame/fgame/core/db"
	"fgame/fgame/core/mongo"
	"fgame/fgame/core/redis"
	stservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"time"

	centerorder "fgame/fgame/gm/gamegm/gm/center/order/service"
	centerservice "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmserverstmodel "fgame/fgame/gm/gamegm/gm/manage/serverdaily/model"
	dailyservice "fgame/fgame/gm/gamegm/gm/manage/serverdaily/service"
	gmplatform "fgame/fgame/gm/gamegm/gm/platform/service"

	log "github.com/Sirupsen/logrus"
)

const (
	oneDay = 24 * int64(time.Hour/time.Millisecond)
)

type ServerDailyStatJob struct {
	ds                       db.DBService //gmdb
	centerDs                 db.DBService //中心db
	rs                       redis.RedisService
	ms                       mongo.MongoService
	msConfig                 *mongo.MongoConfig
	rpStService              stservice.IStaticReportService
	serverService            centerservice.ICenterServerService
	orderservice             centerorder.IOrderService
	gmplatformService        gmplatform.IPlatformService
	gmServerDailyStatService dailyservice.ServerDailyStatService
}

func (m *ServerDailyStatJob) GetId() string {
	return "serverDailyStatJob"
}

func (m *ServerDailyStatJob) Run() error {
	log.Debug("作业ServerDailyStatJob开始运行...")
	lastTime, err := m.gmServerDailyStatService.GetMinDate()
	if err != nil {
		return err
	}
	if lastTime == 0 {
		log.Debug("作业ServerDailyStatJob，开始获取mongo中的数值")
		firstMongoTime, err := m.rpStService.GetOnLineServerFirstTime()
		if err != nil {
			return err
		}
		beginFirstMongoTime, err := timeutils.BeginOfNow(firstMongoTime)
		if err != nil {
			return err
		}
		lastTime = beginFirstMongoTime
	}
	beginNow, err := timeutils.BeginOfDayOfTime(time.Now())
	if err != nil {
		return err
	}
	if lastTime == 0 {
		lastTime = beginNow - oneDay
	}
	for {
		if lastTime >= beginNow {
			break
		}

		begin := lastTime
		end := lastTime + oneDay
		err := m.staticReport(begin, end)
		if err != nil {
			return err
		}
		lastTime = end
	}

	return nil
}

func (m *ServerDailyStatJob) staticReport(begin int64, end int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	serverStatCount, err := m.gmServerDailyStatService.GetStatCountByDate(begin)
	if err != nil {
		return err
	}
	if serverStatCount > 0 { //已经有了，就不计算了
		return nil
	}
	allServer, err := m.serverService.GetAllCenterServerList()
	if err != nil {
		return err
	}

	if len(allServer) == 0 {
		return fmt.Errorf("no server")
	}
	allPlatform, err := m.gmplatformService.GetAllPlatformList()
	if err != nil {
		return err
	}
	maxOnline, err := m.rpStService.GetOnLineStaticDaily(begin, end)
	if err != nil {
		return err
	}
	totalLogin, err := m.rpStService.GetServerLoginPlayerNum(begin, end)
	if err != nil {
		return err
	}
	orderAmount, err := m.orderservice.GetCenterServerOrderStaticDaily(begin, end)
	if err != nil {
		return err
	}

	sdkPlatformMap := make(map[int]int64)
	for _, value := range allPlatform {
		sdkPlatformMap[value.SdkType] = value.CenterPlatformID
	}
	// sdkPlMapJson, _ := json.Marshal(sdkPlatformMap)
	// orderAmountJson, _ := json.Marshal(orderAmount)
	// log.WithFields(log.Fields{
	// 	"begin":       begin,
	// 	"error":       string(sdkPlMapJson),
	// 	"orderAmount": string(orderAmountJson),
	// }).Error("获取的平台")

	saveArray := make([]*gmserverstmodel.ServerDailyStats, 0)
	for _, value := range allServer {
		if value.ServerType != 0 {
			continue
		}
		item := &gmserverstmodel.ServerDailyStats{
			ServerId:   int32(value.ServerId),
			PlatformId: int32(value.Platform),
			CurDate:    begin,
			CreateTime: now,
		}
		for _, onlineValue := range maxOnline {
			if value.Platform != int64(onlineValue.Id.Platform) || value.ServerId != onlineValue.Id.ServerId {
				continue
			}
			item.MaxOnlineNum = int32(onlineValue.MaxPlayer)
		}
		for _, loginValue := range totalLogin {
			if value.Platform != int64(loginValue.Id.Platform) || value.ServerId != loginValue.Id.ServerId {
				continue
			}
			item.LoginNum = int32(loginValue.TotalPlayer)
		}
		for _, orderValue := range orderAmount {
			platformId, exists := sdkPlatformMap[orderValue.SdkType]
			if !exists {
				continue
			}
			if value.Platform != platformId || value.ServerId != orderValue.ServerId {
				continue
			}
			item.OrderGold = orderValue.OrderGold
			item.OrderMoney = orderValue.OrderMoney
			item.OrderNum = orderValue.OrderNum
			item.OrderPlayerNum = orderValue.OrderPlayerNum
		}
		saveArray = append(saveArray, item)
	}

	if len(saveArray) > 0 {
		err := m.gmServerDailyStatService.SaveStatArray(saveArray)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *ServerDailyStatJob) GetTickSecond() int64 {
	return 1
}

/*******接口结束*************/

func NewServerDailyStatJob(ds db.DBService, rs redis.RedisService, ms mongo.MongoService, msc *mongo.MongoConfig, centerDs db.DBService) *ServerDailyStatJob {
	rst := &ServerDailyStatJob{
		ds:       ds,
		rs:       rs,
		ms:       ms,
		msConfig: msc,
		centerDs: centerDs,
	}
	rst.rpStService = stservice.NewReportStatic(rst.ms, rst.msConfig)
	rst.serverService = centerservice.NewCenterServerService(rst.centerDs)
	rst.orderservice = centerorder.NewOrderService(rst.centerDs)
	rst.gmplatformService = gmplatform.NewPlatformService(rst.ds)
	rst.gmServerDailyStatService = dailyservice.NewServerDailyStatService(rst.ds)
	return rst
}
