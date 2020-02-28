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
	logFile, err := os.OpenFile("archivator.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 666)
	if err != nil {
		log.Fatalf("unable to open file or create log file, %v", err)
	}

	log.SetOutput(logFile)

	log.Print("getting command line arguments")
	fileNames := os.Args[1:]
	if fileNames == nil {
		log.Print("nothing to write list of file names is empty")
		return
	}

	var method string

	fmt.Print("Please select archiving method(c-concurrent/s-sequences): ")
	_, err = fmt.Scan(&method)
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
	mutex := &sync.Mutex{}
	for _, fileName := range fileNames {
		fName := fileName
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			archive(fName)
		}()
	}
}

func concurrentArchivator(fileNames []string) {
	for _, fileName := range fileNames {
		go archive(fileName)
	}
}

func archive(fileName string) {
	log.Print("writing started")

	log.Printf("trying to create file %s", fileName+zipFormat)
	archiveFile, err := os.Create(archivesPath + fileName + zipFormat)
	if err != nil {
		log.Printf("unable to create %s archive, %v", fileName, err)
		return
	}
	log.Print("file created")

	defer func() {
		log.Print("closing archive")
		err = archiveFile.Close()
		if err != nil {
			log.Print("unable to close archive")
			return
		}
		log.Print("archive closed")
	}()

	writer := zip.NewWriter(archiveFile)

	defer func() {
		log.Print("closing writer")
		err = writer.Close()
		if err != nil {
			log.Print("unable to close writer")
			return
		}
		log.Print("writer closed")
	}()

	log.Printf("opening file: %s", fileName)
	file, err := os.Open(filesPath + fileName)
	if err != nil {
		log.Print("unable to open file")
		return
	}
	log.Print("file opened")

	log.Print("creating archive writer")
	archive, err := writer.Create(fileName)
	if err != nil {
		log.Printf("unable to create archive writer, %v", err)
		return
	}
	log.Print("archive writer created")

	log.Print("copying from file to archive")
	_, err = io.Copy(archive, file)
	if err != nil {
		log.Printf("unable to copy from file to archive, %v", err)
		return
	}
	log.Print("successfully copied")

	log.Print("writing finished")
}
