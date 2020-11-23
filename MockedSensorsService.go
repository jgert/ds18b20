package ds18b20

import "math/rand"

type MockedSensorsService struct {
	sensors []Sensor
}

func (m MockedSensorsService) Sensors() ([]Sensor, error) {
	return []Sensor{
		MockedSensor{
			name: "Sensor01",
		},
		MockedSensor{
			name: "Sensor02",
		},
		MockedSensor{
			name: "Sensor03",
		},
		MockedSensor{
			name: "Sensor04",
		},
	}, nil
}

func (m MockedSensorsService) ReadTemperature(id string) (float64, error) {
	return rand.Float64(), nil
}
