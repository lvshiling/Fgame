package grpc

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrNoAddrAvailable = status.Error(codes.Unavailable, "there is no address available")

type GrpcBalancer struct {
	mu                 sync.RWMutex
	pinAddr            string
	closed             bool
	notifyCh           chan []grpc.Address
	unhealthyHostPorts map[string]time.Time

	donec chan struct{}
}

func (b *GrpcBalancer) Start(target string, config grpc.BalancerConfig) (err error) {
	return
}

func (b *GrpcBalancer) Up(addr grpc.Address) func(error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return func(err error) {}
	}
	if b.pinAddr != "" {
		return func(err error) {}
	}
	b.pinAddr = addr.Addr

	return func(err error) {
	}
}

func (b *GrpcBalancer) Get(ctx context.Context, opts grpc.BalancerGetOptions) (grpc.Address, func(), error) {
	var (
		addr   string
		closed bool
	)

	// If opts.BlockingWait is false (for fail-fast RPCs), it should return
	// an address it has notified via Notify immediately instead of blocking.
	if !opts.BlockingWait {
		b.mu.RLock()
		closed = b.closed
		addr = b.pinAddr
		b.mu.RUnlock()
		if closed {
			return grpc.Address{Addr: ""}, nil, grpc.ErrClientConnClosing
		}
		if addr == "" {
			return grpc.Address{Addr: ""}, nil, ErrNoAddrAvailable
		}
		return grpc.Address{Addr: addr}, func() {}, nil
	}

	for {
		select {
		case <-b.donec:
			return grpc.Address{Addr: ""}, nil, grpc.ErrClientConnClosing
		case <-ctx.Done():
			return grpc.Address{Addr: ""}, nil, ctx.Err()
		}
		b.mu.RLock()
		closed = b.closed
		addr = b.pinAddr
		b.mu.RUnlock()
		// Close() which sets b.closed = true can be called before Get(), Get() must exit if balancer is closed.
		if closed {
			return grpc.Address{Addr: ""}, nil, grpc.ErrClientConnClosing
		}
		if addr != "" {
			break
		}
	}
	return grpc.Address{Addr: addr}, func() {}, nil
}

func (b *GrpcBalancer) Notify() <-chan []grpc.Address {
	return b.notifyCh
}

func (b *GrpcBalancer) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return nil
	}
	b.closed = true
	b.pinAddr = ""
	close(b.notifyCh)
	return nil
}

func (b *GrpcBalancer) UpdateAddr(ep string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if ep == b.pinAddr {
		return
	}
	b.notifyCh <- []grpc.Address{ep2addr(ep)}
}

func ep2addr(ep string) grpc.Address {
	addr := grpc.Address{
		Addr: getHost(ep),
	}
	return addr
}
