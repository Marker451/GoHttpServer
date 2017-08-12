package models

import "server/resource"
import "database/sql"
type StudentInfo struct {
	Id       int
	ClassNum int
	Score    int
}

func (s StudentInfo) Add() (err error) {
	stmtIns, err := resource.DatabaseMysql.Prepare("INSERT INTO student_info SET id=?,class_num=?,score=?")
	if err != nil {
		return
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(s.Id, s.ClassNum, s.Score)
	return
}

func (s StudentInfo) Update(data StudentInfo) (err error) { 
	stmtUpdate, err := resource.DatabaseMysql.Prepare("UPDATE  student_info SET id=?,class_num=?,score=? WHERE id=?")
	if err != nil {
		return
	}
	defer stmtUpdate.Close()
	_, err = stmtUpdate.Exec(data.Id, data.ClassNum, data.Score, s.Id)
	return
}

func (s StudentInfo) Get() (stu StudentInfo, empty bool, err error) { 
	err = resource.DatabaseMysql.QueryRow("SELECT * FROM student_info WHERE id=?", s.Id).Scan(&stu.Id, &stu.ClassNum,&stu.Score)
		if err ==sql.ErrNoRows{
			empty = true
	}
	return
}

func (s StudentInfo) GetScoreSum() (sumScore int, err error) {

	rows, err := resource.DatabaseMysql.Query("SELECT score FROM student_info WHERE class_num=?", s.ClassNum)
	if err != nil {
		return
	}
	tmpScore := 0
	for rows.Next() {
		err = rows.Scan(&tmpScore)
		sumScore += tmpScore
	}
	return
}

func (s StudentInfo) GetMaxScore() (max int, err error) {
	 err = resource.DatabaseMysql.QueryRow("SELECT MAX(score) FROM student_info").Scan(&max)
	if err != nil {
		return
	}
	return
}
func (s StudentInfo) GetStudentInfoByScore(score int) (stu StudentInfo, err error) {
	 err = resource.DatabaseMysql.QueryRow("SELECT * FROM student_info WHERE score=? ORDER BY id LIMIT 1").Scan(&s.Id, &s.ClassNum, &s.Score)
	if err != nil {
		return
	}
	return
}
