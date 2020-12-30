package grace

import (
	"fmt"
	"testing"
	"time"
)

type Service struct {
	Job
}

func (self *Service) Do() {
	self.GraceShutdown(self.Shutdown)
	for self.Resume {
		fmt.Println("x")
		time.Sleep(time.Second)
	}
}

func (self *Service) Shutdown() {
	fmt.Println("close")
}

func TestGraceJob(t *testing.T) {
	s := &Service{}
	s.Do()
}
