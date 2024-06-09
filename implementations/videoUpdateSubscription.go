package implementations

import "github.com/Nikhil-Giramkar/Pub-Sub-Go/contracts"

type videoUpdateSubscription struct {
	notify  chan contracts.VideoUpdateEvent
	creator *creator
}

func (v videoUpdateSubscription) Unsubscribe() {
	v.creator.unsubscribeToUpdates(v)
}

func (v videoUpdateSubscription) Notify() chan contracts.VideoUpdateEvent {
	return v.notify
}
