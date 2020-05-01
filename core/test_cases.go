package core

import (
	"github.com/stregatto/etcd_check/etcd"
)

type expectedForRaftIndexPerMember struct {
	status        bool
	failedMembers RaftValue
	nagios        string
}

var testCasesIsBetween = []struct {
	value    uint64
	min      uint64
	max      uint64
	expected bool
}{
	{
		10,
		10,
		10,
		true,
	}, {
		10,
		9,
		11,
		true,
	}, {
		11,
		8,
		10,
		false,
	},
	{
		11,
		12,
		13,
		false,
	},
}

var testsCasesRaftIndexPerMember = []struct {
	raftDrift int
	irpm      etcd.RaftIndexPerMember
	expected  expectedForRaftIndexPerMember
}{
	{
		raftDrift: 0, // Test for 0 drift
		irpm: etcd.RaftIndexPerMember{
			"etcd1": {
				111,
				10,
				nil,
			},
			"etcd2": {
				222,
				10,
				nil,
			},
			"etcd3": {
				333,
				10,
				nil,
			},
		},
		expected: struct {
			status        bool
			failedMembers RaftValue
			nagios        string
		}{
			true,        // is expected a true statement, drift is ok
			RaftValue{}, // no failing members returned
			nagiosOk,
		},
	},
	{
		raftDrift: 1, // Test for 1 drift ok
		irpm: etcd.RaftIndexPerMember{
			"etcd1": {
				111,
				10,
				nil,
			},
			"etcd2": {
				222,
				11,
				nil,
			},
			"etcd3": {
				333,
				10,
				nil,
			},
		},
		expected: struct {
			status        bool
			failedMembers RaftValue
			nagios        string
		}{
			true,        // is expected a true statement, drift is ok
			RaftValue{}, // no failing members returned
			nagiosOk,
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, one node is failing, is beyond the other
		irpm: etcd.RaftIndexPerMember{
			"etcd1": {
				111,
				10,
				nil,
			},
			"etcd2": {
				222,
				12,
				nil,
			},
			"etcd3": {
				333,
				10,
				nil,
			},
		},
		expected: struct {
			status        bool
			failedMembers RaftValue
			nagios        string
		}{
			false, // is expected a false statement, drift is not ok
			RaftValue{
				12: {"etcd2"},
			}, // etcd2 member is failing
			nagiosCritical + " 12:[etcd2]",
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, one node is failing, is before the other.
		irpm: etcd.RaftIndexPerMember{
			"etcd1": {
				111,
				10,
				nil,
			},
			"etcd2": {
				222,
				12,
				nil,
			},
			"etcd3": {
				333,
				12,
				nil,
			},
		},
		expected: struct {
			status        bool
			failedMembers RaftValue
			nagios        string
		}{
			false, // is expected a false statement, drift is not ok
			RaftValue{
				10: {"etcd1"},
			}, // etcd2 member is failing
			nagiosCritical + " 10:[etcd1]",
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, all members are failing
		irpm: etcd.RaftIndexPerMember{
			"etcd1": {
				111,
				10,
				nil,
			},
			"etcd2": {
				222,
				12,
				nil,
			},
			"etcd3": {
				333,
				14,
				nil,
			},
		},
		expected: struct {
			status        bool
			failedMembers RaftValue
			nagios        string
		}{
			false, // is expected a false statement, drift is not ok
			RaftValue{
				10: {"etcd1"}, 12: {"etcd2"}, 14: {"etcd3"},
			}, //  all member are failing
			nagiosCritical + " 10:[etcd1] 12:[etcd2] 14:[etcd3]",
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, 2 members are failing
		irpm: etcd.RaftIndexPerMember{
			"etcd1": {
				111,
				10,
				nil,
			},
			"etcd2": {
				222,
				11,
				nil,
			},
			"etcd3": {
				333,
				13,
				nil,
			},
			"etcd4": {
				444,
				10,
				nil,
			},
			"etcd5": {
				555,
				9,
				nil,
			},
		},
		expected: struct {
			status        bool
			failedMembers RaftValue
			nagios        string
		}{
			false, // is expected a false statement, drift is not ok
			RaftValue{
				13: {"etcd3"},
			}, // etcd3 and etcd5 are failing
			nagiosCritical + " 13:[etcd3]",
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, unknown members are failing
		irpm: etcd.RaftIndexPerMember{
			"etcd1": {
				111,
				10,
				nil,
			},
			"etcd2": {
				222,
				11,
				nil,
			},
			"etcd3": {
				333,
				11,
				nil,
			},
			"etcd4": {
				444,
				12,
				nil,
			},
			"etcd5": {
				555,
				15,
				nil,
			},
		},
		expected: struct {
			status        bool
			failedMembers RaftValue
			nagios        string
		}{
			false, // is expected a false statement, drift is not ok
			RaftValue{
				15: {"etcd5"},
			}, // in a way or in another all members are failing, I cannot say which.
			nagiosCritical + " 15:[etcd5]",
		},
	},
}
