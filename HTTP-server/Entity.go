package main

type Student struct {
	Id 		      string    `json:"id"`
	ClassNum      string    `json:"classnum"`
	Score         string    `json:"score"`
}

type class struct {
	ClassNum      string    `json:"classnum"`
	TeacherName   string    `json:"teachername"`
}

type ScoreInfo struct {
	Score  string
}

type TeacherInfo struct {
	TeacherName string
}

type Base struct {
	Code   string
	Msg    string
}

type Students [] Student
type classes  [] class
