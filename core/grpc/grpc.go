package grpc

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type GrpcClientConfig struct {
	Endpoints   []string `json:"endpoints"`
	DialOptions []grpc.DialOption
	DialTimeout time.Duration `json:"dialTimeout"`
}

type GrpcClient struct {
	cfg      *GrpcClientConfig
	ctx      context.Context
	cancel   context.CancelFunc
	conn     *grpc.ClientConn
	callOpts []grpc.CallOption
	balancer grpc.Balancer
}

func (c *GrpcClient) init() (err error) {
	baseCtx := context.TODO()
	ctx, cancel := context.WithCancel(baseCtx)
	c.ctx = ctx
	c.cancel = cancel
	//TODO 修改
	c.callOpts = defaultCallOpts
	conn, err := c.dial(c.cfg.Endpoints[0], grpc.WithBalancer(c.balancer))
	if err != nil {
		c.cancel()
		c.balancer.Close()
		return
	}
	c.conn = conn
	return
}

func getHost(ep string) string {
	url, uerr := url.Parse(ep)
	if uerr != nil || !strings.Contains(ep, "://") {
		return ep
	}
	return url.Host
}

func (c *GrpcClient) dial(endpoint string, dopts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts := c.dialSetupOpts(endpoint, dopts...)
	host := getHost(endpoint)

	opts = append(opts, c.cfg.DialOptions...)

	if err != nil {
		conn, err := grpc.DialContext(c.ctx, host, opts...)
		return nil, err
	}
	return conn, nil
}

func (c *GrpcClient) dialSetupOpts(endpoint string, dopts ...grpc.DialOption) (opts []grpc.DialOption) {
	if c.cfg.DialTimeout > 0 {
		opts = []grpc.DialOption{grpc.WithTimeout(c.cfg.DialTimeout)}
	}
	if c.cfg.DialKeepAliveTime > 0 {
		params := keepalive.ClientParameters{
			Time:    c.cfg.DialKeepAliveTime,
			Timeout: c.cfg.DialKeepAliveTimeout,
		}
		opts = append(opts, grpc.WithKeepaliveParams(params))
	}
	opts = append(opts, dopts...)

	f := func(host string, t time.Duration) (net.Conn, error) {
		proto, host, _ := parseEndpoint(c.balancer.Endpoint(host))
		if host == "" && endpoint != "" {
			// dialing an endpoint not in the balancer; use
			// endpoint passed into dial
			proto, host, _ = parseEndpoint(endpoint)
		}
		if proto == "" {
			return nil, fmt.Errorf("unknown scheme for %q", host)
		}
		select {
		case <-c.ctx.Done():
			return nil, c.ctx.Err()
		default:
		}
		dialer := &net.Dialer{Timeout: t}
		conn, err := dialer.DialContext(c.ctx, proto, host)
		if err != nil {
			select {
			case c.dialerrc <- err:
			default:
			}
		}
		return conn, err
	}
	opts = append(opts, grpc.WithDialer(f))

	creds := c.creds
	if _, _, scheme := parseEndpoint(endpoint); len(scheme) != 0 {
		creds = c.processCreds(scheme)
	}
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	return opts
}

func NewGrpcClient(cfg *GrpcClientConfig) (c *GrpcClient, err error) {
	return nil, nil
}
