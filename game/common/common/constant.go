package common

import (
	"time"
)

//最大等级
const MAX_LEVEL = 999

//最大转数
const MAX_ZHUAN = 999

//常量配置
//万分比
const MAX_RATE = 10000

//毫米
const MILL_METER = 1000

//毫秒
const TIME_RATE = 1000

//角度最小
const MIN_ANGLE = 0

//角度最大
const MAX_ANGLE = 360

//秒
const SECOND = time.Second / time.Millisecond

//分
const MINUTE = 60 * SECOND

//时
const HOUR = 60 * MINUTE

//24小时
const DAY = HOUR * 24

//误差 //浮点数据会有误差 不可能一直等于0
const MIN_DISTANCE_ERROR = 0.2

//误差平方
const MIN_DISTANCE_SQUARE_ERROR = MIN_DISTANCE_ERROR * MIN_DISTANCE_ERROR

//副本存活最高上限
const MAX_FUBEN_TIME = 2 * HOUR

//副本存活最低
const MIN_FUBEN_TIME = MINUTE

//副本失败最小时间
const MIN_FUBEN_FAILTURE_TIME = 30 * SECOND

//最大排行
const MAX_RANK = 100
