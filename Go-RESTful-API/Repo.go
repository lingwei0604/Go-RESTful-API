package main

import (
	"fmt"
	"reflect"
)

func RepoCreateStudent(t Student) Base {

	fmt.Print(t)
	stmt, err := db.Prepare("INSERT Student SET Id=?,ClassNum=?,Score=? ON DUPLICATE KEY UPDATE Id=VALUES(Id)")

	checkErr(err)
	query, err := stmt.Exec(t.Id, t.ClassNum, t.Score)

	checkErr(err)
	v := reflect.ValueOf(query)
	fmt.Println(v)

	var info Base

	if err != nil {
		info.Code = "1"
		info.Msg = "fail"
		return info
	}
	info.Code = "0"
	info.Msg = "Success"
	return info
}

func RepoCreateclass(t class) Base {

	stmt, err := db.Prepare("INSERT class SET ClassNum=?,TeacherName=?")
	checkErr(err)
	query, err := stmt.Exec(t.ClassNum, t.TeacherName)
	checkErr(err)

	v := reflect.ValueOf(query)
	fmt.Println(v)

	var info Base
	if err != nil {
		info.Code = "1"
		info.Msg = "fail"
		return info
	}
	info.Code = "0"
	info.Msg = "Success"
	return info
}

func RepoGetSumScore(i string) []ScoreInfo{

	rows, err := db.Query("select SUM(Score) from Student where ClassNum = (select ClassNum from Student where Id = ?)",i)

	checkErr(err)
	v := reflect.ValueOf(rows)
	fmt.Println(v)

	fac := new(ScoreInfo)
	facs := []ScoreInfo{}
	for rows.Next() {
		err1 := rows.Scan(&fac.Score)
		if err1 != nil {
			panic(err1.Error())
		} else {
			facs = append(facs, *fac)
		}
	}

	return facs
}

func RepoGetTopScoreTeacher() []TeacherInfo {

	rows, err := db.Query("select TeacherName from class where ClassNum = (select ClassNum from Student order by Id,Score limit 1)")

	checkErr(err)
	v := reflect.ValueOf(rows)
	fmt.Println(v)

	fac := new(TeacherInfo)
	facs := []TeacherInfo{}
	for rows.Next() {
		err1 := rows.Scan(&fac.TeacherName)
		if err1 != nil {
			panic(err1.Error())
		} else {
			facs = append(facs, *fac)
		}
	}
	return facs
}
