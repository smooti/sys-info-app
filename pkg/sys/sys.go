package sys

type HardwareInfo struct {
	Manufacturer string
	Model        string
	SerialNumber string
	UUID         string
}

type Collector interface {
	GetHardwareInfo() (*HardwareInfo, error)
}

func NewCollector() Collector {
	return newCollector()
}
