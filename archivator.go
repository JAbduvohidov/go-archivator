package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	zipFormat    = ".zip"
	filesPath    = "./files/"
	archivesPath = "./archives/"
)

func main() {
	fileNames := os.Args[1:]
	if fileNames == nil {
		return
	}

	var method string

	fmt.Print("Please select archiving method(c-concurrent/s-sequences): ")
	_, err := fmt.Scan(&method)
	if err != nil {
		log.Printf("unable to get input: %v", err)
		return
	}
	switch strings.ToLower(strings.TrimSpace(method)) {
	case "c":
		concurrentArchivator(fileNames)
	case "s":
		sequencedArchivator(fileNames)
	default:
		fmt.Println("invalid method selected")
		return
	}

	time.Sleep(time.Second)
}

func sequencedArchivator(fileNames []string) {
	for _, fileName := range fileNames {
		archive(fileName)
	}
}

func concurrentArchivator(fileNames []string) {
	waitGroup := sync.WaitGroup{}
	for _, fileName := range fileNames {
		fName := fileName
		waitGroup.Add(1)
		go func(wg *sync.WaitGroup, fileName string) {
			defer func() {
				waitGroup.Done()
			}()
			archive(fName)
		}(&waitGroup, fileName)
	}
	waitGroup.Wait()
}

func archive(fileName string) {
	archiveFile, err := os.Create(archivesPath + fileName + zipFormat)
	if err != nil {
		return
	}

	defer func() {
		err = archiveFile.Close()
		if err != nil {
			return
		}
	}()

	writer := zip.NewWriter(archiveFile)

	defer func() {
		err = writer.Close()
		if err != nil {
			return
		}
	}()

	file, err := os.Open(filesPath + fileName)
	if err != nil {
		return
	}

	archive, err := writer.Create(fileName)
	if err != nil {
		return
	}
	_, err = io.Copy(archive, file)
	if err != nil {
		return
	}
}
