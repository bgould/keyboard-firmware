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
	ReceiveEvent(ev Event) (bool, error)
}

type EventReceiverFunc func(ev Event) (bool, error)

func (recv EventReceiverFunc) ReceiveEvent(ev Event) (bool, error) {
	return recv(ev)
}
