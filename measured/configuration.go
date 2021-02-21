package measured

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	"com.gft.tsbo-training.src.go/common/ms-framework/microservice"
)

// type deviceHeaders []string

// func (cs *deviceHeaders) String() string { return "" }
// func (cs *deviceHeaders) Set(value string) error {
// 	*cs = append(*cs, value)
// 	return nil
// }

// // HeaderConfiguration ...
// type HeaderConfiguration struct {
// 	HeaderStrings []string `json:"headers"`
// }

// UpstreamConfiguration ...
type UpstreamConfiguration struct {
	Upstream string `json:"logger"`
}

// IUpstreamConfiguration ...
type IUpstreamConfiguration interface {
	GetUpstream() string
}

// DeviceConfiguration ...
type DeviceConfiguration struct {
	DeviceType    string `json:"type"`
	DeviceAddress string `json:"address"`
}

// IDeviceConfiguration ...
type IDeviceConfiguration interface {
	GetDeviceType() string
	GetDeviceAddress() string
}

// Configuration ...
type Configuration struct {
	microservice.Configuration
	// HeaderConfiguration
	UpstreamConfiguration
	DeviceConfiguration
	Interval int
}

// IConfiguration ...
type IConfiguration interface {
	microservice.IConfiguration
	IUpstreamConfiguration
	IDeviceConfiguration
}

// GetUpstream ...
func (cfg UpstreamConfiguration) GetUpstream() string { return cfg.Upstream }

// GetDeviceType ...
func (cfg DeviceConfiguration) GetDeviceType() string { return cfg.DeviceType }

// GetDeviceAddress ...
func (cfg DeviceConfiguration) GetDeviceAddress() string { return cfg.DeviceAddress }

// ---------------------------------------------------------------------------

// InitConfigurationFromArgs ...
func InitConfigurationFromArgs(cfg *Configuration, args []string, flagset *flag.FlagSet) {
	// var dhCli deviceHeaders

	if flagset == nil {
		flagset = flag.NewFlagSet("measured", flag.PanicOnError)
	}

	// flagset.Var(&dhCli, "header", "Device headers")
	pupstream := flagset.String("upstream", "", "Name of the service.")
	pdeviceType := flagset.String("type", "", "Type of device ('thermometer' or 'hygrometer').")
	pDeviceAddress := flagset.String("address", "", "Address of the device.")
	pinterval := flagset.Int("interval", -1, "Interval to generate value.")

	microservice.InitConfigurationFromArgs(&cfg.Configuration, args, flagset)

	if *pinterval >= 0 {
		cfg.Interval = *pinterval
	} else {
		ev := os.Getenv("MS_INTERVAL")
		if len(ev) > 0 {
			cfg.Interval, _ = strconv.Atoi(ev)
		}
	}

	// if len(dhCli) > 0 {
	// 	cfg.HeaderStrings = dhCli
	// }

	if len(*pupstream) > 0 {
		cfg.Upstream = *pupstream
	}

	if len(*pdeviceType) > 0 {
		cfg.DeviceType = *pdeviceType
	}

	if len(*pDeviceAddress) > 0 {
		cfg.DeviceAddress = *pDeviceAddress
	}

	if len(cfg.GetConfigurationFile()) > 0 {
		file, err := os.Open(cfg.GetConfigurationFile())

		if err != nil {
			flagset.Usage()
			panic(fmt.Sprintf(fmt.Sprintf("Error: Failed to open onfiguration file '%s'. Error was %s!\n", cfg.GetConfigurationFile(), err.Error())))
		}

		defer file.Close()
		decoder := json.NewDecoder(file)
		configuration := Configuration{}
		err = decoder.Decode(&configuration)
		if err != nil {
			flagset.Usage()
			panic(fmt.Sprintf(fmt.Sprintf("Error: Failed to parse onfiguration file '%s'. Error was %s!\n", cfg.GetConfigurationFile(), err.Error())))
		}
		if len(cfg.Upstream) == 0 {
			cfg.Upstream = configuration.Upstream
		}

		if len(cfg.DeviceType) == 0 {
			cfg.DeviceType = configuration.DeviceType
		}

		if len(cfg.DeviceAddress) == 0 {
			cfg.DeviceAddress = configuration.DeviceAddress
		}

		if cfg.Interval < 0 {
			cfg.Interval = configuration.Interval
		}

	}

	if len(cfg.Upstream) == 0 {
		cfg.Upstream = os.Getenv("MS_UPSTREAM")
	}

	if len(cfg.DeviceType) == 0 {
		cfg.DeviceType = os.Getenv("MS_DEVICETYPE")
	}

	if len(cfg.DeviceAddress) == 0 {
		cfg.DeviceAddress = os.Getenv("MS_DEVICEADDRESS")
	}
}
