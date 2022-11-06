package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	db      map[int]*Student
	counter int
)

type Student struct {
	name  string
	age   int
	grade int
}

func newStudent() *Student {
	return new(Student)
}

func (std *Student) getName() string {
	return std.name
}

func (std *Student) getAge() int {
	return std.age
}

func (std *Student) getGrade() int {
	return std.grade
}

func (std *Student) putName(name string) {
	std.name = name
}

func (std *Student) putAge(age int) {
	std.age = age
}

func (std *Student) putGrade(grade int) {
	std.grade = grade
}

func handleCtrlZ(c chan os.Signal) {
	sig := <-c
	fmt.Println("\ncatch signal: ", sig)

	for _, s := range db {
		fmt.Println("students from storage:")
		fmt.Println(s.getName(), s.getAge(), s.getGrade())
	}
}

func main() {
	var (
		inName  string
		inAge   int
		inGrade int
	)

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTSTP)
	go handleCtrlZ(c)

	db = make(map[int]*Student)

	for {
		_, err := fmt.Scanf("%s %d %d", &inName, &inAge, &inGrade)
		if err != nil {
			log.Fatal(err)
		}

		counter++

		ns := newStudent()
		ns.putName(inName)
		ns.putAge(inAge)
		ns.putGrade(inGrade)

		db[counter] = ns
	}
}
