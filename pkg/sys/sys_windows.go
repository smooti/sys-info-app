package sys

import (
	"github.com/yusufpapurcu/wmi"
)

type win32_BIOS struct {
	SerialNumber string
}

type win32_ComputerSystem struct {
	Manufacturer string
	Model        string
}

type win32_ComputerSystemProduct struct {
	UUID string
}

type windowsCollector struct{}

func newCollector() Collector {
	return &windowsCollector{}
}

func (w *windowsCollector) GetHardwareInfo() (*HardwareInfo, error) {
	info := &HardwareInfo{}
	var computerSystemInfo []win32_ComputerSystem
	var biosInfo []win32_BIOS
	var computerSystemProductInfo []win32_ComputerSystemProduct

	// Get manufacturer and model information
	err := wmi.Query("SELECT Manufacturer, Model FROM Win32_ComputerSystem", &computerSystemInfo)
	if err != nil {
		return nil, err
	}
	info.Manufacturer = computerSystemInfo[0].Manufacturer
	info.Model = computerSystemInfo[0].Model

	// Get serial number
	err = wmi.Query("SELECT SerialNumber FROM Win32_BIOS", &biosInfo)
	if err != nil {
		return nil, err
	}
	info.SerialNumber = biosInfo[0].SerialNumber

	// Get UUID
	err = wmi.Query("SELECT UUID FROM Win32_ComputerSystemProduct", &computerSystemProductInfo)
	if err != nil {
		return nil, err
	}
	info.UUID = computerSystemProductInfo[0].UUID

	return info, nil
}
