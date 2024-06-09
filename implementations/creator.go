package implementations

import (
	"fmt"
	"sync"

	"github.com/Nikhil-Giramkar/Pub-Sub-Go/contracts"
)

type creator struct {
	name         string
	videos       []contracts.Video
	channels     map[videoUpdateSubscription]bool
	channelsLock sync.RWMutex
}

func (c *creator) Initialize(name string) {
	c.channels = make(map[videoUpdateSubscription]bool)
	c.name = name
}

func (c creator) GetName() string {
	return c.name
}

func (c creator) GetAllVideoTitles() []string {
	var allVideoTitles []string
	for _, video := range c.videos {
		allVideoTitles = append(allVideoTitles, video.GetTitle())
	}
	return allVideoTitles
}

func (c *creator) AddVideo(newTitle, newDescription string) {
	newVideo := video{
		title:       newTitle,
		description: newDescription,
	}
	c.videos = append(c.videos, newVideo)
	fmt.Print("New Video Published\n")
	//Publish when new video added
	newVideoUpdate := &videoUpdateEvent{
		video: newVideo,
	}
	c.publish(newVideoUpdate)
}

func (c creator) SubscribeToCreator() contracts.Subscription[contracts.VideoUpdateEvent] {
	c.channelsLock.Lock()
	defer c.channelsLock.Unlock()
	//Create a subscription object for user for Video-Updates
	subscription := videoUpdateSubscription{
		notify:  make(chan contracts.VideoUpdateEvent),
		creator: &c,
	}
	c.channels[subscription] = true
	return subscription
}

func (c creator) unsubscribeToUpdates(subscription videoUpdateSubscription) {
	c.channelsLock.Lock()
	defer c.channelsLock.Unlock()

	//Unsubscribe the user by deleting his subscription object
	delete(c.channels, subscription)

}

func (c *creator) publish(videoUpdateEvent contracts.VideoUpdateEvent) {
	c.channelsLock.RLock()
	for k, _ := range c.channels {
		go func(subscription videoUpdateSubscription) {
			subscription.notify <- videoUpdateEvent
		}(k)
	}
	c.channelsLock.RUnlock()
}
