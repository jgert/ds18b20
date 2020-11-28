package ds18b20

import "math/rand"

type MockedSensorsService struct {
	sensors []string
}

func (m MockedSensorsService) Sensors() ([]string, error) {
	return []string{
		"mocked-id-01",
		"mocked-id-02",
		"mocked-id-03",
		"mocked-id-04",
	}, nil
}

func (m MockedSensorsService) ReadTemperature(id string) (float64, error) {
	return rand.Float64(), nil
}
