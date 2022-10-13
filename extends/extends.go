package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

type People struct {
	Name string
}

type Person struct {
	Name string
	high int
}

func (stu *Student) printScore() {
	fmt.Println(stu.Score)
}

type Pupil struct {
	Student
}

func (pupil *Pupil) goSchool() {
	fmt.Println("pupil go to school")
}

type UnderGraduate struct {
	Student
	People
	person Person
}

func (graduate *UnderGraduate) goSchool() {
	fmt.Println("graduate go to school")
}

func main() {
	pubil1 := &Pupil{}
	pubil1.Name = "oliver"
	pubil1.Age = 12
	pubil1.Score = 90
	pubil1.printScore()
	pubil1.goSchool()

	underGraduate1 := UnderGraduate{}
	underGraduate1.People.Name = "bob"
	underGraduate1.Age = 20
	underGraduate1.Score = 60
	underGraduate1.person.high = 130
	underGraduate1.printScore()
	underGraduate1.goSchool()

	underGraduate2 := UnderGraduate{Student{
		Name:  "oliver",
		Age:   12,
		Score: 12,
	}, People{
		Name: "bob",
	}, Person{
		Name: "alice",
	}}
	fmt.Println(underGraduate2)

	littleMonkey := LittleMonkey{
		Monkey{
			Name: "oliver",
		},
	}
	littleMonkey.climbing()
	littleMonkey.Flying()
	littleMonkey.Swimming()
}
