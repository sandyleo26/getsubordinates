package main

import (
	"fmt"
)

//Role can be assigned to user
type Role struct {
	ID     int
	Name   string
	Parent int
}

//User represent a user
type User struct {
	ID   int
	Name string
	Role int
}

var roles []Role
var roleMap map[int][]int

var users []User
var roleIDUserIDMap map[int][]int

//setRoles
func setRoles(someRoles []Role) {
	roles = append(roles, someRoles...)
	roleMap = make(map[int][]int, 0)
	for _, role := range someRoles {
		roleMap[role.Parent] = append(roleMap[role.Parent], role.ID)
	}
}

//
func getRole(roleID int) Role {
	return roles[roleID-1]
}

//
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

//setUsers
func setUsers(someUsers []User) {
	users = append(users, someUsers...)
	roleIDUserIDMap = make(map[int][]int, 0)
	for _, user := range someUsers {
		roleIDUserIDMap[user.Role] = append(roleIDUserIDMap[user.Role], user.ID)
	}
}

//
func getUser(userID int) User {
	return users[userID-1]
}

//
func getUsersByID(roleID int) []User {
	userIDs := roleIDUserIDMap[roleID]
	users := []User{}
	for _, userID := range userIDs {
		users = append(users, getUser(userID))
	}
	return users
}

//getSubOrdinates
func getSubOrdinates(userID int) []User {
	user := getUser(userID)
	resutls := []User{}
	descendantRoleIDs := getDescendants(user.Role)
	for _, roleID := range descendantRoleIDs {
		resutls = append(resutls, getUsersByID(roleID)...)
	}
	return resutls
}

func main() {
	fmt.Println("Hello, World!")
}
