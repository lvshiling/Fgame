package monitor

import (
	"net/http"
	"sync/atomic"
	"time"

	"fgame/fgame/gm/gamegm/session"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/xozrc/pkg/httputils"
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
}

func newStats() *stats {
	s := &stats{
		msg:  &msgStats{},
		conn: &connStats{},
	}

	for i := metricTypeMsgIn; i <= metricTypeMsgBytesOut; i++ {
		s.instMetric[i] = &metric{}
	}

	s.done = make(chan struct{}, 0)
	s.Start()
	return s
}

func (s *stats) Start() {
	go func() {
	Loop:
		for {
			select {
			case <-time.After(time.Millisecond * statsMetricCronTime):
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
	s.done <- struct{}{}
}

func (s *stats) trackInstantaneousMetric(mt metricType, current int64) {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	elapse := now - s.instMetric[mt].lastSampleTime
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

//连接统计中间件 连接开启
func handleSessionOpenStats(s session.Session) error {
	gw := MonitorServiceInContext(s.Context())
	atomic.AddInt64(&gw.stats.conn.CurrentConnections, 1)
	atomic.AddInt64(&gw.stats.conn.TotalConnections, 1)
	return nil
}

//连接统计中间件 连接关闭
func handleSessionCloseStats(s session.Session) error {
	gw := MonitorServiceInContext(s.Context())
	atomic.AddInt64(&gw.stats.conn.CurrentConnections, -1)
	return nil
}

//消息数
type msgStats struct {
	InMsgs   int64 `json:"inMsgs"`
	InBytes  int64 `json:"inBytes"`
	OutMsgs  int64 `json:"outMsgs"`
	OutBytes int64 `json:"outBytes"`
}

//消息接收统计中间件
func handleSessionRecvStats(s session.Session, msg []byte) error {
	gw := MonitorServiceInContext(s.Context())
	atomic.AddInt64(&gw.stats.msg.InMsgs, 1)
	atomic.AddInt64(&gw.stats.msg.InBytes, int64(len(msg)))
	return nil
}

//消息发送统计中间件
func handleSessionSendStats(s session.Session, msg []byte) error {
	gw := MonitorServiceInContext(s.Context())
	msgLen := int64(len(msg))
	atomic.AddInt64(&gw.stats.msg.OutMsgs, 1)
	atomic.AddInt64(&gw.stats.msg.OutBytes, msgLen)
	return nil
}

func StatsRouter(r *mux.Router) {
	r.Path("/").Handler(http.HandlerFunc(handleStats))
	r.Path("/conn").Handler(http.HandlerFunc(handleConnStats))
	r.Path("/msg").Handler(http.HandlerFunc(handleMsgStats))
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

//统计http请求
func handleStats(rw http.ResponseWriter, req *http.Request) {
	gw := MonitorServiceInContext(req.Context())
	sr := convertToStatsResult(gw.stats)
	err := httputils.WriteJSON(rw, http.StatusOK, sr)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

//连接统计http请求
func handleConnStats(rw http.ResponseWriter, req *http.Request) {
	gw := MonitorServiceInContext(req.Context())
	err := httputils.WriteJSON(rw, http.StatusOK, gw.stats.conn)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

//接收消息和发送消息http统计
func handleMsgStats(rw http.ResponseWriter, req *http.Request) {
	gw := MonitorServiceInContext(req.Context())
	err := httputils.WriteJSON(rw, http.StatusOK, gw.stats.msg)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

//设置棋牌服务
func SetupQipaiServiceHandler(qps *MonitorService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithMonitorService(ctx, qps)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
