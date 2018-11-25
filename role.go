package main

import "fmt"

//Role represents a role
type Role struct {
	ID     int
	Name   string
	Parent int
}

var roles []Role

//roleMap stores roleID and its direct descendants' roleIDs
var roleMap map[int][]int

var descendantsMap map[int][]int

//getRole return role by roleID
func getRole(roleID int) (Role, error) {
	if roleID > len(roles) {
		return Role{}, fmt.Errorf("roleID (%v) not found", roleID)
	}
	return roles[roleID-1], nil
}

//setRoles initialise global role info
func setRoles(someRoles []Role) error {
	roles = make([]Role, 0)
	roles = append(roles, someRoles...)
	roleMap = make(map[int][]int, 0)
	for _, role := range someRoles {
		// make sure parent exists except 0
		if role.Parent != 0 {
			if _, err := getRole(role.Parent); err != nil {
				return err
			}
		}
		roleMap[role.Parent] = append(roleMap[role.Parent], role.ID)
	}

	if err := buildDescendantsMap(); err != nil {
		return err
	}
	return nil
}

func buildDescendantsMap() error {
	descendantsMap = make(map[int][]int, 0)
	for i := 1; i <= len(roles); i++ {
		marked := map[int]bool{}
		results := []int{}
		queue := []int{}
		queue = append(queue, i)
		for len(queue) != 0 {
			head := queue[0]
			if marked[head] {
				return fmt.Errorf("Loop found at roleID (%v)", head)
			}
			results = append(results, roleMap[head]...)
			queue = append(queue, roleMap[head]...)
			marked[head] = true
			queue = queue[1:]
		}
		descendantsMap[i] = results
	}
	return nil
}

//getDescendants returns descendant role IDs given the parent role ID
func getDescendants(parent int) []int {
	results := []int{}
	queue := []int{}
	queue = append(queue, parent)
	for len(queue) != 0 {
		head := queue[0]
		results = append(results, roleMap[head]...)
		queue = append(queue, roleMap[head]...)
		queue = queue[1:]
	}
	return results
}

//getDescendants is same as getDescendants but return from cached search
func getDescendants2(parent int) []int {
	return descendantsMap[parent]
}
