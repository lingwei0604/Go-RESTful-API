package main

import (
	"fmt"
	"reflect"
)
/*
 achieve repository
 mysql、mongodb etc
*/
func Repo_AddStudent(t Student) Base {
	fmt.Print(t)
	stmt, err := db.Prepare("INSERT student SET id=?,class_num=?,score=? ON DUPLICATE KEY UPDATE id=VALUES(id)")

	checkErr(err)
	query, err := stmt.Exec(t.Id, t.ClassNum, t.Score)

	checkErr(err)
	v := reflect.ValueOf(query)
	fmt.Println(v)

	var status Base

	if err != nil {
		status.code = "0"
		status.msg = "fail"
		return status
	}
	status.code = "1"
	status.msg = "Success"
	return status
}

func Repo_AddClass(t Class) Base {
	stmt, err := db.Prepare("INSERT class SET class_num=?,teacher_name=?")
	checkErr(err)
	query, err := stmt.Exec(t.ClassNum, t.TeacherName)
	checkErr(err)

	v := reflect.ValueOf(query)
	fmt.Println(v)

	var status Base
	if err != nil {
		status.code = "1"
		status.msg = "fail"
		return status
	}
	status.code = "0"
	status.msg = "Success"
	return status
}

func Repo_GetTotalScore(i string) []Score{
	rows, err := db.Query("select SUM(score) from student where class_num = (select class_num from student where id = ?)",i)

	checkErr(err)
	v := reflect.ValueOf(rows)
	fmt.Println(v)

	tmp := new(Score)
	tmps := []Score{}
	for rows.Next() {
		err1 := rows.Scan(&tmp.totalScore)
		if err1 != nil {
			panic(err1.Error())
		} else {
			tmps = append(tmps, *tmp)
		}
	}
	return tmps
}

func Repo_GetTeacherFromScore() []Teacher {
	rows, err := db.Query("select teacher_name from class where class_num = (select class_num from student order by id,score limit 1)")

	checkErr(err)
	v := reflect.ValueOf(rows)
	fmt.Println(v)

	tmp := new(Teacher)
	tmps := []Teacher{}
	for rows.Next() {
		err1 := rows.Scan(&tmp.teacherName)
		if err1 != nil {
			panic(err1.Error())
		} else {
			tmps = append(tmps, *tmp)
		}
	}
	return tmps
}
