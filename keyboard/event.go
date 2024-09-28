package keyboard

type Event struct {
	Pos  Pos
	Made bool
	Time uint32
}

type Pos struct {
	// Layer uint8 // TBD
	Row uint8
	Col uint8
}

type EventReceiver interface {
	// ReceiveEvent is called by the keyboard task loop for matrix state change
	// events. Return value of true indicates that the event is considered to be
	// handled and should not be propagated for handling by the keyboard, and a
	// return value of false indicates the event should be processed normally.
	ReceiveEvent(ev Event) (bool, error)
}

type EventReceiverFunc func(ev Event) (bool, error)

func (recv EventReceiverFunc) ReceiveEvent(ev Event) (bool, error) {
	return recv(ev)
}
