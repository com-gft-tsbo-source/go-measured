package measured

import "strings"

// ###########################################################################
// ###########################################################################
//Measured
// ###########################################################################
// ###########################################################################

// Header ...
type Header struct {
	Key   string
	Value string
}

// HeaderFromString ...
func HeaderFromString(line string) Header {
	var header Header
	pairs := strings.SplitN(line, ":", 2)
	header.Key = pairs[0]
	header.Value = strings.Trim(pairs[1], " ")
	return header
}
