package client

import (
	"fmt"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Config struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type Client struct {
	TradeManager
	conn *grpc.ClientConn
}

const (
	waitBetween = time.Second //100 * time.Millisecond
	maxRetry    = 10
)

func NewClient(cfg *Config) (c *Client, err error) {
	ip := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	var options []grpc.DialOption
	options = append(options, grpc.WithInsecure())
	//TODO:rpc重试
	callOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(waitBetween)),
		grpc_retry.WithCodes(codes.Unavailable, codes.Internal, codes.Aborted),
		grpc_retry.WithMax(maxRetry),
	}
	options = append(options, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(callOpts...)))
	//TODO:重连时间
	maxDelayOption := grpc.WithBackoffMaxDelay(5 * time.Second)
	options = append(options, maxDelayOption)

	conn, err := grpc.Dial(ip, options...)
	if err != nil {
		return
	}
	c = &Client{}
	c.conn = conn
	c.TradeManager = NewTradeManager(c)
	return c, nil
}
