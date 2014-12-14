package main

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"

	"errors"
)

type User struct {
	Email string
	Name  string
}

func markAttendance(c appengine.Context) error {
	u := user.Current(c)
	if u == nil {
		return errors.New("Unauthorized")
	}
	user := &User{
		Email: u.Email,
		Name:  u.String(),
	}

	key := datastore.NewKey(c, "users", u.Email, 0, nil)
	_, err := datastore.Put(c, key, user)
	return err
}

func getAttendeesCount(c appengine.Context) int {
	count, err := datastore.NewQuery("users").KeysOnly().Count(c)
	if err != nil {
		panic(err)
	}
	return count
}
