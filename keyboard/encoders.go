package keyboard

type Encoder interface {
	Value() int
	SetValue(int)
}

type EncoderSubscriber interface {
	EncoderChanged(index int, clockwise bool)
}

type EncoderSubscriberFunc func(index int, clockwise bool)

func (fn EncoderSubscriberFunc) EncoderChanged(index int, clockwise bool) {
	fn(index, clockwise)
}

type encoders struct {
	subcribers []EncoderSubscriber
	encoders   []Encoder
	values     []int
}

func (encs *encoders) Task() {
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
					sub.EncoderChanged(i, clockwise)
				}
			}
		}
	}
}
