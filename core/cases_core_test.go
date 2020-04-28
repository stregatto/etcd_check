package core

import "github.com/stregatto/etcd_check/etcd"

type expected struct {
	status        bool
	failedMembers raftValue
}

var testsCasesRaftIndexPerMember = []struct {
	raftDrift int
	irpm      etcd.RaftIndexPerMember
	expected  expected
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
			failedMembers raftValue
		}{
			status:        true,        // is expected a true statement, drift is ok
			failedMembers: raftValue{}, // no failing members returned
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
			failedMembers raftValue
		}{
			status:        true,        // is expected a true statement, drift is ok
			failedMembers: raftValue{}, // no failing members returned
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
			failedMembers raftValue
		}{
			status: false, // is expected a false statement, drift is not ok
			failedMembers: raftValue{
				12: {"etcd2"},
			}, // etcd2 member is failing
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
			failedMembers raftValue
		}{
			status: false, // is expected a false statement, drift is not ok
			failedMembers: raftValue{
				2: {"etcd2"},
			}, // etcd2 member is failing
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
			failedMembers raftValue
		}{
			status: false, // is expected a false statement, drift is not ok
			failedMembers: raftValue{
				10: {"etcd1"}, 12: {"etcd2"}, 14: {"etcd3"},
			}, //  all member are failing
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
			failedMembers raftValue
		}{
			status: false, // is expected a false statement, drift is not ok
			failedMembers: raftValue{
				13: {"etcd3"}, 9: {"etcd5"},
			}, // etcd3 and etcd5 are failing
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
				12,
				nil,
			},
		},
		expected: struct {
			status        bool
			failedMembers raftValue
		}{
			status: false, // is expected a false statement, drift is not ok
			failedMembers: raftValue{
				10: {"etcd1"},
				11: {"etcd2", "etcd3"},
				12: {"etcd4", "etcd5"},
			}, // in a way or in another all members are failing, I cannot say which.
		},
	},
}
