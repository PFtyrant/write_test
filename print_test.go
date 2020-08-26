package main

import (
	"fmt"
	"os"
	"testing"
)

var (
	testData = []string{"a", "b", "c", "d", "e"}
)

func BenchmarkPrintln(b *testing.B) {
	f, _ := os.OpenFile("first.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer f.Close()
	for i :=0; i < b.N; i++ {
		fmt.Fprintln(f, testData[0])
		fmt.Fprintln(f, testData[1])
		fmt.Fprintln(f, testData[2])
		fmt.Fprintln(f, testData[3])
		fmt.Fprintln(f, testData[4])
	}
}

func BenchmarkConcat(b *testing.B) {
	f, _ := os.OpenFile("second.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer f.Close()
	for i := 0; i < b.N; i++ {
		s := testData[0] + "\n"
		s += testData[1] + "\n"
		s += testData[2] + "\n"
		s += testData[3] + "\n"
		s += testData[4]
		fmt.Fprintln(f, s)
	}
}

func BenchmarkConcatOneLine(b *testing.B) {

	f, _ := os.OpenFile("third.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer f.Close()
	for i := 0; i < b.N; i++ {
		s := testData[0] + "\n" +
			testData[1] + "\n" +
			testData[2] + "\n" +
			testData[3] + "\n" +
			testData[4]
		fmt.Fprintln(f, s)
	}
}

