package models

import (
	"github.com/uadmin/uadmin"
)

type College struct {
	uadmin.Model
	CollegeName        string
	CollegeDescription string
	FieldMajor         string
	CollegeEmail       string
	CollegeNumber      string
	Courses            string
	CourseCount        int
	StudentsEnrolled   int
}

func (c *College) String() string {
	return c.CollegeName
}

func (c *College) Save() {
	uadmin.Save(c)
}
