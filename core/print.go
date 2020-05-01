package core

import (
	"fmt"
	"strings"
)

const (
	nagiosOk = "OK - All members are reachable and healthy"
	// nagiosWarning	= "WARNING" // WARNING is not used
	nagiosCritical = "CRITICAL - Raft algorithm is failing on:"
	// nagiosUnknown	= "UNKNOWN" // UNKNOWN is not used
)

func raftValuesToString(rv RaftValue) string {
	output := fmt.Sprintf("%v", rv)
	output = strings.TrimPrefix(output, "map[")
	output = strings.TrimSuffix(output, "]")
	return output
}

//PrintNagiosRaftChoerence print the output for nagios-nrpe check
func PrintNagiosRaftChoerence(status bool, raftValues RaftValue) string {
	var output string
	if status {
		output = fmt.Sprintf("%s", nagiosOk)
	} else {
		output = fmt.Sprintf("%s %s", nagiosCritical, raftValuesToString(raftValues))
	}
	fmt.Println(output)
	return output
}
