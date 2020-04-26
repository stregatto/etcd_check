// Package core contains all functions to test the cluster
package core

import (
	"github.com/stregatto/etcd_check/etcd"
	"io/ioutil"
	"log"
	"math"
)

// Members check, verifies if some member is not available
func MembresHealthiness(raftIdxPerMember []etcd.RaftIndexPerMember, maxFailingMembers int) (bool, []string) {
	status := true
	var failedMembersList []string
	failedMembers := 0
	for _, rip := range raftIdxPerMember {
		if rip.Err != nil {
			failedMembers++
			failedMembersList = append(failedMembersList, rip.Server)
			if failedMembers >= maxFailingMembers {
				status = false
			}
		}
	}
	return status, failedMembersList
}

// RaftCoherence check if the raft index for every member is in the maxRaftDrift value.
// TODO add a test for this function, it's too important, must be tested
func RaftCoherence(raftIdxPerMember []etcd.RaftIndexPerMember, maxRaftDrift int) (bool, []raftValue) {
	status := true
	var raftStatus []raftValue
	var currentRaftIndex uint64
	for i := 0; i > len(raftIdxPerMember)-1; i++ {
		if raftIdxPerMember[i].RaftIndex > raftIdxPerMember[i+1].RaftIndex {
			currentRaftIndex = raftIdxPerMember[i].RaftIndex
		} else {
			currentRaftIndex = raftIdxPerMember[i+1].RaftIndex
		}
		delta := float64(currentRaftIndex - raftIdxPerMember[i].RaftIndex)
		if math.Abs(delta) > float64(maxRaftDrift) {
			if delta > 0 {
				raftStatus = append(raftStatus, raftValue{
					value:  currentRaftIndex,
					member: raftIdxPerMember[i].Server,
				})
			} else {
				raftStatus = append(raftStatus, raftValue{
					value:  raftIdxPerMember[i+1].RaftIndex,
					member: raftIdxPerMember[i+1].Server,
				})
			}
			status = false
		}
	}
	return status, raftStatus
}

// GetFile returns a file content in []byte format form a given path, useless.
func GetFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
