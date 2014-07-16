package piccolo

import (
	"time"
)

type timingFunc struct {
	FuncJob func()
	Ticker  int
}

var (
	timingCount int
	timingFuncs map[string]*timingFunc
)

func init() {
	timingCount = 0
	timingFuncs = map[string]*timingFunc{}
}

func AddTimingFunc(name string, ticker int, funcJob func()) {
	fn := new(timingFunc)
	fn.FuncJob = funcJob
	fn.Ticker = ticker
	timingFuncs[name] = fn
}

func StartTiming() {
	ticker := time.NewTicker(time.Duration(10) * time.Minute)
	go doTimingJob(ticker.C)
}

func doTimingJob(c <-chan time.Time) {
	for {
		<-c
		timingCount++
		for _, fn := range timingFuncs {
			if timingCount%fn.Ticker == 0 {
				fn.FuncJob()
			}
		}
		if timingCount > 999 {
			timingCount = 0
		}
	}
}
