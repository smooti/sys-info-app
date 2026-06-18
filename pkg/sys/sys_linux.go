package sys

import "errors"

type linuxCollector struct{}

func newCollector() Collector {
	return &linuxCollector{}
}

// GetHardwareInfo implements the Collector interface for Linux
func (l *linuxCollector) GetHardwareInfo() (*HardwareInfo, error) {
	// TODO: Pull from /sys/class/dmi/id/product_uuid, product_serial, etc.
	return nil, errors.New("hardware info collection is not yet implemented on linux")
}
