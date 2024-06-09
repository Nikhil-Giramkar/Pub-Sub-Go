package implementations

import "github.com/Nikhil-Giramkar/Pub-Sub-Go/contracts"

type videoUpdateEvent struct {
	video contracts.Video
}

func (v videoUpdateEvent) GetVideo() contracts.Video {
	return v.video
}
