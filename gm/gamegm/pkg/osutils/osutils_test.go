package osutils_test

import (
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"
)

var _ = fmt.Print

var signalSlice []syscall.Signal = []syscall.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}

var result = 1

func waitSig(t *testing.T, c <-chan os.Signal, sig os.Signal) {
	select {
	case s := <-c:
		if s != sig {
			t.Fatalf("signal was %v, want %v", s, sig)
		}
	case <-time.After(1 * time.Second):
		t.Fatalf("timeout waiting for %v", sig)
	}
}

func TestInterruptHookRun(t *testing.T) {

	// for _, sig := range signalSlice {

	// 	result = 1
	// 	ih := osutils.NewInterruptHooker()

	// 	fir := &first{}
	// 	sec := &second{}
	// 	thi := &third{}

	// 	ih.AddHandler(fir)
	// 	ih.AddHandler(sec)
	// 	ih.AddHandler(thi)
	// 	ih.RemoveHandler(thi)

	// 	c := make(chan os.Signal, 1)
	// 	signal.Notify(c, sig)

	// 	done := make(chan struct{}, 1)
	// 	go func() {
	// 		ih.Run()
	// 		done <- struct{}{}
	// 	}()

	// 	syscall.Kill(syscall.Getpid(), sig)
	// 	waitSig(t, c, sig)
	// 	<-done
	// 	if result == 3 {
	// 		t.Fatalf("interrupt handlers were called in wrong order")
	// 	}
	// 	if result != 4 {
	// 		t.Fatalf("interrupt handlers were not called properly")
	// 	}

	// }
}

type first struct {
}

func (f *first) Run() {
	result += 1
}

type second struct {
}

func (f *second) Run() {
	result *= 2
}

type third struct {
}

func (f *third) Run() {
	result -= 3
}
