package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func writeToFile(fileName string, text string, append bool) {
	flag := os.O_CREATE | os.O_WRONLY
	if append == true {
		flag = os.O_APPEND | flag
	} else {
		flag = os.O_TRUNC | flag
	}
	f, err := os.OpenFile(fileName, flag, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(text + "\n"); err != nil {
		fmt.Println(err)
	}
}

func log(input ...interface{}) {
	dt := time.Now()
	text := fmt.Sprintf("%s - %v", dt.Format("01-02-2006 15:04:05"), input)
	fmt.Println(text)
	writeToFile("main.log", string(text), true)
}

func readFile(path string) []byte {
	b, _ := ioutil.ReadFile(path)
	return b
}

func readCoursesFile() []byte {
	return readFile(courseListFile)
}

func updateCoursesFile(courses map[string]string) {
	text, _ := json.Marshal(courses)
	writeToFile(courseListFile, string(text), false)
}

func createFolder(path string) {
	os.MkdirAll(path, 0777)
}
