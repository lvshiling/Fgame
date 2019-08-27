package cache

import (
	"encoding/json"
	additionsyscommon "fgame/fgame/game/additionsys/common"
	anqitypes "fgame/fgame/game/anqi/types"
	babytypes "fgame/fgame/game/baby/types"
	baguacommon "fgame/fgame/game/bagua/common"
	bodyshieldtypes "fgame/fgame/game/bodyshield/types"
	"fgame/fgame/game/cache/dao"
	cacheentity "fgame/fgame/game/cache/entity"
	dianxingcommon "fgame/fgame/game/dianxing/common"
	fabaocommon "fgame/fgame/game/fabao/common"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	lingtongcommon "fgame/fgame/game/lingtong/common"
	lingtongdevcommon "fgame/fgame/game/lingtongdev/common"
	lingyutypes "fgame/fgame/game/lingyu/types"
	marrytypes "fgame/fgame/game/marry/types"
	massacretypes "fgame/fgame/game/massacre/types"
	mountcommon "fgame/fgame/game/mount/common"
	playercommon "fgame/fgame/game/player/common"
	playertypes "fgame/fgame/game/player/types"
	ringcommon "fgame/fgame/game/ring/common"
	shenfatypes "fgame/fgame/game/shenfa/types"
	shihunfancommon "fgame/fgame/game/shihunfan/common"
	skillcommon "fgame/fgame/game/skill/common"
	soultypes "fgame/fgame/game/soul/types"
	sysskillcommon "fgame/fgame/game/systemskill/common"
	tianmotypes "fgame/fgame/game/tianmo/types"
	viptypes "fgame/fgame/game/vip/types"
	weapontypes "fgame/fgame/game/weapon/types"
	wingcommon "fgame/fgame/game/wing/common"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"
	xianticommon "fgame/fgame/game/xianti/common"
	xianzuncardcommon "fgame/fgame/game/xianzuncard/common"
	xueduncommon "fgame/fgame/game/xuedun/common"
	"sync"

	lru "github.com/hashicorp/golang-lru"
)

type CacheService interface {
	GetPlayerInfoByPlayerId(playerId int64) (info *playercommon.PlayerInfo, err error)
	GetPlayerInfoByName(name string, serverId int32) (info *playercommon.PlayerInfo, err error)
	UpdateCache(p *playercommon.PlayerInfo)
}

type cacheService struct {
	m            sync.Mutex
	lruCache     *lru.TwoQueueCache
	nameLruCache *lru.TwoQueueCache
}

var cacheSize = 500

func (s *cacheService) init() (err error) {
	s.lruCache, err = lru.New2Q(cacheSize)
	if err != nil {
		return
	}
	s.nameLruCache, err = lru.New2Q(cacheSize)
	if err != nil {
		return
	}
	return nil
}

func (s *cacheService) GetPlayerInfoByPlayerId(playerId int64) (info *playercommon.PlayerInfo, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	infoInter, ok := s.lruCache.Get(playerId)
	if ok {
		info, ok = infoInter.(*playercommon.PlayerInfo)
		if ok {
			return
		}
	}
	e, err := dao.GetCacheDao().GetPlayerCacheByPlayerId(playerId)
	if err != nil {
		return
	}
	if e == nil {
		return
	}
	info, err = convertFromCache(e)
	if err != nil {
		return
	}
	s.lruCache.Add(playerId, info)
	return
}

func (s *cacheService) GetPlayerInfoByName(name string, serverId int32) (info *playercommon.PlayerInfo, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	infoInter, ok := s.nameLruCache.Get(name)
	if ok {
		info, ok = infoInter.(*playercommon.PlayerInfo)
		if ok {
			return
		}
	}
	e, err := dao.GetCacheDao().GetPlayerCacheByName(name, serverId)
	if err != nil {
		return
	}
	if e == nil {
		return
	}
	info, err = convertFromCache(e)
	if err != nil {
		return
	}
	s.nameLruCache.Add(name, info)
	return
}

func (s *cacheService) UpdateCache(info *playercommon.PlayerInfo) {
	s.m.Lock()
	defer s.m.Unlock()
	s.lruCache.Add(info.PlayerId, info)
	s.nameLruCache.Add(info.Name, info)
	return
}

func convertFromCache(pe *cacheentity.PlayerCacheEntity) (info *playercommon.PlayerInfo, err error) {
	info = &playercommon.PlayerInfo{
		ServerId:     pe.ServerId,
		PlayerId:     pe.PlayerId,
		Name:         pe.Name,
		Role:         playertypes.RoleType(pe.Role),
		Sex:          playertypes.SexType(pe.Sex),
		Level:        pe.Level,
		Force:        pe.Force,
		AllianceId:   pe.AllianceId,
		AllianceName: pe.AllianceName,
		TeamId:       pe.TeamId,
		IsHuiYuan:    pe.IsHuiYuan,
	}
	baseProperty := make(map[int32]int64)
	err = json.Unmarshal([]byte(pe.BaseProperty), &baseProperty)
	if err != nil {
		return nil, err
	}
	info.BaseProperty = baseProperty
	battleProperty := make(map[int32]int64)
	err = json.Unmarshal([]byte(pe.BattleProperty), &battleProperty)
	if err != nil {
		return nil, err
	}
	info.BattleProperty = battleProperty

	equipmentList := make([]*inventorytypes.EquipmentSlotInfo, 0, 8)
	err = json.Unmarshal([]byte(pe.EquipmentList), &equipmentList)
	if err != nil {
		return nil, err
	}
	info.EquipmentList = equipmentList

	goldEquipList := make([]*goldequiptypes.GoldEquipSlotInfo, 0, 8)
	err = json.Unmarshal([]byte(pe.GoldEquipList), &goldEquipList)
	if err != nil {
		return nil, err
	}
	info.GoldEquipList = goldEquipList

	wushuangList := make([]*wushuangweapontypes.WushuangInfo, 0, 8)
	err = json.Unmarshal([]byte(pe.WushuangListInfo), &wushuangList)
	if err != nil {
		return nil, err
	}
	info.WushuangList = wushuangList
	//TODO: xzk 临时处理bug
	for _, goldEquip := range info.GoldEquipList {
		if goldEquip.PropertyData == nil {
			goldEquip.PropertyData = goldequiptypes.NewGoldEquipPropertyData()
		}
		goldEquip.PropertyData.InitBase()
	}

	xianZunCardList := make([]*xianzuncardcommon.XianZunCardInfo, 0, 8)
	err = json.Unmarshal([]byte(pe.XianZunCardInfo), &xianZunCardList)
	if err != nil {
		return nil, err
	}
	info.XianZunCardList = xianZunCardList

	ringList := make([]*ringcommon.RingInfo, 0, 8)
	err = json.Unmarshal([]byte(pe.RingInfo), &ringList)
	if err != nil {
		return nil, err
	}
	info.RingList = ringList

	mountInfo := &mountcommon.MountInfo{}
	err = json.Unmarshal([]byte(pe.MountInfo), mountInfo)
	if err != nil {
		return nil, err
	}
	info.MountInfo = mountInfo

	wingInfo := &wingcommon.WingInfo{}
	err = json.Unmarshal([]byte(pe.WingInfo), wingInfo)
	if err != nil {
		return nil, err
	}
	info.WingInfo = wingInfo

	bodyShieldInfo := &bodyshieldtypes.BodyShieldInfo{}
	err = json.Unmarshal([]byte(pe.BodyShieldInfo), bodyShieldInfo)
	if err != nil {
		return nil, err
	}
	info.BodyShieldInfo = bodyShieldInfo

	anqiInfo := &anqitypes.AnqiInfo{}
	err = json.Unmarshal([]byte(pe.AnqiInfo), anqiInfo)
	if err != nil {
		return nil, err
	}
	info.AnqiInfo = anqiInfo

	massacreInfo := &massacretypes.MassacreInfo{}
	err = json.Unmarshal([]byte(pe.MassacreInfo), massacreInfo)
	if err != nil {
		return nil, err
	}
	info.MassacreInfo = massacreInfo

	lingyuInfo := &lingyutypes.LingyuInfo{}
	err = json.Unmarshal([]byte(pe.LingyuInfo), lingyuInfo)
	if err != nil {
		return nil, err
	}
	info.LingyuInfo = lingyuInfo

	shenfaInfo := &shenfatypes.ShenfaInfo{}
	err = json.Unmarshal([]byte(pe.ShenfaInfo), shenfaInfo)
	if err != nil {
		return nil, err
	}
	info.ShenfaInfo = shenfaInfo

	allSoulInfo := &soultypes.AllSoulInfo{}
	err = json.Unmarshal([]byte(pe.AllSoulInfo), allSoulInfo)
	if err != nil {
		return nil, err
	}
	info.AllSoulInfo = allSoulInfo

	allWeaponInfo := &weapontypes.AllWeaponInfo{}
	err = json.Unmarshal([]byte(pe.AllWeaponInfo), allWeaponInfo)
	if err != nil {
		return nil, err
	}
	info.AllWeaponInfo = allWeaponInfo
	info.FashionId = pe.FashionId
	info.RealmLevel = pe.RealmLevel

	shieldInfo := &bodyshieldtypes.ShieldInfo{}
	err = json.Unmarshal([]byte(pe.ShieldInfo), shieldInfo)
	if err != nil {
		return nil, err
	}
	info.ShieldInfo = shieldInfo

	featherInfo := &wingcommon.FeatherInfo{}
	err = json.Unmarshal([]byte(pe.FeatherInfo), featherInfo)
	if err != nil {
		return nil, err
	}
	info.FeatherInfo = featherInfo

	marryInfo := &marrytypes.MarryInfo{}
	err = json.Unmarshal([]byte(pe.MarryInfo), marryInfo)
	if err != nil {
		return nil, err
	}
	info.MarryInfo = marryInfo

	skillList := make([]*skillcommon.SkillObjectImpl, 0, 8)
	err = json.Unmarshal([]byte(pe.SkillList), &skillList)
	if err != nil {
		return nil, err
	}
	info.SkillList = skillList

	vipInfo := &viptypes.VipInfo{}
	err = json.Unmarshal([]byte(pe.VipInfo), vipInfo)
	if err != nil {
		return nil, err
	}
	info.VipInfo = vipInfo

	faBaoInfo := &fabaocommon.FaBaoInfo{}
	err = json.Unmarshal([]byte(pe.FaBaoInfo), faBaoInfo)
	if err != nil {
		return nil, err
	}
	info.FaBaoInfo = faBaoInfo

	xueDunInfo := &xueduncommon.XueDunInfo{}
	err = json.Unmarshal([]byte(pe.XueDunInfo), xueDunInfo)
	if err != nil {
		return nil, err
	}
	info.XueDunInfo = xueDunInfo

	xianTiInfo := &xianticommon.XianTiInfo{}
	err = json.Unmarshal([]byte(pe.XianTiInfo), xianTiInfo)
	if err != nil {
		return nil, err
	}
	info.XianTiInfo = xianTiInfo

	baGuaInfo := &baguacommon.BaGuaInfo{}
	err = json.Unmarshal([]byte(pe.BaGuaInfo), baGuaInfo)
	if err != nil {
		return nil, err
	}
	info.BaGuaInfo = baGuaInfo

	dianXingInfo := &dianxingcommon.DianXingInfo{}
	err = json.Unmarshal([]byte(pe.DianXingInfo), dianXingInfo)
	if err != nil {
		return nil, err
	}
	info.DianXingInfo = dianXingInfo

	tianMoTiInfo := &tianmotypes.TianMoInfo{}
	err = json.Unmarshal([]byte(pe.TianMoTiInfo), tianMoTiInfo)
	if err != nil {
		return nil, err
	}
	info.TianMoTiInfo = tianMoTiInfo

	shiHunFanInfo := &shihunfancommon.ShiHunFanInfo{}
	err = json.Unmarshal([]byte(pe.ShiHunFanInfo), shiHunFanInfo)
	if err != nil {
		return nil, err
	}
	info.ShiHunFanInfo = shiHunFanInfo

	allLingTongDevInfo := &lingtongdevcommon.AllLingTongDevInfo{}
	err = json.Unmarshal([]byte(pe.AllLingTongDevInfo), allLingTongDevInfo)
	if err != nil {
		return nil, err
	}
	info.AllLingTongDevInfo = allLingTongDevInfo

	lingTongInfo := &lingtongcommon.LingTongInfo{}
	err = json.Unmarshal([]byte(pe.LingTongInfo), lingTongInfo)
	if err != nil {
		return nil, err
	}
	info.LingTongInfo = lingTongInfo

	allSystemSkillInfo := &sysskillcommon.AllSystemSkillInfo{}
	err = json.Unmarshal([]byte(pe.AllSystemSkillInfo), allSystemSkillInfo)
	if err != nil {
		return nil, err
	}
	info.AllSystemSkillInfo = allSystemSkillInfo

	allAdditionSysInfo := &additionsyscommon.AllAdditionSysInfo{}
	err = json.Unmarshal([]byte(pe.AllAdditionSysInfo), allAdditionSysInfo)
	if err != nil {
		return nil, err
	}
	info.AllAdditionSysInfo = allAdditionSysInfo

	pregnantInfo := &babytypes.PregnantInfo{}
	err = json.Unmarshal([]byte(pe.PregnantInfo), pregnantInfo)
	if err != nil {
		return nil, err
	}
	info.PregnantInfo = pregnantInfo

	return
}

var (
	once sync.Once
	s    *cacheService
)

func Init() (err error) {
	once.Do(func() {
		s = &cacheService{}
		err = s.init()
	})
	return
}

func GetCacheService() CacheService {
	return s
}
