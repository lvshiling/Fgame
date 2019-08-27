package server

import (
	"time"
)

//metric统计参考redis
const (
	//metric样品数
	statsMetricSample = 16
	//metric类型数
	statsMetricCount = 4
	//metric时间 (毫秒)
	statsMetricCronTime = 100
)

type metricType int

const (
	metricTypeMsgIn metricType = iota
	metricTypeMsgBytesIn
	metricTypeMsgOut
	metricTypeMsgBytesOut
)

type metric struct {
	lastSampleTime  int64
	lastSampleCount int64
	idx             int
	samples         [statsMetricSample]int64
}

type stats struct {
	msg        *msgStats
	conn       *connStats
	instMetric [statsMetricCount]*metric
	done       chan struct{}
	t          *time.Ticker
}

func newStats() *stats {
	s := &stats{
		msg:  &msgStats{},
		conn: &connStats{},
	}

	for i := metricTypeMsgIn; i <= metricTypeMsgBytesOut; i++ {
		s.instMetric[i] = &metric{}
	}

	s.done = make(chan struct{})
	s.t = time.NewTicker(time.Millisecond * statsMetricCronTime)
	s.Start()
	return s
}

func (s *stats) Start() {
	go func() {
	Loop:
		for {
			select {
			case <-s.t.C:
				{
					s.trackInstantaneousMetric(metricTypeMsgIn, s.msg.InMsgs)
					s.trackInstantaneousMetric(metricTypeMsgOut, s.msg.OutMsgs)
					s.trackInstantaneousMetric(metricTypeMsgBytesIn, s.msg.InBytes)
					s.trackInstantaneousMetric(metricTypeMsgBytesOut, s.msg.OutBytes)
				}
			case <-s.done:
				{
					break Loop
				}
			}

		}
	}()
}

//停止
func (s *stats) Stop() {
	s.t.Stop()
	close(s.done)
}

func (s *stats) trackInstantaneousMetric(mt metricType, current int64) {

	now := time.Now().UnixNano() / int64(time.Millisecond)
	elapse := now - s.instMetric[mt].lastSampleTime
	//TODO elapse可能为0很奇怪
	if elapse <= 0 {
		return
	}
	ops := current - s.instMetric[mt].lastSampleCount
	opsSec := ops * 1000 / elapse
	s.instMetric[mt].samples[s.instMetric[mt].idx] = opsSec
	s.instMetric[mt].idx++
	s.instMetric[mt].idx %= statsMetricSample
	s.instMetric[mt].lastSampleTime = now
	s.instMetric[mt].lastSampleCount = current
}

//统计连接数
type connStats struct {
	TotalConnections   int64 `json:"totalConnections"`
	CurrentConnections int64 `json:"currentConnections"`
}

//消息数
type msgStats struct {
	InMsgs   int64 `json:"inMsgs"`
	InBytes  int64 `json:"inBytes"`
	OutMsgs  int64 `json:"outMsgs"`
	OutBytes int64 `json:"outBytes"`
}

type statsResult struct {
	InMsgs             int64 `json:"inMsgs"`
	InBytes            int64 `json:"inBytes"`
	OutMsgs            int64 `json:"outMsgs"`
	OutBytes           int64 `json:"outBytes"`
	TotalConnections   int64 `json:"totalConnections"`
	CurrentConnections int64 `json:"currentConnections"`
	InMsgsMetric       int64 `json:"inMsgsMetric"`
	InBytesMetric      int64 `json:"inBytesMetric"`
	OutMsgsMetric      int64 `json:"outMsgsMetric"`
	OutBytesMetric     int64 `json:"outBytesMetric"`
}

func convertToStatsResult(s *stats) *statsResult {
	sr := &statsResult{}
	sr.InMsgs = s.msg.InMsgs
	sr.InBytes = s.msg.InBytes
	sr.OutMsgs = s.msg.OutMsgs
	sr.OutBytes = s.msg.OutBytes
	sr.TotalConnections = s.conn.TotalConnections
	sr.CurrentConnections = s.conn.CurrentConnections
	sr.InMsgsMetric = getMetric(s, metricTypeMsgIn)
	sr.OutMsgsMetric = getMetric(s, metricTypeMsgOut)
	sr.OutBytesMetric = getMetric(s, metricTypeMsgBytesOut)
	sr.InBytesMetric = getMetric(s, metricTypeMsgBytesIn)
	return sr
}

func getMetric(s *stats, mt metricType) int64 {
	var sum int64 = 0
	for i := 0; i < len(s.instMetric[mt].samples); i++ {
		sum += s.instMetric[mt].samples[i]
	}
	return sum / int64(len(s.instMetric[mt].samples))
}
