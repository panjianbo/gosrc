package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type signalHandler func(sig os.Signal, arg interface{})

type signalSet struct {
	handlers map[os.Signal]signalHandler
}

func signalSetNew() *signalSet {
	ss := new(signalSet)
	ss.handlers = make(map[os.Signal]signalHandler)
	return ss
}

func (handlers *signalSet) register(sig os.Signal, handler signalHandler) {
	if _, found := handlers.handlers[sig]; !found {
		handlers.handlers[sig] = handler
	}
}

func (handlers *signalSet) handle(sig os.Signal, arg interface{}) (err error) {
	if _, found := handlers.handlers[sig]; found {
		handlers.handlers[sig](sig, arg)
		return nil
	} else {
		return fmt.Errorf("No handler available for signal %v", sig)
	}
}

func SignalHandleTest() {
	ss := signalSetNew()
	handler := func(sig os.Signal, arg interface{}) {
		fmt.Printf("handle signal:%v\n", sig)
	}

	ss.register(syscall.SIGINT, handler)
	ss.register(syscall.SIGUSR1, handler)
	ss.register(syscall.SIGUSR2, handler)

	chanSigs := make(chan os.Signal)
	var sigs []os.Signal
	for sig := range ss.handlers {
		sigs = append(sigs, sig)
	}
	signal.Notify(chanSigs)

	for {
		sig := <-chanSigs
		err := ss.handle(sig, nil)
		if err != nil {
			fmt.Printf("unknown signal received: %v\n", sig)
		}
	}
}

func main() {
	go SignalHandleTest()
	time.Sleep(time.Hour)
}

