package providers

import (
	"github.com/sujit-baniya/framework/contracts/queue"
	"github.com/sujit-baniya/framework/facades"
)

type QueueServiceProvider struct {
}

func (receiver *QueueServiceProvider) Boot() {

}

func (receiver *QueueServiceProvider) Register() {
	facades.Queue.Register(receiver.Jobs())
}

func (receiver *QueueServiceProvider) Jobs() []queue.Job {
	return []queue.Job{}
}
