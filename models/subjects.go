package models

import (
	"github.com/uadmin/uadmin"
)

type Subjects struct {
	uadmin.Model
	Course      Course
	CourseID    uint
	SubjectName string
	Units       int
}

func (s *Subjects) String() string {
	return s.SubjectName
}

func (s *Subjects) Save() {
	uadmin.Save(s)
}
