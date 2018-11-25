package main

//Role represents a role
type Role struct {
	ID     int
	Name   string
	Parent int
}

var roleMap map[int][]int

//setRoles initialise global role info
func setRoles(someRoles []Role) {
	roleMap = make(map[int][]int, 0)
	for _, role := range someRoles {
		roleMap[role.Parent] = append(roleMap[role.Parent], role.ID)
	}
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
