package main

import (
	"fmt"
)

//User represent a user
type User struct {
	ID   int
	Name string
	Role int
}

var users []User
var roleIDUserIDMap map[int][]int

//setUsers initialise global user info
func setUsers(someUsers []User) {
	users = append(users, someUsers...)
	roleIDUserIDMap = make(map[int][]int, 0)
	for _, user := range someUsers {
		roleIDUserIDMap[user.Role] = append(roleIDUserIDMap[user.Role], user.ID)
	}
}

//getUser returns user info given userID
func getUser(userID int) (User, error) {
	if userID >= len(users) {
		return User{}, fmt.Errorf("userID (%v) not found", userID)
	}
	return users[userID-1], nil
}

//getUsersByRoleID returns users with roleID
func getUsersByRoleID(roleID int) ([]User, error) {
	userIDs := roleIDUserIDMap[roleID]
	results := []User{}
	for _, userID := range userIDs {
		user, err := getUser(userID)
		if err != nil {
			return nil, err
		}
		results = append(results, user)
	}
	return results, nil
}

//getSubOrdinates
func getSubOrdinates(userID int) ([]User, error) {
	user, err := getUser(userID)
	if err != nil {
		return nil, err
	}

	resutls := []User{}
	descendantRoleIDs := getDescendants(user.Role)
	for _, roleID := range descendantRoleIDs {
		found, err := getUsersByRoleID(roleID)
		if err != nil {
			return nil, err
		}
		resutls = append(resutls, found...)
	}
	return resutls, nil
}
