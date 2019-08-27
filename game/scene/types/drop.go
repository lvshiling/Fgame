package types

type DropType int32

const (
	//死亡掉落
	DropTypeAfterDead DropType = iota
	//掉落百分比
	DropTypePercent
	//击杀后物品直接入包,包满则邮件发送给玩家
	DropTypeIntoBag
)

func (dt DropType) Valid() bool {
	switch dt {
	case DropTypeAfterDead,
		DropTypePercent,
		DropTypeIntoBag:
		return true
	}
	return false
}

//掉落判定
type DropJudgeType int32

const (
	// 击杀者或击杀者的队伍获得奖励
	DropJudgeTypeKillerOrTeam DropJudgeType = iota + 1
	//开怪者或开怪者的队伍获得奖励
	DropJudgeTypeOpenerOrTeam
	//累计伤害最高者或累计伤害最高的队伍获得奖励
	DropJudgeTypeMaxHurtOrTeam
	//击杀者获得奖励
	DropJudgeTypeKiller
	//开怪者获得奖励
	DropJudgeTypeOpener
	//累计伤害最高者获得奖励
	DropJudgeTypeMaxHurt
)

func (dt DropJudgeType) Valid() bool {
	switch dt {
	case DropJudgeTypeKillerOrTeam,
		DropJudgeTypeOpenerOrTeam,
		DropJudgeTypeMaxHurtOrTeam,
		DropJudgeTypeKiller,
		DropJudgeTypeOpener,
		DropJudgeTypeMaxHurt:
		return true
	}
	return false
}

//掉落归属
type DropOwnerType int32

const (
	//玩家
	DropOwnerTypePlayer DropOwnerType = iota
	//队伍
	DropOwnerTypeTeam
	//仙盟
	DropOwnerTypeAlliance
)
