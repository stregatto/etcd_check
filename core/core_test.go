package core

import (
	"testing"
)

//areEqual tests two slices of raftValues if are equal or not.
func areEqual(a, b []raftValue) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i := range a {
		if a[i].member != b[i].member {
			return false
		}
		if a[i].value != b[i].value {
			return false
		}
	}
	return true
}

func TestRaftCoherence(t *testing.T) {
	for _, test := range testsCasesRaftIndexPerMember {
		status, raftValues := RaftCoherence(test.irpm, test.raftDrift)
		if (status != test.expected.status) || areEqual(raftValues, test.expected.raftValues) {
			t.Errorf("RaftCoherence test, test: drift %v - %d , expected: %t - %v, got: %t - %v ",
				test.irpm,
				test.raftDrift,
				test.expected.status,
				test.expected.raftValues,
				status,
				raftValues)
		}
	}
}
