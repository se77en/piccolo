package piccolo

import (
	"log"
	"time"
)

type timingFunc struct {
	FuncJob func()
	Ticker  int
}

var (
	timingFuncs map[string]*timingFunc
)

func init() {
	timingFuncs = map[string]*timingFunc{}
}

func AddTimingFunc(name string, ticker int, funcJob func()) {
	fn := new(timingFunc)
	fn.FuncJob = funcJob
	fn.Ticker = ticker
	timingFuncs[name] = fn
}

func StartTiming(interval time.Duration) {
	flag := make(chan bool)
	ticker := time.NewTicker(interval)
	go doTimingJob(ticker.C, flag)
	<-flag
}

func doTimingJob(c <-chan time.Time, flag chan bool) {
	for {
		<-c
		for _, fn := range timingFuncs {
			if fn.Ticker == -1 {
				fn.FuncJob()
				continue
			}
			if fn.Ticker > 0 {
				log.Println(fn.Ticker)
				fn.FuncJob()
				fn.Ticker--
			} else {
				goto DONE
			}
		}
	}
DONE:
	flag <- true
}
