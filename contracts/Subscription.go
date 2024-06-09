package contracts

type Subscription[T any] interface {
	Unsubscribe()
	Notify() chan T
}
