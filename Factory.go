package ds18b20

import "runtime"

func isRaspberry() bool {
	conditions := []bool{
		runtime.GOOS == "linux",
		runtime.GOARCH == "arm",
	}

	var result = true
	for _, v := range conditions {
		result = result && v
	}

	return result
}

func CreateSensorsService() SensorsService {
	if isRaspberry() {
		return W1SensorsService{}
	} else {
		return MockedSensorsService{}
	}
}
