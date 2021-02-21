package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"com.gft.tsbo-training.src.go/common/device/implementation/devicevalue"
	"com.gft.tsbo-training.src.go/measured/measured"
)

// ###########################################################################
// ###########################################################################
// MAIN
// ###########################################################################
// ###########################################################################

func main() {

	var ms measured.Measured
	measured.InitFromArgs(&ms, os.Args, nil)
	ms.HTTPClient.Timeout = time.Duration(750 * time.Millisecond)
	ms.Mutex.Lock()

	go func() {

		if len(ms.GetUpstream()) == 0 {
			return
		}

		for ever := true; ever; ever = true {
			ms.Condition.Wait()
			var data measured.MeasurePut
			measured.InitMeasurePut(&data, "data", &ms)
			msgBytes, _ := json.MarshalIndent(data, "", "  ")
			r, err := http.NewRequest(http.MethodPut, ms.GetUpstream(), strings.NewReader(string(msgBytes)))

			if err != nil {
				ms.GetLogger().Printf("Error: Failed to report value to '%s'!, error was '%s'!\n", ms.GetUpstream(), err.Error())
				continue
			}

			ms.SetRequestHeaders("", r, nil)

			res, err := ms.HTTPClient.Do(r)
			if err != nil {
				ms.GetLogger().Printf("Error: Failed to report value to '%s'!, error was '%s'!\n", ms.GetUpstream(), err.Error())
				continue
			}

			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)

			if err != nil {
				// ms.GetLogger().Printf("!!! Success !!!")
				continue
			}

			if res.StatusCode != http.StatusCreated {

				if res.StatusCode != http.StatusOK {
					ms.GetLogger().Printf("Error: Failed to report value to '%s'!, status was '%d',  message was '%s'!\n", ms.GetUpstream(), res.StatusCode, body)
					continue
				}
			}
		}
	}()

	go func() {
		var useMutex bool

		useMutex = len(ms.GetUpstream()) > 0

		for ever := true; ever; ever = true {
			if useMutex {
				ms.Mutex.Lock()
			}
			ms.Simulate()
			var v devicevalue.DeviceValue
			ms.FillDeviceValue(&v)
			ms.GetLogger().Printf("New value '%s' for device '%s' at '%s'.", v.GetFormatted(), ms.GetDeviceAddress(), v.GetStamp().Format("2006-01-02 15:04:05"))
			if useMutex {
				ms.Mutex.Unlock()
				ms.Condition.Broadcast()
			}
			time.Sleep(time.Duration(ms.GetInterval()) * time.Millisecond)
		}
	}()
	ms.Run()
}
