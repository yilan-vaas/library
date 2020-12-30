package grace

import (
	"os"
	"os/signal"
	"syscall"
)

type Job struct {
	Resume bool
}

func (self *Job) GraceShutdown(callbacks ...func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	self.Resume = true
	go func() {
		<-sigChan
		self.Resume = false
		for _, callback := range callbacks {
			callback()
		}
	}()
}
