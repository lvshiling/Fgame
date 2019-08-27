package types

//权限
type AlliancePrivilege int64

const (
	//解散仙盟
	AlliancePrivilegeJieSan AlliancePrivilege = 1 << iota
	//逐出仙盟
	AlliancePrivilegeZhuChu
	//任命和撤销
	AlliancePrivilegeRenMing
	//一键招人
	AlliancePrivilegeZhaoRen
	//转让盟主
	AlliancePrivilegeZhuanRang
	//弹劾
	AlliancePrivilegeTanHe
	//退盟
	AlliancePiivilegeQuit
)

//职位
type AlliancePosition int32

const (
	//成员
	AlliancePositionMember = iota
	//堂主
	AlliancePositionTangZhu
	//副盟主
	AlliancePositionFuMengZhu
	//盟主
	AlliancePositionMengZhu
	//精英
	AlliancePositionJingYing
)

func (p AlliancePosition) Valid() bool {
	switch p {
	case AlliancePositionMember,
		AlliancePositionTangZhu,
		AlliancePositionFuMengZhu,
		AlliancePositionMengZhu,
		AlliancePositionJingYing:
		return true
	}
	return false
}

var (
	alliancePositionMap = map[AlliancePosition]string{
		AlliancePositionMember:    "普通成员",
		AlliancePositionTangZhu:   "堂主",
		AlliancePositionFuMengZhu: "副盟主",
		AlliancePositionMengZhu:   "盟主",
		AlliancePositionJingYing:  "精英",
	}
)

func (p AlliancePosition) String() string {
	return alliancePositionMap[p]
}

var (
	positionPrivilegeMap = map[AlliancePosition]int64{
		AlliancePositionMember: int64(AlliancePrivilegeTanHe) |
			int64(AlliancePiivilegeQuit),
		AlliancePositionTangZhu: int64(AlliancePrivilegeZhuChu) |
			int64(AlliancePrivilegeZhaoRen) |
			int64(AlliancePrivilegeTanHe) |
			int64(AlliancePiivilegeQuit),
		AlliancePositionFuMengZhu: int64(AlliancePrivilegeZhuChu) |
			int64(AlliancePrivilegeZhaoRen) |
			int64(AlliancePrivilegeTanHe) |
			int64(AlliancePiivilegeQuit),
		AlliancePositionMengZhu: int64(^uint(0) >> 1),
	}
)

func (p AlliancePosition) Privilege() int64 {
	return positionPrivilegeMap[p]
}

var (
	priorityMap = map[AlliancePosition]int32{
		AlliancePositionMember:    0,
		AlliancePositionTangZhu:   1,
		AlliancePositionFuMengZhu: 2,
		AlliancePositionMengZhu:   3,
	}
)

func (p AlliancePosition) Priority(otherPos AlliancePosition) bool {
	return priorityMap[p] > priorityMap[otherPos]
}

//仙术类型
type AllianceSkillType int32

const (
	//血量
	AllianceSkillTypeHP AllianceSkillType = iota + 1
	//攻击
	AllianceSkillTypeAttack
	//防御
	AllianceSkillTypeDefence
)

func (t AllianceSkillType) Valid() bool {
	switch t {
	case AllianceSkillTypeHP,
		AllianceSkillTypeAttack,
		AllianceSkillTypeDefence:
		return true
	}
	return false
}

type AllianceJuanXianType int32

const (
	//令牌捐献
	AllianceJuanXianTypeLingPai = iota + 1
	//银两捐献
	AllianceJuanXianTypeSilver
	//元宝捐献
	AllianceJuanXianTypeGold
)

func (t AllianceJuanXianType) Valid() bool {
	switch t {
	case AllianceJuanXianTypeLingPai,
		AllianceJuanXianTypeSilver,
		AllianceJuanXianTypeGold:
		return true
	}
	return false
}

var (
	juanXianTypeMap = map[AllianceJuanXianType]string{
		AllianceJuanXianTypeLingPai: "仙盟建设令",
		AllianceJuanXianTypeSilver:  "银两",
		AllianceJuanXianTypeGold:    "元宝",
	}
)

func (t AllianceJuanXianType) String() string {
	return juanXianTypeMap[t]
}

type AllianceSceneRewardType int32

const (
	AllianceSceneRewardTypeMember  AllianceSceneRewardType = 1
	AllianceSceneRewardTypeMengZhu                         = 2
)

func (t AllianceSceneRewardType) Valid() bool {
	switch t {
	case AllianceSceneRewardTypeMember,
		AllianceSceneRewardTypeMengZhu:
		return true
	}
	return false
}

//仙盟boss状态
type AllianceBossStatus int32

const (
	//待召唤
	AllianceBossStatusInit AllianceBossStatus = iota
	//已召唤
	AllianceBossStatusSummon
	//已击杀
	AllianceBossStatusDead
)

//仙盟召集类型
type AllianceCallType int32

const (
	AllianceCallTypeCommon AllianceCallType = iota //通用
	AllianceCallTypeYuXi                           //玉玺之战
)

func (t AllianceCallType) Valid() bool {
	switch t {
	case AllianceCallTypeCommon,
		AllianceCallTypeYuXi:
		return true
	}
	return false
}

// 公告常量
const (
	MinAllianceNoticeLen = 0
	MaxAllianceNoticeLen = 50
)

// 仙盟版本
type AllianceVersionType int32

const (
	AllianceVersionTypeOld AllianceVersionType = iota // 老版本仙盟
	AllianceVersionTypeNew                            // 新版本仙盟
)

func (a AllianceVersionType) Valid() bool {
	switch a {
	case AllianceVersionTypeOld,
		AllianceVersionTypeNew:
		return true
	default:
		return false
	}
}

// 新版本仙盟类型
type AllianceNewType int32

const (
	AllianceNewTypeLow  AllianceNewType = iota // 新版本低级仙盟
	AllianceNewTypeHigh                        // 新版本高级仙盟
)

func (a AllianceNewType) Valid() bool {
	switch a {
	case AllianceNewTypeLow,
		AllianceNewTypeHigh:
		return true
	default:
		return false
	}
}
