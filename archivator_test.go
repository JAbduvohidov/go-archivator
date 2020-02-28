package main

import (
	"io/ioutil"
	"testing"
	"time"
)

var fileNames = []string{"file1.txt", "file2.txt", "file3.txt"}

func Test_sequencedArchivator(t *testing.T) {
	sequencedArchivator(fileNames)

	time.Sleep(time.Second)
	files, err := ioutil.ReadDir(archivesPath)
	if err != nil {
		t.Fatalf("unable to read directory: %v", err)
	}
	if len(files) != 3 {
		t.Fatalf("incorrect number of files, expected: 3, got: %d", len(files))
	}
	if files[0].Name() != "file1.txt.zip" {
		t.Fatalf("first element must be file1.txt.zip")
	}
	if files[1].Name() != "file2.txt.zip" {
		t.Fatalf("second element must be file1.txt.zip")
	}
	if files[2].Name() != "file3.txt.zip" {
		t.Fatalf("third element must be file1.txt.zip")
	}
}

func Test_concurrentArchivator(t *testing.T) {
	concurrentArchivator(fileNames)

	time.Sleep(time.Second)
	files, err := ioutil.ReadDir(archivesPath)
	if err != nil {
		t.Fatalf("unable to read directory: %v", err)
	}
	if len(files) != 3 {
		t.Fatalf("incorrect number of files, expected: 3, got: %d", len(files))
	}
	if files[0].Name() != "file1.txt.zip" {
		t.Fatalf("first element must be file1.txt.zip")
	}
	if files[1].Name() != "file2.txt.zip" {
		t.Fatalf("second element must be file1.txt.zip")
	}
	if files[2].Name() != "file3.txt.zip" {
		t.Fatalf("third element must be file1.txt.zip")
	}
}
