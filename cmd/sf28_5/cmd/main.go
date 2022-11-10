package main

import (
	"fmt"
	"io"
	"os"
	"sf28_5/pkg/storage"
	"sf28_5/pkg/student"
)

func main() {
	var (
		inName  string
		inAge   int
		inGrade int
	)

	strg := storage.BuildStorage()

	for {
		_, err := fmt.Scanf("%s %d %d", &inName, &inAge, &inGrade)
		if err == io.EOF {
			strg.PrintAll()
			os.Exit(0)
		}

		ns := student.NewStudent()
		ns.PutName(inName)
		ns.PutAge(inAge)
		ns.PutGrade(inGrade)

		strg.Add(ns)
	}

}
