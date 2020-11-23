package ds18b20

type W1Sensor struct {
	id string
}

func (w W1Sensor) GetID() string {
	return w.id
}
