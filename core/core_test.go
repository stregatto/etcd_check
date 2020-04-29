package core

import (
	"testing"
)

//areEqual tests two slices of failedMembers if are equal or not.
func areEqual(a, b raftValue) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {

		return false
	}
	for k, v := range a {
		for i, m := range v {
			if val, ok := b[k]; ok {
				if m != val[i] {
					return false
				}
			}
		}
	}
	return true
}

func TestIsBetween(t *testing.T) {
	for _, test := range testIsBetween {
		got := IsBetween(test.value, test.min, test.max)
		if got != test.expected {
			t.Errorf("IsBetween test value:%d, min:%d, max: %d, expected:%t, got:%t",
				test.value, test.min, test.max, test.expected, got)
		}
	}
}

func TestRaftCoherence(t *testing.T) {
	for _, test := range testsCasesRaftIndexPerMember {
		status, failedMembers := RaftCoherence(test.irpm, test.raftDrift)
		// Simplified test, no map check
		// removed || !areEqual(failedMembers, test.expected.failedMembers) for now
		if (status != test.expected.status) || !areEqual(failedMembers, test.expected.failedMembers) {
			t.Errorf("RaftCoherence test: drift %v - %d\n"+
				"\texpected:\t%t - %v\n"+
				"\tgot:\t\t%t - %v ",
				test.irpm,
				test.raftDrift,
				test.expected.status,
				test.expected.failedMembers,
				status,
				failedMembers)
		}
	}
}
