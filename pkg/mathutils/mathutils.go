package mathutils

import (
	"fmt"
	"math"
)

//转换成-180-180
func FixedAngle(angle float64) float64 {
	for angle < -180 {
		angle += 360
	}
	for angle > 180 {
		angle -= 360
	}

	return angle
}

//是否在2个角度内
func BetweenAngles(angle float64, minAngle float64, maxAngle float64) bool {
	if minAngle > maxAngle {
		panic(fmt.Errorf("min angle %.2f should no moren than max angle %.2f", minAngle, maxAngle))
	}
	for angle < minAngle {
		angle += 360
	}
	for angle > maxAngle {
		angle -= 360
	}

	if angle < minAngle {
		return false
	}
	if angle > maxAngle {
		return false
	}
	return true
}

func RadianToAngle(radian float64) float64 {
	return radian * 180 / math.Pi
}

func AngleToRadian(angle float64) float64 {
	return angle / 180 * math.Pi
}
