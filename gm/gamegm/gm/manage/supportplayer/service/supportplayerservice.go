package service

import (
	"context"
	"fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	playmodel "fgame/fgame/gm/gamegm/gm/game/player/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fmt"
	"net/http"
	"time"

	"fgame/fgame/gm/gamegm/common"

	"github.com/jinzhu/gorm"

	spmodel "fgame/fgame/gm/gamegm/gm/manage/supportplayer/model"

	"github.com/codegangsta/negroni"
)

type ISupportPlayerService interface {
	GetPlayerList(p_dblink gmdb.GameDbLink, p_serverId int, p_index int, p_name string, p_col int, p_asc int) ([]*playmodel.QueryPlayer, error)
	GetAllGmPlayerList(p_dblink gmdb.GameDbLink, p_serverId int) ([]*playmodel.QueryPlayer, error)
	GetPlayerCount(p_dblink gmdb.GameDbLink, p_serverId int, p_name string) (int, error)
	GetPlayerId(p_dblink gmdb.GameDbLink, p_serverId int, p_name string) (int64, error)
	GetPlayerName(p_dblink gmdb.GameDbLink, p_id int64) (string, error)
	CheckPlayerId(p_dblink gmdb.GameDbLink, p_serverId int, p_playerId int64) (int64, error)

	AddChargeLog(p_channelId int, p_platformId int, p_centerPlatformId int, p_serverId int, p_playerId int64, p_serverName string, p_gold int, p_userName string, p_reason string, p_playerName string) error
	GetChargeLogList(p_channelId int, p_platformId int, p_serverId int, p_playerName string, p_playerId int64, p_pageindex int, p_platformList []int64) ([]*spmodel.PrivilegeChargeLog, error)
	GetChargeLogCount(p_channelId int, p_platformId int, p_serverId int, p_playerName string, p_playerId int64, p_platformList []int64) (int, error)
	CheckPlayerIdFuchi(p_dblink gmdb.GameDbLink, p_id int64) (bool, error)
}

var (
	playerColumnOrderMap = map[int]string{
		1:  "id",
		2:  "userId",
		3:  "serverId",
		4:  "name",
		5:  "role",
		6:  "sex",
		7:  "lastLoginTime",
		8:  "lastLogoutTime",
		9:  "onlineTime",
		10: "offlineTime",
		11: "totalOnlineTime",
		12: "todayOnlineTime",
		13: "createTime",
		14: "level",
		15: "zhuanSheng",
		16: "silver",
		17: "gold",
		18: "bindGold",
		19: "yuanshi",
		20: "allianceName",
		21: "spouseName",
		22: "charm",
		23: "power",
		24: "totalChargeMoney",
		25: "totalChargeGold",
		26: "totalPrivilegeChargeGold",
	}
)

func getColName(p_id int) string {
	if value, ok := playerColumnOrderMap[p_id]; ok {
		return value
	}
	return "id"
}

func getOrderBy(p_ordercol int, p_ordertype int) string {
	colName := getColName(p_ordercol)
	asc := "asc"
	if p_ordertype > 0 {
		asc = "desc"
	}
	return " " + colName + " " + asc
}

type supportPlayerService struct {
	db gmdb.DBService
}

func (m *supportPlayerService) GetPlayerList(p_dblink gmdb.GameDbLink, p_serverId int, p_index int, p_name string, p_col int, p_asc int) ([]*playmodel.QueryPlayer, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	offect := (p_index - 1) * constant.DefaultPageSize
	if offect < 0 {
		offect = 0
	}
	orderStr := getOrderBy(p_col, p_asc)
	rst := make([]*playmodel.QueryPlayer, 0)

	sql := `SELECT
	A.id
	,A.userId
	,A.serverId
	,A.originServerId
	,A.name
	,A.role
	,A.sex
	,A.lastLoginTime
	,A.lastLogoutTime
	,A.onlineTime
	,A.offlineTime
	,A.totalOnlineTime
	,A.todayOnlineTime
	,A.updateTime
	,A.createTime
	,A.deleteTime
	,A.forbid
	,B.level
	,B.zhuanSheng
	,B.silver
	,B.gold
	,B.bindGold
	,0 as yuanshi
	,D.allianceName
	,C.spouseId
	,C.spouseName
	,B.charm
	,B.power
	,A.privilegeType
	,A.totalChargeMoney
	,A.totalPrivilegeChargeGold
	,A.totalChargeGold
FROM 
	t_player A
	LEFT JOIN t_player_property B
	ON A.id = B.playerId AND B.deleteTime=0
	LEFT JOIN t_player_marry C
	ON A.id = C.playerId AND C.deleteTime=0
	LEFT JOIN t_player_alliance D
	ON A.id = D.playerId AND D.deleteTime=0
WHERE
	A.deleteTime=0 AND A.privilegeType IN (1,2) AND A.name LIKE ? and serverId=? ORDER BY ` + orderStr + fmt.Sprintf(" limit %d,%d", offect, constant.DefaultPageSize)

	exdb := db.DB().Raw(sql, "%"+p_name+"%", p_serverId).Scan(&rst)
	// exdb := db.DB().Where("deleteTime=0 and name like ?", "%"+p_name+"%").Order(orderStr).Offset(offect).Limit(constant.DefaultPageSize).Find(&rst)
	if exdb.Error != nil {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *supportPlayerService) GetAllGmPlayerList(p_dblink gmdb.GameDbLink, p_serverId int) ([]*playmodel.QueryPlayer, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}

	rst := make([]*playmodel.QueryPlayer, 0)

	sql := `SELECT
	A.id
	,A.userId
	,A.serverId
	,A.name
	,A.role
	,A.sex
	,A.lastLoginTime
	,A.lastLogoutTime
	,A.onlineTime
	,A.offlineTime
	,A.totalOnlineTime
	,A.todayOnlineTime
	,A.updateTime
	,A.createTime
	,A.deleteTime
	,A.forbid
	,B.level
	,B.zhuanSheng
	,B.silver
	,B.gold
	,B.bindGold
	,0 as yuanshi
	,D.allianceName
	,C.spouseId
	,C.spouseName
	,B.charm
	,B.power
	,A.privilegeType
	,A.totalChargeMoney
	,A.totalPrivilegeChargeGold
	,A.totalChargeGold
FROM 
	t_player A
	LEFT JOIN t_player_property B
	ON A.id = B.playerId AND B.deleteTime=0
	LEFT JOIN t_player_marry C
	ON A.id = C.playerId AND C.deleteTime=0
	LEFT JOIN t_player_alliance D
	ON A.id = D.playerId AND D.deleteTime=0
WHERE
	A.deleteTime=0 AND A.privilegeType IN (1,2) and serverId=? `

	exdb := db.DB().Raw(sql, p_serverId).Scan(&rst)
	// exdb := db.DB().Where("deleteTime=0 and name like ?", "%"+p_name+"%").Order(orderStr).Offset(offect).Limit(constant.DefaultPageSize).Find(&rst)
	if exdb.Error != nil {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *supportPlayerService) GetPlayerCount(p_dblink gmdb.GameDbLink, p_serverId int, p_name string) (int, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := 0
	exdb := db.DB().Table("t_player").Where("deleteTime=0 and privilegeType IN (1,2) and name like ? and serverId=?", "%"+p_name+"%", p_serverId).Count(&rst)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return rst, nil
}

func (m *supportPlayerService) GetPlayerId(p_dblink gmdb.GameDbLink, p_serverId int, p_name string) (int64, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := &playmodel.GamePlayer{}
	exdb := db.DB().Where("deleteTime=0 and name = ? and serverId=?", p_name, p_serverId).First(rst)

	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return 0, exdb.Error
	}
	return rst.Id, nil
}

func (m *supportPlayerService) GetPlayerName(p_dblink gmdb.GameDbLink, p_id int64) (string, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return "", fmt.Errorf("DB服务为空")
	}
	rst := &playmodel.GamePlayer{}
	exdb := db.DB().Where("deleteTime=0 and id = ?", p_id).First(rst)

	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return "", exdb.Error
	}
	return rst.Name, nil
}

func (m *supportPlayerService) CheckPlayerId(p_dblink gmdb.GameDbLink, p_serverId int, p_playerId int64) (int64, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := &playmodel.GamePlayer{}
	exdb := db.DB().Where("deleteTime=0 and id = ? and serverId=?", p_playerId, p_serverId).First(rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return 0, exdb.Error
	}
	return rst.Id, nil
}

func (m *supportPlayerService) CheckPlayerIdFuchi(p_dblink gmdb.GameDbLink, p_id int64) (bool, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return false, fmt.Errorf("DB服务为空")
	}
	rst := 0
	exdb := db.DB().Table("t_player").Where("deleteTime=0 and privilegeType IN (1,2) and id = ?", p_id).Count(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return false, exdb.Error
	}
	return rst > 0, nil
}

func (m *supportPlayerService) AddChargeLog(p_channelId int, p_platformId int, p_centerPlatformId int, p_serverId int, p_playerId int64, p_serverName string, p_gold int, p_userName string, p_reason string, p_playerName string) error {
	now := timeutils.TimeToMillisecond(time.Now())
	info := &spmodel.PrivilegeChargeLog{
		ChannelId:        p_channelId,
		PlatformId:       p_platformId,
		CenterPlatformId: p_centerPlatformId,
		ServerId:         p_serverId,
		PlayerId:         p_playerId,
		ServerName:       p_serverName,
		Gold:             p_gold,
		UserName:         p_userName,
		Reason:           p_reason,
		ChargeTime:       now,
		CreateTime:       now,
		PlayerName:       p_playerName,
	}
	exdb := m.db.DB().Save(info)
	return exdb.Error
}

func (m *supportPlayerService) GetChargeLogList(p_channelId int, p_platformId int, p_serverId int, p_playerName string, p_playerId int64, p_pageindex int, p_platformList []int64) ([]*spmodel.PrivilegeChargeLog, error) {
	rst := make([]*spmodel.PrivilegeChargeLog, 0)
	offset := (p_pageindex - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	where := "deleteTime=0"
	if p_channelId > 0 {
		where += fmt.Sprintf(" and channelId = %d", p_channelId)
	}
	if p_platformId > 0 {
		where += fmt.Sprintf(" and platformId = %d", p_platformId)
	}
	if len(p_platformList) > 0 {
		where += fmt.Sprintf(" and centerPlatformId IN (%s)", common.CombinInt64Array(p_platformList))
	}
	if p_serverId > 0 {
		where += fmt.Sprintf(" and serverId = %d", p_serverId)
	}

	if len(p_playerName) > 0 {
		where += fmt.Sprintf(" and playerName like '%s'", "%"+p_playerName+"%")
	}

	if p_playerId > 0 {
		where += fmt.Sprintf(" and playerId = %d", p_playerId)
	}

	dberr := m.db.DB().Where(where).Order("id desc").Offset(offset).Limit(limit).Find(&rst)
	if dberr.Error != nil {
		return nil, dberr.Error
	}
	return rst, nil
}

func (m *supportPlayerService) GetChargeLogCount(p_channelId int, p_platformId int, p_serverId int, p_playerName string, p_playerId int64, p_platformList []int64) (int, error) {
	rst := 0

	where := "deleteTime=0"
	if p_channelId > 0 {
		where += fmt.Sprintf(" and channelId = %d", p_channelId)
	}
	if p_platformId > 0 {
		where += fmt.Sprintf(" and platformId = %d", p_platformId)
	}
	if len(p_platformList) > 0 {
		where += fmt.Sprintf(" and centerPlatformId IN (%s)", common.CombinInt64Array(p_platformList))
	}
	if p_serverId > 0 {
		where += fmt.Sprintf(" and serverId = %d", p_serverId)
	}

	if len(p_playerName) > 0 {
		where += fmt.Sprintf(" and playerName like '%s'", "%"+p_playerName+"%")
	}

	if p_playerId > 0 {
		where += fmt.Sprintf(" and playerId = %d", p_playerId)
	}

	dberr := m.db.DB().Table("t_charge_log").Where(where).Count(&rst)
	if dberr.Error != nil {
		return 0, dberr.Error
	}
	return rst, nil
}

func (m *supportPlayerService) getdb(p_dblink gmdb.GameDbLink) gmdb.DBService {
	return gmdb.GetDb(p_dblink)
}

func NewSupportPlayerService(p_db gmdb.DBService) ISupportPlayerService {
	rst := &supportPlayerService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	supportPlayerServiceKey = contextKey("supportplayerService")
)

func WithPlayerService(ctx context.Context, ls ISupportPlayerService) context.Context {
	return context.WithValue(ctx, supportPlayerServiceKey, ls)
}

func SupportPlayerServiceInContext(ctx context.Context) ISupportPlayerService {
	us, ok := ctx.Value(supportPlayerServiceKey).(ISupportPlayerService)
	if !ok {
		return nil
	}
	return us
}

func SetupPlayerServiceHandler(ls ISupportPlayerService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithPlayerService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
