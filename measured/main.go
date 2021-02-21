package measured

import (
	"flag"
	"sync"
	"time"

	"com.gft.tsbo-training.src.go/common/device"
	"com.gft.tsbo-training.src.go/common/ms-framework/microservice"
)

// ###########################################################################
// ###########################################################################
//Measured
// ###########################################################################
// ###########################################################################

//Measured Encapsulates the ms-measure data
type Measured struct {
	microservice.MicroService
	// *HeaderConfiguration
	*UpstreamConfiguration
	*DeviceConfiguration

	device.Device
	// Headers   []Header
	starttime time.Time
	HasValue  bool
	Mutex     *sync.Mutex
	Condition *sync.Cond
}

// ###########################################################################

// InitFromArgs Constructor of a new ms-uptime
func InitFromArgs(ms *Measured, args []string, flagset *flag.FlagSet) {
	var cfg Configuration

	if flagset == nil {
		flagset = flag.NewFlagSet("measured", flag.PanicOnError)
	}

	InitConfigurationFromArgs(&cfg, args, flagset)
	ms.HeaderConfiguration = &cfg.HeaderConfiguration
	ms.UpstreamConfiguration = &cfg.UpstreamConfiguration
	ms.DeviceConfiguration = &cfg.DeviceConfiguration

	microservice.Init(&ms.MicroService, &cfg.Configuration, nil)

	if cfg.GetDeviceType() == "thermometer" {
		device.InitThermometer(&ms.Device, cfg.GetDeviceAddress(), -1, -1, cfg.Interval)
	} else if cfg.DeviceType == "hygrometer" {
		device.InitHygrometer(&ms.Device, cfg.GetDeviceAddress(), -1, -1, cfg.Interval)
	} else {
		flagset.Usage()
		panic("Error: Wrong device! Use 'thermometer' or 'hygrometer'.")
	}

	ms.starttime = time.Now()
	ms.HasValue = false
	ms.Mutex = &sync.Mutex{}
	ms.Condition = sync.NewCond(ms.Mutex)

	// for _, line := range ms.HeaderStrings {
	// 	h := HeaderFromString(line)
	// 	ms.Headers = append(ms.Headers, h)
	// }

}
