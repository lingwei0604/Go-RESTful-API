package main

//Immutable model
type Student struct {
	Id 		      string    `json:"id"`
	ClassNum      string    `json:"classnum"`
	Score         string    `json:"score"`
}

type Class struct {
	ClassNum      string    `json:"classnum"`
	TeacherName   string    `json:"teachername"`
}

type Score struct {
	totalScore  string
}

type Teacher struct {
	teacherName string
}

type Base struct {
	code   string
	msg    string
}

type Students [] Student
type classes  [] Class
