package entity

//玩家缓存数据
type PlayerCacheEntity struct {
	Id                 int64  `gorm:"primary_key;column:id"`
	ServerId           int32  `gorm:"column:serverId"`
	PlayerId           int64  `gorm:"column:playerId"`
	Name               string `gorm:"column:name"`
	Role               int32  `gorm:"column:role"`
	Sex                int32  `gorm:"column:sex"`
	Level              int32  `gorm:"column:level"`
	Force              int64  `gorm:"column:force"`
	AllianceId         int64  `gorm:"column:allianceId"`
	IsHuiYuan          int32  `gorm:"column:isHuiYuan"`
	AllianceName       string `gorm:"column:allianceName"`
	TeamId             int64  `gorm:"column:teamId"`
	BaseProperty       string `gorm:"column:baseProperty"`
	BattleProperty     string `gorm:"column:battleProperty"`
	EquipmentList      string `gorm:"column:equipmentList"`
	GoldEquipList      string `gorm:"column:goldEquipList"`
	MountInfo          string `gorm:"column:mountInfo"`
	WingInfo           string `gorm:"column:wingInfo"`
	BodyShieldInfo     string `gorm:"column:bodyShieldInfo"`
	AnqiInfo           string `gorm:"column:anqiInfo"`
	MassacreInfo       string `gorm:"column:massacreInfo"`
	LingyuInfo         string `gorm:"column:lingyuInfo"`
	ShenfaInfo         string `gorm:"column:shenfaInfo"`
	AllSoulInfo        string `gorm:"column:allSoulInfo"`
	AllWeaponInfo      string `gorm:"column:allWeaponInfo"`
	FashionId          int32  `gorm:"column:fashionId"`
	ShieldInfo         string `gorm:"column:shieldInfo"`
	FeatherInfo        string `gorm:"column:featherInfo"`
	MarryInfo          string `gorm:"column:marryInfo"`
	RealmLevel         int32  `gorm:"column:realmLevel"`
	SkillList          string `gorm:"column:skillList"`
	VipInfo            string `gorm:"column:vipInfo"`
	FaBaoInfo          string `gorm:"column:fabaoInfo"`
	XueDunInfo         string `gorm:"column:xuedunInfo"`
	XianTiInfo         string `gorm:"column:xiantiInfo"`
	BaGuaInfo          string `gorm:"column:baguaInfo"`
	DianXingInfo       string `gorm:"column:dianxingInfo"`
	TianMoTiInfo       string `gorm:"column:tianMoTiInfo"`
	ShiHunFanInfo      string `gorm:"column:shihunfanInfo"`
	AllLingTongDevInfo string `gorm:"column:allLingTongDevInfo"`
	LingTongInfo       string `gorm:"column:lingTongInfo"`
	AllSystemSkillInfo string `gorm:"column:allSystemSkillInfo"`
	AllAdditionSysInfo string `gorm:"column:allAdditionSysInfo"`
	PregnantInfo       string `gorm:"column:pregnantInfo"`
	WushuangListInfo   string `gorm:"column:wushuangListInfo"`
	XianZunCardInfo    string `gorm:"column:xianZunCardInfo"`
	RingInfo           string `gorm:"column:ringInfo"`
	UpdateTime         int64  `gorm:"column:updateTime"`
	CreateTime         int64  `gorm:"column:createTime"`
	DeleteTime         int64  `gorm:"column:deleteTime"`
}

func (e *PlayerCacheEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerCacheEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerCacheEntity) TableName() string {
	return "t_player_cache"
}
