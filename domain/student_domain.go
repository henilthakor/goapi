package domain

import "errors"

type Student struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Mark float32 `json:"mark"`
}

func GetStudentDetail(id int) (*Student, error) {
	s := Students[id]
	if s != nil {
		return s, nil
	}
	return nil, errors.New("Student Not Found")
}

func AddNewStudentDetail(s *Student) {
	var id int
	for id, _ = range Students {
	}
	Students[id+1] = s
	return
}

func UpdateStudentDetail(id int, s *Student) error {
	stud := Students[id]
	if stud != nil {
		return errors.New("Student Not Found")
	}
	Students[id] = s
	return nil
}

func DeleteStudentDetail(id int) error {
	s := Students[id]
	if s != nil {
		return errors.New("Student Not Found")
	}
	delete(Students, id)
	return nil
}
