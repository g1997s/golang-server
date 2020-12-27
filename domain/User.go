package domain

import (
	_ "../repository"
	"fmt"
	"time"
)

type User struct {
	Name     string
	Password string
	DOB      string
	Date     string
}

func NewUser(name string, pass string, dob string) (*User, error) {
	date := time.Now().Format("January 2, 2006")
	if name == "" || pass == "" || dob == "" {
		return nil, fmt.Errorf("required field can not be empty")
	}
	return &User{
		Name:     name,
		Password: pass,
		DOB:      dob,
		Date:     date,
	}, nil
}
