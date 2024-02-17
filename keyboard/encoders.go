package keyboard

type Encoders interface {
	EncodersTask()
}

type Encoder interface {
	Value() int
	SetValue(int)
}

type EncodersSubscriber interface {
	EncoderChanged(index int, clockwise bool)
}

type EncodersSubscriberFunc func(index int, clockwise bool)

func (fn EncodersSubscriberFunc) EncoderChanged(index int, clockwise bool) {
	fn(index, clockwise)
}

type encoders struct {
	subcribers []EncodersSubscriber
	encoders   []Encoder
	values     []int
}

func (encs *encoders) EncodersTask() {
	// println("encoder task")
	for i, enc := range encs.encoders {
		_, _ = i, enc
		if newValue, oldValue := enc.Value(), encs.values[i]; newValue != oldValue {
			change := newValue - oldValue
			clockwise := change > 0
			if change < 0 {
				change *= -1
			}
			encs.values[i] = newValue
			for i := 0; i < change; i++ {
				for _, sub := range encs.subcribers {
					// println("encoder value changed", i, clockwise)
					sub.EncoderChanged(i, clockwise)
				}
			}
		}
	}
}
