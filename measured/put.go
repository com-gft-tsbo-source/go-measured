package measured

import (
	"com.gft.tsbo-training.src.go/common/device/implementation/devicedescriptor"
	"com.gft.tsbo-training.src.go/common/device/implementation/devicevalue"
	"com.gft.tsbo-training.src.go/common/ms-framework/microservice"
)

// ###########################################################################
// ###########################################################################
// ###########################################################################
// ###########################################################################
// ###########################################################################

// MeasurePut Encapsulates the reply of ms-measure
type MeasurePut struct {
	microservice.Response
	devicedescriptor.DeviceDescriptor
	devicevalue.DeviceValue
}

// ###########################################################################

// InitMeasurePut Constructor of a response of ms-measure
func InitMeasurePut(r *MeasurePut, status string, ms *Measured) {
	microservice.InitResponseFromMicroService(&r.Response, ms, status)
	devicedescriptor.InitFromDeviceDescriptor(&r.DeviceDescriptor, &ms.Device)
	ms.FillDeviceValue(&r.DeviceValue)
}
