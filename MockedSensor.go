package ds18b20

type MockedSensor struct {
	name string
}

func (m MockedSensor) GetID() string {
	return m.name
}

