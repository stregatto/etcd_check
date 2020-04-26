package core

import "github.com/stregatto/etcd_check/etcd"

type expected struct {
	status  bool
	members []string
}

var testCasesRaftIndexPerMember = []struct {
	raftDrift int
	irpm      []etcd.RaftIndexPerMember
	expected  expected
}{
	{
		raftDrift: 0, // Test for 0 drift
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
				RaftIndex: 10,
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
			status  bool
			members []string
		}{
			status:  true,       // is expected a true statement, drift is ok
			members: []string{}, // no failing members returned
		},
	},
	{
		raftDrift: 1, // Test for 1 drift ok
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
				RaftIndex: 11,
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
			status  bool
			members []string
		}{
			status:  true,       // is expected a true statement, drift is ok
			members: []string{}, // no failing members returned
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
			status  bool
			members []string
		}{
			status:  false,             // is expected a false statement, drift is not ok
			members: []string{"etcd2"}, // etcd2 member is failing
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, one node is failing, is before the other.
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
				RaftIndex: 12,
				Err:       nil,
			},
		},
		expected: struct {
			status  bool
			members []string
		}{
			status:  false,             // is expected a false statement, drift is not ok
			members: []string{"etcd1"}, // etcd2 member is failing
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, all members are failing
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
				RaftIndex: 14,
				Err:       nil,
			},
		},
		expected: struct {
			status  bool
			members []string
		}{
			status:  false,                               // is expected a false statement, drift is not ok
			members: []string{"etcd1", "etcd2", "etcd3"}, //  all member are failing
		},
	},
	{
		raftDrift: 1, // Test for 1 drift, 2 members are failing
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
				RaftIndex: 11,
				Err:       nil,
			},
			{
				Server:    "etcd3",
				MemberId:  222,
				RaftIndex: 13,
				Err:       nil,
			},
			{
				Server:    "etcd4",
				MemberId:  222,
				RaftIndex: 10,
				Err:       nil,
			},
			{
				Server:    "etcd5",
				MemberId:  222,
				RaftIndex: 9,
				Err:       nil,
			},
		},
		expected: struct {
			status  bool
			members []string
		}{
			status:  false,                      // is expected a false statement, drift is not ok
			members: []string{"etcd3", "etcd5"}, //  are failing
		},
	},
}
