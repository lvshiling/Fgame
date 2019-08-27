package service

import (
	"context"
	"fgame/fgame/gm/gamegm/common"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	ordermodel "fgame/fgame/gm/gamegm/gm/center/order/model"
	pubmodel "fgame/fgame/gm/gamegm/gm/center/order/pubmodel"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/codegangsta/negroni"
)

type countInfo struct {
	TotalCount int `gorm:"column:totalCount"`
}

type IOrderService interface {
	GetOrderList(p_orderId string, p_sdkOrderId string, p_sdkType int, p_startTime int64, p_endTime int64, p_index int, p_sdkList []int) ([]*ordermodel.OrderInfo, error)
	GetOrderCount(p_orderId string, p_sdkOrderId string, p_sdkType int, p_startTime int64, p_endTime int64, p_sdkList []int) (int, error)
	GetCenterOrderStatic(p_sdkType int, p_sdkList []int) ([]*ordermodel.CenterOrderStatic, error)
	GetCenterOrderTotalStatic(p_sdkType int, p_sdkList []int) ([]*ordermodel.CenterOrderStatic, error)
	GetOrderAmountPlayerList(p_startTime int64, p_endTime int64, p_startMoney int, p_endMoney int, p_sdkType int, p_serverIndex int, p_sdkList []int) ([]int64, error)

	GetGameOrderList(p_dblink gmdb.GameDbLink, p_serverId int, p_startTime int64, p_endTime int64, p_minAmount int, p_maxAmount int, p_playerId int64, p_userId int64, p_orderId string, p_sdkOrderId string, p_name string, p_index int, p_col int, p_asc int, p_sdkType int, p_sdkList []int) ([]*ordermodel.GameOrderInfo, error)
	GetGameOrderCount(p_dblink gmdb.GameDbLink, p_serverId int, p_startTime int64, p_endTime int64, p_minAmount int, p_maxAmount int, p_playerId int64, p_userId int64, p_orderId string, p_sdkOrderId string, p_name string, p_sdkType int, p_sdkList []int) (int, error)
	GetGameOrderStatic(p_dblink gmdb.GameDbLink, p_serverId int, p_sdkType int, p_sdkList []int) ([]*ordermodel.GameOrderStatic, error)

	GetCenterOrderStaticMultiple(p_serverMap map[int][]int, p_startTime int64, p_endTime int64, p_sdkList []int) ([]*ordermodel.CenterOrderDateStatic, error)
	GetCenterOrderStaticPlatform(p_startTime int64, p_endTime int64, p_sdkList []int) ([]*ordermodel.CenterOrderDatePlatformStatic, error)
	GetPlayerCountDate(p_serverMap map[int][]int, p_startTime int64, p_endTime int64, p_sdkList []int) (map[int64]*pubmodel.UserDateCount, error)

	//计算中心服的订单统计按照skd类型和服务器id，job作业调用
	GetCenterServerOrderStaticDaily(begin int64, end int64) ([]*ordermodel.CenterOrderServerDailyStatic, error)
}

type orderService struct {
	db gmdb.DBService
}

func (m *orderService) GetOrderList(p_orderId string, p_sdkOrderId string, p_sdkType int, p_startTime int64, p_endTime int64, p_index int, p_sdkList []int) ([]*ordermodel.OrderInfo, error) {
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	where := "deleteTime = 0 and status IN (1,2)"
	if len(p_orderId) > 0 {
		where += fmt.Sprintf(" AND orderId LIKE '%s'", "%"+p_orderId+"%")
	}
	if len(p_sdkOrderId) > 0 {
		where += fmt.Sprintf(" AND sdkOrderId LIKE '%s'", "%"+p_sdkOrderId+"%")
	}
	if p_startTime > 0 {
		where += fmt.Sprintf(" and updateTime >= %d", p_startTime)
	}
	if p_endTime > 0 {
		where += fmt.Sprintf(" and updateTime <= %d", p_endTime)
	}
	if p_sdkType > 0 {
		where += fmt.Sprintf(" and sdkType = %d", p_sdkType)
	}
	if len(p_sdkList) > 0 {
		where += fmt.Sprintf(" and sdkType IN (%s)", common.CombinIntArray(p_sdkList))
	}
	rst := make([]*ordermodel.OrderInfo, 0)

	errdb := m.db.DB().Where(where).Order("createTime desc").Offset(offset).Limit(limit).Find(&rst)
	if errdb.Error != nil {
		return nil, errdb.Error
	}
	return rst, nil
}

func (m *orderService) GetOrderCount(p_orderId string, p_sdkOrderId string, p_sdkType int, p_startTime int64, p_endTime int64, p_sdkList []int) (int, error) {
	where := "deleteTime = 0 and status IN (1,2)"
	if len(p_orderId) > 0 {
		where += fmt.Sprintf(" AND orderId LIKE '%s'", "%"+p_orderId+"%")
	}
	if len(p_sdkOrderId) > 0 {
		where += fmt.Sprintf(" AND sdkOrderId LIKE '%s'", "%"+p_sdkOrderId+"%")
	}
	if p_startTime > 0 {
		where += fmt.Sprintf(" and updateTime >= %d", p_startTime)
	}
	if p_endTime > 0 {
		where += fmt.Sprintf(" and updateTime <= %d", p_endTime)
	}
	if p_sdkType > 0 {
		where += fmt.Sprintf(" and sdkType = %d", p_sdkType)
	}
	if len(p_sdkList) > 0 {
		where += fmt.Sprintf(" and sdkType IN (%s)", common.CombinIntArray(p_sdkList))
	}
	rst := 0
	errdb := m.db.DB().Table("t_order").Where(where).Count(&rst)
	if errdb.Error != nil {
		return 0, errdb.Error
	}
	return rst, nil
}

var (
	gameOrderColumnOrderMap = map[int]string{
		1:  "id",
		2:  "serverId",
		3:  "orderId",
		4:  "orderStatus",
		5:  "userId",
		6:  "playerId",
		7:  "chargeId",
		8:  "money",
		9:  "updateTime",
		10: "createTime",
		11: "deleteTime",
		12: "playerLevel",
		13: "gold",
		14: "name",
	}
)

func (m *orderService) GetGameOrderList(p_dblink gmdb.GameDbLink, p_serverId int, p_startTime int64, p_endTime int64, p_minAmount int, p_maxAmount int, p_playerId int64, p_userId int64, p_orderId string, p_sdkOrderId string, p_name string, p_index int, p_col int, p_asc int, p_sdkType int, p_sdkList []int) ([]*ordermodel.GameOrderInfo, error) {
	rst := make([]*ordermodel.GameOrderInfo, 0)
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	offect := (p_index - 1) * constant.DefaultPageSize
	if offect < 0 {
		offect = 0
	}
	orderStr := getGameOrderBy(p_col, p_asc)
	whereStr := ""
	if p_startTime > 0 {
		whereStr += fmt.Sprintf(" and A.createTime >= %d", p_startTime)
	}
	if p_endTime > 0 {
		whereStr += fmt.Sprintf(" and A.createTime <= %d", p_endTime)
	}
	if p_minAmount > 0 {
		whereStr += fmt.Sprintf(" and A.money >= %d", p_minAmount)
	}
	if p_maxAmount > 0 {
		whereStr += fmt.Sprintf(" and A.money <= %d", p_maxAmount)
	}
	if p_playerId > 0 {
		whereStr += fmt.Sprintf(" and A.playerId = %d", p_playerId)
	}
	if p_userId > 0 {
		whereStr += fmt.Sprintf(" and A.userId = %d", p_userId)
	}
	if len(p_orderId) > 0 {
		whereStr += fmt.Sprintf(" and A.orderId = '%s'", p_orderId)
	}
	if len(p_sdkList) > 0 {
		whereStr += fmt.Sprintf(" and B.sdkType IN (%s)", common.CombinIntArray(p_sdkList))
	}
	if p_sdkType > 0 {
		whereStr += fmt.Sprintf(" and B.sdkType = %d", p_sdkType)
	}

	sql := `SELECT
	A.id
	,A.serverId
	,A.orderId
	,A.orderStatus
	,B.userId
	,A.playerId
	,A.chargeId
	,A.money
	,A.updateTime
	,A.createTime
	,A.deleteTime
	,A.playerLevel
	,A.gold
	,B.name
FROM t_order A
INNER JOIN t_player B
ON A.playerId = B.id
WHERE A.deleteTime = 0 AND B.name LIKE ? and A.serverId=? ` + whereStr + " ORDER BY " + orderStr + fmt.Sprintf(" limit %d,%d", offect, constant.DefaultPageSize)

	exdb := db.DB().Raw(sql, "%"+p_name+"%", p_serverId).Scan(&rst)
	// exdb := db.DB().Where("deleteTime=0 and name like ?", "%"+p_name+"%").Order(orderStr).Offset(offect).Limit(constant.DefaultPageSize).Find(&rst)
	if exdb.Error != nil {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *orderService) GetGameOrderCount(p_dblink gmdb.GameDbLink, p_serverId int, p_startTime int64, p_endTime int64, p_minAmount int, p_maxAmount int, p_playerId int64, p_userId int64, p_orderId string, p_sdkOrderId string, p_name string, p_sdkType int, p_sdkList []int) (int, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	whereStr := ""
	if p_startTime > 0 {
		whereStr += fmt.Sprintf(" and A.createTime >= %d", p_startTime)
	}
	if p_endTime > 0 {
		whereStr += fmt.Sprintf(" and A.createTime <= %d", p_endTime)
	}
	if p_minAmount > 0 {
		whereStr += fmt.Sprintf(" and A.money >= %d", p_minAmount)
	}
	if p_maxAmount > 0 {
		whereStr += fmt.Sprintf(" and A.money <= %d", p_maxAmount)
	}
	if p_playerId > 0 {
		whereStr += fmt.Sprintf(" and A.playerId = %d", p_playerId)
	}
	if p_userId > 0 {
		whereStr += fmt.Sprintf(" and A.userId = %d", p_userId)
	}
	if len(p_orderId) > 0 {
		whereStr += fmt.Sprintf(" and A.orderId = '%s'", p_orderId)
	}
	if len(p_sdkList) > 0 {
		whereStr += fmt.Sprintf(" and B.sdkType IN (%s)", common.CombinIntArray(p_sdkList))
	}
	if p_sdkType > 0 {
		whereStr += fmt.Sprintf(" and B.sdkType = %d", p_sdkType)
	}

	sql := `SELECT
	COUNT(1) AS totalCount
FROM t_order A
INNER JOIN t_player B
ON A.playerId = B.id
WHERE A.deleteTime = 0 AND B.name LIKE ? and A.serverId=? ` + whereStr

	info := &countInfo{}

	exdb := db.DB().Raw(sql, "%"+p_name+"%", p_serverId).Scan(&info)
	// exdb := db.DB().Where("deleteTime=0 and name like ?", "%"+p_name+"%").Order(orderStr).Offset(offect).Limit(constant.DefaultPageSize).Find(&rst)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return info.TotalCount, nil
}

func (m *orderService) GetGameOrderStatic(p_dblink gmdb.GameDbLink, p_serverId int, p_sdkType int, p_sdkList []int) ([]*ordermodel.GameOrderStatic, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	rst := make([]*ordermodel.GameOrderStatic, 0)
	today, _ := timeutils.BeginOfNow()
	yestoday, _ := timeutils.BeginOfYesterday()
	beginMonth := timeutils.TimeToMillisecond(time.Now().AddDate(0, -1, 0))
	whereStr := ""
	if len(p_sdkList) > 0 {
		whereStr += fmt.Sprintf(" and B.sdkType IN (%s)", common.CombinIntArray(p_sdkList))
	}
	if p_sdkType > 0 {
		whereStr += fmt.Sprintf(" and B.sdkType = %d", p_sdkType)
	}
	sql := `SELECT SUM(A.money) AS totalAmount
	,COUNT(DISTINCT A.playerId) as totalPerson
	,COUNT(DISTINCT B.id) AS totalRegisterPerson
	,SUM(CASE WHEN A.createTime >= ? THEN A.money ELSE 0 END) AS todayAmount
	,COUNT(DISTINCT CASE WHEN A.createTime >= ? THEN A.playerId ELSE NULL END) AS todayPerson
	,COUNT(DISTINCT CASE WHEN B.createTime >= ? THEN B.id ELSE NULL END) AS todayRegisterPerson
	,SUM(CASE WHEN A.createTime >= ? and A.createTime < ? THEN A.money ELSE 0 END) AS yestodayAmount
	,COUNT(DISTINCT CASE WHEN A.createTime >= ? and A.createTime < ? THEN A.playerId ELSE NULL END) AS yestodayPerson
	,COUNT(DISTINCT CASE WHEN B.createTime >= ? and B.createTime < ? THEN B.id ELSE NULL END) AS yestodayRegisterPerson
	,SUM(CASE WHEN A.createTime >= ? THEN A.money ELSE 0 END) AS monthAmount
	,COUNT(DISTINCT CASE WHEN A.createTime >= ? THEN A.playerId ELSE NULL END) AS monthPerson
FROM t_player B
LEFT JOIN t_order A
ON A.playerId = B.id AND A.serverId=?
WHERE B.serverId=? and B.deleteTime = 0 ` + whereStr
	exdb := db.DB().Raw(sql, today, today, today, yestoday, today, yestoday, today, yestoday, today, beginMonth, beginMonth, p_serverId, p_serverId).Scan(&rst)
	if exdb.Error != nil {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *orderService) GetCenterOrderStatic(p_sdkType int, p_sdkList []int) ([]*ordermodel.CenterOrderStatic, error) {
	rst := make([]*ordermodel.CenterOrderStatic, 0)
	today, _ := timeutils.BeginOfNow()
	yestoday, _ := timeutils.BeginOfYesterday()
	beginMonth := timeutils.TimeToMillisecond(time.Now().AddDate(0, -1, 0))
	where := ""
	if p_sdkType > 0 {
		where += fmt.Sprintf(" and A.platform = %d", p_sdkType)
	}
	if len(p_sdkList) > 0 {
		where += fmt.Sprintf(" and A.platform IN (%s)", common.CombinIntArray(p_sdkList))
	}
	sql := fmt.Sprintf(`SELECT A.platform AS sdkType
	,COUNT(DISTINCT B.id) AS totalRegisterPerson
	,COUNT(DISTINCT CASE WHEN B.createTime >= ? THEN B.id ELSE NULL END) AS todayRegisterPerson
	,COUNT(DISTINCT CASE WHEN B.createTime >= ? and B.createTime < ? THEN B.id ELSE NULL END) AS yestodayRegisterPerson
FROM t_user A
	INNER JOIN t_player B
	ON A.id = B.userId
WHERE A.deleteTime=0 AND B.deleteTime=0 %s
GROUP BY A.platform
ORDER BY A.platform ASC`, where)
	exdb := m.db.DB().Raw(sql, today, yestoday, today).Scan(&rst)
	if exdb.Error != nil {
		return rst, exdb.Error
	}
	orderRst, err := m.getCenterOrderStatic(p_sdkType, p_sdkList, today, yestoday, beginMonth)
	if err != nil {
		return rst, err
	}
	for _, value := range rst {
		for _, orderValue := range orderRst {
			if value.SdkType != orderValue.SdkType {
				continue
			}
			value.MonthAmount = orderValue.MonthAmount
			value.MonthPerson = orderValue.MonthPerson
			value.TotalAmount = orderValue.TotalAmount
			value.TotalPerson = orderValue.TotalPerson
			value.TodayAmount = orderValue.TodayAmount
			value.TodayPerson = orderValue.TodayPerson
			value.YestodayAmount = orderValue.YestodayAmount
			value.YestodayPerson = orderValue.YestodayPerson
		}
	}
	return rst, nil
}

var (
	getCenterOrderStaticSql = `SELECT C.sdkType
	,SUM(money) AS totalAmount
	,COUNT(DISTINCT C.playerId) AS totalPerson
	,SUM(CASE WHEN C.updateTime >= ? THEN C.money ELSE 0 END) AS todayAmount
	,COUNT(DISTINCT CASE WHEN C.updateTime >= ? THEN C.playerId ELSE NULL END) AS todayPerson
	,SUM(CASE WHEN C.updateTime >= ? and C.updateTime < ? THEN C.money ELSE 0 END) AS yestodayAmount
	,COUNT(DISTINCT CASE WHEN C.updateTime >= ? and C.updateTime < ? THEN C.playerId ELSE NULL END) AS yestodayPerson
	,SUM(CASE WHEN C.updateTime >= ? THEN C.money ELSE 0 END) AS monthAmount
	,COUNT(DISTINCT CASE WHEN C.updateTime >= ? THEN C.playerId ELSE NULL END) AS monthPerson
FROM t_order C
WHERE C.status IN (1,2) %s
GROUP BY C.sdkType`
)

func (m *orderService) getCenterOrderStatic(p_sdkType int, p_sdkList []int, p_today int64, p_yestoday int64, p_beginMonth int64) ([]*ordermodel.CenterOrderStatic, error) {
	rst := make([]*ordermodel.CenterOrderStatic, 0)
	where := ""
	if p_sdkType > 0 {
		where += fmt.Sprintf(" and C.sdkType = %d", p_sdkType)
	}
	if len(p_sdkList) > 0 {
		where += fmt.Sprintf(" and C.sdkType IN (%s)", common.CombinIntArray(p_sdkList))
	}
	sql := fmt.Sprintf(getCenterOrderStaticSql, where)
	exdb := m.db.DB().Raw(sql, p_today, p_today, p_yestoday, p_today, p_yestoday, p_today, p_beginMonth, p_beginMonth).Scan(&rst)
	if exdb.Error != nil {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *orderService) GetCenterOrderTotalStatic(p_sdkType int, p_sdkList []int) ([]*ordermodel.CenterOrderStatic, error) {
	rst := make([]*ordermodel.CenterOrderStatic, 0)
	today, _ := timeutils.BeginOfNow()
	yestoday, _ := timeutils.BeginOfYesterday()
	beginMonth := timeutils.TimeToMillisecond(time.Now().AddDate(0, -1, 0))
	where := ""
	if p_sdkType > 0 {
		where += fmt.Sprintf(" and A.platform = %d", p_sdkType)
	}
	if len(p_sdkList) > 0 {
		where += fmt.Sprintf(" and A.platform IN (%s)", common.CombinIntArray(p_sdkList))
	}
	sql := fmt.Sprintf(`SELECT 
	COUNT(DISTINCT B.id) AS totalRegisterPerson
	,COUNT(DISTINCT CASE WHEN B.createTime >= ? THEN B.id ELSE NULL END) AS todayRegisterPerson
	,COUNT(DISTINCT CASE WHEN B.createTime >= ? and B.createTime < ? THEN B.id ELSE NULL END) AS yestodayRegisterPerson
FROM t_user A
	INNER JOIN t_player B
	ON A.id = B.userId
WHERE A.deleteTime=0 AND B.deleteTime=0 %s`, where)
	exdb := m.db.DB().Raw(sql, today, yestoday, today).Scan(&rst)
	if exdb.Error != nil {
		return rst, exdb.Error
	}
	orderRst, err := m.getCenterOrderTotalStatic(p_sdkType, p_sdkList, today, yestoday, beginMonth)
	if err != nil {
		return rst, err
	}
	for _, value := range rst {
		for _, orderValue := range orderRst {
			if value.SdkType != orderValue.SdkType {
				continue
			}
			value.MonthAmount = orderValue.MonthAmount
			value.MonthPerson = orderValue.MonthPerson
			value.TotalAmount = orderValue.TotalAmount
			value.TotalPerson = orderValue.TotalPerson
			value.TodayAmount = orderValue.TodayAmount
			value.TodayPerson = orderValue.TodayPerson
			value.YestodayAmount = orderValue.YestodayAmount
			value.YestodayPerson = orderValue.YestodayPerson
		}
	}
	return rst, nil
}

var (
	getCenterOrderTotalStaticSql = `SELECT 
	SUM(money) AS totalAmount
	,COUNT(DISTINCT C.playerId) AS totalPerson
	,SUM(CASE WHEN C.updateTime >= ? THEN C.money ELSE 0 END) AS todayAmount
	,COUNT(DISTINCT CASE WHEN C.updateTime >= ? THEN C.playerId ELSE NULL END) AS todayPerson
	,SUM(CASE WHEN C.updateTime >= ? and C.updateTime < ? THEN C.money ELSE 0 END) AS yestodayAmount
	,COUNT(DISTINCT CASE WHEN C.updateTime >= ? and C.updateTime < ? THEN C.playerId ELSE NULL END) AS yestodayPerson
	,SUM(CASE WHEN C.updateTime >= ? THEN C.money ELSE 0 END) AS monthAmount
	,COUNT(DISTINCT CASE WHEN C.updateTime >= ? THEN C.playerId ELSE NULL END) AS monthPerson
FROM t_order C
WHERE C.status IN (1,2) %s`
)

func (m *orderService) getCenterOrderTotalStatic(p_sdkType int, p_sdkList []int, p_today int64, p_yestoday int64, p_beginMonth int64) ([]*ordermodel.CenterOrderStatic, error) {
	rst := make([]*ordermodel.CenterOrderStatic, 0)
	where := ""
	if p_sdkType > 0 {
		where += fmt.Sprintf(" and C.sdkType = %d", p_sdkType)
	}
	if len(p_sdkList) > 0 {
		where += fmt.Sprintf(" and C.sdkType IN (%s)", common.CombinIntArray(p_sdkList))
	}
	sql := fmt.Sprintf(getCenterOrderTotalStaticSql, where)
	exdb := m.db.DB().Raw(sql, p_today, p_today, p_yestoday, p_today, p_yestoday, p_today, p_beginMonth, p_beginMonth).Scan(&rst)
	if exdb.Error != nil {
		return rst, exdb.Error
	}
	return rst, nil
}

type orderAmountPlayerList struct {
	PlayerId int64 `gorm:"column:playerId"`
}

func (m *orderService) GetOrderAmountPlayerList(p_startTime int64, p_endTime int64, p_startMoney int, p_endMoney int, p_sdkType int, p_serverIndex int, p_sdkList []int) ([]int64, error) {
	tempList := make([]*orderAmountPlayerList, 0)
	where := ""
	if p_sdkType > 0 {
		where += fmt.Sprintf(" AND A.sdkType=%d", p_sdkType)
	}
	if len(p_sdkList) > 0 {
		where += fmt.Sprintf(" AND A.sdkType IN (%s)", common.CombinIntArray(p_sdkList))
	}
	sql := fmt.Sprintf(`SELECT
		A.playerId
	FROM
		t_order A
	WHERE
		A.deleteTime = 0 AND A.status IN (1,2) AND A.updateTime >= ? and A.updateTime < ? %s
	GROUP BY A.playerId
	HAVING SUM(A.money) >= ? AND SUM(A.money) < ?`, where)
	exdb := m.db.DB().Raw(sql, p_startTime, p_endTime, p_startMoney, p_endMoney).Scan(&tempList)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	rst := make([]int64, 0)
	for _, value := range tempList {
		rst = append(rst, value.PlayerId)
	}
	return rst, nil
}

func (m *orderService) GetCenterOrderStaticMultiple(p_serverMap map[int][]int, p_startTime int64, p_endTime int64, p_sdkList []int) ([]*ordermodel.CenterOrderDateStatic, error) {
	rst := make([]*ordermodel.CenterOrderDateStatic, 0)
	where := ""
	if len(p_sdkList) > 0 {
		where += fmt.Sprintf(" AND A.sdkType IN (%s)", common.CombinIntArray(p_sdkList))
	}
	sdkWhereList := make([]int, 0)
	serverWhere := ""
	index := 0
	for key, value := range p_serverMap {
		sdkWhereList = append(sdkWhereList, key)
		if index > 0 {
			serverWhere += " OR "
		}
		if len(value) > 0 {
			serverWhere += fmt.Sprintf("(A.sdkType=%d AND A.serverId IN (%s))", key, common.CombinIntArray(value))
			index++
		}
	}
	if len(sdkWhereList) > 0 {
		where += fmt.Sprintf(" AND A.sdkType IN (%s)", common.CombinIntArray(sdkWhereList))
	}
	if len(serverWhere) != 0 {
		where += fmt.Sprintf(" AND (%s)", serverWhere)
	}
	sql := `SELECT
	DATE_FORMAT(FROM_UNIXTIME(A.updateTime/1000),'%Y-%m-%d') AS orderDate
	,COUNT(DISTINCT A.playerId) AS orderPlayerNum
	,COUNT(A.id) AS orderNum
	,SUM(A.money) AS orderMoney
	,SUM(A.gold) AS orderGold
	,COUNT(DISTINCT B.id) AS orderNewPlayer
	,SUM(CASE WHEN B.id is NULL THEN 0 ELSE A.money END) AS orderNewMoney
	,COUNT(DISTINCT C.playerId) AS orderFirstPlayer
	,SUM(CASE WHEN C.playerId IS NULL THEN 0 ELSE A.money END) AS orderFirstMoney
FROM
	t_order A
	LEFT JOIN t_player B
	ON A.playerId = B.id 
	AND B.createTime >= UNIX_TIMESTAMP(DATE_FORMAT(FROM_UNIXTIME(A.updateTime/1000),'%Y-%m-%d'))*1000
	AND B.createTime < UNIX_TIMESTAMP(DATE_FORMAT(FROM_UNIXTIME(A.updateTime/1000),'%Y-%m-%d'))*1000+86400000
	LEFT JOIN v_player_firstorder C
	ON A.playerId = C.playerId
	AND C.firstOrderTime >= UNIX_TIMESTAMP(DATE_FORMAT(FROM_UNIXTIME(A.updateTime/1000),'%Y-%m-%d'))*1000
	AND C.firstOrderTime < UNIX_TIMESTAMP(DATE_FORMAT(FROM_UNIXTIME(A.updateTime/1000),'%Y-%m-%d'))*1000+86400000 
WHERE
	A.status IN (1,2) 
	AND A.deleteTime = 0
	AND A.updateTime >= ?
	AND A.updateTime < ? ` + where + ` 
GROUP BY DATE_FORMAT(FROM_UNIXTIME(A.updateTime/1000),'%Y-%m-%d')
ORDER BY OrderDate ASC`
	exdb := m.db.DB().Raw(sql, p_startTime, p_endTime).Scan(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *orderService) GetCenterOrderStaticPlatform(p_startTime int64, p_endTime int64, p_sdkList []int) ([]*ordermodel.CenterOrderDatePlatformStatic, error) {
	rst := make([]*ordermodel.CenterOrderDatePlatformStatic, 0)
	where := ""
	if len(p_sdkList) > 0 {
		where += fmt.Sprintf(" AND A.sdkType IN (%s)", common.CombinIntArray(p_sdkList))
	}

	sql := `SELECT
	sdkType
	,COUNT(DISTINCT A.playerId) AS orderPlayerNum
	,COUNT(A.id) AS orderNum
	,SUM(A.money) AS orderMoney
	,SUM(A.gold) AS orderGold
FROM
	t_order A
WHERE
	A.status IN (1,2) 
	AND A.deleteTime = 0
	AND A.createTime >= ?
	AND A.createTime < ? ` + where + ` 
GROUP BY A.sdkType
ORDER BY A.sdkType ASC`
	exdb := m.db.DB().Raw(sql, p_startTime, p_endTime).Scan(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

type getPlayerCountDateResult struct {
	Count int `gorm:"column:totalCount"`
}

var (
	getPlayerCountDateSql = `SELECT COUNT(1) AS totalCount
	FROM
		t_user A
		INNER JOIN t_player B
		ON A.id = B.userId
	WHERE
		A.deleteTime = 0 AND B.deleteTime = 0 AND B.createTime >= ? AND B.createTime < ?
		%s`
)

//gorm暂时找不到动态列的读取方式，先循环
func (m *orderService) GetPlayerCountDate(p_serverMap map[int][]int, p_startTime int64, p_endTime int64, p_sdkList []int) (map[int64]*pubmodel.UserDateCount, error) {
	rst := make(map[int64]*pubmodel.UserDateCount)
	where := ""
	if len(p_sdkList) > 0 {
		where += fmt.Sprintf(" AND A.platform IN (%s)", common.CombinIntArray(p_sdkList))
	}
	sdkWhereList := make([]int, 0)
	serverWhere := ""
	index := 0
	for key, value := range p_serverMap {
		sdkWhereList = append(sdkWhereList, key)
		if index > 0 {
			serverWhere += " OR "
		}
		if len(value) > 0 {
			serverWhere += fmt.Sprintf("(A.platform=%d AND B.serverId IN (%s))", key, common.CombinIntArray(value))
			index++
		}
	}
	if len(sdkWhereList) > 0 {
		where += fmt.Sprintf(" AND A.platform IN (%s)", common.CombinIntArray(sdkWhereList))
	}
	if len(serverWhere) != 0 {
		where += fmt.Sprintf(" AND (%s)", serverWhere)
	}
	sql := fmt.Sprintf(getPlayerCountDateSql, where)

	totalNum := 0
	for tempBegin := p_startTime; tempBegin < p_endTime; tempBegin = tempBegin + 24*60*60*1000 {
		if tempBegin == p_startTime {
			info1 := &getPlayerCountDateResult{}
			start1 := int64(0)
			end1 := tempBegin
			exdb := m.db.DB().Raw(sql, start1, end1).Scan(&info1)
			if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
				return rst, exdb.Error
			}
			totalNum = totalNum + info1.Count
			rstModel := &pubmodel.UserDateCount{
				DateTime:   0,
				LeiJiCount: totalNum,
				DateCount:  info1.Count,
			}
			rst[0] = rstModel
		}
		info := &getPlayerCountDateResult{}
		start := tempBegin
		end := tempBegin + 24*60*60*1000
		exdb := m.db.DB().Raw(sql, start, end).Scan(&info)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			return rst, exdb.Error
		}
		totalNum = totalNum + info.Count
		rstModelDt := &pubmodel.UserDateCount{
			DateTime:   0,
			LeiJiCount: totalNum,
			DateCount:  info.Count,
		}
		rst[tempBegin] = rstModelDt
	}
	return rst, nil
}

var (
	serverOrderDailySql = `SELECT A.sdkType
	,A.serverId
	,COUNT(DISTINCT A.playerId) AS orderPlayerNum
	,COUNT(A.id) AS orderNum
	,SUM(A.money) AS orderMoney
	,SUM(A.gold) AS orderGold
FROM t_order A
WHERE A.updateTime >= ? and A.updateTime < ? and A.status IN (2) AND A.deleteTime = 0
GROUP BY A.sdkType,A.serverId`
)

func (m *orderService) GetCenterServerOrderStaticDaily(begin int64, end int64) ([]*ordermodel.CenterOrderServerDailyStatic, error) {
	rst := make([]*ordermodel.CenterOrderServerDailyStatic, 0)
	exdb := m.db.DB().Raw(serverOrderDailySql, begin, end).Scan(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *orderService) getdb(p_dblink gmdb.GameDbLink) gmdb.DBService {
	return gmdb.GetDb(p_dblink)
}

func NewOrderService(p_db gmdb.DBService) IOrderService {
	rst := &orderService{
		db: p_db,
	}
	return rst
}

func getGameColName(p_id int) string {
	if value, ok := gameOrderColumnOrderMap[p_id]; ok {
		return value
	}
	return "id"
}

func getGameOrderBy(p_ordercol int, p_ordertype int) string {
	colName := getGameColName(p_ordercol)
	asc := "asc"
	if p_ordertype > 0 {
		asc = "desc"
	}
	return " " + colName + " " + asc
}

type contextKey string

const (
	orderServiceKey = contextKey("OrderService")
)

func WithOrderService(ctx context.Context, ls IOrderService) context.Context {
	return context.WithValue(ctx, orderServiceKey, ls)
}

func OrderServiceInContext(ctx context.Context) IOrderService {
	us, ok := ctx.Value(orderServiceKey).(IOrderService)
	if !ok {
		return nil
	}
	return us
}

func SetupOrderServiceHandler(ls IOrderService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithOrderService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
