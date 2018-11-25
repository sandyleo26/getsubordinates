package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setRoles(t *testing.T) {
	err := setRoles([]Role{
		Role{
			Parent: 0,
		},
		Role{
			Parent: 3,
		},
	})
	assert.Equal(t, "roleID (3) not found", err.Error())
}

func Test_getDescendants(t *testing.T) {
	roleMap = map[int][]int{
		0: []int{1},
		1: []int{2},
		2: []int{3},
		3: []int{4},
	}

	assert.ElementsMatch(t, []int{1, 2, 3, 4}, getDescendants(0))
	assert.ElementsMatch(t, []int{2, 3, 4}, getDescendants(1))
	assert.ElementsMatch(t, []int{3, 4}, getDescendants(2))
	assert.ElementsMatch(t, []int{4}, getDescendants(3))
	assert.ElementsMatch(t, []int{}, getDescendants(4))
}
