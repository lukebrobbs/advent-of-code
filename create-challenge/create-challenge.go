package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
)

var year = flag.Int("year", 2024, "Advent of code year to run")
var day = flag.Int("day", 1, "Advent of code day to run")

type Day struct {
	Day int
}

type FileDetail struct {
	template, name string
}

func main() {
	flag.Parse()
	d := Day{*day}
	y := *year
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir := fmt.Sprintf("%s/challenges/%d/day-%d", cwd, y, d.Day)
	paths := []FileDetail{
		{template: fmt.Sprintf("%s/create-challenge/challenge.tmpl", cwd), name: fmt.Sprintf("%s/day-%d.go", dir, d.Day)},
		{template: fmt.Sprintf("%s/create-challenge/challenge-test.tmpl", cwd), name: fmt.Sprintf("%s/day-%d_test.go", dir, d.Day)},
	}
	os.MkdirAll(dir, os.ModePerm)
	os.WriteFile(fmt.Sprintf("%s/README.md", dir), []byte(""), os.ModePerm)

	for _, file := range paths {

		t := template.Must(template.New(path.Base(file.template)).ParseFiles(file.template))

		f, err := os.Create(file.name)
		if err != nil {
			log.Println("create file: ", err)
			return
		}

		err = t.Execute(f, d)
		if err != nil {
			panic(err)
		}
	}

}
