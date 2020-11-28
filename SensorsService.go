package ds18b20

type SensorsService interface {
	Sensors() ([]string, error)
	ReadTemperature(id string) (float64, error)
}
