package piccolo

import (
	"time"
)

type Piccolo struct {
	Jobs []*Job
	add  chan *Job
}

type Job struct {
	Job    func()
	Ticker int
}

type byTime []*Job

func (s byTime) Len() int {
	return len(s)
}
func (s byTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byTime) Less(i, j int) bool {

}
