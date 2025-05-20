package model

import "myapp/datastore/postgres"

type Enroll struct {
	StdId         int64  `json:"stdid"`
	CourseID      string `json:"cid"`
	Date_Enrolled string `json:"date"`
}

const queryEnrollStd = "INSERT INTO enroll(std_id, course_id, date_enrolled) VALUES($1, $2, $3) RETURNING std_id;"

const (
	queryGetEnroll     = "SELECT std_id, course_id, date_enrolled FROM enroll WHERE std_id=$1 AND course_id=$2;"
	queryGetAllEnrolls = "SELECT std_id, course_id, date_enrolled FROM enroll;"
	queryDeleteEnroll  = "DELETE FROM enroll WHERE std_id=$1 AND course_id=$2 RETURNING std_id;"
	queryUpdateEnroll  = "UPDATE enroll SET course_id=$1, date_enrolled=$2 WHERE std_id=$3 RETURNING std_id;"
)

func (e *Enroll) EnrollStud() error {
	row := postgres.Db.QueryRow(queryEnrollStd, e.StdId, e.CourseID, e.Date_Enrolled)
	return row.Scan(&e.StdId)
}

func GetEnroll(stdID int64, courseID string) (Enroll, error) {
	var e Enroll
	err := postgres.Db.QueryRow(queryGetEnroll, stdID, courseID).Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
	return e, err
}

func GetAllEnrolls() ([]Enroll, error) {
	rows, err := postgres.Db.Query(queryGetAllEnrolls)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrolls []Enroll
	for rows.Next() {
		var e Enroll
		if err := rows.Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled); err != nil {
			return nil, err
		}
		enrolls = append(enrolls, e)
	}
	return enrolls, nil
}

func DeleteEnroll(stdID int64, courseID string) error {
	var id int64
	return postgres.Db.QueryRow(queryDeleteEnroll, stdID, courseID).Scan(&id)
}

func (e *Enroll) UpdateEnroll(oldStdID int64) error {
	row := postgres.Db.QueryRow(queryUpdateEnroll, e.CourseID, e.Date_Enrolled, oldStdID)
	return row.Scan(&e.StdId)
}
