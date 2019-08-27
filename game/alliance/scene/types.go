package scene

//阵营
type AlliancePkCamp int32

const (
	AlliancePkCampDefend AlliancePkCamp = iota + 1
	AlliancePkCampAttack
)

func (c AlliancePkCamp) Camp() int32 {
	return int32(c)
}


