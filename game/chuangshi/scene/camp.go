package scene

//阵营
type ChuangShiPkCamp int32

const (
	ChuangShiPkCampDefend ChuangShiPkCamp = iota + 1
	ChuangShiPkCampAttack
)

func (c ChuangShiPkCamp) Camp() int32 {
	return int32(c)
}
