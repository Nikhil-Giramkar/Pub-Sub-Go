package contracts

type Creator interface {
	GetName() string
	GetAllVideoTitles() []string
	AddVideo(newTitle, newDescription string)
	SubscribeToCreator() Subscription[VideoUpdateEvent]
}
