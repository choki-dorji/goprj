package model

import "myapp/myapp/datastore/postgres"

type Course struct {
	Cid        string `json:"cid"`
	Coursename string `json:"coursename"`
}

const (
	queryInsertCourse = "INSERT INTO course(cid, coursename) VALUES($1, $2) RETURNING cid;"
	queryGetCourses   = "SELECT cid, coursename FROM course;"
	queryGetCourse    = "SELECT cid, coursename FROM course WHERE cid=$1;"
	queryUpdateCourse = "UPDATE course SET coursename=$1 WHERE cid=$2 RETURNING cid;"
	queryDeleteCourse = "DELETE FROM course WHERE cid=$1 RETURNING cid;"
)

func (c *Course) Add() error {
	row := postgres.Db.QueryRow(queryInsertCourse, c.Cid, c.Coursename)
	return row.Scan(&c.Cid)
}

func GetAllCourses() ([]Course, error) {
	rows, err := postgres.Db.Query(queryGetCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var c Course
		if err := rows.Scan(&c.Cid, &c.Coursename); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

func GetCourse(cid string) (Course, error) {
	var c Course
	err := postgres.Db.QueryRow(queryGetCourse, cid).Scan(&c.Cid, &c.Coursename)
	return c, err
}

func (c *Course) Update() error {
	row := postgres.Db.QueryRow(queryUpdateCourse, c.Coursename, c.Cid)
	return row.Scan(&c.Cid)
}

func DeleteCourse(cid string) error {
	var id string
	return postgres.Db.QueryRow(queryDeleteCourse, cid).Scan(&id)
}
