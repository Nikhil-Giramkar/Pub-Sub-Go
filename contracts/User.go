package contracts

type User interface {
	GetName() string
	ListenUpdates()
	StopUpdates()
}
