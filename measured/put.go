package measured

import (
	"github.com/com-gft-tsbo-source/go-common/device/implementation/devicedescriptor"
	"github.com/com-gft-tsbo-source/go-common/device/implementation/devicevalue"
	"github.com/com-gft-tsbo-source/go-common/ms-framework/microservice"
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
	microservice.InitResponseFromMicroService(&r.Response, ms, 200, status)
	devicedescriptor.InitFromDeviceDescriptor(&r.DeviceDescriptor, &ms.Device)
	ms.FillDeviceValue(&r.DeviceValue)
}
