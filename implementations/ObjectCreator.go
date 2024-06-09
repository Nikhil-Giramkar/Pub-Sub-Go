package implementations

import "github.com/Nikhil-Giramkar/Pub-Sub-Go/contracts"

func NewUser(name string, subscription contracts.Subscription[contracts.VideoUpdateEvent]) contracts.User {
	var user user
	user.Initialize(name, subscription)
	return user
}

func NewCreator(name string) contracts.Creator {
	var creator creator
	creator.Initialize(name)
	return &creator
}
