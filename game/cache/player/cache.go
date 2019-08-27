package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	additionsyscommon "fgame/fgame/game/additionsys/common"
	anqitypes "fgame/fgame/game/anqi/types"
	babytypes "fgame/fgame/game/baby/types"
	baguacommon "fgame/fgame/game/bagua/common"
	bodyshieldtypes "fgame/fgame/game/bodyshield/types"
	"fgame/fgame/game/cache/dao"
	cacheentity "fgame/fgame/game/cache/entity"
	dianxingcommon "fgame/fgame/game/dianxing/common"
	fabaocommon "fgame/fgame/game/fabao/common"
	"fgame/fgame/game/global"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	lingtongcommon "fgame/fgame/game/lingtong/common"
	lingtongdevcommon "fgame/fgame/game/lingtongdev/common"
	lingyutypes "fgame/fgame/game/lingyu/types"
	marrytypes "fgame/fgame/game/marry/types"
	massacretypes "fgame/fgame/game/massacre/types"
	mountcommon "fgame/fgame/game/mount/common"
	"fgame/fgame/game/player"
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
	"fgame/fgame/pkg/idutil"
	"fmt"

	"github.com/pkg/errors"
)

//缓存对象
type PlayerCacheObject struct {
	player             player.Player
	Id                 int64
	Name               string
	Role               playertypes.RoleType
	Sex                playertypes.SexType
	Level              int32
	Force              int64
	AllianceId         int64
	AllianceName       string
	TeamId             int64
	IsHuiYuan          int32
	BaseProperty       map[int32]int64
	BattleProperty     map[int32]int64
	EquipmentList      []*inventorytypes.EquipmentSlotInfo
	GoldEquipList      []*goldequiptypes.GoldEquipSlotInfo
	MountInfo          *mountcommon.MountInfo
	WingInfo           *wingcommon.WingInfo
	BodyShieldInfo     *bodyshieldtypes.BodyShieldInfo
	AnqiInfo           *anqitypes.AnqiInfo
	MassacreInfo       *massacretypes.MassacreInfo
	ShenfaInfo         *shenfatypes.ShenfaInfo
	LingyuInfo         *lingyutypes.LingyuInfo
	AllSoulInfo        *soultypes.AllSoulInfo
	AllWeaponInfo      *weapontypes.AllWeaponInfo
	ShieldInfo         *bodyshieldtypes.ShieldInfo
	FeatherInfo        *wingcommon.FeatherInfo
	MarryInfo          *marrytypes.MarryInfo
	FashionId          int32
	SkillList          []*skillcommon.SkillObjectImpl
	VipInfo            *viptypes.VipInfo
	FaBaoInfo          *fabaocommon.FaBaoInfo
	XueDunInfo         *xueduncommon.XueDunInfo
	XianTiInfo         *xianticommon.XianTiInfo
	BaGuaInfo          *baguacommon.BaGuaInfo
	DianXingInfo       *dianxingcommon.DianXingInfo
	TianMoTiInfo       *tianmotypes.TianMoInfo
	ShiHunFanInfo      *shihunfancommon.ShiHunFanInfo
	AllLingTongDevInfo *lingtongdevcommon.AllLingTongDevInfo
	LingTongInfo       *lingtongcommon.LingTongInfo
	AllSystemSkillInfo *sysskillcommon.AllSystemSkillInfo
	AllAdditionSysInfo *additionsyscommon.AllAdditionSysInfo
	PregnantInfo       *babytypes.PregnantInfo
	WushuangList       []*wushuangweapontypes.WushuangInfo
	XianZunCardList    []*xianzuncardcommon.XianZunCardInfo
	RingList           []*ringcommon.RingInfo
	RealmLevel         int32
	UpdateTime         int64
	CreateTime         int64
	DeleteTime         int64
}

func newPlayerCacheObject(pl player.Player) *PlayerCacheObject {
	pmo := &PlayerCacheObject{
		player: pl,
	}
	return pmo
}

func convertPlayerCacheObjectToEntity(o *PlayerCacheObject) (*cacheentity.PlayerCacheEntity, error) {
	baseProperty, err := json.Marshal(o.BaseProperty)
	if err != nil {
		return nil, err
	}
	battleProperty, err := json.Marshal(o.BattleProperty)
	if err != nil {
		return nil, err
	}

	equipmentList, err := json.Marshal(o.EquipmentList)
	if err != nil {
		return nil, err
	}
	goldEquipList, err := json.Marshal(o.GoldEquipList)
	if err != nil {
		return nil, err
	}
	mountInfo, err := json.Marshal(o.MountInfo)
	if err != nil {
		return nil, err
	}
	wingInfo, err := json.Marshal(o.WingInfo)
	if err != nil {
		return nil, err
	}
	bodyShieldInfo, err := json.Marshal(o.BodyShieldInfo)
	if err != nil {
		return nil, err
	}
	anqiInfo, err := json.Marshal(o.AnqiInfo)
	if err != nil {
		return nil, err
	}
	massacreInfo, err := json.Marshal(o.MassacreInfo)
	if err != nil {
		return nil, err
	}
	lingyuInfo, err := json.Marshal(o.LingyuInfo)
	if err != nil {
		return nil, err
	}
	shenfaInfo, err := json.Marshal(o.ShenfaInfo)
	if err != nil {
		return nil, err
	}
	allSoulInfo, err := json.Marshal(o.AllSoulInfo)
	if err != nil {
		return nil, err
	}
	allWeaponInfo, err := json.Marshal(o.AllWeaponInfo)
	if err != nil {
		return nil, err
	}
	shieldInfo, err := json.Marshal(o.ShieldInfo)
	if err != nil {
		return nil, err
	}
	featherInfo, err := json.Marshal(o.FeatherInfo)
	if err != nil {
		return nil, err
	}

	marryInfo, err := json.Marshal(o.MarryInfo)
	if err != nil {
		return nil, err
	}

	skillList, err := json.Marshal(o.SkillList)
	if err != nil {
		return nil, err
	}

	vipInfo, err := json.Marshal(o.VipInfo)
	if err != nil {
		return nil, err
	}

	faBaoInfo, err := json.Marshal(o.FaBaoInfo)
	if err != nil {
		return nil, err
	}

	xueDunInfo, err := json.Marshal(o.XueDunInfo)
	if err != nil {
		return nil, err
	}

	xianTiInfo, err := json.Marshal(o.XianTiInfo)
	if err != nil {
		return nil, err
	}

	baGuaInfo, err := json.Marshal(o.BaGuaInfo)
	if err != nil {
		return nil, err
	}

	dianXingInfo, err := json.Marshal(o.DianXingInfo)
	if err != nil {
		return nil, err
	}

	tianMoTiInfo, err := json.Marshal(o.TianMoTiInfo)
	if err != nil {
		return nil, err
	}

	shiHunFanInfo, err := json.Marshal(o.ShiHunFanInfo)
	if err != nil {
		return nil, err
	}

	allLingTongDevInfo, err := json.Marshal(o.AllLingTongDevInfo)
	if err != nil {
		return nil, err
	}

	lingTongInfo, err := json.Marshal(o.LingTongInfo)
	if err != nil {
		return nil, err
	}

	allSystemSkillInfo, err := json.Marshal(o.AllSystemSkillInfo)
	if err != nil {
		return nil, err
	}

	allAdditionSysInfo, err := json.Marshal(o.AllAdditionSysInfo)
	if err != nil {
		return nil, err
	}

	pregnantInfo, err := json.Marshal(o.PregnantInfo)
	if err != nil {
		return nil, err
	}

	wushuangInfo, err := json.Marshal(o.WushuangList)
	if err != nil {
		return nil, err
	}

	xianZunCardInfo, err := json.Marshal(o.XianZunCardList)
	if err != nil {
		return nil, err
	}

	ringInfo, err := json.Marshal(o.RingList)
	if err != nil {
		return nil, err
	}

	e := &cacheentity.PlayerCacheEntity{
		Id:                 o.Id,
		ServerId:           o.player.GetServerId(),
		PlayerId:           o.player.GetId(),
		Name:               o.Name,
		Role:               int32(o.Role),
		Sex:                int32(o.Sex),
		Level:              o.Level,
		Force:              o.Force,
		AllianceId:         o.AllianceId,
		AllianceName:       o.AllianceName,
		TeamId:             o.TeamId,
		IsHuiYuan:          o.IsHuiYuan,
		BaseProperty:       string(baseProperty),
		BattleProperty:     string(battleProperty),
		EquipmentList:      string(equipmentList),
		GoldEquipList:      string(goldEquipList),
		MountInfo:          string(mountInfo),
		WingInfo:           string(wingInfo),
		BodyShieldInfo:     string(bodyShieldInfo),
		AnqiInfo:           string(anqiInfo),
		MassacreInfo:       string(massacreInfo),
		LingyuInfo:         string(lingyuInfo),
		ShenfaInfo:         string(shenfaInfo),
		AllSoulInfo:        string(allSoulInfo),
		AllWeaponInfo:      string(allWeaponInfo),
		FashionId:          o.FashionId,
		ShieldInfo:         string(shieldInfo),
		FeatherInfo:        string(featherInfo),
		MarryInfo:          string(marryInfo),
		RealmLevel:         o.RealmLevel,
		SkillList:          string(skillList),
		VipInfo:            string(vipInfo),
		FaBaoInfo:          string(faBaoInfo),
		XueDunInfo:         string(xueDunInfo),
		XianTiInfo:         string(xianTiInfo),
		BaGuaInfo:          string(baGuaInfo),
		DianXingInfo:       string(dianXingInfo),
		TianMoTiInfo:       string(tianMoTiInfo),
		ShiHunFanInfo:      string(shiHunFanInfo),
		AllLingTongDevInfo: string(allLingTongDevInfo),
		LingTongInfo:       string(lingTongInfo),
		AllSystemSkillInfo: string(allSystemSkillInfo),
		AllAdditionSysInfo: string(allAdditionSysInfo),
		PregnantInfo:       string(pregnantInfo),
		WushuangListInfo:   string(wushuangInfo),
		XianZunCardInfo:    string(xianZunCardInfo),
		RingInfo:           string(ringInfo),
		UpdateTime:         o.UpdateTime,
		CreateTime:         o.CreateTime,
		DeleteTime:         o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerCacheObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerCacheObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerCacheObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerCacheObjectToEntity(o)
	return e, err
}

func (o *PlayerCacheObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*cacheentity.PlayerCacheEntity)
	o.Id = pe.Id
	o.Name = pe.Name
	o.Role = playertypes.RoleType(pe.Role)
	o.Sex = playertypes.SexType(pe.Sex)
	o.Level = pe.Level
	o.Force = pe.Force
	o.AllianceId = pe.AllianceId
	o.AllianceName = pe.AllianceName
	o.TeamId = pe.TeamId
	o.IsHuiYuan = pe.IsHuiYuan
	baseProperty := make(map[int32]int64)
	err = json.Unmarshal([]byte(pe.BaseProperty), &baseProperty)
	if err != nil {
		return err
	}
	o.BaseProperty = baseProperty
	battleProperty := make(map[int32]int64)
	err = json.Unmarshal([]byte(pe.BattleProperty), &battleProperty)
	if err != nil {
		return err
	}
	o.BattleProperty = battleProperty

	equipmentList := make([]*inventorytypes.EquipmentSlotInfo, 0, 8)
	err = json.Unmarshal([]byte(pe.EquipmentList), &equipmentList)
	if err != nil {
		return err
	}
	o.EquipmentList = equipmentList

	goldEquipList := make([]*goldequiptypes.GoldEquipSlotInfo, 0, 8)
	err = json.Unmarshal([]byte(pe.GoldEquipList), &goldEquipList)
	if err != nil {
		return err
	}
	o.GoldEquipList = goldEquipList

	wushuangListInfo := make([]*wushuangweapontypes.WushuangInfo, 0, 8)
	err = json.Unmarshal([]byte(pe.WushuangListInfo), &wushuangListInfo)
	if err != nil {
		return err
	}
	o.WushuangList = wushuangListInfo

	mountInfo := &mountcommon.MountInfo{}
	err = json.Unmarshal([]byte(pe.MountInfo), mountInfo)
	if err != nil {
		return err
	}
	o.MountInfo = mountInfo

	wingInfo := &wingcommon.WingInfo{}
	err = json.Unmarshal([]byte(pe.WingInfo), wingInfo)
	if err != nil {
		return err
	}
	o.WingInfo = wingInfo

	bodyShieldInfo := &bodyshieldtypes.BodyShieldInfo{}
	err = json.Unmarshal([]byte(pe.BodyShieldInfo), bodyShieldInfo)
	if err != nil {
		return err
	}
	o.BodyShieldInfo = bodyShieldInfo

	anqiInfo := &anqitypes.AnqiInfo{}
	err = json.Unmarshal([]byte(pe.AnqiInfo), anqiInfo)
	if err != nil {
		return err
	}
	o.AnqiInfo = anqiInfo

	massacreInfo := &massacretypes.MassacreInfo{}
	err = json.Unmarshal([]byte(pe.MassacreInfo), massacreInfo)
	if err != nil {
		return err
	}
	o.MassacreInfo = massacreInfo

	lingyuInfo := &lingyutypes.LingyuInfo{}
	err = json.Unmarshal([]byte(pe.LingyuInfo), lingyuInfo)
	if err != nil {
		return err
	}
	o.LingyuInfo = lingyuInfo

	shenfaInfo := &shenfatypes.ShenfaInfo{}
	err = json.Unmarshal([]byte(pe.ShenfaInfo), shenfaInfo)
	if err != nil {
		return err
	}
	o.ShenfaInfo = shenfaInfo

	allSoulInfo := &soultypes.AllSoulInfo{}
	err = json.Unmarshal([]byte(pe.AllSoulInfo), allSoulInfo)
	if err != nil {
		return err
	}
	o.AllSoulInfo = allSoulInfo

	allWeaponInfo := &weapontypes.AllWeaponInfo{}
	err = json.Unmarshal([]byte(pe.AllWeaponInfo), allWeaponInfo)
	if err != nil {
		return err
	}
	o.AllWeaponInfo = allWeaponInfo

	shieldInfo := &bodyshieldtypes.ShieldInfo{}
	err = json.Unmarshal([]byte(pe.ShieldInfo), shieldInfo)
	if err != nil {
		return err
	}
	o.ShieldInfo = shieldInfo

	featherInfo := &wingcommon.FeatherInfo{}
	err = json.Unmarshal([]byte(pe.FeatherInfo), featherInfo)
	if err != nil {
		return err
	}
	o.FeatherInfo = featherInfo

	marryInfo := &marrytypes.MarryInfo{}
	err = json.Unmarshal([]byte(pe.MarryInfo), marryInfo)
	if err != nil {
		return err
	}
	o.MarryInfo = marryInfo

	vipInfo := &viptypes.VipInfo{}
	err = json.Unmarshal([]byte(pe.VipInfo), vipInfo)
	if err != nil {
		return err
	}
	o.VipInfo = vipInfo

	skillList := make([]*skillcommon.SkillObjectImpl, 0, 16)
	err = json.Unmarshal([]byte(pe.SkillList), &skillList)
	if err != nil {
		return err
	}
	o.SkillList = skillList

	faBaoInfo := &fabaocommon.FaBaoInfo{}
	err = json.Unmarshal([]byte(pe.FaBaoInfo), faBaoInfo)
	if err != nil {
		return err
	}
	o.FaBaoInfo = faBaoInfo

	xueDunInfo := &xueduncommon.XueDunInfo{}
	err = json.Unmarshal([]byte(pe.XueDunInfo), xueDunInfo)
	if err != nil {
		return err
	}
	o.XueDunInfo = xueDunInfo

	xianTiInfo := &xianticommon.XianTiInfo{}
	err = json.Unmarshal([]byte(pe.XianTiInfo), xianTiInfo)
	if err != nil {
		return err
	}
	o.XianTiInfo = xianTiInfo

	baGuaInfo := &baguacommon.BaGuaInfo{}
	err = json.Unmarshal([]byte(pe.BaGuaInfo), baGuaInfo)
	if err != nil {
		return err
	}
	o.BaGuaInfo = baGuaInfo

	dianXingInfo := &dianxingcommon.DianXingInfo{}
	err = json.Unmarshal([]byte(pe.DianXingInfo), dianXingInfo)
	if err != nil {
		return err
	}
	o.DianXingInfo = dianXingInfo

	tianMoTiInfo := &tianmotypes.TianMoInfo{}
	err = json.Unmarshal([]byte(pe.TianMoTiInfo), tianMoTiInfo)
	if err != nil {
		return err
	}
	o.TianMoTiInfo = tianMoTiInfo

	shiHunFanInfo := &shihunfancommon.ShiHunFanInfo{}
	err = json.Unmarshal([]byte(pe.ShiHunFanInfo), shiHunFanInfo)
	if err != nil {
		return err
	}
	o.ShiHunFanInfo = shiHunFanInfo

	allLingTongDevInfo := &lingtongdevcommon.AllLingTongDevInfo{}
	err = json.Unmarshal([]byte(pe.AllLingTongDevInfo), allLingTongDevInfo)
	if err != nil {
		return err
	}
	o.AllLingTongDevInfo = allLingTongDevInfo

	lingTongInfo := &lingtongcommon.LingTongInfo{}
	err = json.Unmarshal([]byte(pe.LingTongInfo), lingTongInfo)
	if err != nil {
		return err
	}
	o.LingTongInfo = lingTongInfo

	allSystemSkillInfo := &sysskillcommon.AllSystemSkillInfo{}
	err = json.Unmarshal([]byte(pe.AllSystemSkillInfo), allSystemSkillInfo)
	if err != nil {
		return err
	}
	o.AllSystemSkillInfo = allSystemSkillInfo

	allAdditionSysInfo := &additionsyscommon.AllAdditionSysInfo{}
	err = json.Unmarshal([]byte(pe.AllAdditionSysInfo), allAdditionSysInfo)
	if err != nil {
		return err
	}
	o.AllAdditionSysInfo = allAdditionSysInfo

	pregnantInfo := &babytypes.PregnantInfo{}
	err = json.Unmarshal([]byte(pe.PregnantInfo), pregnantInfo)
	if err != nil {
		return err
	}
	o.PregnantInfo = pregnantInfo

	xianZunCardList := make([]*xianzuncardcommon.XianZunCardInfo, 0, 8)
	err = json.Unmarshal([]byte(pe.XianZunCardInfo), &xianZunCardList)
	if err != nil {
		return err
	}
	o.XianZunCardList = xianZunCardList

	ringList := make([]*ringcommon.RingInfo, 0, 8)
	err = json.Unmarshal([]byte(pe.RingInfo), &ringList)
	if err != nil {
		return err
	}
	o.RingList = ringList

	o.FashionId = pe.FashionId
	o.RealmLevel = pe.RealmLevel
	o.UpdateTime = pe.UpdateTime
	o.CreateTime = pe.CreateTime
	o.DeleteTime = pe.DeleteTime
	return
}

func (o *PlayerCacheObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Cache"))
	}
	obj, ok := e.(playertypes.PlayerDataEntity)
	if !ok {
		panic(fmt.Errorf("cache:强制转化玩家实体应该成功"))
	}

	o.player.AddChangedObject(obj)
	return
}

//玩家缓存管理器
type PlayerCacheDataManager struct {
	p player.Player
	//玩家缓存对象
	playerCacheObject *PlayerCacheObject
}

func (m *PlayerCacheDataManager) Player() player.Player {
	return m.p
}

//同步缓存
func (m *PlayerCacheDataManager) SyncCache() {
	now := global.GetGame().GetTimeService().Now()
	if m.playerCacheObject == nil {
		m.playerCacheObject = newPlayerCacheObject(m.p)
		m.playerCacheObject.Id, _ = idutil.GetId()
		m.playerCacheObject.CreateTime = now
	} else {
		m.playerCacheObject.UpdateTime = now
	}
	m.playerCacheObject.Name = m.p.GetName()
	m.playerCacheObject.Role = m.p.GetRole()
	m.playerCacheObject.Sex = m.p.GetSex()
	m.playerCacheObject.Level = m.p.GetLevel()
	m.playerCacheObject.Force = m.p.GetForce()
	m.playerCacheObject.AllianceId = m.p.GetAllianceId()
	m.playerCacheObject.AllianceName = m.p.GetAllianceName()
	m.playerCacheObject.TeamId = m.p.GetTeamId()
	m.playerCacheObject.BaseProperty = m.p.GetBaseProperties()
	m.playerCacheObject.BattleProperty = m.p.GetBattleProperties()
	m.playerCacheObject.BodyShieldInfo = m.p.GetBodyshieldInfo()
	m.playerCacheObject.AnqiInfo = m.p.GetAnqiInfo()
	m.playerCacheObject.MassacreInfo = m.p.GetMassacreInfo()
	m.playerCacheObject.LingyuInfo = m.p.GetLingyuInfo()
	m.playerCacheObject.ShenfaInfo = m.p.GetShenfaInfo()
	m.playerCacheObject.MountInfo = m.p.GetMountInfo()
	m.playerCacheObject.WingInfo = m.p.GetWingInfo()
	m.playerCacheObject.EquipmentList = m.p.GetEquipmentSlotList()
	m.playerCacheObject.GoldEquipList = m.p.GetGoldEquipSlotList()
	m.playerCacheObject.WushuangList = m.p.GetWushuangListInfo()
	m.playerCacheObject.AllSoulInfo = m.p.GetAllSoulInfo()
	m.playerCacheObject.AllWeaponInfo = m.p.GetAllWeaponInfo()
	m.playerCacheObject.ShieldInfo = m.p.GetShieldInfo()
	m.playerCacheObject.FeatherInfo = m.p.GetFeatherInfo()
	m.playerCacheObject.MarryInfo = m.p.GetMarryInfo()
	m.playerCacheObject.VipInfo = m.p.GetVipInfo()
	m.playerCacheObject.FashionId = m.p.GetFashionId()
	m.playerCacheObject.RealmLevel = m.p.GetRealm()
	m.playerCacheObject.SkillList = m.p.GetSkillList()
	m.playerCacheObject.FaBaoInfo = m.p.GetFaBaoInfo()
	m.playerCacheObject.XueDunInfo = m.p.GetXueDunInfo()
	m.playerCacheObject.XianTiInfo = m.p.GetXianTiInfo()
	m.playerCacheObject.BaGuaInfo = m.p.GetBaGuaInfo()
	m.playerCacheObject.DianXingInfo = m.p.GetDianXingInfo()
	m.playerCacheObject.TianMoTiInfo = m.p.GetTianMoTiInfo()
	m.playerCacheObject.ShiHunFanInfo = m.p.GetShiHunFanInfo()
	m.playerCacheObject.AllLingTongDevInfo = m.p.GetAllLingTongDevInfo()
	m.playerCacheObject.LingTongInfo = m.p.GetLingTongInfo()
	m.playerCacheObject.AllSystemSkillInfo = m.p.GetAllSystemSkillInfo()
	m.playerCacheObject.AllAdditionSysInfo = m.p.GetAllAdditionSysInfo()
	m.playerCacheObject.PregnantInfo = m.p.GetPregnantInfo()
	m.playerCacheObject.XianZunCardList = m.p.GetXianZunCard()
	m.playerCacheObject.RingList = m.p.GetRingInfo()
	if m.p.IsHuiYuanPlus() {
		m.playerCacheObject.IsHuiYuan = 1
	} else {
		m.playerCacheObject.IsHuiYuan = 0
	}
	m.playerCacheObject.SetModified()

	return
}

//加载
func (m *PlayerCacheDataManager) Load() (err error) {
	cacheEntity, err := dao.GetCacheDao().GetPlayerCacheByPlayerId(m.p.GetId())
	if err != nil {
		return
	}
	if cacheEntity != nil {
		m.playerCacheObject = newPlayerCacheObject(m.p)
		err = m.playerCacheObject.FromEntity(cacheEntity)
		if err != nil {
			return
		}
	}
	return nil
}

//加载后
func (m *PlayerCacheDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (m *PlayerCacheDataManager) Heartbeat() {

}

func CreatePlayerCacheDataManager(p player.Player) player.PlayerDataManager {
	pmdm := &PlayerCacheDataManager{}
	pmdm.p = p
	return pmdm
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerCacheDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerCacheDataManager))
}
