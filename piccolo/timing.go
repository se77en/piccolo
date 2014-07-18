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

func StartTiming(interval time.Duration) {
	ticker := time.NewTicker(interval)
	log.Println(ticker.C)
	go doTimingJob()
}

func doTimingJob(c <-chan time.Time) {
	for {
		<-c
		timingCount++
		for _, fn := range timingFuncs {
			log.Println(fn.FuncJob)
			if timingCount%fn.Ticker == 0 {
				fn.FuncJob()
			}
		}
		if timingCount > 999 {
			timingCount = 0
		}
	}
}
