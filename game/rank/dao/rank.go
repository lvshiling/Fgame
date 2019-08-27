package dao

import (
	"encoding/json"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lingtongdev/types"
	marrytypes "fgame/fgame/game/marry/types"
	rankentity "fgame/fgame/game/rank/entity"
	ranktypes "fgame/fgame/game/rank/types"
	"fmt"
	"sync"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

type RankDao interface {
	//排行榜周期重置数据
	GetRankTimeEntity(serverId int32) (*rankentity.RankTimeEntity, error)
	//
	//玩家战斗力排行ds
	GetRankForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerForceData, err error)
	//玩家战斗力排行rs
	GetRedisRankForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerForceData, err error)
	//战斗力排行入rs
	SetRedisRankForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerForceData) (err error)
	//玩家坐骑排行ds
	GetRankMountList(config *ranktypes.RankConfig) (mountList []*rankentity.PlayerOrderData, err error)
	//玩家坐骑排行rs
	GetRedisRankMountList(timestamp int64, config *ranktypes.RankConfig) (mountList []*rankentity.PlayerOrderData, err error)
	//玩家坐骑排行入rs
	SetRedisRankMountList(timestamp int64, config *ranktypes.RankConfig, mountList []*rankentity.PlayerOrderData) (err error)
	//玩家战翼排行ds
	GetRankWingList(config *ranktypes.RankConfig) (wingList []*rankentity.PlayerOrderData, err error)
	//玩家战翼排行rs
	GetRedisRankWingList(timestamp int64, config *ranktypes.RankConfig) (wingList []*rankentity.PlayerOrderData, err error)
	//玩家战翼排行入rs
	SetRedisRankWingList(timestamp int64, config *ranktypes.RankConfig, wingList []*rankentity.PlayerOrderData) (err error)
	//玩家护体盾排行ds
	GetRankBodyShieldList(config *ranktypes.RankConfig) (bodyShieldList []*rankentity.PlayerOrderData, err error)
	//玩家护体盾排行rs
	GetRedisRankBodyShieldList(timestamp int64, config *ranktypes.RankConfig) (bodyShieldList []*rankentity.PlayerOrderData, err error)
	//玩家护体盾排行入rs
	SetRedisRankBodyShieldList(timestamp int64, config *ranktypes.RankConfig, bodyShieldList []*rankentity.PlayerOrderData) (err error)
	//玩家兵魂排行ds
	GetRankWeaponList(config *ranktypes.RankConfig) (weaponList []*rankentity.PlayerWeaponData, err error)
	//玩家兵魂排行rs
	GetRedisRankWeaponList(timestamp int64, config *ranktypes.RankConfig) (weaponList []*rankentity.PlayerWeaponData, err error)
	//玩家兵魂排行入rs
	SetRedisRankWeaponList(timestamp int64, config *ranktypes.RankConfig, weaponList []*rankentity.PlayerWeaponData) (err error)
	//玩家帮派排行ds
	GetRankGangList(config *ranktypes.RankConfig) (gangList []*rankentity.PlayerGangData, err error)
	//玩家帮派排行rs
	GetRedisRankGangList(timestamp int64, config *ranktypes.RankConfig) (gangList []*rankentity.PlayerGangData, err error)
	//玩家帮派排行入rs
	SetRedisRankGangList(timestamp int64, config *ranktypes.RankConfig, gangList []*rankentity.PlayerGangData) (err error)
	//玩家身法排行ds
	GetRankShenFaList(config *ranktypes.RankConfig) (shenFaList []*rankentity.PlayerOrderData, err error)
	//玩家身法排行rs
	GetRedisRankShenFaList(timestamp int64, config *ranktypes.RankConfig) (shenFaList []*rankentity.PlayerOrderData, err error)
	//玩家身法排行入rs
	SetRedisRankShenFaList(timestamp int64, config *ranktypes.RankConfig, shenFaList []*rankentity.PlayerOrderData) (err error)
	//玩家领域排行ds
	GetRankLingYuList(config *ranktypes.RankConfig) (lingYuList []*rankentity.PlayerOrderData, err error)
	//玩家领域排行rs
	GetRedisRankLingYuList(timestamp int64, config *ranktypes.RankConfig) (lingYuList []*rankentity.PlayerOrderData, err error)
	//玩家领域排行入rs
	SetRedisRankLingYuList(timestamp int64, config *ranktypes.RankConfig, lingYuList []*rankentity.PlayerOrderData) (err error)
	//玩家护体仙羽排行ds
	GetRankFeatherList(config *ranktypes.RankConfig) (featherList []*rankentity.PlayerOrderData, err error)
	//玩家护体仙羽排行rs
	GetRedisRankFeatherList(timestamp int64, config *ranktypes.RankConfig) (featherList []*rankentity.PlayerOrderData, err error)
	//玩家领域排行入rs
	SetRedisRankFeatherList(timestamp int64, config *ranktypes.RankConfig, featherList []*rankentity.PlayerOrderData) (err error)
	//玩家神盾尖刺排行ds
	GetRankShieldList(config *ranktypes.RankConfig) (shieldList []*rankentity.PlayerOrderData, err error)
	//玩家神盾尖刺排行rs
	GetRedisRankShieldList(timestamp int64, config *ranktypes.RankConfig) (shieldList []*rankentity.PlayerOrderData, err error)
	//玩家神盾尖刺排行入rs
	SetRedisRankShieldList(timestamp int64, config *ranktypes.RankConfig, shieldList []*rankentity.PlayerOrderData) (err error)
	//玩家活动充值排行ds
	GetRankChargeList(config *ranktypes.RankConfig) (chargeList []*rankentity.PlayerPropertyData, err error)
	//玩家活动充值排行rs
	GetRedisRankChargeList(timestamp int64, config *ranktypes.RankConfig) (chargeList []*rankentity.PlayerPropertyData, err error)
	//玩家活动充值排行入rs
	SetRedisRankChargeList(timestamp int64, config *ranktypes.RankConfig, chargeList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动消费排行ds
	GetRankCostList(config *ranktypes.RankConfig) (costList []*rankentity.PlayerPropertyData, err error)
	//玩家活动消费排行rs
	GetRedisRankCostList(timestamp int64, config *ranktypes.RankConfig) (costList []*rankentity.PlayerPropertyData, err error)
	//玩家活动消费排行入rs
	SetRedisRankCostList(timestamp int64, config *ranktypes.RankConfig, costList []*rankentity.PlayerPropertyData) (err error)
	//玩家暗器排行ds
	GetRankAnQiList(config *ranktypes.RankConfig) (anQiList []*rankentity.PlayerOrderData, err error)
	//玩家暗器排行rs
	GetRedisRankAnQiList(timestamp int64, config *ranktypes.RankConfig) (anQiList []*rankentity.PlayerOrderData, err error)
	//玩家暗器排行入rs
	SetRedisRankAnQiList(timestamp int64, config *ranktypes.RankConfig, anQiList []*rankentity.PlayerOrderData) (err error)
	//玩家魅力排行ds
	GetRankCharmList(config *ranktypes.RankConfig) (charmList []*rankentity.PlayerPropertyData, err error)
	//玩家魅力排行rs
	GetRedisRankCharmList(timestamp int64, config *ranktypes.RankConfig) (charmList []*rankentity.PlayerPropertyData, err error)
	//玩家魅力排行入rs
	SetRedisRankCharmList(timestamp int64, config *ranktypes.RankConfig, charmList []*rankentity.PlayerPropertyData) (err error)
	//玩家表白排行ds
	GetRankMarryDevelopList(config *ranktypes.RankConfig) (marryDevelopList []*rankentity.PlayerPropertyData, err error)
	//玩家表白排行rs
	GetRedisRankMarryDevelopList(timestamp int64, config *ranktypes.RankConfig) (marryDevelopList []*rankentity.PlayerPropertyData, err error)
	//玩家表白排行入rs
	SetRedisRankMarryDevelopList(timestamp int64, config *ranktypes.RankConfig, marryDevelopList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动次数排行ds
	GetRankCountList(config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error)
	//玩家活动次数排行rs
	GetRedisRankCountList(timestamp int64, config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error)
	//玩家活动次数排行入rs
	SetRedisRankCountList(timestamp int64, config *ranktypes.RankConfig, countList []*rankentity.PlayerPropertyData) (err error)
	//玩家法宝排行ds
	GetRankFaBaoList(config *ranktypes.RankConfig) (faBaoList []*rankentity.PlayerOrderData, err error)
	//玩家法宝排行rs
	GetRedisRankFaBaoList(timestamp int64, config *ranktypes.RankConfig) (faBaoList []*rankentity.PlayerOrderData, err error)
	//玩家法宝排行入rs
	SetRedisRankFaBaoList(timestamp int64, config *ranktypes.RankConfig, faBaoList []*rankentity.PlayerOrderData) (err error)
	//玩家仙体排行ds
	GetRankXianTiList(config *ranktypes.RankConfig) (xianTiList []*rankentity.PlayerOrderData, err error)
	//玩家仙体排行rs
	GetRedisRankXianTiList(timestamp int64, config *ranktypes.RankConfig) (xianTiList []*rankentity.PlayerOrderData, err error)
	//玩家仙体排行入rs
	SetRedisRankXianTiList(timestamp int64, config *ranktypes.RankConfig, xianTiList []*rankentity.PlayerOrderData) (err error)
	//玩家等级排行ds
	GetRankLevelList(config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error)
	//玩家等级排行rs
	GetRedisRankLevelList(timestamp int64, config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error)
	//玩家等级排行入rs
	SetRedisRankLevelList(timestamp int64, config *ranktypes.RankConfig, countList []*rankentity.PlayerPropertyData) (err error)
	//玩家灵童养成类排行入ds
	GetRedisRankLingTongDevList(timestamp int64, config *ranktypes.RankConfig) (lingTongDevList []*rankentity.PlayerOrderData, err error)
	//玩家灵童养成类排行rs
	GetRankLingTongDevList(config *ranktypes.RankConfig) (lingTongDevList []*rankentity.PlayerOrderData, err error)
	//玩家灵童养成类排行入rs
	SetRedisRankLingTongDevList(timestamp int64, config *ranktypes.RankConfig, lingTongDevList []*rankentity.PlayerOrderData) (err error)
	//玩家噬魂幡排行ds
	GetRankShiHunFanList(config *ranktypes.RankConfig) (shiHunFanList []*rankentity.PlayerOrderData, err error)
	//玩家噬魂幡排行rs
	GetRedisRankShiHunFanList(timestamp int64, config *ranktypes.RankConfig) (shiHunFanList []*rankentity.PlayerOrderData, err error)
	//玩家噬魂幡排行入rs
	SetRedisRankShiHunFanList(timestamp int64, config *ranktypes.RankConfig, shiHunFanList []*rankentity.PlayerOrderData) (err error)
	//玩家天魔体排行ds
	GetRankTianMoTiList(config *ranktypes.RankConfig) (tianMoTiList []*rankentity.PlayerOrderData, err error)
	//玩家天魔体排行rs
	GetRedisRankTianMoTiList(timestamp int64, config *ranktypes.RankConfig) (tianMoTiList []*rankentity.PlayerOrderData, err error)
	//玩家天魔体排行入rs
	SetRedisRankTianMoTiList(timestamp int64, config *ranktypes.RankConfig, tianMoTiList []*rankentity.PlayerOrderData) (err error)
	//玩家灵童等级排行ds
	GetRankLingTongLevelList(config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error)
	//玩家灵童等级排行rs
	GetRedisRankLingTongLevelList(timestamp int64, config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error)
	//玩家灵童等级排行入rs
	SetRedisRankLingTongLevelList(timestamp int64, config *ranktypes.RankConfig, countList []*rankentity.PlayerPropertyData) (err error)
	//玩家飞升排行ds
	GetRankFeiShengList(config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error)
	//玩家飞升排行rs
	GetRedisRankFeiShengList(timestamp int64, config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error)
	//玩家飞升排行入rs
	SetRedisRankFeiShengList(timestamp int64, config *ranktypes.RankConfig, countList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动灵童战力排行ds
	GetRankLingTongForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动灵童战力排行rs
	GetRedisRankLingTongForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动灵童战力排行入rs
	SetRedisRankLingTongForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动元神金装战力排行ds
	GetRankGoldEquipForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动元神金装战力排行rs
	GetRedisRankGoldEquipForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动元神金装战力排行入rs
	SetRedisRankGoldEquipForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动点星战力排行ds
	GetRankDianXingForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动点星战力排行rs
	GetRedisRankDianXingForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动点星战力排行入rs
	SetRedisRankDianXingForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动神器战力排行ds
	GetRankShenQiForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动神器战力排行rs
	GetRedisRankShenQiForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动神器战力排行入rs
	SetRedisRankShenQiForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动命格战力排行ds
	GetRankMingGeForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动命格战力排行rs
	GetRedisRankMingGeForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动命格战力排行入rs
	SetRedisRankMingGeForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动圣痕战力排行ds
	GetRankShengHenForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动圣痕战力排行rs
	GetRedisRankShengHenForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动圣痕战力排行入rs
	SetRedisRankShengHenForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动阵法战力排行ds
	GetRankZhenFaForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动阵法战力排行rs
	GetRedisRankZhenFaForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动阵法战力排行入rs
	SetRedisRankZhenFaForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动屠龙装战力排行ds
	GetRankTuLongEquipForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动屠龙装战力排行rs
	GetRedisRankTuLongEquipForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动屠龙装战力排行入rs
	SetRedisRankTuLongEquipForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error)
	//玩家活动宝宝战力排行ds
	GetRankBabyForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动宝宝战力排行rs
	GetRedisRankBabyForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error)
	//玩家活动宝宝战力排行入rs
	SetRedisRankBabyForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error)
	//玩家转生排行ds
	GetRankZhuanShengList(config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error)
	//玩家转生排行rs
	GetRedisRankZhuanShengList(timestamp int64, config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error)
	//玩家转生排行入rs
	SetRedisRankZhuanShengList(timestamp int64, config *ranktypes.RankConfig, countList []*rankentity.PlayerPropertyData) (err error)
}

type rankDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func getRankTypeByClassType(classTpye types.LingTongDevSysType) (rankType ranktypes.RankType) {
	switch classTpye {
	case types.LingTongDevSysTypeLingQi:
		rankType = ranktypes.RankTypeLingQi
	case types.LingTongDevSysTypeLingBing:
		rankType = ranktypes.RankTypeLingBing
	case types.LingTongDevSysTypeLingBao:
		rankType = ranktypes.RankTypeLingBao
	case types.LingTongDevSysTypeLingShen:
		rankType = ranktypes.RankTypeLingShen
	case types.LingTongDevSysTypeLingTi:
		rankType = ranktypes.RankTypeLingTi
	case types.LingTongDevSysTypeLingYi:
		rankType = ranktypes.RankTypeLingYi
	case types.LingTongDevSysTypeLingYu:
		rankType = ranktypes.RankTypeLingTongYu
	}
	return
}

func getSqlByRankConfigAndClassType(config *ranktypes.RankConfig, classTpye types.LingTongDevSysType) (rankSql string) {
	if config.ClassType == ranktypes.RankClassTypeLocal {
		switch classTpye {
		case types.LingTongDevSysTypeLingQi,
			types.LingTongDevSysTypeLingBing,
			types.LingTongDevSysTypeLingBao,
			types.LingTongDevSysTypeLingShen,
			types.LingTongDevSysTypeLingTi,
			types.LingTongDevSysTypeLingYi,
			types.LingTongDevSysTypeLingYu:
			rankSql = lingTongDevSql
		}
	} else if config.ClassType == ranktypes.RankClassTypeArea {
		switch classTpye {
		case types.LingTongDevSysTypeLingQi,
			types.LingTongDevSysTypeLingBing,
			types.LingTongDevSysTypeLingBao,
			types.LingTongDevSysTypeLingShen,
			types.LingTongDevSysTypeLingTi,
			types.LingTongDevSysTypeLingYi,
			types.LingTongDevSysTypeLingYu:
			rankSql = areaLingTongDevSql
		}
	} else if config.ClassType == ranktypes.RankClassTypeLocalActivity {
		switch classTpye {
		case types.LingTongDevSysTypeLingQi,
			types.LingTongDevSysTypeLingBing,
			types.LingTongDevSysTypeLingBao,
			types.LingTongDevSysTypeLingShen,
			types.LingTongDevSysTypeLingTi,
			types.LingTongDevSysTypeLingYi,
			types.LingTongDevSysTypeLingYu:
			rankSql = lingTongDevActivitySql
		}
	}
	return
}

//获取key
func getRankKey(config *ranktypes.RankConfig, rankType ranktypes.RankType, timestamp int64) (key string) {
	if config.ClassType == ranktypes.RankClassTypeLocal {
		serverId := global.GetGame().GetServerIndex()
		servKey := coreredis.Combine("serviceId", fmt.Sprintf("%d", serverId))
		rankKey := coreredis.Combine("rankType", fmt.Sprintf("%d", rankType))
		stampKey := coreredis.Combine("time", fmt.Sprintf("%d", timestamp))
		key = coreredis.Join(servKey, rankKey, stampKey)
		return

	} else if config.ClassType == ranktypes.RankClassTypeLocalActivity {
		serverId := global.GetGame().GetServerIndex()
		servKey := coreredis.Combine("serviceId", fmt.Sprintf("%d", serverId))
		groupKey := coreredis.Combine("groupId", fmt.Sprintf("%d", config.GroupId))
		rankKey := coreredis.Combine("rankType", fmt.Sprintf("%d", rankType))
		stampKey := coreredis.Combine("time", fmt.Sprintf("%d", config.MaxExpireTime))
		key = coreredis.Join(servKey, groupKey, rankKey, stampKey)

	} else {
		serverId := global.GetGame().GetServerIndex()
		servKey := coreredis.Combine("serviceId", fmt.Sprintf("%d", serverId))
		rankKey := coreredis.Combine("rankType", fmt.Sprintf("%d", rankType))
		stampKey := coreredis.Combine("time", fmt.Sprintf("%d", timestamp))
		key = coreredis.Join(servKey, rankKey, stampKey)
		return
	}
	return
}

//获取redis数据
func getRedisList(key string) (dataStr string, err error) {
	conn := dao.rs.Pool().Get()
	err = conn.Err()
	if err != nil {
		return
	}
	defer conn.Close()
	dataStr, err = redis.String(conn.Do("GET", key))
	if err != nil {
		if err != redis.ErrNil {
			return
		}
		return "", nil
	}
	return
}

//设置redis数据
func setRedisList(expireTime int64, key string, bytes []byte) (err error) {
	conn := dao.rs.Pool().Get()
	err = conn.Err()
	if err != nil {
		return
	}

	defer conn.Close()
	//set并设置过期时间
	_, err = conn.Do("SET", key, string(bytes), "EX", expireTime)
	if err != nil {
		return
	}
	return
}

//玩家战斗力排行ds
func (dao *rankDao) GetRankForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerForceData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(forceRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&forceList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaForceRankSql, ranktypes.ListLimit).Scan(&forceList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家战斗力排行rs
func (dao *rankDao) GetRedisRankForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerForceData, err error) {
	key := getRankKey(config, ranktypes.RankTypeForce, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &forceList)
	if err != nil {
		return
	}
	return
}

//设置rs-玩家战斗力排行
func (dao *rankDao) SetRedisRankForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerForceData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeForce, timestamp)
	bytes, err := json.Marshal(forceList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家坐骑排行ds
func (dao *rankDao) GetRankMountList(config *ranktypes.RankConfig) (mountList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(mountRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&mountList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(mountRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&mountList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaMountRankSql, ranktypes.ListLimit).Scan(&mountList).Error
		}
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家坐骑排行rs
func (dao *rankDao) GetRedisRankMountList(timestamp int64, config *ranktypes.RankConfig) (mountList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeMount, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &mountList)
	if err != nil {
		return
	}
	return
}

//设置rs-玩家坐骑排行
func (dao *rankDao) SetRedisRankMountList(timestamp int64, config *ranktypes.RankConfig, mountList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeMount, timestamp)
	bytes, err := json.Marshal(mountList)
	if err != nil {
		return err
	}

	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家战翼排行ds
func (dao *rankDao) GetRankWingList(config *ranktypes.RankConfig) (wingList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(wingRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&wingList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(wingRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&wingList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaWingRankSql, ranktypes.ListLimit).Scan(&wingList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家战翼排行rs
func (dao *rankDao) GetRedisRankWingList(timestamp int64, config *ranktypes.RankConfig) (wingList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeWing, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &wingList)
	if err != nil {
		return
	}
	return
}

//设置rs-玩家战翼排行
func (dao *rankDao) SetRedisRankWingList(timestamp int64, config *ranktypes.RankConfig, wingList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeWing, timestamp)

	bytes, err := json.Marshal(wingList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家护体盾排行ds
func (dao *rankDao) GetRankBodyShieldList(config *ranktypes.RankConfig) (bodyShieldList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(bodyShieldRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&bodyShieldList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(bodyShieldRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&bodyShieldList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaBodyShieldRankSql, ranktypes.ListLimit).Scan(&bodyShieldList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家护体盾排行rs
func (dao *rankDao) GetRedisRankBodyShieldList(timestamp int64, config *ranktypes.RankConfig) (bodyShieldList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeBodyShield, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &bodyShieldList)
	if err != nil {
		return
	}
	return
}

//设置rs-玩家护体盾排行
func (dao *rankDao) SetRedisRankBodyShieldList(timestamp int64, config *ranktypes.RankConfig, bodyShieldList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeBodyShield, timestamp)

	bytes, err := json.Marshal(bodyShieldList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家兵魂排行ds
func (dao *rankDao) GetRankWeaponList(config *ranktypes.RankConfig) (weaponList []*rankentity.PlayerWeaponData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(weaponRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&weaponList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaWeaponRankSql, ranktypes.ListLimit).Scan(&weaponList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家兵魂排行rs
func (dao *rankDao) GetRedisRankWeaponList(timestamp int64, config *ranktypes.RankConfig) (weaponList []*rankentity.PlayerWeaponData, err error) {
	key := getRankKey(config, ranktypes.RankTypeWeapon, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &weaponList)
	if err != nil {
		return
	}
	return
}

//设置rs-玩家兵魂排行
func (dao *rankDao) SetRedisRankWeaponList(timestamp int64, config *ranktypes.RankConfig, weaponList []*rankentity.PlayerWeaponData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeWeapon, timestamp)

	bytes, err := json.Marshal(weaponList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家帮派排行ds
func (dao *rankDao) GetRankGangList(config *ranktypes.RankConfig) (gangList []*rankentity.PlayerGangData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(gangRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&gangList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaGangRankSql, ranktypes.ListLimit).Scan(&gangList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家帮派排行rs
func (dao *rankDao) GetRedisRankGangList(timestamp int64, config *ranktypes.RankConfig) (gangList []*rankentity.PlayerGangData, err error) {
	key := getRankKey(config, ranktypes.RankTypeGang, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &gangList)
	if err != nil {
		return
	}
	return
}

//设置rs-玩家帮派排行
func (dao *rankDao) SetRedisRankGangList(timestamp int64, config *ranktypes.RankConfig, gangList []*rankentity.PlayerGangData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeGang, timestamp)

	bytes, err := json.Marshal(gangList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家活动充值排行ds
func (dao *rankDao) GetRankChargeList(config *ranktypes.RankConfig) (chargeList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(chargeRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&chargeList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(chargeRankActivitySql, global.GetGame().GetServerIndex(), config.GroupId, config.MinCondition, config.EndTime, ranktypes.ListLimit).Scan(&chargeList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaChargeRankSql, ranktypes.ListLimit).Scan(&chargeList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动充值排行rs
func (dao *rankDao) GetRedisRankChargeList(timestamp int64, config *ranktypes.RankConfig) (chargeList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeCharge, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &chargeList)
	if err != nil {
		return
	}
	return
}

//设置rs-玩家活动充值排行
func (dao *rankDao) SetRedisRankChargeList(timestamp int64, config *ranktypes.RankConfig, chargeList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeCharge, timestamp)

	bytes, err := json.Marshal(chargeList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//设置rs-玩家活动灵童战力排行
func (dao *rankDao) SetRedisRankLingTongForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeLingTongForce, timestamp)

	bytes, err := json.Marshal(forceList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家活动灵童战力排行rs
func (dao *rankDao) GetRedisRankLingTongForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeLingTongForce, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &forceList)
	if err != nil {
		return
	}
	return
}

//玩家活动灵童战力排行ds
func (dao *rankDao) GetRankLingTongForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		err = dao.ds.DB().Raw(lingTongForceRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeLocalActivity:
		err = dao.ds.DB().Raw(lingTongForceRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeArea:
		err = dao.ds.DB().Raw(areaLingTongForceRankSql, ranktypes.ListLimit).Scan(&forceList).Error
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//设置rs-玩家活动元神金装战力排行
func (dao *rankDao) SetRedisRankGoldEquipForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeGoldEquipForce, timestamp)

	bytes, err := json.Marshal(forceList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家活动元神金装战力排行rs
func (dao *rankDao) GetRedisRankGoldEquipForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeGoldEquipForce, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &forceList)
	if err != nil {
		return
	}
	return
}

//玩家活动元神金装战力排行ds
func (dao *rankDao) GetRankGoldEquipForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		err = dao.ds.DB().Raw(goldEquipForceRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeLocalActivity:
		err = dao.ds.DB().Raw(goldEquipForceRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeArea:
		err = dao.ds.DB().Raw(areaGoldEquipForceRankSql, ranktypes.ListLimit).Scan(&forceList).Error
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动点星战力排行ds
func (dao *rankDao) GetRankDianXingForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		err = dao.ds.DB().Raw(dianXingForceRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeLocalActivity:
		err = dao.ds.DB().Raw(dianXingForceRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeArea:
		err = dao.ds.DB().Raw(areaDianXingForceRankSql, ranktypes.ListLimit).Scan(&forceList).Error
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动点星战力排行rs
func (dao *rankDao) GetRedisRankDianXingForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeDianXingForce, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &forceList)
	if err != nil {
		return
	}
	return
}

//玩家活动点星战力排行入rs
func (dao *rankDao) SetRedisRankDianXingForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeDianXingForce, timestamp)

	bytes, err := json.Marshal(forceList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家活动神器战力排行ds
func (dao *rankDao) GetRankShenQiForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		err = dao.ds.DB().Raw(shenQiForceRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeLocalActivity:
		err = dao.ds.DB().Raw(shenQiForceRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeArea:
		err = dao.ds.DB().Raw(areaShenQiForceRankSql, ranktypes.ListLimit).Scan(&forceList).Error
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动神器战力排行rs
func (dao *rankDao) GetRedisRankShenQiForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeShenQiForce, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &forceList)
	if err != nil {
		return
	}
	return
}

//玩家活动神器战力排行入rs
func (dao *rankDao) SetRedisRankShenQiForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeShenQiForce, timestamp)

	bytes, err := json.Marshal(forceList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家活动命格战力排行ds
func (dao *rankDao) GetRankMingGeForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		err = dao.ds.DB().Raw(mingGeForceRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeLocalActivity:
		err = dao.ds.DB().Raw(mingGeForceRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeArea:
		err = dao.ds.DB().Raw(areaMingGeForceRankSql, ranktypes.ListLimit).Scan(&forceList).Error
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动命格战力排行rs
func (dao *rankDao) GetRedisRankMingGeForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeMingGeForce, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &forceList)
	if err != nil {
		return
	}
	return
}

//玩家活动命格战力排行入rs
func (dao *rankDao) SetRedisRankMingGeForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeMingGeForce, timestamp)

	bytes, err := json.Marshal(forceList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家活动圣痕战力排行ds
func (dao *rankDao) GetRankShengHenForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		err = dao.ds.DB().Raw(shengHenForceRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeLocalActivity:
		err = dao.ds.DB().Raw(shengHenForceRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeArea:
		err = dao.ds.DB().Raw(areaShengHenForceRankSql, ranktypes.ListLimit).Scan(&forceList).Error
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动圣痕战力排行rs
func (dao *rankDao) GetRedisRankShengHenForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeShengHenForce, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &forceList)
	if err != nil {
		return
	}
	return
}

//玩家活动圣痕战力排行入rs
func (dao *rankDao) SetRedisRankShengHenForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeShengHenForce, timestamp)

	bytes, err := json.Marshal(forceList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家活动阵法战力排行ds
func (dao *rankDao) GetRankZhenFaForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		err = dao.ds.DB().Raw(zhenFaForceRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeLocalActivity:
		err = dao.ds.DB().Raw(zhenFaForceRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeArea:
		err = dao.ds.DB().Raw(areaZhenFaForceRankSql, ranktypes.ListLimit).Scan(&forceList).Error
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动阵法战力排行rs
func (dao *rankDao) GetRedisRankZhenFaForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeZhenFaForce, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &forceList)
	if err != nil {
		return
	}
	return
}

//玩家活动阵法战力排行入rs
func (dao *rankDao) SetRedisRankZhenFaForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeZhenFaForce, timestamp)

	bytes, err := json.Marshal(forceList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家活动屠龙装战力排行ds
func (dao *rankDao) GetRankTuLongEquipForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		err = dao.ds.DB().Raw(tuLongEquipForceRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeLocalActivity:
		err = dao.ds.DB().Raw(tuLongEquipForceRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeArea:
		err = dao.ds.DB().Raw(areaTuLongEquipForceRankSql, ranktypes.ListLimit).Scan(&forceList).Error
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动屠龙装战力排行rs
func (dao *rankDao) GetRedisRankTuLongEquipForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeTuLongEquipForce, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &forceList)
	if err != nil {
		return
	}
	return
}

//玩家活动屠龙装战力排行入rs
func (dao *rankDao) SetRedisRankTuLongEquipForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeTuLongEquipForce, timestamp)

	bytes, err := json.Marshal(forceList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家活动宝宝战力排行ds
func (dao *rankDao) GetRankBabyForceList(config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		err = dao.ds.DB().Raw(babyForceRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeLocalActivity:
		err = dao.ds.DB().Raw(babyForceRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&forceList).Error
	case ranktypes.RankClassTypeArea:
		err = dao.ds.DB().Raw(areaBabyForceRankSql, ranktypes.ListLimit).Scan(&forceList).Error
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动宝宝战力排行rs
func (dao *rankDao) GetRedisRankBabyForceList(timestamp int64, config *ranktypes.RankConfig) (forceList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeBabyForce, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &forceList)
	if err != nil {
		return
	}
	return
}

//玩家活动宝宝战力排行入rs
func (dao *rankDao) SetRedisRankBabyForceList(timestamp int64, config *ranktypes.RankConfig, forceList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeBabyForce, timestamp)

	bytes, err := json.Marshal(forceList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家消费充值排行ds
func (dao *rankDao) GetRankCostList(config *ranktypes.RankConfig) (costList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(costRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&costList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(costRankActivitySql, global.GetGame().GetServerIndex(), config.GroupId, config.MinCondition, config.EndTime, ranktypes.ListLimit).Scan(&costList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaCostRankSql, ranktypes.ListLimit).Scan(&costList).Error
		}
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动消费排行rs
func (dao *rankDao) GetRedisRankCostList(timestamp int64, config *ranktypes.RankConfig) (costList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeCost, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &costList)
	if err != nil {
		return
	}
	return
}

//玩家活动消费排行入rs
func (dao *rankDao) SetRedisRankCostList(timestamp int64, config *ranktypes.RankConfig, costList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeCost, timestamp)

	bytes, err := json.Marshal(costList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家次数排行ds
func (dao *rankDao) GetRankCountList(config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(countRankSql, global.GetGame().GetServerIndex(), config.GroupId, config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(countRankActivitySql, global.GetGame().GetServerIndex(), config.GroupId, config.MinCondition, config.EndTime, ranktypes.ListLimit).Scan(&countList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaCountRankSql, global.GetGame().GetServerIndex(), config.GroupId, config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家活动次数排行rs
func (dao *rankDao) GetRedisRankCountList(timestamp int64, config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error) {
	//由于存已经用cost，所以拿也是用cost
	key := getRankKey(config, ranktypes.RankTypeCount, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &countList)
	if err != nil {
		return
	}
	return
}

//玩家活动次数排行入rs
func (dao *rankDao) SetRedisRankCountList(timestamp int64, config *ranktypes.RankConfig, countList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeCount, timestamp)

	bytes, err := json.Marshal(countList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家等级排行ds
func (dao *rankDao) GetRankLevelList(config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(levelRankSql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(levelRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaLevelRankSql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家等级排行rs
func (dao *rankDao) GetRedisRankLevelList(timestamp int64, config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeLevel, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &countList)
	if err != nil {
		return
	}
	return
}

//玩家等级排行入rs
func (dao *rankDao) SetRedisRankLevelList(timestamp int64, config *ranktypes.RankConfig, countList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeLevel, timestamp)

	bytes, err := json.Marshal(countList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家身法排行ds
func (dao *rankDao) GetRankShenFaList(config *ranktypes.RankConfig) (shenFaList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(shenFaRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&shenFaList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(shenFaRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&shenFaList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaShenFaRankSql, ranktypes.ListLimit).Scan(&shenFaList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家身法排行rs
func (dao *rankDao) GetRedisRankShenFaList(timestamp int64, config *ranktypes.RankConfig) (shenFaList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeShenFa, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &shenFaList)
	if err != nil {
		return
	}
	return
}

//玩家身法排行入rs
func (dao *rankDao) SetRedisRankShenFaList(timestamp int64, config *ranktypes.RankConfig, shenFaList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeShenFa, timestamp)

	bytes, err := json.Marshal(shenFaList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家领域排行ds
func (dao *rankDao) GetRankLingYuList(config *ranktypes.RankConfig) (lingYuList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(lingYuRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&lingYuList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(lingYuRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&lingYuList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaLingYuRankSql, ranktypes.ListLimit).Scan(&lingYuList).Error
		}
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家领域排行rs
func (dao *rankDao) GetRedisRankLingYuList(timestamp int64, config *ranktypes.RankConfig) (lingYuList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeLingYu, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &lingYuList)
	if err != nil {
		return
	}
	return
}

//玩家领域排行入rs
func (dao *rankDao) SetRedisRankLingYuList(timestamp int64, config *ranktypes.RankConfig, lingYuList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeLingYu, timestamp)

	bytes, err := json.Marshal(lingYuList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家护体仙羽排行ds
func (dao *rankDao) GetRankFeatherList(config *ranktypes.RankConfig) (featherList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(featherRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&featherList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(featherRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&featherList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaFeatherRankSql, ranktypes.ListLimit).Scan(&featherList).Error
		}
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家护体仙羽排行rs
func (dao *rankDao) GetRedisRankFeatherList(timestamp int64, config *ranktypes.RankConfig) (featherList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeFeather, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &featherList)
	if err != nil {
		return
	}
	return
}

//玩家护体仙羽排行入rs
func (dao *rankDao) SetRedisRankFeatherList(timestamp int64, config *ranktypes.RankConfig, featherList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeFeather, timestamp)

	bytes, err := json.Marshal(featherList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家神盾尖刺排行ds
func (dao *rankDao) GetRankShieldList(config *ranktypes.RankConfig) (shieldList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(shieldRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&shieldList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(shieldRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&shieldList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaShieldRankSql, ranktypes.ListLimit).Scan(&shieldList).Error
		}
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家神盾尖刺排行rs
func (dao *rankDao) GetRedisRankShieldList(timestamp int64, config *ranktypes.RankConfig) (shieldList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeShield, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &shieldList)
	if err != nil {
		return
	}
	return
}

//玩家神盾尖刺排行入rs
func (dao *rankDao) SetRedisRankShieldList(timestamp int64, config *ranktypes.RankConfig, shieldList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeShield, timestamp)

	bytes, err := json.Marshal(shieldList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家暗器排行ds
func (dao *rankDao) GetRankAnQiList(config *ranktypes.RankConfig) (anQiList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(anQiRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&anQiList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(anQiRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&anQiList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaAnQiRankSql, ranktypes.ListLimit).Scan(&anQiList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return

}

//玩家暗器排行rs
func (dao *rankDao) GetRedisRankAnQiList(timestamp int64, config *ranktypes.RankConfig) (anQiList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeAnQi, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &anQiList)
	if err != nil {
		return
	}
	return
}

//玩家暗器排行入rs
func (dao *rankDao) SetRedisRankAnQiList(timestamp int64, config *ranktypes.RankConfig, anQiList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeAnQi, timestamp)

	bytes, err := json.Marshal(anQiList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家魅力排行ds
func (dao *rankDao) GetRankCharmList(config *ranktypes.RankConfig) (charmList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(charmRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&charmList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(charmRankActivitySql, global.GetGame().GetServerIndex(), config.GroupId, marrytypes.MarryStatusTypeMarried, config.MinCondition, config.EndTime, ranktypes.ListLimit).Scan(&charmList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaCharmRankSql, ranktypes.ListLimit).Scan(&charmList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return

}

//玩家魅力排行rs
func (dao *rankDao) GetRedisRankCharmList(timestamp int64, config *ranktypes.RankConfig) (charmList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeCharm, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &charmList)
	if err != nil {
		return
	}
	return
}

//玩家魅力排行入rs
func (dao *rankDao) SetRedisRankCharmList(timestamp int64, config *ranktypes.RankConfig, charmList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeCharm, timestamp)

	bytes, err := json.Marshal(charmList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家表白排行ds
func (dao *rankDao) GetRankMarryDevelopList(config *ranktypes.RankConfig) (marryDevelopList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(marryDevelopRankSql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&marryDevelopList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(marryDevelopRankActivitySql, global.GetGame().GetServerIndex(), config.GroupId, config.MinCondition, config.EndTime, ranktypes.ListLimit).Scan(&marryDevelopList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaMarryDevelopRankSql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&marryDevelopList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return

}

//玩家表白排行rs
func (dao *rankDao) GetRedisRankMarryDevelopList(timestamp int64, config *ranktypes.RankConfig) (marryDevelopList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeMarryDevelop, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &marryDevelopList)
	if err != nil {
		return
	}
	return
}

//玩家表白排行入rs
func (dao *rankDao) SetRedisRankMarryDevelopList(timestamp int64, config *ranktypes.RankConfig, marryDevelopList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeMarryDevelop, timestamp)

	bytes, err := json.Marshal(marryDevelopList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家法宝排行ds
func (dao *rankDao) GetRankFaBaoList(config *ranktypes.RankConfig) (faBaoList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(faBaoRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&faBaoList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(faBaoRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&faBaoList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaFaBaoRankSql, ranktypes.ListLimit).Scan(&faBaoList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家法宝排行rs
func (dao *rankDao) GetRedisRankFaBaoList(timestamp int64, config *ranktypes.RankConfig) (faBaoList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeFaBao, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &faBaoList)
	if err != nil {
		return
	}
	return
}

//玩家法宝排行入rs
func (dao *rankDao) SetRedisRankFaBaoList(timestamp int64, config *ranktypes.RankConfig, faBaoList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeFaBao, timestamp)

	bytes, err := json.Marshal(faBaoList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家仙体排行ds
func (dao *rankDao) GetRankXianTiList(config *ranktypes.RankConfig) (xianTiList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(xianTiRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&xianTiList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(xianTiRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&xianTiList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaXianTiRankSql, ranktypes.ListLimit).Scan(&xianTiList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家仙体排行rs
func (dao *rankDao) GetRedisRankXianTiList(timestamp int64, config *ranktypes.RankConfig) (xianTiList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeXianTi, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &xianTiList)
	if err != nil {
		return
	}
	return
}

//玩家仙体排行入rs
func (dao *rankDao) SetRedisRankXianTiList(timestamp int64, config *ranktypes.RankConfig, xianTiList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeXianTi, timestamp)

	bytes, err := json.Marshal(xianTiList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家灵童养成类排行入ds
func (dao *rankDao) GetRankLingTongDevList(config *ranktypes.RankConfig) (lingTongDevList []*rankentity.PlayerOrderData, err error) {
	longTongDevType := config.RankType.GetLongTongDevType()
	sqlStr := getSqlByRankConfigAndClassType(config, longTongDevType)
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(sqlStr, longTongDevType, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&lingTongDevList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(sqlStr, longTongDevType, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&lingTongDevList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(sqlStr, longTongDevType, ranktypes.ListLimit).Scan(&lingTongDevList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家灵童养成类排行rs
func (dao *rankDao) GetRedisRankLingTongDevList(timestamp int64, config *ranktypes.RankConfig) (lingTongDevList []*rankentity.PlayerOrderData, err error) {
	rankType := config.RankType
	key := getRankKey(config, rankType, timestamp)

	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &lingTongDevList)
	if err != nil {
		return
	}
	return
}

//玩家灵童养成类排行入rs
func (dao *rankDao) SetRedisRankLingTongDevList(timestamp int64, config *ranktypes.RankConfig, lingTongDevList []*rankentity.PlayerOrderData) (err error) {
	rankType := config.RankType
	key := getRankKey(config, rankType, timestamp)
	bytes, err := json.Marshal(lingTongDevList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家噬魂幡排行ds
func (dao *rankDao) GetRankShiHunFanList(config *ranktypes.RankConfig) (shiHunFanList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(shiHunFanRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&shiHunFanList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(shiHunFanRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&shiHunFanList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaShiHunFanRankSql, ranktypes.ListLimit).Scan(&shiHunFanList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return

}

//玩家噬魂幡排行rs
func (dao *rankDao) GetRedisRankShiHunFanList(timestamp int64, config *ranktypes.RankConfig) (shiHunFanList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeShiHunFan, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &shiHunFanList)
	if err != nil {
		return
	}
	return
}

//玩家噬魂幡排行入rs
func (dao *rankDao) SetRedisRankShiHunFanList(timestamp int64, config *ranktypes.RankConfig, shiHunFanList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeShiHunFan, timestamp)

	bytes, err := json.Marshal(shiHunFanList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家天魔体排行ds
func (dao *rankDao) GetRankTianMoTiList(config *ranktypes.RankConfig) (tianMoTiList []*rankentity.PlayerOrderData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(tianmoTiRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&tianMoTiList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(tianMoTiRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&tianMoTiList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaTianMoTiRankSql, ranktypes.ListLimit).Scan(&tianMoTiList).Error
		}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家天魔体排行rs
func (dao *rankDao) GetRedisRankTianMoTiList(timestamp int64, config *ranktypes.RankConfig) (tianMoTiList []*rankentity.PlayerOrderData, err error) {
	key := getRankKey(config, ranktypes.RankTypeTianMoTi, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &tianMoTiList)
	if err != nil {
		return
	}
	return
}

//玩家天魔体排行入rs
func (dao *rankDao) SetRedisRankTianMoTiList(timestamp int64, config *ranktypes.RankConfig, tianMoTiList []*rankentity.PlayerOrderData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeTianMoTi, timestamp)

	bytes, err := json.Marshal(tianMoTiList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家灵童等级排行ds
func (dao *rankDao) GetRankLingTongLevelList(config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(lingTongLevelRankSql, global.GetGame().GetServerIndex(), ranktypes.ListLimit).Scan(&countList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaLingTongLevelRankSql, ranktypes.ListLimit).Scan(&countList).Error
		}
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家灵童等级排行rs
func (dao *rankDao) GetRedisRankLingTongLevelList(timestamp int64, config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeLingTongLevel, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &countList)
	if err != nil {
		return
	}
	return
}

//玩家灵童等级排行入rs
func (dao *rankDao) SetRedisRankLingTongLevelList(timestamp int64, config *ranktypes.RankConfig, countList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeLingTongLevel, timestamp)

	bytes, err := json.Marshal(countList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家飞升排行ds
func (dao *rankDao) GetRankFeiShengList(config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(feiShengRankSql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(feiShengRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaFeiShengRankSql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家飞升排行rs
func (dao *rankDao) GetRedisRankFeiShengList(timestamp int64, config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, ranktypes.RankTypeFeiSheng, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &countList)
	if err != nil {
		return
	}
	return
}

//玩家飞升排行入rs
func (dao *rankDao) SetRedisRankFeiShengList(timestamp int64, config *ranktypes.RankConfig, countList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, ranktypes.RankTypeFeiSheng, timestamp)

	bytes, err := json.Marshal(countList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//玩家转生排行ds
func (dao *rankDao) GetRankZhuanShengList(config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error) {
	switch config.ClassType {
	case ranktypes.RankClassTypeLocal:
		{
			err = dao.ds.DB().Raw(zhuanshengRankSql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	case ranktypes.RankClassTypeLocalActivity:
		{
			err = dao.ds.DB().Raw(zhuanshengRankActivitySql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	case ranktypes.RankClassTypeArea:
		{
			err = dao.ds.DB().Raw(areaZhuanshengRankSql, global.GetGame().GetServerIndex(), config.MinCondition, ranktypes.ListLimit).Scan(&countList).Error
		}
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//玩家转生排行rs
func (dao *rankDao) GetRedisRankZhuanShengList(timestamp int64, config *ranktypes.RankConfig) (countList []*rankentity.PlayerPropertyData, err error) {
	key := getRankKey(config, config.RankType, timestamp)
	dataStr, err := getRedisList(key)
	if err != nil || dataStr == "" {
		return
	}
	err = json.Unmarshal([]byte(dataStr), &countList)
	if err != nil {
		return
	}
	return
}

//玩家转生排行入rs
func (dao *rankDao) SetRedisRankZhuanShengList(timestamp int64, config *ranktypes.RankConfig, countList []*rankentity.PlayerPropertyData) (err error) {
	key := getRankKey(config, config.RankType, timestamp)

	bytes, err := json.Marshal(countList)
	if err != nil {
		return err
	}
	err = setRedisList(config.MaxExpireTime, key, bytes)
	if err != nil {
		return
	}
	return
}

//
const (
	dbName = "rank"
)

func (dao *rankDao) GetRankTimeEntity(serverId int32) (entity *rankentity.RankTimeEntity, err error) {
	entity = &rankentity.RankTimeEntity{}
	err = dao.ds.DB().First(entity, "serverId=? and  deleteTime=0", serverId).Error
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
	dao  *rankDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &rankDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetRankDao() RankDao {
	return dao
}
