package core

import (
	"fmt"
	"strings"
)

const (
	nagiosTextOk = "OK - All members are reachable and healthy"
	// nagiosTextWarning			= "WARNING" // WARNING is not used
	nagiosTextCritical = "CRITICAL - Raft algorithm is failing on:"
	// nagiosTextUnknown			= "UNKNOWN" // UNKNOWN is not used
	nagiosTextCriticalMembers = "CRITICAL - ETCD cluster is failing"
)

const (
	nagiosExitOk = 0
	// nagiosExitWarning	= 1	// WARNING is not used
	nagiosExitCritical = 2
	// nagiosExitUnknown	= 3 // Unknown is not used
)

func raftValuesToString(rv RaftValue) string {
	output := fmt.Sprintf("%v", rv)
	output = strings.TrimPrefix(output, "map[")
	output = strings.TrimSuffix(output, "]")
	return output
}

// PrintNagiosFailingMembers prints the failing members for nagios-nrpe check
func PrintNagiosFailingMembers(failingMembers []string, maxFailingMembers int) int {
	output := fmt.Sprintf("%s , more than %d members are failing: %v", nagiosTextCriticalMembers, maxFailingMembers, failingMembers)
	println(output)
	return nagiosExitCritical
}

//PrintNagiosRaftCoherence prints the output for nagios-nrpe check, return exit code
func PrintNagiosRaftCoherence(status bool, raftValues RaftValue) (int, string) {
	var output string
	var exitCode int
	if status {
		output = fmt.Sprintf("%s", nagiosTextOk)
		exitCode = nagiosExitOk
	} else {
		output = fmt.Sprintf("%s %s", nagiosTextCritical, raftValuesToString(raftValues))
		exitCode = nagiosExitCritical
	}
	fmt.Println(output)
	return exitCode, output
}
