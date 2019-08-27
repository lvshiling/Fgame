package wxapi

import (
	"context"
	"encoding/json"
	httputils "fgame/fgame/pkg/httputils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type RemoteConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}
type RemoteService struct {
	cfg *RemoteConfig
}

const (
	exchangePath = "/api/exchange/exchange"
)

func (m *RemoteService) GetOrder(wxId string, code string) (returnCode int32, returnMsg string, orderId string, money int32, clientIp string, err error) {

	host := m.cfg.Host
	port := m.cfg.Port

	type exchangeForm struct {
		WxId string `json:"wxId"`
		Code string `json:"code"`
	}

	getExchangePath := fmt.Sprintf("http://%s:%d%s", host, port, exchangePath)
	form := &exchangeForm{
		WxId: wxId,
		Code: code,
	}

	result, err := httputils.PostJsonWithRawMessage(getExchangePath, nil, form)
	if err != nil {
		return
	}
	if result.ErrorCode != 0 {
		log.WithFields(
			log.Fields{
				"wxId":      wxId,
				"code":      code,
				"errorCode": result.ErrorCode,
				"errorMsg":  result.ErrorMsg,
			}).Warn("wxapi:获取订单号,错误")
		returnCode = int32(result.ErrorCode)
		returnMsg = result.ErrorMsg
		return
	}

	type getExchangeResponse struct {
		OrderId string `json:"orderId"`
		Money   int32  `json:"money"`
		WxId    string `json:"wxId"`
		Code    string `json:"code"`
	}
	res := &getExchangeResponse{}
	err = json.Unmarshal(result.Result, res)
	if err != nil {
		return
	}
	orderId = res.OrderId
	money = res.Money
	clientIp = host
	return
}

func NewRemoteService(cfg *RemoteConfig) *RemoteService {
	s := &RemoteService{
		cfg: cfg,
	}

	return s
}

type RemoteServiceKey string

const (
	remoteServiceKey = RemoteServiceKey("remoteServiceKey")
)

func RemoteServiceInContext(ctx context.Context) *RemoteService {
	c, ok := ctx.Value(remoteServiceKey).(*RemoteService)
	if !ok {
		return nil
	}
	return c
}
func WithRemoteService(ctx context.Context, ss *RemoteService) context.Context {
	return context.WithValue(ctx, remoteServiceKey, ss)
}
