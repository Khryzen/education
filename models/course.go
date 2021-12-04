package models

import (
	"github.com/uadmin/uadmin"
)

type Course struct {
	uadmin.Model
	CourseName        string
	CourseDescription string
	CourseType        string
	Semesters         int
	TotalUnits        int
}

func (c *Course) String() string {
	return c.CourseName
}

func (c *Course) Save() {
	uadmin.Save(c)
}
