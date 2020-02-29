package main

import (
	"testing"
)

func Benchmark_sequencedArchivator(b *testing.B) {
	for i:=0; i< b.N; i++{
		sequencedArchivator([]string{
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
		})
	}
}

func Benchmark_concurrentArchivator(b *testing.B) {
	for i:=0; i< b.N; i++{
		concurrentArchivator([]string{
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
			"file1.txt",
			"file2.txt",
			"file3.txt",
		})
	}
}
