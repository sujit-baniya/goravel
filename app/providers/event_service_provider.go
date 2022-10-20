package providers

import (
	contractevent "github.com/sujit-baniya/framework/contracts/event"
	"github.com/sujit-baniya/framework/facades"
)

type EventServiceProvider struct {
}

func (receiver *EventServiceProvider) Boot() {

}

func (receiver *EventServiceProvider) Register() {
	facades.Event.Register(receiver.listen())
}

func (receiver *EventServiceProvider) listen() map[contractevent.Event][]contractevent.Listener {
	return map[contractevent.Event][]contractevent.Listener{}
}
