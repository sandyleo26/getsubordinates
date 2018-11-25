package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getSubOrdinates(t *testing.T) {
	roles := []Role{
		Role{
			ID:     1,
			Name:   "System Administrator",
			Parent: 0,
		},
		Role{
			ID:     2,
			Name:   "Location Manager",
			Parent: 1,
		},
		Role{
			ID:     3,
			Name:   "Supervisor",
			Parent: 2,
		},
		Role{
			ID:     4,
			Name:   "Employee",
			Parent: 3,
		},
		Role{
			ID:     5,
			Name:   "Trainer",
			Parent: 3,
		},
	}

	users := []User{
		User{
			ID:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		User{
			ID:   2,
			Name: "Emily Employee",
			Role: 4,
		},
		User{
			ID:   3,
			Name: "Sam Supervisor",
			Role: 3,
		},
		User{
			ID:   4,
			Name: "Mary Manager",
			Role: 2,
		},
		User{
			ID:   5,
			Name: "Steve Trainer",
			Role: 5,
		},
	}

	setRoles(roles)
	setUsers(users)
	assert.Equal(t, []User{users[1], users[4]}, getSubOrdinates(3))
	assert.ElementsMatch(t, users[1:], getSubOrdinates(1))
}
