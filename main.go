package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Lession struct {
	Slug  string `json:"slug`
	Title string `json:"title`
	Index int    `json:"index`
	Hash  string `json:"hash`
}

type CourseInfo struct {
	Title      string             `json:"title"`
	LessonData map[string]Lession `json:"lessonData"`
}

const downloadFolder = "download"
const courseInfoFileName = "info.json"
const courseListFile = "courses.json"
const fileExt = ".webm"

func main() {
	courses := map[string]string{}

	if err := json.Unmarshal(readCoursesFile(), &courses); err != nil {
		panic(err)
	}

	for courseName, courseValue := range courses {
		if courseValue != "true" {
			continue
		}
		log(courseName)

		courseFolder := downloadFolder + "/" + courseName
		createFolder(courseFolder)

		awsFileKey := courseName + "/" + courseInfoFileName
		infoPath := courseFolder + "/" + courseInfoFileName
		log("download info file", awsFileKey, infoPath)
		download(awsFileKey, infoPath)

		info := CourseInfo{}
		if err := json.Unmarshal(readFile(infoPath), &info); err != nil {
			panic(err)
		}
		log("read info file", infoPath)

		downloadCount := 0
		totalCount := len(info.LessonData)
		for hash, data := range info.LessonData {
			awsFileKey = courseName + "/" + hash + fileExt
			title := strings.ReplaceAll(data.Title, "/", "-")
			filePath := courseFolder + "/" + fmt.Sprintf("%02d - ", data.Index) + title + fileExt
			log("download video file", awsFileKey, filePath)
			download(awsFileKey, filePath)

			downloadCount = downloadCount + 1
			courses[courseName] = fmt.Sprintf("%d/%d", downloadCount, totalCount)
			updateCoursesFile(courses)
		}

		courses[courseName] = courses[courseName] + " finished"
		updateCoursesFile(courses)
	}
}
