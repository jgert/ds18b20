package ds18b20

type SensorsService interface {
	Sensors() ([]Sensor, error)
	ReadTemperature(id string) (float64, error)
}
