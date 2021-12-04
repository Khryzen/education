package models

import (
	"github.com/uadmin/uadmin"
)

type Subjects struct {
	uadmin.Model
	SubjectName string
	Units       int
}

func (s *Subjects) String() string {
	return s.SubjectName
}

func (s *Subjects) Save() {
	uadmin.Save(s)
}
