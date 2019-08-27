package validator

import "fmt"

//最小值验证
func MinValidate(val float64, min float64, minInclude bool) (err error) {
	if minInclude {
		if val < min {
			return fmt.Errorf("%.2f  less than %.2f", val, min)
		}
	} else {
		if val <= min {
			return fmt.Errorf("%.2f no more than %.2f", val, min)
		}
	}
	return nil
}

//最大值验证
func MaxValidate(val float64, max float64, maxInclude bool) (err error) {
	if maxInclude {
		if val > max {
			return fmt.Errorf("%.2f more than  %.2f", val, max)
		}
	} else {
		if val >= max {
			return fmt.Errorf("%.2f no less than %.2f", val, max)
		}
	}
	return nil
}

//范围验证
func RangeValidate(val float64, min float64, minInclude bool, max float64, maxInclude bool) (err error) {
	if err = MinValidate(val, min, minInclude); err != nil {
		return
	}
	if err = MaxValidate(val, max, maxInclude); err != nil {
		return
	}
	return nil
}
