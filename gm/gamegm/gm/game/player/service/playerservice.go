package service

import (
	"context"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	playmodel "fgame/fgame/gm/gamegm/gm/game/player/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/codegangsta/negroni"
)

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
		25: "yesterdayChargeMoney",
		26: "todayChargeMoney",
		27: "totalChargeGold",
		28: "totalPrivilegeChargeGold",
		29: "sdkType",
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

type IPlayerService interface {
	GetPlayerList(p_dblink gmdb.GameDbLink, p_serverId int, p_index int, p_name string, p_col int, p_asc int) ([]*playmodel.QueryPlayer, error)
	GetPlayerCount(p_dblink gmdb.GameDbLink, p_serverId int, p_name string) (int, error)

	GetPlayerFengJinList(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string, p_index int) ([]*playmodel.GamePlayer, error)
	GetPlayerFengJinCount(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string) (int, error)

	GetPlayerJinYanList(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string, p_index int) ([]*playmodel.GamePlayer, error)
	GetPlayerJinYanCount(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string) (int, error)

	GetPlayerIgnoreList(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string, p_index int) ([]*playmodel.GamePlayer, error)
	GetPlayerIgnoreCount(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string) (int, error)

	GetPlayerInfo(p_dblink gmdb.GameDbLink, p_playerId int64) (*playmodel.GamePlayer, error)
	GetPlayerInfoByUserId(p_dblink gmdb.GameDbLink, p_serverId int, p_userId int64) (*playmodel.GamePlayer, error)
	GetOriginPlayerInfoByUserId(p_dblink gmdb.GameDbLink, p_serverId int, p_userId int64) (*playmodel.GamePlayer, error)
	GetPlayerMailList(p_dblink gmdb.GameDbLink, p_playerId int64, p_begin int64, p_end int64, p_page int) ([]*playmodel.PlayerMail, error)
	GetPlayerMailCount(p_dblink gmdb.GameDbLink, p_playerId int64, p_begin int64, p_end int64) (int, error)

	GetPlayerLevelStatic(p_dblink gmdb.GameDbLink, p_serverId int) ([]*playmodel.QueryPlayerLevelStatic, error)
	GetPlayerQuestPlayerCountStatic(p_dblink gmdb.GameDbLink, p_serverId int) ([]*playmodel.QueryPlayerQuestStatic, error)

	GetPlayerTopThreePower(p_dblink gmdb.GameDbLink, p_serverId int) ([]*playmodel.QueryPlayerPower, error)
}

type playerService struct {
}

func (m *playerService) GetPlayerList(p_dblink gmdb.GameDbLink, p_serverId int, p_index int, p_name string, p_col int, p_asc int) ([]*playmodel.QueryPlayer, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	offect := (p_index - 1) * constant.DefaultPageSize
	if offect < 0 {
		offect = 0
	}
	beginNow, _ := timeutils.BeginOfDayOfTime(time.Now())
	beginYestoday := beginNow - int64(time.Hour*24/time.Millisecond)
	orderStr := getOrderBy(p_col, p_asc)
	rst := make([]*playmodel.QueryPlayer, 0)

	sql := fmt.Sprintf(`SELECT
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
	,A.totalChargeMoney
	,A.totalChargeGold
	,A.totalPrivilegeChargeGold
	,A.ip
	,A.originServerId
	,CASE WHEN A.chargeTime >= %d THEN A.todayChargeMoney WHEN A.chargeTime >= %d AND A.chargeTime < %d THEN A.yesterdayChargeMoney ELSE 0 END AS todayChargeMoney
	,CASE WHEN A.chargeTime >= %d THEN A.yesterdayChargeMoney ELSE 0 END AS yesterdayChargeMoney
	,A.sdkType
FROM 
	t_player A
	LEFT JOIN t_player_property B
	ON A.id = B.playerId AND B.deleteTime=0
	LEFT JOIN t_player_marry C
	ON A.id = C.playerId AND C.deleteTime=0
	LEFT JOIN t_player_alliance D
	ON A.id = D.playerId AND D.deleteTime=0
WHERE
	A.deleteTime=0 AND A.name LIKE ? and serverId=? ORDER BY `, beginNow, beginYestoday, beginNow, beginNow)
	sql += orderStr + fmt.Sprintf(" limit %d,%d", offect, constant.DefaultPageSize)

	exdb := db.DB().Raw(sql, "%"+p_name+"%", p_serverId).Scan(&rst)
	// exdb := db.DB().Where("deleteTime=0 and name like ?", "%"+p_name+"%").Order(orderStr).Offset(offect).Limit(constant.DefaultPageSize).Find(&rst)
	if exdb.Error != nil {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerCount(p_dblink gmdb.GameDbLink, p_serverId int, p_name string) (int, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := 0
	exdb := db.DB().Table("t_player").Where("deleteTime=0 and name like ? and serverId=?", "%"+p_name+"%", p_serverId).Count(&rst)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerFengJinList(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string, p_index int) ([]*playmodel.GamePlayer, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	limit := constant.DefaultPageSize
	offect := (p_index - 1) * limit
	if offect < 0 {
		offect = 0
	}
	rst := make([]*playmodel.GamePlayer, 0)
	exdb := db.DB().Where("deleteTime=0 and forbidTime > 0 and  serverId=? and name like ? and forbidText like ?", p_serverid, "%"+p_playerName+"%", "%"+p_reason+"%").Offset(offect).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerFengJinCount(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string) (int, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := 0
	exdb := db.DB().Table("t_player").Where("deleteTime=0 and forbidTime > 0 and  serverId=? and name like ? and forbidText like ?", p_serverid, "%"+p_playerName+"%", "%"+p_reason+"%").Count(&rst)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerJinYanList(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string, p_index int) ([]*playmodel.GamePlayer, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	limit := constant.DefaultPageSize
	offect := (p_index - 1) * limit
	if offect < 0 {
		offect = 0
	}
	rst := make([]*playmodel.GamePlayer, 0)
	exdb := db.DB().Where("deleteTime=0 and forbidChatTime > 0 and  serverId=? and name like ? and forbidChatText like ?", p_serverid, "%"+p_playerName+"%", "%"+p_reason+"%").Offset(offect).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerJinYanCount(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string) (int, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := 0
	exdb := db.DB().Table("t_player").Where("deleteTime=0 and forbidChatTime > 0 and  serverId=? and name like ? and forbidChatText like ?", p_serverid, "%"+p_playerName+"%", "%"+p_reason+"%").Count(&rst)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerIgnoreList(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string, p_index int) ([]*playmodel.GamePlayer, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	limit := constant.DefaultPageSize
	offect := (p_index - 1) * limit
	if offect < 0 {
		offect = 0
	}
	rst := make([]*playmodel.GamePlayer, 0)
	exdb := db.DB().Where("deleteTime=0 and ignoreChatTime > 0 and  serverId=? and name like ? and ignoreChatText like ?", p_serverid, "%"+p_playerName+"%", "%"+p_reason+"%").Offset(offect).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerIgnoreCount(p_dblink gmdb.GameDbLink, p_serverid int32, p_playerName string, p_reason string) (int, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := 0
	exdb := db.DB().Table("t_player").Where("deleteTime=0 and ignoreChatTime > 0 and  serverId=? and name like ? and ignoreChatText like ?", p_serverid, "%"+p_playerName+"%", "%"+p_reason+"%").Count(&rst)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerInfo(p_dblink gmdb.GameDbLink, p_playerId int64) (*playmodel.GamePlayer, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	rst := &playmodel.GamePlayer{}
	exdb := db.DB().Table("t_player").Where("id = ?", p_playerId).First(&rst)
	if exdb.Error != nil {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerInfoByUserId(p_dblink gmdb.GameDbLink, p_serverId int, p_userId int64) (*playmodel.GamePlayer, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	rst := &playmodel.GamePlayer{}
	exdb := db.DB().Table("t_player").Where("serverId = ? and userId = ?", p_serverId, p_userId).First(&rst)
	if exdb.Error != nil {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetOriginPlayerInfoByUserId(p_dblink gmdb.GameDbLink, p_serverId int, p_userId int64) (*playmodel.GamePlayer, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	rst := &playmodel.GamePlayer{}
	exdb := db.DB().Table("t_player").Where("originServerId = ? and userId = ?", p_serverId, p_userId).First(&rst)
	if exdb.Error != nil {
		if exdb.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerMailList(p_dblink gmdb.GameDbLink, p_playerId int64, p_begin int64, p_end int64, p_index int) ([]*playmodel.PlayerMail, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	limit := constant.DefaultPageSize
	offect := (p_index - 1) * limit
	if offect < 0 {
		offect = 0
	}
	rst := make([]*playmodel.PlayerMail, 0)
	exdb := db.DB().Where("playerId = ? and  createTime >= ? and createTime < ?", p_playerId, p_begin, p_end).Offset(offect).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *playerService) GetPlayerMailCount(p_dblink gmdb.GameDbLink, p_playerId int64, p_begin int64, p_end int64) (int, error) {
	db := m.getdb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := 0
	exdb := db.DB().Table("t_player_email").Where("playerId = ? and  createTime >= ? and createTime < ?", p_playerId, p_begin, p_end).Count(&rst)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return rst, nil
}

var (
	playerLevelStatic = `SELECT A.level,COUNT(1) AS playerCount
	FROM t_player_property A
	INNER JOIN t_player B
	ON A.playerId = B.id
	WHERE B.serverId = ?
	GROUP BY A.level
	ORDER BY level ASC`
)

func (m *playerService) GetPlayerLevelStatic(p_dblink gmdb.GameDbLink, p_serverId int) ([]*playmodel.QueryPlayerLevelStatic, error) {
	rst := make([]*playmodel.QueryPlayerLevelStatic, 0)
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	exdb := db.DB().Raw(playerLevelStatic, p_serverId).Scan(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

var (
	playerQuestStatic = `SELECT A.questId,COUNT(distinct A.playerId) AS playerCount 
	FROM t_player_quest A
	INNER JOIN t_player B
	ON A.playerId = B.Id
	WHERE A.questState >=1 AND A.questState <=2 AND B.serverId=?
	GROUP BY A.questId
	ORDER BY A.questId ASC`
)

func (m *playerService) GetPlayerQuestPlayerCountStatic(p_dblink gmdb.GameDbLink, p_serverId int) ([]*playmodel.QueryPlayerQuestStatic, error) {
	rst := make([]*playmodel.QueryPlayerQuestStatic, 0)
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	exdb := db.DB().Raw(playerQuestStatic, p_serverId).Scan(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

var (
	getPlayerTopThreePowerSql = `SELECT
	B.playerId
	,B.power
FROM 
	t_player A
	INNER JOIN t_player_property B
	ON A.id = B.playerId
WHERE
	A.serverId = ?
ORDER BY B.power DESC
LIMIT 0,3`
)

func (m *playerService) GetPlayerTopThreePower(p_dblink gmdb.GameDbLink, p_serverId int) ([]*playmodel.QueryPlayerPower, error) {
	rst := make([]*playmodel.QueryPlayerPower, 0)
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	exdb := db.DB().Raw(getPlayerTopThreePowerSql, p_serverId).Scan(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *playerService) getdb(p_dblink gmdb.GameDbLink) gmdb.DBService {
	return gmdb.GetDb(p_dblink)
}

func NewPlayService() IPlayerService {
	rst := &playerService{}
	return rst
}

type contextKey string

const (
	playServiceKey = contextKey("playerService")
)

func WithPlayerService(ctx context.Context, ls IPlayerService) context.Context {
	return context.WithValue(ctx, playServiceKey, ls)
}

func PlayerServiceInContext(ctx context.Context) IPlayerService {
	us, ok := ctx.Value(playServiceKey).(IPlayerService)
	if !ok {
		return nil
	}
	return us
}

func SetupPlayerServiceHandler(ls IPlayerService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithPlayerService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
