package model

import "myapp/dataStore/postgres"

type Student struct {
	StdId int64 `json:"stdid"`
	FirstName string `json:"fname"`
	LastName string `json:"lname"`
	Email string `json:"email"`
}

type Course struct {
	Cid string `json:"cid"`
	CourseName string `json:"coursename"`
}

//query for student table
const queryinsertUser = "INSERT INTO student(stdid, firstname, lastname, email) VALUES ($1, $2, $3, $4);"
const queryGetUser = "SELECT * FROM student WHERE stdid=$1;"
const queryUpdateUser = "UPDATE student SET stdid=$1, firstname=$2, lastname=$3, email=$4 WHERE stdid=$5 RETURNING stdid;"
const queryDeleteUser = "DELETE FROM student WHERE stdid=$1 RETURNING stdid;"

//query for course table
const queryinsertCourse = "INSERT INTO course (cid, coursename) VALUES ($1, $2);"
const queryGetCourse = "SELECT * FROM course WHERE cid=$1;"
const queryUpdateCourse = "UPDATE course SET cid=$1, coursename=$2 WHERE cid=$3 RETURNING cid;"
const quertDeleteCourse = "DELETE FROM course WHERE cid=$1 RETURNING cid;"
const queryAllCourse = "SELECT * FROM course;"


//responsible for interacting with database
//stud -> s
func (s *Student) Create() error {
		_, err := postgres.Db.Exec(queryinsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
		return err 
}

func (s *Student) Read() error {
		row := postgres.Db.QueryRow(queryGetUser, s.StdId)
		return row.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
}

func(s *Student) Update(oldID int64) error {
		row := postgres.Db.QueryRow(queryUpdateUser, s.StdId, s.FirstName, s.LastName, s.Email, oldID)
		return row.Scan(&s.StdId)
}

func (s *Student) Delete() error {
	if err := postgres.Db.QueryRow(queryDeleteUser, s.StdId).Scan(&s.StdId); err != nil {
		return err
	}
	return nil
}

func GetAllStudents() ([]Student, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM student;")
	if getErr != nil {
		return nil, getErr
	}
	//create a slice of type student
	students := []Student{}

	for rows.Next() {
		var s Student
		dbErr := rows.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
		if dbErr != nil {
			return nil, dbErr
		}

		students = append(students, s)
	}
	rows.Close()
	return students, nil
}



// Course
func (c *Course) Create() error {
	_, err := postgres.Db.Exec(queryinsertCourse, c.Cid, c.CourseName)
	return err
}

func (c *Course) Read() error {
	row := postgres.Db.QueryRow(queryGetCourse, c.Cid)
	return row.Scan(&c.Cid, &c.CourseName)
}

func (c *Course) Update(oldCid string) error {
	err := postgres.Db.QueryRow(queryUpdateCourse, c.Cid, c.CourseName, oldCid).Scan(&c.Cid)
	return err
}

func (c *Course) Delete() error {
	_, err := postgres.Db.Exec(quertDeleteCourse, c.Cid)
	return err
}

func GetAllCourse() ([]Course, error) {
	rows, err := postgres.Db.Query(queryAllCourse)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []Course{}
	for rows.Next() {
		var c Course
		if err := rows.Scan(&c.Cid, &c.CourseName); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}