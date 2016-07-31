// Package hid provides access to Human Interface Devices.
package hid

// DeviceInfo provides general information about a device
type DeviceInfo struct {
	// Path contains a Platform-specific device path which is used to identify the device
	Path string
	// VendorId contains the USB Vendor ID of the device
	VendorId uint16
	// ProductId contains the USB Product ID of the device
	ProductId uint16
	// VersionNumber contains the Version / Release Number of the device
	VersionNumber uint16
	// Manufacturer of the USB device
	Manufacturer string
	// Product contains the product name of the device
	Product string

	UsagePage uint16
	Usage     uint16

	InputReportLength   uint16
	OutputReportLength  uint16
	FeatureReportLength uint16
}

// Device interface for an opened HID USB device
type Device interface {
	// Close closes the device and release all keept resources.
	Close()

	// Write writes and output report to device.
	Write([]byte) error

	// ReadCh returns a channel that will return input reports from the device.
	ReadCh() <-chan []byte
}
