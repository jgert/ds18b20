package ds18b20

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type W1SensorsService struct {
}

func (w W1SensorsService) Sensors() ([]Sensor, error) {

	data, err := ioutil.ReadFile("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
	if err != nil {
		return []Sensor{}, err
	}

	names := strings.Split(string(data), "\n")
	if len(names) > 0 {
		names = names[:len(names)-1]
	}

	sensors := make([]Sensor, len(names))
	for i, v := range names {
		sensors[i] = W1Sensor{id: v}
	}

	return sensors, nil
}

func (w W1SensorsService) ReadTemperature(id string) (float64, error) {
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/" + id + "/temperature")
	if err != nil {
		return 0, err
	}

	raw := string(data)

	c, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0.0, err
	}

	return c / 1000.0, nil
}
