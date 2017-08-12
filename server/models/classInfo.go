package models

import "server/resource"
import "database/sql"

type ClassInfo struct {
	ClassNum int
	Teacher  string
}

func (c ClassInfo) Add() (err error) {
	stmtIns, err := resource.DatabaseMysql.Prepare("INSERT INTO class_info SET Class_num=?, teacher=?")
	if err != nil {
		return
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(c.ClassNum, c.Teacher)
	return
}

func (c ClassInfo) Update(data ClassInfo) (err error) { //by id
	stmtUpdate, err := resource.DatabaseMysql.Prepare("UPDATE  class_info SET class_num=?,teacher=? WHERE class_num=?")
	if err != nil {
		return
	}
	defer stmtUpdate.Close()
	_, err = stmtUpdate.Exec(data.ClassNum, data.Teacher, c.ClassNum)
	return
}

func (c ClassInfo) Get() (class ClassInfo, empty bool, err error) { //by id
	 err = resource.DatabaseMysql.QueryRow("SELECT * FROM class_info WHERE class_num=?", c.ClassNum).Scan(&class.ClassNum, &class.Teacher)
	if err == sql.ErrNoRows{
		empty =  true
	}
	return
}
