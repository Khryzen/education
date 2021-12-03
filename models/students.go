package models

import "github.com/uadmin/uadmin"

type Students struct {
	uadmin.Model
	Firstname  string
	Middlename string
	Lastname   string
	Email      string
	Cnumber    string
	Address    string
	Course     string
	YearLevel  string
}

func (s *Students) String() string {
	return s.Firstname + " " + s.Lastname
}

func (s *Students) Save() {
	uadmin.Save(s)
}
