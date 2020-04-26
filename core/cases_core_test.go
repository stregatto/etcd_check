package core

import "github.com/stregatto/etcd_check/etcd"

type expected struct {
	status     bool
	raftValues []raftValue
}

var testsCasesRaftIndexPerMember = []struct {
	raftDrift int
	irpm      []etcd.RaftIndexPerMember
	expected  expected
}{
	{
		raftDrift: 0, // Test for 0 drift
		irpm: []etcd.RaftIndexPerMember{
			{
				"etcd1",
				111,
				10,
				nil,
			},
			{
				"etcd2",
				222,
				10,
				nil,
			},
			{
				"etcd3",
				222,
				10,
				nil,
			},
		},
		expected: struct {
			status     bool
			raftValues []raftValue
		}{
			status:     true,          // is expected a true statement, drift is ok
			raftValues: []raftValue{}, // no failing members returned
		},
	},
	{
		raftDrift: 1, // Test for 1 drift ok
		irpm: []etcd.RaftIndexPerMember{
			{
				"etcd1",
				111,
				10,
				nil,
			},
			{
				"etcd2",
				222,
				11,
				nil,
			},
			{
				"etcd3",
				222,
				10,
				nil,
			},
		},
		expected: struct {
			status     bool
			raftValues []raftValue
		}{
			status:     true,          // is expected a true statement, drift is ok
			raftValues: []raftValue{}, // no failing members returned
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, one node is failing, is beyond the other
		irpm: []etcd.RaftIndexPerMember{
			{
				Server:    "etcd1",
				MemberId:  111,
				RaftIndex: 10,
				Err:       nil,
			},
			{
				Server:    "etcd2",
				MemberId:  222,
				RaftIndex: 12,
				Err:       nil,
			},
			{
				Server:    "etcd3",
				MemberId:  222,
				RaftIndex: 10,
				Err:       nil,
			},
		},
		expected: struct {
			status     bool
			raftValues []raftValue
		}{
			status: false, // is expected a false statement, drift is not ok
			raftValues: []raftValue{
				{12, "etcd2"},
			}, // etcd2 member is failing
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, one node is failing, is before the other.
		irpm: []etcd.RaftIndexPerMember{
			{
				"etcd1",
				111,
				10,
				nil,
			},
			{
				"etcd2",
				222,
				12,
				nil,
			},
			{
				"etcd3",
				222,
				12,
				nil,
			},
		},
		expected: struct {
			status     bool
			raftValues []raftValue
		}{
			status: false, // is expected a false statement, drift is not ok
			raftValues: []raftValue{
				{2, "etcd2"},
			}, // etcd2 member is failing
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, all members are failing
		irpm: []etcd.RaftIndexPerMember{
			{
				"etcd1",
				111,
				10,
				nil,
			},
			{
				"etcd2",
				222,
				12,
				nil,
			},
			{
				"etcd3",
				222,
				14,
				nil,
			},
		},
		expected: struct {
			status     bool
			raftValues []raftValue
		}{
			status: false, // is expected a false statement, drift is not ok
			raftValues: []raftValue{
				{10, "etcd1"}, {12, "etcd2"}, {14, "etcd3"},
			}, //  all member are failing
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, 2 members are failing
		irpm: []etcd.RaftIndexPerMember{
			{
				"etcd1",
				111,
				10,
				nil,
			},
			{
				"etcd2",
				222,
				11,
				nil,
			},
			{
				"etcd3",
				222,
				13,
				nil,
			},
			{
				"etcd4",
				222,
				10,
				nil,
			},
			{
				"etcd5",
				222,
				9,
				nil,
			},
		},
		expected: struct {
			status     bool
			raftValues []raftValue
		}{
			status: false, // is expected a false statement, drift is not ok
			raftValues: []raftValue{
				{13, "etcd3"}, {9, "etcd5"},
			}, // etcd3 and etcd5 are failing
		},
	},
}
