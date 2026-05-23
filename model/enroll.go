package model

import (
	"myapp/dataStore/postgres"
)

type Enroll struct {
	StdId         int64  `json:"stdid"`
	CourseID      string `json:"cid"`
	Date_Enrolled string `json:"date"`
}

const queryEnrollStd = "insert into enroll(std_id,course_id, date_enrolled) values ($1,$2,$3);"

func (e *Enroll) EnrollStud() error {
	//fmt.Println(e)
	_, err := postgres.Db.Exec(queryEnrollStd, e.StdId, e.CourseID, e.Date_Enrolled)
	return err
}
