package main

import (
	"bufio"
	"fmt"
	"learngo/pipeline"
	"os"
)

func main() {
	// const filename = "large.in"
	// const n = 1000000
	const filename = "small.in"
	const n = 64
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		count++
		if count < 10 {
			fmt.Println(v)
		}
	}
}

func MergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(12, 13, 42, 111, 5)),
		pipeline.InMemSort(pipeline.ArraySource(12, 23, 24, 11, 15)))
	for num := range p {
		fmt.Println(num)
	}
}
