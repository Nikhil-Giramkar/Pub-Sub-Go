package implementations

import (
	"fmt"

	"github.com/Nikhil-Giramkar/Pub-Sub-Go/contracts"
)

type user struct {
	name         string
	subscription contracts.Subscription[contracts.VideoUpdateEvent]
}

func (u *user) Initialize(name string, subscription contracts.Subscription[contracts.VideoUpdateEvent]) {
	u.name = name
	u.subscription = subscription
}
func (u user) GetName() string {
	return u.name
}


func (u user) ListenUpdates() {
	videoUpdate := <-u.subscription.Notify()
	fmt.Printf("%s receieved new video - %s\n", u.name, videoUpdate.GetVideo().GetTitle())
}
